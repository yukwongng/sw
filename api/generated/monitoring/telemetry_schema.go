// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringApiServer is a auto generated package.
Input file: telemetry.proto
*/
package monitoring

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapTelemetry = map[string]*api.Struct{

	"monitoring.FlowExportPolicy": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version": api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"format":      api.CLIInfo{Path: "Spec.Targets[].Format", Skip: false, Insert: "", Help: ""},
			"interval":    api.CLIInfo{Path: "Spec.Targets[].Interval", Skip: false, Insert: "", Help: ""},
			"kind":        api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.FlowExportSpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportSpec{}) },
		Fields: map[string]api.Field{
			"Targets": api.Field{Name: "Targets", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "targets", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportTarget"},
		},
	},
	"monitoring.FlowExportStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportStatus{}) },
		Fields:    map[string]api.Field{},
	},
	"monitoring.FlowExportTarget": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportTarget{}) },
		Fields: map[string]api.Field{
			"Interval": api.Field{Name: "Interval", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "interval", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Format": api.Field{Name: "Format", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "format", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Exports": api.Field{Name: "Exports", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "exports", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ExportConfig"},
		},
	},
	"monitoring.FwlogExport": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogExport{}) },
		Fields: map[string]api.Field{
			"Targets": api.Field{Name: "Targets", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "targets", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ExportConfig"},

			"Format": api.Field{Name: "Format", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "format", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Filter": api.Field{Name: "Filter", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "export-filter", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"SyslogConfig": api.Field{Name: "SyslogConfig", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "syslog-config", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.SyslogExportConfig"},
		},
	},
	"monitoring.FwlogPolicy": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":    api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"export-filter":  api.CLIInfo{Path: "Spec.Exports[].Filter", Skip: false, Insert: "", Help: ""},
			"filter":         api.CLIInfo{Path: "Spec.Filter", Skip: false, Insert: "", Help: ""},
			"format":         api.CLIInfo{Path: "Spec.Exports[].Format", Skip: false, Insert: "", Help: ""},
			"kind":           api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"retention-time": api.CLIInfo{Path: "Spec.RetentionTime", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.FwlogSpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogSpec{}) },
		Fields: map[string]api.Field{
			"RetentionTime": api.Field{Name: "RetentionTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "retention-time", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Filter": api.Field{Name: "Filter", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "filter", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Exports": api.Field{Name: "Exports", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "exports", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogExport"},
		},
	},
	"monitoring.FwlogStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogStatus{}) },
		Fields:    map[string]api.Field{},
	},
	"monitoring.StatsPolicy": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.StatsSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.StatsStatus"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":               api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"compaction-interval ":      api.CLIInfo{Path: "Spec.CompactionInterval", Skip: false, Insert: "", Help: ""},
			"downsample-retention-time": api.CLIInfo{Path: "Spec.DownSampleRetentionTime", Skip: false, Insert: "", Help: ""},
			"kind":           api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"retention-time": api.CLIInfo{Path: "Spec.RetentionTime", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.StatsSpec": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsSpec{}) },
		Fields: map[string]api.Field{
			"CompactionInterval": api.Field{Name: "CompactionInterval", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "compaction-interval ", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"RetentionTime": api.Field{Name: "RetentionTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "retention-time", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"DownSampleRetentionTime": api.Field{Name: "DownSampleRetentionTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "downsample-retention-time", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"monitoring.StatsStatus": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsStatus{}) },
		Fields:    map[string]api.Field{},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapTelemetry)
}
