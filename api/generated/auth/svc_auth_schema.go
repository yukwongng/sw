// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package authApiServer is a auto generated package.
Input file: svc_auth.proto
*/
package auth

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapSvc_auth = map[string]*api.Struct{

	"auth.AuthenticationPolicyList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AuthenticationPolicyList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AuthenticationPolicy"},
		},
	},
	"auth.AutoMsgAuthenticationPolicyWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAuthenticationPolicyWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AutoMsgAuthenticationPolicyWatchHelper.WatchEvent"},
		},
	},
	"auth.AutoMsgAuthenticationPolicyWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgAuthenticationPolicyWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AuthenticationPolicy"},
		},
	},
	"auth.AutoMsgRoleBindingWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgRoleBindingWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AutoMsgRoleBindingWatchHelper.WatchEvent"},
		},
	},
	"auth.AutoMsgRoleBindingWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgRoleBindingWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RoleBinding"},
		},
	},
	"auth.AutoMsgRoleWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgRoleWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AutoMsgRoleWatchHelper.WatchEvent"},
		},
	},
	"auth.AutoMsgRoleWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgRoleWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Role"},
		},
	},
	"auth.AutoMsgUserWatchHelper": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgUserWatchHelper{}) },
		Fields: map[string]api.Field{
			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.AutoMsgUserWatchHelper.WatchEvent"},
		},
	},
	"auth.AutoMsgUserWatchHelper.WatchEvent": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(AutoMsgUserWatchHelper_WatchEvent{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.User"},
		},
	},
	"auth.RoleBindingList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleBindingList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.RoleBinding"},
		},
	},
	"auth.RoleList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(RoleList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.Role"},
		},
	},
	"auth.UserList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(UserList{}) },
		Fields: map[string]api.Field{
			"T": api.Field{Name: "T", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ListMeta": api.Field{Name: "ListMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ListMeta"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "auth.User"},
		},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSvc_auth)
}
