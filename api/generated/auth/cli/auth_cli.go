// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package authCliUtilsBackend is a auto generated package.
Input file: auth.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateAuthenticationPolicyFlags specifies flags for AuthenticationPolicy create operation
var CreateAuthenticationPolicyFlags = []gen.CliFlag{
	{
		ID:     "authenticator-order",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "base-dn",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "bind-dn",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "bind-password",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "email",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "enabled",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "fullname",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "group",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "group-object-class",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "nas-id",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "token-expiry",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "user",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "user-object-class",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeAuthenticationPolicyOper(obj interface{}) error {
	if v, ok := obj.(*auth.AuthenticationPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = auth.AuthenticationPolicyStatus{}
	}
	return nil
}

// CreateRoleFlags specifies flags for Role create operation
var CreateRoleFlags = []gen.CliFlag{
	{
		ID:     "actions",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "resource-group",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "resource-kind",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "resource-names",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "resource-namespace",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "resource-tenant",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeRoleOper(obj interface{}) error {
	if v, ok := obj.(*auth.Role); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = auth.RoleStatus{}
	}
	return nil
}

// CreateRoleBindingFlags specifies flags for RoleBinding create operation
var CreateRoleBindingFlags = []gen.CliFlag{
	{
		ID:     "role",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "user-groups",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "users",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeRoleBindingOper(obj interface{}) error {
	if v, ok := obj.(*auth.RoleBinding); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = auth.RoleBindingStatus{}
	}
	return nil
}

// CreateUserFlags specifies flags for User create operation
var CreateUserFlags = []gen.CliFlag{
	{
		ID:     "email",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "fullname",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "password",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "type",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeUserOper(obj interface{}) error {
	if v, ok := obj.(*auth.User); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = auth.UserStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("auth.AuthenticationPolicy", "create", CreateAuthenticationPolicyFlags)
	cl.AddRemoveObjOperFunc("auth.AuthenticationPolicy", removeAuthenticationPolicyOper)

	cl.AddCliInfo("auth.Role", "create", CreateRoleFlags)
	cl.AddRemoveObjOperFunc("auth.Role", removeRoleOper)

	cl.AddCliInfo("auth.RoleBinding", "create", CreateRoleBindingFlags)
	cl.AddRemoveObjOperFunc("auth.RoleBinding", removeRoleBindingOper)

	cl.AddCliInfo("auth.User", "create", CreateUserFlags)
	cl.AddRemoveObjOperFunc("auth.User", removeUserOper)

}
