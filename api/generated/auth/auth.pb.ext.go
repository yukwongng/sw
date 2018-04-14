// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package auth is a auto generated package.
Input file: protos/auth.proto
*/
package auth

import (
	fmt "fmt"

	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"

	"github.com/pensando/sw/venice/globals"
	validators "github.com/pensando/sw/venice/utils/apigen/validators"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

var _ validators.DummyVar
var funcMapAuth = make(map[string]map[string][]func(interface{}) bool)

// MakeKey generates a KV store key for the object
func (m *AuthenticationPolicy) MakeKey(prefix string) string {
	return fmt.Sprint(globals.RootPrefix, "/", prefix, "/", "authn-policy/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Role) MakeKey(prefix string) string {
	return fmt.Sprint(globals.RootPrefix, "/", prefix, "/", "roles/", m.Tenant, "/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *RoleBinding) MakeKey(prefix string) string {
	return fmt.Sprint(globals.RootPrefix, "/", prefix, "/", "role-bindings/", m.Tenant, "/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *User) MakeKey(prefix string) string {
	return fmt.Sprint(globals.RootPrefix, "/", prefix, "/", "users/", m.Tenant, "/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *AuthenticationPolicyList) MakeKey(prefix string) string {
	obj := AuthenticationPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *RoleBindingList) MakeKey(prefix string) string {
	obj := RoleBinding{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *RoleList) MakeKey(prefix string) string {
	obj := Role{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *UserList) MakeKey(prefix string) string {
	obj := User{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgAuthenticationPolicyWatchHelper) MakeKey(prefix string) string {
	obj := AuthenticationPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgRoleBindingWatchHelper) MakeKey(prefix string) string {
	obj := RoleBinding{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgRoleWatchHelper) MakeKey(prefix string) string {
	obj := Role{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgUserWatchHelper) MakeKey(prefix string) string {
	obj := User{}
	return obj.MakeKey(prefix)
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

// Clone clones the object into into or creates one of into is nil
func (m *AuthenticationPolicyList) Clone(into interface{}) (interface{}, error) {
	var out *AuthenticationPolicyList
	var ok bool
	if into == nil {
		out = &AuthenticationPolicyList{}
	} else {
		out, ok = into.(*AuthenticationPolicyList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
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

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgAuthenticationPolicyWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgAuthenticationPolicyWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgAuthenticationPolicyWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgAuthenticationPolicyWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgRoleBindingWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgRoleBindingWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgRoleBindingWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgRoleBindingWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgRoleWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgRoleWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgRoleWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgRoleWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
}

// Clone clones the object into into or creates one of into is nil
func (m *AutoMsgUserWatchHelper) Clone(into interface{}) (interface{}, error) {
	var out *AutoMsgUserWatchHelper
	var ok bool
	if into == nil {
		out = &AutoMsgUserWatchHelper{}
	} else {
		out, ok = into.(*AutoMsgUserWatchHelper)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
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

// Clone clones the object into into or creates one of into is nil
func (m *RoleBindingList) Clone(into interface{}) (interface{}, error) {
	var out *RoleBindingList
	var ok bool
	if into == nil {
		out = &RoleBindingList{}
	} else {
		out, ok = into.(*RoleBindingList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
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

// Clone clones the object into into or creates one of into is nil
func (m *RoleList) Clone(into interface{}) (interface{}, error) {
	var out *RoleList
	var ok bool
	if into == nil {
		out = &RoleList{}
	} else {
		out, ok = into.(*RoleList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
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

// Clone clones the object into into or creates one of into is nil
func (m *UserList) Clone(into interface{}) (interface{}, error) {
	var out *UserList
	var ok bool
	if into == nil {
		out = &UserList{}
	} else {
		out, ok = into.(*UserList)
		if !ok {
			return nil, fmt.Errorf("mismatched object types")
		}
	}
	*out = *m
	return out, nil
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

// Validators

func (m *AuthenticationPolicy) Validate(ver string, ignoreStatus bool) bool {
	if !m.Spec.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *AuthenticationPolicyList) Validate(ver string, ignoreStatus bool) bool {
	for _, v := range m.Items {
		if !v.Validate(ver, ignoreStatus) {
			return false
		}
	}
	return true
}

func (m *AuthenticationPolicySpec) Validate(ver string, ignoreStatus bool) bool {
	if !m.Authenticators.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *AuthenticationPolicyStatus) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *Authenticators) Validate(ver string, ignoreStatus bool) bool {
	if vs, ok := funcMapAuth["Authenticators"][ver]; ok {
		for _, v := range vs {
			if !v(m) {
				return false
			}
		}
	} else if vs, ok := funcMapAuth["Authenticators"]["all"]; ok {
		for _, v := range vs {
			if !v(m) {
				return false
			}
		}
	}
	return true
}

func (m *AutoMsgAuthenticationPolicyWatchHelper) Validate(ver string, ignoreStatus bool) bool {
	if m.Object != nil && !m.Object.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *AutoMsgRoleBindingWatchHelper) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *AutoMsgRoleWatchHelper) Validate(ver string, ignoreStatus bool) bool {
	if m.Object != nil && !m.Object.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *AutoMsgUserWatchHelper) Validate(ver string, ignoreStatus bool) bool {
	if m.Object != nil && !m.Object.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *Ldap) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *LdapAttributeMapping) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *Local) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *Permission) Validate(ver string, ignoreStatus bool) bool {
	if vs, ok := funcMapAuth["Permission"][ver]; ok {
		for _, v := range vs {
			if !v(m) {
				return false
			}
		}
	} else if vs, ok := funcMapAuth["Permission"]["all"]; ok {
		for _, v := range vs {
			if !v(m) {
				return false
			}
		}
	}
	return true
}

func (m *Radius) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *Role) Validate(ver string, ignoreStatus bool) bool {
	if !m.Spec.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *RoleBinding) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *RoleBindingList) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *RoleBindingSpec) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *RoleBindingStatus) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *RoleList) Validate(ver string, ignoreStatus bool) bool {
	for _, v := range m.Items {
		if !v.Validate(ver, ignoreStatus) {
			return false
		}
	}
	return true
}

func (m *RoleSpec) Validate(ver string, ignoreStatus bool) bool {
	for _, v := range m.Permissions {
		if !v.Validate(ver, ignoreStatus) {
			return false
		}
	}
	return true
}

func (m *RoleStatus) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *TLSOptions) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func (m *User) Validate(ver string, ignoreStatus bool) bool {
	if !m.Spec.Validate(ver, ignoreStatus) {
		return false
	}
	return true
}

func (m *UserList) Validate(ver string, ignoreStatus bool) bool {
	for _, v := range m.Items {
		if !v.Validate(ver, ignoreStatus) {
			return false
		}
	}
	return true
}

func (m *UserSpec) Validate(ver string, ignoreStatus bool) bool {
	if vs, ok := funcMapAuth["UserSpec"][ver]; ok {
		for _, v := range vs {
			if !v(m) {
				return false
			}
		}
	} else if vs, ok := funcMapAuth["UserSpec"]["all"]; ok {
		for _, v := range vs {
			if !v(m) {
				return false
			}
		}
	}
	return true
}

func (m *UserStatus) Validate(ver string, ignoreStatus bool) bool {
	return true
}

func init() {
	funcMapAuth = make(map[string]map[string][]func(interface{}) bool)

	funcMapAuth["Authenticators"] = make(map[string][]func(interface{}) bool)
	funcMapAuth["Authenticators"]["all"] = append(funcMapAuth["Authenticators"]["all"], func(i interface{}) bool {
		m := i.(*Authenticators)

		for _, v := range m.AuthenticatorOrder {
			if _, ok := Authenticators_AuthenticatorType_value[v]; !ok {
				return false
			}
		}
		return true
	})

	funcMapAuth["Permission"] = make(map[string][]func(interface{}) bool)
	funcMapAuth["Permission"]["all"] = append(funcMapAuth["Permission"]["all"], func(i interface{}) bool {
		m := i.(*Permission)

		for _, v := range m.Actions {
			if _, ok := Permission_ActionType_value[v]; !ok {
				return false
			}
		}
		return true
	})

	funcMapAuth["Permission"]["all"] = append(funcMapAuth["Permission"]["all"], func(i interface{}) bool {
		m := i.(*Permission)

		if _, ok := Permission_ResrcKind_value[m.ResourceKind]; !ok {
			return false
		}
		return true
	})

	funcMapAuth["UserSpec"] = make(map[string][]func(interface{}) bool)
	funcMapAuth["UserSpec"]["all"] = append(funcMapAuth["UserSpec"]["all"], func(i interface{}) bool {
		m := i.(*UserSpec)

		if _, ok := UserSpec_UserType_value[m.Type]; !ok {
			return false
		}
		return true
	})

}
