// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package stagingApiServer is a auto generated package.
Input file: staging.proto
*/
package staging

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapStaging = map[string]*api.Struct{

	"staging.Buffer": &api.Struct{
		Kind: "Buffer", APIGroup: "staging", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(Buffer{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.BufferSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.BufferStatus"},

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
			"Contact":           api.CLIInfo{Path: "Spec.Contact", Skip: false, Insert: "", Help: ""},
			"api-version":       api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"error":             api.CLIInfo{Path: "Status.Errors[].Errors", Skip: false, Insert: "", Help: ""},
			"generation-id":     api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":              api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":            api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"method":            api.CLIInfo{Path: "Status.Errors[].ItemId.Method", Skip: false, Insert: "", Help: ""},
			"name":              api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":         api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"resource-version":  api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":         api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"tenant":            api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uri":               api.CLIInfo{Path: "Status.Errors[].ItemId.URI", Skip: false, Insert: "", Help: ""},
			"uuid":              api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
			"validation-result": api.CLIInfo{Path: "Status.ValidationResult", Skip: false, Insert: "", Help: ""},
		},
	},
	"staging.BufferSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(BufferSpec{}) },
		Fields: map[string]api.Field{
			"Contact": api.Field{Name: "Contact", CLITag: api.CLIInfo{ID: "Contact", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"staging.BufferStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(BufferStatus{}) },
		Fields: map[string]api.Field{
			"ValidationResult": api.Field{Name: "ValidationResult", CLITag: api.CLIInfo{ID: "validation-result", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "validation-result", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Errors": api.Field{Name: "Errors", CLITag: api.CLIInfo{ID: "errors", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "errors", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.ValidationError"},

			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{ID: "items", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "items", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.Item"},
		},
	},
	"staging.ClearAction": &api.Struct{
		Kind: "ClearAction", APIGroup: "staging", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ClearAction{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.ClearActionSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.ClearActionStatus"},

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
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"method":           api.CLIInfo{Path: "Spec.Items[].Method", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"reason":           api.CLIInfo{Path: "Status.Reason", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"status":           api.CLIInfo{Path: "Status.Status", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uri":              api.CLIInfo{Path: "Spec.Items[].URI", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"staging.ClearActionSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ClearActionSpec{}) },
		Fields: map[string]api.Field{
			"Items": api.Field{Name: "Items", CLITag: api.CLIInfo{ID: "items", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "items", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.ItemId"},
		},
	},
	"staging.ClearActionStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ClearActionStatus{}) },
		Fields: map[string]api.Field{
			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Reason": api.Field{Name: "Reason", CLITag: api.CLIInfo{ID: "reason", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "reason", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"staging.CommitAction": &api.Struct{
		Kind: "CommitAction", APIGroup: "staging", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(CommitAction{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.CommitActionSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "staging.CommitActionStatus"},

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
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"reason":           api.CLIInfo{Path: "Status.Reason", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"status":           api.CLIInfo{Path: "Status.Status", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"staging.CommitActionSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(CommitActionSpec{}) },
		Fields: map[string]api.Field{},
	},
	"staging.CommitActionStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(CommitActionStatus{}) },
		Fields: map[string]api.Field{
			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Reason": api.Field{Name: "Reason", CLITag: api.CLIInfo{ID: "reason", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "reason", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"staging.Item": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(Item{}) },
		Fields: map[string]api.Field{
			"ItemId": api.Field{Name: "ItemId", CLITag: api.CLIInfo{ID: "Id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "staging.ItemId"},

			"Object": api.Field{Name: "Object", CLITag: api.CLIInfo{ID: "object", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "object", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Any"},

			"URI": api.Field{Name: "URI", CLITag: api.CLIInfo{ID: "uri", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uri", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Method": api.Field{Name: "Method", CLITag: api.CLIInfo{ID: "method", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "method", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"staging.ItemId": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ItemId{}) },
		Fields: map[string]api.Field{
			"URI": api.Field{Name: "URI", CLITag: api.CLIInfo{ID: "uri", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uri", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Method": api.Field{Name: "Method", CLITag: api.CLIInfo{ID: "method", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "method", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"staging.ValidationError": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ValidationError{}) },
		Fields: map[string]api.Field{
			"ItemId": api.Field{Name: "ItemId", CLITag: api.CLIInfo{ID: "Id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "staging.ItemId"},

			"Errors": api.Field{Name: "Errors", CLITag: api.CLIInfo{ID: "error", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "error", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"URI": api.Field{Name: "URI", CLITag: api.CLIInfo{ID: "uri", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uri", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Method": api.Field{Name: "Method", CLITag: api.CLIInfo{ID: "method", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "method", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

var keyMapStaging = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapStaging)
	schema.AddPaths(keyMapStaging)
}
