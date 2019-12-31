// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: route.proto
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

type RoutingConfigReactor interface {
	CreateRoutingConfig(routingconfigObj *netproto.RoutingConfig) error     // creates an RoutingConfig
	FindRoutingConfig(meta api.ObjectMeta) (*netproto.RoutingConfig, error) // finds an RoutingConfig
	ListRoutingConfig() []*netproto.RoutingConfig                           // lists all RoutingConfigs
	UpdateRoutingConfig(routingconfigObj *netproto.RoutingConfig) error     // updates an RoutingConfig
	DeleteRoutingConfig(routingconfigObj, ns, name string) error            // deletes an RoutingConfig
	GetWatchOptions(cts context.Context, kind string) api.ListWatchOptions
}
type RoutingConfigOStream struct {
	sync.Mutex
	stream netproto.RoutingConfigApiV1_RoutingConfigOperUpdateClient
}

// WatchRoutingConfigs runs RoutingConfig watcher loop
func (client *NimbusClient) WatchRoutingConfigs(ctx context.Context, reactor RoutingConfigReactor) {
	// setup wait group
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()
	client.debugStats.AddInt("ActiveRoutingConfigWatch", 1)

	// make sure rpc client is good
	if client.rpcClient == nil || client.rpcClient.ClientConn == nil || client.rpcClient.ClientConn.GetState() != connectivity.Ready {
		log.Errorf("RPC client is disconnected. Exiting watch")
		return
	}

	// start the watch
	watchOptions := reactor.GetWatchOptions(ctx, "RoutingConfig")
	routingconfigRPCClient := netproto.NewRoutingConfigApiV1Client(client.rpcClient.ClientConn)
	stream, err := routingconfigRPCClient.WatchRoutingConfigs(ctx, &watchOptions)
	if err != nil {
		log.Errorf("Error watching RoutingConfig. Err: %v", err)
		return
	}

	// start oper update stream
	opStream, err := routingconfigRPCClient.RoutingConfigOperUpdate(ctx)
	if err != nil {
		log.Errorf("Error starting RoutingConfig oper updates. Err: %v", err)
		return
	}

	ostream := &RoutingConfigOStream{stream: opStream}

	// get a list of objects
	objList, err := routingconfigRPCClient.ListRoutingConfigs(ctx, &watchOptions)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok || st.Code() == codes.Unavailable {
			log.Errorf("Error getting RoutingConfig list. Err: %v", err)
			return
		}
	} else {
		// perform a diff of the states
		client.diffRoutingConfigs(objList, reactor, ostream)
	}

	// start grpc stream recv
	recvCh := make(chan *netproto.RoutingConfigEvent, evChanLength)
	go client.watchRoutingConfigRecvLoop(stream, recvCh)

	// loop till the end
	for {
		evtWork := func(evt *netproto.RoutingConfigEvent) {
			client.debugStats.AddInt("RoutingConfigWatchEvents", 1)
			log.Infof("Ctrlerif: agent %s got RoutingConfig watch event: Type: {%+v} RoutingConfig:{%+v}", client.clientName, evt.EventType, evt.RoutingConfig.ObjectMeta)
			client.lockObject(evt.RoutingConfig.GetObjectKind(), evt.RoutingConfig.ObjectMeta)
			go client.processRoutingConfigEvent(*evt, reactor, ostream)
			//Give it some time to increment waitgrp
			time.Sleep(100 * time.Microsecond)
		}
		//Give priority to evnt work.
		select {
		case evt, ok := <-recvCh:
			if !ok {
				log.Warnf("RoutingConfig Watch channel closed. Exisint RoutingConfigWatch")
				return
			}
			evtWork(evt)
		// periodic resync (Disabling as we have aggregate watch support)
		case <-time.After(resyncInterval):
			//Give priority to evt work
			//Wait for batch interval for inflight work
			time.Sleep(5 * DefaultWatchHoldInterval)
			select {
			case evt, ok := <-recvCh:
				if !ok {
					log.Warnf("RoutingConfig Watch channel closed. Exisint RoutingConfigWatch")
					return
				}
				evtWork(evt)
				continue
			default:
			}
			// get a list of objects
			objList, err := routingconfigRPCClient.ListRoutingConfigs(ctx, &watchOptions)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok || st.Code() == codes.Unavailable {
					log.Errorf("Error getting RoutingConfig list. Err: %v", err)
					return
				}
			} else {
				client.debugStats.AddInt("RoutingConfigWatchResyncs", 1)
				// perform a diff of the states
				client.diffRoutingConfigs(objList, reactor, ostream)
			}
		}
	}
}

// watchRoutingConfigRecvLoop receives from stream and write it to a channel
func (client *NimbusClient) watchRoutingConfigRecvLoop(stream netproto.RoutingConfigApiV1_WatchRoutingConfigsClient, recvch chan<- *netproto.RoutingConfigEvent) {
	defer close(recvch)
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// loop till the end
	for {
		// receive from stream
		objList, err := stream.Recv()
		if err != nil {
			log.Errorf("Error receiving from watch channel. Exiting RoutingConfig watch. Err: %v", err)
			return
		}
		for _, evt := range objList.RoutingConfigEvents {
			recvch <- evt
		}
	}
}

// diffRoutingConfig diffs local state with controller state
// FIXME: this is not handling deletes today
func (client *NimbusClient) diffRoutingConfigs(objList *netproto.RoutingConfigList, reactor RoutingConfigReactor, ostream *RoutingConfigOStream) {
	// build a map of objects
	objmap := make(map[string]*netproto.RoutingConfig)
	for _, obj := range objList.RoutingConfigs {
		key := obj.ObjectMeta.GetKey()
		objmap[key] = obj
	}

	// see if we need to delete any locally found object
	localObjs := reactor.ListRoutingConfig()
	for _, lobj := range localObjs {
		ctby, ok := lobj.ObjectMeta.Labels["CreatedBy"]
		if ok && ctby == "Venice" {
			key := lobj.ObjectMeta.GetKey()
			if _, ok := objmap[key]; !ok {
				evt := netproto.RoutingConfigEvent{
					EventType:     api.EventType_DeleteEvent,
					RoutingConfig: *lobj,
				}
				log.Infof("diffRoutingConfigs(): Deleting object %+v", lobj.ObjectMeta)
				client.lockObject(evt.RoutingConfig.GetObjectKind(), evt.RoutingConfig.ObjectMeta)
				client.processRoutingConfigEvent(evt, reactor, ostream)
			}
		} else {
			log.Infof("Not deleting non-venice object %+v", lobj.ObjectMeta)
		}
	}

	// add/update all new objects
	for _, obj := range objList.RoutingConfigs {
		evt := netproto.RoutingConfigEvent{
			EventType:     api.EventType_UpdateEvent,
			RoutingConfig: *obj,
		}
		client.lockObject(evt.RoutingConfig.GetObjectKind(), evt.RoutingConfig.ObjectMeta)
		client.processRoutingConfigEvent(evt, reactor, ostream)
	}
}

// processRoutingConfigEvent handles RoutingConfig event
func (client *NimbusClient) processRoutingConfigEvent(evt netproto.RoutingConfigEvent, reactor RoutingConfigReactor, ostream *RoutingConfigOStream) error {
	var err error
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// add venice label to the object
	evt.RoutingConfig.ObjectMeta.Labels = make(map[string]string)
	evt.RoutingConfig.ObjectMeta.Labels["CreatedBy"] = "Venice"

	// unlock the object once we are done
	defer client.unlockObject(evt.RoutingConfig.GetObjectKind(), evt.RoutingConfig.ObjectMeta)

	// retry till successful
	for iter := 0; iter < maxOpretry; iter++ {
		switch evt.EventType {
		case api.EventType_CreateEvent:
			fallthrough
		case api.EventType_UpdateEvent:
			_, err = reactor.FindRoutingConfig(evt.RoutingConfig.ObjectMeta)
			if err != nil {
				// create the RoutingConfig
				err = reactor.CreateRoutingConfig(&evt.RoutingConfig)
				if err != nil {
					log.Errorf("Error creating the RoutingConfig {%+v}. Err: %v", evt.RoutingConfig.ObjectMeta, err)
					client.debugStats.AddInt("CreateRoutingConfigError", 1)
				} else {
					client.debugStats.AddInt("CreateRoutingConfig", 1)
				}
			} else {
				// update the RoutingConfig
				err = reactor.UpdateRoutingConfig(&evt.RoutingConfig)
				if err != nil {
					log.Errorf("Error updating the RoutingConfig {%+v}. Err: %v", evt.RoutingConfig.GetKey(), err)
					client.debugStats.AddInt("UpdateRoutingConfigError", 1)
				} else {
					client.debugStats.AddInt("UpdateRoutingConfig", 1)
				}
			}

		case api.EventType_DeleteEvent:
			// delete the object
			err = reactor.DeleteRoutingConfig(evt.RoutingConfig.Tenant, evt.RoutingConfig.Namespace, evt.RoutingConfig.Name)
			if err == state.ErrObjectNotFound { // give idempotency to caller
				log.Debugf("RoutingConfig {%+v} not found", evt.RoutingConfig.ObjectMeta)
				err = nil
			}
			if err != nil {
				log.Errorf("Error deleting the RoutingConfig {%+v}. Err: %v", evt.RoutingConfig.ObjectMeta, err)
				client.debugStats.AddInt("DeleteRoutingConfigError", 1)
			} else {
				client.debugStats.AddInt("DeleteRoutingConfig", 1)
			}
		}

		if ostream == nil {
			return err
		}
		// send oper status and return if there is no error
		if err == nil {
			robj := netproto.RoutingConfigEvent{
				EventType: evt.EventType,
				RoutingConfig: netproto.RoutingConfig{
					TypeMeta:   evt.RoutingConfig.TypeMeta,
					ObjectMeta: evt.RoutingConfig.ObjectMeta,
					Status:     evt.RoutingConfig.Status,
				},
			}

			// send oper status
			ostream.Lock()
			modificationTime, _ := types.TimestampProto(time.Now())
			robj.RoutingConfig.ObjectMeta.ModTime = api.Timestamp{Timestamp: *modificationTime}
			err := ostream.stream.Send(&robj)
			if err != nil {
				log.Errorf("failed to send RoutingConfig oper Status, %s", err)
				client.debugStats.AddInt("RoutingConfigOperSendError", 1)
			} else {
				client.debugStats.AddInt("RoutingConfigOperSent", 1)
			}
			ostream.Unlock()

			return err
		}

		// else, retry after some time, with backoff
		time.Sleep(time.Second * time.Duration(2*iter))
	}

	return nil
}

func (client *NimbusClient) processRoutingConfigDynamic(evt api.EventType,
	object *netproto.RoutingConfig, reactor RoutingConfigReactor) error {

	routingconfigEvt := netproto.RoutingConfigEvent{
		EventType:     evt,
		RoutingConfig: *object,
	}

	// add venice label to the object
	routingconfigEvt.RoutingConfig.ObjectMeta.Labels = make(map[string]string)
	routingconfigEvt.RoutingConfig.ObjectMeta.Labels["CreatedBy"] = "Venice"

	client.lockObject(routingconfigEvt.RoutingConfig.GetObjectKind(), routingconfigEvt.RoutingConfig.ObjectMeta)

	err := client.processRoutingConfigEvent(routingconfigEvt, reactor, nil)
	modificationTime, _ := types.TimestampProto(time.Now())
	object.ObjectMeta.ModTime = api.Timestamp{Timestamp: *modificationTime}

	return err
}
