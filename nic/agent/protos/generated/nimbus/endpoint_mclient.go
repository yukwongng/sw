// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: endpoint.proto
*/

package nimbus

import (
	"context"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/state"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/utils/log"
	"google.golang.org/grpc/connectivity"
)

type EndpointReactor interface {
	CreateEndpoint(endpointObj *netproto.Endpoint) error          // creates an Endpoint
	FindEndpoint(meta api.ObjectMeta) (*netproto.Endpoint, error) // finds an Endpoint
	ListEndpoint() []*netproto.Endpoint                           // lists all Endpoints
	UpdateEndpoint(endpointObj *netproto.Endpoint) error          // updates an Endpoint
	DeleteEndpoint(endpointObj, ns, name string) error            // deletes an Endpoint
}

// WatchEndpoints runs Endpoint watcher loop
func (client *NimbusClient) WatchEndpoints(ctx context.Context, reactor EndpointReactor) {
	// setup wait group
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()
	client.debugStats.AddInt("ActiveEndpointWatch", 1)

	// make sure rpc client is good
	if client.rpcClient == nil || client.rpcClient.ClientConn == nil || client.rpcClient.ClientConn.GetState() != connectivity.Ready {
		log.Errorf("RPC client is disconnected. Exiting watch")
		return
	}

	// start the watch
	endpointRPCClient := netproto.NewEndpointApiClient(client.rpcClient.ClientConn)
	stream, err := endpointRPCClient.WatchEndpoints(ctx, &api.ObjectMeta{})
	if err != nil {
		log.Errorf("Error watching Endpoint. Err: %v", err)
		return
	}

	// start oper update stream
	ostream, err := endpointRPCClient.EndpointOperUpdate(ctx)
	if err != nil {
		log.Errorf("Error starting Endpoint oper updates. Err: %v", err)
		return
	}

	// get a list of objects
	objList, err := endpointRPCClient.ListEndpoints(ctx, &api.ObjectMeta{})
	if err != nil {
		log.Errorf("Error getting Endpoint list. Err: %v", err)
		return
	}

	// perform a diff of the states
	client.diffEndpoints(objList, reactor, ostream)

	// start grpc stream recv
	recvCh := make(chan *netproto.EndpointEvent, evChanLength)
	go client.watchEndpointRecvLoop(stream, recvCh)

	// loop till the end
	for {
		select {
		case evt, ok := <-recvCh:
			if !ok {
				log.Warnf("Endpoint Watch channel closed. Exisint EndpointWatch")
				return
			}
			client.debugStats.AddInt("EndpointWatchEvents", 1)

			log.Infof("Ctrlerif: agent %s got Endpoint watch event: Type: {%+v} Endpoint:{%+v}", client.clientName, evt.EventType, evt.Endpoint)

			client.lockObject(evt.Endpoint.GetObjectKind(), evt.Endpoint.ObjectMeta)
			client.processEndpointEvent(*evt, reactor, ostream)
		// periodic resync
		case <-time.After(resyncInterval):
			// get a list of objects
			objList, err := endpointRPCClient.ListEndpoints(ctx, &api.ObjectMeta{})
			if err != nil {
				log.Errorf("Error getting Endpoint list. Err: %v", err)
				return
			}
			client.debugStats.AddInt("EndpointWatchResyncs", 1)

			// perform a diff of the states
			client.diffEndpoints(objList, reactor, ostream)
		}
	}
}

// watchEndpointRecvLoop receives from stream and write it to a channel
func (client *NimbusClient) watchEndpointRecvLoop(stream netproto.EndpointApi_WatchEndpointsClient, recvch chan<- *netproto.EndpointEvent) {
	defer close(recvch)
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// loop till the end
	for {
		// receive from stream
		evt, err := stream.Recv()
		if err != nil {
			log.Errorf("Error receiving from watch channel. Exiting Endpoint watch. Err: %v", err)
			return
		}

		recvch <- evt
	}
}

// diffEndpoint diffs local state with controller state
// FIXME: this is not handling deletes today
func (client *NimbusClient) diffEndpoints(objList *netproto.EndpointList, reactor EndpointReactor, ostream netproto.EndpointApi_EndpointOperUpdateClient) {
	// build a map of objects
	objmap := make(map[string]*netproto.Endpoint)
	for _, obj := range objList.Endpoints {
		key := obj.ObjectMeta.GetKey()
		objmap[key] = obj
	}

	// see if we need to delete any locally found object
	localObjs := reactor.ListEndpoint()
	for _, lobj := range localObjs {
		ctby, ok := lobj.ObjectMeta.Labels["CreatedBy"]
		if ok && ctby == "Venice" {
			key := lobj.ObjectMeta.GetKey()
			if _, ok := objmap[key]; !ok {
				evt := netproto.EndpointEvent{
					EventType: api.EventType_DeleteEvent,
					Endpoint:  *lobj,
				}
				log.Infof("diffEndpoints(): Deleting object %+v", lobj.ObjectMeta)
				client.lockObject(evt.Endpoint.GetObjectKind(), evt.Endpoint.ObjectMeta)
				client.processEndpointEvent(evt, reactor, ostream)
			}
		} else {
			log.Infof("Not deleting non-venice object %+v", lobj.ObjectMeta)
		}
	}

	// add/update all new objects
	for _, obj := range objList.Endpoints {
		evt := netproto.EndpointEvent{
			EventType: api.EventType_CreateEvent,
			Endpoint:  *obj,
		}
		client.lockObject(evt.Endpoint.GetObjectKind(), evt.Endpoint.ObjectMeta)
		client.processEndpointEvent(evt, reactor, ostream)
	}
}

// processEndpointEvent handles Endpoint event
func (client *NimbusClient) processEndpointEvent(evt netproto.EndpointEvent, reactor EndpointReactor, ostream netproto.EndpointApi_EndpointOperUpdateClient) {
	var err error
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// add venice label to the object
	evt.Endpoint.ObjectMeta.Labels = make(map[string]string)
	evt.Endpoint.ObjectMeta.Labels["CreatedBy"] = "Venice"

	// unlock the object once we are done
	defer client.unlockObject(evt.Endpoint.GetObjectKind(), evt.Endpoint.ObjectMeta)

	// retry till successful
	for iter := 0; iter < maxOpretry; iter++ {
		switch evt.EventType {
		case api.EventType_CreateEvent:
			fallthrough
		case api.EventType_UpdateEvent:
			_, err = reactor.FindEndpoint(evt.Endpoint.ObjectMeta)
			if err != nil {
				// create the Endpoint
				err = reactor.CreateEndpoint(&evt.Endpoint)
				if err != nil {
					log.Errorf("Error creating the Endpoint {%+v}. Err: %v", evt.Endpoint, err)
					client.debugStats.AddInt("CreateEndpointError", 1)
				} else {
					client.debugStats.AddInt("CreateEndpoint", 1)
				}
			} else {
				// update the Endpoint
				err = reactor.UpdateEndpoint(&evt.Endpoint)
				if err != nil {
					log.Errorf("Error updating the Endpoint {%+v}. Err: %v", evt.Endpoint, err)
					client.debugStats.AddInt("UpdateEndpointError", 1)
				} else {
					client.debugStats.AddInt("UpdateEndpoint", 1)
				}
			}

		case api.EventType_DeleteEvent:
			// delete the object
			err = reactor.DeleteEndpoint(evt.Endpoint.Tenant, evt.Endpoint.Namespace, evt.Endpoint.Name)
			if err == state.ErrObjectNotFound { // give idempotency to caller
				log.Debugf("Endpoint {%+v} not found", evt.Endpoint.ObjectMeta)
				err = nil
			}
			if err != nil {
				log.Errorf("Error deleting the Endpoint {%+v}. Err: %v", evt.Endpoint.ObjectMeta, err)
				client.debugStats.AddInt("DeleteEndpointError", 1)
			} else {
				client.debugStats.AddInt("DeleteEndpoint", 1)
			}
		}

		// send oper status and return if there is no error
		if err == nil {
			robj := netproto.EndpointEvent{
				EventType: evt.EventType,
				Endpoint: netproto.Endpoint{
					TypeMeta:   evt.Endpoint.TypeMeta,
					ObjectMeta: evt.Endpoint.ObjectMeta,
					Status:     evt.Endpoint.Status,
				},
			}

			// send oper status
			err := ostream.Send(&robj)
			if err != nil {
				log.Errorf("failed to send Endpoint oper Status, %s", err)
				client.debugStats.AddInt("EndpointOperSendError", 1)
			} else {
				client.debugStats.AddInt("EndpointOperSent", 1)
			}

			return
		}

		// else, retry after some time, with backoff
		time.Sleep(time.Second * time.Duration(2*iter))
	}
}
