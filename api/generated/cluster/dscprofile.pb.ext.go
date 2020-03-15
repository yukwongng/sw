// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package cluster is a auto generated package.
Input file: dscprofile.proto
*/
package cluster

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

// DSCProfileSpec_Fwd_Mode_normal is a map of normalized values for the enum
var DSCProfileSpec_Fwd_Mode_normal = map[string]string{
	"insertion":   "insertion",
	"transparent": "transparent",
}

var DSCProfileSpec_Fwd_Mode_vname = map[int32]string{
	0: "transparent",
	1: "insertion",
}

var DSCProfileSpec_Fwd_Mode_vvalue = map[string]int32{
	"transparent": 0,
	"insertion":   1,
}

func (x DSCProfileSpec_Fwd_Mode) String() string {
	return DSCProfileSpec_Fwd_Mode_vname[int32(x)]
}

// DSCProfileSpec_FlowPolicy_Mode_normal is a map of normalized values for the enum
var DSCProfileSpec_FlowPolicy_Mode_normal = map[string]string{
	"basenet":   "basenet",
	"enforced":  "enforced",
	"flowaware": "flowaware",
}

var DSCProfileSpec_FlowPolicy_Mode_vname = map[int32]string{
	0: "basenet",
	1: "flowaware",
	2: "enforced",
}

var DSCProfileSpec_FlowPolicy_Mode_vvalue = map[string]int32{
	"basenet":   0,
	"flowaware": 1,
	"enforced":  2,
}

func (x DSCProfileSpec_FlowPolicy_Mode) String() string {
	return DSCProfileSpec_FlowPolicy_Mode_vname[int32(x)]
}

var _ validators.DummyVar
var validatorMapDscprofile = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *DSCProfile) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "dscprofiles/", m.Name)
}

func (m *DSCProfile) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/dscprofiles/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *DSCProfile) Clone(into interface{}) (interface{}, error) {
	var out *DSCProfile
	var ok bool
	if into == nil {
		out = &DSCProfile{}
	} else {
		out, ok = into.(*DSCProfile)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*DSCProfile))
	return out, nil
}

// Default sets up the defaults for the object
func (m *DSCProfile) Defaults(ver string) bool {
	var ret bool
	m.Kind = "DSCProfile"
	ret = m.Tenant != "" || m.Namespace != ""
	if ret {
		m.Tenant, m.Namespace = "", ""
	}
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *DSCProfileSpec) Clone(into interface{}) (interface{}, error) {
	var out *DSCProfileSpec
	var ok bool
	if into == nil {
		out = &DSCProfileSpec{}
	} else {
		out, ok = into.(*DSCProfileSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*DSCProfileSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *DSCProfileSpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.FlowPolicyMode = "basenet"
		m.FwdMode = "transparent"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *DSCProfileStatus) Clone(into interface{}) (interface{}, error) {
	var out *DSCProfileStatus
	var ok bool
	if into == nil {
		out = &DSCProfileStatus{}
	} else {
		out, ok = into.(*DSCProfileStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*DSCProfileStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *DSCProfileStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PropagationStatus) Clone(into interface{}) (interface{}, error) {
	var out *PropagationStatus
	var ok bool
	if into == nil {
		out = &PropagationStatus{}
	} else {
		out, ok = into.(*PropagationStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PropagationStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PropagationStatus) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *DSCProfile) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *DSCProfile) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Tenant != "" {
		ret = append(ret, errors.New("Tenant not allowed for DSCProfile"))
	}
	if m.Namespace != "" {
		ret = append(ret, errors.New("Namespace not allowed for DSCProfile"))
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

func (m *DSCProfile) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *DSCProfileSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *DSCProfileSpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapDscprofile["DSCProfileSpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapDscprofile["DSCProfileSpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *DSCProfileSpec) Normalize() {

	m.FlowPolicyMode = DSCProfileSpec_FlowPolicy_Mode_normal[strings.ToLower(m.FlowPolicyMode)]

	m.FwdMode = DSCProfileSpec_Fwd_Mode_normal[strings.ToLower(m.FwdMode)]

}

func (m *DSCProfileStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *DSCProfileStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *DSCProfileStatus) Normalize() {

}

func (m *PropagationStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *PropagationStatus) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *PropagationStatus) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&DSCProfile{},
	)

	validatorMapDscprofile = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapDscprofile["DSCProfileSpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapDscprofile["DSCProfileSpec"]["all"] = append(validatorMapDscprofile["DSCProfileSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*DSCProfileSpec)

		if _, ok := DSCProfileSpec_FlowPolicy_Mode_vvalue[m.FlowPolicyMode]; !ok {
			vals := []string{}
			for k1, _ := range DSCProfileSpec_FlowPolicy_Mode_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"FlowPolicyMode", vals)
		}
		return nil
	})

	validatorMapDscprofile["DSCProfileSpec"]["all"] = append(validatorMapDscprofile["DSCProfileSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*DSCProfileSpec)

		if _, ok := DSCProfileSpec_Fwd_Mode_vvalue[m.FwdMode]; !ok {
			vals := []string{}
			for k1, _ := range DSCProfileSpec_Fwd_Mode_vvalue {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"FwdMode", vals)
		}
		return nil
	})

}
