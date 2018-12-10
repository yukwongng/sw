// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityCliUtilsBackend is a auto generated package.
Input file: app.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/security"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateAppFlags specifies flags for App create operation
var CreateAppFlags = []gen.CliFlag{
	{
		ID:     "drop-multi-question-packets",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "map-entry-timeout",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "max-call-duration",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "program-id",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "protocol",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeAppOper(obj interface{}) error {
	if v, ok := obj.(*security.App); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = security.AppStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("security.App", "create", CreateAppFlags)
	cl.AddRemoveObjOperFunc("security.App", removeAppOper)

}
