// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventpolicy.proto

package monitoring

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"
import fields "github.com/pensando/sw/api/fields"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ------------------------------- Event Policy --------------------------------
type MonitoringExportFormat int32

const (
	// ui-hint: BSD
	MonitoringExportFormat_SYSLOG_BSD MonitoringExportFormat = 0
	// ui-hint: RFC5424
	MonitoringExportFormat_SYSLOG_RFC5424 MonitoringExportFormat = 1
)

var MonitoringExportFormat_name = map[int32]string{
	0: "SYSLOG_BSD",
	1: "SYSLOG_RFC5424",
}
var MonitoringExportFormat_value = map[string]int32{
	"SYSLOG_BSD":     0,
	"SYSLOG_RFC5424": 1,
}

func (MonitoringExportFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEventpolicy, []int{0}
}

// Event Policy represents the policy definition for Events.
// Event Client module will be consumer of this policy.
type EventPolicy struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	//
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	// Spec contains the configuration of an event policy.
	Spec EventPolicySpec `protobuf:"bytes,3,opt,name=Spec,json=spec,omitempty" json:"spec,omitempty"`
	// Status contains the current state of an event policy.
	Status EventPolicyStatus `protobuf:"bytes,4,opt,name=Status,json=status,omitempty" json:"status,omitempty"`
}

func (m *EventPolicy) Reset()                    { *m = EventPolicy{} }
func (m *EventPolicy) String() string            { return proto.CompactTextString(m) }
func (*EventPolicy) ProtoMessage()               {}
func (*EventPolicy) Descriptor() ([]byte, []int) { return fileDescriptorEventpolicy, []int{0} }

func (m *EventPolicy) GetSpec() EventPolicySpec {
	if m != nil {
		return m.Spec
	}
	return EventPolicySpec{}
}

func (m *EventPolicy) GetStatus() EventPolicyStatus {
	if m != nil {
		return m.Status
	}
	return EventPolicyStatus{}
}

// EventPolicySpec is the specification of an Event Policy.
type EventPolicySpec struct {
	// event export format, SYSLOG_BSD default
	Format string `protobuf:"bytes,1,opt,name=Format,json=format, omitempty,proto3" json:"format, omitempty"`
	// export events matched by the selector
	Selector *fields.Selector `protobuf:"bytes,2,opt,name=Selector,json=selector,omitempty" json:"selector,omitempty"`
	// export target ip/port/protocol
	Targets []*ExportConfig `protobuf:"bytes,3,rep,name=Targets,json=targets,omitempty" json:"targets,omitempty"`
	// once we support other formats, it should be one of the supported configs
	// syslog specific configuration
	SyslogConfig *SyslogExportConfig `protobuf:"bytes,4,opt,name=SyslogConfig,json=config,omitempty" json:"config,omitempty"`
}

func (m *EventPolicySpec) Reset()                    { *m = EventPolicySpec{} }
func (m *EventPolicySpec) String() string            { return proto.CompactTextString(m) }
func (*EventPolicySpec) ProtoMessage()               {}
func (*EventPolicySpec) Descriptor() ([]byte, []int) { return fileDescriptorEventpolicy, []int{1} }

func (m *EventPolicySpec) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *EventPolicySpec) GetSelector() *fields.Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *EventPolicySpec) GetTargets() []*ExportConfig {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *EventPolicySpec) GetSyslogConfig() *SyslogExportConfig {
	if m != nil {
		return m.SyslogConfig
	}
	return nil
}

// EventPolicyStatus
type EventPolicyStatus struct {
}

func (m *EventPolicyStatus) Reset()                    { *m = EventPolicyStatus{} }
func (m *EventPolicyStatus) String() string            { return proto.CompactTextString(m) }
func (*EventPolicyStatus) ProtoMessage()               {}
func (*EventPolicyStatus) Descriptor() ([]byte, []int) { return fileDescriptorEventpolicy, []int{2} }

func init() {
	proto.RegisterType((*EventPolicy)(nil), "monitoring.EventPolicy")
	proto.RegisterType((*EventPolicySpec)(nil), "monitoring.EventPolicySpec")
	proto.RegisterType((*EventPolicyStatus)(nil), "monitoring.EventPolicyStatus")
	proto.RegisterEnum("monitoring.MonitoringExportFormat", MonitoringExportFormat_name, MonitoringExportFormat_value)
}
func (m *EventPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEventpolicy(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintEventpolicy(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintEventpolicy(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintEventpolicy(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *EventPolicySpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPolicySpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Format) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEventpolicy(dAtA, i, uint64(len(m.Format)))
		i += copy(dAtA[i:], m.Format)
	}
	if m.Selector != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEventpolicy(dAtA, i, uint64(m.Selector.Size()))
		n5, err := m.Selector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.Targets) > 0 {
		for _, msg := range m.Targets {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintEventpolicy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.SyslogConfig != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEventpolicy(dAtA, i, uint64(m.SyslogConfig.Size()))
		n6, err := m.SyslogConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *EventPolicyStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPolicyStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintEventpolicy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EventPolicy) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovEventpolicy(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovEventpolicy(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovEventpolicy(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovEventpolicy(uint64(l))
	return n
}

func (m *EventPolicySpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Format)
	if l > 0 {
		n += 1 + l + sovEventpolicy(uint64(l))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovEventpolicy(uint64(l))
	}
	if len(m.Targets) > 0 {
		for _, e := range m.Targets {
			l = e.Size()
			n += 1 + l + sovEventpolicy(uint64(l))
		}
	}
	if m.SyslogConfig != nil {
		l = m.SyslogConfig.Size()
		n += 1 + l + sovEventpolicy(uint64(l))
	}
	return n
}

func (m *EventPolicyStatus) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovEventpolicy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEventpolicy(x uint64) (n int) {
	return sovEventpolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventpolicy
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
			return fmt.Errorf("proto: EventPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
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
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
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
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
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
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
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
			skippy, err := skipEventpolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventpolicy
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
func (m *EventPolicySpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventpolicy
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
			return fmt.Errorf("proto: EventPolicySpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPolicySpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Format = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &fields.Selector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Targets = append(m.Targets, &ExportConfig{})
			if err := m.Targets[len(m.Targets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyslogConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventpolicy
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
				return ErrInvalidLengthEventpolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SyslogConfig == nil {
				m.SyslogConfig = &SyslogExportConfig{}
			}
			if err := m.SyslogConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventpolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventpolicy
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
func (m *EventPolicyStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventpolicy
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
			return fmt.Errorf("proto: EventPolicyStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPolicyStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEventpolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventpolicy
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
func skipEventpolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventpolicy
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
					return 0, ErrIntOverflowEventpolicy
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
					return 0, ErrIntOverflowEventpolicy
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
				return 0, ErrInvalidLengthEventpolicy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEventpolicy
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
				next, err := skipEventpolicy(dAtA[start:])
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
	ErrInvalidLengthEventpolicy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventpolicy   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("eventpolicy.proto", fileDescriptorEventpolicy) }

var fileDescriptorEventpolicy = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x4c,
	0x18, 0xac, 0xdb, 0xaa, 0xfd, 0xff, 0x4d, 0x48, 0x93, 0x2d, 0x44, 0x4e, 0x0a, 0x76, 0x14, 0x09,
	0x14, 0x50, 0x62, 0xa3, 0x50, 0x10, 0xe2, 0xe8, 0x92, 0x70, 0xa1, 0xa4, 0xc4, 0xb9, 0x54, 0x20,
	0x81, 0xe3, 0x6c, 0xcc, 0x22, 0x7b, 0xd7, 0xf2, 0xae, 0x0b, 0x11, 0xe2, 0x08, 0xea, 0xb3, 0xf0,
	0x24, 0x3d, 0x56, 0x9c, 0x51, 0x84, 0x7c, 0x42, 0xb9, 0xf1, 0x06, 0xc8, 0x6b, 0x87, 0x3a, 0x49,
	0xcb, 0x6d, 0xbf, 0xf1, 0xcc, 0xec, 0xe4, 0x9b, 0x0d, 0x28, 0xa1, 0x13, 0x44, 0xb8, 0x4f, 0x5d,
	0x6c, 0x4f, 0x34, 0x3f, 0xa0, 0x9c, 0x42, 0xe0, 0x51, 0x82, 0x39, 0x0d, 0x30, 0x71, 0xaa, 0x37,
	0x1d, 0x4a, 0x1d, 0x17, 0xe9, 0x96, 0x8f, 0x75, 0x8b, 0x10, 0xca, 0x2d, 0x8e, 0x29, 0x61, 0x09,
	0xb3, 0xda, 0x71, 0x30, 0x7f, 0x17, 0x0e, 0x35, 0x9b, 0x7a, 0xba, 0x8f, 0x08, 0xb3, 0xc8, 0x88,
	0xea, 0xec, 0x83, 0x7e, 0x82, 0x08, 0xb6, 0x91, 0x1e, 0x72, 0xec, 0xb2, 0x58, 0xea, 0x20, 0x92,
	0x55, 0xeb, 0x98, 0xd8, 0x6e, 0x38, 0x42, 0x73, 0x9b, 0x56, 0xc6, 0xc6, 0xa1, 0x0e, 0xd5, 0x05,
	0x3c, 0x0c, 0xc7, 0x62, 0x12, 0x83, 0x38, 0xa5, 0xf4, 0xdb, 0x57, 0xdc, 0x1a, 0x67, 0xf4, 0x10,
	0xb7, 0x52, 0xda, 0xfd, 0x7f, 0xd0, 0xc6, 0x18, 0xb9, 0x23, 0xa6, 0x33, 0xe4, 0x22, 0x9b, 0xd3,
	0x20, 0x55, 0xe4, 0xd1, 0x47, 0x9f, 0x06, 0x3c, 0x99, 0xea, 0x3f, 0xd6, 0x41, 0xae, 0x13, 0x2f,
	0xe7, 0x48, 0x2c, 0x07, 0x3e, 0x02, 0xd2, 0x40, 0x96, 0x6a, 0x52, 0x23, 0xd7, 0xbe, 0xa6, 0x59,
	0x3e, 0xd6, 0x06, 0x13, 0x1f, 0x1d, 0x22, 0x6e, 0x19, 0xbb, 0x67, 0x53, 0x75, 0xed, 0x7c, 0xaa,
	0x4a, 0xb3, 0xa9, 0xba, 0xdd, 0xc4, 0xc4, 0xc5, 0x04, 0xf5, 0xe7, 0x07, 0xd8, 0x05, 0x52, 0x4f,
	0x5e, 0x17, 0xba, 0x1d, 0xa1, 0xeb, 0x0d, 0xdf, 0x23, 0x9b, 0x0b, 0x65, 0x35, 0xa3, 0x2c, 0xc4,
	0xd9, 0x9b, 0xd4, 0xc3, 0x1c, 0x79, 0x3e, 0x9f, 0xf4, 0x97, 0x66, 0xf8, 0x12, 0x6c, 0x9a, 0x3e,
	0xb2, 0xe5, 0x0d, 0x61, 0xb5, 0xa7, 0x5d, 0xb4, 0xa4, 0x65, 0x62, 0xc6, 0x14, 0xa3, 0x1c, 0xdb,
	0xc6, 0x96, 0xcc, 0x47, 0x76, 0xd6, 0x72, 0x71, 0x86, 0xc7, 0x60, 0xcb, 0xe4, 0x16, 0x0f, 0x99,
	0xbc, 0x29, 0x4c, 0x6f, 0x5d, 0x65, 0x2a, 0x48, 0x86, 0x9c, 0xda, 0x16, 0x99, 0x98, 0x33, 0xc6,
	0x2b, 0xc8, 0x13, 0xf5, 0xfb, 0x97, 0xca, 0x1e, 0xc8, 0xe9, 0x9f, 0x7a, 0xda, 0x00, 0x11, 0x8b,
	0xf0, 0xcf, 0x30, 0x2f, 0x9e, 0x5a, 0x2b, 0x79, 0x6b, 0xf5, 0xdf, 0xeb, 0x60, 0x67, 0x29, 0x37,
	0x7c, 0x05, 0xb6, 0xba, 0x34, 0xf0, 0x2c, 0x2e, 0xf6, 0xfc, 0xbf, 0xf1, 0xf8, 0xdb, 0xd7, 0x8a,
	0x6a, 0xf2, 0xa0, 0x43, 0x42, 0xaf, 0x71, 0xf8, 0x37, 0x5a, 0x47, 0xd4, 0x94, 0x30, 0xef, 0xce,
	0xa6, 0x6a, 0x69, 0x2c, 0x8e, 0xcd, 0xda, 0x45, 0xa0, 0x55, 0x08, 0xbe, 0x00, 0xff, 0x99, 0x69,
	0xdf, 0x69, 0x1d, 0x45, 0x2d, 0x79, 0x07, 0xda, 0x1c, 0x37, 0xca, 0xb3, 0xa9, 0x0a, 0xe7, 0xaf,
	0x22, 0xf3, 0xfb, 0x2e, 0xc1, 0xa0, 0x09, 0xb6, 0x07, 0x56, 0xe0, 0x20, 0xce, 0xe4, 0x8d, 0xda,
	0x46, 0x23, 0xd7, 0x96, 0x17, 0xb6, 0x27, 0x32, 0x1e, 0x50, 0x32, 0xc6, 0x8e, 0x71, 0x23, 0x0e,
	0xc9, 0x13, 0x72, 0xc6, 0x75, 0x15, 0x82, 0xaf, 0x41, 0xde, 0x9c, 0x30, 0x97, 0x3a, 0x89, 0x32,
	0xed, 0x45, 0xc9, 0x3a, 0x27, 0xdf, 0x17, 0xfc, 0xaf, 0xc7, 0xa5, 0xd8, 0xe2, 0x9c, 0x2d, 0x65,
	0x19, 0xa9, 0xef, 0x82, 0xd2, 0x4a, 0xab, 0xf7, 0xde, 0x82, 0xf2, 0xe5, 0xab, 0x85, 0x0a, 0x00,
	0xe6, 0xb1, 0xf9, 0xbc, 0xf7, 0xec, 0x8d, 0x61, 0x3e, 0x2d, 0xae, 0x55, 0x0b, 0xd1, 0x69, 0x05,
	0x30, 0x71, 0x7d, 0x6b, 0xc8, 0x46, 0xf0, 0x0e, 0x28, 0xa4, 0xdf, 0xfb, 0xdd, 0x83, 0x87, 0xfb,
	0xed, 0xfd, 0xa2, 0x54, 0x85, 0xd1, 0x69, 0xa5, 0x90, 0x72, 0x82, 0xb1, 0x1d, 0xa3, 0x46, 0xf1,
	0x2c, 0x52, 0xa4, 0xf3, 0x48, 0x91, 0x7e, 0x46, 0x8a, 0xf4, 0x2b, 0x52, 0xd6, 0x8e, 0xa4, 0xe1,
	0x96, 0xf8, 0x93, 0x3d, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xe4, 0x09, 0xed, 0x80, 0x04,
	0x00, 0x00,
}
