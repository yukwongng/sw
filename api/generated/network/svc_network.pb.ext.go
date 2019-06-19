// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package network is a auto generated package.
Input file: svc_network.proto
*/
package network

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
func (m *LbPolicyList) MakeKey(prefix string) string {
	obj := LbPolicy{}
	return obj.MakeKey(prefix)
}

func (m *LbPolicyList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *NetworkInterfaceList) MakeKey(prefix string) string {
	obj := NetworkInterface{}
	return obj.MakeKey(prefix)
}

func (m *NetworkInterfaceList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *NetworkList) MakeKey(prefix string) string {
	obj := Network{}
	return obj.MakeKey(prefix)
}

func (m *NetworkList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *ServiceList) MakeKey(prefix string) string {
	obj := Service{}
	return obj.MakeKey(prefix)
}

func (m *ServiceList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *VirtualRouterList) MakeKey(prefix string) string {
	obj := VirtualRouter{}
	return obj.MakeKey(prefix)
}

func (m *VirtualRouterList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgLbPolicyWatchHelper) MakeKey(prefix string) string {
	obj := LbPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgNetworkInterfaceWatchHelper) MakeKey(prefix string) string {
	obj := NetworkInterface{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgNetworkWatchHelper) MakeKey(prefix string) string {
	obj := Network{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgServiceWatchHelper) MakeKey(prefix string) string {
	obj := Service{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgVirtualRouterWatchHelper) MakeKey(prefix string) string {
	obj := VirtualRouter{}
	return obj.MakeKey(prefix)
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgLbPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgLbPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgLbPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgLbPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgLbPolicyWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgLbPolicyWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgLbPolicyWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgLbPolicyWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgLbPolicyWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgLbPolicyWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgLbPolicyWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgLbPolicyWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNetworkInterfaceWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNetworkInterfaceWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgNetworkInterfaceWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgNetworkInterfaceWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNetworkInterfaceWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNetworkInterfaceWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNetworkInterfaceWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNetworkInterfaceWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgNetworkInterfaceWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgNetworkInterfaceWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNetworkInterfaceWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNetworkInterfaceWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNetworkWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNetworkWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgNetworkWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgNetworkWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNetworkWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNetworkWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNetworkWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNetworkWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgNetworkWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgNetworkWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNetworkWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNetworkWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgServiceWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgServiceWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgServiceWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgServiceWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgServiceWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgServiceWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgServiceWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgServiceWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgServiceWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgServiceWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgServiceWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgServiceWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgVirtualRouterWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgVirtualRouterWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgVirtualRouterWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgVirtualRouterWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgVirtualRouterWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgVirtualRouterWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgVirtualRouterWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgVirtualRouterWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgVirtualRouterWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgVirtualRouterWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgVirtualRouterWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgVirtualRouterWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LbPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *LbPolicyList
	var ok bool
	if into == nil {
		out = &LbPolicyList{}
	} else {
		out, ok = into.(*LbPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*LbPolicyList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *LbPolicyList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *NetworkInterfaceList) Clone(into interface{}) (interface{}, error) {
	var out *NetworkInterfaceList
	var ok bool
	if into == nil {
		out = &NetworkInterfaceList{}
	} else {
		out, ok = into.(*NetworkInterfaceList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*NetworkInterfaceList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *NetworkInterfaceList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *NetworkList) Clone(into interface{}) (interface{}, error) {
	var out *NetworkList
	var ok bool
	if into == nil {
		out = &NetworkList{}
	} else {
		out, ok = into.(*NetworkList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*NetworkList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *NetworkList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *ServiceList) Clone(into interface{}) (interface{}, error) {
	var out *ServiceList
	var ok bool
	if into == nil {
		out = &ServiceList{}
	} else {
		out, ok = into.(*ServiceList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*ServiceList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *ServiceList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *VirtualRouterList) Clone(into interface{}) (interface{}, error) {
	var out *VirtualRouterList
	var ok bool
	if into == nil {
		out = &VirtualRouterList{}
	} else {
		out, ok = into.(*VirtualRouterList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*VirtualRouterList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *VirtualRouterList) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *AutoMsgLbPolicyWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgLbPolicyWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgLbPolicyWatchHelper) Normalize() {

	for _, v := range m.Events {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *AutoMsgLbPolicyWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgLbPolicyWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgLbPolicyWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgNetworkInterfaceWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgNetworkInterfaceWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgNetworkInterfaceWatchHelper) Normalize() {

	for _, v := range m.Events {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *AutoMsgNetworkInterfaceWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgNetworkInterfaceWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgNetworkInterfaceWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgNetworkWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgNetworkWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgNetworkWatchHelper) Normalize() {

	for _, v := range m.Events {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *AutoMsgNetworkWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgNetworkWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgNetworkWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgServiceWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgServiceWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgServiceWatchHelper) Normalize() {

	for _, v := range m.Events {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *AutoMsgServiceWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgServiceWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgServiceWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgVirtualRouterWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgVirtualRouterWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgVirtualRouterWatchHelper) Normalize() {

	for _, v := range m.Events {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *AutoMsgVirtualRouterWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgVirtualRouterWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgVirtualRouterWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *LbPolicyList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *LbPolicyList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *LbPolicyList) Normalize() {

	for _, v := range m.Items {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *NetworkInterfaceList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *NetworkInterfaceList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *NetworkInterfaceList) Normalize() {

	for _, v := range m.Items {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *NetworkList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *NetworkList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *NetworkList) Normalize() {

	for _, v := range m.Items {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *ServiceList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *ServiceList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *ServiceList) Normalize() {

	for _, v := range m.Items {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *VirtualRouterList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *VirtualRouterList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *VirtualRouterList) Normalize() {

	for _, v := range m.Items {
		if v != nil {
			v.Normalize()
		}
	}

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

}
