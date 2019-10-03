// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: app.proto
*/

package nimbus

import (
	"context"
	"sync"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/state"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/utils/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"
)

type AppReactor interface {
	CreateApp(appObj *netproto.App) error               // creates an App
	FindApp(meta api.ObjectMeta) (*netproto.App, error) // finds an App
	ListApp() []*netproto.App                           // lists all Apps
	UpdateApp(appObj *netproto.App) error               // updates an App
	DeleteApp(appObj, ns, name string) error            // deletes an App
	GetWatchOptions(cts context.Context, kind string) api.ObjectMeta
}
type AppOStream struct {
	sync.Mutex
	stream netproto.AppApi_AppOperUpdateClient
}

// WatchApps runs App watcher loop
func (client *NimbusClient) WatchApps(ctx context.Context, reactor AppReactor) {
	// setup wait group
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()
	client.debugStats.AddInt("ActiveAppWatch", 1)

	// make sure rpc client is good
	if client.rpcClient == nil || client.rpcClient.ClientConn == nil || client.rpcClient.ClientConn.GetState() != connectivity.Ready {
		log.Errorf("RPC client is disconnected. Exiting watch")
		return
	}

	// start the watch
	ometa := reactor.GetWatchOptions(ctx, "App")
	appRPCClient := netproto.NewAppApiClient(client.rpcClient.ClientConn)
	stream, err := appRPCClient.WatchApps(ctx, &ometa)
	if err != nil {
		log.Errorf("Error watching App. Err: %v", err)
		return
	}

	// start oper update stream
	opStream, err := appRPCClient.AppOperUpdate(ctx)
	if err != nil {
		log.Errorf("Error starting App oper updates. Err: %v", err)
		return
	}

	ostream := &AppOStream{stream: opStream}

	// get a list of objects
	objList, err := appRPCClient.ListApps(ctx, &ometa)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok || st.Code() == codes.Unavailable {
			log.Errorf("Error getting App list. Err: %v", err)
			return
		}
	} else {
		// perform a diff of the states
		client.diffApps(objList, reactor, ostream)
	}

	// start grpc stream recv
	recvCh := make(chan *netproto.AppEvent, evChanLength)
	go client.watchAppRecvLoop(stream, recvCh)

	// loop till the end
	for {
		evtWork := func(evt *netproto.AppEvent) {
			client.debugStats.AddInt("AppWatchEvents", 1)
			log.Infof("Ctrlerif: agent %s got App watch event: Type: {%+v} App:{%+v}", client.clientName, evt.EventType, evt.App.ObjectMeta)
			client.lockObject(evt.App.GetObjectKind(), evt.App.ObjectMeta)
			go client.processAppEvent(*evt, reactor, ostream)
			//Give it some time to increment waitgrp
			time.Sleep(100 * time.Microsecond)
		}
		//Give priority to evnt work.
		select {
		case evt, ok := <-recvCh:
			if !ok {
				log.Warnf("App Watch channel closed. Exisint AppWatch")
				return
			}
			evtWork(evt)
		// periodic resync
		case <-time.After(resyncInterval):
			//Give priority to evt work
			//Wait for batch interval for inflight work
			time.Sleep(DefaultWatchHoldInterval)
			select {
			case evt, ok := <-recvCh:
				if !ok {
					log.Warnf("App Watch channel closed. Exisint AppWatch")
					return
				}
				evtWork(evt)
				continue
			default:
			}
			// get a list of objects
			objList, err := appRPCClient.ListApps(ctx, &ometa)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok || st.Code() == codes.Unavailable {
					log.Errorf("Error getting App list. Err: %v", err)
					return
				}
			} else {
				client.debugStats.AddInt("AppWatchResyncs", 1)
				// perform a diff of the states
				client.diffApps(objList, reactor, ostream)
			}
		}
	}
}

// watchAppRecvLoop receives from stream and write it to a channel
func (client *NimbusClient) watchAppRecvLoop(stream netproto.AppApi_WatchAppsClient, recvch chan<- *netproto.AppEvent) {
	defer close(recvch)
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// loop till the end
	for {
		// receive from stream
		objList, err := stream.Recv()
		if err != nil {
			log.Errorf("Error receiving from watch channel. Exiting App watch. Err: %v", err)
			return
		}
		for _, evt := range objList.AppEvents {
			recvch <- evt
		}
	}
}

// diffApp diffs local state with controller state
// FIXME: this is not handling deletes today
func (client *NimbusClient) diffApps(objList *netproto.AppList, reactor AppReactor, ostream *AppOStream) {
	// build a map of objects
	objmap := make(map[string]*netproto.App)
	for _, obj := range objList.Apps {
		key := obj.ObjectMeta.GetKey()
		objmap[key] = obj
	}

	// see if we need to delete any locally found object
	localObjs := reactor.ListApp()
	for _, lobj := range localObjs {
		ctby, ok := lobj.ObjectMeta.Labels["CreatedBy"]
		if ok && ctby == "Venice" {
			key := lobj.ObjectMeta.GetKey()
			if _, ok := objmap[key]; !ok {
				evt := netproto.AppEvent{
					EventType: api.EventType_DeleteEvent,
					App:       *lobj,
				}
				log.Infof("diffApps(): Deleting object %+v", lobj.ObjectMeta)
				client.lockObject(evt.App.GetObjectKind(), evt.App.ObjectMeta)
				client.processAppEvent(evt, reactor, ostream)
			}
		} else {
			log.Infof("Not deleting non-venice object %+v", lobj.ObjectMeta)
		}
	}

	// add/update all new objects
	for _, obj := range objList.Apps {
		evt := netproto.AppEvent{
			EventType: api.EventType_UpdateEvent,
			App:       *obj,
		}
		client.lockObject(evt.App.GetObjectKind(), evt.App.ObjectMeta)
		client.processAppEvent(evt, reactor, ostream)
	}
}

// processAppEvent handles App event
func (client *NimbusClient) processAppEvent(evt netproto.AppEvent, reactor AppReactor, ostream *AppOStream) {
	var err error
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// add venice label to the object
	evt.App.ObjectMeta.Labels = make(map[string]string)
	evt.App.ObjectMeta.Labels["CreatedBy"] = "Venice"

	// unlock the object once we are done
	defer client.unlockObject(evt.App.GetObjectKind(), evt.App.ObjectMeta)

	// retry till successful
	for iter := 0; iter < maxOpretry; iter++ {
		switch evt.EventType {
		case api.EventType_CreateEvent:
			fallthrough
		case api.EventType_UpdateEvent:
			_, err = reactor.FindApp(evt.App.ObjectMeta)
			if err != nil {
				// create the App
				err = reactor.CreateApp(&evt.App)
				if err != nil {
					log.Errorf("Error creating the App {%+v}. Err: %v", evt.App.ObjectMeta, err)
					client.debugStats.AddInt("CreateAppError", 1)
				} else {
					client.debugStats.AddInt("CreateApp", 1)
				}
			} else {
				// update the App
				err = reactor.UpdateApp(&evt.App)
				if err != nil {
					log.Errorf("Error updating the App {%+v}. Err: %v", evt.App.GetKey(), err)
					client.debugStats.AddInt("UpdateAppError", 1)
				} else {
					client.debugStats.AddInt("UpdateApp", 1)
				}
			}

		case api.EventType_DeleteEvent:
			// delete the object
			err = reactor.DeleteApp(evt.App.Tenant, evt.App.Namespace, evt.App.Name)
			if err == state.ErrObjectNotFound { // give idempotency to caller
				log.Debugf("App {%+v} not found", evt.App.ObjectMeta)
				err = nil
			}
			if err != nil {
				log.Errorf("Error deleting the App {%+v}. Err: %v", evt.App.ObjectMeta, err)
				client.debugStats.AddInt("DeleteAppError", 1)
			} else {
				client.debugStats.AddInt("DeleteApp", 1)
			}
		}

		// send oper status and return if there is no error
		if err == nil {
			robj := netproto.AppEvent{
				EventType: evt.EventType,
				App: netproto.App{
					TypeMeta:   evt.App.TypeMeta,
					ObjectMeta: evt.App.ObjectMeta,
					Status:     evt.App.Status,
				},
			}

			// send oper status
			ostream.Lock()
			modificationTime, _ := types.TimestampProto(time.Now())
			robj.App.ObjectMeta.ModTime = api.Timestamp{Timestamp: *modificationTime}
			err := ostream.stream.Send(&robj)
			if err != nil {
				log.Errorf("failed to send App oper Status, %s", err)
				client.debugStats.AddInt("AppOperSendError", 1)
			} else {
				client.debugStats.AddInt("AppOperSent", 1)
			}
			ostream.Unlock()

			return
		}

		// else, retry after some time, with backoff
		time.Sleep(time.Second * time.Duration(2*iter))
	}
}
