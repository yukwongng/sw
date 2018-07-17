// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: svc_monitoring.proto
*/
package monitoring

import (
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"

	"github.com/pensando/sw/venice/globals"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// MakeKey generates a KV store key for the object
func (m *AlertDestinationList) MakeKey(prefix string) string {
	obj := AlertDestination{}
	return obj.MakeKey(prefix)
}

func (m *AlertDestinationList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AlertList) MakeKey(prefix string) string {
	obj := Alert{}
	return obj.MakeKey(prefix)
}

func (m *AlertList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AlertPolicyList) MakeKey(prefix string) string {
	obj := AlertPolicy{}
	return obj.MakeKey(prefix)
}

func (m *AlertPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *EventPolicyList) MakeKey(prefix string) string {
	obj := EventPolicy{}
	return obj.MakeKey(prefix)
}

func (m *EventPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *FlowExportPolicyList) MakeKey(prefix string) string {
	obj := FlowExportPolicy{}
	return obj.MakeKey(prefix)
}

func (m *FlowExportPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *FwlogPolicyList) MakeKey(prefix string) string {
	obj := FwlogPolicy{}
	return obj.MakeKey(prefix)
}

func (m *FwlogPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *MirrorSessionList) MakeKey(prefix string) string {
	obj := MirrorSession{}
	return obj.MakeKey(prefix)
}

func (m *MirrorSessionList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *StatsPolicyList) MakeKey(prefix string) string {
	obj := StatsPolicy{}
	return obj.MakeKey(prefix)
}

func (m *StatsPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgAlertDestinationWatchHelper) MakeKey(prefix string) string {
	obj := AlertDestination{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgAlertPolicyWatchHelper) MakeKey(prefix string) string {
	obj := AlertPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgAlertWatchHelper) MakeKey(prefix string) string {
	obj := Alert{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgEventPolicyWatchHelper) MakeKey(prefix string) string {
	obj := EventPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgFlowExportPolicyWatchHelper) MakeKey(prefix string) string {
	obj := FlowExportPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgFwlogPolicyWatchHelper) MakeKey(prefix string) string {
	obj := FwlogPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgMirrorSessionWatchHelper) MakeKey(prefix string) string {
	obj := MirrorSession{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgStatsPolicyWatchHelper) MakeKey(prefix string) string {
	obj := StatsPolicy{}
	return obj.MakeKey(prefix)
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertDestinationList) Clone(into interface{}) (interface{}, error) {
	var out *AlertDestinationList
	var ok bool
	if into == nil {
		out = &AlertDestinationList{}
	} else {
		out, ok = into.(*AlertDestinationList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertDestinationList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertList) Clone(into interface{}) (interface{}, error) {
	var out *AlertList
	var ok bool
	if into == nil {
		out = &AlertList{}
	} else {
		out, ok = into.(*AlertList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *AlertPolicyList
	var ok bool
	if into == nil {
		out = &AlertPolicyList{}
	} else {
		out, ok = into.(*AlertPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAlertDestinationWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAlertDestinationWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgAlertDestinationWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgAlertDestinationWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAlertDestinationWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAlertDestinationWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAlertDestinationWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgAlertDestinationWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgAlertDestinationWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAlertDestinationWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAlertPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAlertPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgAlertPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgAlertPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAlertPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAlertPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAlertPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgAlertPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgAlertPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAlertPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAlertWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAlertWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgAlertWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgAlertWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAlertWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAlertWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAlertWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgAlertWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgAlertWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgAlertWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgEventPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgEventPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgEventPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgEventPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgEventPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgEventPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgEventPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgEventPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgEventPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgEventPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgFlowExportPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgFlowExportPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgFlowExportPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgFlowExportPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgFlowExportPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgFlowExportPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgFlowExportPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgFlowExportPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgFlowExportPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgFlowExportPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgFwlogPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgFwlogPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgFwlogPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgFwlogPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgFwlogPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgFwlogPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgFwlogPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgFwlogPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgFwlogPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgFwlogPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgMirrorSessionWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgMirrorSessionWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgMirrorSessionWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgMirrorSessionWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgMirrorSessionWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgMirrorSessionWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgMirrorSessionWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgMirrorSessionWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgMirrorSessionWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgMirrorSessionWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgStatsPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgStatsPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgStatsPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgStatsPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgStatsPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgStatsPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgStatsPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgStatsPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgStatsPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgStatsPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *EventPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *EventPolicyList
	var ok bool
	if into == nil {
		out = &EventPolicyList{}
	} else {
		out, ok = into.(*EventPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *EventPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *FlowExportPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *FlowExportPolicyList
	var ok bool
	if into == nil {
		out = &FlowExportPolicyList{}
	} else {
		out, ok = into.(*FlowExportPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *FlowExportPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *FwlogPolicyList
	var ok bool
	if into == nil {
		out = &FwlogPolicyList{}
	} else {
		out, ok = into.(*FwlogPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *MirrorSessionList) Clone(into interface{}) (interface{}, error) {
	var out *MirrorSessionList
	var ok bool
	if into == nil {
		out = &MirrorSessionList{}
	} else {
		out, ok = into.(*MirrorSessionList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *MirrorSessionList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *StatsPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *StatsPolicyList
	var ok bool
	if into == nil {
		out = &StatsPolicyList{}
	} else {
		out, ok = into.(*StatsPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *StatsPolicyList) Defaults(ver string) bool {
	return false
}

// Validators

func (m *AlertDestinationList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AlertList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AlertPolicyList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAlertDestinationWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAlertDestinationWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAlertPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAlertPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAlertWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgAlertWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgEventPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgEventPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgFlowExportPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgFlowExportPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgFwlogPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgFwlogPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgMirrorSessionWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgMirrorSessionWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Object != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Object"
		if errs := m.Object.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgStatsPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *AutoMsgStatsPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *EventPolicyList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *FlowExportPolicyList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *FwlogPolicyList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *MirrorSessionList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%d]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *StatsPolicyList) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

}
