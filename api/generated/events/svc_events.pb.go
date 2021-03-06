// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: svc_events.proto

package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"
import _ "github.com/pensando/sw/api/generated/cluster"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request for GET /event
type GetEventRequest struct {
	//
	UUID string `protobuf:"bytes,1,opt,name=UUID,json=uuid,omitempty,proto3" json:"uuid,omitempty"`
}

func (m *GetEventRequest) Reset()                    { *m = GetEventRequest{} }
func (m *GetEventRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEventRequest) ProtoMessage()               {}
func (*GetEventRequest) Descriptor() ([]byte, []int) { return fileDescriptorSvcEvents, []int{0} }

func (m *GetEventRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*GetEventRequest)(nil), "events.GetEventRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventsV1 service

type EventsV1Client interface {
	AutoWatchSvcEventsV1(ctx context.Context, in *api.AggWatchOptions, opts ...grpc.CallOption) (EventsV1_AutoWatchSvcEventsV1Client, error)
	// Get specific event
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error)
	//
	GetEvents(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (*EventList, error)
}

type eventsV1Client struct {
	cc *grpc.ClientConn
}

func NewEventsV1Client(cc *grpc.ClientConn) EventsV1Client {
	return &eventsV1Client{cc}
}

func (c *eventsV1Client) AutoWatchSvcEventsV1(ctx context.Context, in *api.AggWatchOptions, opts ...grpc.CallOption) (EventsV1_AutoWatchSvcEventsV1Client, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventsV1_serviceDesc.Streams[0], c.cc, "/events.EventsV1/AutoWatchSvcEventsV1", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsV1AutoWatchSvcEventsV1Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventsV1_AutoWatchSvcEventsV1Client interface {
	Recv() (*api.WatchEventList, error)
	grpc.ClientStream
}

type eventsV1AutoWatchSvcEventsV1Client struct {
	grpc.ClientStream
}

func (x *eventsV1AutoWatchSvcEventsV1Client) Recv() (*api.WatchEventList, error) {
	m := new(api.WatchEventList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventsV1Client) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := grpc.Invoke(ctx, "/events.EventsV1/GetEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsV1Client) GetEvents(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := grpc.Invoke(ctx, "/events.EventsV1/GetEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventsV1 service

type EventsV1Server interface {
	AutoWatchSvcEventsV1(*api.AggWatchOptions, EventsV1_AutoWatchSvcEventsV1Server) error
	// Get specific event
	GetEvent(context.Context, *GetEventRequest) (*Event, error)
	//
	GetEvents(context.Context, *api.ListWatchOptions) (*EventList, error)
}

func RegisterEventsV1Server(s *grpc.Server, srv EventsV1Server) {
	s.RegisterService(&_EventsV1_serviceDesc, srv)
}

func _EventsV1_AutoWatchSvcEventsV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.AggWatchOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsV1Server).AutoWatchSvcEventsV1(m, &eventsV1AutoWatchSvcEventsV1Server{stream})
}

type EventsV1_AutoWatchSvcEventsV1Server interface {
	Send(*api.WatchEventList) error
	grpc.ServerStream
}

type eventsV1AutoWatchSvcEventsV1Server struct {
	grpc.ServerStream
}

func (x *eventsV1AutoWatchSvcEventsV1Server) Send(m *api.WatchEventList) error {
	return x.ServerStream.SendMsg(m)
}

func _EventsV1_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsV1Server).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventsV1/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsV1Server).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsV1_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ListWatchOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsV1Server).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventsV1/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsV1Server).GetEvents(ctx, req.(*api.ListWatchOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventsV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.EventsV1",
	HandlerType: (*EventsV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _EventsV1_GetEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _EventsV1_GetEvents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AutoWatchSvcEventsV1",
			Handler:       _EventsV1_AutoWatchSvcEventsV1_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "svc_events.proto",
}

func (m *GetEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEventRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UUID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSvcEvents(dAtA, i, uint64(len(m.UUID)))
		i += copy(dAtA[i:], m.UUID)
	}
	return i, nil
}

func encodeVarintSvcEvents(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetEventRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovSvcEvents(uint64(l))
	}
	return n
}

func sovSvcEvents(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSvcEvents(x uint64) (n int) {
	return sovSvcEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSvcEvents
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
			return fmt.Errorf("proto: GetEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcEvents
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
				return ErrInvalidLengthSvcEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSvcEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSvcEvents
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
func skipSvcEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSvcEvents
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
					return 0, ErrIntOverflowSvcEvents
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
					return 0, ErrIntOverflowSvcEvents
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
				return 0, ErrInvalidLengthSvcEvents
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSvcEvents
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
				next, err := skipSvcEvents(dAtA[start:])
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
	ErrInvalidLengthSvcEvents = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSvcEvents   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("svc_events.proto", fileDescriptorSvcEvents) }

var fileDescriptorSvcEvents = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x45, 0xeb, 0x36, 0xac, 0xeb, 0xee, 0x58, 0x7f, 0x24, 0x48, 0x94, 0x80, 0xe0,
	0xc1, 0xcd, 0xec, 0xea, 0x5f, 0xb0, 0xab, 0x8b, 0x28, 0x82, 0x5a, 0x69, 0x05, 0x2f, 0x32, 0x4d,
	0xc7, 0xe9, 0x40, 0x32, 0x33, 0x3a, 0x6f, 0x52, 0x8a, 0x78, 0x89, 0xbd, 0x78, 0xd5, 0x9b, 0x27,
	0xcf, 0x1e, 0x3d, 0x79, 0xf4, 0xe8, 0x51, 0xf0, 0x2e, 0x12, 0x3c, 0x88, 0x7f, 0x85, 0x64, 0x9a,
	0x48, 0x5b, 0xd8, 0xdb, 0x7c, 0x3f, 0xf3, 0xde, 0xf7, 0x7d, 0x5f, 0x32, 0xde, 0xb6, 0xc9, 0x93,
	0x67, 0x2c, 0x67, 0x12, 0x4c, 0xac, 0x5f, 0x2a, 0x50, 0xb8, 0xb3, 0x50, 0xc1, 0x25, 0xae, 0x14,
	0x4f, 0x19, 0xa1, 0x5a, 0x10, 0x2a, 0xa5, 0x02, 0x0a, 0x42, 0xc9, 0xba, 0x2a, 0x38, 0xe2, 0x02,
	0x26, 0x76, 0x14, 0x27, 0x2a, 0x23, 0x9a, 0x49, 0x43, 0xe5, 0x58, 0x11, 0x33, 0x25, 0x39, 0x93,
	0x22, 0x61, 0xc4, 0x82, 0x48, 0x4d, 0xd5, 0xca, 0x99, 0x5c, 0xee, 0x26, 0x42, 0x26, 0xa9, 0x1d,
	0xb3, 0xc6, 0x66, 0x77, 0xc9, 0x86, 0x2b, 0xae, 0x88, 0xc3, 0x23, 0xfb, 0xdc, 0x29, 0x27, 0xdc,
	0xa9, 0x2e, 0xbf, 0x7a, 0xcc, 0xd4, 0x2a, 0x63, 0xc6, 0x80, 0xd6, 0x65, 0x9b, 0xc0, 0x24, 0x95,
	0xd0, 0xa8, 0xe5, 0xf5, 0xa2, 0x5b, 0xde, 0x99, 0x3b, 0x0c, 0x8e, 0x2a, 0xd4, 0x67, 0x2f, 0x2c,
	0x33, 0x80, 0xf7, 0xbc, 0x13, 0x83, 0xc1, 0xdd, 0xdb, 0x17, 0xd1, 0x15, 0x74, 0xad, 0x7b, 0x88,
	0xff, 0xfe, 0xbc, 0xbc, 0x65, 0xad, 0x18, 0x5f, 0x57, 0x99, 0x00, 0x96, 0x69, 0x98, 0xf5, 0xd7,
	0xf4, 0x8d, 0xb7, 0x6d, 0x6f, 0xc3, 0x59, 0x98, 0xe1, 0x3e, 0x1e, 0x7a, 0xbd, 0x03, 0x0b, 0xea,
	0x09, 0x85, 0x64, 0xf2, 0x38, 0x4f, 0xfe, 0xf3, 0x5e, 0x4c, 0xb5, 0x88, 0x0f, 0x38, 0x77, 0x37,
	0x0f, 0xb4, 0xfb, 0x00, 0xc1, 0x59, 0x47, 0x1d, 0x72, 0x95, 0xf7, 0x85, 0x81, 0x68, 0xe7, 0xf3,
	0xdc, 0x3f, 0x39, 0xad, 0xd8, 0x97, 0xb9, 0x8f, 0xbe, 0xce, 0xfd, 0xd6, 0x1e, 0xc2, 0xf7, 0xbc,
	0x8d, 0x26, 0x29, 0xbe, 0x10, 0xd7, 0x4b, 0xac, 0x65, 0x0f, 0x4e, 0x37, 0x17, 0x8e, 0x46, 0xe7,
	0x8b, 0x1f, 0xbf, 0xdf, 0xb7, 0xb7, 0xf1, 0x16, 0x59, 0x60, 0xf2, 0xaa, 0xda, 0xec, 0x35, 0x7e,
	0xe4, 0x75, 0x9b, 0x4e, 0x83, 0xcf, 0xb9, 0x08, 0xd5, 0xe0, 0x95, 0x64, 0x3b, 0x2b, 0x56, 0x2e,
	0x57, 0xe0, 0xec, 0x7a, 0xd1, 0xa9, 0xda, 0xee, 0x69, 0x17, 0x37, 0xc7, 0xa0, 0xf3, 0xee, 0x8d,
	0xdf, 0xce, 0xf7, 0x0f, 0xe3, 0x0f, 0x85, 0xbf, 0xa9, 0x99, 0xdc, 0x35, 0x7a, 0xc6, 0x53, 0x6a,
	0xcc, 0xc7, 0xc2, 0x6f, 0x7d, 0x2a, 0xfc, 0xfa, 0x1d, 0x7d, 0x2b, 0x43, 0xf4, 0xbd, 0x0c, 0xd1,
	0xaf, 0x32, 0x44, 0x7f, 0xca, 0xb0, 0xf5, 0x10, 0x8d, 0x3a, 0xee, 0x4f, 0xdc, 0xfc, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x22, 0x11, 0xcf, 0x7c, 0x02, 0x00, 0x00,
}
