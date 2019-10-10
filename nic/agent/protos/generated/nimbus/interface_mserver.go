// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: interface.proto
*/

package nimbus

import (
	"context"
	"errors"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	hdr "github.com/pensando/sw/venice/utils/histogram"
	"github.com/pensando/sw/venice/utils/log"
	memdb "github.com/pensando/sw/venice/utils/memdb"
	"github.com/pensando/sw/venice/utils/netutils"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// FindInterface finds an Interface by object meta
func (ms *MbusServer) FindInterface(objmeta *api.ObjectMeta) (*netproto.Interface, error) {
	// find the object
	obj, err := ms.memDB.FindObject("Interface", objmeta)
	if err != nil {
		return nil, err
	}

	return InterfaceFromObj(obj)
}

// ListInterfaces lists all Interfaces in the mbus
func (ms *MbusServer) ListInterfaces(ctx context.Context, filterFn func(memdb.Object) bool) ([]*netproto.Interface, error) {
	var objlist []*netproto.Interface

	// walk all objects
	objs := ms.memDB.ListObjects("Interface", filterFn)
	for _, oo := range objs {
		obj, err := InterfaceFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// InterfaceStatusReactor is the reactor interface implemented by controllers
type InterfaceStatusReactor interface {
	OnInterfaceCreateReq(nodeID string, objinfo *netproto.Interface) error
	OnInterfaceUpdateReq(nodeID string, objinfo *netproto.Interface) error
	OnInterfaceDeleteReq(nodeID string, objinfo *netproto.Interface) error
	OnInterfaceOperUpdate(nodeID string, objinfo *netproto.Interface) error
	OnInterfaceOperDelete(nodeID string, objinfo *netproto.Interface) error
	GetWatchFilter(kind string, ometa *api.ObjectMeta) func(memdb.Object) bool
}

type InterfaceNodeStatus struct {
	nodeID        string
	watcher       *memdb.Watcher
	opSentStatus  map[api.EventType]*EventStatus
	opAckedStatus map[api.EventType]*EventStatus
}

// InterfaceTopic is the Interface topic on message bus
type InterfaceTopic struct {
	sync.Mutex
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor InterfaceStatusReactor // status event reactor
	nodeStatus    map[string]*InterfaceNodeStatus
}

// AddInterfaceTopic returns a network RPC server
func AddInterfaceTopic(server *MbusServer, reactor InterfaceStatusReactor) (*InterfaceTopic, error) {
	// RPC handler instance
	handler := InterfaceTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
		nodeStatus:    make(map[string]*InterfaceNodeStatus),
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterInterfaceApiServer(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

func (eh *InterfaceTopic) registerWatcher(nodeID string, watcher *memdb.Watcher) {
	eh.Lock()
	defer eh.Unlock()

	eh.nodeStatus[nodeID] = &InterfaceNodeStatus{nodeID: nodeID, watcher: watcher}
	eh.nodeStatus[nodeID].opSentStatus = make(map[api.EventType]*EventStatus)
	eh.nodeStatus[nodeID].opAckedStatus = make(map[api.EventType]*EventStatus)
}

func (eh *InterfaceTopic) unRegisterWatcher(nodeID string) {
	eh.Lock()
	defer eh.Unlock()

	delete(eh.nodeStatus, nodeID)
}

//update recv object status
func (eh *InterfaceTopic) updateAckedObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

	eh.Lock()
	defer eh.Unlock()
	var evStatus *EventStatus

	nodeStatus, ok := eh.nodeStatus[nodeID]
	if !ok {
		//Watcher already unregistered.
		return
	}

	evStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		nodeStatus.opAckedStatus[event] = &EventStatus{}
		evStatus = nodeStatus.opAckedStatus[event]
	}

	if LatencyMeasurementEnabled {
		rcvdTime, _ := objMeta.ModTime.Time()
		sendTime, _ := objMeta.CreationTime.Time()
		delta := rcvdTime.Sub(sendTime)

		hdr.Record(nodeID+"_"+"Interface", delta)
		hdr.Record("Interface", delta)
		hdr.Record(nodeID, delta)
	}

	new, _ := strconv.Atoi(objMeta.ResourceVersion)
	//for create/delete keep track of last one sent to, this may not be full proof
	//  Create could be processed asynchoronusly by client and can come out of order.
	//  For now should be ok as at least we make sure all messages are processed.
	//For update keep track of only last one as nimbus client periodically pulls
	if evStatus.LastObjectMeta != nil {
		current, _ := strconv.Atoi(evStatus.LastObjectMeta.ResourceVersion)
		if current > new {
			return
		}
	}
	evStatus.LastObjectMeta = objMeta
}

//update recv object status
func (eh *InterfaceTopic) updateSentObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

	eh.Lock()
	defer eh.Unlock()
	var evStatus *EventStatus

	nodeStatus, ok := eh.nodeStatus[nodeID]
	if !ok {
		//Watcher already unregistered.
		return
	}

	evStatus, ok = nodeStatus.opSentStatus[event]
	if !ok {
		nodeStatus.opSentStatus[event] = &EventStatus{}
		evStatus = nodeStatus.opSentStatus[event]
	}

	new, _ := strconv.Atoi(objMeta.ResourceVersion)
	//for create/delete keep track of last one sent to, this may not be full proof
	//  Create could be processed asynchoronusly by client and can come out of order.
	//  For now should be ok as at least we make sure all messages are processed.
	//For update keep track of only last one as nimbus client periodically pulls
	if evStatus.LastObjectMeta != nil {
		current, _ := strconv.Atoi(evStatus.LastObjectMeta.ResourceVersion)
		if current > new {
			return
		}
	}
	evStatus.LastObjectMeta = objMeta
}

//update recv object status
func (eh *InterfaceTopic) WatcherInConfigSync(nodeID string, event api.EventType) bool {

	var ok bool
	var evStatus *EventStatus
	var evAckStatus *EventStatus

	eh.Lock()
	defer eh.Unlock()

	nodeStatus, ok := eh.nodeStatus[nodeID]
	if !ok {
		return true
	}

	evStatus, ok = nodeStatus.opSentStatus[event]
	if !ok {
		//nothing sent, so insync
		return true
	}

	//In-flight object still exists
	if len(nodeStatus.watcher.Channel) != 0 {
		log.Infof("watcher %v still has objects in in-flight %v(%v)", nodeID, "Interface", event)
		return false
	}

	evAckStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		//nothing received, failed.
		log.Infof("watcher %v still has not received anything %v(%v)", nodeID, "Interface", event)
		return false
	}

	if evAckStatus.LastObjectMeta.ResourceVersion != evStatus.LastObjectMeta.ResourceVersion {
		log.Infof("watcher %v resource version mismatch for %v(%v)  sent %v: recived %v",
			nodeID, "Interface", event, evStatus.LastObjectMeta.ResourceVersion,
			evAckStatus.LastObjectMeta.ResourceVersion)
		return false
	}

	return true
}

/*
//GetSentEventStatus
func (eh *InterfaceTopic) GetSentEventStatus(nodeID string, event api.EventType) *EventStatus {

    eh.Lock()
    defer eh.Unlock()
    var evStatus *EventStatus

    objStatus, ok := eh.opSentStatus[nodeID]
    if ok {
        evStatus, ok = objStatus.opStatus[event]
        if ok {
            return evStatus
        }
    }
    return nil
}


//GetAckedEventStatus
func (eh *InterfaceTopic) GetAckedEventStatus(nodeID string, event api.EventType) *EventStatus {

    eh.Lock()
    defer eh.Unlock()
    var evStatus *EventStatus

    objStatus, ok := eh.opAckedStatus[nodeID]
    if ok {
        evStatus, ok = objStatus.opStatus[event]
        if ok {
            return evStatus
        }
    }
    return nil
}

*/

// CreateInterface creates Interface
func (eh *InterfaceTopic) CreateInterface(ctx context.Context, objinfo *netproto.Interface) (*netproto.Interface, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received CreateInterface from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnInterfaceCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("Interface", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateInterface updates Interface
func (eh *InterfaceTopic) UpdateInterface(ctx context.Context, objinfo *netproto.Interface) (*netproto.Interface, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateInterface from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("Interface", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnInterfaceUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteInterface deletes an Interface
func (eh *InterfaceTopic) DeleteInterface(ctx context.Context, objinfo *netproto.Interface) (*netproto.Interface, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteInterface from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("Interface", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnInterfaceDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// InterfaceFromObj converts memdb object to Interface
func InterfaceFromObj(obj memdb.Object) (*netproto.Interface, error) {
	switch obj.(type) {
	case *netproto.Interface:
		eobj := obj.(*netproto.Interface)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetInterface returns a specific Interface
func (eh *InterfaceTopic) GetInterface(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.Interface, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("Interface", objmeta)
	if err != nil {
		return nil, err
	}

	return InterfaceFromObj(obj)
}

// ListInterfaces lists all Interfaces matching object selector
func (eh *InterfaceTopic) ListInterfaces(ctx context.Context, objsel *api.ObjectMeta) (*netproto.InterfaceList, error) {
	var objlist netproto.InterfaceList
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	filterFn := func(memdb.Object) bool {
		return true
	}

	if eh.statusReactor != nil {
		filterFn = eh.statusReactor.GetWatchFilter("Interface", objsel)
	}

	// walk all objects
	objs := eh.server.memDB.ListObjects("Interface", filterFn)
	//creationTime, _ := types.TimestampProto(time.Now())
	for _, oo := range objs {
		obj, err := InterfaceFromObj(oo)
		if err == nil {
			//obj.CreationTime = api.Timestamp{Timestamp: *creationTime}
			objlist.Interfaces = append(objlist.Interfaces, obj)
			//record the last object sent to check config sync
			eh.updateSentObjStatus(nodeID, api.EventType_UpdateEvent, &obj.ObjectMeta)
		}
	}

	return &objlist, nil
}

// WatchInterfaces watches Interfaces and sends streaming resp
func (eh *InterfaceTopic) WatchInterfaces(ometa *api.ObjectMeta, stream netproto.InterfaceApi_WatchInterfacesServer) error {
	// watch for changes
	watcher := memdb.Watcher{}
	watcher.Channel = make(chan memdb.Event, memdb.WatchLen)
	defer close(watcher.Channel)

	if eh.statusReactor != nil {
		watcher.Filter = eh.statusReactor.GetWatchFilter("Interface", ometa)
	} else {
		watcher.Filter = func(memdb.Object) bool {
			return true
		}
	}

	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	watcher.Name = nodeID
	eh.server.memDB.WatchObjects("Interface", &watcher)
	defer eh.server.memDB.StopWatchObjects("Interface", &watcher)

	// get a list of all Interfaces
	objlist, err := eh.ListInterfaces(context.Background(), ometa)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	eh.registerWatcher(nodeID, &watcher)
	defer eh.unRegisterWatcher(nodeID)

	// increment stats
	eh.server.Stats("Interface", "ActiveWatch").Inc()
	eh.server.Stats("Interface", "WatchConnect").Inc()
	defer eh.server.Stats("Interface", "ActiveWatch").Dec()
	defer eh.server.Stats("Interface", "WatchDisconnect").Inc()

	// walk all Interfaces and send it out
	watchEvts := netproto.InterfaceEventList{}
	for _, obj := range objlist.Interfaces {
		watchEvt := netproto.InterfaceEvent{
			EventType: api.EventType_CreateEvent,
			Interface: *obj,
		}
		watchEvts.InterfaceEvents = append(watchEvts.InterfaceEvents, &watchEvt)
	}
	if len(watchEvts.InterfaceEvents) > 0 {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending Interface to stream. Err: %v", err)
			return err
		}
	}
	timer := time.NewTimer(DefaultWatchHoldInterval)
	if !timer.Stop() {
		<-timer.C
	}

	running := false
	watchEvts = netproto.InterfaceEventList{}
	sendToStream := func() error {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending Interface to stream. Err: %v", err)
			return err
		}
		watchEvts = netproto.InterfaceEventList{}
		return nil
	}

	// loop forever on watch channel
	for {
		select {
		// read from channel
		case evt, ok := <-watcher.Channel:
			if !ok {
				log.Errorf("Error reading from channel. Closing watch")
				return errors.New("Error reading from channel")
			}

			// convert the events
			var etype api.EventType
			switch evt.EventType {
			case memdb.CreateEvent:
				etype = api.EventType_CreateEvent
			case memdb.UpdateEvent:
				etype = api.EventType_UpdateEvent
			case memdb.DeleteEvent:
				etype = api.EventType_DeleteEvent
			}

			// get the object
			obj, err := InterfaceFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.InterfaceEvent{
				EventType: etype,
				Interface: *obj,
			}
			watchEvts.InterfaceEvents = append(watchEvts.InterfaceEvents, &watchEvt)
			if !running {
				running = true
				timer.Reset(DefaultWatchHoldInterval)
			}
			if len(watchEvts.InterfaceEvents) >= DefaultWatchBatchSize {
				if err = sendToStream(); err != nil {
					return err
				}
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(DefaultWatchHoldInterval)
			}
			eh.updateSentObjStatus(nodeID, etype, &obj.ObjectMeta)
		case <-timer.C:
			running = false
			if err = sendToStream(); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// done
}

// updateInterfaceOper triggers oper update callbacks
func (eh *InterfaceTopic) updateInterfaceOper(oper *netproto.InterfaceEvent, nodeID string) error {
	eh.updateAckedObjStatus(nodeID, oper.EventType, &oper.Interface.ObjectMeta)
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("Interface", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			return eh.statusReactor.OnInterfaceOperUpdate(nodeID, &oper.Interface)
		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("Interface", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			eh.statusReactor.OnInterfaceOperDelete(nodeID, &oper.Interface)
		}
	}

	return nil
}

func (eh *InterfaceTopic) InterfaceOperUpdate(stream netproto.InterfaceApi_InterfaceOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("%v InterfaceOperUpdate stream ended. closing..", nodeID)
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from %v InterfaceOperUpdate stream. Err: %v", nodeID, err)
			return err
		}

		err = eh.updateInterfaceOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating Interface oper state. Err: %v", err)
		}
	}
}
