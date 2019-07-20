// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: eventpolicy.proto
*/
package monitoring

import (
	"errors"
	fmt "fmt"
	"strings"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// MonitoringExportFormat_normal is a map of normalized values for the enum
var MonitoringExportFormat_normal = map[string]string{
	"syslog-bsd":     "syslog-bsd",
	"syslog-rfc5424": "syslog-rfc5424",
}

var MonitoringExportFormat_vname = map[int32]string{
	0: "syslog-bsd",
	1: "syslog-rfc5424",
}

var MonitoringExportFormat_vvalue = map[string]int32{
	"syslog-bsd":     0,
	"syslog-rfc5424": 1,
}

func (x MonitoringExportFormat) String() string {
	return MonitoringExportFormat_vname[int32(x)]
}

var _ validators.DummyVar
var validatorMapEventpolicy = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *EventPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "event-policy/", m.Tenant, "/", m.Name)
}

func (m *EventPolicy) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/event-policy/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *EventPolicy) Clone(into interface{}) (interface{}, error) {
	var out *EventPolicy
	var ok bool
	if into == nil {
		out = &EventPolicy{}
	} else {
		out, ok = into.(*EventPolicy)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EventPolicy))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EventPolicy) Defaults(ver string) bool {
	var ret bool
	m.Kind = "EventPolicy"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *EventPolicySpec) Clone(into interface{}) (interface{}, error) {
	var out *EventPolicySpec
	var ok bool
	if into == nil {
		out = &EventPolicySpec{}
	} else {
		out, ok = into.(*EventPolicySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EventPolicySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EventPolicySpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Format = "syslog-bsd"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *EventPolicyStatus) Clone(into interface{}) (interface{}, error) {
	var out *EventPolicyStatus
	var ok bool
	if into == nil {
		out = &EventPolicyStatus{}
	} else {
		out, ok = into.(*EventPolicyStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EventPolicyStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EventPolicyStatus) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *EventPolicy) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	tenant = m.Tenant

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "meta.tenant"
		uref, ok := resp[tag]
		if !ok {
			uref = apiintf.ReferenceObj{
				RefType: apiintf.ReferenceType("NamedRef"),
			}
		}

		if m.Tenant != "" {
			uref.Refs = append(uref.Refs, globals.ConfigRootPrefix+"/cluster/"+"tenants/"+m.Tenant)
		}

		if len(uref.Refs) > 0 {
			resp[tag] = uref
		}
	}
}

func (m *EventPolicy) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for EventPolicy"))
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "ObjectMeta"
		if errs := m.ObjectMeta.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}

	if !ignoreSpec {

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *EventPolicy) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *EventPolicySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EventPolicySpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Selector != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Selector"
			if errs := m.Selector.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}

	if m.SyslogConfig != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "SyslogConfig"
			if errs := m.SyslogConfig.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	for k, v := range m.Targets {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sTargets[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapEventpolicy["EventPolicySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapEventpolicy["EventPolicySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *EventPolicySpec) Normalize() {

	m.Format = MonitoringExportFormat_normal[strings.ToLower(m.Format)]

	if m.Selector != nil {
		m.Selector.Normalize()
	}

	if m.SyslogConfig != nil {
		m.SyslogConfig.Normalize()
	}

	for k, v := range m.Targets {
		if v != nil {
			v.Normalize()
			m.Targets[k] = v
		}
	}

}

func (m *EventPolicyStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EventPolicyStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *EventPolicyStatus) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&EventPolicy{},
	)

	validatorMapEventpolicy = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapEventpolicy["EventPolicySpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapEventpolicy["EventPolicySpec"]["all"] = append(validatorMapEventpolicy["EventPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*EventPolicySpec)

		if _, ok := MonitoringExportFormat_vvalue[m.Format]; !ok {
			vals := []string{}
			for k1, _ := range MonitoringExportFormat_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Format", vals)
		}
		return nil
	})

}
