// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orchestration.proto

/*
	Package orchestration is a generated protocol buffer package.

	Service name

	It is generated from these files:
		orchestration.proto
		svc_orchestration.proto

	It has these top-level messages:
		Orchestrator
		OrchestratorSpec
		OrchestratorStatus
		AutoMsgOrchestratorWatchHelper
		OrchestratorList
*/
package orchestration

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"
import monitoring "github.com/pensando/sw/api/generated/monitoring"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
type OrchestratorSpec_OrchestratorType int32

const (
	// ui-hint: vcenter
	OrchestratorSpec_VCenter OrchestratorSpec_OrchestratorType = 0
)

var OrchestratorSpec_OrchestratorType_name = map[int32]string{
	0: "VCenter",
}
var OrchestratorSpec_OrchestratorType_value = map[string]int32{
	"VCenter": 0,
}

func (OrchestratorSpec_OrchestratorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOrchestration, []int{1, 0}
}

//
type OrchestratorStatus_ConnectionStatus int32

const (
	// ui-hint: Unknown
	OrchestratorStatus_Unknown OrchestratorStatus_ConnectionStatus = 0
	// ui-hint: Success
	OrchestratorStatus_Success OrchestratorStatus_ConnectionStatus = 1
	// ui-hint: Failure
	OrchestratorStatus_Failure OrchestratorStatus_ConnectionStatus = 2
	// ui-hint: Degraded
	OrchestratorStatus_Degraded OrchestratorStatus_ConnectionStatus = 3
)

var OrchestratorStatus_ConnectionStatus_name = map[int32]string{
	0: "Unknown",
	1: "Success",
	2: "Failure",
	3: "Degraded",
}
var OrchestratorStatus_ConnectionStatus_value = map[string]int32{
	"Unknown":  0,
	"Success":  1,
	"Failure":  2,
	"Degraded": 3,
}

func (OrchestratorStatus_ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOrchestration, []int{2, 0}
}

// --------------------------------- ORCHESTRATOR ---------------------------------------------
//
// Orchestrator represents the config object which allows Venice to connect to the appropriate orchestrator
//
//
type Orchestrator struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	//
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	// Spec contains the configuration of the vcenter.
	Spec OrchestratorSpec `protobuf:"bytes,3,opt,name=Spec,json=spec,omitempty" json:"spec,omitempty"`
	// Status contains the current state of the cluster.
	Status OrchestratorStatus `protobuf:"bytes,4,opt,name=Status,json=status,omitempty" json:"status,omitempty"`
}

func (m *Orchestrator) Reset()                    { *m = Orchestrator{} }
func (m *Orchestrator) String() string            { return proto.CompactTextString(m) }
func (*Orchestrator) ProtoMessage()               {}
func (*Orchestrator) Descriptor() ([]byte, []int) { return fileDescriptorOrchestration, []int{0} }

func (m *Orchestrator) GetSpec() OrchestratorSpec {
	if m != nil {
		return m.Spec
	}
	return OrchestratorSpec{}
}

func (m *Orchestrator) GetStatus() OrchestratorStatus {
	if m != nil {
		return m.Status
	}
	return OrchestratorStatus{}
}

// OrchestratorSpec contains the configuration of the cluster.
type OrchestratorSpec struct {
	// Type of orchestrator
	Type string `protobuf:"bytes,1,opt,name=Type,json=type,proto3" json:"type"`
	// URI of the orchestrator
	URI string `protobuf:"bytes,2,opt,name=URI,json=uri,proto3" json:"uri"`
	// Credentials for the orchestrator
	Credentials *monitoring.ExternalCred `protobuf:"bytes,3,opt,name=Credentials,json=credentials,omitempty" json:"credentials,omitempty"`
	// Namespaces that will be managed by this orchestrator. "all_namespaces" will manage all namespaces.
	ManageNamespaces []string `protobuf:"bytes,4,rep,name=ManageNamespaces,json=manage-namespaces,omitempty" json:"manage-namespaces,omitempty"`
}

func (m *OrchestratorSpec) Reset()                    { *m = OrchestratorSpec{} }
func (m *OrchestratorSpec) String() string            { return proto.CompactTextString(m) }
func (*OrchestratorSpec) ProtoMessage()               {}
func (*OrchestratorSpec) Descriptor() ([]byte, []int) { return fileDescriptorOrchestration, []int{1} }

func (m *OrchestratorSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OrchestratorSpec) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (m *OrchestratorSpec) GetCredentials() *monitoring.ExternalCred {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *OrchestratorSpec) GetManageNamespaces() []string {
	if m != nil {
		return m.ManageNamespaces
	}
	return nil
}

// OrchestratorStatus contains the current state of connection with the orchestrator.
type OrchestratorStatus struct {
	//
	Status string `protobuf:"bytes,1,opt,name=Status,json=connection-status,proto3" json:"connection-status"`
	//
	LastTransitionTime *api.Timestamp `protobuf:"bytes,2,opt,name=LastTransitionTime,json=last-transition-time,omitempty" json:"last-transition-time,omitempty"`
	//
	Message string `protobuf:"bytes,3,opt,name=Message,json=message,omitempty,proto3" json:"message,omitempty"`
	//
	OrchID int32 `protobuf:"varint,4,opt,name=OrchID,json=orch-id,omitempty,proto3" json:"orch-id,omitempty"`
	//
	DiscoveredNamespaces []string `protobuf:"bytes,5,rep,name=DiscoveredNamespaces,json=discovered-namespaces,omitempty" json:"discovered-namespaces,omitempty"`
	//
	IncompatibleDSCs []string `protobuf:"bytes,6,rep,name=IncompatibleDSCs,json=incompatible-dscs,omitempty" json:"incompatible-dscs,omitempty"`
}

func (m *OrchestratorStatus) Reset()                    { *m = OrchestratorStatus{} }
func (m *OrchestratorStatus) String() string            { return proto.CompactTextString(m) }
func (*OrchestratorStatus) ProtoMessage()               {}
func (*OrchestratorStatus) Descriptor() ([]byte, []int) { return fileDescriptorOrchestration, []int{2} }

func (m *OrchestratorStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrchestratorStatus) GetLastTransitionTime() *api.Timestamp {
	if m != nil {
		return m.LastTransitionTime
	}
	return nil
}

func (m *OrchestratorStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OrchestratorStatus) GetOrchID() int32 {
	if m != nil {
		return m.OrchID
	}
	return 0
}

func (m *OrchestratorStatus) GetDiscoveredNamespaces() []string {
	if m != nil {
		return m.DiscoveredNamespaces
	}
	return nil
}

func (m *OrchestratorStatus) GetIncompatibleDSCs() []string {
	if m != nil {
		return m.IncompatibleDSCs
	}
	return nil
}

func init() {
	proto.RegisterType((*Orchestrator)(nil), "orchestration.Orchestrator")
	proto.RegisterType((*OrchestratorSpec)(nil), "orchestration.OrchestratorSpec")
	proto.RegisterType((*OrchestratorStatus)(nil), "orchestration.OrchestratorStatus")
	proto.RegisterEnum("orchestration.OrchestratorSpec_OrchestratorType", OrchestratorSpec_OrchestratorType_name, OrchestratorSpec_OrchestratorType_value)
	proto.RegisterEnum("orchestration.OrchestratorStatus_ConnectionStatus", OrchestratorStatus_ConnectionStatus_name, OrchestratorStatus_ConnectionStatus_value)
}
func (m *Orchestrator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Orchestrator) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintOrchestration(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintOrchestration(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintOrchestration(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintOrchestration(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *OrchestratorSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrchestratorSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.URI) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(len(m.URI)))
		i += copy(dAtA[i:], m.URI)
	}
	if m.Credentials != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(m.Credentials.Size()))
		n5, err := m.Credentials.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.ManageNamespaces) > 0 {
		for _, s := range m.ManageNamespaces {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *OrchestratorStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrchestratorStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if m.LastTransitionTime != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(m.LastTransitionTime.Size()))
		n6, err := m.LastTransitionTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.OrchID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintOrchestration(dAtA, i, uint64(m.OrchID))
	}
	if len(m.DiscoveredNamespaces) > 0 {
		for _, s := range m.DiscoveredNamespaces {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.IncompatibleDSCs) > 0 {
		for _, s := range m.IncompatibleDSCs {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintOrchestration(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Orchestrator) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovOrchestration(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovOrchestration(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovOrchestration(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovOrchestration(uint64(l))
	return n
}

func (m *OrchestratorSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovOrchestration(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovOrchestration(uint64(l))
	}
	if m.Credentials != nil {
		l = m.Credentials.Size()
		n += 1 + l + sovOrchestration(uint64(l))
	}
	if len(m.ManageNamespaces) > 0 {
		for _, s := range m.ManageNamespaces {
			l = len(s)
			n += 1 + l + sovOrchestration(uint64(l))
		}
	}
	return n
}

func (m *OrchestratorStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovOrchestration(uint64(l))
	}
	if m.LastTransitionTime != nil {
		l = m.LastTransitionTime.Size()
		n += 1 + l + sovOrchestration(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovOrchestration(uint64(l))
	}
	if m.OrchID != 0 {
		n += 1 + sovOrchestration(uint64(m.OrchID))
	}
	if len(m.DiscoveredNamespaces) > 0 {
		for _, s := range m.DiscoveredNamespaces {
			l = len(s)
			n += 1 + l + sovOrchestration(uint64(l))
		}
	}
	if len(m.IncompatibleDSCs) > 0 {
		for _, s := range m.IncompatibleDSCs {
			l = len(s)
			n += 1 + l + sovOrchestration(uint64(l))
		}
	}
	return n
}

func sovOrchestration(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOrchestration(x uint64) (n int) {
	return sovOrchestration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Orchestrator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrchestration
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
			return fmt.Errorf("proto: Orchestrator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Orchestrator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
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
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
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
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
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
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
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
			skippy, err := skipOrchestration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrchestration
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
func (m *OrchestratorSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrchestration
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
			return fmt.Errorf("proto: OrchestratorSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrchestratorSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Credentials == nil {
				m.Credentials = &monitoring.ExternalCred{}
			}
			if err := m.Credentials.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManageNamespaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManageNamespaces = append(m.ManageNamespaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrchestration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrchestration
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
func (m *OrchestratorStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrchestration
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
			return fmt.Errorf("proto: OrchestratorStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrchestratorStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTransitionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastTransitionTime == nil {
				m.LastTransitionTime = &api.Timestamp{}
			}
			if err := m.LastTransitionTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrchID", wireType)
			}
			m.OrchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrchID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscoveredNamespaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DiscoveredNamespaces = append(m.DiscoveredNamespaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncompatibleDSCs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrchestration
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
				return ErrInvalidLengthOrchestration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncompatibleDSCs = append(m.IncompatibleDSCs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrchestration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrchestration
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
func skipOrchestration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrchestration
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
					return 0, ErrIntOverflowOrchestration
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
					return 0, ErrIntOverflowOrchestration
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
				return 0, ErrInvalidLengthOrchestration
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOrchestration
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
				next, err := skipOrchestration(dAtA[start:])
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
	ErrInvalidLengthOrchestration = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrchestration   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("orchestration.proto", fileDescriptorOrchestration) }

var fileDescriptorOrchestration = []byte{
	// 811 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xae, 0x9b, 0xfe, 0xd8, 0x4e, 0xb3, 0xc5, 0x3b, 0xcb, 0x8f, 0x38, 0xa0, 0x04, 0x82, 0x80,
	0x16, 0xc5, 0xb6, 0x16, 0x24, 0x0e, 0x2b, 0xb8, 0xb8, 0xed, 0x4a, 0x95, 0xb6, 0x5b, 0x94, 0xb4,
	0x1c, 0x90, 0x10, 0x9a, 0x8c, 0x5f, 0xdd, 0x01, 0x7b, 0xc6, 0x9a, 0x19, 0x77, 0xb7, 0x12, 0x57,
	0xaa, 0xfd, 0x5b, 0x7a, 0xe3, 0x3f, 0xe0, 0xd8, 0xe3, 0x8a, 0x13, 0xa7, 0x08, 0xe5, 0x84, 0xfa,
	0x17, 0x70, 0x44, 0x33, 0x76, 0xb3, 0x4e, 0xd3, 0xed, 0x5e, 0x2c, 0xbf, 0xef, 0x7d, 0xdf, 0x97,
	0xf1, 0xf7, 0xde, 0x04, 0x3d, 0x14, 0x92, 0x9e, 0x80, 0xd2, 0x92, 0x68, 0x26, 0x78, 0x90, 0x4b,
	0xa1, 0x05, 0xbe, 0x3f, 0x03, 0xb6, 0x3f, 0x4a, 0x84, 0x48, 0x52, 0x08, 0x49, 0xce, 0x42, 0xc2,
	0xb9, 0xd0, 0x16, 0x56, 0x25, 0xb9, 0xbd, 0x9b, 0x30, 0x7d, 0x52, 0x8c, 0x02, 0x2a, 0xb2, 0x30,
	0x07, 0xae, 0x08, 0x8f, 0x45, 0xa8, 0x9e, 0x87, 0xa7, 0xc0, 0x19, 0x85, 0xb0, 0xd0, 0x2c, 0x55,
	0x46, 0x9a, 0x00, 0xaf, 0xab, 0x43, 0xc6, 0x69, 0x5a, 0xc4, 0x70, 0x6d, 0xe3, 0xd7, 0x6c, 0x12,
	0x91, 0x88, 0xd0, 0xc2, 0xa3, 0xe2, 0xd8, 0x56, 0xb6, 0xb0, 0x6f, 0x15, 0xfd, 0xb3, 0x37, 0xfc,
	0xaa, 0x39, 0x63, 0x06, 0x9a, 0x54, 0xb4, 0x26, 0xbc, 0xc8, 0x85, 0xd4, 0x65, 0xd5, 0xfb, 0x7b,
	0x11, 0x35, 0x0f, 0xa6, 0x9f, 0x26, 0x24, 0xfe, 0x06, 0x39, 0x87, 0x2d, 0xe7, 0x63, 0x67, 0x73,
	0xfd, 0xab, 0xfb, 0x01, 0xc9, 0x59, 0x70, 0x78, 0x96, 0xc3, 0x3e, 0x68, 0x12, 0x3d, 0xbc, 0x1c,
	0x77, 0x17, 0x5e, 0x8d, 0xbb, 0xce, 0xd5, 0xb8, 0xbb, 0xda, 0x67, 0x3c, 0x65, 0x1c, 0x06, 0xd7,
	0x2f, 0xf8, 0x09, 0x72, 0x0e, 0x5a, 0x8b, 0x56, 0xf7, 0x8e, 0xd5, 0x1d, 0x8c, 0x7e, 0x01, 0xaa,
	0xad, 0xb2, 0x5d, 0x53, 0x6e, 0x98, 0xa3, 0xf4, 0x45, 0xc6, 0x34, 0x64, 0xb9, 0x3e, 0x1b, 0xdc,
	0xa8, 0xf1, 0x11, 0x5a, 0x1a, 0xe6, 0x40, 0x5b, 0x0d, 0x6b, 0xd5, 0x0d, 0x66, 0x87, 0x51, 0x3f,
	0xaa, 0xa1, 0x45, 0xef, 0x1b, 0x6b, 0x63, 0xab, 0x72, 0xa0, 0x75, 0xdb, 0xd9, 0x1a, 0xff, 0x84,
	0x56, 0x86, 0x9a, 0xe8, 0x42, 0xb5, 0x96, 0xac, 0xf1, 0x27, 0x77, 0x19, 0x5b, 0x62, 0xd4, 0xaa,
	0xac, 0x5d, 0x65, 0xeb, 0x9a, 0xf9, 0x1c, 0xf2, 0x18, 0xff, 0xf5, 0xbb, 0xb7, 0x81, 0x9b, 0xa2,
	0xe6, 0xd2, 0xfb, 0x6f, 0x11, 0xb9, 0x37, 0xcf, 0x8b, 0x9f, 0xa1, 0x25, 0x13, 0xa8, 0x4d, 0x78,
	0x2d, 0x7a, 0x7c, 0x71, 0xee, 0x7d, 0x39, 0xd4, 0x72, 0x97, 0x17, 0xd9, 0xe6, 0x4d, 0xee, 0xcc,
	0x99, 0x8c, 0x68, 0xeb, 0xcf, 0x73, 0xcf, 0x1c, 0x68, 0x49, 0x9f, 0xe5, 0x30, 0xb0, 0x4f, 0xbc,
	0x85, 0x1a, 0x47, 0x83, 0x3d, 0x1b, 0xfc, 0x5a, 0xf4, 0xc1, 0xc5, 0xb9, 0xd7, 0x1c, 0x6a, 0xf9,
	0x14, 0xf8, 0xe6, 0xa3, 0xbe, 0xff, 0x68, 0xeb, 0x6a, 0xdc, 0x6d, 0x14, 0x92, 0x0d, 0xcc, 0x03,
	0xff, 0x8c, 0xd6, 0xb7, 0x25, 0xc4, 0xc0, 0x35, 0x23, 0xa9, 0xaa, 0x02, 0x6e, 0x05, 0x99, 0xe0,
	0x4c, 0x0b, 0xc9, 0x78, 0x12, 0xec, 0xbe, 0xd0, 0x20, 0x39, 0x49, 0x0d, 0x2d, 0xf2, 0xae, 0xc6,
	0xdd, 0xf7, 0xe8, 0x6b, 0x41, 0xed, 0xfb, 0x6f, 0x87, 0xf1, 0x8f, 0xc8, 0xdd, 0x27, 0x9c, 0x24,
	0xf0, 0x8c, 0x64, 0xa0, 0x72, 0x42, 0xc1, 0xa4, 0xdd, 0xd8, 0x5c, 0x8b, 0x3e, 0xbd, 0x2c, 0x87,
	0xff, 0x61, 0x66, 0xfb, 0x3e, 0x9f, 0x12, 0x6a, 0xae, 0x77, 0x35, 0x7b, 0xfd, 0xd9, 0x2c, 0x4d,
	0x1c, 0xb8, 0x85, 0x56, 0x7f, 0xd8, 0x06, 0xae, 0x41, 0xba, 0x0b, 0xed, 0xf5, 0xc9, 0x4b, 0x6f,
	0xf5, 0x94, 0xda, 0xb2, 0xf7, 0xc7, 0x32, 0xc2, 0xf3, 0x13, 0xc5, 0x27, 0xd3, 0x25, 0x28, 0xe3,
	0xdf, 0xbd, 0x38, 0xf7, 0xfa, 0xb7, 0xc6, 0x6f, 0x59, 0xc1, 0xb6, 0xe0, 0x1c, 0xa8, 0xd9, 0x8f,
	0x12, 0x30, 0x79, 0x3e, 0xa0, 0x53, 0xd0, 0x2f, 0xd7, 0x60, 0x30, 0x0f, 0x61, 0x89, 0xf0, 0x53,
	0xa2, 0xf4, 0xa1, 0x24, 0x5c, 0x31, 0xd3, 0x38, 0x64, 0x19, 0x54, 0xd7, 0x63, 0xa3, 0xbc, 0x56,
	0x2c, 0x03, 0xa5, 0x49, 0x96, 0x47, 0x9f, 0x57, 0xe1, 0x74, 0x52, 0xa2, 0xb4, 0xaf, 0xa7, 0x12,
	0x5f, 0xb3, 0x0c, 0x6a, 0xf9, 0xbc, 0xa5, 0x8f, 0xbf, 0x43, 0xab, 0xfb, 0xa0, 0x14, 0x49, 0xc0,
	0xce, 0x76, 0x2d, 0xf2, 0x2a, 0xe3, 0x07, 0x59, 0x09, 0xd7, 0xbc, 0xe6, 0x21, 0xfc, 0x2d, 0x5a,
	0x31, 0x11, 0xec, 0xed, 0xd8, 0x1b, 0xb2, 0xfc, 0x5a, 0x6d, 0x96, 0xda, 0x67, 0x71, 0x5d, 0x3d,
	0x07, 0xe1, 0x63, 0xf4, 0xee, 0x0e, 0x53, 0x54, 0x9c, 0x82, 0x84, 0xb8, 0x36, 0xff, 0x65, 0x3b,
	0xff, 0x2f, 0x2a, 0xaf, 0x6e, 0x3c, 0xe5, 0xdc, 0xbe, 0x03, 0x6f, 0x23, 0x98, 0x1d, 0xdb, 0xe3,
	0x54, 0x64, 0x39, 0xd1, 0x6c, 0x94, 0xc2, 0xce, 0x70, 0x5b, 0xb5, 0x56, 0x66, 0x77, 0x8c, 0xd5,
	0xfa, 0x7e, 0xac, 0xe8, 0xcc, 0x8e, 0xdd, 0xd1, 0xec, 0xfd, 0x86, 0xdc, 0x9b, 0x13, 0x37, 0x3b,
	0x76, 0xc4, 0x7f, 0xe5, 0xe2, 0x39, 0xbf, 0xde, 0xb1, 0xa2, 0x2c, 0x4d, 0x67, 0x58, 0x50, 0x0a,
	0x4a, 0xb9, 0x4e, 0xd9, 0x51, 0x65, 0x69, 0x3a, 0x4f, 0x08, 0x4b, 0x0b, 0x09, 0xee, 0x62, 0xd9,
	0x39, 0x2e, 0x4b, 0xdc, 0x46, 0xf7, 0x76, 0x20, 0x91, 0x24, 0x86, 0xd8, 0x6d, 0xb4, 0x9b, 0x93,
	0x97, 0xde, 0xbd, 0xb8, 0xaa, 0x23, 0xf7, 0x72, 0xd2, 0x71, 0x5e, 0x4d, 0x3a, 0xce, 0x3f, 0x93,
	0x8e, 0xf3, 0xef, 0xa4, 0xb3, 0xf0, 0xbd, 0x33, 0x5a, 0xb1, 0x7f, 0xd2, 0x5f, 0xff, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xcb, 0x4e, 0x09, 0x93, 0x06, 0x00, 0x00,
}
