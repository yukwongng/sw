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
	AutoWatchSvcEventsV1(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (EventsV1_AutoWatchSvcEventsV1Client, error)
	// http://<...>/events/v1/events/12345 will be translated to a gRPC query - GetEvent(uuid:"12345")
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error)
	// http://<...>/events/v1/events
	GetEvents(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (*EventList, error)
}

type eventsV1Client struct {
	cc *grpc.ClientConn
}

func NewEventsV1Client(cc *grpc.ClientConn) EventsV1Client {
	return &eventsV1Client{cc}
}

func (c *eventsV1Client) AutoWatchSvcEventsV1(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (EventsV1_AutoWatchSvcEventsV1Client, error) {
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
	AutoWatchSvcEventsV1(*api.ListWatchOptions, EventsV1_AutoWatchSvcEventsV1Server) error
	// http://<...>/events/v1/events/12345 will be translated to a gRPC query - GetEvent(uuid:"12345")
	GetEvent(context.Context, *GetEventRequest) (*Event, error)
	// http://<...>/events/v1/events
	GetEvents(context.Context, *api.ListWatchOptions) (*EventList, error)
}

func RegisterEventsV1Server(s *grpc.Server, srv EventsV1Server) {
	s.RegisterService(&_EventsV1_serviceDesc, srv)
}

func _EventsV1_AutoWatchSvcEventsV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.ListWatchOptions)
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
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0xcb, 0xd3, 0x30,
	0x1c, 0xc6, 0xdf, 0x0c, 0x9d, 0xef, 0xca, 0x9c, 0x5b, 0x9c, 0x3f, 0x5a, 0x64, 0x4a, 0x41, 0x10,
	0x71, 0xcd, 0xa6, 0x7f, 0x81, 0xd3, 0x21, 0x8a, 0xa0, 0x4e, 0xe6, 0xc0, 0x8b, 0x64, 0x5d, 0xec,
	0x02, 0x6d, 0x12, 0xcd, 0x37, 0x1d, 0x43, 0xbc, 0xd4, 0x9d, 0x3c, 0xea, 0xcd, 0x93, 0x67, 0x8f,
	0x9e, 0x3c, 0x7a, 0xf4, 0x28, 0x78, 0x17, 0x19, 0x9e, 0xfc, 0x2b, 0xa4, 0x59, 0x2b, 0x73, 0xb0,
	0xf7, 0x96, 0xe7, 0xc9, 0xe7, 0xfb, 0xf4, 0x49, 0x1a, 0xa7, 0xa9, 0xd3, 0xf0, 0x19, 0x4b, 0x99,
	0x00, 0x1d, 0xa8, 0x97, 0x12, 0x24, 0xae, 0x6e, 0x94, 0x77, 0x21, 0x92, 0x32, 0x8a, 0x19, 0xa1,
	0x8a, 0x13, 0x2a, 0x84, 0x04, 0x0a, 0x5c, 0x8a, 0x82, 0xf2, 0x86, 0x11, 0x87, 0xb9, 0x99, 0x06,
	0xa1, 0x4c, 0x88, 0x62, 0x42, 0x53, 0x31, 0x93, 0x44, 0x2f, 0x48, 0xca, 0x04, 0x0f, 0x19, 0x31,
	0xc0, 0x63, 0x9d, 0x8f, 0x46, 0x4c, 0x6c, 0x4f, 0x13, 0x2e, 0xc2, 0xd8, 0xcc, 0x58, 0x19, 0xd3,
	0xdd, 0x8a, 0x89, 0x64, 0x24, 0x89, 0xb5, 0xa7, 0xe6, 0xb9, 0x55, 0x56, 0xd8, 0x55, 0x81, 0x5f,
	0xde, 0xf3, 0xd5, 0xbc, 0x63, 0xc2, 0x80, 0x16, 0x58, 0x70, 0x04, 0x66, 0x09, 0x4d, 0x80, 0x09,
	0x2a, 0xa0, 0xe0, 0xeb, 0xdb, 0x17, 0xe0, 0xdf, 0x72, 0x4e, 0xdd, 0x61, 0x30, 0xcc, 0xad, 0x11,
	0x7b, 0x61, 0x98, 0x06, 0xdc, 0x73, 0x8e, 0x8d, 0xc7, 0x77, 0x6f, 0x9f, 0x47, 0x97, 0xd0, 0x95,
	0xda, 0x00, 0xff, 0xf9, 0x79, 0xb1, 0x61, 0x0c, 0x9f, 0x5d, 0x93, 0x09, 0x07, 0x96, 0x28, 0x58,
	0x8e, 0x76, 0xf4, 0xf5, 0xb7, 0x15, 0xe7, 0xd0, 0x46, 0xe8, 0x27, 0x7d, 0x3c, 0x71, 0xda, 0x37,
	0x0d, 0xc8, 0x09, 0x85, 0x70, 0xfe, 0x38, 0x0d, 0xff, 0xf9, 0x67, 0x02, 0xaa, 0x78, 0x70, 0x9f,
	0x6b, 0xb0, 0x5b, 0x0f, 0x94, 0xbd, 0x23, 0xef, 0xb4, 0xb5, 0xad, 0x65, 0xd1, 0x1c, 0xf0, 0x5b,
	0x9f, 0x57, 0xee, 0xf1, 0x45, 0xee, 0x7d, 0x59, 0xb9, 0xe8, 0xeb, 0xca, 0x3d, 0xe8, 0x21, 0x7c,
	0xcf, 0x39, 0x2c, 0xab, 0xe2, 0x73, 0x41, 0x71, 0x8a, 0x9d, 0xf2, 0xde, 0xc9, 0x72, 0xc3, 0xba,
	0xfe, 0xd9, 0xec, 0xc7, 0xef, 0xf7, 0x95, 0x26, 0x6e, 0x90, 0x8d, 0x4d, 0x5e, 0xe5, 0x47, 0x7b,
	0x8d, 0x1f, 0x39, 0xb5, 0x72, 0x52, 0xef, 0x6b, 0xd6, 0xfa, 0x2f, 0xca, 0xf6, 0xf2, 0x6c, 0x5c,
	0xdb, 0x3f, 0x51, 0xc4, 0x3d, 0xad, 0xe1, 0x72, 0xe9, 0x55, 0xdf, 0xbd, 0x71, 0x2b, 0x69, 0x7f,
	0x70, 0xf5, 0x43, 0xe6, 0xd6, 0x15, 0x13, 0x5d, 0xad, 0x96, 0x51, 0x4c, 0xb5, 0xfe, 0x98, 0xb9,
	0x07, 0x9f, 0x32, 0xb7, 0x78, 0x6a, 0xdf, 0xd6, 0x1d, 0xf4, 0x7d, 0xdd, 0x41, 0xbf, 0xd6, 0x1d,
	0xf4, 0x10, 0x4d, 0xab, 0xf6, 0x37, 0xdc, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x8f, 0x0d,
	0xf0, 0x9b, 0x02, 0x00, 0x00,
}
