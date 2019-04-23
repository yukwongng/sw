// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: alerts.proto
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

	"github.com/pensando/sw/api/generated/events"

	validators "github.com/pensando/sw/venice/utils/apigen/validators"

	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/runtime"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// AlertSpec_AlertState_normal is a map of normalized values for the enum
var AlertSpec_AlertState_normal = map[string]string{
	"ACKNOWLEDGED": "ACKNOWLEDGED",
	"OPEN":         "OPEN",
	"RESOLVED":     "RESOLVED",
	"acknowledged": "ACKNOWLEDGED",
	"open":         "OPEN",
	"resolved":     "RESOLVED",
}

var _ validators.DummyVar
var validatorMapAlerts = make(map[string]map[string][]func(string, interface{}) error)

// MakeKey generates a KV store key for the object
func (m *Alert) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "alerts/", m.Tenant, "/", m.Name)
}

func (m *Alert) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/alerts/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *AlertDestination) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "alertDestinations/", m.Tenant, "/", m.Name)
}

func (m *AlertDestination) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/alertDestinations/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *AlertPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "alertPolicies/", m.Tenant, "/", m.Name)
}

func (m *AlertPolicy) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/alertPolicies/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *Alert) Clone(into interface{}) (interface{}, error) {
	var out *Alert
	var ok bool
	if into == nil {
		out = &Alert{}
	} else {
		out, ok = into.(*Alert)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*Alert))
	return out, nil
}

// Default sets up the defaults for the object
func (m *Alert) Defaults(ver string) bool {
	var ret bool
	m.Kind = "Alert"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Spec.Defaults(ver) || ret
	ret = m.Status.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertDestination) Clone(into interface{}) (interface{}, error) {
	var out *AlertDestination
	var ok bool
	if into == nil {
		out = &AlertDestination{}
	} else {
		out, ok = into.(*AlertDestination)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertDestination))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertDestination) Defaults(ver string) bool {
	var ret bool
	m.Kind = "AlertDestination"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertDestinationSpec) Clone(into interface{}) (interface{}, error) {
	var out *AlertDestinationSpec
	var ok bool
	if into == nil {
		out = &AlertDestinationSpec{}
	} else {
		out, ok = into.(*AlertDestinationSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertDestinationSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertDestinationSpec) Defaults(ver string) bool {
	var ret bool
	if m.SyslogExport != nil {
		ret = m.SyslogExport.Defaults(ver) || ret
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertDestinationStatus) Clone(into interface{}) (interface{}, error) {
	var out *AlertDestinationStatus
	var ok bool
	if into == nil {
		out = &AlertDestinationStatus{}
	} else {
		out, ok = into.(*AlertDestinationStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertDestinationStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertDestinationStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertPolicy) Clone(into interface{}) (interface{}, error) {
	var out *AlertPolicy
	var ok bool
	if into == nil {
		out = &AlertPolicy{}
	} else {
		out, ok = into.(*AlertPolicy)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertPolicy))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertPolicy) Defaults(ver string) bool {
	var ret bool
	m.Kind = "AlertPolicy"
	ret = m.Tenant != "default" || m.Namespace != "default"
	if ret {
		m.Tenant, m.Namespace = "default", "default"
	}
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertPolicySpec) Clone(into interface{}) (interface{}, error) {
	var out *AlertPolicySpec
	var ok bool
	if into == nil {
		out = &AlertPolicySpec{}
	} else {
		out, ok = into.(*AlertPolicySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertPolicySpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertPolicySpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Enable = true
		m.Severity = "INFO"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertPolicyStatus) Clone(into interface{}) (interface{}, error) {
	var out *AlertPolicyStatus
	var ok bool
	if into == nil {
		out = &AlertPolicyStatus{}
	} else {
		out, ok = into.(*AlertPolicyStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertPolicyStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertPolicyStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertReason) Clone(into interface{}) (interface{}, error) {
	var out *AlertReason
	var ok bool
	if into == nil {
		out = &AlertReason{}
	} else {
		out, ok = into.(*AlertReason)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertReason))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertReason) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertSource) Clone(into interface{}) (interface{}, error) {
	var out *AlertSource
	var ok bool
	if into == nil {
		out = &AlertSource{}
	} else {
		out, ok = into.(*AlertSource)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertSource))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertSource) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertSpec) Clone(into interface{}) (interface{}, error) {
	var out *AlertSpec
	var ok bool
	if into == nil {
		out = &AlertSpec{}
	} else {
		out, ok = into.(*AlertSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertSpec))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertSpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.State = "OPEN"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AlertStatus) Clone(into interface{}) (interface{}, error) {
	var out *AlertStatus
	var ok bool
	if into == nil {
		out = &AlertStatus{}
	} else {
		out, ok = into.(*AlertStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AlertStatus))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AlertStatus) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Severity = "INFO"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AuditInfo) Clone(into interface{}) (interface{}, error) {
	var out *AuditInfo
	var ok bool
	if into == nil {
		out = &AuditInfo{}
	} else {
		out, ok = into.(*AuditInfo)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*AuditInfo))
	return out, nil
}

// Default sets up the defaults for the object
func (m *AuditInfo) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *EmailExport) Clone(into interface{}) (interface{}, error) {
	var out *EmailExport
	var ok bool
	if into == nil {
		out = &EmailExport{}
	} else {
		out, ok = into.(*EmailExport)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*EmailExport))
	return out, nil
}

// Default sets up the defaults for the object
func (m *EmailExport) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *MatchedRequirement) Clone(into interface{}) (interface{}, error) {
	var out *MatchedRequirement
	var ok bool
	if into == nil {
		out = &MatchedRequirement{}
	} else {
		out, ok = into.(*MatchedRequirement)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*MatchedRequirement))
	return out, nil
}

// Default sets up the defaults for the object
func (m *MatchedRequirement) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *SNMPExport) Clone(into interface{}) (interface{}, error) {
	var out *SNMPExport
	var ok bool
	if into == nil {
		out = &SNMPExport{}
	} else {
		out, ok = into.(*SNMPExport)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SNMPExport))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SNMPExport) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *SyslogExport) Clone(into interface{}) (interface{}, error) {
	var out *SyslogExport
	var ok bool
	if into == nil {
		out = &SyslogExport{}
	} else {
		out, ok = into.(*SyslogExport)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *(ref.DeepCopy(m).(*SyslogExport))
	return out, nil
}

// Default sets up the defaults for the object
func (m *SyslogExport) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Format = "SYSLOG_BSD"
	}
	return ret
}

// Validators and Requirements

func (m *Alert) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	tenant = m.Tenant

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "status"

		m.Status.References(tenant, tag, resp)

	}
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

func (m *Alert) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for Alert"))
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "ObjectMeta"
		if errs := m.ObjectMeta.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if !ignoreStatus {

		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Status"
		if errs := m.Status.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *Alert) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

	m.Status.Normalize()

}

func (m *AlertDestination) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

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

func (m *AlertDestination) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for AlertDestination"))
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "ObjectMeta"
		if errs := m.ObjectMeta.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AlertDestination) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *AlertDestinationSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AlertDestinationSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.SNMPExport != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "SNMPExport"
			if errs := m.SNMPExport.Validate(ver, npath, ignoreStatus); errs != nil {
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
			if errs := m.Selector.Validate(ver, npath, ignoreStatus); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	if m.SyslogExport != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "SyslogExport"
			if errs := m.SyslogExport.Validate(ver, npath, ignoreStatus); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *AlertDestinationSpec) Normalize() {

	if m.SNMPExport != nil {
		m.SNMPExport.Normalize()
	}

	if m.Selector != nil {
		m.Selector.Normalize()
	}

	if m.SyslogExport != nil {
		m.SyslogExport.Normalize()
	}

}

func (m *AlertDestinationStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AlertDestinationStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *AlertDestinationStatus) Normalize() {

}

func (m *AlertPolicy) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	tenant = m.Tenant

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "spec"

		m.Spec.References(tenant, tag, resp)

	}
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

func (m *AlertPolicy) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error

	if m.Namespace != "default" {
		ret = append(ret, errors.New("Only Namespace default is allowed for AlertPolicy"))
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "ObjectMeta"
		if errs := m.ObjectMeta.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Spec"
		if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AlertPolicy) Normalize() {

	m.ObjectMeta.Normalize()

	m.Spec.Normalize()

}

func (m *AlertPolicySpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "destinations"
		uref, ok := resp[tag]
		if !ok {
			uref = apiintf.ReferenceObj{
				RefType: apiintf.ReferenceType("NamedRef"),
			}
		}

		for _, v := range m.Destinations {

			uref.Refs = append(uref.Refs, globals.ConfigRootPrefix+"/monitoring/"+"alertDestinations/"+tenant+"/"+v)

		}
		if len(uref.Refs) > 0 {
			resp[tag] = uref
		}
	}
}

func (m *AlertPolicySpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Requirements {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sRequirements[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapAlerts["AlertPolicySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAlerts["AlertPolicySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *AlertPolicySpec) Normalize() {

	for _, v := range m.Requirements {
		if v != nil {
			v.Normalize()
		}
	}

	m.Severity = events.SeverityLevel_normal[strings.ToLower(m.Severity)]

}

func (m *AlertPolicyStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AlertPolicyStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *AlertPolicyStatus) Normalize() {

}

func (m *AlertReason) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "alert-policy-id"
		uref, ok := resp[tag]
		if !ok {
			uref = apiintf.ReferenceObj{
				RefType: apiintf.ReferenceType("NamedRef"),
			}
		}

		if m.PolicyID != "" {
			uref.Refs = append(uref.Refs, globals.ConfigRootPrefix+"/monitoring/"+"alertPolicies/"+tenant+"/"+m.PolicyID)
		}

		if len(uref.Refs) > 0 {
			resp[tag] = uref
		}
	}
}

func (m *AlertReason) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.MatchedRequirements {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sMatchedRequirements[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *AlertReason) Normalize() {

	for _, v := range m.MatchedRequirements {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *AlertSource) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AlertSource) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *AlertSource) Normalize() {

}

func (m *AlertSpec) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AlertSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapAlerts["AlertSpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAlerts["AlertSpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *AlertSpec) Normalize() {

	m.State = AlertSpec_AlertState_normal[strings.ToLower(m.State)]

}

func (m *AlertStatus) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		tag := path + dlmtr + "reason"

		m.Reason.References(tenant, tag, resp)

	}
}

func (m *AlertStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error

	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Reason"
		if errs := m.Reason.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapAlerts["AlertStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAlerts["AlertStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *AlertStatus) Normalize() {

	m.Reason.Normalize()

	m.Severity = events.SeverityLevel_normal[strings.ToLower(m.Severity)]

}

func (m *AuditInfo) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *AuditInfo) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *AuditInfo) Normalize() {

}

func (m *EmailExport) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *EmailExport) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *EmailExport) Normalize() {

}

func (m *MatchedRequirement) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *MatchedRequirement) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Requirement != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Requirement"
			if errs := m.Requirement.Validate(ver, npath, ignoreStatus); errs != nil {
				ret = append(ret, errs...)
			}
		}
	}
	return ret
}

func (m *MatchedRequirement) Normalize() {

	if m.Requirement != nil {
		m.Requirement.Normalize()
	}

}

func (m *SNMPExport) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SNMPExport) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.SNMPTrapServers {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sSNMPTrapServers[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *SNMPExport) Normalize() {

	for _, v := range m.SNMPTrapServers {
		if v != nil {
			v.Normalize()
		}
	}

}

func (m *SyslogExport) References(tenant string, path string, resp map[string]apiintf.ReferenceObj) {

}

func (m *SyslogExport) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Config != nil {
		{
			dlmtr := "."
			if path == "" {
				dlmtr = ""
			}
			npath := path + dlmtr + "Config"
			if errs := m.Config.Validate(ver, npath, ignoreStatus); errs != nil {
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
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapAlerts["SyslogExport"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAlerts["SyslogExport"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *SyslogExport) Normalize() {

	if m.Config != nil {
		m.Config.Normalize()
	}

	m.Format = MonitoringExportFormat_normal[strings.ToLower(m.Format)]

	for _, v := range m.Targets {
		if v != nil {
			v.Normalize()
		}
	}

}

// Transformers

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&Alert{},
		&AlertDestination{},
		&AlertPolicy{},
	)

	validatorMapAlerts = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapAlerts["AlertPolicySpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapAlerts["AlertPolicySpec"]["all"] = append(validatorMapAlerts["AlertPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*AlertPolicySpec)

		if _, ok := events.SeverityLevel_value[m.Severity]; !ok {
			vals := []string{}
			for k1, _ := range events.SeverityLevel_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Severity", vals)
		}
		return nil
	})

	validatorMapAlerts["AlertSpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapAlerts["AlertSpec"]["all"] = append(validatorMapAlerts["AlertSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*AlertSpec)

		if _, ok := AlertSpec_AlertState_value[m.State]; !ok {
			vals := []string{}
			for k1, _ := range AlertSpec_AlertState_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"State", vals)
		}
		return nil
	})

	validatorMapAlerts["AlertStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapAlerts["AlertStatus"]["all"] = append(validatorMapAlerts["AlertStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*AlertStatus)

		if _, ok := events.SeverityLevel_value[m.Severity]; !ok {
			vals := []string{}
			for k1, _ := range events.SeverityLevel_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Severity", vals)
		}
		return nil
	})

	validatorMapAlerts["SyslogExport"] = make(map[string][]func(string, interface{}) error)
	validatorMapAlerts["SyslogExport"]["all"] = append(validatorMapAlerts["SyslogExport"]["all"], func(path string, i interface{}) error {
		m := i.(*SyslogExport)

		if _, ok := MonitoringExportFormat_value[m.Format]; !ok {
			vals := []string{}
			for k1, _ := range MonitoringExportFormat_value {
				vals = append(vals, k1)
			}
			return fmt.Errorf("%v did not match allowed strings %v", path+"."+"Format", vals)
		}
		return nil
	})

}
