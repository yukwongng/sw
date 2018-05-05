// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metric.proto

/*
Package metric is a generated protocol buffer package.

It is generated from these files:
	metric.proto

It has these top-level messages:
	LineBundle
	Field
	MetricPoint
	MetricBundle
*/
package metric

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import api "github.com/pensando/sw/api"
import api1 "github.com/pensando/sw/api"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LineBundle struct {
	DbName string   `protobuf:"bytes,1,opt,name=dbName" json:"dbName,omitempty"`
	SrcID  uint64   `protobuf:"varint,2,opt,name=srcID" json:"srcID,omitempty"`
	Lines  []string `protobuf:"bytes,3,rep,name=Lines" json:"Lines,omitempty"`
}

func (m *LineBundle) Reset()                    { *m = LineBundle{} }
func (m *LineBundle) String() string            { return proto.CompactTextString(m) }
func (*LineBundle) ProtoMessage()               {}
func (*LineBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LineBundle) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *LineBundle) GetSrcID() uint64 {
	if m != nil {
		return m.SrcID
	}
	return 0
}

func (m *LineBundle) GetLines() []string {
	if m != nil {
		return m.Lines
	}
	return nil
}

// Field is one of (uint64, float64, string, bool)
type Field struct {
	// Types that are valid to be assigned to F:
	//	*Field_Int64
	//	*Field_Float64
	//	*Field_String_
	//	*Field_Bool
	F isField_F `protobuf_oneof:"f"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isField_F interface{ isField_F() }

type Field_Int64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=Int64,oneof"`
}
type Field_Float64 struct {
	Float64 float64 `protobuf:"fixed64,2,opt,name=Float64,oneof"`
}
type Field_String_ struct {
	String_ string `protobuf:"bytes,3,opt,name=String,oneof"`
}
type Field_Bool struct {
	Bool bool `protobuf:"varint,4,opt,name=Bool,oneof"`
}

func (*Field_Int64) isField_F()   {}
func (*Field_Float64) isField_F() {}
func (*Field_String_) isField_F() {}
func (*Field_Bool) isField_F()    {}

func (m *Field) GetF() isField_F {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *Field) GetInt64() int64 {
	if x, ok := m.GetF().(*Field_Int64); ok {
		return x.Int64
	}
	return 0
}

func (m *Field) GetFloat64() float64 {
	if x, ok := m.GetF().(*Field_Float64); ok {
		return x.Float64
	}
	return 0
}

func (m *Field) GetString_() string {
	if x, ok := m.GetF().(*Field_String_); ok {
		return x.String_
	}
	return ""
}

func (m *Field) GetBool() bool {
	if x, ok := m.GetF().(*Field_Bool); ok {
		return x.Bool
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Field) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Field_OneofMarshaler, _Field_OneofUnmarshaler, _Field_OneofSizer, []interface{}{
		(*Field_Int64)(nil),
		(*Field_Float64)(nil),
		(*Field_String_)(nil),
		(*Field_Bool)(nil),
	}
}

func _Field_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Field)
	// f
	switch x := m.F.(type) {
	case *Field_Int64:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64))
	case *Field_Float64:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Float64))
	case *Field_String_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.String_)
	case *Field_Bool:
		t := uint64(0)
		if x.Bool {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Field.F has unexpected type %T", x)
	}
	return nil
}

func _Field_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Field)
	switch tag {
	case 1: // f.Int64
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.F = &Field_Int64{int64(x)}
		return true, err
	case 2: // f.Float64
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.F = &Field_Float64{math.Float64frombits(x)}
		return true, err
	case 3: // f.String
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.F = &Field_String_{x}
		return true, err
	case 4: // f.Bool
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.F = &Field_Bool{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Field_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Field)
	// f
	switch x := m.F.(type) {
	case *Field_Int64:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int64))
	case *Field_Float64:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *Field_String_:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.String_)))
		n += len(x.String_)
	case *Field_Bool:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// MetricPoint contains a set of tags and fields, and a timestamp
type MetricPoint struct {
	Name   string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags   map[string]string `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Fields map[string]*Field `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	When   *api1.Timestamp   `protobuf:"bytes,4,opt,name=when" json:"when,omitempty"`
}

func (m *MetricPoint) Reset()                    { *m = MetricPoint{} }
func (m *MetricPoint) String() string            { return proto.CompactTextString(m) }
func (*MetricPoint) ProtoMessage()               {}
func (*MetricPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetricPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricPoint) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MetricPoint) GetFields() map[string]*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *MetricPoint) GetWhen() *api1.Timestamp {
	if m != nil {
		return m.When
	}
	return nil
}

// MetricBundle is a set of metric points to be written to the same db.
type MetricBundle struct {
	DbName   string         `protobuf:"bytes,1,opt,name=dbName" json:"dbName,omitempty"`
	Reporter string         `protobuf:"bytes,2,opt,name=reporter" json:"reporter,omitempty"`
	Metrics  []*MetricPoint `protobuf:"bytes,3,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *MetricBundle) Reset()                    { *m = MetricBundle{} }
func (m *MetricBundle) String() string            { return proto.CompactTextString(m) }
func (*MetricBundle) ProtoMessage()               {}
func (*MetricBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetricBundle) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *MetricBundle) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *MetricBundle) GetMetrics() []*MetricPoint {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*LineBundle)(nil), "metric.LineBundle")
	proto.RegisterType((*Field)(nil), "metric.Field")
	proto.RegisterType((*MetricPoint)(nil), "metric.MetricPoint")
	proto.RegisterType((*MetricBundle)(nil), "metric.MetricBundle")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MetricApi service

type MetricApiClient interface {
	WriteMetrics(ctx context.Context, in *MetricBundle, opts ...grpc.CallOption) (*api.Empty, error)
	WriteLines(ctx context.Context, in *LineBundle, opts ...grpc.CallOption) (*api.Empty, error)
}

type metricApiClient struct {
	cc *grpc.ClientConn
}

func NewMetricApiClient(cc *grpc.ClientConn) MetricApiClient {
	return &metricApiClient{cc}
}

func (c *metricApiClient) WriteMetrics(ctx context.Context, in *MetricBundle, opts ...grpc.CallOption) (*api.Empty, error) {
	out := new(api.Empty)
	err := grpc.Invoke(ctx, "/metric.MetricApi/WriteMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricApiClient) WriteLines(ctx context.Context, in *LineBundle, opts ...grpc.CallOption) (*api.Empty, error) {
	out := new(api.Empty)
	err := grpc.Invoke(ctx, "/metric.MetricApi/WriteLines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetricApi service

type MetricApiServer interface {
	WriteMetrics(context.Context, *MetricBundle) (*api.Empty, error)
	WriteLines(context.Context, *LineBundle) (*api.Empty, error)
}

func RegisterMetricApiServer(s *grpc.Server, srv MetricApiServer) {
	s.RegisterService(&_MetricApi_serviceDesc, srv)
}

func _MetricApi_WriteMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricApiServer).WriteMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metric.MetricApi/WriteMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricApiServer).WriteMetrics(ctx, req.(*MetricBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricApi_WriteLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricApiServer).WriteLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metric.MetricApi/WriteLines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricApiServer).WriteLines(ctx, req.(*LineBundle))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metric.MetricApi",
	HandlerType: (*MetricApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteMetrics",
			Handler:    _MetricApi_WriteMetrics_Handler,
		},
		{
			MethodName: "WriteLines",
			Handler:    _MetricApi_WriteLines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metric.proto",
}

func init() { proto.RegisterFile("metric.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x9b, 0x26, 0xeb, 0xd6, 0x4b, 0x41, 0xc8, 0x54, 0x53, 0x14, 0x81, 0xa8, 0x82, 0x10,
	0x7d, 0x59, 0x22, 0x02, 0x62, 0x88, 0x37, 0x2a, 0x3a, 0x75, 0x12, 0xa0, 0xca, 0x4c, 0xe2, 0xd9,
	0x4d, 0xdd, 0xcc, 0x22, 0xb1, 0x43, 0xec, 0x6c, 0xea, 0xf7, 0xe2, 0x03, 0x22, 0xdb, 0xc9, 0x96,
	0x49, 0x05, 0xde, 0xfc, 0x3f, 0xff, 0x7f, 0xe7, 0xbb, 0xf3, 0xc1, 0xa4, 0xa4, 0xaa, 0x66, 0x59,
	0x5c, 0xd5, 0x42, 0x09, 0x34, 0xb2, 0x2a, 0x7c, 0x96, 0x0b, 0x91, 0x17, 0x34, 0x21, 0x15, 0x4b,
	0x08, 0xe7, 0x42, 0x11, 0xc5, 0x04, 0x97, 0xd6, 0x15, 0x2e, 0x73, 0xa6, 0xae, 0x9b, 0x4d, 0x9c,
	0x89, 0x32, 0xa9, 0x28, 0x97, 0x84, 0x6f, 0x45, 0x22, 0x6f, 0x93, 0x1b, 0xca, 0x59, 0x46, 0x93,
	0x46, 0xb1, 0x42, 0x6a, 0x34, 0xa7, 0xbc, 0x4f, 0x27, 0x8c, 0x67, 0x45, 0xb3, 0xa5, 0x5d, 0x9a,
	0xd7, 0x7f, 0x49, 0xa3, 0x1f, 0xcd, 0x44, 0x59, 0x0a, 0xde, 0x1a, 0x5f, 0xfd, 0xc3, 0x58, 0x52,
	0x45, 0x5a, 0xdb, 0x59, 0xcf, 0x96, 0x8b, 0x5c, 0x24, 0x26, 0xbc, 0x69, 0x76, 0x46, 0x19, 0x61,
	0x4e, 0xd6, 0x1e, 0xad, 0x01, 0xbe, 0x30, 0x4e, 0x17, 0x0d, 0xdf, 0x16, 0x14, 0x9d, 0xc2, 0x68,
	0xbb, 0xf9, 0x46, 0x4a, 0x1a, 0x38, 0x33, 0x67, 0x3e, 0xc6, 0xad, 0x42, 0x53, 0x38, 0x92, 0x75,
	0x76, 0xf9, 0x39, 0x18, 0xce, 0x9c, 0xb9, 0x87, 0xad, 0xd0, 0x51, 0xcd, 0xca, 0xc0, 0x9d, 0xb9,
	0xf3, 0x31, 0xb6, 0x22, 0xaa, 0xe0, 0xe8, 0x82, 0xd1, 0x62, 0x8b, 0x4e, 0xe1, 0xe8, 0x92, 0xab,
	0xf7, 0xef, 0x4c, 0x2e, 0x77, 0x35, 0xc0, 0x56, 0xa2, 0x10, 0x8e, 0x2f, 0x0a, 0x41, 0xf4, 0x8d,
	0x4e, 0xe7, 0xac, 0x06, 0xb8, 0x0b, 0xa0, 0x00, 0x46, 0xdf, 0x55, 0xcd, 0x78, 0x1e, 0xb8, 0xba,
	0x80, 0xd5, 0x00, 0xb7, 0x1a, 0x4d, 0xc1, 0x5b, 0x08, 0x51, 0x04, 0xde, 0xcc, 0x99, 0x9f, 0xac,
	0x06, 0xd8, 0xa8, 0x85, 0x0b, 0xce, 0x2e, 0xfa, 0x3d, 0x04, 0xff, 0xab, 0xf9, 0xb2, 0xb5, 0x60,
	0x5c, 0x21, 0x04, 0x1e, 0xbf, 0xef, 0xc1, 0x9c, 0xd1, 0x1b, 0xf0, 0x14, 0xc9, 0x65, 0x30, 0x9c,
	0xb9, 0x73, 0x3f, 0x7d, 0x1e, 0xb7, 0x1f, 0xde, 0xc3, 0xe2, 0x2b, 0x92, 0xcb, 0x25, 0x57, 0xf5,
	0x1e, 0x1b, 0x2b, 0x3a, 0x87, 0xd1, 0x4e, 0x37, 0x62, 0xfb, 0xf3, 0xd3, 0x17, 0x87, 0x20, 0xd3,
	0x6a, 0x8b, 0xb5, 0x76, 0x14, 0x81, 0x77, 0x7b, 0x4d, 0xb9, 0x29, 0xd5, 0x4f, 0x1f, 0xc7, 0xa4,
	0x62, 0xf1, 0x15, 0x2b, 0xa9, 0x54, 0xa4, 0xac, 0xb0, 0xb9, 0x0b, 0xcf, 0x61, 0x7c, 0xf7, 0x1e,
	0x7a, 0x02, 0xee, 0x4f, 0xba, 0x6f, 0xeb, 0xd5, 0x47, 0x3d, 0xda, 0x1b, 0x52, 0x34, 0xd4, 0x4c,
	0x68, 0x8c, 0xad, 0xf8, 0x38, 0xfc, 0xe0, 0x84, 0x2b, 0xf0, 0x7b, 0x6f, 0x1e, 0x40, 0x5f, 0xf6,
	0x51, 0x3f, 0x7d, 0xd4, 0x55, 0x6d, 0xa8, 0x5e, 0xa6, 0xe8, 0x17, 0x4c, 0x6c, 0x27, 0xff, 0xf9,
	0xfc, 0x10, 0x4e, 0x6a, 0x5a, 0x89, 0x5a, 0xd1, 0xba, 0x2d, 0xe7, 0x4e, 0xa3, 0x33, 0x38, 0xb6,
	0xe9, 0xbb, 0x21, 0x3d, 0x3d, 0x30, 0x24, 0xdc, 0x79, 0x52, 0x01, 0x63, 0x1b, 0xff, 0x54, 0x31,
	0x94, 0xc2, 0xe4, 0x47, 0xcd, 0x14, 0xb5, 0x11, 0x89, 0xa6, 0x0f, 0x51, 0x5b, 0x55, 0x08, 0x66,
	0x7c, 0xcb, 0xb2, 0x52, 0xfb, 0x68, 0x80, 0x62, 0x00, 0xc3, 0x98, 0x55, 0x43, 0xa8, 0x23, 0xee,
	0x57, 0xf8, 0xa1, 0x7f, 0xed, 0x6c, 0x46, 0x66, 0xcf, 0xdf, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xef, 0x0a, 0x6c, 0xb3, 0xe3, 0x03, 0x00, 0x00,
}
