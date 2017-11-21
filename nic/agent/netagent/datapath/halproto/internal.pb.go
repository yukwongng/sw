// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProgramAddressReq struct {
	// Handle. E.g.: 'iris', 'p4plus'.
	Handle string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
	// Program name to resolve
	ProgName string `protobuf:"bytes,2,opt,name=prog_name,json=progName,proto3" json:"prog_name,omitempty"`
	// If resolve_label is false, returns the base address of the program.
	// Otherwise returns the PC offset (14 bits) of the program.
	ResolveLabel bool `protobuf:"varint,3,opt,name=resolve_label,json=resolveLabel,proto3" json:"resolve_label,omitempty"`
	// Label name to resolve
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (m *ProgramAddressReq) Reset()                    { *m = ProgramAddressReq{} }
func (m *ProgramAddressReq) String() string            { return proto.CompactTextString(m) }
func (*ProgramAddressReq) ProtoMessage()               {}
func (*ProgramAddressReq) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *ProgramAddressReq) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *ProgramAddressReq) GetProgName() string {
	if m != nil {
		return m.ProgName
	}
	return ""
}

func (m *ProgramAddressReq) GetResolveLabel() bool {
	if m != nil {
		return m.ResolveLabel
	}
	return false
}

func (m *ProgramAddressReq) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ProgramAddressResp struct {
	Addr int64 `protobuf:"varint,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *ProgramAddressResp) Reset()                    { *m = ProgramAddressResp{} }
func (m *ProgramAddressResp) String() string            { return proto.CompactTextString(m) }
func (*ProgramAddressResp) ProtoMessage()               {}
func (*ProgramAddressResp) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *ProgramAddressResp) GetAddr() int64 {
	if m != nil {
		return m.Addr
	}
	return 0
}

type GetProgramAddressRequestMsg struct {
	Reqs []*ProgramAddressReq `protobuf:"bytes,1,rep,name=reqs" json:"reqs,omitempty"`
}

func (m *GetProgramAddressRequestMsg) Reset()         { *m = GetProgramAddressRequestMsg{} }
func (m *GetProgramAddressRequestMsg) String() string { return proto.CompactTextString(m) }
func (*GetProgramAddressRequestMsg) ProtoMessage()    {}
func (*GetProgramAddressRequestMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorInternal, []int{2}
}

func (m *GetProgramAddressRequestMsg) GetReqs() []*ProgramAddressReq {
	if m != nil {
		return m.Reqs
	}
	return nil
}

type ProgramAddressResponseMsg struct {
	Resps []*ProgramAddressResp `protobuf:"bytes,1,rep,name=resps" json:"resps,omitempty"`
}

func (m *ProgramAddressResponseMsg) Reset()         { *m = ProgramAddressResponseMsg{} }
func (m *ProgramAddressResponseMsg) String() string { return proto.CompactTextString(m) }
func (*ProgramAddressResponseMsg) ProtoMessage()    {}
func (*ProgramAddressResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorInternal, []int{3}
}

func (m *ProgramAddressResponseMsg) GetResps() []*ProgramAddressResp {
	if m != nil {
		return m.Resps
	}
	return nil
}

type HbmAddressReq struct {
	// handle that specifies region in config (json) file
	Handle string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
}

func (m *HbmAddressReq) Reset()                    { *m = HbmAddressReq{} }
func (m *HbmAddressReq) String() string            { return proto.CompactTextString(m) }
func (*HbmAddressReq) ProtoMessage()               {}
func (*HbmAddressReq) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *HbmAddressReq) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

type HbmAddressResp struct {
	Addr  int64  `protobuf:"varint,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Size_ uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (m *HbmAddressResp) Reset()                    { *m = HbmAddressResp{} }
func (m *HbmAddressResp) String() string            { return proto.CompactTextString(m) }
func (*HbmAddressResp) ProtoMessage()               {}
func (*HbmAddressResp) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

func (m *HbmAddressResp) GetAddr() int64 {
	if m != nil {
		return m.Addr
	}
	return 0
}

func (m *HbmAddressResp) GetSize_() uint32 {
	if m != nil {
		return m.Size_
	}
	return 0
}

type AllocHbmAddressRequestMsg struct {
	Reqs []*HbmAddressReq `protobuf:"bytes,1,rep,name=reqs" json:"reqs,omitempty"`
}

func (m *AllocHbmAddressRequestMsg) Reset()         { *m = AllocHbmAddressRequestMsg{} }
func (m *AllocHbmAddressRequestMsg) String() string { return proto.CompactTextString(m) }
func (*AllocHbmAddressRequestMsg) ProtoMessage()    {}
func (*AllocHbmAddressRequestMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorInternal, []int{6}
}

func (m *AllocHbmAddressRequestMsg) GetReqs() []*HbmAddressReq {
	if m != nil {
		return m.Reqs
	}
	return nil
}

type AllocHbmAddressResponseMsg struct {
	Resps []*HbmAddressResp `protobuf:"bytes,1,rep,name=resps" json:"resps,omitempty"`
}

func (m *AllocHbmAddressResponseMsg) Reset()         { *m = AllocHbmAddressResponseMsg{} }
func (m *AllocHbmAddressResponseMsg) String() string { return proto.CompactTextString(m) }
func (*AllocHbmAddressResponseMsg) ProtoMessage()    {}
func (*AllocHbmAddressResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorInternal, []int{7}
}

func (m *AllocHbmAddressResponseMsg) GetResps() []*HbmAddressResp {
	if m != nil {
		return m.Resps
	}
	return nil
}

func init() {
	proto.RegisterType((*ProgramAddressReq)(nil), "internal.ProgramAddressReq")
	proto.RegisterType((*ProgramAddressResp)(nil), "internal.ProgramAddressResp")
	proto.RegisterType((*GetProgramAddressRequestMsg)(nil), "internal.GetProgramAddressRequestMsg")
	proto.RegisterType((*ProgramAddressResponseMsg)(nil), "internal.ProgramAddressResponseMsg")
	proto.RegisterType((*HbmAddressReq)(nil), "internal.HbmAddressReq")
	proto.RegisterType((*HbmAddressResp)(nil), "internal.HbmAddressResp")
	proto.RegisterType((*AllocHbmAddressRequestMsg)(nil), "internal.AllocHbmAddressRequestMsg")
	proto.RegisterType((*AllocHbmAddressResponseMsg)(nil), "internal.AllocHbmAddressResponseMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Internal service

type InternalClient interface {
	// Program resolution related APIs
	GetProgramAddress(ctx context.Context, in *GetProgramAddressRequestMsg, opts ...grpc.CallOption) (*ProgramAddressResponseMsg, error)
	// Program resolution related APIs
	AllocHbmAddress(ctx context.Context, in *AllocHbmAddressRequestMsg, opts ...grpc.CallOption) (*AllocHbmAddressResponseMsg, error)
}

type internalClient struct {
	cc *grpc.ClientConn
}

func NewInternalClient(cc *grpc.ClientConn) InternalClient {
	return &internalClient{cc}
}

func (c *internalClient) GetProgramAddress(ctx context.Context, in *GetProgramAddressRequestMsg, opts ...grpc.CallOption) (*ProgramAddressResponseMsg, error) {
	out := new(ProgramAddressResponseMsg)
	err := grpc.Invoke(ctx, "/internal.Internal/GetProgramAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) AllocHbmAddress(ctx context.Context, in *AllocHbmAddressRequestMsg, opts ...grpc.CallOption) (*AllocHbmAddressResponseMsg, error) {
	out := new(AllocHbmAddressResponseMsg)
	err := grpc.Invoke(ctx, "/internal.Internal/AllocHbmAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Internal service

type InternalServer interface {
	// Program resolution related APIs
	GetProgramAddress(context.Context, *GetProgramAddressRequestMsg) (*ProgramAddressResponseMsg, error)
	// Program resolution related APIs
	AllocHbmAddress(context.Context, *AllocHbmAddressRequestMsg) (*AllocHbmAddressResponseMsg, error)
}

func RegisterInternalServer(s *grpc.Server, srv InternalServer) {
	s.RegisterService(&_Internal_serviceDesc, srv)
}

func _Internal_GetProgramAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgramAddressRequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).GetProgramAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Internal/GetProgramAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).GetProgramAddress(ctx, req.(*GetProgramAddressRequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_AllocHbmAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocHbmAddressRequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).AllocHbmAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Internal/AllocHbmAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).AllocHbmAddress(ctx, req.(*AllocHbmAddressRequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Internal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Internal",
	HandlerType: (*InternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProgramAddress",
			Handler:    _Internal_GetProgramAddress_Handler,
		},
		{
			MethodName: "AllocHbmAddress",
			Handler:    _Internal_AllocHbmAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal.proto",
}

func (m *ProgramAddressReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProgramAddressReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Handle) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Handle)))
		i += copy(dAtA[i:], m.Handle)
	}
	if len(m.ProgName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.ProgName)))
		i += copy(dAtA[i:], m.ProgName)
	}
	if m.ResolveLabel {
		dAtA[i] = 0x18
		i++
		if m.ResolveLabel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Label) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Label)))
		i += copy(dAtA[i:], m.Label)
	}
	return i, nil
}

func (m *ProgramAddressResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProgramAddressResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Addr != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.Addr))
	}
	return i, nil
}

func (m *GetProgramAddressRequestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetProgramAddressRequestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Reqs) > 0 {
		for _, msg := range m.Reqs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInternal(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ProgramAddressResponseMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProgramAddressResponseMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Resps) > 0 {
		for _, msg := range m.Resps {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInternal(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HbmAddressReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HbmAddressReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Handle) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Handle)))
		i += copy(dAtA[i:], m.Handle)
	}
	return i, nil
}

func (m *HbmAddressResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HbmAddressResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Addr != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.Addr))
	}
	if m.Size_ != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.Size_))
	}
	return i, nil
}

func (m *AllocHbmAddressRequestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocHbmAddressRequestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Reqs) > 0 {
		for _, msg := range m.Reqs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInternal(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *AllocHbmAddressResponseMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocHbmAddressResponseMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Resps) > 0 {
		for _, msg := range m.Resps {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInternal(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Internal(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Internal(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintInternal(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProgramAddressReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Handle)
	if l > 0 {
		n += 1 + l + sovInternal(uint64(l))
	}
	l = len(m.ProgName)
	if l > 0 {
		n += 1 + l + sovInternal(uint64(l))
	}
	if m.ResolveLabel {
		n += 2
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovInternal(uint64(l))
	}
	return n
}

func (m *ProgramAddressResp) Size() (n int) {
	var l int
	_ = l
	if m.Addr != 0 {
		n += 1 + sovInternal(uint64(m.Addr))
	}
	return n
}

func (m *GetProgramAddressRequestMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Reqs) > 0 {
		for _, e := range m.Reqs {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func (m *ProgramAddressResponseMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Resps) > 0 {
		for _, e := range m.Resps {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func (m *HbmAddressReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Handle)
	if l > 0 {
		n += 1 + l + sovInternal(uint64(l))
	}
	return n
}

func (m *HbmAddressResp) Size() (n int) {
	var l int
	_ = l
	if m.Addr != 0 {
		n += 1 + sovInternal(uint64(m.Addr))
	}
	if m.Size_ != 0 {
		n += 1 + sovInternal(uint64(m.Size_))
	}
	return n
}

func (m *AllocHbmAddressRequestMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Reqs) > 0 {
		for _, e := range m.Reqs {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func (m *AllocHbmAddressResponseMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Resps) > 0 {
		for _, e := range m.Resps {
			l = e.Size()
			n += 1 + l + sovInternal(uint64(l))
		}
	}
	return n
}

func sovInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternal(x uint64) (n int) {
	return sovInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProgramAddressReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: ProgramAddressReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProgramAddressReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveLabel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ResolveLabel = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *ProgramAddressResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: ProgramAddressResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProgramAddressResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			m.Addr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Addr |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *GetProgramAddressRequestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: GetProgramAddressRequestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetProgramAddressRequestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reqs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reqs = append(m.Reqs, &ProgramAddressReq{})
			if err := m.Reqs[len(m.Reqs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *ProgramAddressResponseMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: ProgramAddressResponseMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProgramAddressResponseMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resps = append(m.Resps, &ProgramAddressResp{})
			if err := m.Resps[len(m.Resps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *HbmAddressReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: HbmAddressReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HbmAddressReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *HbmAddressResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: HbmAddressResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HbmAddressResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			m.Addr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Addr |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *AllocHbmAddressRequestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: AllocHbmAddressRequestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocHbmAddressRequestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reqs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reqs = append(m.Reqs, &HbmAddressReq{})
			if err := m.Reqs[len(m.Reqs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func (m *AllocHbmAddressResponseMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: AllocHbmAddressResponseMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocHbmAddressResponseMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
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
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resps = append(m.Resps, &HbmAddressResp{})
			if err := m.Resps[len(m.Resps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
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
func skipInternal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
				return 0, ErrInvalidLengthInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternal
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
				next, err := skipInternal(dAtA[start:])
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
	ErrInvalidLengthInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x65, 0x3e, 0x7e, 0x52, 0xee, 0xf7, 0xc1, 0x17, 0x26, 0x46, 0x4b, 0x31, 0x0d, 0x19, 0x34,
	0x36, 0x31, 0xc1, 0x04, 0x37, 0x6e, 0x71, 0x23, 0x26, 0x88, 0xa6, 0x4b, 0x17, 0x92, 0xc1, 0x4e,
	0x80, 0x64, 0x68, 0xcb, 0x4c, 0x75, 0xe1, 0xd6, 0x97, 0xf0, 0x91, 0xdc, 0x98, 0xf8, 0x08, 0x06,
	0x5f, 0xc4, 0x74, 0x5a, 0x7e, 0x4a, 0x81, 0xb8, 0xea, 0xdc, 0x3b, 0xe7, 0x9e, 0x39, 0xe7, 0xdc,
	0x14, 0xca, 0x63, 0x37, 0x60, 0xc2, 0xa5, 0xbc, 0xe9, 0x0b, 0x2f, 0xf0, 0xb0, 0x36, 0xaf, 0xc9,
	0x2b, 0x82, 0xca, 0x9d, 0xf0, 0x86, 0x82, 0x4e, 0xda, 0x8e, 0x23, 0x98, 0x94, 0x36, 0x9b, 0xe2,
	0x7d, 0x28, 0x8c, 0xa8, 0xeb, 0x70, 0xa6, 0xa3, 0x3a, 0xb2, 0x8a, 0x76, 0x5c, 0xe1, 0x1a, 0x14,
	0x7d, 0xe1, 0x0d, 0xfb, 0x2e, 0x9d, 0x30, 0xfd, 0x8f, 0xba, 0xd2, 0xc2, 0x46, 0x8f, 0x4e, 0x18,
	0x6e, 0x40, 0x49, 0x30, 0xe9, 0xf1, 0x67, 0xd6, 0xe7, 0x74, 0xc0, 0xb8, 0x9e, 0xad, 0x23, 0x4b,
	0xb3, 0xff, 0xc5, 0xcd, 0x6e, 0xd8, 0xc3, 0x7b, 0x90, 0x8f, 0x2e, 0x73, 0x6a, 0x3a, 0x2a, 0x88,
	0x05, 0x78, 0x5d, 0x84, 0xf4, 0x31, 0x86, 0x1c, 0x75, 0x1c, 0xa1, 0x34, 0x64, 0x6d, 0x75, 0x26,
	0x3d, 0xa8, 0x5d, 0xb1, 0x20, 0xa5, 0xf8, 0x89, 0xc9, 0xe0, 0x46, 0x0e, 0xf1, 0x19, 0xe4, 0x04,
	0x9b, 0x4a, 0x1d, 0xd5, 0xb3, 0xd6, 0xdf, 0x56, 0xad, 0xb9, 0xf0, 0x9d, 0x9a, 0xb0, 0x15, 0x90,
	0xdc, 0x42, 0x35, 0xfd, 0xb2, 0xe7, 0x4a, 0x16, 0xb2, 0xb5, 0x20, 0x2f, 0x98, 0xf4, 0xe7, 0x74,
	0x87, 0xdb, 0xe9, 0xa4, 0x6f, 0x47, 0x50, 0x72, 0x02, 0xa5, 0xce, 0xe0, 0x17, 0x59, 0x92, 0x0b,
	0x28, 0xaf, 0x02, 0x37, 0xfb, 0x0d, 0x7b, 0x72, 0xfc, 0x12, 0x85, 0x5d, 0xb2, 0xd5, 0x99, 0x74,
	0xa0, 0xda, 0xe6, 0xdc, 0x7b, 0x4c, 0xbc, 0x33, 0x4f, 0xe0, 0x34, 0x91, 0xc0, 0xc1, 0x52, 0x72,
	0x02, 0x1d, 0xbb, 0xef, 0x82, 0x91, 0x62, 0x5a, 0xda, 0x6f, 0x26, 0xed, 0xeb, 0x9b, 0xb9, 0x16,
	0xd6, 0x5b, 0x1f, 0x08, 0xb4, 0xeb, 0x18, 0x82, 0x29, 0x54, 0x52, 0x8b, 0xc2, 0xc7, 0x4b, 0x8a,
	0x1d, 0x5b, 0x34, 0x1a, 0xbb, 0x82, 0x8e, 0xd5, 0x91, 0x0c, 0x7e, 0x80, 0xff, 0x6b, 0xea, 0xf1,
	0xca, 0xe4, 0xd6, 0x88, 0x8c, 0xa3, 0x1d, 0xa0, 0x15, 0xfe, 0x4b, 0xe3, 0x7d, 0x66, 0xa2, 0xcf,
	0x99, 0x89, 0xbe, 0x66, 0x26, 0x7a, 0xfb, 0x36, 0x33, 0xf7, 0xda, 0x88, 0x72, 0xf5, 0x07, 0x0d,
	0x0a, 0xea, 0x73, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x7a, 0x5e, 0xe4, 0x5a, 0x03, 0x00,
	0x00,
}
