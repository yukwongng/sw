// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// events that HAL generates and app(s) can listen to
type EventId int32

const (
	EventId_EVENT_ID_NONE     EventId = 0
	EventId_EVENT_ID_ENDPOINT EventId = 1
	EventId_EVENT_ID_PORT     EventId = 2
)

var EventId_name = map[int32]string{
	0: "EVENT_ID_NONE",
	1: "EVENT_ID_ENDPOINT",
	2: "EVENT_ID_PORT",
}
var EventId_value = map[string]int32{
	"EVENT_ID_NONE":     0,
	"EVENT_ID_ENDPOINT": 1,
	"EVENT_ID_PORT":     2,
}

func (x EventId) String() string {
	return proto.EnumName(EventId_name, int32(x))
}
func (EventId) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type EventOp int32

const (
	EventOp_EVENT_OP_NONE        EventOp = 0
	EventOp_EVENT_OP_SUBSCRIBE   EventOp = 1
	EventOp_EVENT_OP_UNSUBSCRIBE EventOp = 2
)

var EventOp_name = map[int32]string{
	0: "EVENT_OP_NONE",
	1: "EVENT_OP_SUBSCRIBE",
	2: "EVENT_OP_UNSUBSCRIBE",
}
var EventOp_value = map[string]int32{
	"EVENT_OP_NONE":        0,
	"EVENT_OP_SUBSCRIBE":   1,
	"EVENT_OP_UNSUBSCRIBE": 2,
}

func (x EventOp) String() string {
	return proto.EnumName(EventOp_name, int32(x))
}
func (EventOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

// EventSpec captures the event of interest to the app
type EventRequest struct {
	EventId        EventId `protobuf:"varint,1,opt,name=event_id,json=eventId,enum=event.EventId" json:"event_id,omitempty"`
	EventOperation EventOp `protobuf:"varint,2,opt,name=event_operation,json=eventOperation,enum=event.EventOp" json:"event_operation,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *EventRequest) GetEventId() EventId {
	if m != nil {
		return m.EventId
	}
	return EventId_EVENT_ID_NONE
}

func (m *EventRequest) GetEventOperation() EventOp {
	if m != nil {
		return m.EventOperation
	}
	return EventOp_EVENT_OP_NONE
}

type EndpointEvent struct {
	L2SegmentHandle uint64 `protobuf:"fixed64,1,opt,name=l2_segment_handle,json=l2SegmentHandle" json:"l2_segment_handle,omitempty"`
	MacAddress      uint64 `protobuf:"varint,2,opt,name=mac_address,json=macAddress" json:"mac_address,omitempty"`
}

func (m *EndpointEvent) Reset()                    { *m = EndpointEvent{} }
func (m *EndpointEvent) String() string            { return proto.CompactTextString(m) }
func (*EndpointEvent) ProtoMessage()               {}
func (*EndpointEvent) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *EndpointEvent) GetL2SegmentHandle() uint64 {
	if m != nil {
		return m.L2SegmentHandle
	}
	return 0
}

func (m *EndpointEvent) GetMacAddress() uint64 {
	if m != nil {
		return m.MacAddress
	}
	return 0
}

type EventResponse struct {
	ApiStatus ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	EventId   EventId   `protobuf:"varint,2,opt,name=event_id,json=eventId,enum=event.EventId" json:"event_id,omitempty"`
	// Types that are valid to be assigned to EventInfo:
	//	*EventResponse_EpEvent
	EventInfo isEventResponse_EventInfo `protobuf_oneof:"event_info"`
}

func (m *EventResponse) Reset()                    { *m = EventResponse{} }
func (m *EventResponse) String() string            { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()               {}
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type isEventResponse_EventInfo interface{ isEventResponse_EventInfo() }

type EventResponse_EpEvent struct {
	EpEvent *EndpointEvent `protobuf:"bytes,3,opt,name=ep_event,json=epEvent,oneof"`
}

func (*EventResponse_EpEvent) isEventResponse_EventInfo() {}

func (m *EventResponse) GetEventInfo() isEventResponse_EventInfo {
	if m != nil {
		return m.EventInfo
	}
	return nil
}

func (m *EventResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *EventResponse) GetEventId() EventId {
	if m != nil {
		return m.EventId
	}
	return EventId_EVENT_ID_NONE
}

func (m *EventResponse) GetEpEvent() *EndpointEvent {
	if x, ok := m.GetEventInfo().(*EventResponse_EpEvent); ok {
		return x.EpEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EventResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EventResponse_OneofMarshaler, _EventResponse_OneofUnmarshaler, _EventResponse_OneofSizer, []interface{}{
		(*EventResponse_EpEvent)(nil),
	}
}

func _EventResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EventResponse)
	// event_info
	switch x := m.EventInfo.(type) {
	case *EventResponse_EpEvent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EpEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EventResponse.EventInfo has unexpected type %T", x)
	}
	return nil
}

func _EventResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EventResponse)
	switch tag {
	case 3: // event_info.ep_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EndpointEvent)
		err := b.DecodeMessage(msg)
		m.EventInfo = &EventResponse_EpEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EventResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EventResponse)
	// event_info
	switch x := m.EventInfo.(type) {
	case *EventResponse_EpEvent:
		s := proto.Size(x.EpEvent)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*EventRequest)(nil), "halproto.EventRequest")
	proto.RegisterType((*EndpointEvent)(nil), "halproto.EndpointEvent")
	proto.RegisterType((*EventResponse)(nil), "halproto.EventResponse")
	proto.RegisterEnum("halproto.EventId", EventId_name, EventId_value)
	proto.RegisterEnum("halproto.EventOp", EventOp_name, EventOp_value)
}

func init() { proto.RegisterFile("event.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0xaf, 0xd2, 0x30,
	0x1c, 0xc5, 0xe9, 0xf4, 0x5e, 0xf0, 0xbb, 0x0b, 0x77, 0x54, 0x34, 0x0b, 0x2f, 0x92, 0x3d, 0x21,
	0x0f, 0xa8, 0xf3, 0xc1, 0x37, 0x13, 0x26, 0x53, 0x96, 0x98, 0x8d, 0x74, 0xe0, 0x83, 0x31, 0x59,
	0x2a, 0xab, 0xb2, 0x64, 0x74, 0x75, 0x2d, 0x26, 0xfe, 0x4f, 0xfe, 0x91, 0x86, 0x76, 0xfc, 0x8a,
	0xc9, 0x7d, 0xda, 0xe9, 0xa7, 0x67, 0xa7, 0xf9, 0x9e, 0x16, 0x6c, 0xf6, 0x9b, 0x71, 0x35, 0x15,
	0x75, 0xa5, 0x2a, 0x7c, 0xa3, 0x17, 0x43, 0x5b, 0xfd, 0x11, 0x4c, 0x1a, 0xe6, 0xd5, 0x70, 0x17,
	0x1e, 0x28, 0x61, 0xbf, 0xf6, 0x4c, 0x2a, 0xfc, 0x12, 0x3a, 0xda, 0x95, 0x15, 0xb9, 0x8b, 0x46,
	0x68, 0xdc, 0xf3, 0x7b, 0x53, 0x93, 0xa1, 0x6d, 0x51, 0x4e, 0xda, 0xcc, 0x08, 0xfc, 0x0e, 0xee,
	0x8d, 0xb5, 0x12, 0xac, 0xa6, 0xaa, 0xa8, 0xb8, 0x6b, 0xfd, 0xff, 0x47, 0x22, 0x48, 0x8f, 0x19,
	0xd1, 0xb8, 0xbc, 0x6f, 0xd0, 0x0d, 0x79, 0x2e, 0xaa, 0x82, 0x2b, 0x6d, 0xc1, 0x13, 0xe8, 0x97,
	0x7e, 0x26, 0xd9, 0xcf, 0xdd, 0x21, 0x6e, 0x4b, 0x79, 0x5e, 0x32, 0x7d, 0xfa, 0x2d, 0xb9, 0x2f,
	0xfd, 0xd4, 0xf0, 0x85, 0xc6, 0xf8, 0x05, 0xd8, 0x3b, 0xba, 0xc9, 0x68, 0x9e, 0xd7, 0x4c, 0x4a,
	0x7d, 0xe2, 0x63, 0x02, 0x3b, 0xba, 0x99, 0x19, 0xe2, 0xfd, 0x45, 0xd0, 0x6d, 0x46, 0x92, 0xa2,
	0xe2, 0x92, 0xe1, 0x57, 0x00, 0x54, 0x14, 0x99, 0x54, 0x54, 0xed, 0x65, 0x33, 0x95, 0x33, 0x35,
	0x2d, 0xcc, 0x44, 0x91, 0x6a, 0x4e, 0x9e, 0xd0, 0xa3, 0xbc, 0x2a, 0xc1, 0x7a, 0xb8, 0x84, 0x37,
	0xd0, 0x61, 0x22, 0xd3, 0x2b, 0xf7, 0xd1, 0x08, 0x8d, 0x6d, 0x7f, 0x70, 0xb4, 0x5e, 0x8e, 0xb8,
	0x68, 0x91, 0x36, 0x13, 0x5a, 0x06, 0x77, 0x00, 0x4d, 0x3a, 0xff, 0x51, 0x4d, 0x3e, 0x42, 0xbb,
	0x09, 0xc5, 0x7d, 0xe8, 0x86, 0x5f, 0xc2, 0x78, 0x95, 0x45, 0xf3, 0x2c, 0x4e, 0xe2, 0xd0, 0x69,
	0xe1, 0x67, 0xd0, 0x3f, 0xa1, 0x30, 0x9e, 0x2f, 0x93, 0x28, 0x5e, 0x39, 0xe8, 0xca, 0xb9, 0x4c,
	0xc8, 0xca, 0xb1, 0x26, 0x71, 0x93, 0x93, 0x88, 0xf3, 0x6e, 0xb2, 0x3c, 0xe6, 0x3c, 0x07, 0x7c,
	0x42, 0xe9, 0x3a, 0x48, 0x3f, 0x90, 0x28, 0x08, 0x1d, 0x84, 0x5d, 0x18, 0x9c, 0xf8, 0x3a, 0x3e,
	0xef, 0x58, 0xfe, 0x27, 0xb8, 0x31, 0x97, 0xf3, 0x1e, 0x6c, 0x2d, 0x3e, 0x17, 0x52, 0x31, 0x8e,
	0x9f, 0x5e, 0x36, 0xd1, 0xbc, 0x9a, 0xe1, 0xe0, 0x1a, 0x9a, 0xde, 0xbd, 0xd6, 0x18, 0xbd, 0x46,
	0x01, 0x7c, 0xed, 0x6c, 0x69, 0xa9, 0x5f, 0xdb, 0xf7, 0x5b, 0xfd, 0x79, 0xfb, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x90, 0x1f, 0x5a, 0x97, 0x02, 0x00, 0x00,
}
