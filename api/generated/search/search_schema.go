// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package searchApiServer is a auto generated package.
Input file: search.proto
*/
package search

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapSearch = map[string]*api.Struct{

	"search.Category": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Category{}) },
		Fields:    map[string]api.Field{},
	},
	"search.CategoryAggregation": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(CategoryAggregation{}) },
		Fields: map[string]api.Field{
			"Categories": api.Field{Name: "Categories", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "categories", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "search.KindAggregation"},
		},
	},
	"search.CategoryAggregation.CategoriesEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.KindAggregation"},
		},
	},
	"search.CategoryPreview": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(CategoryPreview{}) },
		Fields: map[string]api.Field{
			"Categories": api.Field{Name: "Categories", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "categories", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "search.KindPreview"},
		},
	},
	"search.CategoryPreview.CategoriesEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.KindPreview"},
		},
	},
	"search.Entry": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Entry{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"O": api.Field{Name: "O", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},
		},
	},
	"search.EntryList": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(EntryList{}) },
		Fields: map[string]api.Field{
			"Entries": api.Field{Name: "Entries", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "entries", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.Entry"},
		},
	},
	"search.Error": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Error{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Reason": api.Field{Name: "Reason", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "reason", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"search.Kind": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(Kind{}) },
		Fields:    map[string]api.Field{},
	},
	"search.KindAggregation": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(KindAggregation{}) },
		Fields: map[string]api.Field{
			"Kinds": api.Field{Name: "Kinds", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kinds", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "search.EntryList"},
		},
	},
	"search.KindAggregation.KindsEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.EntryList"},
		},
	},
	"search.KindPreview": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(KindPreview{}) },
		Fields: map[string]api.Field{
			"Kinds": api.Field{Name: "Kinds", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kinds", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "TYPE_INT64"},
		},
	},
	"search.KindPreview.KindsEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT64"},
		},
	},
	"search.PolicyMatchEntry": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(PolicyMatchEntry{}) },
		Fields: map[string]api.Field{
			"Rule": api.Field{Name: "Rule", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rule", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGRule"},

			"Index": api.Field{Name: "Index", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "index", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},
		},
	},
	"search.PolicySearchRequest": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(PolicySearchRequest{}) },
		Fields: map[string]api.Field{
			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"SGPolicy": api.Field{Name: "SGPolicy", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "sg-policy", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"App": api.Field{Name: "App", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "app", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"FromIPAddress": api.Field{Name: "FromIPAddress", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from-ip-address", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ToIPAddress": api.Field{Name: "ToIPAddress", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "to-ip-address", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"FromSecurityGroup": api.Field{Name: "FromSecurityGroup", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from-security-group", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ToSecurityGroup": api.Field{Name: "ToSecurityGroup", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "to-security-group", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"search.PolicySearchResponse": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(PolicySearchResponse{}) },
		Fields: map[string]api.Field{
			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Results": api.Field{Name: "Results", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "results", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "search.PolicyMatchEntry"},
		},
	},
	"search.PolicySearchResponse.ResultsEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.PolicyMatchEntry"},
		},
	},
	"search.SearchQuery": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(SearchQuery{}) },
		Fields: map[string]api.Field{
			"Texts": api.Field{Name: "Texts", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "texts", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.TextRequirement"},

			"Categories": api.Field{Name: "Categories", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "categories", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Kinds": api.Field{Name: "Kinds", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kinds", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Fields": api.Field{Name: "Fields", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "fields", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "fields.Selector"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "labels.Selector"},
		},
	},
	"search.SearchRequest": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(SearchRequest{}) },
		Fields: map[string]api.Field{
			"QueryString": api.Field{Name: "QueryString", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "query-string", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"From": api.Field{Name: "From", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT32"},

			"MaxResults": api.Field{Name: "MaxResults", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "max-results", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT32"},

			"SortBy": api.Field{Name: "SortBy", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "sort-by", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Mode": api.Field{Name: "Mode", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mode", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Query": api.Field{Name: "Query", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "query", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.SearchQuery"},
		},
	},
	"search.SearchResponse": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(SearchResponse{}) },
		Fields: map[string]api.Field{
			"TotalHits": api.Field{Name: "TotalHits", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "total-hits", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT64"},

			"ActualHits": api.Field{Name: "ActualHits", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "actual-hits", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT64"},

			"TimeTakenMsecs": api.Field{Name: "TimeTakenMsecs", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "time-taken-msecs", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT64"},

			"Error": api.Field{Name: "Error", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "error", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.Error"},

			"Entries": api.Field{Name: "Entries", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "entries", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.Entry"},

			"PreviewEntries": api.Field{Name: "PreviewEntries", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "preview-entries", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.TenantPreview"},

			"AggregatedEntries": api.Field{Name: "AggregatedEntries", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "aggregated-entries", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.TenantAggregation"},
		},
	},
	"search.TenantAggregation": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(TenantAggregation{}) },
		Fields: map[string]api.Field{
			"Tenants": api.Field{Name: "Tenants", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenants", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "search.CategoryAggregation"},
		},
	},
	"search.TenantAggregation.TenantsEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.CategoryAggregation"},
		},
	},
	"search.TenantPreview": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(TenantPreview{}) },
		Fields: map[string]api.Field{
			"Tenants": api.Field{Name: "Tenants", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenants", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: false, KeyType: "TYPE_STRING", Type: "search.CategoryPreview"},
		},
	},
	"search.TenantPreview.TenantsEntry": &api.Struct{
		Fields: map[string]api.Field{
			"key": api.Field{Name: "key", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"value": api.Field{Name: "value", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "search.CategoryPreview"},
		},
	},
	"search.TextRequirement": &api.Struct{
		GetTypeFn: func() reflect.Type { return reflect.TypeOf(TextRequirement{}) },
		Fields: map[string]api.Field{
			"Text": api.Field{Name: "Text", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "text", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSearch)
}
