// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package cluster is a auto generated package.
Input file: svc_cluster.proto
*/
package cluster

import (
	"context"
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
func (m *ClusterList) MakeKey(prefix string) string {
	obj := Cluster{}
	return obj.MakeKey(prefix)
}

func (m *ClusterList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *HostList) MakeKey(prefix string) string {
	obj := Host{}
	return obj.MakeKey(prefix)
}

func (m *HostList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *NodeList) MakeKey(prefix string) string {
	obj := Node{}
	return obj.MakeKey(prefix)
}

func (m *NodeList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *SmartNICList) MakeKey(prefix string) string {
	obj := SmartNIC{}
	return obj.MakeKey(prefix)
}

func (m *SmartNICList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *TenantList) MakeKey(prefix string) string {
	obj := Tenant{}
	return obj.MakeKey(prefix)
}

func (m *TenantList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *VersionList) MakeKey(prefix string) string {
	obj := Version{}
	return obj.MakeKey(prefix)
}

func (m *VersionList) MakeURI(ver, prefix string) string {
	return fmt.Sprint("/", globals.ConfigURIPrefix, "/", prefix, "/", ver)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgClusterWatchHelper) MakeKey(prefix string) string {
	obj := Cluster{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgHostWatchHelper) MakeKey(prefix string) string {
	obj := Host{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgNodeWatchHelper) MakeKey(prefix string) string {
	obj := Node{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgSmartNICWatchHelper) MakeKey(prefix string) string {
	obj := SmartNIC{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgTenantWatchHelper) MakeKey(prefix string) string {
	obj := Tenant{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgVersionWatchHelper) MakeKey(prefix string) string {
	obj := Version{}
	return obj.MakeKey(prefix)
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgClusterWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgClusterWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgClusterWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgClusterWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgClusterWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgClusterWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgClusterWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgClusterWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgClusterWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgClusterWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgClusterWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgClusterWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgHostWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgHostWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgHostWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgHostWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgHostWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgHostWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgHostWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgHostWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgHostWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgHostWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgHostWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgHostWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNodeWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNodeWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgNodeWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgNodeWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNodeWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNodeWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgNodeWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgNodeWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgNodeWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgNodeWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgNodeWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgNodeWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgSmartNICWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgSmartNICWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgSmartNICWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgSmartNICWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgSmartNICWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgSmartNICWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgSmartNICWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgSmartNICWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgSmartNICWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgSmartNICWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgSmartNICWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgSmartNICWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgTenantWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgTenantWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgTenantWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgTenantWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgTenantWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgTenantWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgTenantWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgTenantWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgTenantWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgTenantWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgTenantWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgTenantWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgVersionWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgVersionWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgVersionWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgVersionWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgVersionWatchHelper))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgVersionWatchHelper) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgVersionWatchHelper_WatchEvent) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgVersionWatchHelper_WatchEvent
	var ok bool
	if into == nil {
		out = &AutoMsgVersionWatchHelper_WatchEvent{}
	} else {
		out, ok = into.(*AutoMsgVersionWatchHelper_WatchEvent)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AutoMsgVersionWatchHelper_WatchEvent))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AutoMsgVersionWatchHelper_WatchEvent) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *ClusterList) Clone(into interface{}) (interface{}, error) {
	var out *ClusterList
	var ok bool
	if into == nil {
		out = &ClusterList{}
	} else {
		out, ok = into.(*ClusterList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*ClusterList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *ClusterList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *HostList) Clone(into interface{}) (interface{}, error) {
	var out *HostList
	var ok bool
	if into == nil {
		out = &HostList{}
	} else {
		out, ok = into.(*HostList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*HostList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *HostList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *NodeList) Clone(into interface{}) (interface{}, error) {
	var out *NodeList
	var ok bool
	if into == nil {
		out = &NodeList{}
	} else {
		out, ok = into.(*NodeList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*NodeList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *NodeList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *SmartNICList) Clone(into interface{}) (interface{}, error) {
	var out *SmartNICList
	var ok bool
	if into == nil {
		out = &SmartNICList{}
	} else {
		out, ok = into.(*SmartNICList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SmartNICList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SmartNICList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TenantList) Clone(into interface{}) (interface{}, error) {
	var out *TenantList
	var ok bool
	if into == nil {
		out = &TenantList{}
	} else {
		out, ok = into.(*TenantList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*TenantList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *TenantList) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *VersionList) Clone(into interface{}) (interface{}, error) {
	var out *VersionList
	var ok bool
	if into == nil {
		out = &VersionList{}
	} else {
		out, ok = into.(*VersionList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*VersionList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *VersionList) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *AutoMsgClusterWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgClusterWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgClusterWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgClusterWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgClusterWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgClusterWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgHostWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgHostWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgHostWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgHostWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgHostWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgHostWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgNodeWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgNodeWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgNodeWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgNodeWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgNodeWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgNodeWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgSmartNICWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgSmartNICWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgSmartNICWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgSmartNICWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgSmartNICWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgSmartNICWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgTenantWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgTenantWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgTenantWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgTenantWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgTenantWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgTenantWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *AutoMsgVersionWatchHelper) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgVersionWatchHelper) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgVersionWatchHelper) Normalize() {

	for k, v := range m.Events {
		if v != nil {
			v.Normalize()
			m.Events[k] = v
		}
	}

}

func (m *AutoMsgVersionWatchHelper_WatchEvent) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AutoMsgVersionWatchHelper_WatchEvent) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *AutoMsgVersionWatchHelper_WatchEvent) Normalize() {

	if m.Object != nil {
		m.Object.Normalize()
	}

}

func (m *ClusterList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *ClusterList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *ClusterList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *HostList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *HostList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *HostList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *NodeList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *NodeList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *NodeList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *SmartNICList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SmartNICList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *SmartNICList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *TenantList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *TenantList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *TenantList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

func (m *VersionList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *VersionList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
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

func (m *VersionList) Normalize() {

	for k, v := range m.Items {
		if v != nil {
			v.Normalize()
			m.Items[k] = v
		}
	}

}

// Transformers

func (m *AutoMsgClusterWatchHelper) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	for i, v := range m.Events {
		c := *v
		if err := c.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
		m.Events[i] = &c
	}
	return nil
}

func (m *AutoMsgClusterWatchHelper_WatchEvent) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {

	if m.Object == nil {
		return nil
	}
	if err := m.Object.ApplyStorageTransformer(ctx, toStorage); err != nil {
		return err
	}
	return nil
}

func (m *ClusterList) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	for i, v := range m.Items {
		c := *v
		if err := c.ApplyStorageTransformer(ctx, toStorage); err != nil {
			return err
		}
		m.Items[i] = &c
	}
	return nil
}

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

}
