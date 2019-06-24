// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vrf.proto

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

type VrfSpec_Type int32

const (
	VrfSpec_CUSTOMER VrfSpec_Type = 0
	VrfSpec_INFRA    VrfSpec_Type = 1
)

var VrfSpec_Type_name = map[int32]string{
	0: "CUSTOMER",
	1: "INFRA",
}
var VrfSpec_Type_value = map[string]int32{
	"CUSTOMER": 0,
	"INFRA":    1,
}

func (x VrfSpec_Type) String() string {
	return proto.EnumName(VrfSpec_Type_name, int32(x))
}
func (VrfSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorVrf, []int{1, 0} }

// Vrf object
type Vrf struct {
	api.TypeMeta   `protobuf:"bytes,1,opt,name=TypeMeta,embedded=TypeMeta" json:",inline"`
	api.ObjectMeta `protobuf:"bytes,2,opt,name=ObjectMeta,embedded=ObjectMeta" json:"meta,omitempty"`
	Spec           VrfSpec   `protobuf:"bytes,3,opt,name=Spec" json:"spec,omitempty"`
	Status         VrfStatus `protobuf:"bytes,4,opt,name=Status" json:"status,omitempty"`
}

func (m *Vrf) Reset()                    { *m = Vrf{} }
func (m *Vrf) String() string            { return proto.CompactTextString(m) }
func (*Vrf) ProtoMessage()               {}
func (*Vrf) Descriptor() ([]byte, []int) { return fileDescriptorVrf, []int{0} }

func (m *Vrf) GetSpec() VrfSpec {
	if m != nil {
		return m.Spec
	}
	return VrfSpec{}
}

func (m *Vrf) GetStatus() VrfStatus {
	if m != nil {
		return m.Status
	}
	return VrfStatus{}
}

// VrfSpec captures all the vrf level configuration
type VrfSpec struct {
	// Type of the Vrf.
	// Infra type creates a overlay VRF in the datapath. This is automatically created on bringup.
	// Customer type creates a VRF in the datapath.
	// default and infra vrf under default tenant are automatically created during init time.
	VrfType string `protobuf:"bytes,1,opt,name=VrfType,proto3" json:"vrf-type,omitempty"`
}

func (m *VrfSpec) Reset()                    { *m = VrfSpec{} }
func (m *VrfSpec) String() string            { return proto.CompactTextString(m) }
func (*VrfSpec) ProtoMessage()               {}
func (*VrfSpec) Descriptor() ([]byte, []int) { return fileDescriptorVrf, []int{1} }

func (m *VrfSpec) GetVrfType() string {
	if m != nil {
		return m.VrfType
	}
	return ""
}

// Vrf Status
type VrfStatus struct {
	// VRF ID in the datapath
	VrfID uint64 `protobuf:"varint,1,opt,name=VrfID,proto3" json:"vrf-id,omitempty"`
}

func (m *VrfStatus) Reset()                    { *m = VrfStatus{} }
func (m *VrfStatus) String() string            { return proto.CompactTextString(m) }
func (*VrfStatus) ProtoMessage()               {}
func (*VrfStatus) Descriptor() ([]byte, []int) { return fileDescriptorVrf, []int{2} }

func (m *VrfStatus) GetVrfID() uint64 {
	if m != nil {
		return m.VrfID
	}
	return 0
}

type VrfList struct {
	Vrfs []*Vrf `protobuf:"bytes,1,rep,name=vrfs" json:"vrfs,omitempty"`
}

func (m *VrfList) Reset()                    { *m = VrfList{} }
func (m *VrfList) String() string            { return proto.CompactTextString(m) }
func (*VrfList) ProtoMessage()               {}
func (*VrfList) Descriptor() ([]byte, []int) { return fileDescriptorVrf, []int{3} }

func (m *VrfList) GetVrfs() []*Vrf {
	if m != nil {
		return m.Vrfs
	}
	return nil
}

// vrf watch event
type VrfEvent struct {
	EventType api.EventType `protobuf:"varint,1,opt,name=EventType,proto3,enum=api.EventType" json:"event-type,omitempty"`
	Vrf       Vrf           `protobuf:"bytes,2,opt,name=Vrf" json:"vrf,omitempty"`
}

func (m *VrfEvent) Reset()                    { *m = VrfEvent{} }
func (m *VrfEvent) String() string            { return proto.CompactTextString(m) }
func (*VrfEvent) ProtoMessage()               {}
func (*VrfEvent) Descriptor() ([]byte, []int) { return fileDescriptorVrf, []int{4} }

func (m *VrfEvent) GetEventType() api.EventType {
	if m != nil {
		return m.EventType
	}
	return api.EventType_CreateEvent
}

func (m *VrfEvent) GetVrf() Vrf {
	if m != nil {
		return m.Vrf
	}
	return Vrf{}
}

func init() {
	proto.RegisterType((*Vrf)(nil), "netproto.Vrf")
	proto.RegisterType((*VrfSpec)(nil), "netproto.VrfSpec")
	proto.RegisterType((*VrfStatus)(nil), "netproto.VrfStatus")
	proto.RegisterType((*VrfList)(nil), "netproto.VrfList")
	proto.RegisterType((*VrfEvent)(nil), "netproto.VrfEvent")
	proto.RegisterEnum("netproto.VrfSpec_Type", VrfSpec_Type_name, VrfSpec_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VrfApi service

type VrfApiClient interface {
	GetVrf(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*Vrf, error)
	ListVrfs(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*VrfList, error)
	WatchVrfs(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (VrfApi_WatchVrfsClient, error)
	VrfOperUpdate(ctx context.Context, opts ...grpc.CallOption) (VrfApi_VrfOperUpdateClient, error)
}

type vrfApiClient struct {
	cc *grpc.ClientConn
}

func NewVrfApiClient(cc *grpc.ClientConn) VrfApiClient {
	return &vrfApiClient{cc}
}

func (c *vrfApiClient) GetVrf(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*Vrf, error) {
	out := new(Vrf)
	err := grpc.Invoke(ctx, "/netproto.VrfApi/GetVrf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vrfApiClient) ListVrfs(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (*VrfList, error) {
	out := new(VrfList)
	err := grpc.Invoke(ctx, "/netproto.VrfApi/ListVrfs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vrfApiClient) WatchVrfs(ctx context.Context, in *api.ObjectMeta, opts ...grpc.CallOption) (VrfApi_WatchVrfsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VrfApi_serviceDesc.Streams[0], c.cc, "/netproto.VrfApi/WatchVrfs", opts...)
	if err != nil {
		return nil, err
	}
	x := &vrfApiWatchVrfsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VrfApi_WatchVrfsClient interface {
	Recv() (*VrfEvent, error)
	grpc.ClientStream
}

type vrfApiWatchVrfsClient struct {
	grpc.ClientStream
}

func (x *vrfApiWatchVrfsClient) Recv() (*VrfEvent, error) {
	m := new(VrfEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vrfApiClient) VrfOperUpdate(ctx context.Context, opts ...grpc.CallOption) (VrfApi_VrfOperUpdateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VrfApi_serviceDesc.Streams[1], c.cc, "/netproto.VrfApi/VrfOperUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &vrfApiVrfOperUpdateClient{stream}
	return x, nil
}

type VrfApi_VrfOperUpdateClient interface {
	Send(*VrfEvent) error
	CloseAndRecv() (*api.TypeMeta, error)
	grpc.ClientStream
}

type vrfApiVrfOperUpdateClient struct {
	grpc.ClientStream
}

func (x *vrfApiVrfOperUpdateClient) Send(m *VrfEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vrfApiVrfOperUpdateClient) CloseAndRecv() (*api.TypeMeta, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(api.TypeMeta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for VrfApi service

type VrfApiServer interface {
	GetVrf(context.Context, *api.ObjectMeta) (*Vrf, error)
	ListVrfs(context.Context, *api.ObjectMeta) (*VrfList, error)
	WatchVrfs(*api.ObjectMeta, VrfApi_WatchVrfsServer) error
	VrfOperUpdate(VrfApi_VrfOperUpdateServer) error
}

func RegisterVrfApiServer(s *grpc.Server, srv VrfApiServer) {
	s.RegisterService(&_VrfApi_serviceDesc, srv)
}

func _VrfApi_GetVrf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VrfApiServer).GetVrf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netproto.VrfApi/GetVrf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VrfApiServer).GetVrf(ctx, req.(*api.ObjectMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _VrfApi_ListVrfs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ObjectMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VrfApiServer).ListVrfs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netproto.VrfApi/ListVrfs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VrfApiServer).ListVrfs(ctx, req.(*api.ObjectMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _VrfApi_WatchVrfs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.ObjectMeta)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VrfApiServer).WatchVrfs(m, &vrfApiWatchVrfsServer{stream})
}

type VrfApi_WatchVrfsServer interface {
	Send(*VrfEvent) error
	grpc.ServerStream
}

type vrfApiWatchVrfsServer struct {
	grpc.ServerStream
}

func (x *vrfApiWatchVrfsServer) Send(m *VrfEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _VrfApi_VrfOperUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VrfApiServer).VrfOperUpdate(&vrfApiVrfOperUpdateServer{stream})
}

type VrfApi_VrfOperUpdateServer interface {
	SendAndClose(*api.TypeMeta) error
	Recv() (*VrfEvent, error)
	grpc.ServerStream
}

type vrfApiVrfOperUpdateServer struct {
	grpc.ServerStream
}

func (x *vrfApiVrfOperUpdateServer) SendAndClose(m *api.TypeMeta) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vrfApiVrfOperUpdateServer) Recv() (*VrfEvent, error) {
	m := new(VrfEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VrfApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "netproto.VrfApi",
	HandlerType: (*VrfApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVrf",
			Handler:    _VrfApi_GetVrf_Handler,
		},
		{
			MethodName: "ListVrfs",
			Handler:    _VrfApi_ListVrfs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchVrfs",
			Handler:       _VrfApi_WatchVrfs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "VrfOperUpdate",
			Handler:       _VrfApi_VrfOperUpdate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "vrf.proto",
}

func (m *Vrf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vrf) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintVrf(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintVrf(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintVrf(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintVrf(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *VrfSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VrfSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VrfType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintVrf(dAtA, i, uint64(len(m.VrfType)))
		i += copy(dAtA[i:], m.VrfType)
	}
	return i, nil
}

func (m *VrfStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VrfStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.VrfID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintVrf(dAtA, i, uint64(m.VrfID))
	}
	return i, nil
}

func (m *VrfList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VrfList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Vrfs) > 0 {
		for _, msg := range m.Vrfs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintVrf(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *VrfEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VrfEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EventType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintVrf(dAtA, i, uint64(m.EventType))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintVrf(dAtA, i, uint64(m.Vrf.Size()))
	n5, err := m.Vrf.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func encodeVarintVrf(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Vrf) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovVrf(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovVrf(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovVrf(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovVrf(uint64(l))
	return n
}

func (m *VrfSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.VrfType)
	if l > 0 {
		n += 1 + l + sovVrf(uint64(l))
	}
	return n
}

func (m *VrfStatus) Size() (n int) {
	var l int
	_ = l
	if m.VrfID != 0 {
		n += 1 + sovVrf(uint64(m.VrfID))
	}
	return n
}

func (m *VrfList) Size() (n int) {
	var l int
	_ = l
	if len(m.Vrfs) > 0 {
		for _, e := range m.Vrfs {
			l = e.Size()
			n += 1 + l + sovVrf(uint64(l))
		}
	}
	return n
}

func (m *VrfEvent) Size() (n int) {
	var l int
	_ = l
	if m.EventType != 0 {
		n += 1 + sovVrf(uint64(m.EventType))
	}
	l = m.Vrf.Size()
	n += 1 + l + sovVrf(uint64(l))
	return n
}

func sovVrf(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVrf(x uint64) (n int) {
	return sovVrf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vrf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrf
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
			return fmt.Errorf("proto: Vrf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vrf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
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
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
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
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
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
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
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
			skippy, err := skipVrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVrf
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
func (m *VrfSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrf
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
			return fmt.Errorf("proto: VrfSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VrfSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VrfType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VrfType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVrf
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
func (m *VrfStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrf
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
			return fmt.Errorf("proto: VrfStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VrfStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VrfID", wireType)
			}
			m.VrfID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VrfID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVrf
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
func (m *VrfList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrf
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
			return fmt.Errorf("proto: VrfList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VrfList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vrfs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vrfs = append(m.Vrfs, &Vrf{})
			if err := m.Vrfs[len(m.Vrfs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVrf
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
func (m *VrfEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrf
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
			return fmt.Errorf("proto: VrfEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VrfEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrf
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
				return fmt.Errorf("proto: wrong wireType = %d for field Vrf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrf
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
				return ErrInvalidLengthVrf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Vrf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVrf
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
func skipVrf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVrf
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
					return 0, ErrIntOverflowVrf
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
					return 0, ErrIntOverflowVrf
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
				return 0, ErrInvalidLengthVrf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVrf
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
				next, err := skipVrf(dAtA[start:])
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
	ErrInvalidLengthVrf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVrf   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("vrf.proto", fileDescriptorVrf) }

var fileDescriptorVrf = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x4e, 0xdb, 0x4c,
	0x10, 0xc7, 0x6d, 0x62, 0x42, 0xb2, 0xdf, 0x07, 0x4d, 0x17, 0xa8, 0x82, 0x55, 0x25, 0xd4, 0x52,
	0x25, 0x8a, 0xc0, 0x46, 0x50, 0xd1, 0x13, 0x95, 0x70, 0x1b, 0x2a, 0xd4, 0x02, 0x55, 0x02, 0xee,
	0xd9, 0x71, 0xc6, 0xc6, 0x6d, 0xb2, 0x5e, 0xad, 0xd7, 0xa9, 0xa2, 0x8a, 0x53, 0xa5, 0x1e, 0xfa,
	0x04, 0x3c, 0x43, 0x8f, 0x7d, 0x83, 0xde, 0x38, 0xf2, 0x04, 0x51, 0x45, 0x6f, 0x3c, 0x45, 0xb5,
	0x1b, 0x27, 0xd8, 0x40, 0x4f, 0xde, 0x1d, 0xff, 0xff, 0xbf, 0x9d, 0x99, 0x9d, 0x45, 0xe5, 0x3e,
	0xf3, 0x4d, 0xca, 0x22, 0x1e, 0xe1, 0x12, 0x01, 0x2e, 0x57, 0xfa, 0xe3, 0x20, 0x8a, 0x82, 0x2e,
	0x58, 0x2e, 0x0d, 0x2d, 0x97, 0x90, 0x88, 0xbb, 0x3c, 0x8c, 0x48, 0x3c, 0xd2, 0xe9, 0x8d, 0x20,
	0xe4, 0xa7, 0x49, 0xdb, 0xf4, 0xa2, 0x9e, 0x45, 0x81, 0xc4, 0x2e, 0xe9, 0x44, 0x56, 0xfc, 0xd9,
	0xea, 0x03, 0x09, 0x3d, 0xb0, 0x12, 0x1e, 0x76, 0x63, 0x61, 0x0d, 0x80, 0x64, 0xdd, 0x56, 0x48,
	0xbc, 0x6e, 0xd2, 0x81, 0x31, 0x66, 0x3d, 0x83, 0x09, 0xa2, 0x20, 0xb2, 0x64, 0xb8, 0x9d, 0xf8,
	0x72, 0x27, 0x37, 0x72, 0x95, 0xca, 0x9f, 0xfe, 0xe3, 0x54, 0x91, 0x63, 0x0f, 0xb8, 0x3b, 0x92,
	0x19, 0xe7, 0x53, 0xa8, 0xe0, 0x30, 0x1f, 0xbf, 0x44, 0xa5, 0xe3, 0x01, 0x85, 0x03, 0xe0, 0x6e,
	0x55, 0x5d, 0x56, 0x57, 0xfe, 0xdb, 0x9c, 0x35, 0x5d, 0x1a, 0x9a, 0xe3, 0xa0, 0x3d, 0x7f, 0x31,
	0xac, 0x2b, 0x97, 0xc3, 0xba, 0x7a, 0x3d, 0xac, 0xcf, 0xac, 0x85, 0xa4, 0x1b, 0x12, 0x68, 0x4e,
	0x3c, 0xf8, 0x2d, 0x42, 0x47, 0xed, 0x8f, 0xe0, 0x71, 0x49, 0x98, 0x92, 0x84, 0x07, 0x92, 0x70,
	0x13, 0xb6, 0xf5, 0x0c, 0x63, 0x4e, 0x24, 0xb1, 0x16, 0xf5, 0x42, 0x0e, 0x3d, 0xca, 0x07, 0xcd,
	0x8c, 0x1d, 0xef, 0x20, 0xad, 0x45, 0xc1, 0xab, 0x16, 0x24, 0xe6, 0xa1, 0x39, 0x6e, 0xb4, 0xe9,
	0x30, 0x5f, 0xfc, 0xb0, 0x1f, 0x09, 0x90, 0x80, 0xc4, 0x14, 0xbc, 0x0c, 0x44, 0xda, 0x70, 0x03,
	0x15, 0x5b, 0xdc, 0xe5, 0x49, 0x5c, 0xd5, 0x24, 0x60, 0x3e, 0x0f, 0x90, 0xbf, 0xec, 0x6a, 0x8a,
	0xa8, 0xc4, 0x72, 0x9f, 0x81, 0xa4, 0x66, 0x63, 0x80, 0x66, 0xd2, 0xf3, 0xf0, 0x81, 0x5c, 0x8a,
	0x62, 0x65, 0x73, 0xca, 0xf6, 0xd6, 0x8f, 0x6f, 0x4b, 0x8b, 0x2d, 0xce, 0x1a, 0x24, 0xe9, 0xad,
	0xa4, 0x2a, 0xd9, 0xab, 0x67, 0x17, 0xa3, 0xf2, 0x70, 0x9f, 0xf9, 0xeb, 0x7c, 0x40, 0x21, 0x03,
	0x1e, 0x33, 0x8c, 0x3a, 0xd2, 0xc4, 0x17, 0xff, 0x8f, 0x4a, 0xaf, 0x4e, 0x5a, 0xc7, 0x47, 0x07,
	0x8d, 0x66, 0x45, 0xc1, 0x65, 0x34, 0xbd, 0x7f, 0xb8, 0xd7, 0xdc, 0xad, 0xa8, 0xc6, 0x0b, 0x54,
	0x9e, 0x64, 0x8a, 0x57, 0xd1, 0xb4, 0xc3, 0xfc, 0xfd, 0xd7, 0xf2, 0x68, 0xcd, 0x5e, 0x10, 0x49,
	0x0b, 0x7a, 0xd8, 0xc9, 0xb0, 0x47, 0x12, 0x63, 0x4d, 0x26, 0xfa, 0x2e, 0x8c, 0x39, 0x7e, 0x82,
	0xb4, 0x3e, 0xf3, 0xe3, 0xaa, 0xba, 0x5c, 0x90, 0xb7, 0x99, 0xed, 0x41, 0x53, 0xfe, 0x32, 0xbe,
	0xab, 0xa8, 0xe4, 0x30, 0xbf, 0xd1, 0x07, 0xc2, 0xf1, 0x1e, 0x2a, 0xcb, 0xc5, 0xa4, 0xca, 0xb9,
	0xcd, 0x39, 0x79, 0x81, 0x93, 0xa8, 0x5d, 0xbd, 0x1e, 0xd6, 0x17, 0x40, 0x6c, 0x6f, 0x97, 0x76,
	0x63, 0xc5, 0xdb, 0x72, 0xa0, 0xd2, 0x11, 0xc8, 0x1f, 0x6b, 0x2f, 0xa6, 0x4d, 0x9f, 0xed, 0x33,
	0x3f, 0xe3, 0x16, 0x86, 0xcd, 0x5f, 0x53, 0xa8, 0xe8, 0x30, 0x7f, 0x97, 0x86, 0x78, 0x15, 0x15,
	0xdf, 0x00, 0x17, 0x63, 0x79, 0x7b, 0x84, 0xf4, 0x3c, 0xd0, 0x50, 0xf0, 0x06, 0x2a, 0x89, 0x72,
	0x1d, 0xe6, 0xc7, 0x77, 0xd5, 0xf9, 0xd1, 0x11, 0x3a, 0x43, 0xc1, 0xcf, 0x51, 0xf9, 0x83, 0xcb,
	0xbd, 0xd3, 0xfb, 0x2d, 0x38, 0x67, 0x91, 0x65, 0x19, 0xca, 0x86, 0x8a, 0xb7, 0xd1, 0xac, 0xc3,
	0xfc, 0x23, 0x0a, 0xec, 0x84, 0x76, 0x5c, 0x0e, 0xf8, 0x1e, 0xa1, 0x9e, 0x7f, 0x33, 0x86, 0xb2,
	0xa2, 0xea, 0x9f, 0x7e, 0x7e, 0x5d, 0x0a, 0x46, 0x6f, 0x4c, 0xeb, 0x8a, 0x7b, 0xd1, 0x68, 0x14,
	0x73, 0x5c, 0xec, 0x40, 0x17, 0x38, 0xe0, 0x02, 0x4d, 0xb8, 0xbe, 0x63, 0x7d, 0xb9, 0xc9, 0xc0,
	0x3c, 0x06, 0xe2, 0x12, 0x7e, 0x96, 0x8b, 0x1d, 0xba, 0x3d, 0x88, 0xa9, 0xeb, 0xc1, 0xdd, 0xf0,
	0x99, 0x51, 0x12, 0x6f, 0x5a, 0x5c, 0xa8, 0x5d, 0xb9, 0xb8, 0xaa, 0xa9, 0x97, 0x57, 0x35, 0xf5,
	0xf7, 0x55, 0x4d, 0x3d, 0xff, 0x53, 0x53, 0xde, 0xab, 0xed, 0xa2, 0xcc, 0x71, 0xeb, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x90, 0x61, 0xb2, 0x5c, 0xba, 0x04, 0x00, 0x00,
}
