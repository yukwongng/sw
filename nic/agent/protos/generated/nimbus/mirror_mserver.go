// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: mirror.proto
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

// FindInterfaceMirrorSession finds an InterfaceMirrorSession by object meta
func (ms *MbusServer) FindInterfaceMirrorSession(objmeta *api.ObjectMeta) (*netproto.InterfaceMirrorSession, error) {
	// find the object
	obj, err := ms.memDB.FindObject("InterfaceMirrorSession", objmeta)
	if err != nil {
		return nil, err
	}

	return InterfaceMirrorSessionFromObj(obj)
}

// ListInterfaceMirrorSessions lists all InterfaceMirrorSessions in the mbus
func (ms *MbusServer) ListInterfaceMirrorSessions(ctx context.Context, nodeID string, filters []memdb.FilterFn) ([]*netproto.InterfaceMirrorSession, error) {
	var objlist []*netproto.InterfaceMirrorSession

	// walk all objects
	objs := ms.memDB.ListObjectsForReceiver("InterfaceMirrorSession", nodeID, filters)
	for _, oo := range objs {
		obj, err := InterfaceMirrorSessionFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// ListInterfaceMirrorSessionsNoFilter lists all InterfaceMirrorSessions in the mbus without applying a watch filter
func (ms *MbusServer) ListInterfaceMirrorSessionsNoFilter(ctx context.Context) ([]*netproto.InterfaceMirrorSession, error) {
	var objlist []*netproto.InterfaceMirrorSession

	// walk all objects
	objs := ms.memDB.ListObjects("InterfaceMirrorSession", nil)
	for _, oo := range objs {
		obj, err := InterfaceMirrorSessionFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// InterfaceMirrorSessionStatusReactor is the reactor interface implemented by controllers
type InterfaceMirrorSessionStatusReactor interface {
	OnInterfaceMirrorSessionCreateReq(nodeID string, objinfo *netproto.InterfaceMirrorSession) error
	OnInterfaceMirrorSessionUpdateReq(nodeID string, objinfo *netproto.InterfaceMirrorSession) error
	OnInterfaceMirrorSessionDeleteReq(nodeID string, objinfo *netproto.InterfaceMirrorSession) error
	OnInterfaceMirrorSessionOperUpdate(nodeID string, objinfo *netproto.InterfaceMirrorSession) error
	OnInterfaceMirrorSessionOperDelete(nodeID string, objinfo *netproto.InterfaceMirrorSession) error
	GetAgentWatchFilter(ctx context.Context, kind string, watchOptions *api.ListWatchOptions) []memdb.FilterFn
}

type InterfaceMirrorSessionNodeStatus struct {
	nodeID        string
	watcher       *memdb.Watcher
	opSentStatus  map[api.EventType]*EventStatus
	opAckedStatus map[api.EventType]*EventStatus
}

// InterfaceMirrorSessionTopic is the InterfaceMirrorSession topic on message bus
type InterfaceMirrorSessionTopic struct {
	sync.Mutex
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor InterfaceMirrorSessionStatusReactor // status event reactor
	nodeStatus    map[string]*InterfaceMirrorSessionNodeStatus
}

// AddInterfaceMirrorSessionTopic returns a network RPC server
func AddInterfaceMirrorSessionTopic(server *MbusServer, reactor InterfaceMirrorSessionStatusReactor) (*InterfaceMirrorSessionTopic, error) {
	// RPC handler instance
	handler := InterfaceMirrorSessionTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
		nodeStatus:    make(map[string]*InterfaceMirrorSessionNodeStatus),
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterInterfaceMirrorSessionApiV1Server(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

func (eh *InterfaceMirrorSessionTopic) registerWatcher(nodeID string, watcher *memdb.Watcher) {
	eh.Lock()
	defer eh.Unlock()

	eh.nodeStatus[nodeID] = &InterfaceMirrorSessionNodeStatus{nodeID: nodeID, watcher: watcher}
	eh.nodeStatus[nodeID].opSentStatus = make(map[api.EventType]*EventStatus)
	eh.nodeStatus[nodeID].opAckedStatus = make(map[api.EventType]*EventStatus)
}

func (eh *InterfaceMirrorSessionTopic) unRegisterWatcher(nodeID string) {
	eh.Lock()
	defer eh.Unlock()

	delete(eh.nodeStatus, nodeID)
}

//update recv object status
func (eh *InterfaceMirrorSessionTopic) updateAckedObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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

		hdr.Record(nodeID+"_"+"InterfaceMirrorSession", delta)
		hdr.Record("InterfaceMirrorSession", delta)
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
func (eh *InterfaceMirrorSessionTopic) updateSentObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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
func (eh *InterfaceMirrorSessionTopic) WatcherInConfigSync(nodeID string, event api.EventType) bool {

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
		log.Infof("watcher %v still has objects in in-flight %v(%v)", nodeID, "InterfaceMirrorSession", event)
		return false
	}

	evAckStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		//nothing received, failed.
		log.Infof("watcher %v still has not received anything %v(%v)", nodeID, "InterfaceMirrorSession", event)
		return false
	}

	if evAckStatus.LastObjectMeta.ResourceVersion < evStatus.LastObjectMeta.ResourceVersion {
		log.Infof("watcher %v resource version mismatch for %v(%v)  sent %v: recived %v",
			nodeID, "InterfaceMirrorSession", event, evStatus.LastObjectMeta.ResourceVersion,
			evAckStatus.LastObjectMeta.ResourceVersion)
		return false
	}

	return true
}

/*
//GetSentEventStatus
func (eh *InterfaceMirrorSessionTopic) GetSentEventStatus(nodeID string, event api.EventType) *EventStatus {

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
func (eh *InterfaceMirrorSessionTopic) GetAckedEventStatus(nodeID string, event api.EventType) *EventStatus {

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

// CreateInterfaceMirrorSession creates InterfaceMirrorSession
func (eh *InterfaceMirrorSessionTopic) CreateInterfaceMirrorSession(ctx context.Context, objinfo *netproto.InterfaceMirrorSession) (*netproto.InterfaceMirrorSession, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received CreateInterfaceMirrorSession from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnInterfaceMirrorSessionCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("InterfaceMirrorSession", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateInterfaceMirrorSession updates InterfaceMirrorSession
func (eh *InterfaceMirrorSessionTopic) UpdateInterfaceMirrorSession(ctx context.Context, objinfo *netproto.InterfaceMirrorSession) (*netproto.InterfaceMirrorSession, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateInterfaceMirrorSession from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("InterfaceMirrorSession", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnInterfaceMirrorSessionUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteInterfaceMirrorSession deletes an InterfaceMirrorSession
func (eh *InterfaceMirrorSessionTopic) DeleteInterfaceMirrorSession(ctx context.Context, objinfo *netproto.InterfaceMirrorSession) (*netproto.InterfaceMirrorSession, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteInterfaceMirrorSession from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("InterfaceMirrorSession", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnInterfaceMirrorSessionDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// InterfaceMirrorSessionFromObj converts memdb object to InterfaceMirrorSession
func InterfaceMirrorSessionFromObj(obj memdb.Object) (*netproto.InterfaceMirrorSession, error) {
	switch obj.(type) {
	case *netproto.InterfaceMirrorSession:
		eobj := obj.(*netproto.InterfaceMirrorSession)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetInterfaceMirrorSession returns a specific InterfaceMirrorSession
func (eh *InterfaceMirrorSessionTopic) GetInterfaceMirrorSession(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.InterfaceMirrorSession, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("InterfaceMirrorSession", objmeta)
	if err != nil {
		return nil, err
	}

	return InterfaceMirrorSessionFromObj(obj)
}

// ListInterfaceMirrorSessions lists all InterfaceMirrorSessions matching object selector
func (eh *InterfaceMirrorSessionTopic) ListInterfaceMirrorSessions(ctx context.Context, objsel *api.ListWatchOptions) (*netproto.InterfaceMirrorSessionList, error) {
	var objlist netproto.InterfaceMirrorSessionList
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	filters := []memdb.FilterFn{}

	filterFn := func(obj, prev memdb.Object) bool {
		return true
	}

	if eh.statusReactor != nil {
		filters = eh.statusReactor.GetAgentWatchFilter(ctx, "netproto.InterfaceMirrorSession", objsel)
	} else {
		filters = append(filters, filterFn)
	}

	// walk all objects
	objs := eh.server.memDB.ListObjectsForReceiver("InterfaceMirrorSession", nodeID, filters)
	//creationTime, _ := types.TimestampProto(time.Now())
	for _, oo := range objs {
		obj, err := InterfaceMirrorSessionFromObj(oo)
		if err == nil {
			//obj.CreationTime = api.Timestamp{Timestamp: *creationTime}
			objlist.InterfaceMirrorSessions = append(objlist.InterfaceMirrorSessions, obj)
			//record the last object sent to check config sync
			eh.updateSentObjStatus(nodeID, api.EventType_UpdateEvent, &obj.ObjectMeta)
		}
	}

	return &objlist, nil
}

// WatchInterfaceMirrorSessions watches InterfaceMirrorSessions and sends streaming resp
func (eh *InterfaceMirrorSessionTopic) WatchInterfaceMirrorSessions(watchOptions *api.ListWatchOptions, stream netproto.InterfaceMirrorSessionApiV1_WatchInterfaceMirrorSessionsServer) error {
	// watch for changes
	watcher := memdb.Watcher{}
	watcher.Channel = make(chan memdb.Event, memdb.WatchLen)
	watcher.Filters = make(map[string][]memdb.FilterFn)
	defer close(watcher.Channel)

	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	if eh.statusReactor != nil {
		watcher.Filters["InterfaceMirrorSession"] = eh.statusReactor.GetAgentWatchFilter(ctx, "InterfaceMirrorSession", watchOptions)
	} else {
		filt := func(obj, prev memdb.Object) bool {
			return true
		}
		watcher.Filters["InterfaceMirrorSession"] = append(watcher.Filters["InterfaceMirrorSession"], filt)
	}

	watcher.Name = nodeID
	err := eh.server.memDB.WatchObjects("InterfaceMirrorSession", &watcher)
	if err != nil {
		log.Errorf("Error Starting watch for kind %v Err: %v", "InterfaceMirrorSession", err)
		return err
	}
	defer eh.server.memDB.StopWatchObjects("InterfaceMirrorSession", &watcher)

	// get a list of all InterfaceMirrorSessions
	objlist, err := eh.ListInterfaceMirrorSessions(context.Background(), watchOptions)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	eh.registerWatcher(nodeID, &watcher)
	defer eh.unRegisterWatcher(nodeID)

	// increment stats
	eh.server.Stats("InterfaceMirrorSession", "ActiveWatch").Inc()
	eh.server.Stats("InterfaceMirrorSession", "WatchConnect").Inc()
	defer eh.server.Stats("InterfaceMirrorSession", "ActiveWatch").Dec()
	defer eh.server.Stats("InterfaceMirrorSession", "WatchDisconnect").Inc()

	// walk all InterfaceMirrorSessions and send it out
	watchEvts := netproto.InterfaceMirrorSessionEventList{}
	for _, obj := range objlist.InterfaceMirrorSessions {
		watchEvt := netproto.InterfaceMirrorSessionEvent{
			EventType: api.EventType_CreateEvent,

			InterfaceMirrorSession: *obj,
		}
		watchEvts.InterfaceMirrorSessionEvents = append(watchEvts.InterfaceMirrorSessionEvents, &watchEvt)
	}
	if len(watchEvts.InterfaceMirrorSessionEvents) > 0 {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending InterfaceMirrorSession to stream. Err: %v", err)
			return err
		}
	}
	timer := time.NewTimer(DefaultWatchHoldInterval)
	if !timer.Stop() {
		<-timer.C
	}

	running := false
	watchEvts = netproto.InterfaceMirrorSessionEventList{}
	sendToStream := func() error {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending InterfaceMirrorSession to stream. Err: %v", err)
			return err
		}
		watchEvts = netproto.InterfaceMirrorSessionEventList{}
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
			obj, err := InterfaceMirrorSessionFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.InterfaceMirrorSessionEvent{
				EventType: etype,

				InterfaceMirrorSession: *obj,
			}
			watchEvts.InterfaceMirrorSessionEvents = append(watchEvts.InterfaceMirrorSessionEvents, &watchEvt)
			if !running {
				running = true
				timer.Reset(DefaultWatchHoldInterval)
			}
			if len(watchEvts.InterfaceMirrorSessionEvents) >= DefaultWatchBatchSize {
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

// updateInterfaceMirrorSessionOper triggers oper update callbacks
func (eh *InterfaceMirrorSessionTopic) updateInterfaceMirrorSessionOper(oper *netproto.InterfaceMirrorSessionEvent, nodeID string) error {
	eh.updateAckedObjStatus(nodeID, oper.EventType, &oper.InterfaceMirrorSession.ObjectMeta)
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("InterfaceMirrorSession", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {

			return eh.statusReactor.OnInterfaceMirrorSessionOperUpdate(nodeID, &oper.InterfaceMirrorSession)

		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("InterfaceMirrorSession", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {

			eh.statusReactor.OnInterfaceMirrorSessionOperDelete(nodeID, &oper.InterfaceMirrorSession)

		}
	}

	return nil
}

func (eh *InterfaceMirrorSessionTopic) InterfaceMirrorSessionOperUpdate(stream netproto.InterfaceMirrorSessionApiV1_InterfaceMirrorSessionOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("%v InterfaceMirrorSessionOperUpdate stream ended. closing..", nodeID)
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from %v InterfaceMirrorSessionOperUpdate stream. Err: %v", nodeID, err)
			return err
		}

		err = eh.updateInterfaceMirrorSessionOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating InterfaceMirrorSession oper state. Err: %v", err)
		}
	}
}

// FindMirrorSession finds an MirrorSession by object meta
func (ms *MbusServer) FindMirrorSession(objmeta *api.ObjectMeta) (*netproto.MirrorSession, error) {
	// find the object
	obj, err := ms.memDB.FindObject("MirrorSession", objmeta)
	if err != nil {
		return nil, err
	}

	return MirrorSessionFromObj(obj)
}

// ListMirrorSessions lists all MirrorSessions in the mbus
func (ms *MbusServer) ListMirrorSessions(ctx context.Context, nodeID string, filters []memdb.FilterFn) ([]*netproto.MirrorSession, error) {
	var objlist []*netproto.MirrorSession

	// walk all objects
	objs := ms.memDB.ListObjectsForReceiver("MirrorSession", nodeID, filters)
	for _, oo := range objs {
		obj, err := MirrorSessionFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// ListMirrorSessionsNoFilter lists all MirrorSessions in the mbus without applying a watch filter
func (ms *MbusServer) ListMirrorSessionsNoFilter(ctx context.Context) ([]*netproto.MirrorSession, error) {
	var objlist []*netproto.MirrorSession

	// walk all objects
	objs := ms.memDB.ListObjects("MirrorSession", nil)
	for _, oo := range objs {
		obj, err := MirrorSessionFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// MirrorSessionStatusReactor is the reactor interface implemented by controllers
type MirrorSessionStatusReactor interface {
	OnMirrorSessionCreateReq(nodeID string, objinfo *netproto.MirrorSession) error
	OnMirrorSessionUpdateReq(nodeID string, objinfo *netproto.MirrorSession) error
	OnMirrorSessionDeleteReq(nodeID string, objinfo *netproto.MirrorSession) error
	OnMirrorSessionOperUpdate(nodeID string, objinfo *netproto.MirrorSession) error
	OnMirrorSessionOperDelete(nodeID string, objinfo *netproto.MirrorSession) error
	GetAgentWatchFilter(ctx context.Context, kind string, watchOptions *api.ListWatchOptions) []memdb.FilterFn
}

type MirrorSessionNodeStatus struct {
	nodeID        string
	watcher       *memdb.Watcher
	opSentStatus  map[api.EventType]*EventStatus
	opAckedStatus map[api.EventType]*EventStatus
}

// MirrorSessionTopic is the MirrorSession topic on message bus
type MirrorSessionTopic struct {
	sync.Mutex
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor MirrorSessionStatusReactor // status event reactor
	nodeStatus    map[string]*MirrorSessionNodeStatus
}

// AddMirrorSessionTopic returns a network RPC server
func AddMirrorSessionTopic(server *MbusServer, reactor MirrorSessionStatusReactor) (*MirrorSessionTopic, error) {
	// RPC handler instance
	handler := MirrorSessionTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
		nodeStatus:    make(map[string]*MirrorSessionNodeStatus),
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterMirrorSessionApiV2Server(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

func (eh *MirrorSessionTopic) registerWatcher(nodeID string, watcher *memdb.Watcher) {
	eh.Lock()
	defer eh.Unlock()

	eh.nodeStatus[nodeID] = &MirrorSessionNodeStatus{nodeID: nodeID, watcher: watcher}
	eh.nodeStatus[nodeID].opSentStatus = make(map[api.EventType]*EventStatus)
	eh.nodeStatus[nodeID].opAckedStatus = make(map[api.EventType]*EventStatus)
}

func (eh *MirrorSessionTopic) unRegisterWatcher(nodeID string) {
	eh.Lock()
	defer eh.Unlock()

	delete(eh.nodeStatus, nodeID)
}

//update recv object status
func (eh *MirrorSessionTopic) updateAckedObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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

		hdr.Record(nodeID+"_"+"MirrorSession", delta)
		hdr.Record("MirrorSession", delta)
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
func (eh *MirrorSessionTopic) updateSentObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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
func (eh *MirrorSessionTopic) WatcherInConfigSync(nodeID string, event api.EventType) bool {

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
		log.Infof("watcher %v still has objects in in-flight %v(%v)", nodeID, "MirrorSession", event)
		return false
	}

	evAckStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		//nothing received, failed.
		log.Infof("watcher %v still has not received anything %v(%v)", nodeID, "MirrorSession", event)
		return false
	}

	if evAckStatus.LastObjectMeta.ResourceVersion < evStatus.LastObjectMeta.ResourceVersion {
		log.Infof("watcher %v resource version mismatch for %v(%v)  sent %v: recived %v",
			nodeID, "MirrorSession", event, evStatus.LastObjectMeta.ResourceVersion,
			evAckStatus.LastObjectMeta.ResourceVersion)
		return false
	}

	return true
}

/*
//GetSentEventStatus
func (eh *MirrorSessionTopic) GetSentEventStatus(nodeID string, event api.EventType) *EventStatus {

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
func (eh *MirrorSessionTopic) GetAckedEventStatus(nodeID string, event api.EventType) *EventStatus {

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

// CreateMirrorSession creates MirrorSession
func (eh *MirrorSessionTopic) CreateMirrorSession(ctx context.Context, objinfo *netproto.MirrorSession) (*netproto.MirrorSession, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received CreateMirrorSession from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnMirrorSessionCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("MirrorSession", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateMirrorSession updates MirrorSession
func (eh *MirrorSessionTopic) UpdateMirrorSession(ctx context.Context, objinfo *netproto.MirrorSession) (*netproto.MirrorSession, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateMirrorSession from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("MirrorSession", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnMirrorSessionUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteMirrorSession deletes an MirrorSession
func (eh *MirrorSessionTopic) DeleteMirrorSession(ctx context.Context, objinfo *netproto.MirrorSession) (*netproto.MirrorSession, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteMirrorSession from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("MirrorSession", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnMirrorSessionDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// MirrorSessionFromObj converts memdb object to MirrorSession
func MirrorSessionFromObj(obj memdb.Object) (*netproto.MirrorSession, error) {
	switch obj.(type) {
	case *netproto.MirrorSession:
		eobj := obj.(*netproto.MirrorSession)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetMirrorSession returns a specific MirrorSession
func (eh *MirrorSessionTopic) GetMirrorSession(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.MirrorSession, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("MirrorSession", objmeta)
	if err != nil {
		return nil, err
	}

	return MirrorSessionFromObj(obj)
}

// ListMirrorSessions lists all MirrorSessions matching object selector
func (eh *MirrorSessionTopic) ListMirrorSessions(ctx context.Context, objsel *api.ListWatchOptions) (*netproto.MirrorSessionList, error) {
	var objlist netproto.MirrorSessionList
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	filters := []memdb.FilterFn{}

	filterFn := func(obj, prev memdb.Object) bool {
		return true
	}

	if eh.statusReactor != nil {
		filters = eh.statusReactor.GetAgentWatchFilter(ctx, "netproto.MirrorSession", objsel)
	} else {
		filters = append(filters, filterFn)
	}

	// walk all objects
	objs := eh.server.memDB.ListObjectsForReceiver("MirrorSession", nodeID, filters)
	//creationTime, _ := types.TimestampProto(time.Now())
	for _, oo := range objs {
		obj, err := MirrorSessionFromObj(oo)
		if err == nil {
			//obj.CreationTime = api.Timestamp{Timestamp: *creationTime}
			objlist.MirrorSessions = append(objlist.MirrorSessions, obj)
			//record the last object sent to check config sync
			eh.updateSentObjStatus(nodeID, api.EventType_UpdateEvent, &obj.ObjectMeta)
		}
	}

	return &objlist, nil
}

// WatchMirrorSessions watches MirrorSessions and sends streaming resp
func (eh *MirrorSessionTopic) WatchMirrorSessions(watchOptions *api.ListWatchOptions, stream netproto.MirrorSessionApiV2_WatchMirrorSessionsServer) error {
	// watch for changes
	watcher := memdb.Watcher{}
	watcher.Channel = make(chan memdb.Event, memdb.WatchLen)
	watcher.Filters = make(map[string][]memdb.FilterFn)
	defer close(watcher.Channel)

	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	if eh.statusReactor != nil {
		watcher.Filters["MirrorSession"] = eh.statusReactor.GetAgentWatchFilter(ctx, "MirrorSession", watchOptions)
	} else {
		filt := func(obj, prev memdb.Object) bool {
			return true
		}
		watcher.Filters["MirrorSession"] = append(watcher.Filters["MirrorSession"], filt)
	}

	watcher.Name = nodeID
	err := eh.server.memDB.WatchObjects("MirrorSession", &watcher)
	if err != nil {
		log.Errorf("Error Starting watch for kind %v Err: %v", "MirrorSession", err)
		return err
	}
	defer eh.server.memDB.StopWatchObjects("MirrorSession", &watcher)

	// get a list of all MirrorSessions
	objlist, err := eh.ListMirrorSessions(context.Background(), watchOptions)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	eh.registerWatcher(nodeID, &watcher)
	defer eh.unRegisterWatcher(nodeID)

	// increment stats
	eh.server.Stats("MirrorSession", "ActiveWatch").Inc()
	eh.server.Stats("MirrorSession", "WatchConnect").Inc()
	defer eh.server.Stats("MirrorSession", "ActiveWatch").Dec()
	defer eh.server.Stats("MirrorSession", "WatchDisconnect").Inc()

	// walk all MirrorSessions and send it out
	watchEvts := netproto.MirrorSessionEventList{}
	for _, obj := range objlist.MirrorSessions {
		watchEvt := netproto.MirrorSessionEvent{
			EventType: api.EventType_CreateEvent,

			MirrorSession: obj,
		}
		watchEvts.MirrorSessionEvents = append(watchEvts.MirrorSessionEvents, &watchEvt)
	}
	if len(watchEvts.MirrorSessionEvents) > 0 {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending MirrorSession to stream. Err: %v", err)
			return err
		}
	}
	timer := time.NewTimer(DefaultWatchHoldInterval)
	if !timer.Stop() {
		<-timer.C
	}

	running := false
	watchEvts = netproto.MirrorSessionEventList{}
	sendToStream := func() error {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending MirrorSession to stream. Err: %v", err)
			return err
		}
		watchEvts = netproto.MirrorSessionEventList{}
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
			obj, err := MirrorSessionFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.MirrorSessionEvent{
				EventType: etype,

				MirrorSession: obj,
			}
			watchEvts.MirrorSessionEvents = append(watchEvts.MirrorSessionEvents, &watchEvt)
			if !running {
				running = true
				timer.Reset(DefaultWatchHoldInterval)
			}
			if len(watchEvts.MirrorSessionEvents) >= DefaultWatchBatchSize {
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

// updateMirrorSessionOper triggers oper update callbacks
func (eh *MirrorSessionTopic) updateMirrorSessionOper(oper *netproto.MirrorSessionEvent, nodeID string) error {
	eh.updateAckedObjStatus(nodeID, oper.EventType, &oper.MirrorSession.ObjectMeta)
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("MirrorSession", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {

			return eh.statusReactor.OnMirrorSessionOperUpdate(nodeID, oper.MirrorSession)

		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("MirrorSession", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {

			eh.statusReactor.OnMirrorSessionOperDelete(nodeID, oper.MirrorSession)

		}
	}

	return nil
}

func (eh *MirrorSessionTopic) MirrorSessionOperUpdate(stream netproto.MirrorSessionApiV2_MirrorSessionOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("%v MirrorSessionOperUpdate stream ended. closing..", nodeID)
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from %v MirrorSessionOperUpdate stream. Err: %v", nodeID, err)
			return err
		}

		err = eh.updateMirrorSessionOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating MirrorSession oper state. Err: %v", err)
		}
	}
}
