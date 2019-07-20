// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package workload is a auto generated package.
Input file: svc_workload.proto
*/
package workload

import (
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// MakeKey generates a KV store key for the object
func (m *EndpointList) MakeKey(prefix string) string {
	obj := Endpoint{}
	return obj.MakeKey(prefix)
}

func (m *EndpointList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *WorkloadList) MakeKey(prefix string) string {
	obj := Workload{}
	return obj.MakeKey(prefix)
}

func (m *WorkloadList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgEndpointWatchHelper) MakeKey(prefix string) string {
	obj := Endpoint{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgWorkloadWatchHelper) MakeKey(prefix string) string {
	obj := Workload{}
	return obj.MakeKey(prefix)
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgEndpointWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgEndpointWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgEndpointWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgEndpointWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgEndpointWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgEndpointWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgEndpointWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgEndpointWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgEndpointWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgEndpointWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgEndpointWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgEndpointWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgWorkloadWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgWorkloadWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgWorkloadWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgWorkloadWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgWorkloadWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgWorkloadWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgWorkloadWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgWorkloadWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgWorkloadWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgWorkloadWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgWorkloadWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgWorkloadWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *EndpointList) Clone(into interface{}) (interface{}, error) {
	var out *EndpointList
	var ok bool
	if into == nil {
		out = &EndpointList{}
	} else {
		out, ok = into.(*EndpointList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EndpointList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EndpointList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *WorkloadList) Clone(into interface{}) (interface{}, error) {
	var out *WorkloadList
	var ok bool
	if into == nil {
		out = &WorkloadList{}
	} else {
		out, ok = into.(*WorkloadList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*WorkloadList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *WorkloadList) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *AutoMsgEndpointWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgEndpointWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgEndpointWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgEndpointWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgEndpointWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgEndpointWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgWorkloadWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "events"

		for _, v := range m.Events {
			if v != nil {
				v.References(tenant, tag, resp)
			}
		}
	}
}

func (m *AutoMsgWorkloadWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Events {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sEvents[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AutoMsgWorkloadWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgWorkloadWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "object"

		if m.Object != nil {
			m.Object.References(tenant, tag, resp)
		}

	}
}

func (m *AutoMsgWorkloadWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Object != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Object"
			if errs := m.Object.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AutoMsgWorkloadWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *EndpointList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EndpointList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *EndpointList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *WorkloadList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "items"

		for _, v := range m.Items {
			if v != nil {
				v.References(tenant, tag, resp)
			}
		}
	}
}

func (m *WorkloadList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Items {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sItems[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *WorkloadList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

}
