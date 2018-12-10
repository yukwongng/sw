// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityCliUtilsBackend is a auto generated package.
Input file: sgpolicy.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/security"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateSGPolicyFlags specifies flags for SGPolicy create operation
var CreateSGPolicyFlags = []gen.CliFlag{
	{
		ID:     "action",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "apps",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "attach-groups",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "attach-tenant",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "from-ip",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "from-security-groups",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "to-ip",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "to-security-groups",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeSGPolicyOper(obj interface{}) error {
	if v, ok := obj.(*security.SGPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = security.SGPolicyStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("security.SGPolicy", "create", CreateSGPolicyFlags)
	cl.AddRemoveObjOperFunc("security.SGPolicy", removeSGPolicyOper)

}
