// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package telemetry_query is a auto generated package.
Input file: telemetry_query.proto
*/
package telemetry_query

import (
	fmt "fmt"
	"strings"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// TsdbFunctionType_normal is a map of normalized values for the enum
var TsdbFunctionType_normal = map[string]string{
	"MAX":  "MAX",
	"MEAN": "MEAN",
	"NONE": "NONE",
	"max":  "MAX",
	"mean": "MEAN",
	"none": "NONE",
}

// SortOrder_normal is a map of normalized values for the enum
var SortOrder_normal = map[string]string{
	"Ascending":  "Ascending",
	"Descending": "Descending",
	"ascending":  "Ascending",
	"descending": "Descending",
}

// FwlogActions_normal is a map of normalized values for the enum
var FwlogActions_normal = map[string]string{
	"allow":         "allow",
	"deny":          "deny",
	"implicit_deny": "implicit_deny",
	"reject":        "reject",
}

// FwlogDirections_normal is a map of normalized values for the enum
var FwlogDirections_normal = map[string]string{
	"from_host":   "from_host",
	"from_uplink": "from_uplink",
}

var _ validators.DummyVar
var validatorMapTelemetry_query = make(map[string]map[string][]func(string, interface{}) error)

// Clone clones the object into into or creates one of into is nil
func (m *Fwlog) Clone(into interface{}) (interface{}, error) {
	var out *Fwlog
	var ok bool
	if into == nil {
		out = &Fwlog{}
	} else {
		out, ok = into.(*Fwlog)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*Fwlog))
	return out, nil
}

// Default sets up the defaults for the object
func (m *Fwlog) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Action = "allow"
		m.Direction = "from_host"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogsQueryList) Clone(into interface{}) (interface{}, error) {
	var out *FwlogsQueryList
	var ok bool
	if into == nil {
		out = &FwlogsQueryList{}
	} else {
		out, ok = into.(*FwlogsQueryList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogsQueryList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogsQueryList) Defaults(ver string) bool {
	var ret bool
	for k := range m.Queries {
		if m.Queries[k] != nil {
			i := m.Queries[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogsQueryResponse) Clone(into interface{}) (interface{}, error) {
	var out *FwlogsQueryResponse
	var ok bool
	if into == nil {
		out = &FwlogsQueryResponse{}
	} else {
		out, ok = into.(*FwlogsQueryResponse)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogsQueryResponse))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogsQueryResponse) Defaults(ver string) bool {
	var ret bool
	for k := range m.Results {
		if m.Results[k] != nil {
			i := m.Results[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogsQueryResult) Clone(into interface{}) (interface{}, error) {
	var out *FwlogsQueryResult
	var ok bool
	if into == nil {
		out = &FwlogsQueryResult{}
	} else {
		out, ok = into.(*FwlogsQueryResult)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogsQueryResult))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogsQueryResult) Defaults(ver string) bool {
	var ret bool
	for k := range m.Logs {
		if m.Logs[k] != nil {
			i := m.Logs[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *FwlogsQuerySpec) Clone(into interface{}) (interface{}, error) {
	var out *FwlogsQuerySpec
	var ok bool
	if into == nil {
		out = &FwlogsQuerySpec{}
	} else {
		out, ok = into.(*FwlogsQuerySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*FwlogsQuerySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *FwlogsQuerySpec) Defaults(ver string) bool {
	var ret bool
	if m.Pagination != nil {
		ret = m.Pagination.Defaults(ver) || ret
	}
	ret = true
	switch ver {
	default:
		for k := range m.Actions {
			m.Actions[k] = "allow"
		}
		for k := range m.Directions {
			m.Directions[k] = "from_host"
		}
		m.SortOrder = "Descending"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *MetricsQueryList) Clone(into interface{}) (interface{}, error) {
	var out *MetricsQueryList
	var ok bool
	if into == nil {
		out = &MetricsQueryList{}
	} else {
		out, ok = into.(*MetricsQueryList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*MetricsQueryList))
	return out, nil
}

// Default sets up the defaults for the object
func (m *MetricsQueryList) Defaults(ver string) bool {
	var ret bool
	for k := range m.Queries {
		if m.Queries[k] != nil {
			i := m.Queries[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *MetricsQueryResponse) Clone(into interface{}) (interface{}, error) {
	var out *MetricsQueryResponse
	var ok bool
	if into == nil {
		out = &MetricsQueryResponse{}
	} else {
		out, ok = into.(*MetricsQueryResponse)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*MetricsQueryResponse))
	return out, nil
}

// Default sets up the defaults for the object
func (m *MetricsQueryResponse) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *MetricsQueryResult) Clone(into interface{}) (interface{}, error) {
	var out *MetricsQueryResult
	var ok bool
	if into == nil {
		out = &MetricsQueryResult{}
	} else {
		out, ok = into.(*MetricsQueryResult)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*MetricsQueryResult))
	return out, nil
}

// Default sets up the defaults for the object
func (m *MetricsQueryResult) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *MetricsQuerySpec) Clone(into interface{}) (interface{}, error) {
	var out *MetricsQuerySpec
	var ok bool
	if into == nil {
		out = &MetricsQuerySpec{}
	} else {
		out, ok = into.(*MetricsQuerySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*MetricsQuerySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *MetricsQuerySpec) Defaults(ver string) bool {
	var ret bool
	if m.Pagination != nil {
		ret = m.Pagination.Defaults(ver) || ret
	}
	ret = true
	switch ver {
	default:
		m.Function = "NONE"
		m.SortOrder = "Ascending"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *PaginationSpec) Clone(into interface{}) (interface{}, error) {
	var out *PaginationSpec
	var ok bool
	if into == nil {
		out = &PaginationSpec{}
	} else {
		out, ok = into.(*PaginationSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*PaginationSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *PaginationSpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Offset = 0
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *ResultSeries) Clone(into interface{}) (interface{}, error) {
	var out *ResultSeries
	var ok bool
	if into == nil {
		out = &ResultSeries{}
	} else {
		out, ok = into.(*ResultSeries)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*ResultSeries))
	return out, nil
}

// Default sets up the defaults for the object
func (m *ResultSeries) Defaults(ver string) bool {
	return false
}

// Validators and Requirements

func (m *Fwlog) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *Fwlog) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapTelemetry_query["Fwlog"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTelemetry_query["Fwlog"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *Fwlog) Normalize() {

	m.Action = FwlogActions_normal[strings.ToLower(m.Action)]

	m.Direction = FwlogDirections_normal[strings.ToLower(m.Direction)]

}

func (m *FwlogsQueryList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FwlogsQueryList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Queries {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sQueries[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *FwlogsQueryList) Normalize() {

	for _, v := range m.Queries {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *FwlogsQueryResponse) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FwlogsQueryResponse) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Results {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sResults[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *FwlogsQueryResponse) Normalize() {

	for _, v := range m.Results {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *FwlogsQueryResult) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FwlogsQueryResult) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Logs {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sLogs[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *FwlogsQueryResult) Normalize() {

	for _, v := range m.Logs {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *FwlogsQuerySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *FwlogsQuerySpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Pagination != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Pagination"
			if errs := m.Pagination.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	if vs, ok := validatorMapTelemetry_query["FwlogsQuerySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTelemetry_query["FwlogsQuerySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *FwlogsQuerySpec) Normalize() {

	for k, v := range m.Actions {
		m.Actions[k] = FwlogActions_normal[strings.ToLower(v)]
	}

	for k, v := range m.Directions {
		m.Directions[k] = FwlogDirections_normal[strings.ToLower(v)]
	}

	if m.Pagination != nil {
		m.Pagination.Normalize()
	}

	m.SortOrder = SortOrder_normal[strings.ToLower(m.SortOrder)]

}

func (m *MetricsQueryList) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *MetricsQueryList) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	for k, v := range m.Queries {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sQueries[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *MetricsQueryList) Normalize() {

	for _, v := range m.Queries {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *MetricsQueryResponse) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *MetricsQueryResponse) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *MetricsQueryResponse) Normalize() {

}

func (m *MetricsQueryResult) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *MetricsQueryResult) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *MetricsQueryResult) Normalize() {

}

func (m *MetricsQuerySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *MetricsQuerySpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error

	if m.Pagination != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Pagination"
			if errs := m.Pagination.Validate(ver, npath, ignoreStatus, ignoreSpec); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}

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
	if vs, ok := validatorMapTelemetry_query["MetricsQuerySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTelemetry_query["MetricsQuerySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *MetricsQuerySpec) Normalize() {

	m.Function = TsdbFunctionType_normal[strings.ToLower(m.Function)]

	if m.Pagination != nil {
		m.Pagination.Normalize()
	}

	if m.Selector != nil {
		m.Selector.Normalize()
	}

	m.SortOrder = SortOrder_normal[strings.ToLower(m.SortOrder)]

}

func (m *PaginationSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *PaginationSpec) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	if vs, ok := validatorMapTelemetry_query["PaginationSpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapTelemetry_query["PaginationSpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *PaginationSpec) Normalize() {

}

func (m *ResultSeries) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *ResultSeries) Validate(ver, path string, ignoreStatus bool, ignoreSpec bool) []error {
	var ret []error
	return ret
}

func (m *ResultSeries) Normalize() {

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes()

	validatorMapTelemetry_query = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapTelemetry_query["Fwlog"] = make(map[string][]func(string, interface{}) error)
	validatorMapTelemetry_query["Fwlog"]["all"] = append(validatorMapTelemetry_query["Fwlog"]["all"], func(path string, i interface{}) error {
		m := i.(*Fwlog)

		if _, ok := FwlogActions_value[m.Action]; !ok {
			vals := []string{}
			for k1, _ := range FwlogActions_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Action", vals)
		}
		return nil
	})

	validatorMapTelemetry_query["Fwlog"]["all"] = append(validatorMapTelemetry_query["Fwlog"]["all"], func(path string, i interface{}) error {
		m := i.(*Fwlog)

		if _, ok := FwlogDirections_value[m.Direction]; !ok {
			vals := []string{}
			for k1, _ := range FwlogDirections_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Direction", vals)
		}
		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)

		for k, v := range m.Actions {
			if _, ok := FwlogActions_value[v]; !ok {
				vals := []string{}
				for k1, _ := range FwlogActions_value {
					vals = append(vals, k1)
				}
				return fmt.Errorf("%v[%v] did not match allowed strings %v", path+"."+"Actions", k, vals)
			}
		}
		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)
		for k, v := range m.DestIPs {
			if err := validators.IPAddr(v); err != nil {
				return fmt.Errorf("%v[%v] failed validation: %s", path+"."+"DestIPs", k, err.Error())
			}
		}

		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "65535")

		for _, v := range m.DestPorts {
			if err := validators.IntRange(v, args); err != nil {
				return fmt.Errorf("%v failed validation: %s", path+"."+"DestPorts", err.Error())
			}
		}
		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)

		for k, v := range m.Directions {
			if _, ok := FwlogDirections_value[v]; !ok {
				vals := []string{}
				for k1, _ := range FwlogDirections_value {
					vals = append(vals, k1)
				}
				return fmt.Errorf("%v[%v] did not match allowed strings %v", path+"."+"Directions", k, vals)
			}
		}
		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)

		if _, ok := SortOrder_value[m.SortOrder]; !ok {
			vals := []string{}
			for k1, _ := range SortOrder_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"SortOrder", vals)
		}
		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)
		for k, v := range m.SourceIPs {
			if err := validators.IPAddr(v); err != nil {
				return fmt.Errorf("%v[%v] failed validation: %s", path+"."+"SourceIPs", k, err.Error())
			}
		}

		return nil
	})

	validatorMapTelemetry_query["FwlogsQuerySpec"]["all"] = append(validatorMapTelemetry_query["FwlogsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*FwlogsQuerySpec)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "65535")

		for _, v := range m.SourcePorts {
			if err := validators.IntRange(v, args); err != nil {
				return fmt.Errorf("%v failed validation: %s", path+"."+"SourcePorts", err.Error())
			}
		}
		return nil
	})

	validatorMapTelemetry_query["MetricsQuerySpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapTelemetry_query["MetricsQuerySpec"]["all"] = append(validatorMapTelemetry_query["MetricsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*MetricsQuerySpec)
		args := make([]string, 0)
		args = append(args, "name")

		for _, v := range m.Fields {
			if err := validators.EmptyOr(validators.RegExp, v, args); err != nil {
				return fmt.Errorf("%v failed validation: %s", path+"."+"Fields", err.Error())
			}
		}
		return nil
	})

	validatorMapTelemetry_query["MetricsQuerySpec"]["all"] = append(validatorMapTelemetry_query["MetricsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*MetricsQuerySpec)

		if _, ok := TsdbFunctionType_value[m.Function]; !ok {
			vals := []string{}
			for k1, _ := range TsdbFunctionType_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Function", vals)
		}
		return nil
	})

	validatorMapTelemetry_query["MetricsQuerySpec"]["all"] = append(validatorMapTelemetry_query["MetricsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*MetricsQuerySpec)
		args := make([]string, 0)
		args = append(args, "name")

		if err := validators.EmptyOr(validators.RegExp, m.GroupbyField, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"GroupbyField", err.Error())
		}
		return nil
	})

	validatorMapTelemetry_query["MetricsQuerySpec"]["all"] = append(validatorMapTelemetry_query["MetricsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*MetricsQuerySpec)
		args := make([]string, 0)
		args = append(args, "0")
		args = append(args, "0")

		if err := validators.EmptyOr(validators.Duration, m.GroupbyTime, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"GroupbyTime", err.Error())
		}
		return nil
	})

	validatorMapTelemetry_query["MetricsQuerySpec"]["all"] = append(validatorMapTelemetry_query["MetricsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*MetricsQuerySpec)
		args := make([]string, 0)
		args = append(args, "name")

		if err := validators.EmptyOr(validators.RegExp, m.Name, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Name", err.Error())
		}
		return nil
	})

	validatorMapTelemetry_query["MetricsQuerySpec"]["all"] = append(validatorMapTelemetry_query["MetricsQuerySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*MetricsQuerySpec)

		if _, ok := SortOrder_value[m.SortOrder]; !ok {
			vals := []string{}
			for k1, _ := range SortOrder_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"SortOrder", vals)
		}
		return nil
	})

	validatorMapTelemetry_query["PaginationSpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapTelemetry_query["PaginationSpec"]["all"] = append(validatorMapTelemetry_query["PaginationSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*PaginationSpec)
		args := make([]string, 0)
		args = append(args, "1")

		if err := validators.IntMin(m.Count, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Count", err.Error())
		}
		return nil
	})

	validatorMapTelemetry_query["PaginationSpec"]["all"] = append(validatorMapTelemetry_query["PaginationSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*PaginationSpec)
		args := make([]string, 0)
		args = append(args, "0")

		if err := validators.IntMin(m.Offset, args); err != nil {
			return fmt.Errorf("%v failed validation: %s", path+"."+"Offset", err.Error())
		}
		return nil
	})

}
