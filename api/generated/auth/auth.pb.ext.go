// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package auth is a auto generated package.
Input file: auth.proto
*/
package auth

import (
	"context"
	"errors"
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"

	"github.com/pensando/sw/venice/globals"
	validators "github.com/pensando/sw/venice/utils/apigen/validators"
	"github.com/pensando/sw/venice/utils/runtime"
	"github.com/pensando/sw/venice/utils/transformers/storage"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

var _ validators.DummyVar
var validatorMapAuth = make(map[string]map[string][]func(string, interface{}) error)

var storageTransformersMapAuth = make(map[string][]func(ctx context.Context, i interface{}, toStorage bool) error)

// MakeKey generates a KV store key for the object
func (m *AuthenticationPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "authn-policy", "/Singleton")
}

func (m *AuthenticationPolicy) MakeURI(cat, ver, prefix string) string {
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/authn-policy")
}

// MakeKey generates a KV store key for the object
func (m *PasswordChangeRequest) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "users/", m.Tenant, "/", m.Name)
}

func (m *PasswordChangeRequest) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/users/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *PasswordResetRequest) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "users/", m.Tenant, "/", m.Name)
}

func (m *PasswordResetRequest) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/users/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *Role) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "roles/", m.Tenant, "/", m.Name)
}

func (m *Role) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/roles/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *RoleBinding) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "role-bindings/", m.Tenant, "/", m.Name)
}

func (m *RoleBinding) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/role-bindings/", in.Name)
}

// MakeKey generates a KV store key for the object
func (m *User) MakeKey(prefix string) string {
	return fmt.Sprint(globals.ConfigRootPrefix, "/", prefix, "/", "users/", m.Tenant, "/", m.Name)
}

func (m *User) MakeURI(cat, ver, prefix string) string {
	in := m
	return fmt.Sprint("/", cat, "/", prefix, "/", ver, "/tenant/", in.Tenant, "/users/", in.Name)
}

// Clone clones the object into into or creates one of into is nil
func (m *AuthenticationPolicy) Clone(into interface{}) (interface{}, error) {
	var out *AuthenticationPolicy
	var ok bool
	if into == nil {
		out = &AuthenticationPolicy{}
	} else {
		out, ok = into.(*AuthenticationPolicy)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AuthenticationPolicy) Defaults(ver string) bool {
	m.Kind = "AuthenticationPolicy"
	m.Tenant, m.Namespace = "", ""
	var ret bool
	ret = m.Spec.Defaults(ver) || ret
	ret = m.Status.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AuthenticationPolicySpec) Clone(into interface{}) (interface{}, error) {
	var out *AuthenticationPolicySpec
	var ok bool
	if into == nil {
		out = &AuthenticationPolicySpec{}
	} else {
		out, ok = into.(*AuthenticationPolicySpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AuthenticationPolicySpec) Defaults(ver string) bool {
	var ret bool
	ret = m.Authenticators.Defaults(ver) || ret
	ret = true
	switch ver {
	default:
		m.TokenExpiry = "144h"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *AuthenticationPolicyStatus) Clone(into interface{}) (interface{}, error) {
	var out *AuthenticationPolicyStatus
	var ok bool
	if into == nil {
		out = &AuthenticationPolicyStatus{}
	} else {
		out, ok = into.(*AuthenticationPolicyStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *AuthenticationPolicyStatus) Defaults(ver string) bool {
	var ret bool
	for k := range m.LdapServers {
		if m.LdapServers[k] != nil {
			i := m.LdapServers[k]
			ret = i.Defaults(ver) || ret
		}
	}
	for k := range m.RadiusServers {
		if m.RadiusServers[k] != nil {
			i := m.RadiusServers[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *Authenticators) Clone(into interface{}) (interface{}, error) {
	var out *Authenticators
	var ok bool
	if into == nil {
		out = &Authenticators{}
	} else {
		out, ok = into.(*Authenticators)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Authenticators) Defaults(ver string) bool {
	var ret bool
	if m.Radius != nil {
		ret = m.Radius.Defaults(ver) || ret
	}
	ret = true
	switch ver {
	default:
		for k := range m.AuthenticatorOrder {
			m.AuthenticatorOrder[k] = "LOCAL"
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *Ldap) Clone(into interface{}) (interface{}, error) {
	var out *Ldap
	var ok bool
	if into == nil {
		out = &Ldap{}
	} else {
		out, ok = into.(*Ldap)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Ldap) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LdapAttributeMapping) Clone(into interface{}) (interface{}, error) {
	var out *LdapAttributeMapping
	var ok bool
	if into == nil {
		out = &LdapAttributeMapping{}
	} else {
		out, ok = into.(*LdapAttributeMapping)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *LdapAttributeMapping) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LdapServer) Clone(into interface{}) (interface{}, error) {
	var out *LdapServer
	var ok bool
	if into == nil {
		out = &LdapServer{}
	} else {
		out, ok = into.(*LdapServer)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *LdapServer) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *LdapServerStatus) Clone(into interface{}) (interface{}, error) {
	var out *LdapServerStatus
	var ok bool
	if into == nil {
		out = &LdapServerStatus{}
	} else {
		out, ok = into.(*LdapServerStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *LdapServerStatus) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Result = "Connect_Success"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *Local) Clone(into interface{}) (interface{}, error) {
	var out *Local
	var ok bool
	if into == nil {
		out = &Local{}
	} else {
		out, ok = into.(*Local)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Local) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PasswordChangeRequest) Clone(into interface{}) (interface{}, error) {
	var out *PasswordChangeRequest
	var ok bool
	if into == nil {
		out = &PasswordChangeRequest{}
	} else {
		out, ok = into.(*PasswordChangeRequest)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *PasswordChangeRequest) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PasswordCredential) Clone(into interface{}) (interface{}, error) {
	var out *PasswordCredential
	var ok bool
	if into == nil {
		out = &PasswordCredential{}
	} else {
		out, ok = into.(*PasswordCredential)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *PasswordCredential) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *PasswordResetRequest) Clone(into interface{}) (interface{}, error) {
	var out *PasswordResetRequest
	var ok bool
	if into == nil {
		out = &PasswordResetRequest{}
	} else {
		out, ok = into.(*PasswordResetRequest)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *PasswordResetRequest) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *Permission) Clone(into interface{}) (interface{}, error) {
	var out *Permission
	var ok bool
	if into == nil {
		out = &Permission{}
	} else {
		out, ok = into.(*Permission)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Permission) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		for k := range m.Actions {
			m.Actions[k] = "AllActions"
		}
		m.ResourceNamespace = "_All_"
		m.ResourceTenant = "default"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *Radius) Clone(into interface{}) (interface{}, error) {
	var out *Radius
	var ok bool
	if into == nil {
		out = &Radius{}
	} else {
		out, ok = into.(*Radius)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Radius) Defaults(ver string) bool {
	var ret bool
	for k := range m.Servers {
		if m.Servers[k] != nil {
			i := m.Servers[k]
			ret = i.Defaults(ver) || ret
		}
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *RadiusServer) Clone(into interface{}) (interface{}, error) {
	var out *RadiusServer
	var ok bool
	if into == nil {
		out = &RadiusServer{}
	} else {
		out, ok = into.(*RadiusServer)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RadiusServer) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.AuthMethod = "PAP"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *RadiusServerStatus) Clone(into interface{}) (interface{}, error) {
	var out *RadiusServerStatus
	var ok bool
	if into == nil {
		out = &RadiusServerStatus{}
	} else {
		out, ok = into.(*RadiusServerStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RadiusServerStatus) Defaults(ver string) bool {
	var ret bool
	if m.Server != nil {
		ret = m.Server.Defaults(ver) || ret
	}
	ret = true
	switch ver {
	default:
		m.Result = "Connect_Success"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *Role) Clone(into interface{}) (interface{}, error) {
	var out *Role
	var ok bool
	if into == nil {
		out = &Role{}
	} else {
		out, ok = into.(*Role)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *Role) Defaults(ver string) bool {
	m.Kind = "Role"
	m.Tenant, m.Namespace = "default", "default"
	var ret bool
	ret = m.Spec.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleBinding) Clone(into interface{}) (interface{}, error) {
	var out *RoleBinding
	var ok bool
	if into == nil {
		out = &RoleBinding{}
	} else {
		out, ok = into.(*RoleBinding)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleBinding) Defaults(ver string) bool {
	m.Kind = "RoleBinding"
	m.Tenant, m.Namespace = "default", "default"
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleBindingSpec) Clone(into interface{}) (interface{}, error) {
	var out *RoleBindingSpec
	var ok bool
	if into == nil {
		out = &RoleBindingSpec{}
	} else {
		out, ok = into.(*RoleBindingSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleBindingSpec) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleBindingStatus) Clone(into interface{}) (interface{}, error) {
	var out *RoleBindingStatus
	var ok bool
	if into == nil {
		out = &RoleBindingStatus{}
	} else {
		out, ok = into.(*RoleBindingStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleBindingStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleSpec) Clone(into interface{}) (interface{}, error) {
	var out *RoleSpec
	var ok bool
	if into == nil {
		out = &RoleSpec{}
	} else {
		out, ok = into.(*RoleSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleSpec) Defaults(ver string) bool {
	var ret bool
	for k := range m.Permissions {
		i := m.Permissions[k]
		ret = i.Defaults(ver) || ret
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *RoleStatus) Clone(into interface{}) (interface{}, error) {
	var out *RoleStatus
	var ok bool
	if into == nil {
		out = &RoleStatus{}
	} else {
		out, ok = into.(*RoleStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *RoleStatus) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *TLSOptions) Clone(into interface{}) (interface{}, error) {
	var out *TLSOptions
	var ok bool
	if into == nil {
		out = &TLSOptions{}
	} else {
		out, ok = into.(*TLSOptions)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *TLSOptions) Defaults(ver string) bool {
	return false
}

// Clone clones the object into into or creates one of into is nil
func (m *User) Clone(into interface{}) (interface{}, error) {
	var out *User
	var ok bool
	if into == nil {
		out = &User{}
	} else {
		out, ok = into.(*User)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *User) Defaults(ver string) bool {
	m.Kind = "User"
	m.Tenant, m.Namespace = "default", "default"
	var ret bool
	ret = m.Spec.Defaults(ver) || ret
	ret = m.Status.Defaults(ver) || ret
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *UserSpec) Clone(into interface{}) (interface{}, error) {
	var out *UserSpec
	var ok bool
	if into == nil {
		out = &UserSpec{}
	} else {
		out, ok = into.(*UserSpec)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *UserSpec) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		m.Type = "Local"
	}
	return ret
}

// Clone clones the object into into or creates one of into is nil
func (m *UserStatus) Clone(into interface{}) (interface{}, error) {
	var out *UserStatus
	var ok bool
	if into == nil {
		out = &UserStatus{}
	} else {
		out, ok = into.(*UserStatus)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Default sets up the defaults for the object
func (m *UserStatus) Defaults(ver string) bool {
	var ret bool
	ret = true
	switch ver {
	default:
		for k := range m.Authenticators {
			m.Authenticators[k] = "LOCAL"
		}
	}
	return ret
}

// Validators

func (m *AuthenticationPolicy) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		ret = m.ObjectMeta.Validate(ver, path+dlmtr+"ObjectMeta", ignoreStatus)
	}
	if m.Tenant != "" {
		ret = append(ret, errors.New("Tenant not allowed for AuthenticationPolicy"))
	}

	dlmtr := "."
	if path == "" {
		dlmtr = ""
	}
	npath := path + dlmtr + "Spec"
	if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
		ret = append(ret, errs...)
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

func (m *AuthenticationPolicySpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error

	dlmtr := "."
	if path == "" {
		dlmtr = ""
	}
	npath := path + dlmtr + "Authenticators"
	if errs := m.Authenticators.Validate(ver, npath, ignoreStatus); errs != nil {
		ret = append(ret, errs...)
	}
	if vs, ok := validatorMapAuth["AuthenticationPolicySpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["AuthenticationPolicySpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *AuthenticationPolicyStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.LdapServers {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sLdapServers[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	for k, v := range m.RadiusServers {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sRadiusServers[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *Authenticators) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Radius != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Radius"
		if errs := m.Radius.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapAuth["Authenticators"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["Authenticators"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *Ldap) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *LdapAttributeMapping) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *LdapServer) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *LdapServerStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapAuth["LdapServerStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["LdapServerStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *Local) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *PasswordChangeRequest) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *PasswordCredential) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *PasswordResetRequest) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *Permission) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapAuth["Permission"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["Permission"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *Radius) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Servers {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sServers[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *RadiusServer) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapAuth["RadiusServer"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["RadiusServer"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *RadiusServerStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if m.Server != nil {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := path + dlmtr + "Server"
		if errs := m.Server.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	if vs, ok := validatorMapAuth["RadiusServerStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["RadiusServerStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *Role) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		ret = m.ObjectMeta.Validate(ver, path+dlmtr+"ObjectMeta", ignoreStatus)
	}

	dlmtr := "."
	if path == "" {
		dlmtr = ""
	}
	npath := path + dlmtr + "Spec"
	if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
		ret = append(ret, errs...)
	}
	return ret
}

func (m *RoleBinding) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		ret = m.ObjectMeta.Validate(ver, path+dlmtr+"ObjectMeta", ignoreStatus)
	}
	return ret
}

func (m *RoleBindingSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *RoleBindingStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *RoleSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	for k, v := range m.Permissions {
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		npath := fmt.Sprintf("%s%sPermissions[%v]", path, dlmtr, k)
		if errs := v.Validate(ver, npath, ignoreStatus); errs != nil {
			ret = append(ret, errs...)
		}
	}
	return ret
}

func (m *RoleStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *TLSOptions) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	return ret
}

func (m *User) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	{
		dlmtr := "."
		if path == "" {
			dlmtr = ""
		}
		ret = m.ObjectMeta.Validate(ver, path+dlmtr+"ObjectMeta", ignoreStatus)
	}

	dlmtr := "."
	if path == "" {
		dlmtr = ""
	}
	npath := path + dlmtr + "Spec"
	if errs := m.Spec.Validate(ver, npath, ignoreStatus); errs != nil {
		ret = append(ret, errs...)
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

func (m *UserSpec) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapAuth["UserSpec"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["UserSpec"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

func (m *UserStatus) Validate(ver, path string, ignoreStatus bool) []error {
	var ret []error
	if vs, ok := validatorMapAuth["UserStatus"][ver]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	} else if vs, ok := validatorMapAuth["UserStatus"]["all"]; ok {
		for _, v := range vs {
			if err := v(path, m); err != nil {
				ret = append(ret, err)
			}
		}
	}
	return ret
}

// Transformers

func (m *AuthenticationPolicy) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	if err := m.Spec.ApplyStorageTransformer(ctx, toStorage); err != nil {
		return err
	}
	return nil
}

type storageAuthenticationPolicyTransformer struct{}

var StorageAuthenticationPolicyTransformer storageAuthenticationPolicyTransformer

func (st *storageAuthenticationPolicyTransformer) TransformFromStorage(ctx context.Context, i interface{}) (interface{}, error) {
	r := i.(AuthenticationPolicy)
	err := r.ApplyStorageTransformer(ctx, false)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (st *storageAuthenticationPolicyTransformer) TransformToStorage(ctx context.Context, i interface{}) (interface{}, error) {
	r := i.(AuthenticationPolicy)
	err := r.ApplyStorageTransformer(ctx, true)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *AuthenticationPolicySpec) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	if vs, ok := storageTransformersMapAuth["AuthenticationPolicySpec"]; ok {
		for _, v := range vs {
			if err := v(ctx, m, toStorage); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *User) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	if err := m.Spec.ApplyStorageTransformer(ctx, toStorage); err != nil {
		return err
	}
	return nil
}

type storageUserTransformer struct{}

var StorageUserTransformer storageUserTransformer

func (st *storageUserTransformer) TransformFromStorage(ctx context.Context, i interface{}) (interface{}, error) {
	r := i.(User)
	err := r.ApplyStorageTransformer(ctx, false)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (st *storageUserTransformer) TransformToStorage(ctx context.Context, i interface{}) (interface{}, error) {
	r := i.(User)
	err := r.ApplyStorageTransformer(ctx, true)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *UserSpec) ApplyStorageTransformer(ctx context.Context, toStorage bool) error {
	if vs, ok := storageTransformersMapAuth["UserSpec"]; ok {
		for _, v := range vs {
			if err := v(ctx, m, toStorage); err != nil {
				return err
			}
		}
	}
	return nil
}

func init() {
	scheme := runtime.GetDefaultScheme()
	scheme.AddKnownTypes(
		&AuthenticationPolicy{},
		&PasswordChangeRequest{},
		&PasswordResetRequest{},
		&Role{},
		&RoleBinding{},
		&User{},
	)

	validatorMapAuth = make(map[string]map[string][]func(string, interface{}) error)

	validatorMapAuth["AuthenticationPolicySpec"] = make(map[string][]func(string, interface{}) error)

	validatorMapAuth["AuthenticationPolicySpec"]["all"] = append(validatorMapAuth["AuthenticationPolicySpec"]["all"], func(path string, i interface{}) error {
		m := i.(*AuthenticationPolicySpec)
		if !validators.Duration(m.TokenExpiry) {
			return fmt.Errorf("%v validation failed", path+"."+"TokenExpiry")
		}
		return nil
	})

	validatorMapAuth["Authenticators"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["Authenticators"]["all"] = append(validatorMapAuth["Authenticators"]["all"], func(path string, i interface{}) error {
		m := i.(*Authenticators)

		for k, v := range m.AuthenticatorOrder {
			if _, ok := Authenticators_AuthenticatorType_value[v]; !ok {
				return fmt.Errorf("%v[%v] did not match allowed strings", path+"."+"AuthenticatorOrder", k)
			}
		}
		return nil
	})

	validatorMapAuth["LdapServerStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["LdapServerStatus"]["all"] = append(validatorMapAuth["LdapServerStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*LdapServerStatus)

		if _, ok := LdapServerStatus_LdapResult_value[m.Result]; !ok {
			return errors.New("LdapServerStatus.Result did not match allowed strings")
		}
		return nil
	})

	validatorMapAuth["Permission"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["Permission"]["all"] = append(validatorMapAuth["Permission"]["all"], func(path string, i interface{}) error {
		m := i.(*Permission)

		for k, v := range m.Actions {
			if _, ok := Permission_ActionType_value[v]; !ok {
				return fmt.Errorf("%v[%v] did not match allowed strings", path+"."+"Actions", k)
			}
		}
		return nil
	})

	validatorMapAuth["RadiusServer"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["RadiusServer"]["all"] = append(validatorMapAuth["RadiusServer"]["all"], func(path string, i interface{}) error {
		m := i.(*RadiusServer)

		if _, ok := Radius_AuthMethod_value[m.AuthMethod]; !ok {
			return errors.New("RadiusServer.AuthMethod did not match allowed strings")
		}
		return nil
	})

	validatorMapAuth["RadiusServerStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["RadiusServerStatus"]["all"] = append(validatorMapAuth["RadiusServerStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*RadiusServerStatus)

		if _, ok := RadiusServerStatus_RadiusResult_value[m.Result]; !ok {
			return errors.New("RadiusServerStatus.Result did not match allowed strings")
		}
		return nil
	})

	validatorMapAuth["UserSpec"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["UserSpec"]["all"] = append(validatorMapAuth["UserSpec"]["all"], func(path string, i interface{}) error {
		m := i.(*UserSpec)

		if _, ok := UserSpec_UserType_value[m.Type]; !ok {
			return errors.New("UserSpec.Type did not match allowed strings")
		}
		return nil
	})

	validatorMapAuth["UserStatus"] = make(map[string][]func(string, interface{}) error)
	validatorMapAuth["UserStatus"]["all"] = append(validatorMapAuth["UserStatus"]["all"], func(path string, i interface{}) error {
		m := i.(*UserStatus)

		for k, v := range m.Authenticators {
			if _, ok := Authenticators_AuthenticatorType_value[v]; !ok {
				return fmt.Errorf("%v[%v] did not match allowed strings", path+"."+"Authenticators", k)
			}
		}
		return nil
	})

	{
		AuthenticationPolicySpecSecretTx, err := storage.NewSecretValueTransformer()
		if err != nil {
			log.Fatalf("Error instantiating SecretStorageTransformer: %v", err)
		}
		storageTransformersMapAuth["AuthenticationPolicySpec"] = append(storageTransformersMapAuth["AuthenticationPolicySpec"],
			func(ctx context.Context, i interface{}, toStorage bool) error {
				var data []byte
				var err error
				m := i.(*AuthenticationPolicySpec)

				if toStorage {
					data, err = AuthenticationPolicySpecSecretTx.TransformToStorage(ctx, []byte(m.Secret))
				} else {
					data, err = AuthenticationPolicySpecSecretTx.TransformFromStorage(ctx, []byte(m.Secret))
				}
				m.Secret = []byte(data)

				return err
			})
	}

	{
		UserSpecPasswordTx, err := storage.NewSecretValueTransformer()
		if err != nil {
			log.Fatalf("Error instantiating SecretStorageTransformer: %v", err)
		}
		storageTransformersMapAuth["UserSpec"] = append(storageTransformersMapAuth["UserSpec"],
			func(ctx context.Context, i interface{}, toStorage bool) error {
				var data []byte
				var err error
				m := i.(*UserSpec)

				if toStorage {
					data, err = UserSpecPasswordTx.TransformToStorage(ctx, []byte(m.Password))
				} else {
					data, err = UserSpecPasswordTx.TransformFromStorage(ctx, []byte(m.Password))
				}
				m.Password = string(data)

				return err
			})
	}

}
