// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

package rpcserver

import (
	"context"

	"github.com/pensando/sw/api"
	evtsapi "github.com/pensando/sw/api/generated/events"
	"github.com/pensando/sw/venice/utils/events"
	"github.com/pensando/sw/venice/utils/log"
)

// EvtsProxyRPCHandler handles all event proxy RPC calls
type EvtsProxyRPCHandler struct {
	dispatcher events.Dispatcher
	logger     log.Logger
}

// NewEvtsProxyRPCHandler returns a events proxy RPC handler
func NewEvtsProxyRPCHandler(evtsDispatcher events.Dispatcher, logger log.Logger) (*EvtsProxyRPCHandler, error) {
	return &EvtsProxyRPCHandler{
		dispatcher: evtsDispatcher,
		logger:     logger,
	}, nil
}

// ForwardEvent forwards the given event to the dispatcher.
func (e *EvtsProxyRPCHandler) ForwardEvent(ctx context.Context, event *evtsapi.Event) (*api.Empty, error) {
	err := e.dispatcher.Action(*event)
	if err != nil {
		e.logger.Errorf("failed to forward event {%s} from the proxy, err: %v", event.GetUUID(), err)
	}

	return &api.Empty{}, err
}

// ForwardEvents forwards the given list of events to the dispatcher.
func (e *EvtsProxyRPCHandler) ForwardEvents(ctx context.Context, events *evtsapi.EventList) (*api.Empty, error) {
	for _, event := range events.GetItems() {
		temp := *event
		if err := e.dispatcher.Action(temp); err != nil {
			e.logger.Errorf("failed to forward event {%s} from the proxy, err: %v", temp.GetUUID(), err)

			// dispatcher could have been stopped; throw error so, that the caller retries.
			// any failure here could result in duplicate events from the recorder
			return &api.Empty{}, err
		}
	}

	return &api.Empty{}, nil
}

// Stop stops/closes all the underlying connections with rest of venice components
func (e *EvtsProxyRPCHandler) Stop() {
	e.dispatcher.Shutdown() // closes all the event channels and offset trackers; no more exports
}
