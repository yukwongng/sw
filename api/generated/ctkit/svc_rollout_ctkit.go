// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package ctkit is a auto generated package.
Input file: svc_rollout.proto
*/
package ctkit

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/rollout"
	"github.com/pensando/sw/venice/utils/balancer"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Rollout is a wrapper object that implements additional functionality
type Rollout struct {
	sync.Mutex
	rollout.Rollout
	HandlerCtx interface{} // additional state handlers can store
	ctrler     *ctrlerCtx  // reference back to the controller instance
}

func (obj *Rollout) Write() error {
	// if there is no API server to connect to, we are done
	if (obj.ctrler == nil) || (obj.ctrler.resolver == nil) || obj.ctrler.apisrvURL == "" {
		return nil
	}

	apicl, err := obj.ctrler.apiClient()
	if err != nil {
		obj.ctrler.logger.Errorf("Error creating API server clent. Err: %v", err)
		return err
	}

	obj.ctrler.stats.Counter("Rollout_Writes").Inc()

	// write to api server
	if obj.ObjectMeta.ResourceVersion != "" {
		// update it
		for i := 0; i < maxApisrvWriteRetry; i++ {
			_, err = apicl.RolloutV1().Rollout().UpdateStatus(context.Background(), &obj.Rollout)
			if err == nil {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}
	} else {
		//  create
		_, err = apicl.RolloutV1().Rollout().Create(context.Background(), &obj.Rollout)
	}

	return nil
}

// RolloutHandler is the event handler for Rollout object
type RolloutHandler interface {
	OnRolloutCreate(obj *Rollout) error
	OnRolloutUpdate(oldObj *Rollout, newObj *rollout.Rollout) error
	OnRolloutDelete(obj *Rollout) error
}

// handleRolloutEvent handles Rollout events from watcher
func (ct *ctrlerCtx) handleRolloutEvent(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *rollout.Rollout:
		eobj := evt.Object.(*rollout.Rollout)
		kind := "Rollout"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		rolloutHandler := handler.(RolloutHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &Rollout{
					Rollout:    *eobj,
					HandlerCtx: nil,
					ctrler:     ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("Rollout_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				err = rolloutHandler.OnRolloutCreate(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					ct.delObject(kind, eobj.GetKey())
					return err
				}
			} else {
				obj := fobj.(*Rollout)

				ct.stats.Counter("Rollout_Updated_Events").Inc()

				// call the event handler
				obj.Lock()
				err = rolloutHandler.OnRolloutUpdate(obj, eobj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					return err
				}
			}
		case kvstore.Deleted:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				ct.logger.Errorf("Object %s/%s not found durng delete. Err: %v", kind, eobj.GetKey(), err)
				return err
			}

			obj := fobj.(*Rollout)

			ct.stats.Counter("Rollout_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			err = rolloutHandler.OnRolloutDelete(obj)
			obj.Unlock()
			if err != nil {
				ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
			}

			ct.delObject(kind, eobj.GetKey())
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on Rollout watch channel", tp)
	}

	return nil
}

// handleRolloutEventParallel handles Rollout events from watcher
func (ct *ctrlerCtx) handleRolloutEventParallel(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *rollout.Rollout:
		eobj := evt.Object.(*rollout.Rollout)
		kind := "Rollout"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		rolloutHandler := handler.(RolloutHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &Rollout{
					Rollout:    *eobj,
					HandlerCtx: nil,
					ctrler:     ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("Rollout_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				go func() {
					err = rolloutHandler.OnRolloutCreate(obj)
					obj.Unlock()
					if err != nil {
						ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
						ct.delObject(kind, eobj.GetKey())
					}
				}()
			} else {
				obj := fobj.(*Rollout)

				ct.stats.Counter("Rollout_Updated_Events").Inc()

				// call the event handler
				obj.Lock()
				go func() {
					err = rolloutHandler.OnRolloutUpdate(obj, eobj)
					obj.Unlock()
					if err != nil {
						ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					}
				}()
			}
		case kvstore.Deleted:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				ct.logger.Errorf("Object %s/%s not found durng delete. Err: %v", kind, eobj.GetKey(), err)
				return err
			}

			obj := fobj.(*Rollout)

			ct.stats.Counter("Rollout_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			go func() {
				err = rolloutHandler.OnRolloutDelete(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
				}

				ct.delObject(kind, eobj.GetKey())
			}()
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on Rollout watch channel", tp)
	}

	return nil
}

// diffRollout does a diff of Rollout objects between local cache and API server
func (ct *ctrlerCtx) diffRollout(apicl apiclient.Services) {
	opts := api.ListWatchOptions{}

	// get a list of all objects from API server
	objlist, err := apicl.RolloutV1().Rollout().List(context.Background(), &opts)
	if err != nil {
		ct.logger.Errorf("Error getting a list of objects. Err: %v", err)
		return
	}

	ct.logger.Infof("diffRollout(): RolloutList returned %d objects", len(objlist))

	// build an object map
	objmap := make(map[string]*rollout.Rollout)
	for _, obj := range objlist {
		objmap[obj.GetKey()] = obj
	}

	// if an object is in our local cache and not in API server, trigger delete for it
	for _, obj := range ct.Rollout().List() {
		_, ok := objmap[obj.GetKey()]
		if !ok {
			ct.logger.Infof("diffRollout(): Deleting existing object %#v since its not in apiserver", obj.GetKey())
			evt := kvstore.WatchEvent{
				Type:   kvstore.Deleted,
				Key:    obj.GetKey(),
				Object: &obj.Rollout,
			}
			ct.handleRolloutEvent(&evt)
		}
	}

	// trigger create event for all others
	for _, obj := range objlist {
		ct.logger.Infof("diffRollout(): Adding object %#v", obj.GetKey())
		evt := kvstore.WatchEvent{
			Type:   kvstore.Created,
			Key:    obj.GetKey(),
			Object: obj,
		}
		ct.handleRolloutEvent(&evt)
	}
}

func (ct *ctrlerCtx) runRolloutWatcher() {
	kind := "Rollout"

	// if there is no API server to connect to, we are done
	if (ct.resolver == nil) || ct.apisrvURL == "" {
		return
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	ct.Lock()
	ct.watchCancel[kind] = cancel
	ct.Unlock()
	opts := api.ListWatchOptions{}
	logger := ct.logger.WithContext("submodule", "RolloutWatcher")

	// create a grpc client
	apiclt, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
	if err == nil {
		ct.diffRollout(apiclt)
	}

	// setup wait group
	ct.waitGrp.Add(1)

	// start a goroutine
	go func() {
		defer ct.waitGrp.Done()
		ct.stats.Counter("Rollout_Watch").Inc()
		defer ct.stats.Counter("Rollout_Watch").Dec()

		// loop forever
		for {
			// create a grpc client
			apicl, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
			if err != nil {
				logger.Warnf("Failed to connect to gRPC server [%s]\n", ct.apisrvURL)
				ct.stats.Counter("Rollout_ApiClientErr").Inc()
			} else {
				logger.Infof("API client connected {%+v}", apicl)

				// Rollout object watcher
				wt, werr := apicl.RolloutV1().Rollout().Watch(ctx, &opts)
				if werr != nil {
					logger.Errorf("Failed to start %s watch (%s)\n", kind, werr)
					// wait for a second and retry connecting to api server
					apicl.Close()
					time.Sleep(time.Second)
					continue
				}
				ct.Lock()
				ct.watchers[kind] = wt
				ct.Unlock()

				// perform a diff with API server and local cache
				time.Sleep(time.Millisecond * 100)
				ct.diffRollout(apicl)

				// handle api server watch events
			innerLoop:
				for {
					// wait for events
					select {
					case evt, ok := <-wt.EventChan():
						if !ok {
							logger.Error("Error receiving from apisrv watcher")
							ct.stats.Counter("Rollout_WatchErrors").Inc()
							break innerLoop
						}

						// handle event in parallel
						ct.handleRolloutEventParallel(evt)
					}
				}
				apicl.Close()
			}

			// if stop flag is set, we are done
			if ct.stoped {
				logger.Infof("Exiting API server watcher")
				return
			}

			// wait for a second and retry connecting to api server
			time.Sleep(time.Second)
		}
	}()
}

// WatchRollout starts watch on Rollout object
func (ct *ctrlerCtx) WatchRollout(handler RolloutHandler) error {
	kind := "Rollout"

	// see if we already have a watcher
	ct.Lock()
	_, ok := ct.watchers[kind]
	ct.Unlock()
	if ok {
		return fmt.Errorf("Rollout watcher already exists")
	}

	// save handler
	ct.Lock()
	ct.handlers[kind] = handler
	ct.Unlock()

	// run Rollout watcher in a go routine
	ct.runRolloutWatcher()

	return nil
}

// RolloutAPI returns
type RolloutAPI interface {
	Create(obj *rollout.Rollout) error
	Update(obj *rollout.Rollout) error
	Delete(obj *rollout.Rollout) error
	Find(meta *api.ObjectMeta) (*Rollout, error)
	List() []*Rollout
	Watch(handler RolloutHandler) error
}

// dummy struct that implements RolloutAPI
type rolloutAPI struct {
	ct *ctrlerCtx
}

// Create creates Rollout object
func (api *rolloutAPI) Create(obj *rollout.Rollout) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.RolloutV1().Rollout().Create(context.Background(), obj)
		if err != nil && strings.Contains(err.Error(), "AlreadyExists") {
			_, err = apicl.RolloutV1().Rollout().Update(context.Background(), obj)
		}
		if err != nil {
			return err
		}
	}

	return api.ct.handleRolloutEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Created})
}

// Update triggers update on Rollout object
func (api *rolloutAPI) Update(obj *rollout.Rollout) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.RolloutV1().Rollout().Update(context.Background(), obj)
		if err != nil {
			return err
		}
	}

	return api.ct.handleRolloutEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Updated})
}

// Delete deletes Rollout object
func (api *rolloutAPI) Delete(obj *rollout.Rollout) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		apicl.RolloutV1().Rollout().Delete(context.Background(), &obj.ObjectMeta)
	}

	return api.ct.handleRolloutEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Deleted})
}

// Find returns an object by meta
func (api *rolloutAPI) Find(meta *api.ObjectMeta) (*Rollout, error) {
	// find the object
	obj, err := api.ct.FindObject("Rollout", meta)
	if err != nil {
		return nil, err
	}

	// asset type
	switch obj.(type) {
	case *Rollout:
		hobj := obj.(*Rollout)
		return hobj, nil
	default:
		return nil, errors.New("incorrect object type")
	}
}

// List returns a list of all Rollout objects
func (api *rolloutAPI) List() []*Rollout {
	var objlist []*Rollout

	objs := api.ct.ListObjects("Rollout")
	for _, obj := range objs {
		switch tp := obj.(type) {
		case *Rollout:
			eobj := obj.(*Rollout)
			objlist = append(objlist, eobj)
		default:
			log.Fatalf("Got invalid object type %v while looking for Rollout", tp)
		}
	}

	return objlist
}

// Watch sets up a event handlers for Rollout object
func (api *rolloutAPI) Watch(handler RolloutHandler) error {
	return api.ct.WatchRollout(handler)
}

// Rollout returns RolloutAPI
func (ct *ctrlerCtx) Rollout() RolloutAPI {
	return &rolloutAPI{ct: ct}
}

// RolloutAction is a wrapper object that implements additional functionality
type RolloutAction struct {
	sync.Mutex
	rollout.RolloutAction
	HandlerCtx interface{} // additional state handlers can store
	ctrler     *ctrlerCtx  // reference back to the controller instance
}

func (obj *RolloutAction) Write() error {
	// if there is no API server to connect to, we are done
	if (obj.ctrler == nil) || (obj.ctrler.resolver == nil) || obj.ctrler.apisrvURL == "" {
		return nil
	}

	apicl, err := obj.ctrler.apiClient()
	if err != nil {
		obj.ctrler.logger.Errorf("Error creating API server clent. Err: %v", err)
		return err
	}

	obj.ctrler.stats.Counter("RolloutAction_Writes").Inc()

	// write to api server
	if obj.ObjectMeta.ResourceVersion != "" {
		// update it
		for i := 0; i < maxApisrvWriteRetry; i++ {
			_, err = apicl.RolloutV1().RolloutAction().UpdateStatus(context.Background(), &obj.RolloutAction)
			if err == nil {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}
	} else {
		//  create
		_, err = apicl.RolloutV1().RolloutAction().Create(context.Background(), &obj.RolloutAction)
	}

	return nil
}

// RolloutActionHandler is the event handler for RolloutAction object
type RolloutActionHandler interface {
	OnRolloutActionCreate(obj *RolloutAction) error
	OnRolloutActionUpdate(oldObj *RolloutAction, newObj *rollout.RolloutAction) error
	OnRolloutActionDelete(obj *RolloutAction) error
}

// handleRolloutActionEvent handles RolloutAction events from watcher
func (ct *ctrlerCtx) handleRolloutActionEvent(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *rollout.RolloutAction:
		eobj := evt.Object.(*rollout.RolloutAction)
		kind := "RolloutAction"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		rolloutactionHandler := handler.(RolloutActionHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &RolloutAction{
					RolloutAction: *eobj,
					HandlerCtx:    nil,
					ctrler:        ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("RolloutAction_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				err = rolloutactionHandler.OnRolloutActionCreate(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					ct.delObject(kind, eobj.GetKey())
					return err
				}
			} else {
				obj := fobj.(*RolloutAction)

				ct.stats.Counter("RolloutAction_Updated_Events").Inc()

				// call the event handler
				obj.Lock()
				err = rolloutactionHandler.OnRolloutActionUpdate(obj, eobj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					return err
				}
			}
		case kvstore.Deleted:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				ct.logger.Errorf("Object %s/%s not found durng delete. Err: %v", kind, eobj.GetKey(), err)
				return err
			}

			obj := fobj.(*RolloutAction)

			ct.stats.Counter("RolloutAction_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			err = rolloutactionHandler.OnRolloutActionDelete(obj)
			obj.Unlock()
			if err != nil {
				ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
			}

			ct.delObject(kind, eobj.GetKey())
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on RolloutAction watch channel", tp)
	}

	return nil
}

// handleRolloutActionEventParallel handles RolloutAction events from watcher
func (ct *ctrlerCtx) handleRolloutActionEventParallel(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *rollout.RolloutAction:
		eobj := evt.Object.(*rollout.RolloutAction)
		kind := "RolloutAction"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		rolloutactionHandler := handler.(RolloutActionHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &RolloutAction{
					RolloutAction: *eobj,
					HandlerCtx:    nil,
					ctrler:        ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("RolloutAction_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				go func() {
					err = rolloutactionHandler.OnRolloutActionCreate(obj)
					obj.Unlock()
					if err != nil {
						ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
						ct.delObject(kind, eobj.GetKey())
					}
				}()
			} else {
				obj := fobj.(*RolloutAction)

				ct.stats.Counter("RolloutAction_Updated_Events").Inc()

				// call the event handler
				obj.Lock()
				go func() {
					err = rolloutactionHandler.OnRolloutActionUpdate(obj, eobj)
					obj.Unlock()
					if err != nil {
						ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					}
				}()
			}
		case kvstore.Deleted:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				ct.logger.Errorf("Object %s/%s not found durng delete. Err: %v", kind, eobj.GetKey(), err)
				return err
			}

			obj := fobj.(*RolloutAction)

			ct.stats.Counter("RolloutAction_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			go func() {
				err = rolloutactionHandler.OnRolloutActionDelete(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
				}

				ct.delObject(kind, eobj.GetKey())
			}()
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on RolloutAction watch channel", tp)
	}

	return nil
}

// diffRolloutAction does a diff of RolloutAction objects between local cache and API server
func (ct *ctrlerCtx) diffRolloutAction(apicl apiclient.Services) {
	opts := api.ListWatchOptions{}

	// get a list of all objects from API server
	objlist, err := apicl.RolloutV1().RolloutAction().List(context.Background(), &opts)
	if err != nil {
		ct.logger.Errorf("Error getting a list of objects. Err: %v", err)
		return
	}

	ct.logger.Infof("diffRolloutAction(): RolloutActionList returned %d objects", len(objlist))

	// build an object map
	objmap := make(map[string]*rollout.RolloutAction)
	for _, obj := range objlist {
		objmap[obj.GetKey()] = obj
	}

	// if an object is in our local cache and not in API server, trigger delete for it
	for _, obj := range ct.RolloutAction().List() {
		_, ok := objmap[obj.GetKey()]
		if !ok {
			ct.logger.Infof("diffRolloutAction(): Deleting existing object %#v since its not in apiserver", obj.GetKey())
			evt := kvstore.WatchEvent{
				Type:   kvstore.Deleted,
				Key:    obj.GetKey(),
				Object: &obj.RolloutAction,
			}
			ct.handleRolloutActionEvent(&evt)
		}
	}

	// trigger create event for all others
	for _, obj := range objlist {
		ct.logger.Infof("diffRolloutAction(): Adding object %#v", obj.GetKey())
		evt := kvstore.WatchEvent{
			Type:   kvstore.Created,
			Key:    obj.GetKey(),
			Object: obj,
		}
		ct.handleRolloutActionEvent(&evt)
	}
}

func (ct *ctrlerCtx) runRolloutActionWatcher() {
	kind := "RolloutAction"

	// if there is no API server to connect to, we are done
	if (ct.resolver == nil) || ct.apisrvURL == "" {
		return
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	ct.Lock()
	ct.watchCancel[kind] = cancel
	ct.Unlock()
	opts := api.ListWatchOptions{}
	logger := ct.logger.WithContext("submodule", "RolloutActionWatcher")

	// create a grpc client
	apiclt, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
	if err == nil {
		ct.diffRolloutAction(apiclt)
	}

	// setup wait group
	ct.waitGrp.Add(1)

	// start a goroutine
	go func() {
		defer ct.waitGrp.Done()
		ct.stats.Counter("RolloutAction_Watch").Inc()
		defer ct.stats.Counter("RolloutAction_Watch").Dec()

		// loop forever
		for {
			// create a grpc client
			apicl, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
			if err != nil {
				logger.Warnf("Failed to connect to gRPC server [%s]\n", ct.apisrvURL)
				ct.stats.Counter("RolloutAction_ApiClientErr").Inc()
			} else {
				logger.Infof("API client connected {%+v}", apicl)

				// RolloutAction object watcher
				wt, werr := apicl.RolloutV1().RolloutAction().Watch(ctx, &opts)
				if werr != nil {
					logger.Errorf("Failed to start %s watch (%s)\n", kind, werr)
					// wait for a second and retry connecting to api server
					apicl.Close()
					time.Sleep(time.Second)
					continue
				}
				ct.Lock()
				ct.watchers[kind] = wt
				ct.Unlock()

				// perform a diff with API server and local cache
				time.Sleep(time.Millisecond * 100)
				ct.diffRolloutAction(apicl)

				// handle api server watch events
			innerLoop:
				for {
					// wait for events
					select {
					case evt, ok := <-wt.EventChan():
						if !ok {
							logger.Error("Error receiving from apisrv watcher")
							ct.stats.Counter("RolloutAction_WatchErrors").Inc()
							break innerLoop
						}

						// handle event in parallel
						ct.handleRolloutActionEventParallel(evt)
					}
				}
				apicl.Close()
			}

			// if stop flag is set, we are done
			if ct.stoped {
				logger.Infof("Exiting API server watcher")
				return
			}

			// wait for a second and retry connecting to api server
			time.Sleep(time.Second)
		}
	}()
}

// WatchRolloutAction starts watch on RolloutAction object
func (ct *ctrlerCtx) WatchRolloutAction(handler RolloutActionHandler) error {
	kind := "RolloutAction"

	// see if we already have a watcher
	ct.Lock()
	_, ok := ct.watchers[kind]
	ct.Unlock()
	if ok {
		return fmt.Errorf("RolloutAction watcher already exists")
	}

	// save handler
	ct.Lock()
	ct.handlers[kind] = handler
	ct.Unlock()

	// run RolloutAction watcher in a go routine
	ct.runRolloutActionWatcher()

	return nil
}

// RolloutActionAPI returns
type RolloutActionAPI interface {
	Create(obj *rollout.RolloutAction) error
	Update(obj *rollout.RolloutAction) error
	Delete(obj *rollout.RolloutAction) error
	Find(meta *api.ObjectMeta) (*RolloutAction, error)
	List() []*RolloutAction
	Watch(handler RolloutActionHandler) error
}

// dummy struct that implements RolloutActionAPI
type rolloutactionAPI struct {
	ct *ctrlerCtx
}

// Create creates RolloutAction object
func (api *rolloutactionAPI) Create(obj *rollout.RolloutAction) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.RolloutV1().RolloutAction().Create(context.Background(), obj)
		if err != nil && strings.Contains(err.Error(), "AlreadyExists") {
			_, err = apicl.RolloutV1().RolloutAction().Update(context.Background(), obj)
		}
		if err != nil {
			return err
		}
	}

	return api.ct.handleRolloutActionEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Created})
}

// Update triggers update on RolloutAction object
func (api *rolloutactionAPI) Update(obj *rollout.RolloutAction) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.RolloutV1().RolloutAction().Update(context.Background(), obj)
		if err != nil {
			return err
		}
	}

	return api.ct.handleRolloutActionEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Updated})
}

// Delete deletes RolloutAction object
func (api *rolloutactionAPI) Delete(obj *rollout.RolloutAction) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		apicl.RolloutV1().RolloutAction().Delete(context.Background(), &obj.ObjectMeta)
	}

	return api.ct.handleRolloutActionEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Deleted})
}

// Find returns an object by meta
func (api *rolloutactionAPI) Find(meta *api.ObjectMeta) (*RolloutAction, error) {
	// find the object
	obj, err := api.ct.FindObject("RolloutAction", meta)
	if err != nil {
		return nil, err
	}

	// asset type
	switch obj.(type) {
	case *RolloutAction:
		hobj := obj.(*RolloutAction)
		return hobj, nil
	default:
		return nil, errors.New("incorrect object type")
	}
}

// List returns a list of all RolloutAction objects
func (api *rolloutactionAPI) List() []*RolloutAction {
	var objlist []*RolloutAction

	objs := api.ct.ListObjects("RolloutAction")
	for _, obj := range objs {
		switch tp := obj.(type) {
		case *RolloutAction:
			eobj := obj.(*RolloutAction)
			objlist = append(objlist, eobj)
		default:
			log.Fatalf("Got invalid object type %v while looking for RolloutAction", tp)
		}
	}

	return objlist
}

// Watch sets up a event handlers for RolloutAction object
func (api *rolloutactionAPI) Watch(handler RolloutActionHandler) error {
	return api.ct.WatchRolloutAction(handler)
}

// RolloutAction returns RolloutActionAPI
func (ct *ctrlerCtx) RolloutAction() RolloutActionAPI {
	return &rolloutactionAPI{ct: ct}
}
