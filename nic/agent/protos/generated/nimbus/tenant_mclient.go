// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: tenant.proto
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

type TenantReactor interface {
	CreateTenant(tenantObj *netproto.Tenant) error            // creates an Tenant
	FindTenant(meta api.ObjectMeta) (*netproto.Tenant, error) // finds an Tenant
	ListTenant() []*netproto.Tenant                           // lists all Tenants
	UpdateTenant(tenantObj *netproto.Tenant) error            // updates an Tenant
	DeleteTenant(tenantObj, ns, name string) error            // deletes an Tenant
	GetWatchOptions(cts context.Context, kind string) api.ObjectMeta
}
type TenantOStream struct {
	sync.Mutex
	stream netproto.TenantApi_TenantOperUpdateClient
}

// WatchTenants runs Tenant watcher loop
func (client *NimbusClient) WatchTenants(ctx context.Context, reactor TenantReactor) {
	// setup wait group
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()
	client.debugStats.AddInt("ActiveTenantWatch", 1)

	// make sure rpc client is good
	if client.rpcClient == nil || client.rpcClient.ClientConn == nil || client.rpcClient.ClientConn.GetState() != connectivity.Ready {
		log.Errorf("RPC client is disconnected. Exiting watch")
		return
	}

	// start the watch
	ometa := reactor.GetWatchOptions(ctx, "Tenant")
	tenantRPCClient := netproto.NewTenantApiClient(client.rpcClient.ClientConn)
	stream, err := tenantRPCClient.WatchTenants(ctx, &ometa)
	if err != nil {
		log.Errorf("Error watching Tenant. Err: %v", err)
		return
	}

	// start oper update stream
	opStream, err := tenantRPCClient.TenantOperUpdate(ctx)
	if err != nil {
		log.Errorf("Error starting Tenant oper updates. Err: %v", err)
		return
	}

	ostream := &TenantOStream{stream: opStream}

	// get a list of objects
	objList, err := tenantRPCClient.ListTenants(ctx, &ometa)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok || st.Code() == codes.Unavailable {
			log.Errorf("Error getting Tenant list. Err: %v", err)
			return
		}
	} else {
		// perform a diff of the states
		client.diffTenants(objList, reactor, ostream)
	}

	// start grpc stream recv
	recvCh := make(chan *netproto.TenantEvent, evChanLength)
	go client.watchTenantRecvLoop(stream, recvCh)

	// loop till the end
	for {
		evtWork := func(evt *netproto.TenantEvent) {
			client.debugStats.AddInt("TenantWatchEvents", 1)
			log.Infof("Ctrlerif: agent %s got Tenant watch event: Type: {%+v} Tenant:{%+v}", client.clientName, evt.EventType, evt.Tenant.ObjectMeta)
			client.lockObject(evt.Tenant.GetObjectKind(), evt.Tenant.ObjectMeta)
			go client.processTenantEvent(*evt, reactor, ostream)
			//Give it some time to increment waitgrp
			time.Sleep(100 * time.Microsecond)
		}
		//Give priority to evnt work.
		select {
		case evt, ok := <-recvCh:
			if !ok {
				log.Warnf("Tenant Watch channel closed. Exisint TenantWatch")
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
					log.Warnf("Tenant Watch channel closed. Exisint TenantWatch")
					return
				}
				evtWork(evt)
				continue
			default:
			}
			// get a list of objects
			objList, err := tenantRPCClient.ListTenants(ctx, &ometa)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok || st.Code() == codes.Unavailable {
					log.Errorf("Error getting Tenant list. Err: %v", err)
					return
				}
			} else {
				client.debugStats.AddInt("TenantWatchResyncs", 1)
				// perform a diff of the states
				client.diffTenants(objList, reactor, ostream)
			}
		}
	}
}

// watchTenantRecvLoop receives from stream and write it to a channel
func (client *NimbusClient) watchTenantRecvLoop(stream netproto.TenantApi_WatchTenantsClient, recvch chan<- *netproto.TenantEvent) {
	defer close(recvch)
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// loop till the end
	for {
		// receive from stream
		objList, err := stream.Recv()
		if err != nil {
			log.Errorf("Error receiving from watch channel. Exiting Tenant watch. Err: %v", err)
			return
		}
		for _, evt := range objList.TenantEvents {
			recvch <- evt
		}
	}
}

// diffTenant diffs local state with controller state
// FIXME: this is not handling deletes today
func (client *NimbusClient) diffTenants(objList *netproto.TenantList, reactor TenantReactor, ostream *TenantOStream) {
	// build a map of objects
	objmap := make(map[string]*netproto.Tenant)
	for _, obj := range objList.Tenants {
		key := obj.ObjectMeta.GetKey()
		objmap[key] = obj
	}

	// see if we need to delete any locally found object
	localObjs := reactor.ListTenant()
	for _, lobj := range localObjs {
		ctby, ok := lobj.ObjectMeta.Labels["CreatedBy"]
		if ok && ctby == "Venice" {
			key := lobj.ObjectMeta.GetKey()
			if _, ok := objmap[key]; !ok {
				evt := netproto.TenantEvent{
					EventType: api.EventType_DeleteEvent,
					Tenant:    *lobj,
				}
				log.Infof("diffTenants(): Deleting object %+v", lobj.ObjectMeta)
				client.lockObject(evt.Tenant.GetObjectKind(), evt.Tenant.ObjectMeta)
				client.processTenantEvent(evt, reactor, ostream)
			}
		} else {
			log.Infof("Not deleting non-venice object %+v", lobj.ObjectMeta)
		}
	}

	// add/update all new objects
	for _, obj := range objList.Tenants {
		evt := netproto.TenantEvent{
			EventType: api.EventType_UpdateEvent,
			Tenant:    *obj,
		}
		client.lockObject(evt.Tenant.GetObjectKind(), evt.Tenant.ObjectMeta)
		client.processTenantEvent(evt, reactor, ostream)
	}
}

// processTenantEvent handles Tenant event
func (client *NimbusClient) processTenantEvent(evt netproto.TenantEvent, reactor TenantReactor, ostream *TenantOStream) {
	var err error
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// add venice label to the object
	evt.Tenant.ObjectMeta.Labels = make(map[string]string)
	evt.Tenant.ObjectMeta.Labels["CreatedBy"] = "Venice"

	// unlock the object once we are done
	defer client.unlockObject(evt.Tenant.GetObjectKind(), evt.Tenant.ObjectMeta)

	// retry till successful
	for iter := 0; iter < maxOpretry; iter++ {
		switch evt.EventType {
		case api.EventType_CreateEvent:
			fallthrough
		case api.EventType_UpdateEvent:
			_, err = reactor.FindTenant(evt.Tenant.ObjectMeta)
			if err != nil {
				// create the Tenant
				err = reactor.CreateTenant(&evt.Tenant)
				if err != nil {
					log.Errorf("Error creating the Tenant {%+v}. Err: %v", evt.Tenant.ObjectMeta, err)
					client.debugStats.AddInt("CreateTenantError", 1)
				} else {
					client.debugStats.AddInt("CreateTenant", 1)
				}
			} else {
				// update the Tenant
				err = reactor.UpdateTenant(&evt.Tenant)
				if err != nil {
					log.Errorf("Error updating the Tenant {%+v}. Err: %v", evt.Tenant.GetKey(), err)
					client.debugStats.AddInt("UpdateTenantError", 1)
				} else {
					client.debugStats.AddInt("UpdateTenant", 1)
				}
			}

		case api.EventType_DeleteEvent:
			// delete the object
			err = reactor.DeleteTenant(evt.Tenant.Tenant, evt.Tenant.Namespace, evt.Tenant.Name)
			if err == state.ErrObjectNotFound { // give idempotency to caller
				log.Debugf("Tenant {%+v} not found", evt.Tenant.ObjectMeta)
				err = nil
			}
			if err != nil {
				log.Errorf("Error deleting the Tenant {%+v}. Err: %v", evt.Tenant.ObjectMeta, err)
				client.debugStats.AddInt("DeleteTenantError", 1)
			} else {
				client.debugStats.AddInt("DeleteTenant", 1)
			}
		}

		// send oper status and return if there is no error
		if err == nil {
			robj := netproto.TenantEvent{
				EventType: evt.EventType,
				Tenant: netproto.Tenant{
					TypeMeta:   evt.Tenant.TypeMeta,
					ObjectMeta: evt.Tenant.ObjectMeta,
					Status:     evt.Tenant.Status,
				},
			}

			// send oper status
			ostream.Lock()
			modificationTime, _ := types.TimestampProto(time.Now())
			robj.Tenant.ObjectMeta.ModTime = api.Timestamp{Timestamp: *modificationTime}
			err := ostream.stream.Send(&robj)
			if err != nil {
				log.Errorf("failed to send Tenant oper Status, %s", err)
				client.debugStats.AddInt("TenantOperSendError", 1)
			} else {
				client.debugStats.AddInt("TenantOperSent", 1)
			}
			ostream.Unlock()

			return
		}

		// else, retry after some time, with backoff
		time.Sleep(time.Second * time.Duration(2*iter))
	}
}
