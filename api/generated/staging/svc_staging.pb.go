// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: svc_staging.proto

package staging

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/pensando/sw/api/generated/cluster"
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

// AutoMsgBufferWatchHelper is a wrapper object for watch events for Buffer objects
type AutoMsgBufferWatchHelper struct {
	Events []*AutoMsgBufferWatchHelper_WatchEvent `protobuf:"bytes,1,rep,name=Events,json=events" json:"events"`
}

func (m *AutoMsgBufferWatchHelper) Reset()         { *m = AutoMsgBufferWatchHelper{} }
func (m *AutoMsgBufferWatchHelper) String() string { return proto.CompactTextString(m) }
func (*AutoMsgBufferWatchHelper) ProtoMessage()    {}
func (*AutoMsgBufferWatchHelper) Descriptor() ([]byte, []int) {
	return fileDescriptorSvcStaging, []int{0}
}

func (m *AutoMsgBufferWatchHelper) GetEvents() []*AutoMsgBufferWatchHelper_WatchEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type AutoMsgBufferWatchHelper_WatchEvent struct {
	Type   string  `protobuf:"bytes,1,opt,name=Type,proto3" json:"type,omitempty"`
	Object *Buffer `protobuf:"bytes,2,opt,name=Object" json:"object,omitempty"`
}

func (m *AutoMsgBufferWatchHelper_WatchEvent) Reset()         { *m = AutoMsgBufferWatchHelper_WatchEvent{} }
func (m *AutoMsgBufferWatchHelper_WatchEvent) String() string { return proto.CompactTextString(m) }
func (*AutoMsgBufferWatchHelper_WatchEvent) ProtoMessage()    {}
func (*AutoMsgBufferWatchHelper_WatchEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorSvcStaging, []int{0, 0}
}

func (m *AutoMsgBufferWatchHelper_WatchEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AutoMsgBufferWatchHelper_WatchEvent) GetObject() *Buffer {
	if m != nil {
		return m.Object
	}
	return nil
}

// BufferList is a container object for list of Buffer objects
type BufferList struct {
	api.TypeMeta `protobuf:"bytes,2,opt,name=T,json=,inline,embedded=T" json:",inline"`
	api.ListMeta `protobuf:"bytes,3,opt,name=ListMeta,json=list-meta,inline,embedded=ListMeta" json:"list-meta,inline"`
	// List of Buffer objects
	Items []*Buffer `protobuf:"bytes,4,rep,name=Items,json=items" json:"items"`
}

func (m *BufferList) Reset()                    { *m = BufferList{} }
func (m *BufferList) String() string            { return proto.CompactTextString(m) }
func (*BufferList) ProtoMessage()               {}
func (*BufferList) Descriptor() ([]byte, []int) { return fileDescriptorSvcStaging, []int{1} }

func (m *BufferList) GetItems() []*Buffer {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*AutoMsgBufferWatchHelper)(nil), "staging.AutoMsgBufferWatchHelper")
	proto.RegisterType((*AutoMsgBufferWatchHelper_WatchEvent)(nil), "staging.AutoMsgBufferWatchHelper.WatchEvent")
	proto.RegisterType((*BufferList)(nil), "staging.BufferList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StagingV1 service

type StagingV1Client interface {
	// Create Buffer object
	AutoAddBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error)
	// Delete Buffer object
	AutoDeleteBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error)
	// Get Buffer object
	AutoGetBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error)
	// Label Buffer object
	AutoLabelBuffer(ctx context.Context, in *api.Label, opts ...grpc.CallOption) (*Buffer, error)
	// List Buffer objects
	AutoListBuffer(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (*BufferList, error)
	// Update Buffer object
	AutoUpdateBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error)
	// Watch Buffer objects. Supports WebSockets or HTTP long poll
	AutoWatchBuffer(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (StagingV1_AutoWatchBufferClient, error)
	AutoWatchSvcStagingV1(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (StagingV1_AutoWatchSvcStagingV1Client, error)
	// Create/Update/Delete multiple objects as part of a single request
	Bulkedit(ctx context.Context, in *BulkEditAction, opts ...grpc.CallOption) (*BulkEditAction, error)
	// Clear operations from a configuration buffer
	Clear(ctx context.Context, in *ClearAction, opts ...grpc.CallOption) (*ClearAction, error)
	// Commit a staged configuration buffer
	Commit(ctx context.Context, in *CommitAction, opts ...grpc.CallOption) (*CommitAction, error)
}

type stagingV1Client struct {
	cc *grpc.ClientConn
}

func NewStagingV1Client(cc *grpc.ClientConn) StagingV1Client {
	return &stagingV1Client{cc}
}

func (c *stagingV1Client) AutoAddBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error) {
	out := new(Buffer)
	err := grpc.Invoke(ctx, "/staging.StagingV1/AutoAddBuffer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) AutoDeleteBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error) {
	out := new(Buffer)
	err := grpc.Invoke(ctx, "/staging.StagingV1/AutoDeleteBuffer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) AutoGetBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error) {
	out := new(Buffer)
	err := grpc.Invoke(ctx, "/staging.StagingV1/AutoGetBuffer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) AutoLabelBuffer(ctx context.Context, in *api.Label, opts ...grpc.CallOption) (*Buffer, error) {
	out := new(Buffer)
	err := grpc.Invoke(ctx, "/staging.StagingV1/AutoLabelBuffer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) AutoListBuffer(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (*BufferList, error) {
	out := new(BufferList)
	err := grpc.Invoke(ctx, "/staging.StagingV1/AutoListBuffer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) AutoUpdateBuffer(ctx context.Context, in *Buffer, opts ...grpc.CallOption) (*Buffer, error) {
	out := new(Buffer)
	err := grpc.Invoke(ctx, "/staging.StagingV1/AutoUpdateBuffer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) AutoWatchBuffer(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (StagingV1_AutoWatchBufferClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StagingV1_serviceDesc.Streams[0], c.cc, "/staging.StagingV1/AutoWatchBuffer", opts...)
	if err != nil {
		return nil, err
	}
	x := &stagingV1AutoWatchBufferClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StagingV1_AutoWatchBufferClient interface {
	Recv() (*AutoMsgBufferWatchHelper, error)
	grpc.ClientStream
}

type stagingV1AutoWatchBufferClient struct {
	grpc.ClientStream
}

func (x *stagingV1AutoWatchBufferClient) Recv() (*AutoMsgBufferWatchHelper, error) {
	m := new(AutoMsgBufferWatchHelper)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stagingV1Client) AutoWatchSvcStagingV1(ctx context.Context, in *api.ListWatchOptions, opts ...grpc.CallOption) (StagingV1_AutoWatchSvcStagingV1Client, error) {
	stream, err := grpc.NewClientStream(ctx, &_StagingV1_serviceDesc.Streams[1], c.cc, "/staging.StagingV1/AutoWatchSvcStagingV1", opts...)
	if err != nil {
		return nil, err
	}
	x := &stagingV1AutoWatchSvcStagingV1Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StagingV1_AutoWatchSvcStagingV1Client interface {
	Recv() (*api.WatchEventList, error)
	grpc.ClientStream
}

type stagingV1AutoWatchSvcStagingV1Client struct {
	grpc.ClientStream
}

func (x *stagingV1AutoWatchSvcStagingV1Client) Recv() (*api.WatchEventList, error) {
	m := new(api.WatchEventList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stagingV1Client) Bulkedit(ctx context.Context, in *BulkEditAction, opts ...grpc.CallOption) (*BulkEditAction, error) {
	out := new(BulkEditAction)
	err := grpc.Invoke(ctx, "/staging.StagingV1/Bulkedit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) Clear(ctx context.Context, in *ClearAction, opts ...grpc.CallOption) (*ClearAction, error) {
	out := new(ClearAction)
	err := grpc.Invoke(ctx, "/staging.StagingV1/Clear", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stagingV1Client) Commit(ctx context.Context, in *CommitAction, opts ...grpc.CallOption) (*CommitAction, error) {
	out := new(CommitAction)
	err := grpc.Invoke(ctx, "/staging.StagingV1/Commit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StagingV1 service

type StagingV1Server interface {
	// Create Buffer object
	AutoAddBuffer(context.Context, *Buffer) (*Buffer, error)
	// Delete Buffer object
	AutoDeleteBuffer(context.Context, *Buffer) (*Buffer, error)
	// Get Buffer object
	AutoGetBuffer(context.Context, *Buffer) (*Buffer, error)
	// Label Buffer object
	AutoLabelBuffer(context.Context, *api.Label) (*Buffer, error)
	// List Buffer objects
	AutoListBuffer(context.Context, *api.ListWatchOptions) (*BufferList, error)
	// Update Buffer object
	AutoUpdateBuffer(context.Context, *Buffer) (*Buffer, error)
	// Watch Buffer objects. Supports WebSockets or HTTP long poll
	AutoWatchBuffer(*api.ListWatchOptions, StagingV1_AutoWatchBufferServer) error
	AutoWatchSvcStagingV1(*api.ListWatchOptions, StagingV1_AutoWatchSvcStagingV1Server) error
	// Create/Update/Delete multiple objects as part of a single request
	Bulkedit(context.Context, *BulkEditAction) (*BulkEditAction, error)
	// Clear operations from a configuration buffer
	Clear(context.Context, *ClearAction) (*ClearAction, error)
	// Commit a staged configuration buffer
	Commit(context.Context, *CommitAction) (*CommitAction, error)
}

func RegisterStagingV1Server(s *grpc.Server, srv StagingV1Server) {
	s.RegisterService(&_StagingV1_serviceDesc, srv)
}

func _StagingV1_AutoAddBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Buffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).AutoAddBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/AutoAddBuffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).AutoAddBuffer(ctx, req.(*Buffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_AutoDeleteBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Buffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).AutoDeleteBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/AutoDeleteBuffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).AutoDeleteBuffer(ctx, req.(*Buffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_AutoGetBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Buffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).AutoGetBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/AutoGetBuffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).AutoGetBuffer(ctx, req.(*Buffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_AutoLabelBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.Label)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).AutoLabelBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/AutoLabelBuffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).AutoLabelBuffer(ctx, req.(*api.Label))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_AutoListBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ListWatchOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).AutoListBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/AutoListBuffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).AutoListBuffer(ctx, req.(*api.ListWatchOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_AutoUpdateBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Buffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).AutoUpdateBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/AutoUpdateBuffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).AutoUpdateBuffer(ctx, req.(*Buffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_AutoWatchBuffer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.ListWatchOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StagingV1Server).AutoWatchBuffer(m, &stagingV1AutoWatchBufferServer{stream})
}

type StagingV1_AutoWatchBufferServer interface {
	Send(*AutoMsgBufferWatchHelper) error
	grpc.ServerStream
}

type stagingV1AutoWatchBufferServer struct {
	grpc.ServerStream
}

func (x *stagingV1AutoWatchBufferServer) Send(m *AutoMsgBufferWatchHelper) error {
	return x.ServerStream.SendMsg(m)
}

func _StagingV1_AutoWatchSvcStagingV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(api.ListWatchOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StagingV1Server).AutoWatchSvcStagingV1(m, &stagingV1AutoWatchSvcStagingV1Server{stream})
}

type StagingV1_AutoWatchSvcStagingV1Server interface {
	Send(*api.WatchEventList) error
	grpc.ServerStream
}

type stagingV1AutoWatchSvcStagingV1Server struct {
	grpc.ServerStream
}

func (x *stagingV1AutoWatchSvcStagingV1Server) Send(m *api.WatchEventList) error {
	return x.ServerStream.SendMsg(m)
}

func _StagingV1_Bulkedit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkEditAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).Bulkedit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/Bulkedit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).Bulkedit(ctx, req.(*BulkEditAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).Clear(ctx, req.(*ClearAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _StagingV1_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StagingV1Server).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staging.StagingV1/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StagingV1Server).Commit(ctx, req.(*CommitAction))
	}
	return interceptor(ctx, in, info, handler)
}

var _StagingV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "staging.StagingV1",
	HandlerType: (*StagingV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AutoAddBuffer",
			Handler:    _StagingV1_AutoAddBuffer_Handler,
		},
		{
			MethodName: "AutoDeleteBuffer",
			Handler:    _StagingV1_AutoDeleteBuffer_Handler,
		},
		{
			MethodName: "AutoGetBuffer",
			Handler:    _StagingV1_AutoGetBuffer_Handler,
		},
		{
			MethodName: "AutoLabelBuffer",
			Handler:    _StagingV1_AutoLabelBuffer_Handler,
		},
		{
			MethodName: "AutoListBuffer",
			Handler:    _StagingV1_AutoListBuffer_Handler,
		},
		{
			MethodName: "AutoUpdateBuffer",
			Handler:    _StagingV1_AutoUpdateBuffer_Handler,
		},
		{
			MethodName: "Bulkedit",
			Handler:    _StagingV1_Bulkedit_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _StagingV1_Clear_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _StagingV1_Commit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AutoWatchBuffer",
			Handler:       _StagingV1_AutoWatchBuffer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AutoWatchSvcStagingV1",
			Handler:       _StagingV1_AutoWatchSvcStagingV1_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "svc_staging.proto",
}

func (m *AutoMsgBufferWatchHelper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AutoMsgBufferWatchHelper) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSvcStaging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *AutoMsgBufferWatchHelper_WatchEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AutoMsgBufferWatchHelper_WatchEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSvcStaging(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Object != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSvcStaging(dAtA, i, uint64(m.Object.Size()))
		n1, err := m.Object.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *BufferList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BufferList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x12
	i++
	i = encodeVarintSvcStaging(dAtA, i, uint64(m.TypeMeta.Size()))
	n2, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSvcStaging(dAtA, i, uint64(m.ListMeta.Size()))
	n3, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x22
			i++
			i = encodeVarintSvcStaging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintSvcStaging(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AutoMsgBufferWatchHelper) Size() (n int) {
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovSvcStaging(uint64(l))
		}
	}
	return n
}

func (m *AutoMsgBufferWatchHelper_WatchEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSvcStaging(uint64(l))
	}
	if m.Object != nil {
		l = m.Object.Size()
		n += 1 + l + sovSvcStaging(uint64(l))
	}
	return n
}

func (m *BufferList) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovSvcStaging(uint64(l))
	l = m.ListMeta.Size()
	n += 1 + l + sovSvcStaging(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovSvcStaging(uint64(l))
		}
	}
	return n
}

func sovSvcStaging(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSvcStaging(x uint64) (n int) {
	return sovSvcStaging(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AutoMsgBufferWatchHelper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSvcStaging
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
			return fmt.Errorf("proto: AutoMsgBufferWatchHelper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AutoMsgBufferWatchHelper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcStaging
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
				return ErrInvalidLengthSvcStaging
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &AutoMsgBufferWatchHelper_WatchEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSvcStaging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSvcStaging
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
func (m *AutoMsgBufferWatchHelper_WatchEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSvcStaging
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
			return fmt.Errorf("proto: WatchEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcStaging
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
				return ErrInvalidLengthSvcStaging
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcStaging
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
				return ErrInvalidLengthSvcStaging
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Object == nil {
				m.Object = &Buffer{}
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSvcStaging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSvcStaging
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
func (m *BufferList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSvcStaging
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
			return fmt.Errorf("proto: BufferList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BufferList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcStaging
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
				return ErrInvalidLengthSvcStaging
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcStaging
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
				return ErrInvalidLengthSvcStaging
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSvcStaging
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
				return ErrInvalidLengthSvcStaging
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Buffer{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSvcStaging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSvcStaging
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
func skipSvcStaging(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSvcStaging
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
					return 0, ErrIntOverflowSvcStaging
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
					return 0, ErrIntOverflowSvcStaging
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
				return 0, ErrInvalidLengthSvcStaging
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSvcStaging
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
				next, err := skipSvcStaging(dAtA[start:])
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
	ErrInvalidLengthSvcStaging = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSvcStaging   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("svc_staging.proto", fileDescriptorSvcStaging) }

var fileDescriptorSvcStaging = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x3d, 0x49, 0xec, 0x26, 0x93, 0xc6, 0x49, 0xa6, 0x8d, 0xe2, 0x5d, 0x2a, 0x7b, 0xb3,
	0x80, 0x14, 0x2c, 0xd7, 0x9b, 0x1a, 0xc1, 0x21, 0xb7, 0x38, 0x04, 0x0a, 0xb4, 0x24, 0x72, 0x43,
	0x83, 0x72, 0x81, 0xf5, 0x7a, 0xbc, 0x19, 0xba, 0xbf, 0xd8, 0x9d, 0x75, 0x89, 0x50, 0x2e, 0xd8,
	0xbd, 0x70, 0x42, 0x70, 0xe3, 0xc8, 0x05, 0x09, 0x71, 0xca, 0x01, 0xa1, 0x9e, 0x38, 0x56, 0x9c,
	0x2a, 0x71, 0x40, 0xea, 0xc1, 0x42, 0x51, 0x0f, 0xc8, 0x7f, 0x05, 0x9a, 0x99, 0xdd, 0xb5, 0x1d,
	0xc7, 0x89, 0x7b, 0xf1, 0xce, 0x9b, 0x7d, 0xef, 0xfb, 0x3e, 0x6f, 0xde, 0xcc, 0x78, 0xe1, 0x72,
	0xd0, 0x32, 0x3e, 0x0f, 0xa8, 0x6e, 0x12, 0xc7, 0x2c, 0x7b, 0xbe, 0x4b, 0x5d, 0x74, 0x2d, 0x32,
	0xe5, 0x5b, 0xa6, 0xeb, 0x9a, 0x16, 0xd6, 0x74, 0x8f, 0x68, 0xba, 0xe3, 0xb8, 0x54, 0xa7, 0xc4,
	0x75, 0x02, 0xe1, 0x26, 0xef, 0x98, 0x84, 0x1e, 0x85, 0xf5, 0xb2, 0xe1, 0xda, 0x9a, 0x87, 0x9d,
	0x40, 0x77, 0x1a, 0xae, 0x16, 0x3c, 0xd6, 0x5a, 0xd8, 0x21, 0x06, 0xd6, 0x42, 0x4a, 0xac, 0x80,
	0x85, 0x9a, 0xd8, 0x19, 0x8c, 0xd6, 0x88, 0x63, 0x58, 0x61, 0x03, 0xc7, 0x32, 0xb7, 0x07, 0x64,
	0x4c, 0xd7, 0x74, 0x35, 0x3e, 0x5d, 0x0f, 0x9b, 0xdc, 0xe2, 0x06, 0x1f, 0x45, 0xee, 0xd7, 0x29,
	0x76, 0x74, 0x87, 0x46, 0xd6, 0xc2, 0x10, 0xb9, 0xfc, 0xe6, 0x18, 0x24, 0x56, 0x80, 0x8d, 0xa9,
	0x2e, 0xdc, 0xd4, 0x97, 0x00, 0xe6, 0xb6, 0x42, 0xea, 0xde, 0x0f, 0xcc, 0x6a, 0xd8, 0x6c, 0x62,
	0xff, 0x40, 0xa7, 0xc6, 0xd1, 0x5d, 0x6c, 0x79, 0xd8, 0x47, 0x7b, 0x30, 0xb3, 0xd3, 0xc2, 0x0e,
	0x0d, 0x72, 0x40, 0x99, 0x5e, 0x9f, 0xaf, 0x94, 0xca, 0x71, 0x8e, 0x71, 0x21, 0x65, 0x3e, 0xe6,
	0x41, 0x55, 0xd8, 0xeb, 0x16, 0x32, 0x98, 0xc7, 0xd7, 0xa2, 0xa7, 0x8c, 0x21, 0xec, 0x7b, 0x20,
	0x05, 0xce, 0xec, 0x1f, 0x7b, 0x38, 0x07, 0x14, 0xb0, 0x3e, 0x57, 0x45, 0xbd, 0x6e, 0x21, 0x4b,
	0x8f, 0x3d, 0x5c, 0x72, 0x6d, 0x42, 0xb1, 0xed, 0xd1, 0x63, 0xf4, 0x0e, 0xcc, 0xec, 0xd6, 0xbf,
	0xc4, 0x06, 0xcd, 0x4d, 0x29, 0x60, 0x7d, 0xbe, 0xb2, 0x98, 0x10, 0x88, 0xd4, 0xd5, 0x9b, 0xbd,
	0x6e, 0x61, 0xc9, 0xe5, 0x2e, 0xfd, 0xb0, 0xcd, 0xc5, 0x17, 0x4f, 0xa4, 0xf9, 0xc7, 0x2c, 0xd1,
	0x11, 0xc7, 0x52, 0xff, 0x01, 0x10, 0x8a, 0x88, 0x7b, 0x24, 0xa0, 0xe8, 0x5d, 0x08, 0xf6, 0x23,
	0xc5, 0x85, 0xb2, 0xee, 0x91, 0x32, 0xc3, 0xb8, 0x8f, 0xa9, 0x5e, 0xbd, 0xf1, 0xac, 0x5b, 0x48,
	0x3d, 0xef, 0x16, 0x40, 0xaf, 0x5b, 0xb8, 0x56, 0x22, 0x8e, 0x45, 0x1c, 0x5c, 0x8b, 0x07, 0x68,
	0x17, 0xce, 0xb2, 0x78, 0xe6, 0x99, 0x9b, 0x1e, 0x08, 0x8f, 0x27, 0xab, 0xb7, 0x06, 0xc2, 0x97,
	0x2c, 0x12, 0xd0, 0xdb, 0x6c, 0xad, 0x63, 0x9d, 0x91, 0x19, 0xb4, 0x01, 0xd3, 0x1f, 0x52, 0x6c,
	0x07, 0xb9, 0x19, 0xbe, 0xc0, 0x23, 0xe5, 0xcd, 0xf5, 0xba, 0x85, 0x34, 0xab, 0x2a, 0xa8, 0x89,
	0xc7, 0x66, 0xf6, 0xc5, 0x13, 0x09, 0x32, 0x1d, 0x51, 0x59, 0xa5, 0x9d, 0x85, 0x73, 0x0f, 0x44,
	0xd0, 0xc3, 0x3b, 0xe8, 0x04, 0x2e, 0xb0, 0xd6, 0x6c, 0x35, 0x1a, 0x42, 0x00, 0x9d, 0x57, 0x94,
	0xcf, 0x4f, 0xa8, 0x1f, 0x9d, 0x76, 0xa4, 0x8c, 0xe1, 0x63, 0x9d, 0xe2, 0x3f, 0x3a, 0x12, 0xf8,
	0xb3, 0x23, 0x81, 0x6f, 0xff, 0x7e, 0xf9, 0xe3, 0x54, 0x05, 0xa6, 0x36, 0x41, 0xf1, 0x70, 0x91,
	0x3f, 0xd4, 0x59, 0xad, 0xce, 0xdd, 0x03, 0x55, 0xd6, 0xc4, 0xb6, 0xd3, 0xbe, 0xd9, 0x2d, 0xef,
	0xf3, 0xd1, 0x49, 0xfc, 0x0e, 0x7d, 0x0f, 0xe0, 0x12, 0xcb, 0xff, 0x1e, 0xb6, 0x30, 0xc5, 0x13,
	0x23, 0x1c, 0x32, 0x84, 0x06, 0x8f, 0x19, 0x42, 0xa8, 0xc2, 0xd4, 0x66, 0xea, 0x70, 0x95, 0xfd,
	0x16, 0x97, 0xe3, 0x24, 0x2c, 0xef, 0x27, 0xba, 0x8d, 0x4f, 0x8a, 0xaf, 0x8f, 0x27, 0x49, 0x9c,
	0xd0, 0x77, 0x40, 0x2c, 0xc9, 0x07, 0x98, 0x4e, 0xcc, 0x73, 0x70, 0xda, 0x91, 0xa6, 0x4d, 0x4c,
	0xc7, 0xc1, 0xa0, 0x51, 0x18, 0x34, 0x11, 0xcc, 0xc7, 0x70, 0x91, 0xb1, 0xdc, 0xd3, 0xeb, 0xd8,
	0x8a, 0x68, 0xa0, 0xd8, 0x40, 0x6c, 0x66, 0x14, 0xe4, 0xb5, 0xd3, 0x8e, 0x94, 0xb6, 0xd8, 0xbb,
	0x08, 0x25, 0xf5, 0x57, 0x47, 0xca, 0x44, 0x91, 0x1d, 0x00, 0xb3, 0x5c, 0x8d, 0x04, 0x71, 0x69,
	0x2b, 0xc9, 0x6e, 0xe4, 0xa7, 0x6c, 0xd7, 0xe3, 0xd7, 0x8c, 0x7c, 0xe3, 0x9c, 0x2e, 0x73, 0x50,
	0xdf, 0x3f, 0xed, 0x48, 0x33, 0x6c, 0x23, 0x0d, 0x55, 0xb9, 0xc1, 0xab, 0xcc, 0xf2, 0x2a, 0x93,
	0x9e, 0xa3, 0xcb, 0x7a, 0x7e, 0x57, 0xb4, 0xfc, 0x53, 0xaf, 0xa1, 0xbf, 0x42, 0xcb, 0x11, 0x6b,
	0x79, 0xc8, 0x63, 0xe2, 0xd2, 0xd0, 0x17, 0x62, 0x75, 0x38, 0xfa, 0xe5, 0x05, 0xad, 0x5d, 0x79,
	0x11, 0xa9, 0xcb, 0x6c, 0xe9, 0xf8, 0x15, 0x10, 0xeb, 0x6f, 0x00, 0xf4, 0x19, 0x5c, 0x49, 0x32,
	0x3c, 0x68, 0x19, 0xfd, 0x73, 0x33, 0x76, 0xe1, 0xd8, 0x74, 0xff, 0xc6, 0xe2, 0x0b, 0x77, 0xa1,
	0xf2, 0xef, 0x00, 0xce, 0x56, 0x43, 0xeb, 0x11, 0x6e, 0x10, 0x8a, 0x56, 0x07, 0xaa, 0xb5, 0x1e,
	0xed, 0x34, 0x08, 0xdd, 0x32, 0x98, 0xa0, 0x3c, 0xee, 0x85, 0xfa, 0xf5, 0xe8, 0x21, 0xec, 0x77,
	0x9a, 0x37, 0xe6, 0xa1, 0x38, 0x8e, 0x6b, 0xe2, 0x38, 0xca, 0x23, 0x7b, 0x4b, 0xab, 0x47, 0xd9,
	0xd5, 0xd2, 0x04, 0x3b, 0x31, 0xf1, 0x46, 0xbf, 0x00, 0x98, 0xde, 0xb6, 0xb0, 0xee, 0xa3, 0x9b,
	0x09, 0x1c, 0xb7, 0x23, 0xe4, 0x0b, 0x67, 0x55, 0xff, 0x4a, 0xde, 0x3d, 0xc1, 0x9b, 0x17, 0xbc,
	0xab, 0xa3, 0x04, 0x06, 0x93, 0x53, 0xdf, 0x9a, 0x04, 0x96, 0xbb, 0xa2, 0xdf, 0x00, 0xcc, 0x6c,
	0xbb, 0xb6, 0x4d, 0x28, 0x5a, 0xe9, 0x43, 0xf1, 0x89, 0x88, 0xf5, 0xe2, 0x69, 0x95, 0x5e, 0x09,
	0x5b, 0x13, 0xb0, 0x05, 0x01, 0x9b, 0xbb, 0x80, 0x80, 0xeb, 0xa9, 0xc5, 0x89, 0x68, 0xb9, 0xaf,
	0xdc, 0x9b, 0xfa, 0xa1, 0x2d, 0x4d, 0xb5, 0xee, 0xfc, 0xd4, 0x96, 0xe2, 0xaf, 0x88, 0x9f, 0xdb,
	0x71, 0xd6, 0x5f, 0xdb, 0xd2, 0x1a, 0x8c, 0x0f, 0xf2, 0x8c, 0xe7, 0x06, 0x14, 0xf1, 0xb3, 0x88,
	0xd8, 0xb5, 0x83, 0xa2, 0xbb, 0xf0, 0x69, 0x5b, 0xaa, 0xa1, 0xc8, 0x49, 0xce, 0x44, 0x08, 0xd7,
	0x07, 0x4b, 0x2b, 0x0e, 0x59, 0x95, 0x37, 0x84, 0xa5, 0xe8, 0x0a, 0x4b, 0x88, 0x1b, 0x8a, 0xe1,
	0x3a, 0x4d, 0x62, 0x86, 0x3e, 0xff, 0xdc, 0x50, 0x04, 0xe9, 0xd3, 0xb6, 0x74, 0x90, 0xa8, 0xa6,
	0x45, 0x17, 0xe6, 0x07, 0x7a, 0x5b, 0x1c, 0x34, 0x2a, 0x25, 0x6e, 0x28, 0xae, 0x87, 0x85, 0x48,
	0xa0, 0x34, 0x7d, 0xd7, 0x56, 0xf4, 0x71, 0xd2, 0x61, 0x22, 0x3d, 0x9b, 0x6c, 0xc8, 0xec, 0xf0,
	0x66, 0x2f, 0x9e, 0xb3, 0x2b, 0x5b, 0xdb, 0xbc, 0x37, 0x9a, 0xb8, 0x42, 0x34, 0xf1, 0xe7, 0xa1,
	0xd8, 0xa1, 0x45, 0x89, 0x67, 0x61, 0x45, 0xfc, 0xbb, 0x07, 0x8a, 0x1e, 0x28, 0x9e, 0xee, 0x53,
	0xc5, 0x6d, 0xb2, 0x02, 0x89, 0x63, 0x5a, 0x58, 0xf1, 0xf1, 0x57, 0x21, 0x0e, 0x68, 0x75, 0xe9,
	0xd9, 0x59, 0x1e, 0x3c, 0x3f, 0xcb, 0x83, 0x7f, 0xcf, 0xf2, 0xe0, 0xbf, 0xb3, 0x7c, 0x6a, 0x0f,
	0xd4, 0x33, 0xfc, 0x0b, 0xe7, 0xed, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x8a, 0xbe, 0xf8,
	0xd7, 0x09, 0x00, 0x00,
}
