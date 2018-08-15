// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package networkApiServer is a auto generated package.
Input file: svc_network.proto
*/
package network

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapSvc_network = map[string]*api.Struct{

	"network.AutoMsgLbPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgLbPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.AutoMsgLbPolicyWatchHelper.WatchEvent"},
		},
	},
	"network.AutoMsgLbPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgLbPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.LbPolicy"},
		},
	},
	"network.AutoMsgNetworkWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgNetworkWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.AutoMsgNetworkWatchHelper.WatchEvent"},
		},
	},
	"network.AutoMsgNetworkWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgNetworkWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.Network"},
		},
	},
	"network.AutoMsgServiceWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgServiceWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.AutoMsgServiceWatchHelper.WatchEvent"},
		},
	},
	"network.AutoMsgServiceWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgServiceWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.Service"},
		},
	},
	"network.LbPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(LbPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.LbPolicy"},
		},
	},
	"network.NetworkList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(NetworkList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.Network"},
		},
	},
	"network.ServiceList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(ServiceList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.Service"},
		},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSvc_network)
}
