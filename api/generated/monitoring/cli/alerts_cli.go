// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringCliUtilsBackend is a auto generated package.
Input file: alerts.proto
*/
package cli

import (
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/monitoring"
	"github.com/pensando/sw/venice/cli/gen"
)

// CreateAlertFlags specifies flags for Alert create operation
var CreateAlertFlags = []gen.CliFlag{
	{
		ID:     "state",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeAlertOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.Alert); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.AlertStatus{}
	}
	return nil
}

// CreateAlertDestinationFlags specifies flags for AlertDestination create operation
var CreateAlertDestinationFlags = []gen.CliFlag{
	{
		ID:     "email-list",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "format",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeAlertDestinationOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.AlertDestination); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.AlertDestinationStatus{}
	}
	return nil
}

// CreateAlertPolicyFlags specifies flags for AlertPolicy create operation
var CreateAlertPolicyFlags = []gen.CliFlag{
	{
		ID:     "auto-resolve",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "clear-duration",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "destinations",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "enable",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "message",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "persistence-duration",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "resource",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "severity",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeAlertPolicyOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.AlertPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.AlertPolicyStatus{}
	}
	return nil
}

// CreateStatsAlertPolicyFlags specifies flags for StatsAlertPolicy create operation
var CreateStatsAlertPolicyFlags = []gen.CliFlag{
	{
		ID:     "destinations",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "enable",
		Type:   "Bool",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "field-name",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "function",
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
		ID:     "kind",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "operator",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "raise-value",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "severity",
		Type:   "StringSlice",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
	{
		ID:     "window",
		Type:   "String",
		Help:   "",
		Skip:   false,
		Insert: "",
	},
}

func removeStatsAlertPolicyOper(obj interface{}) error {
	if v, ok := obj.(*monitoring.StatsAlertPolicy); ok {
		v.UUID = ""
		v.ResourceVersion = ""
		v.CreationTime = api.Timestamp{}
		v.ModTime = api.Timestamp{}
		v.Status = monitoring.StatsAlertPolicyStatus{}
	}
	return nil
}

func init() {
	cl := gen.GetInfo()

	cl.AddCliInfo("monitoring.Alert", "create", CreateAlertFlags)
	cl.AddRemoveObjOperFunc("monitoring.Alert", removeAlertOper)

	cl.AddCliInfo("monitoring.AlertDestination", "create", CreateAlertDestinationFlags)
	cl.AddRemoveObjOperFunc("monitoring.AlertDestination", removeAlertDestinationOper)

	cl.AddCliInfo("monitoring.AlertPolicy", "create", CreateAlertPolicyFlags)
	cl.AddRemoveObjOperFunc("monitoring.AlertPolicy", removeAlertPolicyOper)

	cl.AddCliInfo("monitoring.StatsAlertPolicy", "create", CreateStatsAlertPolicyFlags)
	cl.AddRemoveObjOperFunc("monitoring.StatsAlertPolicy", removeStatsAlertPolicyOper)

}
