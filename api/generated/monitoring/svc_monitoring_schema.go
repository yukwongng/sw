// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringApiServer is a auto generated package.
Input file: svc_monitoring.proto
*/
package monitoring

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapSvc_monitoring = map[string]*api.Struct{

	"monitoring.AlertDestinationList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AlertDestinationList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AlertDestination"},
		},
	},
	"monitoring.AlertList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AlertList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.Alert"},
		},
	},
	"monitoring.AlertPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AlertPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AlertPolicy"},
		},
	},
	"monitoring.AutoMsgAlertDestinationWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAlertDestinationWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgAlertDestinationWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgAlertDestinationWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAlertDestinationWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AlertDestination"},
		},
	},
	"monitoring.AutoMsgAlertPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAlertPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgAlertPolicyWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgAlertPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAlertPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AlertPolicy"},
		},
	},
	"monitoring.AutoMsgAlertWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAlertWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgAlertWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgAlertWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAlertWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.Alert"},
		},
	},
	"monitoring.AutoMsgEventPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgEventPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgEventPolicyWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgEventPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgEventPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.EventPolicy"},
		},
	},
	"monitoring.AutoMsgFlowExportPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgFlowExportPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgFlowExportPolicyWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgFlowExportPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgFlowExportPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportPolicy"},
		},
	},
	"monitoring.AutoMsgFwlogPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgFwlogPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgFwlogPolicyWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgFwlogPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgFwlogPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogPolicy"},
		},
	},
	"monitoring.AutoMsgMirrorSessionWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgMirrorSessionWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgMirrorSessionWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgMirrorSessionWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgMirrorSessionWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.MirrorSession"},
		},
	},
	"monitoring.AutoMsgStatsPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgStatsPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgStatsPolicyWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgStatsPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgStatsPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.StatsPolicy"},
		},
	},
	"monitoring.AutoMsgTroubleshootingSessionWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgTroubleshootingSessionWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AutoMsgTroubleshootingSessionWatchHelper.WatchEvent"},
		},
	},
	"monitoring.AutoMsgTroubleshootingSessionWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgTroubleshootingSessionWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TroubleshootingSession"},
		},
	},
	"monitoring.EventPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(EventPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.EventPolicy"},
		},
	},
	"monitoring.FlowExportPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportPolicy"},
		},
	},
	"monitoring.FwlogPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogPolicy"},
		},
	},
	"monitoring.MirrorSessionList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(MirrorSessionList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.MirrorSession"},
		},
	},
	"monitoring.StatsPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.StatsPolicy"},
		},
	},
	"monitoring.TroubleshootingSessionList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(TroubleshootingSessionList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TroubleshootingSession"},
		},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSvc_monitoring)
}
