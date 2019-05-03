package mock

import (
	"sync"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/events/generated/eventtypes"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"
)

// It is recommended to use this mock events recorder only in the unit-tests and in some integ. test which are not
// testing the events. Events recorded using the mock recorder will not make it to venice rather
// it will be logged as a info message using the logger.

// Event is a record of an event generated by the mock recorder
type Event struct {
	EventType string
	Category  string
	Severity  string
	Message   string
	ObjRef    interface{}
}

// Recorder is a mock implementation of events recorder
type Recorder struct {
	name   string     // name (unique id/key) of the recorder
	logger log.Logger // logger
	events []Event
	evLock sync.Mutex
}

// NewRecorder is a mock events recorder that logs events using the given logger.
// It is used in the unit tests
func NewRecorder(name string, l log.Logger) *Recorder {
	return &Recorder{name: name, logger: l}
}

// Event mock implementation that records the event using a logger
func (mr *Recorder) Event(eventType eventtypes.EventType, message string, objRef interface{}) {
	var objRefMeta *api.ObjectMeta
	var objRefKind string
	var err error

	// derive reference object details from the given object
	if objRef != nil {
		objRefMeta, err = runtime.GetObjectMeta(objRef)
		if err != nil {
			mr.logger.Debugf("{%s} failed to get the object meta from reference object, err: %v", mr.name, err)
		}

		objRefKind = objRef.(runtime.Object).GetObjectKind()
	}

	eTypeAttrs := eventtypes.GetEventTypeAttrs(eventType)

	if objRefMeta != nil {
		mr.logger.InfoLog(
			"recorder_name", mr.name,
			"type", eTypeAttrs.EType,
			"category", eTypeAttrs.Category,
			"severity", eTypeAttrs.Severity,
			"message", message,
			"object-ref.tenant", objRefMeta.GetTenant(),
			"object-ref.namespace", objRefMeta.GetNamespace(),
			"object-ref.kind", objRefKind,
			"object-ref.name", objRefMeta.GetName())
	} else {
		mr.logger.InfoLog(
			"recorder_name", mr.name,
			"type", eTypeAttrs.EType,
			"category", eTypeAttrs.Category,
			"severity", eTypeAttrs.Severity,
			"message", message)
	}
	event := Event{
		EventType: eTypeAttrs.EType,
		Category:  eTypeAttrs.Category,
		Severity:  eTypeAttrs.Severity,
		Message:   message,
		ObjRef:    objRef,
	}
	mr.evLock.Lock()
	defer mr.evLock.Unlock()
	mr.events = append(mr.events, event)
}

// StartExport mock impl
func (mr *Recorder) StartExport() {}

// Close mock impl
func (mr *Recorder) Close() {}

// GetEvents return a copy of the event list
func (mr *Recorder) GetEvents() []Event {
	return mr.events
}

// ClearEvents clears the event list
func (mr *Recorder) ClearEvents() {
	mr.events = []Event{}
}
