// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: app.proto
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

// FindApp finds an App by object meta
func (ms *MbusServer) FindApp(objmeta *api.ObjectMeta) (*netproto.App, error) {
	// find the object
	obj, err := ms.memDB.FindObject("App", objmeta)
	if err != nil {
		return nil, err
	}

	return AppFromObj(obj)
}

// ListApps lists all Apps in the mbus
func (ms *MbusServer) ListApps(ctx context.Context, filterFn func(memdb.Object) bool) ([]*netproto.App, error) {
	var objlist []*netproto.App

	// walk all objects
	objs := ms.memDB.ListObjects("App", filterFn)
	for _, oo := range objs {
		obj, err := AppFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// AppStatusReactor is the reactor interface implemented by controllers
type AppStatusReactor interface {
	OnAppCreateReq(nodeID string, objinfo *netproto.App) error
	OnAppUpdateReq(nodeID string, objinfo *netproto.App) error
	OnAppDeleteReq(nodeID string, objinfo *netproto.App) error
	OnAppOperUpdate(nodeID string, objinfo *netproto.App) error
	OnAppOperDelete(nodeID string, objinfo *netproto.App) error
	GetWatchFilter(kind string, ometa *api.ObjectMeta) func(memdb.Object) bool
}

type AppNodeStatus struct {
	nodeID        string
	watcher       *memdb.Watcher
	opSentStatus  map[api.EventType]*EventStatus
	opAckedStatus map[api.EventType]*EventStatus
}

// AppTopic is the App topic on message bus
type AppTopic struct {
	sync.Mutex
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor AppStatusReactor // status event reactor
	nodeStatus    map[string]*AppNodeStatus
}

// AddAppTopic returns a network RPC server
func AddAppTopic(server *MbusServer, reactor AppStatusReactor) (*AppTopic, error) {
	// RPC handler instance
	handler := AppTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
		nodeStatus:    make(map[string]*AppNodeStatus),
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterAppApiServer(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

func (eh *AppTopic) registerWatcher(nodeID string, watcher *memdb.Watcher) {
	eh.Lock()
	defer eh.Unlock()

	eh.nodeStatus[nodeID] = &AppNodeStatus{nodeID: nodeID, watcher: watcher}
	eh.nodeStatus[nodeID].opSentStatus = make(map[api.EventType]*EventStatus)
	eh.nodeStatus[nodeID].opAckedStatus = make(map[api.EventType]*EventStatus)
}

func (eh *AppTopic) unRegisterWatcher(nodeID string) {
	eh.Lock()
	defer eh.Unlock()

	delete(eh.nodeStatus, nodeID)
}

//update recv object status
func (eh *AppTopic) updateAckedObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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

	rcvdTime, _ := objMeta.ModTime.Time()
	sendTime, _ := objMeta.CreationTime.Time()
	delta := rcvdTime.Sub(sendTime)

	hdr.Record(nodeID+"_"+"App", delta)
	hdr.Record("App", delta)
	hdr.Record(nodeID, delta)

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
func (eh *AppTopic) updateSentObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

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
func (eh *AppTopic) WatcherInConfigSync(nodeID string, event api.EventType) bool {

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
		log.Infof("watcher %v still has objects in in-flight %v(%v)", nodeID, "App", event)
		return false
	}

	evAckStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		//nothing received, failed.
		log.Infof("watcher %v still has not received anything %v(%v)", nodeID, "App", event)
		return false
	}

	if evAckStatus.LastObjectMeta.ResourceVersion != evStatus.LastObjectMeta.ResourceVersion {
		log.Infof("watcher %v resource version mismatch for %v(%v)  sent %v: recived %v",
			nodeID, "App", event, evStatus.LastObjectMeta.ResourceVersion,
			evAckStatus.LastObjectMeta.ResourceVersion)
		return false
	}

	return true
}

/*
//GetSentEventStatus
func (eh *AppTopic) GetSentEventStatus(nodeID string, event api.EventType) *EventStatus {

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
func (eh *AppTopic) GetAckedEventStatus(nodeID string, event api.EventType) *EventStatus {

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

// CreateApp creates App
func (eh *AppTopic) CreateApp(ctx context.Context, objinfo *netproto.App) (*netproto.App, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received CreateApp from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnAppCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("App", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateApp updates App
func (eh *AppTopic) UpdateApp(ctx context.Context, objinfo *netproto.App) (*netproto.App, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateApp from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("App", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnAppUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteApp deletes an App
func (eh *AppTopic) DeleteApp(ctx context.Context, objinfo *netproto.App) (*netproto.App, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteApp from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("App", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnAppDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// AppFromObj converts memdb object to App
func AppFromObj(obj memdb.Object) (*netproto.App, error) {
	switch obj.(type) {
	case *netproto.App:
		eobj := obj.(*netproto.App)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetApp returns a specific App
func (eh *AppTopic) GetApp(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.App, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("App", objmeta)
	if err != nil {
		return nil, err
	}

	return AppFromObj(obj)
}

// ListApps lists all Apps matching object selector
func (eh *AppTopic) ListApps(ctx context.Context, objsel *api.ObjectMeta) (*netproto.AppList, error) {
	var objlist netproto.AppList
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	filterFn := func(memdb.Object) bool {
		return true
	}

	if eh.statusReactor != nil {
		filterFn = eh.statusReactor.GetWatchFilter("App", objsel)
	}

	// walk all objects
	objs := eh.server.memDB.ListObjects("App", filterFn)
	//creationTime, _ := types.TimestampProto(time.Now())
	for _, oo := range objs {
		obj, err := AppFromObj(oo)
		if err == nil {
			//obj.CreationTime = api.Timestamp{Timestamp: *creationTime}
			objlist.Apps = append(objlist.Apps, obj)
			//record the last object sent to check config sync
			eh.updateSentObjStatus(nodeID, api.EventType_UpdateEvent, &obj.ObjectMeta)
		}
	}

	return &objlist, nil
}

// WatchApps watches Apps and sends streaming resp
func (eh *AppTopic) WatchApps(ometa *api.ObjectMeta, stream netproto.AppApi_WatchAppsServer) error {
	// watch for changes
	watcher := memdb.Watcher{}
	watcher.Channel = make(chan memdb.Event, memdb.WatchLen)
	defer close(watcher.Channel)

	if eh.statusReactor != nil {
		watcher.Filter = eh.statusReactor.GetWatchFilter("App", ometa)
	} else {
		watcher.Filter = func(memdb.Object) bool {
			return true
		}
	}

	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	watcher.Name = nodeID
	eh.server.memDB.WatchObjects("App", &watcher)
	defer eh.server.memDB.StopWatchObjects("App", &watcher)

	// get a list of all Apps
	objlist, err := eh.ListApps(context.Background(), ometa)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	eh.registerWatcher(nodeID, &watcher)
	defer eh.unRegisterWatcher(nodeID)

	// increment stats
	eh.server.Stats("App", "ActiveWatch").Inc()
	eh.server.Stats("App", "WatchConnect").Inc()
	defer eh.server.Stats("App", "ActiveWatch").Dec()
	defer eh.server.Stats("App", "WatchDisconnect").Inc()

	// walk all Apps and send it out
	watchEvts := netproto.AppEventList{}
	for _, obj := range objlist.Apps {
		watchEvt := netproto.AppEvent{
			EventType: api.EventType_CreateEvent,
			App:       *obj,
		}
		watchEvts.AppEvents = append(watchEvts.AppEvents, &watchEvt)
	}
	if len(watchEvts.AppEvents) > 0 {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending App to stream. Err: %v", err)
			return err
		}
	}
	timer := time.NewTimer(DefaultWatchHoldInterval)
	if !timer.Stop() {
		<-timer.C
	}

	running := false
	watchEvts = netproto.AppEventList{}
	sendToStream := func() error {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending App to stream. Err: %v", err)
			return err
		}
		watchEvts = netproto.AppEventList{}
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
			obj, err := AppFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.AppEvent{
				EventType: etype,
				App:       *obj,
			}
			watchEvts.AppEvents = append(watchEvts.AppEvents, &watchEvt)
			if !running {
				running = true
				timer.Reset(DefaultWatchHoldInterval)
			}
			if len(watchEvts.AppEvents) >= DefaultWatchBatchSize {
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

// updateAppOper triggers oper update callbacks
func (eh *AppTopic) updateAppOper(oper *netproto.AppEvent, nodeID string) error {
	eh.updateAckedObjStatus(nodeID, oper.EventType, &oper.App.ObjectMeta)
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("App", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			return eh.statusReactor.OnAppOperUpdate(nodeID, &oper.App)
		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("App", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			eh.statusReactor.OnAppOperDelete(nodeID, &oper.App)
		}
	}

	return nil
}

func (eh *AppTopic) AppOperUpdate(stream netproto.AppApi_AppOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("%v AppOperUpdate stream ended. closing..", nodeID)
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from %v AppOperUpdate stream. Err: %v", nodeID, err)
			return err
		}

		err = eh.updateAppOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating App oper state. Err: %v", err)
		}
	}
}
