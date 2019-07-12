// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventtypes.proto

/*
Package eventtypes is a generated protocol buffer package.

It is generated from these files:
	eventtypes.proto

It has these top-level messages:
*/
package eventtypes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import gogoproto "github.com/gogo/protobuf/gogoproto"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import eventattrs "github.com/pensando/sw/events/generated/eventattrs"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// goproto_enum_prefix from public import gogo.proto
var E_GoprotoEnumPrefix = gogoproto.E_GoprotoEnumPrefix

// goproto_enum_stringer from public import gogo.proto
var E_GoprotoEnumStringer = gogoproto.E_GoprotoEnumStringer

// enum_stringer from public import gogo.proto
var E_EnumStringer = gogoproto.E_EnumStringer

// enum_customname from public import gogo.proto
var E_EnumCustomname = gogoproto.E_EnumCustomname

// enumdecl from public import gogo.proto
var E_Enumdecl = gogoproto.E_Enumdecl

// enumvalue_customname from public import gogo.proto
var E_EnumvalueCustomname = gogoproto.E_EnumvalueCustomname

// goproto_getters_all from public import gogo.proto
var E_GoprotoGettersAll = gogoproto.E_GoprotoGettersAll

// goproto_enum_prefix_all from public import gogo.proto
var E_GoprotoEnumPrefixAll = gogoproto.E_GoprotoEnumPrefixAll

// goproto_stringer_all from public import gogo.proto
var E_GoprotoStringerAll = gogoproto.E_GoprotoStringerAll

// verbose_equal_all from public import gogo.proto
var E_VerboseEqualAll = gogoproto.E_VerboseEqualAll

// face_all from public import gogo.proto
var E_FaceAll = gogoproto.E_FaceAll

// gostring_all from public import gogo.proto
var E_GostringAll = gogoproto.E_GostringAll

// populate_all from public import gogo.proto
var E_PopulateAll = gogoproto.E_PopulateAll

// stringer_all from public import gogo.proto
var E_StringerAll = gogoproto.E_StringerAll

// onlyone_all from public import gogo.proto
var E_OnlyoneAll = gogoproto.E_OnlyoneAll

// equal_all from public import gogo.proto
var E_EqualAll = gogoproto.E_EqualAll

// description_all from public import gogo.proto
var E_DescriptionAll = gogoproto.E_DescriptionAll

// testgen_all from public import gogo.proto
var E_TestgenAll = gogoproto.E_TestgenAll

// benchgen_all from public import gogo.proto
var E_BenchgenAll = gogoproto.E_BenchgenAll

// marshaler_all from public import gogo.proto
var E_MarshalerAll = gogoproto.E_MarshalerAll

// unmarshaler_all from public import gogo.proto
var E_UnmarshalerAll = gogoproto.E_UnmarshalerAll

// stable_marshaler_all from public import gogo.proto
var E_StableMarshalerAll = gogoproto.E_StableMarshalerAll

// sizer_all from public import gogo.proto
var E_SizerAll = gogoproto.E_SizerAll

// goproto_enum_stringer_all from public import gogo.proto
var E_GoprotoEnumStringerAll = gogoproto.E_GoprotoEnumStringerAll

// enum_stringer_all from public import gogo.proto
var E_EnumStringerAll = gogoproto.E_EnumStringerAll

// unsafe_marshaler_all from public import gogo.proto
var E_UnsafeMarshalerAll = gogoproto.E_UnsafeMarshalerAll

// unsafe_unmarshaler_all from public import gogo.proto
var E_UnsafeUnmarshalerAll = gogoproto.E_UnsafeUnmarshalerAll

// goproto_extensions_map_all from public import gogo.proto
var E_GoprotoExtensionsMapAll = gogoproto.E_GoprotoExtensionsMapAll

// goproto_unrecognized_all from public import gogo.proto
var E_GoprotoUnrecognizedAll = gogoproto.E_GoprotoUnrecognizedAll

// gogoproto_import from public import gogo.proto
var E_GogoprotoImport = gogoproto.E_GogoprotoImport

// protosizer_all from public import gogo.proto
var E_ProtosizerAll = gogoproto.E_ProtosizerAll

// compare_all from public import gogo.proto
var E_CompareAll = gogoproto.E_CompareAll

// typedecl_all from public import gogo.proto
var E_TypedeclAll = gogoproto.E_TypedeclAll

// enumdecl_all from public import gogo.proto
var E_EnumdeclAll = gogoproto.E_EnumdeclAll

// goproto_registration from public import gogo.proto
var E_GoprotoRegistration = gogoproto.E_GoprotoRegistration

// goproto_getters from public import gogo.proto
var E_GoprotoGetters = gogoproto.E_GoprotoGetters

// goproto_stringer from public import gogo.proto
var E_GoprotoStringer = gogoproto.E_GoprotoStringer

// verbose_equal from public import gogo.proto
var E_VerboseEqual = gogoproto.E_VerboseEqual

// face from public import gogo.proto
var E_Face = gogoproto.E_Face

// gostring from public import gogo.proto
var E_Gostring = gogoproto.E_Gostring

// populate from public import gogo.proto
var E_Populate = gogoproto.E_Populate

// stringer from public import gogo.proto
var E_Stringer = gogoproto.E_Stringer

// onlyone from public import gogo.proto
var E_Onlyone = gogoproto.E_Onlyone

// equal from public import gogo.proto
var E_Equal = gogoproto.E_Equal

// description from public import gogo.proto
var E_Description = gogoproto.E_Description

// testgen from public import gogo.proto
var E_Testgen = gogoproto.E_Testgen

// benchgen from public import gogo.proto
var E_Benchgen = gogoproto.E_Benchgen

// marshaler from public import gogo.proto
var E_Marshaler = gogoproto.E_Marshaler

// unmarshaler from public import gogo.proto
var E_Unmarshaler = gogoproto.E_Unmarshaler

// stable_marshaler from public import gogo.proto
var E_StableMarshaler = gogoproto.E_StableMarshaler

// sizer from public import gogo.proto
var E_Sizer = gogoproto.E_Sizer

// unsafe_marshaler from public import gogo.proto
var E_UnsafeMarshaler = gogoproto.E_UnsafeMarshaler

// unsafe_unmarshaler from public import gogo.proto
var E_UnsafeUnmarshaler = gogoproto.E_UnsafeUnmarshaler

// goproto_extensions_map from public import gogo.proto
var E_GoprotoExtensionsMap = gogoproto.E_GoprotoExtensionsMap

// goproto_unrecognized from public import gogo.proto
var E_GoprotoUnrecognized = gogoproto.E_GoprotoUnrecognized

// protosizer from public import gogo.proto
var E_Protosizer = gogoproto.E_Protosizer

// compare from public import gogo.proto
var E_Compare = gogoproto.E_Compare

// typedecl from public import gogo.proto
var E_Typedecl = gogoproto.E_Typedecl

// nullable from public import gogo.proto
var E_Nullable = gogoproto.E_Nullable

// embed from public import gogo.proto
var E_Embed = gogoproto.E_Embed

// customtype from public import gogo.proto
var E_Customtype = gogoproto.E_Customtype

// customname from public import gogo.proto
var E_Customname = gogoproto.E_Customname

// jsontag from public import gogo.proto
var E_Jsontag = gogoproto.E_Jsontag

// moretags from public import gogo.proto
var E_Moretags = gogoproto.E_Moretags

// casttype from public import gogo.proto
var E_Casttype = gogoproto.E_Casttype

// castkey from public import gogo.proto
var E_Castkey = gogoproto.E_Castkey

// castvalue from public import gogo.proto
var E_Castvalue = gogoproto.E_Castvalue

// stdtime from public import gogo.proto
var E_Stdtime = gogoproto.E_Stdtime

// stdduration from public import gogo.proto
var E_Stdduration = gogoproto.E_Stdduration

//
type EventType int32

const (
	// ----------------------------- Cluster events -------------------------- //
	ELECTION_STARTED EventType = 0
	//
	ELECTION_CANCELLED EventType = 1
	//
	ELECTION_NOTIFICATION_FAILED EventType = 2
	//
	ELECTION_STOPPED EventType = 3
	//
	LEADER_ELECTED EventType = 4
	//
	LEADER_LOST EventType = 5
	//
	LEADER_CHANGED EventType = 6
	// ------------------------------- Node events ----------------------------- //
	NODE_JOINED EventType = 7
	//
	NODE_DISJOINED EventType = 8
	//
	NODE_HEALTHY EventType = 9
	//
	NODE_UNREACHABLE EventType = 10
	// ------------------------------- Quorum events ----------------------------- //
	QUORUM_MEMBER_ADD EventType = 15
	//
	QUORUM_MEMBER_REMOVE EventType = 16
	//
	QUORUM_MEMBER_HEALTHY EventType = 17
	//
	QUORUM_MEMBER_UNHEALTHY EventType = 18
	//
	UNSUPPORTED_QUORUM_SIZE EventType = 19
	//
	QUORUM_UNHEALTHY EventType = 20
	// ------------------------------- Diagnostic events ----------------------------- //
	MODULE_CREATION_FAILED EventType = 24
	// -------------------------------- Host/NIC events -------------------------- //
	NIC_ADMITTED EventType = 100
	//
	NIC_REJECTED EventType = 101
	//
	NIC_UNREACHABLE EventType = 102
	//
	NIC_HEALTHY EventType = 103
	//
	NIC_UNHEALTHY EventType = 104
	//
	HOST_SMART_NIC_SPEC_CONFLICT EventType = 105
	// ----------------------------- API Gateway events ---------------------- //
	AUTO_GENERATED_TLS_CERT EventType = 200
	// --------------------------- Auth/Audit events ------------------------- //
	LOGIN_FAILED EventType = 300
	//
	AUDITING_FAILED EventType = 301
	// --------------------------- HAL/Link events --------------------------- //
	LINK_UP EventType = 400
	//
	LINK_DOWN EventType = 401
	// ------------------------------ System events -------------------------- //
	SERVICE_STARTED EventType = 500
	//
	SERVICE_STOPPED EventType = 501
	//
	NAPLES_SERVICE_STOPPED EventType = 502
	//
	SERVICE_PENDING EventType = 503
	//
	SERVICE_RUNNING EventType = 504
	//
	SERVICE_UNRESPONSIVE EventType = 505
	//
	SYSTEM_COLDBOOT EventType = 506
	// ------------------------------ Rollout events -------------------------- //
	ROLLOUT_STARTED EventType = 701
	//
	ROLLOUT_SUCCESS EventType = 702
	//
	ROLLOUT_FAILED EventType = 703
	//
	ROLLOUT_SUSPENDED EventType = 704
)

var EventType_name = map[int32]string{
	0:   "ELECTION_STARTED",
	1:   "ELECTION_CANCELLED",
	2:   "ELECTION_NOTIFICATION_FAILED",
	3:   "ELECTION_STOPPED",
	4:   "LEADER_ELECTED",
	5:   "LEADER_LOST",
	6:   "LEADER_CHANGED",
	7:   "NODE_JOINED",
	8:   "NODE_DISJOINED",
	9:   "NODE_HEALTHY",
	10:  "NODE_UNREACHABLE",
	15:  "QUORUM_MEMBER_ADD",
	16:  "QUORUM_MEMBER_REMOVE",
	17:  "QUORUM_MEMBER_HEALTHY",
	18:  "QUORUM_MEMBER_UNHEALTHY",
	19:  "UNSUPPORTED_QUORUM_SIZE",
	20:  "QUORUM_UNHEALTHY",
	24:  "MODULE_CREATION_FAILED",
	100: "NIC_ADMITTED",
	101: "NIC_REJECTED",
	102: "NIC_UNREACHABLE",
	103: "NIC_HEALTHY",
	104: "NIC_UNHEALTHY",
	105: "HOST_SMART_NIC_SPEC_CONFLICT",
	200: "AUTO_GENERATED_TLS_CERT",
	300: "LOGIN_FAILED",
	301: "AUDITING_FAILED",
	400: "LINK_UP",
	401: "LINK_DOWN",
	500: "SERVICE_STARTED",
	501: "SERVICE_STOPPED",
	502: "NAPLES_SERVICE_STOPPED",
	503: "SERVICE_PENDING",
	504: "SERVICE_RUNNING",
	505: "SERVICE_UNRESPONSIVE",
	506: "SYSTEM_COLDBOOT",
	701: "ROLLOUT_STARTED",
	702: "ROLLOUT_SUCCESS",
	703: "ROLLOUT_FAILED",
	704: "ROLLOUT_SUSPENDED",
}
var EventType_value = map[string]int32{
	"ELECTION_STARTED":             0,
	"ELECTION_CANCELLED":           1,
	"ELECTION_NOTIFICATION_FAILED": 2,
	"ELECTION_STOPPED":             3,
	"LEADER_ELECTED":               4,
	"LEADER_LOST":                  5,
	"LEADER_CHANGED":               6,
	"NODE_JOINED":                  7,
	"NODE_DISJOINED":               8,
	"NODE_HEALTHY":                 9,
	"NODE_UNREACHABLE":             10,
	"QUORUM_MEMBER_ADD":            15,
	"QUORUM_MEMBER_REMOVE":         16,
	"QUORUM_MEMBER_HEALTHY":        17,
	"QUORUM_MEMBER_UNHEALTHY":      18,
	"UNSUPPORTED_QUORUM_SIZE":      19,
	"QUORUM_UNHEALTHY":             20,
	"MODULE_CREATION_FAILED":       24,
	"NIC_ADMITTED":                 100,
	"NIC_REJECTED":                 101,
	"NIC_UNREACHABLE":              102,
	"NIC_HEALTHY":                  103,
	"NIC_UNHEALTHY":                104,
	"HOST_SMART_NIC_SPEC_CONFLICT": 105,
	"AUTO_GENERATED_TLS_CERT":      200,
	"LOGIN_FAILED":                 300,
	"AUDITING_FAILED":              301,
	"LINK_UP":                      400,
	"LINK_DOWN":                    401,
	"SERVICE_STARTED":              500,
	"SERVICE_STOPPED":              501,
	"NAPLES_SERVICE_STOPPED":       502,
	"SERVICE_PENDING":              503,
	"SERVICE_RUNNING":              504,
	"SERVICE_UNRESPONSIVE":         505,
	"SYSTEM_COLDBOOT":              506,
	"ROLLOUT_STARTED":              701,
	"ROLLOUT_SUCCESS":              702,
	"ROLLOUT_FAILED":               703,
	"ROLLOUT_SUSPENDED":            704,
}

func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEventtypes, []int{0} }

var E_Severity = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*eventattrs.Severity)(nil),
	Field:         10005,
	Name:          "eventtypes.severity",
	Tag:           "varint,10005,opt,name=severity,enum=eventattrs.Severity",
	Filename:      "eventtypes.proto",
}

var E_Category = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*eventattrs.Category)(nil),
	Field:         10006,
	Name:          "eventtypes.category",
	Tag:           "varint,10006,opt,name=category,enum=eventattrs.Category",
	Filename:      "eventtypes.proto",
}

var E_Desc = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         10007,
	Name:          "eventtypes.desc",
	Tag:           "bytes,10007,opt,name=desc",
	Filename:      "eventtypes.proto",
}

func init() {
	proto.RegisterEnum("eventtypes.EventType", EventType_name, EventType_value)
	proto.RegisterExtension(E_Severity)
	proto.RegisterExtension(E_Category)
	proto.RegisterExtension(E_Desc)
}
func (x EventType) String() string {
	s, ok := EventType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("eventtypes.proto", fileDescriptorEventtypes) }

var fileDescriptorEventtypes = []byte{
	// 1569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xcf, 0x6f, 0xdc, 0xc6,
	0x15, 0xc7, 0xb5, 0x5a, 0x39, 0x89, 0xc7, 0xae, 0x44, 0x33, 0xaa, 0xe5, 0x30, 0xee, 0x66, 0x92,
	0xb6, 0xb6, 0x1b, 0x23, 0x72, 0x9b, 0xc8, 0x2d, 0x1c, 0xb4, 0x8e, 0x69, 0x72, 0x2c, 0x31, 0xe5,
	0x92, 0x1b, 0x92, 0xeb, 0xc0, 0xed, 0x81, 0xe5, 0x92, 0x6f, 0x77, 0x27, 0xe1, 0x72, 0x36, 0x9c,
	0xa1, 0x04, 0xf5, 0x2f, 0xe8, 0xa9, 0x68, 0x0f, 0x6d, 0xff, 0x80, 0xfa, 0xa0, 0x43, 0x0b, 0xf8,
	0x1c, 0xa0, 0x3f, 0x8e, 0x39, 0xf6, 0x4f, 0x28, 0xfc, 0x07, 0x14, 0x10, 0xd0, 0xdf, 0xa7, 0x62,
	0x66, 0xc9, 0x15, 0xa5, 0xd8, 0x85, 0x8f, 0x24, 0xe6, 0xfb, 0x99, 0xef, 0x7b, 0xf3, 0xde, 0x9b,
	0x41, 0x1a, 0xec, 0x43, 0x21, 0xc4, 0xe1, 0x1c, 0xf8, 0xf6, 0xbc, 0x64, 0x82, 0xe9, 0xe8, 0xe4,
	0x8f, 0x81, 0x26, 0x6c, 0xc2, 0x16, 0xff, 0x0d, 0x3c, 0x61, 0x6c, 0x92, 0xc3, 0x2d, 0xf5, 0x35,
	0xaa, 0xc6, 0xb7, 0x32, 0xe0, 0x69, 0x49, 0xe7, 0x82, 0x95, 0xf5, 0x0a, 0x2d, 0x11, 0xa2, 0xa4,
	0xa3, 0x4a, 0x34, 0xac, 0xb7, 0xff, 0xb6, 0x89, 0x10, 0x91, 0xb8, 0x48, 0xe2, 0xf4, 0x3d, 0xa4,
	0x11, 0x97, 0x58, 0x91, 0xe3, 0x7b, 0x71, 0x18, 0x99, 0x41, 0x44, 0x6c, 0x6d, 0xc5, 0x78, 0xf7,
	0xe8, 0x78, 0x6d, 0xe5, 0xc9, 0xf1, 0xda, 0xca, 0xe7, 0xc7, 0x6b, 0xd7, 0x5c, 0x48, 0x32, 0x28,
	0x31, 0xe4, 0x90, 0x0a, 0xca, 0x0a, 0xcc, 0x45, 0x52, 0x0a, 0xc8, 0x30, 0x2d, 0xb0, 0x98, 0x02,
	0x4e, 0xf3, 0x8a, 0x0b, 0x28, 0xf5, 0x1f, 0x20, 0x7d, 0x49, 0xb2, 0x4c, 0xcf, 0x22, 0xae, 0x4b,
	0x6c, 0xad, 0x63, 0x7c, 0xf3, 0xe8, 0x78, 0xad, 0x53, 0xb3, 0x5e, 0x3b, 0xcb, 0x4a, 0x93, 0x22,
	0x85, 0x3c, 0x87, 0x4c, 0x7f, 0x84, 0xae, 0x2e, 0xe5, 0x9e, 0x1f, 0x39, 0x0f, 0x1c, 0xcb, 0x54,
	0x1f, 0x0f, 0x4c, 0x47, 0x82, 0x56, 0x8d, 0xef, 0xb5, 0x40, 0x37, 0x1f, 0x24, 0x34, 0x87, 0x0c,
	0x0b, 0x86, 0x39, 0x14, 0x19, 0xce, 0xcf, 0x70, 0x0b, 0x26, 0xe8, 0x98, 0xa6, 0x89, 0xfc, 0xd0,
	0xef, 0x9c, 0x8a, 0xd1, 0x1f, 0x0c, 0x88, 0xad, 0x75, 0x8d, 0xaf, 0xb7, 0x70, 0x5b, 0x5f, 0x8e,
	0x91, 0xcd, 0xe7, 0x90, 0xe9, 0x77, 0xd1, 0xba, 0x4b, 0x4c, 0x9b, 0x04, 0xb1, 0x22, 0x10, 0x5b,
	0x5b, 0x33, 0xde, 0x6e, 0x25, 0xa7, 0xd7, 0x16, 0x42, 0x86, 0xc7, 0xac, 0x3c, 0x95, 0x14, 0x82,
	0x2e, 0xd4, 0x7a, 0xd7, 0x0f, 0x23, 0xed, 0x9c, 0xb1, 0xd3, 0x12, 0xdf, 0xf0, 0x58, 0x06, 0x38,
	0x67, 0x5c, 0xd4, 0xfe, 0xf9, 0x94, 0xce, 0x71, 0x56, 0x95, 0xb4, 0x98, 0x28, 0x4a, 0x63, 0xa7,
	0x65, 0xc3, 0xda, 0x33, 0xbd, 0x5d, 0x62, 0x6b, 0x2f, 0x3d, 0xd3, 0x46, 0x3a, 0x4d, 0x8a, 0xc9,
	0xc9, 0xd1, 0x2c, 0xf5, 0x3b, 0xe8, 0x82, 0xe7, 0xdb, 0x24, 0xfe, 0xd0, 0x77, 0x3c, 0x62, 0x6b,
	0x2f, 0xab, 0xe0, 0x1b, 0xf1, 0x96, 0xb2, 0xf1, 0x09, 0xa3, 0x85, 0x4c, 0x68, 0xcb, 0xfc, 0x07,
	0x68, 0x5d, 0xa9, 0x6c, 0x27, 0xac, 0x85, 0xaf, 0x18, 0x37, 0x5b, 0x59, 0x7b, 0x43, 0x09, 0x33,
	0xca, 0x6b, 0xed, 0xb8, 0x64, 0xb3, 0x53, 0x80, 0x77, 0xd0, 0x45, 0x05, 0xd8, 0x23, 0xa6, 0x1b,
	0xed, 0x3d, 0xd2, 0xce, 0x1b, 0xaf, 0xb7, 0xf6, 0xdd, 0x50, 0x72, 0xca, 0xf1, 0x14, 0x92, 0x5c,
	0x4c, 0x0f, 0xf5, 0xdb, 0x48, 0x53, 0xcb, 0x87, 0x5e, 0x40, 0x4c, 0x6b, 0xcf, 0xbc, 0xef, 0x12,
	0x0d, 0x19, 0x6f, 0x1c, 0x1d, 0xaf, 0xad, 0xd6, 0x92, 0x57, 0x1b, 0x49, 0x55, 0x94, 0x90, 0xa4,
	0xd3, 0x64, 0x94, 0x83, 0x7e, 0x07, 0x5d, 0xfa, 0x68, 0xe8, 0x07, 0xc3, 0x7e, 0xdc, 0x27, 0xfd,
	0xfb, 0x24, 0x88, 0x4d, 0xdb, 0xd6, 0x36, 0x8c, 0xb7, 0x5a, 0x5b, 0x5d, 0x36, 0xb3, 0x0c, 0x32,
	0x3c, 0x83, 0xd9, 0x08, 0x4a, 0x59, 0x34, 0x9f, 0x55, 0xac, 0xac, 0x66, 0xfa, 0x3d, 0xb4, 0x79,
	0x5a, 0x1a, 0x90, 0xbe, 0xff, 0x90, 0x68, 0x9a, 0x71, 0xad, 0xa5, 0x36, 0x02, 0x98, 0xb1, 0xfd,
	0x13, 0xbd, 0x8a, 0xb3, 0x26, 0x58, 0xe8, 0xab, 0xa7, 0x09, 0x4d, 0xac, 0x97, 0x8c, 0x1b, 0x2d,
	0xc4, 0xd5, 0x8f, 0xd4, 0xf2, 0x86, 0x40, 0x39, 0x2e, 0xd8, 0xc1, 0x32, 0xf0, 0x5d, 0xb4, 0x75,
	0x1a, 0x32, 0xf4, 0x1a, 0x8c, 0xae, 0xce, 0xb9, 0x89, 0xbf, 0xf7, 0x4c, 0x4c, 0x55, 0x34, 0x20,
	0x1f, 0x6d, 0x0d, 0xbd, 0x70, 0x38, 0x18, 0xf8, 0xb2, 0x91, 0xe3, 0x1a, 0x1a, 0x3a, 0x3f, 0x22,
	0xda, 0xab, 0xaa, 0xa9, 0x1b, 0xd0, 0xb5, 0x1a, 0xc4, 0xe9, 0x4f, 0x55, 0x3e, 0x47, 0x90, 0xb3,
	0x03, 0xcc, 0xab, 0xf9, 0x9c, 0xa9, 0xde, 0x9e, 0xd1, 0x82, 0xce, 0xaa, 0x99, 0xde, 0x47, 0x5a,
	0x0d, 0x39, 0xb1, 0xb4, 0xa9, 0x3a, 0xb1, 0x21, 0xdd, 0xac, 0x49, 0x19, 0x03, 0xe9, 0x46, 0xe0,
	0x69, 0xb2, 0x0f, 0x18, 0x0a, 0x56, 0x4d, 0xa6, 0x4d, 0x80, 0xb5, 0x5f, 0xae, 0x7b, 0xe8, 0x72,
	0xdf, 0xb7, 0x87, 0x2e, 0x89, 0xad, 0x80, 0x9c, 0x6a, 0xef, 0x2b, 0xca, 0x5e, 0x53, 0x59, 0xd7,
	0xfa, 0x2c, 0xab, 0x72, 0xc0, 0x69, 0x09, 0xaa, 0x85, 0x55, 0x5f, 0x65, 0x34, 0x99, 0x14, 0x8c,
	0x0b, 0x9a, 0x72, 0x3c, 0x56, 0xed, 0xaf, 0xdf, 0x41, 0x17, 0x3d, 0xc7, 0x8a, 0x4d, 0xbb, 0xef,
	0x44, 0xb2, 0x39, 0x33, 0xe3, 0x7a, 0x2b, 0xe9, 0xaf, 0x7b, 0x8e, 0x85, 0x93, 0x6c, 0x46, 0x85,
	0x58, 0x8c, 0x8a, 0x76, 0x6d, 0xee, 0x2e, 0xa4, 0x01, 0xf9, 0x70, 0xd1, 0xd7, 0x60, 0xdc, 0x6e,
	0x19, 0xf8, 0x96, 0x89, 0xc3, 0x59, 0x52, 0x8a, 0x86, 0xc0, 0xb9, 0x74, 0x51, 0xc2, 0x67, 0x15,
	0x70, 0x81, 0x0f, 0x12, 0x8e, 0x4b, 0xf8, 0x44, 0xb5, 0xbc, 0xfe, 0x1e, 0xda, 0x90, 0xa0, 0x76,
	0xd1, 0x8e, 0x8d, 0x5e, 0x2b, 0x43, 0xba, 0x84, 0x9c, 0xa9, 0xd9, 0xbb, 0xe8, 0x82, 0x14, 0x35,
	0x29, 0x9d, 0x18, 0xef, 0xb4, 0x7c, 0xbf, 0x29, 0x05, 0x8b, 0xd4, 0x61, 0x51, 0x26, 0x05, 0xa7,
	0x32, 0x07, 0x8b, 0x10, 0xae, 0x8b, 0xb2, 0x82, 0xeb, 0xba, 0x89, 0xbe, 0xb2, 0xd8, 0xb4, 0x21,
	0x4c, 0x8d, 0xed, 0x96, 0xfd, 0xb7, 0xfe, 0x1f, 0x61, 0x9c, 0xe4, 0x1c, 0xae, 0xeb, 0x3f, 0xef,
	0xa0, 0xab, 0x7b, 0x7e, 0x18, 0xc5, 0x61, 0xdf, 0x0c, 0xa2, 0x58, 0xe2, 0xc2, 0x01, 0xb1, 0x62,
	0xcb, 0xf7, 0x1e, 0xb8, 0x8e, 0x15, 0x69, 0xd4, 0xc8, 0x5b, 0xc8, 0x9f, 0x44, 0x53, 0xc0, 0xfc,
	0x90, 0x0b, 0x98, 0xe1, 0x69, 0xc2, 0x71, 0x06, 0x62, 0x31, 0xf2, 0x12, 0x9c, 0xb2, 0x62, 0x9c,
	0xd3, 0x54, 0xe0, 0x11, 0x88, 0x03, 0x80, 0xc5, 0xec, 0x59, 0xa6, 0x8f, 0xcf, 0x21, 0x5d, 0xce,
	0x62, 0x8e, 0xd9, 0x18, 0x67, 0x74, 0x3c, 0x86, 0x12, 0x0a, 0x81, 0xf7, 0xe4, 0xec, 0x63, 0x23,
	0x99, 0x49, 0xae, 0x8f, 0xd1, 0x96, 0x39, 0x8c, 0xfc, 0x78, 0x97, 0x78, 0x24, 0x30, 0x65, 0xfd,
	0x46, 0x6e, 0x18, 0x5b, 0x24, 0x88, 0xb4, 0x2f, 0x3a, 0xc6, 0x5e, 0xcb, 0xcb, 0xf7, 0xcd, 0x4a,
	0x30, 0x3c, 0x81, 0x02, 0xca, 0x44, 0x5a, 0x48, 0xa1, 0xac, 0xc7, 0x7d, 0x5d, 0xcc, 0x72, 0x80,
	0x56, 0xbc, 0x9e, 0xc7, 0xe6, 0xc0, 0xc1, 0xbb, 0x89, 0x80, 0x83, 0xe4, 0x10, 0x47, 0x6e, 0xa8,
	0x7f, 0x1b, 0x5d, 0x74, 0xfd, 0x5d, 0x67, 0x59, 0x7a, 0xbf, 0x5b, 0x35, 0xbe, 0xd6, 0xca, 0xfe,
	0xa5, 0x21, 0x87, 0x12, 0xe7, 0x6c, 0x42, 0x8b, 0xa6, 0xcc, 0xee, 0xa2, 0x0d, 0x73, 0x68, 0x3b,
	0x91, 0xe3, 0xed, 0x36, 0xa2, 0xdf, 0xaf, 0xaa, 0xfe, 0x6e, 0xce, 0xf8, 0xea, 0xc7, 0x25, 0x15,
	0x72, 0x57, 0x36, 0xc6, 0x66, 0x95, 0x51, 0xa1, 0xae, 0xd8, 0x46, 0x7f, 0x13, 0xbd, 0xec, 0x3a,
	0xde, 0x0f, 0xe3, 0xe1, 0x40, 0xfb, 0x45, 0x77, 0xb9, 0x59, 0x47, 0x6e, 0x36, 0x60, 0xa5, 0x90,
	0x9e, 0x73, 0x5a, 0x7c, 0x0a, 0x19, 0xae, 0xe6, 0xfa, 0x0e, 0x3a, 0xaf, 0x16, 0xdb, 0xfe, 0xc7,
	0x9e, 0xf6, 0xcb, 0xae, 0xf1, 0x8d, 0x3a, 0x70, 0xb9, 0xfc, 0x8a, 0x5a, 0x2e, 0xd7, 0xca, 0x5b,
	0x58, 0x54, 0x5c, 0x4a, 0x33, 0x76, 0x50, 0xe8, 0xdf, 0x41, 0x1b, 0x21, 0x09, 0x1e, 0x3a, 0x16,
	0x59, 0x5e, 0xe3, 0x7f, 0xef, 0xaa, 0x71, 0xdb, 0x7d, 0x72, 0xbc, 0xb6, 0x2a, 0xc7, 0x6d, 0x08,
	0xe5, 0x3e, 0x4d, 0xa1, 0xb9, 0xbf, 0x4f, 0x4b, 0x16, 0xb7, 0xe2, 0x3f, 0x16, 0x92, 0xce, 0x97,
	0x25, 0x8b, 0xeb, 0xf0, 0x03, 0x74, 0xd9, 0x33, 0x07, 0x2e, 0x09, 0xe3, 0xb3, 0xca, 0x7f, 0x76,
	0xd5, 0xc0, 0x5d, 0xad, 0x95, 0x97, 0xbd, 0x64, 0x9e, 0x03, 0xc7, 0xfc, 0x0c, 0xa0, 0xb5, 0xe7,
	0x80, 0x78, 0xb6, 0xe3, 0xed, 0x6a, 0xff, 0x7a, 0x8e, 0xcd, 0x39, 0x14, 0x19, 0x2d, 0x26, 0x6d,
	0x49, 0x30, 0xf4, 0x3c, 0x29, 0xf9, 0xf7, 0x73, 0x24, 0x65, 0x55, 0x14, 0x52, 0xf2, 0x63, 0xb4,
	0xd9, 0x48, 0x64, 0x5b, 0x86, 0x03, 0xdf, 0x0b, 0x9d, 0x87, 0x44, 0xfb, 0x4f, 0xd7, 0xb8, 0xd7,
	0x32, 0xb9, 0xd3, 0xe8, 0x64, 0x67, 0xf2, 0x39, 0x2b, 0x38, 0xdd, 0x07, 0x9c, 0x55, 0x20, 0xdb,
	0x24, 0x4f, 0xd2, 0x4f, 0xe5, 0x69, 0xd6, 0x25, 0x5f, 0x02, 0x67, 0x55, 0x99, 0x02, 0xd7, 0x77,
	0xd0, 0x46, 0xf8, 0x28, 0x8c, 0x48, 0x3f, 0xb6, 0x7c, 0xd7, 0xbe, 0xef, 0xfb, 0x91, 0xf6, 0xdf,
	0xae, 0x6a, 0xf8, 0x26, 0x6d, 0x7a, 0xb8, 0xd0, 0xa4, 0x2c, 0xcf, 0xf0, 0x88, 0xb1, 0x7a, 0x4a,
	0x04, 0xbe, 0xeb, 0xfa, 0xc3, 0x68, 0x79, 0x3e, 0x7f, 0x38, 0xb7, 0x2c, 0x85, 0xae, 0x2c, 0x85,
	0x80, 0xe5, 0x39, 0xab, 0x04, 0xa6, 0x05, 0x15, 0x54, 0xd6, 0xb5, 0x7e, 0xaf, 0x25, 0x1a, 0x5a,
	0x16, 0x09, 0x43, 0xed, 0x8f, 0xe7, 0x96, 0x17, 0xbf, 0x14, 0xf5, 0x1a, 0x11, 0xaf, 0xd2, 0x14,
	0x38, 0x1f, 0x57, 0x79, 0x7e, 0x88, 0x53, 0x36, 0x9b, 0xe7, 0x20, 0x09, 0xb7, 0xd0, 0x7a, 0x43,
	0xa8, 0x0b, 0xf7, 0x4f, 0xe7, 0x0c, 0xa3, 0xce, 0x81, 0x04, 0xac, 0x37, 0x80, 0xba, 0x54, 0x6f,
	0xa3, 0x4b, 0x27, 0x5b, 0x86, 0xf2, 0x8c, 0x88, 0xad, 0xfd, 0xf9, 0x39, 0x4e, 0x79, 0xc5, 0xe5,
	0x29, 0x41, 0x66, 0xbc, 0xf6, 0xb3, 0xdf, 0xf6, 0x56, 0x8e, 0x1e, 0xf7, 0x56, 0x9e, 0x3c, 0xee,
	0x75, 0x3e, 0x7f, 0xdc, 0x3b, 0xbf, 0x7c, 0x62, 0xbe, 0x1f, 0xa1, 0x57, 0x38, 0xec, 0x43, 0x49,
	0xc5, 0xa1, 0xfe, 0xe6, 0xf6, 0xe2, 0xc5, 0xba, 0xdd, 0xbc, 0x58, 0xb7, 0x49, 0x51, 0xcd, 0x1e,
	0x26, 0x79, 0x05, 0xfe, 0x5c, 0xcd, 0x87, 0x2b, 0xbf, 0xf2, 0x70, 0xe7, 0xc6, 0xfa, 0xbb, 0x9b,
	0xdb, 0xea, 0xcd, 0x2b, 0x9f, 0xaf, 0x7c, 0x3b, 0xac, 0xf5, 0xc1, 0x92, 0x24, 0xa9, 0xb2, 0xdd,
	0x27, 0xac, 0x7c, 0x21, 0xea, 0xaf, 0x9f, 0x41, 0xb5, 0x6a, 0x7d, 0xb0, 0x24, 0xbd, 0xff, 0x5d,
	0xb4, 0x26, 0x9f, 0xd0, 0x2f, 0x42, 0xfc, 0x8d, 0x24, 0x9e, 0x0f, 0xd4, 0xfa, 0xfb, 0x17, 0xbf,
	0x78, 0xda, 0xeb, 0xfc, 0xe5, 0x69, 0xaf, 0xf3, 0xd7, 0xa7, 0xbd, 0xce, 0x60, 0x65, 0xf4, 0x92,
	0xd2, 0xbd, 0xf7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x59, 0xc4, 0xdf, 0xcb, 0x0b, 0x00,
	0x00,
}
