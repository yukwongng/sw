// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: namespace.proto

package netproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NamespaceSpec_Type int32

const (
	NamespaceSpec_CUSTOMER NamespaceSpec_Type = 0
	NamespaceSpec_INFRA    NamespaceSpec_Type = 1
)

var NamespaceSpec_Type_name = map[int32]string{
	0: "CUSTOMER",
	1: "INFRA",
}
var NamespaceSpec_Type_value = map[string]int32{
	"CUSTOMER": 0,
	"INFRA":    1,
}

func (x NamespaceSpec_Type) String() string {
	return proto.EnumName(NamespaceSpec_Type_name, int32(x))
}
func (NamespaceSpec_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorNamespace, []int{1, 0}
}

// Namespace object
type Namespace struct {
	api.TypeMeta   `protobuf:"bytes,1,opt,name=TypeMeta,embedded=TypeMeta" json:",inline"`
	api.ObjectMeta `protobuf:"bytes,2,opt,name=ObjectMeta,embedded=ObjectMeta" json:"meta,omitempty"`
	Spec           NamespaceSpec   `protobuf:"bytes,3,opt,name=Spec" json:"spec,omitempty"`
	Status         NamespaceStatus `protobuf:"bytes,4,opt,name=Status" json:"status,omitempty"`
}

func (m *Namespace) Reset()                    { *m = Namespace{} }
func (m *Namespace) String() string            { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()               {}
func (*Namespace) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{0} }

func (m *Namespace) GetSpec() NamespaceSpec {
	if m != nil {
		return m.Spec
	}
	return NamespaceSpec{}
}

func (m *Namespace) GetStatus() NamespaceStatus {
	if m != nil {
		return m.Status
	}
	return NamespaceStatus{}
}

// NamespaceSpec captures all the namespace level configuration
type NamespaceSpec struct {
	// Type of the Namespace.
	// Infra type creates a overlay VRF in the datapath. This is automatically created on bringup.
	// Customer type creates a VRF in the datapath.
	// default and infra namespace under default tenant are automatically created during init time.
	NamespaceType string `protobuf:"bytes,1,opt,name=NamespaceType,proto3" json:"namespace-type,omitempty"`
}

func (m *NamespaceSpec) Reset()                    { *m = NamespaceSpec{} }
func (m *NamespaceSpec) String() string            { return proto.CompactTextString(m) }
func (*NamespaceSpec) ProtoMessage()               {}
func (*NamespaceSpec) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{1} }

func (m *NamespaceSpec) GetNamespaceType() string {
	if m != nil {
		return m.NamespaceType
	}
	return ""
}

// Namespace Status
type NamespaceStatus struct {
	// VRF ID in the datapath
	NamespaceID uint64 `protobuf:"varint,1,opt,name=NamespaceID,proto3" json:"namespace-id,omitempty"`
}

func (m *NamespaceStatus) Reset()                    { *m = NamespaceStatus{} }
func (m *NamespaceStatus) String() string            { return proto.CompactTextString(m) }
func (*NamespaceStatus) ProtoMessage()               {}
func (*NamespaceStatus) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{2} }

func (m *NamespaceStatus) GetNamespaceID() uint64 {
	if m != nil {
		return m.NamespaceID
	}
	return 0
}

type NamespaceList struct {
	Namespaces []*Namespace `protobuf:"bytes,1,rep,name=namespaces" json:"namespaces,omitempty"`
}

func (m *NamespaceList) Reset()                    { *m = NamespaceList{} }
func (m *NamespaceList) String() string            { return proto.CompactTextString(m) }
func (*NamespaceList) ProtoMessage()               {}
func (*NamespaceList) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{3} }

func (m *NamespaceList) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

// namespace watch event
type NamespaceEvent struct {
	EventType api.EventType `protobuf:"varint,1,opt,name=EventType,proto3,enum=api.EventType" json:"event-type,omitempty"`
	Namespace Namespace     `protobuf:"bytes,2,opt,name=Namespace" json:"namespace,omitempty"`
}

func (m *NamespaceEvent) Reset()                    { *m = NamespaceEvent{} }
func (m *NamespaceEvent) String() string            { return proto.CompactTextString(m) }
func (*NamespaceEvent) ProtoMessage()               {}
func (*NamespaceEvent) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{4} }

func (m *NamespaceEvent) GetEventType() api.EventType {
	if m != nil {
		return m.EventType
	}
	return api.EventType_CreateEvent
}

func (m *NamespaceEvent) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace{}
}

// namespace watch events batched
type NamespaceEventList struct {
	NamespaceEvents []*NamespaceEvent `protobuf:"bytes,1,rep,name=NamespaceEvents" json:"NamespaceEvents,omitempty"`
}

func (m *NamespaceEventList) Reset()                    { *m = NamespaceEventList{} }
func (m *NamespaceEventList) String() string            { return proto.CompactTextString(m) }
func (*NamespaceEventList) ProtoMessage()               {}
func (*NamespaceEventList) Descriptor() ([]byte, []int) { return fileDescriptorNamespace, []int{5} }

func (m *NamespaceEventList) GetNamespaceEvents() []*NamespaceEvent {
	if m != nil {
		return m.NamespaceEvents
	}
	return nil
}

func init() {
	proto.RegisterType((*Namespace)(nil), "netproto.Namespace")
	proto.RegisterType((*NamespaceSpec)(nil), "netproto.NamespaceSpec")
	proto.RegisterType((*NamespaceStatus)(nil), "netproto.NamespaceStatus")
	proto.RegisterType((*NamespaceList)(nil), "netproto.NamespaceList")
	proto.RegisterType((*NamespaceEvent)(nil), "netproto.NamespaceEvent")
	proto.RegisterType((*NamespaceEventList)(nil), "netproto.NamespaceEventList")
	proto.RegisterEnum("netproto.NamespaceSpec_Type", NamespaceSpec_Type_name, NamespaceSpec_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NamespaceApi service

type NamespaceApiClient interface {
	GetNamespace(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*Namespace, error)
	ListNamespaces(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*NamespaceList, error)
	WatchNamespaces(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (NamespaceApi_WatchNamespacesClient, error)
	UpdateNamespace(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*Namespace, error)
	NamespaceOperUpdate(ctx context.Context, opts ...grpc.CallOption) (NamespaceApi_NamespaceOperUpdateClient, error)
}

type namespaceApiClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceApiClient(cc *grpc.ClientConn) NamespaceApiClient {
	return &namespaceApiClient{cc}
}

func (c *namespaceApiClient) GetNamespace(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := grpc.Invoke(ctx, "/netproto.NamespaceApi/GetNamespace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceApiClient) ListNamespaces(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*NamespaceList, error) {
	out := new(NamespaceList)
	err := grpc.Invoke(ctx, "/netproto.NamespaceApi/ListNamespaces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceApiClient) WatchNamespaces(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (NamespaceApi_WatchNamespacesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NamespaceApi_serviceDesc.Streams[0], c.cc, "/netproto.NamespaceApi/WatchNamespaces", opts...)
	if err != nil {
		return nil, err
	}
	x := &namespaceApiWatchNamespacesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NamespaceApi_WatchNamespacesClient interface {
	Recv() (*NamespaceEventList, error)
	grpc.ClientStream
}

type namespaceApiWatchNamespacesClient struct {
	grpc.ClientStream
}

func (x *namespaceApiWatchNamespacesClient) Recv() (*NamespaceEventList, error) {
	m := new(NamespaceEventList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *namespaceApiClient) UpdateNamespace(ctx context.Context, in *Namespace, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := grpc.Invoke(ctx, "/netproto.NamespaceApi/UpdateNamespace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceApiClient) NamespaceOperUpdate(ctx context.Context, opts ...grpc.CallOption) (NamespaceApi_NamespaceOperUpdateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NamespaceApi_serviceDesc.Streams[1], c.cc, "/netproto.NamespaceApi/NamespaceOperUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &namespaceApiNamespaceOperUpdateClient{stream}
	return x, nil
}

type NamespaceApi_NamespaceOperUpdateClient interface {
	Send(*NamespaceEvent) error
	CloseAndRecv() (*api.TypeMeta, error)
	grpc.ClientStream
}

type namespaceApiNamespaceOperUpdateClient struct {
	grpc.ClientStream
}

func (x *namespaceApiNamespaceOperUpdateClient) Send(m *NamespaceEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *namespaceApiNamespaceOperUpdateClient) CloseAndRecv() (*api.TypeMeta, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(api.TypeMeta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for NamespaceApi service

type NamespaceApiServer interface {
	GetNamespace(context.Context, *api.ObjectMeta) (*Namespace, error)
	ListNamespaces(context.Context, *api.ObjectMeta) (*NamespaceList, error)
	WatchNamespaces(*api.ObjectMeta, NamespaceApi_WatchNamespacesServer) error
	UpdateNamespace(context.Context, *Namespace) (*Namespace, error)
	NamespaceOperUpdate(NamespaceApi_NamespaceOperUpdateServer) error
}

func RegisterNamespaceApiServer(s *grpc.Server, srv NamespaceApiServer) {
	s.RegisterService(&_NamespaceApi_serviceDesc, srv)
}

func _NamespaceApi_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceApiServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netproto.NamespaceApi/GetNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceApiServer).GetNamespace(ctx, req.(*api.ObjectMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceApi_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceApiServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netproto.NamespaceApi/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceApiServer).ListNamespaces(ctx, req.(*api.ObjectMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceApi_WatchNamespaces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.ObjectMeta)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NamespaceApiServer).WatchNamespaces(m, &namespaceApiWatchNamespacesServer{stream})
}

type NamespaceApi_WatchNamespacesServer interface {
	Send(*NamespaceEventList) error
	grpc.ServerStream
}

type namespaceApiWatchNamespacesServer struct {
	grpc.ServerStream
}

func (x *namespaceApiWatchNamespacesServer) Send(m *NamespaceEventList) error {
	return x.ServerStream.SendMsg(m)
}

func _NamespaceApi_UpdateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceApiServer).UpdateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netproto.NamespaceApi/UpdateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceApiServer).UpdateNamespace(ctx, req.(*Namespace))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceApi_NamespaceOperUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NamespaceApiServer).NamespaceOperUpdate(&namespaceApiNamespaceOperUpdateServer{stream})
}

type NamespaceApi_NamespaceOperUpdateServer interface {
	SendAndClose(*api.TypeMeta) error
	Recv() (*NamespaceEvent, error)
	grpc.ServerStream
}

type namespaceApiNamespaceOperUpdateServer struct {
	grpc.ServerStream
}

func (x *namespaceApiNamespaceOperUpdateServer) SendAndClose(m *api.TypeMeta) error {
	return x.ServerStream.SendMsg(m)
}

func (x *namespaceApiNamespaceOperUpdateServer) Recv() (*NamespaceEvent, error) {
	m := new(NamespaceEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NamespaceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "netproto.NamespaceApi",
	HandlerType: (*NamespaceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNamespace",
			Handler:    _NamespaceApi_GetNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _NamespaceApi_ListNamespaces_Handler,
		},
		{
			MethodName: "UpdateNamespace",
			Handler:    _NamespaceApi_UpdateNamespace_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchNamespaces",
			Handler:       _NamespaceApi_WatchNamespaces_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NamespaceOperUpdate",
			Handler:       _NamespaceApi_NamespaceOperUpdate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "namespace.proto",
}

func (m *Namespace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Namespace) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintNamespace(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintNamespace(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintNamespace(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintNamespace(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *NamespaceSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NamespaceType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.NamespaceType)))
		i += copy(dAtA[i:], m.NamespaceType)
	}
	return i, nil
}

func (m *NamespaceStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NamespaceID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.NamespaceID))
	}
	return i, nil
}

func (m *NamespaceList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for _, msg := range m.Namespaces {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NamespaceEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EventType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNamespace(dAtA, i, uint64(m.EventType))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintNamespace(dAtA, i, uint64(m.Namespace.Size()))
	n5, err := m.Namespace.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func (m *NamespaceEventList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceEventList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NamespaceEvents) > 0 {
		for _, msg := range m.NamespaceEvents {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNamespace(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Namespace) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovNamespace(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovNamespace(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovNamespace(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovNamespace(uint64(l))
	return n
}

func (m *NamespaceSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.NamespaceType)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	return n
}

func (m *NamespaceStatus) Size() (n int) {
	var l int
	_ = l
	if m.NamespaceID != 0 {
		n += 1 + sovNamespace(uint64(m.NamespaceID))
	}
	return n
}

func (m *NamespaceList) Size() (n int) {
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for _, e := range m.Namespaces {
			l = e.Size()
			n += 1 + l + sovNamespace(uint64(l))
		}
	}
	return n
}

func (m *NamespaceEvent) Size() (n int) {
	var l int
	_ = l
	if m.EventType != 0 {
		n += 1 + sovNamespace(uint64(m.EventType))
	}
	l = m.Namespace.Size()
	n += 1 + l + sovNamespace(uint64(l))
	return n
}

func (m *NamespaceEventList) Size() (n int) {
	var l int
	_ = l
	if len(m.NamespaceEvents) > 0 {
		for _, e := range m.NamespaceEvents {
			l = e.Size()
			n += 1 + l + sovNamespace(uint64(l))
		}
	}
	return n
}

func sovNamespace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Namespace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Namespace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Namespace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceID", wireType)
			}
			m.NamespaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NamespaceID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespaces = append(m.Namespaces, &Namespace{})
			if err := m.Namespaces[len(m.Namespaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventType |= (api.EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Namespace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceEventList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceEventList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceEventList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceEvents = append(m.NamespaceEvents, &NamespaceEvent{})
			if err := m.NamespaceEvents[len(m.NamespaceEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNamespace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNamespace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNamespace
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNamespace(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNamespace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("namespace.proto", fileDescriptorNamespace) }

var fileDescriptorNamespace = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xd1, 0x4e, 0x13, 0x4d,
	0x14, 0xc7, 0xbb, 0xb4, 0x1f, 0x1f, 0x1d, 0xa0, 0x6d, 0xa6, 0x06, 0x97, 0x85, 0xb4, 0x64, 0x13,
	0x4d, 0x4d, 0x60, 0xd7, 0x94, 0xc4, 0x2b, 0xd4, 0xb0, 0x50, 0x0c, 0x41, 0xa8, 0x69, 0x21, 0x7a,
	0xe1, 0x85, 0xdb, 0xed, 0xb1, 0xac, 0x69, 0x67, 0x27, 0xec, 0x2c, 0x86, 0x18, 0xae, 0x4c, 0x7c,
	0x05, 0x7d, 0x04, 0xc3, 0x8d, 0x89, 0x4f, 0xc1, 0x25, 0x4f, 0xd0, 0x18, 0xbc, 0xe3, 0x29, 0xcc,
	0x4c, 0xdb, 0xd9, 0x69, 0xdd, 0x7a, 0xd5, 0x99, 0xd3, 0xf3, 0xff, 0x9d, 0xff, 0x39, 0x3b, 0x33,
	0x28, 0x4f, 0xdc, 0x1e, 0x84, 0xd4, 0xf5, 0xc0, 0xa2, 0x67, 0x01, 0x0b, 0xf0, 0x1c, 0x01, 0x26,
	0x56, 0xc6, 0x6a, 0x27, 0x08, 0x3a, 0x5d, 0xb0, 0x5d, 0xea, 0xdb, 0x2e, 0x21, 0x01, 0x73, 0x99,
	0x1f, 0x90, 0x70, 0x90, 0x67, 0xd4, 0x3a, 0x3e, 0x3b, 0x8d, 0x5a, 0x96, 0x17, 0xf4, 0x6c, 0x0a,
	0x24, 0x74, 0x49, 0x3b, 0xb0, 0xc3, 0x8f, 0xf6, 0x39, 0x10, 0xdf, 0x03, 0x3b, 0x62, 0x7e, 0x37,
	0xe4, 0xd2, 0x0e, 0x10, 0x55, 0x6d, 0xfb, 0xc4, 0xeb, 0x46, 0x6d, 0x18, 0x61, 0x36, 0x14, 0x4c,
	0x27, 0xe8, 0x04, 0xb6, 0x08, 0xb7, 0xa2, 0xf7, 0x62, 0x27, 0x36, 0x62, 0x35, 0x4c, 0x7f, 0x30,
	0xa5, 0x2a, 0xf7, 0xd8, 0x03, 0xe6, 0x0e, 0xd2, 0xcc, 0xab, 0x19, 0x94, 0x3d, 0x1a, 0x35, 0x86,
	0x9f, 0xa1, 0xb9, 0xe3, 0x0b, 0x0a, 0x87, 0xc0, 0x5c, 0x5d, 0x5b, 0xd3, 0x2a, 0xf3, 0xd5, 0x45,
	0xcb, 0xa5, 0xbe, 0x35, 0x0a, 0x3a, 0xc5, 0xeb, 0x7e, 0x39, 0x75, 0xd3, 0x2f, 0x6b, 0x77, 0xfd,
	0xf2, 0xff, 0xeb, 0x3e, 0xe9, 0xfa, 0x04, 0x1a, 0x52, 0x83, 0x0f, 0x10, 0xaa, 0xb7, 0x3e, 0x80,
	0xc7, 0x04, 0x61, 0x46, 0x10, 0xf2, 0x82, 0x10, 0x87, 0x1d, 0x43, 0x61, 0xe4, 0xb8, 0x95, 0xf5,
	0xa0, 0xe7, 0x33, 0xe8, 0x51, 0x76, 0xd1, 0x50, 0xe4, 0x78, 0x07, 0x65, 0x9a, 0x14, 0x3c, 0x3d,
	0x2d, 0x30, 0xf7, 0xad, 0xd1, 0xb8, 0x2d, 0xe9, 0x97, 0xff, 0xed, 0x2c, 0x71, 0x1c, 0x47, 0x85,
	0x14, 0x3c, 0x05, 0x25, 0xc4, 0xf8, 0x00, 0xcd, 0x36, 0x99, 0xcb, 0xa2, 0x50, 0xcf, 0x08, 0xcc,
	0x72, 0x12, 0x46, 0x24, 0x38, 0xfa, 0x10, 0x54, 0x08, 0xc5, 0x5e, 0x41, 0x0d, 0x11, 0xe6, 0x57,
	0x0d, 0x2d, 0x8e, 0x15, 0xc7, 0xa0, 0x04, 0xf8, 0x14, 0xc4, 0xd4, 0xb2, 0xce, 0xf3, 0xab, 0x2f,
	0xcb, 0x2b, 0x4d, 0x76, 0x56, 0x23, 0x51, 0xaf, 0x32, 0xa6, 0x10, 0xa3, 0x7c, 0x74, 0x3d, 0xe8,
	0x5e, 0x97, 0x47, 0x6a, 0x83, 0x5d, 0x50, 0x50, 0x2a, 0x8e, 0x53, 0xcd, 0x32, 0xca, 0xf0, 0x5f,
	0xbc, 0x80, 0xe6, 0x76, 0x4e, 0x9a, 0xc7, 0xf5, 0xc3, 0x5a, 0xa3, 0x90, 0xc2, 0x59, 0xf4, 0xdf,
	0xfe, 0xd1, 0x5e, 0x63, 0xbb, 0xa0, 0x99, 0x75, 0x94, 0x9f, 0x68, 0x07, 0x6f, 0xa1, 0x79, 0x19,
	0xda, 0xdf, 0x15, 0xc6, 0x32, 0x8e, 0x71, 0xd7, 0x2f, 0x2f, 0xc5, 0x55, 0xfd, 0xb6, 0x52, 0x53,
	0x4d, 0x37, 0x77, 0x95, 0xc6, 0x5e, 0xfa, 0x21, 0xc3, 0x9b, 0x08, 0x49, 0x5d, 0xa8, 0x6b, 0x6b,
	0xe9, 0xca, 0x7c, 0xb5, 0x98, 0x30, 0xcc, 0x86, 0x92, 0x66, 0x7e, 0xd7, 0x50, 0x4e, 0xfe, 0x53,
	0x3b, 0x07, 0xc2, 0xf0, 0x1e, 0xca, 0x8a, 0x85, 0x9c, 0x56, 0xae, 0x9a, 0x13, 0x27, 0x44, 0x46,
	0x1d, 0xfd, 0xae, 0x5f, 0xbe, 0x07, 0x7c, 0x3b, 0x39, 0x96, 0x58, 0x8a, 0x8f, 0x94, 0x73, 0x3b,
	0x3c, 0x69, 0x49, 0x76, 0x9c, 0x95, 0xe1, 0x57, 0x2d, 0x4a, 0x5b, 0x2a, 0x4f, 0xe6, 0x99, 0x6f,
	0x10, 0x1e, 0x77, 0x2a, 0xba, 0x76, 0x94, 0xb9, 0x8a, 0xe8, 0xa8, 0x75, 0x3d, 0xa1, 0x96, 0x48,
	0x68, 0x4c, 0x0a, 0xaa, 0x3f, 0xd2, 0x68, 0x41, 0xc6, 0xb6, 0xa9, 0x8f, 0x9f, 0xa0, 0x85, 0x17,
	0xc0, 0xe2, 0x5b, 0x37, 0x79, 0x43, 0x8c, 0xa4, 0x46, 0xcc, 0x14, 0xde, 0x42, 0x39, 0x6e, 0x4a,
	0x86, 0xc2, 0xbf, 0x95, 0x49, 0xb7, 0x84, 0x6b, 0xcc, 0x14, 0xde, 0x45, 0xf9, 0xd7, 0x2e, 0xf3,
	0x4e, 0xff, 0x25, 0x5f, 0x9d, 0xd6, 0xd5, 0x80, 0xf1, 0x58, 0xc3, 0x4f, 0x51, 0xfe, 0x84, 0xb6,
	0x5d, 0x06, 0xb1, 0xfd, 0x24, 0xb7, 0xd3, 0x5a, 0x70, 0x50, 0x51, 0x6e, 0xeb, 0x14, 0xce, 0x06,
	0x2c, 0x3c, 0x75, 0x9a, 0xc6, 0xf8, 0xfb, 0x63, 0xa6, 0x2a, 0x9a, 0xf1, 0xee, 0xe7, 0xe7, 0xe5,
	0xb7, 0xea, 0xab, 0x95, 0xe9, 0xf2, 0x4f, 0x95, 0xa1, 0x41, 0xc8, 0xf0, 0x6c, 0x1b, 0xba, 0xc0,
	0x00, 0xa7, 0x69, 0xc4, 0x8c, 0x87, 0xf6, 0xa7, 0xb8, 0x39, 0xeb, 0x18, 0x88, 0x4b, 0xd8, 0xe5,
	0x58, 0x8c, 0x43, 0x2e, 0xcd, 0x1c, 0x7f, 0x14, 0xe3, 0x63, 0xeb, 0x14, 0xae, 0x6f, 0x4b, 0xda,
	0xcd, 0x6d, 0x49, 0xfb, 0x75, 0x5b, 0xd2, 0xbe, 0xfd, 0x2e, 0xa5, 0x5e, 0x69, 0xad, 0x59, 0x61,
	0x6f, 0xf3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x77, 0x5e, 0xb2, 0x07, 0x06, 0x00, 0x00,
}
