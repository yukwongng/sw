// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: network.proto
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

// FindNetwork finds an Network by object meta
func (ms *MbusServer) FindNetwork(objmeta *api.ObjectMeta) (*netproto.Network, error) {
	// find the object
	obj, err := ms.memDB.FindObject("Network", objmeta)
	if err != nil {
		return nil, err
	}

	return NetworkFromObj(obj)
}

// ListNetworks lists all Networks in the mbus
func (ms *MbusServer) ListNetworks(ctx context.Context, filterFn func(memdb.Object) bool) ([]*netproto.Network, error) {
	var objlist []*netproto.Network

	// walk all objects
	objs := ms.memDB.ListObjects("Network", filterFn)
	for _, oo := range objs {
		obj, err := NetworkFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// NetworkStatusReactor is the reactor interface implemented by controllers
type NetworkStatusReactor interface {
	OnNetworkCreateReq(nodeID string, objinfo *netproto.Network) error
	OnNetworkUpdateReq(nodeID string, objinfo *netproto.Network) error
	OnNetworkDeleteReq(nodeID string, objinfo *netproto.Network) error
	OnNetworkOperUpdate(nodeID string, objinfo *netproto.Network) error
	OnNetworkOperDelete(nodeID string, objinfo *netproto.Network) error
	GetWatchFilter(kind string, ometa *api.ObjectMeta) func(memdb.Object) bool
}

type NetworkNodeStatus struct {
	nodeID        string
	watcher       *memdb.Watcher
	opSentStatus  map[api.EventType]*EventStatus
	opAckedStatus map[api.EventType]*EventStatus
}

// NetworkTopic is the Network topic on message bus
type NetworkTopic struct {
	sync.Mutex
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor NetworkStatusReactor // status event reactor
	nodeStatus    map[string]*NetworkNodeStatus
}

// AddNetworkTopic returns a network RPC server
func AddNetworkTopic(server *MbusServer, reactor NetworkStatusReactor) (*NetworkTopic, error) {
	// RPC handler instance
	handler := NetworkTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
		nodeStatus:    make(map[string]*NetworkNodeStatus),
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterNetworkApiServer(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

func (eh *NetworkTopic) registerWatcher(nodeID string, watcher *memdb.Watcher) {
	eh.Lock()
	defer eh.Unlock()

	eh.nodeStatus[nodeID] = &NetworkNodeStatus{nodeID: nodeID, watcher: watcher}
	eh.nodeStatus[nodeID].opSentStatus = make(map[api.EventType]*EventStatus)
	eh.nodeStatus[nodeID].opAckedStatus = make(map[api.EventType]*EventStatus)
}

func (eh *NetworkTopic) unRegisterWatcher(nodeID string) {
	eh.Lock()
	defer eh.Unlock()

	delete(eh.nodeStatus, nodeID)
}

//update recv object status
func (eh *NetworkTopic) updateAckedObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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

		hdr.Record(nodeID+"_"+"Network", delta)
		hdr.Record("Network", delta)
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
func (eh *NetworkTopic) updateSentObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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
func (eh *NetworkTopic) WatcherInConfigSync(nodeID string, event api.EventType) bool {

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
		log.Infof("watcher %v still has objects in in-flight %v(%v)", nodeID, "Network", event)
		return false
	}

	evAckStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		//nothing received, failed.
		log.Infof("watcher %v still has not received anything %v(%v)", nodeID, "Network", event)
		return false
	}

	if evAckStatus.LastObjectMeta.ResourceVersion != evStatus.LastObjectMeta.ResourceVersion {
		log.Infof("watcher %v resource version mismatch for %v(%v)  sent %v: recived %v",
			nodeID, "Network", event, evStatus.LastObjectMeta.ResourceVersion,
			evAckStatus.LastObjectMeta.ResourceVersion)
		return false
	}

	return true
}

/*
//GetSentEventStatus
func (eh *NetworkTopic) GetSentEventStatus(nodeID string, event api.EventType) *EventStatus {

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
func (eh *NetworkTopic) GetAckedEventStatus(nodeID string, event api.EventType) *EventStatus {

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

// CreateNetwork creates Network
func (eh *NetworkTopic) CreateNetwork(ctx context.Context, objinfo *netproto.Network) (*netproto.Network, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received CreateNetwork from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnNetworkCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("Network", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateNetwork updates Network
func (eh *NetworkTopic) UpdateNetwork(ctx context.Context, objinfo *netproto.Network) (*netproto.Network, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateNetwork from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("Network", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnNetworkUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteNetwork deletes an Network
func (eh *NetworkTopic) DeleteNetwork(ctx context.Context, objinfo *netproto.Network) (*netproto.Network, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteNetwork from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("Network", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnNetworkDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// NetworkFromObj converts memdb object to Network
func NetworkFromObj(obj memdb.Object) (*netproto.Network, error) {
	switch obj.(type) {
	case *netproto.Network:
		eobj := obj.(*netproto.Network)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetNetwork returns a specific Network
func (eh *NetworkTopic) GetNetwork(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.Network, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("Network", objmeta)
	if err != nil {
		return nil, err
	}

	return NetworkFromObj(obj)
}

// ListNetworks lists all Networks matching object selector
func (eh *NetworkTopic) ListNetworks(ctx context.Context, objsel *api.ObjectMeta) (*netproto.NetworkList, error) {
	var objlist netproto.NetworkList
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	filterFn := func(memdb.Object) bool {
		return true
	}

	if eh.statusReactor != nil {
		filterFn = eh.statusReactor.GetWatchFilter("Network", objsel)
	}

	// walk all objects
	objs := eh.server.memDB.ListObjects("Network", filterFn)
	//creationTime, _ := types.TimestampProto(time.Now())
	for _, oo := range objs {
		obj, err := NetworkFromObj(oo)
		if err == nil {
			//obj.CreationTime = api.Timestamp{Timestamp: *creationTime}
			objlist.Networks = append(objlist.Networks, obj)
			//record the last object sent to check config sync
			eh.updateSentObjStatus(nodeID, api.EventType_UpdateEvent, &obj.ObjectMeta)
		}
	}

	return &objlist, nil
}

// WatchNetworks watches Networks and sends streaming resp
func (eh *NetworkTopic) WatchNetworks(ometa *api.ObjectMeta, stream netproto.NetworkApi_WatchNetworksServer) error {
	// watch for changes
	watcher := memdb.Watcher{}
	watcher.Channel = make(chan memdb.Event, memdb.WatchLen)
	defer close(watcher.Channel)

	if eh.statusReactor != nil {
		watcher.Filter = eh.statusReactor.GetWatchFilter("Network", ometa)
	} else {
		watcher.Filter = func(memdb.Object) bool {
			return true
		}
	}

	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	watcher.Name = nodeID
	eh.server.memDB.WatchObjects("Network", &watcher)
	defer eh.server.memDB.StopWatchObjects("Network", &watcher)

	// get a list of all Networks
	objlist, err := eh.ListNetworks(context.Background(), ometa)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	eh.registerWatcher(nodeID, &watcher)
	defer eh.unRegisterWatcher(nodeID)

	// increment stats
	eh.server.Stats("Network", "ActiveWatch").Inc()
	eh.server.Stats("Network", "WatchConnect").Inc()
	defer eh.server.Stats("Network", "ActiveWatch").Dec()
	defer eh.server.Stats("Network", "WatchDisconnect").Inc()

	// walk all Networks and send it out
	watchEvts := netproto.NetworkEventList{}
	for _, obj := range objlist.Networks {
		watchEvt := netproto.NetworkEvent{
			EventType: api.EventType_CreateEvent,
			Network:   *obj,
		}
		watchEvts.NetworkEvents = append(watchEvts.NetworkEvents, &watchEvt)
	}
	if len(watchEvts.NetworkEvents) > 0 {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending Network to stream. Err: %v", err)
			return err
		}
	}
	timer := time.NewTimer(DefaultWatchHoldInterval)
	if !timer.Stop() {
		<-timer.C
	}

	running := false
	watchEvts = netproto.NetworkEventList{}
	sendToStream := func() error {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending Network to stream. Err: %v", err)
			return err
		}
		watchEvts = netproto.NetworkEventList{}
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
			obj, err := NetworkFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.NetworkEvent{
				EventType: etype,
				Network:   *obj,
			}
			watchEvts.NetworkEvents = append(watchEvts.NetworkEvents, &watchEvt)
			if !running {
				running = true
				timer.Reset(DefaultWatchHoldInterval)
			}
			if len(watchEvts.NetworkEvents) >= DefaultWatchBatchSize {
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

// updateNetworkOper triggers oper update callbacks
func (eh *NetworkTopic) updateNetworkOper(oper *netproto.NetworkEvent, nodeID string) error {
	eh.updateAckedObjStatus(nodeID, oper.EventType, &oper.Network.ObjectMeta)
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("Network", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			return eh.statusReactor.OnNetworkOperUpdate(nodeID, &oper.Network)
		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("Network", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			eh.statusReactor.OnNetworkOperDelete(nodeID, &oper.Network)
		}
	}

	return nil
}

func (eh *NetworkTopic) NetworkOperUpdate(stream netproto.NetworkApi_NetworkOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("%v NetworkOperUpdate stream ended. closing..", nodeID)
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from %v NetworkOperUpdate stream. Err: %v", nodeID, err)
			return err
		}

		err = eh.updateNetworkOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating Network oper state. Err: %v", err)
		}
	}
}
