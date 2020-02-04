// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package networkApiServer is a auto generated package.
Input file: networkinterface.proto
*/
package network

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapNetworkinterface = map[string]*api.Struct{

	"network.NetworkInterface": &api.Struct{
		Kind: "NetworkInterface", APIGroup: "network", Scopes: []string{"Cluster"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(NetworkInterface{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.NetworkInterfaceSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.NetworkInterfaceStatus"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{ID: "kind", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{ID: "api-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{ID: "tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{ID: "namespace", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{ID: "generation-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"ResourceVersion": api.Field{Name: "ResourceVersion", CLITag: api.CLIInfo{ID: "resource-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"UUID": api.Field{Name: "UUID", CLITag: api.CLIInfo{ID: "uuid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uuid", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{ID: "labels", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Mutable: true, Map: true, Inline: false, FromInline: true, KeyType: "TYPE_STRING", Type: "TYPE_STRING"},

			"CreationTime": api.Field{Name: "CreationTime", CLITag: api.CLIInfo{ID: "creation-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "creation-time", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"ModTime": api.Field{Name: "ModTime", CLITag: api.CLIInfo{ID: "mod-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mod-time", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"SelfLink": api.Field{Name: "SelfLink", CLITag: api.CLIInfo{ID: "self-link", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "self-link", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},

		CLITags: map[string]api.CLIInfo{
			"admin-status":     api.CLIInfo{Path: "Spec.AdminStatus", Skip: false, Insert: "", Help: ""},
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"attach-network":   api.CLIInfo{Path: "Spec.AttachNetwork", Skip: false, Insert: "", Help: ""},
			"attach-tenant":    api.CLIInfo{Path: "Spec.AttachTenant", Skip: false, Insert: "", Help: ""},
			"cable-type":       api.CLIInfo{Path: "Status.IFUplinkStatus.TransceiverStatus.TranceiverCableType", Skip: false, Insert: "", Help: ""},
			"cluster-node":     api.CLIInfo{Path: "Status.ClusterNode", Skip: false, Insert: "", Help: ""},
			"device-id":        api.CLIInfo{Path: "Status.IFHostStatus.DeviceID", Skip: false, Insert: "", Help: ""},
			"dsc":              api.CLIInfo{Path: "Status.Name", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"host-ifname":      api.CLIInfo{Path: "Status.IFHostStatus.HostIfName", Skip: false, Insert: "", Help: ""},
			"ip-alloc-type":    api.CLIInfo{Path: "Spec.IPAllocType", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"link-speed":       api.CLIInfo{Path: "Status.IFUplinkStatus.LinkSpeed", Skip: false, Insert: "", Help: ""},
			"mac-address":      api.CLIInfo{Path: "Spec.MACAddress", Skip: false, Insert: "", Help: ""},
			"mirror-enabled":   api.CLIInfo{Path: "Status.MirrorEnabled", Skip: false, Insert: "", Help: ""},
			"mtu":              api.CLIInfo{Path: "Spec.MTU", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"oper-status":      api.CLIInfo{Path: "Status.OperStatus", Skip: false, Insert: "", Help: ""},
			"pid":              api.CLIInfo{Path: "Status.IFUplinkStatus.TransceiverStatus.TranceiverPid", Skip: false, Insert: "", Help: ""},
			"primary-mac":      api.CLIInfo{Path: "Status.PrimaryMac", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"rx-pause-enabled": api.CLIInfo{Path: "Spec.Pause.RxPauseEnabled", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"speed":            api.CLIInfo{Path: "Spec.Speed", Skip: false, Insert: "", Help: ""},
			"state":            api.CLIInfo{Path: "Status.IFUplinkStatus.TransceiverStatus.TransceiverState", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"tx-pause-enabled": api.CLIInfo{Path: "Spec.Pause.TxPauseEnabled", Skip: false, Insert: "", Help: ""},
			"type":             api.CLIInfo{Path: "Spec.Pause.Type", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"network.NetworkInterfaceHostStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(NetworkInterfaceHostStatus{}) },
		Fields: map[string]api.Field{
			"HostIfName": api.Field{Name: "HostIfName", CLITag: api.CLIInfo{ID: "host-ifname", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "host-ifname", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"DeviceID": api.Field{Name: "DeviceID", CLITag: api.CLIInfo{ID: "device-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "device-id", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"network.NetworkInterfaceSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(NetworkInterfaceSpec{}) },
		Fields: map[string]api.Field{
			"AdminStatus": api.Field{Name: "AdminStatus", CLITag: api.CLIInfo{ID: "admin-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "admin-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Speed": api.Field{Name: "Speed", CLITag: api.CLIInfo{ID: "speed", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "speed", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MTU": api.Field{Name: "MTU", CLITag: api.CLIInfo{ID: "mtu", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mtu", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"Pause": api.Field{Name: "Pause", CLITag: api.CLIInfo{ID: "pause", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pause", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.PauseSpec"},

			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{ID: "type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AttachTenant": api.Field{Name: "AttachTenant", CLITag: api.CLIInfo{ID: "attach-tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attach-tenant", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AttachNetwork": api.Field{Name: "AttachNetwork", CLITag: api.CLIInfo{ID: "attach-network", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attach-network", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"IPAllocType": api.Field{Name: "IPAllocType", CLITag: api.CLIInfo{ID: "ip-alloc-type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ip-alloc-type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"IPConfig": api.Field{Name: "IPConfig", CLITag: api.CLIInfo{ID: "ip-config", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ip-config", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "cluster.IPConfig"},

			"MACAddress": api.Field{Name: "MACAddress", CLITag: api.CLIInfo{ID: "mac-address", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mac-address", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"network.NetworkInterfaceStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(NetworkInterfaceStatus{}) },
		Fields: map[string]api.Field{
			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "dsc", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "dsc", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"DSC": api.Field{Name: "DSC", CLITag: api.CLIInfo{ID: "dsc", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "dsc", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{ID: "type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"OperStatus": api.Field{Name: "OperStatus", CLITag: api.CLIInfo{ID: "oper-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "oper-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"PrimaryMac": api.Field{Name: "PrimaryMac", CLITag: api.CLIInfo{ID: "primary-mac", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "primary-mac", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"IFHostStatus": api.Field{Name: "IFHostStatus", CLITag: api.CLIInfo{ID: "if-host-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "if-host-status", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.NetworkInterfaceHostStatus"},

			"IFUplinkStatus": api.Field{Name: "IFUplinkStatus", CLITag: api.CLIInfo{ID: "if-uplink-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "if-uplink-status", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.NetworkInterfaceUplinkStatus"},

			"MirrorEnabled": api.Field{Name: "MirrorEnabled", CLITag: api.CLIInfo{ID: "mirror-enabled", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mirror-enabled", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"ClusterNode": api.Field{Name: "ClusterNode", CLITag: api.CLIInfo{ID: "cluster-node", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "cluster-node", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"network.NetworkInterfaceUplinkStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(NetworkInterfaceUplinkStatus{}) },
		Fields: map[string]api.Field{
			"LinkSpeed": api.Field{Name: "LinkSpeed", CLITag: api.CLIInfo{ID: "link-speed", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "link-speed", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TransceiverStatus": api.Field{Name: "TransceiverStatus", CLITag: api.CLIInfo{ID: "transceiver-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "transceiver-status", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "network.TransceiverStatus"},
		},
	},
	"network.PauseSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(PauseSpec{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{ID: "type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TxPauseEnabled": api.Field{Name: "TxPauseEnabled", CLITag: api.CLIInfo{ID: "tx-pause-enabled", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tx-pause-enabled", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"RxPauseEnabled": api.Field{Name: "RxPauseEnabled", CLITag: api.CLIInfo{ID: "rx-pause-enabled", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rx-pause-enabled", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},
		},
	},
	"network.TransceiverStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TransceiverStatus{}) },
		Fields: map[string]api.Field{
			"TransceiverState": api.Field{Name: "TransceiverState", CLITag: api.CLIInfo{ID: "state", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "state", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TranceiverCableType": api.Field{Name: "TranceiverCableType", CLITag: api.CLIInfo{ID: "cable-type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "cable-type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TranceiverPid": api.Field{Name: "TranceiverPid", CLITag: api.CLIInfo{ID: "pid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pid", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

var keyMapNetworkinterface = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapNetworkinterface)
	schema.AddPaths(keyMapNetworkinterface)
}
