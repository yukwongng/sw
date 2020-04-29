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
	//
	CLOCK_SYNC_FAILED EventType = 11
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
	// ------------------------ Config Snapshot/Restore events ------------------------ //
	CONFIG_RESTORED EventType = 30
	//
	CONFIG_RESTORE_ABORTED EventType = 31
	//
	CONFIG_RESTORE_FAILED EventType = 32
	// -------------------------------- Host/DSC events -------------------------- //
	DSC_ADMITTED EventType = 100
	//
	DSC_REJECTED EventType = 101
	//
	DSC_UNREACHABLE EventType = 102
	//
	DSC_HEALTHY EventType = 103
	//
	DSC_UNHEALTHY EventType = 104
	//
	HOST_DSC_SPEC_CONFLICT EventType = 105
	//
	DSC_DEADMITTED EventType = 106
	//
	DSC_DECOMMISSIONED EventType = 107
	// ----------------------------- API Gateway events ---------------------- //
	AUTO_GENERATED_TLS_CERT EventType = 200
	// --------------------------- Auth/Audit events ------------------------- //
	LOGIN_FAILED EventType = 300
	//
	AUDITING_FAILED EventType = 301
	//
	PASSWORD_CHANGED EventType = 302
	//
	PASSWORD_RESET EventType = 303
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
	//
	SYSTEM_RESOURCE_USAGE EventType = 507
	//
	NAPLES_FATAL_INTERRUPT EventType = 508
	//
	NAPLES_CATTRIP_INTERRUPT EventType = 509
	//
	NAPLES_OVER_TEMP EventType = 510
	//
	NAPLES_OVER_TEMP_EXIT EventType = 511
	//
	NAPLES_PANIC_EVENT EventType = 512
	//
	NAPLES_POST_DIAG_FAILURE_EVENT EventType = 513
	//
	NAPLES_INFO_PCIEHEALTH_EVENT EventType = 514
	//
	NAPLES_WARN_PCIEHEALTH_EVENT EventType = 515
	//
	NAPLES_ERR_PCIEHEALTH_EVENT EventType = 516
	// ------------------------------ Resource events -------------------------- //
	DISK_THRESHOLD_EXCEEDED EventType = 601
	// ------------------------------ Rollout events -------------------------- //
	ROLLOUT_STARTED EventType = 701
	//
	ROLLOUT_SUCCESS EventType = 702
	//
	ROLLOUT_FAILED EventType = 703
	//
	ROLLOUT_SUSPENDED EventType = 704
	// ------------------------------ Config events -------------------------- //
	CONFIG_FAIL EventType = 801
	// ------------------------------ Session Limit events -------------------------- //
	TCP_HALF_OPEN_SESSION_LIMIT_APPROACH EventType = 901
	//
	TCP_HALF_OPEN_SESSION_LIMIT_REACHED EventType = 902
	//
	UDP_ACTIVE_SESSION_LIMIT_APPROACH EventType = 903
	//
	UDP_ACTIVE_SESSION_LIMIT_REACHED EventType = 904
	//
	ICMP_ACTIVE_SESSION_LIMIT_APPROACH EventType = 905
	//
	ICMP_ACTIVE_SESSION_LIMIT_REACHED EventType = 906
	//
	OTHER_ACTIVE_SESSION_LIMIT_APPROACH EventType = 907
	//
	OTHER_ACTIVE_SESSION_LIMIT_REACHED EventType = 908
	// ------------------------- Orchestration events ------------------------- //
	ORCH_CONNECTION_ERROR EventType = 1001
	//
	ORCH_LOGIN_FAILURE EventType = 1002
	//
	ORCH_CONFIG_PUSH_FAILURE EventType = 1003
	//
	ORCH_INVALID_ACTION EventType = 1004
	//
	MIGRATION_FAILED EventType = 1005
	//
	MIGRATION_TIMED_OUT EventType = 1006
	//
	ORCH_ALREADY_MANAGED EventType = 1007
	//
	ORCH_UNSUPPORTED_VERSION EventType = 1008
	//
	ORCH_DSC_NOT_ADMITTED EventType = 1009
	//
	ORCH_DSC_MODE_INCOMPATIBLE EventType = 1010
	// ------------------------- Controller events ------------------------- //
	COLLECTOR_REACHABLE EventType = 1101
	//
	COLLECTOR_UNREACHABLE EventType = 1102
	// ------------------------- Control Plane events ------------------------- //
	BGP_SESSION_ESTABLISHED EventType = 1201
	//
	BGP_SESSION_DOWN EventType = 1202
)

var EventType_name = map[int32]string{
	0:    "ELECTION_STARTED",
	1:    "ELECTION_CANCELLED",
	2:    "ELECTION_NOTIFICATION_FAILED",
	3:    "ELECTION_STOPPED",
	4:    "LEADER_ELECTED",
	5:    "LEADER_LOST",
	6:    "LEADER_CHANGED",
	7:    "NODE_JOINED",
	8:    "NODE_DISJOINED",
	9:    "NODE_HEALTHY",
	10:   "NODE_UNREACHABLE",
	11:   "CLOCK_SYNC_FAILED",
	15:   "QUORUM_MEMBER_ADD",
	16:   "QUORUM_MEMBER_REMOVE",
	17:   "QUORUM_MEMBER_HEALTHY",
	18:   "QUORUM_MEMBER_UNHEALTHY",
	19:   "UNSUPPORTED_QUORUM_SIZE",
	20:   "QUORUM_UNHEALTHY",
	24:   "MODULE_CREATION_FAILED",
	30:   "CONFIG_RESTORED",
	31:   "CONFIG_RESTORE_ABORTED",
	32:   "CONFIG_RESTORE_FAILED",
	100:  "DSC_ADMITTED",
	101:  "DSC_REJECTED",
	102:  "DSC_UNREACHABLE",
	103:  "DSC_HEALTHY",
	104:  "DSC_UNHEALTHY",
	105:  "HOST_DSC_SPEC_CONFLICT",
	106:  "DSC_DEADMITTED",
	107:  "DSC_DECOMMISSIONED",
	200:  "AUTO_GENERATED_TLS_CERT",
	300:  "LOGIN_FAILED",
	301:  "AUDITING_FAILED",
	302:  "PASSWORD_CHANGED",
	303:  "PASSWORD_RESET",
	400:  "LINK_UP",
	401:  "LINK_DOWN",
	500:  "SERVICE_STARTED",
	501:  "SERVICE_STOPPED",
	502:  "NAPLES_SERVICE_STOPPED",
	503:  "SERVICE_PENDING",
	504:  "SERVICE_RUNNING",
	505:  "SERVICE_UNRESPONSIVE",
	506:  "SYSTEM_COLDBOOT",
	507:  "SYSTEM_RESOURCE_USAGE",
	508:  "NAPLES_FATAL_INTERRUPT",
	509:  "NAPLES_CATTRIP_INTERRUPT",
	510:  "NAPLES_OVER_TEMP",
	511:  "NAPLES_OVER_TEMP_EXIT",
	512:  "NAPLES_PANIC_EVENT",
	513:  "NAPLES_POST_DIAG_FAILURE_EVENT",
	514:  "NAPLES_INFO_PCIEHEALTH_EVENT",
	515:  "NAPLES_WARN_PCIEHEALTH_EVENT",
	516:  "NAPLES_ERR_PCIEHEALTH_EVENT",
	601:  "DISK_THRESHOLD_EXCEEDED",
	701:  "ROLLOUT_STARTED",
	702:  "ROLLOUT_SUCCESS",
	703:  "ROLLOUT_FAILED",
	704:  "ROLLOUT_SUSPENDED",
	801:  "CONFIG_FAIL",
	901:  "TCP_HALF_OPEN_SESSION_LIMIT_APPROACH",
	902:  "TCP_HALF_OPEN_SESSION_LIMIT_REACHED",
	903:  "UDP_ACTIVE_SESSION_LIMIT_APPROACH",
	904:  "UDP_ACTIVE_SESSION_LIMIT_REACHED",
	905:  "ICMP_ACTIVE_SESSION_LIMIT_APPROACH",
	906:  "ICMP_ACTIVE_SESSION_LIMIT_REACHED",
	907:  "OTHER_ACTIVE_SESSION_LIMIT_APPROACH",
	908:  "OTHER_ACTIVE_SESSION_LIMIT_REACHED",
	1001: "ORCH_CONNECTION_ERROR",
	1002: "ORCH_LOGIN_FAILURE",
	1003: "ORCH_CONFIG_PUSH_FAILURE",
	1004: "ORCH_INVALID_ACTION",
	1005: "MIGRATION_FAILED",
	1006: "MIGRATION_TIMED_OUT",
	1007: "ORCH_ALREADY_MANAGED",
	1008: "ORCH_UNSUPPORTED_VERSION",
	1009: "ORCH_DSC_NOT_ADMITTED",
	1010: "ORCH_DSC_MODE_INCOMPATIBLE",
	1101: "COLLECTOR_REACHABLE",
	1102: "COLLECTOR_UNREACHABLE",
	1201: "BGP_SESSION_ESTABLISHED",
	1202: "BGP_SESSION_DOWN",
}
var EventType_value = map[string]int32{
	"ELECTION_STARTED":                     0,
	"ELECTION_CANCELLED":                   1,
	"ELECTION_NOTIFICATION_FAILED":         2,
	"ELECTION_STOPPED":                     3,
	"LEADER_ELECTED":                       4,
	"LEADER_LOST":                          5,
	"LEADER_CHANGED":                       6,
	"NODE_JOINED":                          7,
	"NODE_DISJOINED":                       8,
	"NODE_HEALTHY":                         9,
	"NODE_UNREACHABLE":                     10,
	"CLOCK_SYNC_FAILED":                    11,
	"QUORUM_MEMBER_ADD":                    15,
	"QUORUM_MEMBER_REMOVE":                 16,
	"QUORUM_MEMBER_HEALTHY":                17,
	"QUORUM_MEMBER_UNHEALTHY":              18,
	"UNSUPPORTED_QUORUM_SIZE":              19,
	"QUORUM_UNHEALTHY":                     20,
	"MODULE_CREATION_FAILED":               24,
	"CONFIG_RESTORED":                      30,
	"CONFIG_RESTORE_ABORTED":               31,
	"CONFIG_RESTORE_FAILED":                32,
	"DSC_ADMITTED":                         100,
	"DSC_REJECTED":                         101,
	"DSC_UNREACHABLE":                      102,
	"DSC_HEALTHY":                          103,
	"DSC_UNHEALTHY":                        104,
	"HOST_DSC_SPEC_CONFLICT":               105,
	"DSC_DEADMITTED":                       106,
	"DSC_DECOMMISSIONED":                   107,
	"AUTO_GENERATED_TLS_CERT":              200,
	"LOGIN_FAILED":                         300,
	"AUDITING_FAILED":                      301,
	"PASSWORD_CHANGED":                     302,
	"PASSWORD_RESET":                       303,
	"LINK_UP":                              400,
	"LINK_DOWN":                            401,
	"SERVICE_STARTED":                      500,
	"SERVICE_STOPPED":                      501,
	"NAPLES_SERVICE_STOPPED":               502,
	"SERVICE_PENDING":                      503,
	"SERVICE_RUNNING":                      504,
	"SERVICE_UNRESPONSIVE":                 505,
	"SYSTEM_COLDBOOT":                      506,
	"SYSTEM_RESOURCE_USAGE":                507,
	"NAPLES_FATAL_INTERRUPT":               508,
	"NAPLES_CATTRIP_INTERRUPT":             509,
	"NAPLES_OVER_TEMP":                     510,
	"NAPLES_OVER_TEMP_EXIT":                511,
	"NAPLES_PANIC_EVENT":                   512,
	"NAPLES_POST_DIAG_FAILURE_EVENT":       513,
	"NAPLES_INFO_PCIEHEALTH_EVENT":         514,
	"NAPLES_WARN_PCIEHEALTH_EVENT":         515,
	"NAPLES_ERR_PCIEHEALTH_EVENT":          516,
	"DISK_THRESHOLD_EXCEEDED":              601,
	"ROLLOUT_STARTED":                      701,
	"ROLLOUT_SUCCESS":                      702,
	"ROLLOUT_FAILED":                       703,
	"ROLLOUT_SUSPENDED":                    704,
	"CONFIG_FAIL":                          801,
	"TCP_HALF_OPEN_SESSION_LIMIT_APPROACH": 901,
	"TCP_HALF_OPEN_SESSION_LIMIT_REACHED":  902,
	"UDP_ACTIVE_SESSION_LIMIT_APPROACH":    903,
	"UDP_ACTIVE_SESSION_LIMIT_REACHED":     904,
	"ICMP_ACTIVE_SESSION_LIMIT_APPROACH":   905,
	"ICMP_ACTIVE_SESSION_LIMIT_REACHED":    906,
	"OTHER_ACTIVE_SESSION_LIMIT_APPROACH":  907,
	"OTHER_ACTIVE_SESSION_LIMIT_REACHED":   908,
	"ORCH_CONNECTION_ERROR":                1001,
	"ORCH_LOGIN_FAILURE":                   1002,
	"ORCH_CONFIG_PUSH_FAILURE":             1003,
	"ORCH_INVALID_ACTION":                  1004,
	"MIGRATION_FAILED":                     1005,
	"MIGRATION_TIMED_OUT":                  1006,
	"ORCH_ALREADY_MANAGED":                 1007,
	"ORCH_UNSUPPORTED_VERSION":             1008,
	"ORCH_DSC_NOT_ADMITTED":                1009,
	"ORCH_DSC_MODE_INCOMPATIBLE":           1010,
	"COLLECTOR_REACHABLE":                  1101,
	"COLLECTOR_UNREACHABLE":                1102,
	"BGP_SESSION_ESTABLISHED":              1201,
	"BGP_SESSION_DOWN":                     1202,
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

var E_SuppressMm = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10008,
	Name:          "eventtypes.suppress_mm",
	Tag:           "varint,10008,opt,name=suppress_mm,json=suppressMm",
	Filename:      "eventtypes.proto",
}

func init() {
	proto.RegisterEnum("eventtypes.EventType", EventType_name, EventType_value)
	proto.RegisterExtension(E_Severity)
	proto.RegisterExtension(E_Category)
	proto.RegisterExtension(E_Desc)
	proto.RegisterExtension(E_SuppressMm)
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
	// 2905 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0x49, 0x74, 0xdc, 0xc6,
	0x99, 0xc7, 0xb9, 0xb4, 0x44, 0xa9, 0xe4, 0x11, 0x21, 0x68, 0x35, 0x24, 0xd3, 0x25, 0x69, 0x46,
	0xe3, 0xb1, 0x6c, 0xca, 0x96, 0xc6, 0x9b, 0xc6, 0x1e, 0x0b, 0x04, 0x8a, 0x24, 0x2c, 0x34, 0xd0,
	0x02, 0xd0, 0x94, 0x65, 0x3d, 0x3f, 0x3c, 0x10, 0x28, 0x76, 0x43, 0x42, 0xa3, 0xda, 0x28, 0x80,
	0x1c, 0xce, 0x69, 0x26, 0xce, 0xe2, 0x2c, 0x87, 0xe4, 0x90, 0xe5, 0x9a, 0x17, 0x1f, 0x74, 0xc8,
	0xe2, 0xe4, 0xe8, 0xbc, 0x24, 0x3e, 0xfa, 0x92, 0xbc, 0x5c, 0x73, 0xcb, 0xd3, 0x29, 0x7b, 0xc2,
	0xbc, 0xec, 0xfb, 0xab, 0x42, 0x01, 0x8d, 0x26, 0x29, 0x59, 0xb9, 0x89, 0x7a, 0xfd, 0xff, 0xe1,
	0x5f, 0x5f, 0x55, 0x7d, 0x4b, 0x01, 0x09, 0xaf, 0xe3, 0x34, 0xcf, 0x37, 0x87, 0x98, 0xce, 0x0f,
	0x33, 0x92, 0x13, 0x19, 0x8c, 0xfe, 0x47, 0x01, 0x3d, 0xd2, 0x23, 0xe5, 0xff, 0x2b, 0xb0, 0x47,
	0x48, 0x2f, 0xc1, 0x17, 0xf8, 0x5f, 0xab, 0xc5, 0xda, 0x85, 0x08, 0xd3, 0x30, 0x8b, 0x87, 0x39,
	0xc9, 0xc4, 0x2f, 0xa4, 0x20, 0xcf, 0xb3, 0x78, 0xb5, 0xc8, 0x2b, 0xd6, 0xe3, 0xdf, 0x9c, 0x07,
	0x00, 0x31, 0x9c, 0xc7, 0x70, 0xb2, 0x09, 0x24, 0x64, 0x22, 0xcd, 0x33, 0x6c, 0xcb, 0x77, 0x3d,
	0xd5, 0xf1, 0x90, 0x2e, 0x4d, 0x28, 0xcf, 0xde, 0xd9, 0x6a, 0x4d, 0xbc, 0xb3, 0xd5, 0x9a, 0x78,
	0x77, 0xab, 0x75, 0xce, 0xc4, 0x41, 0x84, 0x33, 0x88, 0x13, 0x1c, 0xe6, 0x31, 0x49, 0x21, 0xcd,
	0x83, 0x2c, 0xc7, 0x11, 0x8c, 0x53, 0x98, 0xf7, 0x31, 0x0c, 0x93, 0x82, 0xe6, 0x38, 0x7b, 0x6f,
	0xab, 0x35, 0x29, 0xab, 0x40, 0xae, 0x69, 0x9a, 0x6a, 0x69, 0xc8, 0x34, 0x91, 0x2e, 0x4d, 0x2a,
	0xff, 0x71, 0x67, 0xab, 0x35, 0x29, 0x78, 0x0f, 0x6f, 0xe7, 0x85, 0x41, 0x1a, 0xe2, 0x24, 0xc1,
	0x11, 0x47, 0xbc, 0x0e, 0x4e, 0xd5, 0x08, 0xcb, 0xf6, 0x8c, 0x45, 0x43, 0x53, 0xf9, 0x1f, 0x8b,
	0xaa, 0xc1, 0x60, 0x53, 0xca, 0x7f, 0x35, 0x60, 0xe7, 0x17, 0x83, 0x38, 0xc1, 0x11, 0xcc, 0x09,
	0xa4, 0x38, 0x8d, 0x60, 0xb2, 0x8d, 0x9d, 0x92, 0x3c, 0x5e, 0x8b, 0xc3, 0x80, 0xfd, 0xc1, 0xf1,
	0x2f, 0x8d, 0xad, 0xd7, 0xee, 0x74, 0x90, 0x2e, 0x4d, 0x2b, 0xff, 0xde, 0x40, 0x1e, 0xdf, 0xb9,
	0x5e, 0x32, 0x1c, 0x0a, 0x77, 0x0b, 0xe0, 0xa0, 0x89, 0x54, 0x1d, 0x39, 0x3e, 0xa7, 0x20, 0x5d,
	0x6a, 0x29, 0xf3, 0x8d, 0x60, 0xcd, 0x35, 0xc5, 0x38, 0x82, 0x6b, 0x24, 0xdb, 0x11, 0x24, 0x03,
	0x1c, 0x10, 0x0c, 0xd3, 0x76, 0x3d, 0x69, 0x8f, 0xf2, 0x7c, 0x03, 0xf0, 0x98, 0x45, 0x22, 0x0c,
	0x13, 0x42, 0x73, 0xb1, 0x16, 0xda, 0x8f, 0x87, 0x30, 0x2a, 0xb2, 0x38, 0xed, 0x71, 0x52, 0x65,
	0x6b, 0x9b, 0x1d, 0x6d, 0x59, 0xb5, 0x96, 0x90, 0x2e, 0xed, 0xdd, 0xd5, 0x4e, 0xd8, 0x0f, 0xd2,
	0xde, 0x68, 0xcb, 0xc6, 0x18, 0xcf, 0x83, 0x03, 0x96, 0xad, 0x23, 0xff, 0x15, 0xdb, 0xb0, 0x90,
	0x2e, 0xcd, 0xf0, 0x60, 0x54, 0x80, 0xe3, 0xdc, 0xce, 0x2d, 0x12, 0xa7, 0x2c, 0xc8, 0xdb, 0x16,
	0xf2, 0x32, 0x38, 0xc8, 0x95, 0xba, 0xe1, 0x0a, 0xf1, 0x3e, 0xe5, 0x7c, 0x23, 0x92, 0x8f, 0x72,
	0x71, 0x14, 0x53, 0xa1, 0x5f, 0xcb, 0xc8, 0xa0, 0x09, 0x91, 0x9f, 0x06, 0x0f, 0x71, 0xc0, 0x32,
	0x52, 0x4d, 0x6f, 0xf9, 0x86, 0xb4, 0x5f, 0x79, 0xb4, 0xf1, 0xed, 0x59, 0x2e, 0x8f, 0x29, 0xec,
	0xe3, 0x20, 0xc9, 0xfb, 0x9b, 0xfc, 0x9b, 0x2f, 0x00, 0x89, 0x4b, 0xba, 0x96, 0x83, 0x54, 0x6d,
	0x59, 0x5d, 0x30, 0x91, 0x04, 0x94, 0xb3, 0x77, 0xb6, 0x5a, 0x53, 0x42, 0x76, 0xb8, 0x92, 0x15,
	0x69, 0x86, 0x83, 0xb0, 0x1f, 0xac, 0x26, 0x98, 0x4b, 0x97, 0xc0, 0x21, 0xcd, 0xb4, 0xb5, 0xab,
	0xbe, 0x7b, 0xc3, 0xd2, 0xaa, 0xe3, 0x74, 0x40, 0x79, 0xaa, 0xe1, 0x18, 0x72, 0xed, 0xda, 0xe8,
	0x4c, 0x6d, 0xa6, 0x61, 0x3f, 0x23, 0x69, 0xfc, 0xbf, 0xcc, 0x34, 0x09, 0x6f, 0x8b, 0x33, 0x74,
	0xe8, 0x5a, 0xd7, 0x76, 0xba, 0x6d, 0xbf, 0x8d, 0xda, 0x0b, 0xc8, 0xf1, 0x55, 0x5d, 0x97, 0x66,
	0x95, 0x73, 0x0d, 0xef, 0xc7, 0xd4, 0x28, 0xc2, 0x11, 0x1c, 0xe0, 0xc1, 0x2a, 0xce, 0x18, 0xe9,
	0x8d, 0x82, 0x64, 0xc5, 0x80, 0xcb, 0x75, 0x70, 0x64, 0x5c, 0xee, 0xa0, 0xb6, 0xbd, 0x82, 0x24,
	0x49, 0x79, 0xbc, 0x41, 0x50, 0x1c, 0x3c, 0x20, 0xeb, 0x23, 0x06, 0x0f, 0x5e, 0x83, 0xb2, 0x04,
	0x8e, 0x8e, 0x53, 0xaa, 0x20, 0x1e, 0x52, 0x9e, 0x68, 0x60, 0x4e, 0x5d, 0xe3, 0x92, 0x8a, 0x12,
	0x53, 0x98, 0x92, 0x8d, 0xb1, 0x88, 0x5e, 0x05, 0xc7, 0xc7, 0x41, 0x5d, 0xab, 0x42, 0xc9, 0xfc,
	0x30, 0x55, 0x81, 0x9d, 0xdb, 0x15, 0x55, 0xa4, 0x4d, 0x98, 0x0b, 0x8e, 0x77, 0x2d, 0xb7, 0xdb,
	0xe9, 0xd8, 0x2c, 0x93, 0xf8, 0x02, 0xec, 0x1a, 0xaf, 0x21, 0xe9, 0x30, 0xcf, 0x2a, 0x15, 0xec,
	0x9c, 0x80, 0x51, 0x16, 0xda, 0x98, 0xc2, 0x55, 0x9c, 0x90, 0x0d, 0x48, 0x8b, 0xe1, 0x90, 0xf0,
	0xe4, 0x32, 0x88, 0xd3, 0x78, 0x20, 0x96, 0x7a, 0x0d, 0x48, 0x02, 0x34, 0xb2, 0x76, 0x84, 0xa7,
	0x81, 0x8a, 0x76, 0x5e, 0xd0, 0x22, 0x82, 0x99, 0xab, 0x1c, 0xf6, 0x83, 0x75, 0x0c, 0x71, 0x4a,
	0x8a, 0x5e, 0xbf, 0x5a, 0xac, 0xf0, 0x4d, 0x39, 0xd2, 0x01, 0xc7, 0xda, 0xb6, 0xde, 0x35, 0x91,
	0xaf, 0x39, 0x68, 0x2c, 0xbf, 0x9c, 0xe0, 0x36, 0xab, 0x03, 0x71, 0xae, 0x4d, 0xa2, 0x22, 0xc1,
	0x30, 0xcc, 0x30, 0xcf, 0x21, 0xfc, 0x42, 0x47, 0x71, 0xd0, 0x4b, 0x09, 0xcd, 0xe3, 0x90, 0x8a,
	0xb3, 0x22, 0xae, 0xc3, 0xac, 0x66, 0x5b, 0x8b, 0xc6, 0x92, 0xef, 0x20, 0xd7, 0xb3, 0x1d, 0xa4,
	0x4b, 0x73, 0xe3, 0x5b, 0xaa, 0x91, 0x74, 0x2d, 0xee, 0x15, 0x59, 0x89, 0xda, 0x08, 0x28, 0xcc,
	0x30, 0xcd, 0x49, 0x26, 0x00, 0xd7, 0xc1, 0xb1, 0x71, 0x80, 0xaf, 0x2e, 0xf0, 0x38, 0x4a, 0x8f,
	0x8e, 0x27, 0xbd, 0x71, 0x8e, 0x60, 0x40, 0x32, 0xc4, 0x0d, 0x72, 0xb0, 0xca, 0x03, 0x29, 0x02,
	0x78, 0x74, 0x1b, 0x58, 0x2c, 0x16, 0x8e, 0xef, 0xc9, 0x07, 0x71, 0x1b, 0x8b, 0x7d, 0x09, 0x3c,
	0xa4, 0xbb, 0x9a, 0xaf, 0xea, 0x6d, 0xc3, 0x63, 0x0e, 0x23, 0x7e, 0xf3, 0xab, 0x95, 0x9e, 0xd4,
	0x5d, 0x0d, 0x06, 0xd1, 0x20, 0xce, 0xf3, 0xf2, 0x22, 0x6d, 0x4f, 0x1d, 0x57, 0x4a, 0xb9, 0x83,
	0x5e, 0x29, 0xb3, 0x28, 0xe6, 0x27, 0xad, 0x5a, 0xe0, 0x99, 0x4a, 0x4e, 0x69, 0x69, 0xe4, 0x8d,
	0x02, 0xd3, 0x5c, 0x04, 0xec, 0x16, 0xcf, 0xac, 0xf2, 0x73, 0x60, 0x96, 0x11, 0x9a, 0x79, 0x60,
	0x4d, 0x39, 0xd3, 0x58, 0x8d, 0xcc, 0x20, 0xbb, 0xa4, 0x81, 0x0b, 0xe0, 0x00, 0x13, 0x56, 0x07,
	0xa9, 0xa7, 0xcc, 0x35, 0x8c, 0x1f, 0x14, 0xa2, 0xe6, 0x99, 0xbe, 0x04, 0xfe, 0xa5, 0xfc, 0x52,
	0x25, 0xe9, 0x2b, 0xb0, 0xf1, 0x1d, 0xa9, 0xfe, 0x4e, 0x53, 0xf4, 0xe6, 0x24, 0x38, 0xb6, 0x6c,
	0xbb, 0x9e, 0xcf, 0xa4, 0x6e, 0x07, 0x69, 0x3e, 0xdb, 0x01, 0xd3, 0xd0, 0x3c, 0x29, 0x56, 0x7a,
	0x8d, 0xb5, 0xde, 0xf4, 0xfa, 0x18, 0xd2, 0x4d, 0x9a, 0xe3, 0x01, 0xec, 0x07, 0x14, 0x46, 0x38,
	0x2f, 0x4b, 0x47, 0x00, 0x43, 0x92, 0xae, 0x25, 0x71, 0x98, 0xc3, 0x55, 0x9c, 0x6f, 0x60, 0x5c,
	0xe6, 0x6e, 0xf6, 0x35, 0x3a, 0xc4, 0x61, 0x5d, 0xd6, 0x28, 0x24, 0x6b, 0x30, 0x8a, 0xd7, 0xd6,
	0x70, 0x86, 0xd3, 0x1c, 0x2e, 0xb3, 0xd2, 0x41, 0x56, 0x59, 0x94, 0xa8, 0x7c, 0x05, 0xb0, 0x05,
	0xf9, 0x3a, 0xaa, 0xf7, 0xe9, 0xd6, 0x58, 0x76, 0xe0, 0xb4, 0x08, 0x3f, 0x59, 0x6f, 0xd5, 0x8e,
	0x14, 0xbd, 0x04, 0xe4, 0x92, 0xa0, 0xd9, 0xed, 0xb6, 0xe1, 0xba, 0x86, 0xcd, 0xf2, 0xfc, 0x6d,
	0xe5, 0x42, 0x83, 0x72, 0xb6, 0xa4, 0x84, 0x64, 0x20, 0x76, 0x6c, 0x37, 0xd0, 0x2d, 0x70, 0x5c,
	0xed, 0x7a, 0xb6, 0xbf, 0x84, 0x2c, 0xe4, 0xa8, 0x2c, 0x39, 0x78, 0xa6, 0xeb, 0x6b, 0xc8, 0xf1,
	0xa4, 0xf7, 0x27, 0x15, 0xb3, 0x11, 0x91, 0x17, 0xd5, 0x22, 0x27, 0xb0, 0x87, 0x53, 0x76, 0xec,
	0x70, 0x04, 0x43, 0x9c, 0x89, 0x22, 0x2e, 0x32, 0x05, 0x2b, 0x85, 0x05, 0x15, 0xd5, 0x55, 0xed,
	0x18, 0x70, 0x29, 0xc8, 0xf1, 0x46, 0xb0, 0x09, 0x3d, 0xd3, 0xe5, 0xc1, 0x7f, 0x0a, 0x3c, 0x64,
	0xda, 0x4b, 0x46, 0x7d, 0xa7, 0xbf, 0x3c, 0xa5, 0x3c, 0xd2, 0xf0, 0x7b, 0xa8, 0x4b, 0x71, 0x06,
	0x13, 0xd2, 0x8b, 0xab, 0x23, 0x2d, 0x2f, 0x80, 0x59, 0xb5, 0xab, 0x1b, 0x9e, 0x61, 0x2d, 0x55,
	0xa2, 0xaf, 0x4c, 0xf1, 0x50, 0x55, 0xdb, 0x7c, 0xea, 0x7a, 0x16, 0xe7, 0xec, 0xcb, 0x64, 0x0d,
	0xaa, 0x45, 0x14, 0xe7, 0xbc, 0x89, 0x6a, 0x5e, 0x89, 0x4b, 0x40, 0xea, 0xa8, 0xae, 0x7b, 0xdd,
	0x76, 0xf4, 0xba, 0x1c, 0x7f, 0x75, 0x4a, 0x39, 0xd5, 0x58, 0x9a, 0xd4, 0x09, 0x28, 0xdd, 0x20,
	0x59, 0x54, 0x55, 0x64, 0xf9, 0x02, 0x38, 0x58, 0x8b, 0x1c, 0xe4, 0x22, 0x4f, 0xfa, 0xda, 0x94,
	0xa2, 0x34, 0x24, 0x07, 0x6b, 0x49, 0x86, 0x29, 0xce, 0xe5, 0x0b, 0x60, 0xc6, 0x34, 0xac, 0xab,
	0x7e, 0xb7, 0x23, 0x7d, 0x7a, 0x5a, 0x39, 0x2d, 0x96, 0x35, 0xc9, 0x96, 0xd5, 0x21, 0x59, 0xce,
	0x22, 0x94, 0xc4, 0xe9, 0x6d, 0x1c, 0xc1, 0x62, 0x28, 0xea, 0xfb, 0x7e, 0x2e, 0xd0, 0xed, 0xeb,
	0x96, 0xf4, 0x99, 0x69, 0xe5, 0x31, 0x01, 0x67, 0x92, 0x13, 0x5c, 0xc2, 0x7e, 0xcf, 0xba, 0xba,
	0xbc, 0xa0, 0x4c, 0x1e, 0x91, 0x8d, 0xb2, 0x33, 0x78, 0x06, 0xcc, 0xba, 0xc8, 0x59, 0x31, 0x34,
	0x54, 0xb7, 0x86, 0xbf, 0x99, 0xe6, 0x25, 0x7a, 0xfa, 0x9d, 0xad, 0xd6, 0x14, 0x2b, 0xd1, 0x2e,
	0xce, 0xd6, 0xe3, 0x10, 0x57, 0x3d, 0xe1, 0x4e, 0x59, 0xd9, 0x61, 0xfd, 0xb6, 0x94, 0x4d, 0xee,
	0x94, 0x8d, 0x5a, 0x2b, 0x0d, 0x1c, 0xb3, 0xd4, 0x8e, 0x89, 0x5c, 0x7f, 0xbb, 0xfa, 0x77, 0xd3,
	0xbc, 0xb6, 0x4e, 0x09, 0xf5, 0x31, 0x2b, 0x18, 0x26, 0x98, 0x42, 0xba, 0x0b, 0xe4, 0xe9, 0xd1,
	0xb7, 0x3b, 0xc8, 0xd2, 0x0d, 0x6b, 0x49, 0xfa, 0xfd, 0xb4, 0x72, 0x72, 0x37, 0xcb, 0x43, 0x9c,
	0x46, 0x71, 0xda, 0x6b, 0xda, 0x75, 0xba, 0x96, 0xc5, 0x24, 0x7f, 0xb8, 0xc7, 0x2a, 0xb3, 0x22,
	0x4d, 0xe3, 0xb4, 0xc7, 0xbf, 0x74, 0x13, 0x1c, 0xa9, 0x64, 0x2c, 0x07, 0xb9, 0x1d, 0xdb, 0x72,
	0x8d, 0x15, 0x24, 0xfd, 0x71, 0x5a, 0xb9, 0xd2, 0x30, 0xfb, 0x9f, 0x95, 0x96, 0xa5, 0x21, 0x3a,
	0x24, 0x29, 0x8d, 0xd7, 0x31, 0x8c, 0x0a, 0xcc, 0xb2, 0x62, 0x12, 0x84, 0xb7, 0xd9, 0x79, 0x12,
	0xd7, 0x3f, 0xc3, 0x94, 0x14, 0x59, 0x88, 0xa9, 0xfc, 0x3c, 0x98, 0x75, 0x6f, 0xb8, 0x1e, 0x6a,
	0xfb, 0x9a, 0x6d, 0xea, 0x0b, 0xb6, 0xed, 0x49, 0x7f, 0x9a, 0xe6, 0xd9, 0xad, 0x0a, 0xa1, 0xec,
	0x96, 0x9a, 0x90, 0x24, 0x11, 0x5c, 0x25, 0xa4, 0x0a, 0x3e, 0x02, 0x47, 0x85, 0xd2, 0x41, 0xae,
	0xdd, 0x75, 0x98, 0x3d, 0x57, 0x5d, 0x42, 0xd2, 0x9f, 0xa7, 0xeb, 0x2e, 0x9c, 0xe9, 0x1f, 0x71,
	0xc7, 0xbf, 0x09, 0x0b, 0x1a, 0xf4, 0xca, 0x6e, 0x2b, 0xee, 0xf5, 0x65, 0x54, 0x6f, 0xc6, 0xa2,
	0xea, 0xa9, 0xa6, 0x6f, 0x58, 0x1e, 0x72, 0x9c, 0x6e, 0xc7, 0x93, 0xfe, 0x52, 0x9e, 0xa0, 0x6a,
	0x7d, 0xa7, 0xc4, 0x66, 0xb0, 0xd4, 0x15, 0xc0, 0xb5, 0x20, 0x0f, 0x12, 0x18, 0xa7, 0x39, 0xce,
	0xb2, 0x62, 0x98, 0xcb, 0x2b, 0xe0, 0x84, 0xc0, 0x68, 0xaa, 0xe7, 0x39, 0x46, 0xa7, 0x01, 0xfa,
	0xeb, 0xb4, 0xf2, 0x5c, 0x03, 0x74, 0x5e, 0x18, 0xc2, 0x69, 0x48, 0x0a, 0x06, 0x60, 0x37, 0x9f,
	0x4f, 0x2f, 0xc3, 0xf2, 0xe4, 0xf3, 0xcb, 0x57, 0x46, 0x4a, 0x36, 0x80, 0x24, 0xb8, 0xf6, 0x0a,
	0x72, 0x7c, 0x0f, 0xb5, 0x3b, 0xd2, 0xdf, 0xa6, 0x95, 0x8b, 0x0d, 0xde, 0x39, 0xc1, 0xcb, 0xf1,
	0x80, 0xd7, 0xaf, 0x22, 0xe3, 0xab, 0x0b, 0x56, 0xc9, 0x3a, 0x86, 0x79, 0x3f, 0xc3, 0xb4, 0x4f,
	0x92, 0x68, 0x5e, 0xb6, 0xc0, 0xd1, 0xed, 0x28, 0x1f, 0xbd, 0x6a, 0x78, 0xd2, 0xdf, 0x4b, 0xde,
	0xc4, 0xfd, 0x79, 0x65, 0xdf, 0xd2, 0xe0, 0x2d, 0x02, 0x59, 0xf0, 0x3a, 0xaa, 0x65, 0x68, 0x3e,
	0x5a, 0x41, 0x96, 0x27, 0xfd, 0x5f, 0x4b, 0x79, 0xb2, 0x61, 0xee, 0xb4, 0x80, 0x0d, 0x83, 0x34,
	0x0e, 0x21, 0x29, 0x33, 0xfb, 0x30, 0xc3, 0xeb, 0x31, 0x29, 0x28, 0xdf, 0x4e, 0xb9, 0x0d, 0xe6,
	0x2a, 0x0e, 0x2f, 0x23, 0x86, 0x5a, 0x66, 0xa6, 0xae, 0x83, 0x04, 0xf3, 0xff, 0x5b, 0xf5, 0x5d,
	0xe6, 0x3b, 0x51, 0x31, 0x59, 0x01, 0x60, 0x2d, 0x0a, 0xcc, 0x59, 0xd1, 0x14, 0x09, 0xee, 0x26,
	0x38, 0x25, 0x70, 0x86, 0xb5, 0x68, 0xfb, 0x1d, 0xcd, 0x40, 0x65, 0x39, 0x13, 0xb0, 0x0f, 0xb5,
	0xea, 0x31, 0x84, 0xc1, 0x9e, 0x70, 0x77, 0xad, 0x48, 0xc3, 0x30, 0xc6, 0x65, 0xba, 0x28, 0x6b,
	0x1d, 0xe4, 0x83, 0x29, 0x9b, 0xd9, 0x04, 0xfc, 0xba, 0xea, 0x58, 0x3b, 0xe1, 0x6f, 0xb6, 0x94,
	0xcb, 0x0d, 0xa7, 0xf3, 0x0f, 0x08, 0xdf, 0x08, 0x32, 0x76, 0xdd, 0xe4, 0xd7, 0xc0, 0x49, 0x81,
	0x47, 0x8e, 0xb3, 0x93, 0xfe, 0xe1, 0xd2, 0xfa, 0xd4, 0x3f, 0x6b, 0x3d, 0xcb, 0x48, 0x26, 0xab,
	0xe0, 0xb8, 0x6e, 0xb8, 0x57, 0x7d, 0x6f, 0xd9, 0x41, 0xee, 0xb2, 0x6d, 0xea, 0x3e, 0x7a, 0x55,
	0x43, 0x48, 0x47, 0xba, 0xf4, 0x83, 0x56, 0x3d, 0x57, 0xec, 0x61, 0xa3, 0x90, 0x1e, 0xd3, 0xdb,
	0xa3, 0x9d, 0x86, 0xf8, 0x7f, 0x42, 0x8c, 0x23, 0x1c, 0xc9, 0x97, 0xc0, 0xac, 0x63, 0x9b, 0xa6,
	0xdd, 0xf5, 0xea, 0x34, 0xf9, 0xad, 0x3d, 0x75, 0xc1, 0x99, 0x66, 0x99, 0xd9, 0x21, 0x49, 0x42,
	0x8a, 0x1c, 0xc6, 0x69, 0x9c, 0xc7, 0xac, 0xa8, 0xc9, 0x57, 0x1a, 0xa2, 0xae, 0xa6, 0x21, 0xd7,
	0x95, 0xbe, 0xbd, 0xa7, 0xee, 0x16, 0x99, 0x68, 0xae, 0x12, 0xd1, 0x22, 0x0c, 0x31, 0xa5, 0x6b,
	0x45, 0x92, 0x6c, 0xc2, 0x90, 0x0c, 0x86, 0x09, 0xce, 0xcb, 0xca, 0x51, 0x11, 0x44, 0xc5, 0xfa,
	0xce, 0x1e, 0x5e, 0x39, 0xa6, 0x04, 0xe0, 0x60, 0x05, 0x10, 0x47, 0xe0, 0x19, 0x70, 0x68, 0xf4,
	0x49, 0x97, 0xa5, 0x47, 0xa4, 0x4b, 0xef, 0xdd, 0xc3, 0x29, 0x2d, 0x28, 0x4b, 0x90, 0x38, 0x92,
	0x2f, 0x82, 0x03, 0xa2, 0x79, 0x64, 0x9f, 0x91, 0xbe, 0xb8, 0x97, 0x77, 0x3f, 0x6c, 0x2f, 0x5b,
	0xef, 0x6e, 0xb5, 0x8e, 0x8c, 0xf7, 0x8c, 0xe2, 0x53, 0x1e, 0xf8, 0x57, 0x4f, 0xeb, 0xf8, 0xcb,
	0xaa, 0xb9, 0xe8, 0xdb, 0x1d, 0x64, 0xf9, 0x2e, 0xe2, 0xad, 0x83, 0x6f, 0x1a, 0x6d, 0xc3, 0xf3,
	0xd5, 0x4e, 0xc7, 0xb1, 0x55, 0x6d, 0x59, 0xfa, 0xc8, 0xcc, 0xd8, 0xd3, 0xc0, 0x23, 0x9e, 0xd6,
	0x81, 0x2e, 0x2e, 0xbb, 0x3e, 0x33, 0x1e, 0xc4, 0x39, 0x0c, 0x86, 0xc3, 0x8c, 0x04, 0x61, 0x9f,
	0x9d, 0x03, 0x1b, 0x9c, 0xbd, 0x1f, 0x95, 0x37, 0x82, 0x48, 0x97, 0x3e, 0x3a, 0xa3, 0xfc, 0x5b,
	0xa3, 0x70, 0x3f, 0xbc, 0x13, 0xca, 0x1b, 0x42, 0x1c, 0xc9, 0xd7, 0xc0, 0xe9, 0xae, 0xde, 0xf1,
	0x55, 0xcd, 0x33, 0x56, 0xd0, 0xbd, 0x3c, 0x7e, 0x6c, 0x9b, 0xc7, 0xae, 0x7e, 0x3f, 0x8f, 0x26,
	0x80, 0xf7, 0x44, 0x56, 0x06, 0xdf, 0xda, 0x66, 0x70, 0x27, 0xb1, 0x32, 0xe8, 0x82, 0x33, 0x86,
	0xd6, 0xfe, 0x20, 0x87, 0x1f, 0x9f, 0xe1, 0x07, 0xa7, 0x72, 0x38, 0xc7, 0x24, 0xf7, 0xb1, 0x68,
	0x81, 0xd3, 0xf7, 0x86, 0x56, 0x1e, 0x3f, 0x31, 0x53, 0xd7, 0x5c, 0x3e, 0xba, 0xec, 0xc2, 0xac,
	0x4c, 0x76, 0xc1, 0x59, 0xdb, 0x5b, 0x66, 0x63, 0xf0, 0x7d, 0x5d, 0x7e, 0x72, 0x66, 0xfc, 0x71,
	0x80, 0x6b, 0xee, 0x63, 0xb3, 0x03, 0xce, 0xdc, 0x07, 0x5b, 0xf9, 0xfc, 0x54, 0xf9, 0x5e, 0x51,
	0xf9, 0x3c, 0xb9, 0x1b, 0xb5, 0x32, 0x6a, 0x80, 0xa3, 0xb6, 0xa3, 0x2d, 0xb3, 0x4e, 0xdc, 0x12,
	0x4f, 0x40, 0xc8, 0x71, 0x6c, 0x47, 0xfa, 0xf1, 0x4c, 0x9d, 0x9d, 0xf7, 0xb2, 0xec, 0x3c, 0x7a,
	0x54, 0x0a, 0x49, 0x9a, 0xe2, 0x30, 0x67, 0xff, 0x24, 0x59, 0xd8, 0xc7, 0x34, 0xcf, 0x82, 0x9c,
	0x64, 0x72, 0x17, 0xc8, 0x1c, 0x35, 0x6a, 0x33, 0xbb, 0x0e, 0x92, 0x7e, 0x32, 0xa3, 0xbc, 0xd8,
	0xe0, 0x3c, 0x65, 0xf2, 0x16, 0x33, 0xcc, 0x70, 0x84, 0xd3, 0x3c, 0x0e, 0x12, 0x0a, 0x37, 0x30,
	0x2b, 0x1b, 0xe9, 0x7a, 0x90, 0xc4, 0x65, 0xfb, 0xda, 0xa4, 0xce, 0xcb, 0xaf, 0x83, 0x13, 0x95,
	0x43, 0x76, 0xe1, 0x3a, 0x5d, 0x77, 0xb9, 0x86, 0xff, 0x74, 0x46, 0xf9, 0x6f, 0x11, 0x3f, 0x06,
	0xbf, 0x38, 0x32, 0x39, 0x2c, 0x68, 0x1f, 0x52, 0x32, 0xc0, 0x7c, 0x62, 0xa8, 0x2f, 0x23, 0xdd,
	0xe1, 0xda, 0x03, 0x87, 0x39, 0xde, 0xb0, 0x56, 0x54, 0xd3, 0xd0, 0x79, 0x64, 0x6d, 0x4b, 0xfa,
	0xd9, 0x4c, 0x9d, 0x9e, 0x19, 0x79, 0x9e, 0xb7, 0xc7, 0x43, 0x9c, 0xad, 0x91, 0x6c, 0xc0, 0x72,
	0x67, 0x0a, 0x83, 0xf2, 0x25, 0x2c, 0xef, 0x07, 0x79, 0x39, 0xf7, 0xe7, 0xa3, 0x51, 0x9d, 0xf5,
	0xbd, 0x6d, 0x63, 0xc9, 0x19, 0x9b, 0xa2, 0x7f, 0x3e, 0x53, 0xf7, 0xbd, 0x0c, 0x29, 0xb5, 0xe3,
	0x9e, 0xc8, 0x10, 0xa5, 0x6d, 0xf9, 0x05, 0x70, 0x78, 0x24, 0xf2, 0x8c, 0x36, 0xd2, 0x7d, 0xbb,
	0xeb, 0x49, 0xbf, 0x98, 0xa9, 0x1b, 0x45, 0xa6, 0x3b, 0x3c, 0xd2, 0x79, 0x31, 0xf3, 0x42, 0x8a,
	0x5c, 0xee, 0x82, 0x23, 0x7c, 0x15, 0xaa, 0xe9, 0x20, 0x55, 0xbf, 0xe1, 0xb7, 0x55, 0x4b, 0x65,
	0xbd, 0xf6, 0x2f, 0xb7, 0x2d, 0x43, 0x4d, 0x49, 0xde, 0xc7, 0x19, 0xec, 0xb8, 0x6d, 0x38, 0x08,
	0x36, 0xe1, 0x2a, 0x86, 0x83, 0x20, 0x0d, 0x7a, 0xd5, 0x4b, 0x1a, 0x0d, 0x06, 0x18, 0xa6, 0xc1,
	0x00, 0xd3, 0x61, 0x10, 0x62, 0xf9, 0xaa, 0x88, 0x7d, 0xf3, 0xfd, 0x62, 0x05, 0x39, 0xec, 0xcc,
	0x49, 0xbf, 0x9a, 0xa9, 0x67, 0x01, 0x86, 0x86, 0xdd, 0x74, 0xf4, 0x52, 0xd1, 0x8c, 0x30, 0x5c,
	0xc7, 0x19, 0x3b, 0x78, 0xf2, 0x4d, 0x71, 0xd4, 0xd8, 0xec, 0x64, 0xd9, 0xde, 0x68, 0x4e, 0xde,
	0x9a, 0x51, 0x5e, 0x6e, 0xdc, 0x82, 0x4b, 0xd7, 0x49, 0x76, 0x3b, 0x21, 0x41, 0x04, 0x87, 0x19,
	0x59, 0x8f, 0xc5, 0xf0, 0x44, 0x52, 0x18, 0xc0, 0x3e, 0xab, 0xe3, 0x1b, 0x71, 0xde, 0x87, 0x29,
	0x19, 0xcd, 0xd2, 0xba, 0xab, 0xc9, 0x8b, 0x40, 0xa9, 0xe1, 0x6d, 0x5b, 0x47, 0xbe, 0x61, 0x69,
	0x76, 0xbb, 0xa3, 0x7a, 0x06, 0x9b, 0x82, 0x7f, 0xbd, 0x2d, 0xbb, 0xb0, 0xe1, 0x6c, 0x20, 0x9e,
	0xc4, 0xe2, 0x94, 0xd5, 0x8f, 0x20, 0x8f, 0x57, 0x13, 0xcc, 0x5a, 0x95, 0xc3, 0x9a, 0x6d, 0x9a,
	0x48, 0xf3, 0x6c, 0xc7, 0x1f, 0x8d, 0xd1, 0xdf, 0xdd, 0x57, 0xcf, 0x88, 0x6c, 0x46, 0x38, 0xa5,
	0x91, 0x24, 0xc1, 0x21, 0x5b, 0x5d, 0x3d, 0x48, 0x97, 0xe3, 0x9d, 0xee, 0x6a, 0xbc, 0xe7, 0x5c,
	0x06, 0x47, 0x47, 0x9c, 0xe6, 0x40, 0xfe, 0xbd, 0x92, 0x54, 0x4d, 0x1b, 0x70, 0x44, 0x62, 0x27,
	0x68, 0x27, 0x4d, 0x5e, 0x06, 0xc7, 0x17, 0x96, 0x3a, 0xf5, 0x55, 0x47, 0xae, 0xa7, 0x2e, 0x98,
	0x86, 0xcb, 0x2e, 0xfa, 0xd7, 0xf7, 0xd7, 0xd5, 0x91, 0xb1, 0x94, 0x85, 0xa5, 0x0e, 0xa4, 0xe2,
	0x9a, 0xc7, 0x14, 0x62, 0x9a, 0x07, 0xab, 0x49, 0x4c, 0xfb, 0xa2, 0x0f, 0xbe, 0x0c, 0xa4, 0x26,
	0x89, 0x0f, 0x3f, 0xdf, 0xd8, 0xcf, 0x0b, 0x7a, 0x65, 0xe7, 0xf0, 0x36, 0x44, 0x35, 0xf7, 0x28,
	0x0f, 0xbf, 0xf5, 0xa5, 0xb9, 0x89, 0x3b, 0x6f, 0xcf, 0x4d, 0xbc, 0xf3, 0xf6, 0xdc, 0xe4, 0xbb,
	0x6f, 0xcf, 0xed, 0xaf, 0xdf, 0xcb, 0x2f, 0x7b, 0x60, 0x1f, 0xc5, 0xeb, 0x38, 0x8b, 0xf3, 0x4d,
	0xf9, 0xf4, 0x7c, 0xf9, 0xfc, 0x3e, 0x5f, 0x3d, 0xbf, 0xcf, 0xa3, 0xb4, 0x18, 0xac, 0x04, 0x49,
	0x81, 0xed, 0x21, 0xbf, 0x80, 0x27, 0x3e, 0x6b, 0xc1, 0xc9, 0xc7, 0x0e, 0x5e, 0x3c, 0x32, 0xcf,
	0xfb, 0x24, 0xd6, 0xcd, 0xd2, 0x79, 0x57, 0xe8, 0x9d, 0x9a, 0xc4, 0xa8, 0x6c, 0xba, 0xed, 0x91,
	0xec, 0x81, 0xa8, 0x9f, 0xdb, 0x85, 0xaa, 0x09, 0xbd, 0x53, 0x93, 0x2e, 0x3f, 0x0b, 0x5a, 0x11,
	0xa6, 0xe1, 0x83, 0x10, 0x3f, 0xcf, 0x88, 0xfb, 0x1d, 0xfe, 0xfb, 0xcb, 0x1a, 0x38, 0xc0, 0x4e,
	0x77, 0x86, 0x29, 0xf5, 0x07, 0x83, 0x07, 0x91, 0x7f, 0x81, 0xc9, 0xf7, 0x39, 0xa0, 0x92, 0xb5,
	0x07, 0x0b, 0xd2, 0xfb, 0x77, 0xe7, 0x26, 0xbf, 0x7f, 0x77, 0x6e, 0xf2, 0x87, 0x77, 0xe7, 0x26,
	0x7f, 0x74, 0x77, 0x6e, 0xa2, 0x33, 0xb1, 0xba, 0x97, 0x13, 0x2e, 0xfd, 0x23, 0x00, 0x00, 0xff,
	0xff, 0x33, 0xd5, 0x18, 0x4c, 0xe1, 0x18, 0x00, 0x00,
}
