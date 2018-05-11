// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: match.proto

package netproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/pensando/sw/api"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MatchSelector_Type int32

const (
	MatchSelector_NONE   MatchSelector_Type = 0
	MatchSelector_L4PORT MatchSelector_Type = 1
)

var MatchSelector_Type_name = map[int32]string{
	0: "NONE",
	1: "L4PORT",
}
var MatchSelector_Type_value = map[string]int32{
	"NONE":   0,
	"L4PORT": 1,
}

func (x MatchSelector_Type) String() string {
	return proto.EnumName(MatchSelector_Type_name, int32(x))
}
func (MatchSelector_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorMatch, []int{0, 0} }

// Common MatchSelector. ToDo Add ICMP Match criteria
type MatchSelector struct {
	// Automatically interpret the string as an octet, a CIDR or an hyphen separated range
	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"address,omitempty"`
	// Match on the security group
	SecurityGroup string `protobuf:"bytes,2,opt,name=SecurityGroup,proto3" json:"security-group,omitempty"`
	// Match on the App info
	App       string `protobuf:"bytes,3,opt,name=App,proto3" json:"app"`
	AppConfig string `protobuf:"bytes,4,opt,name=AppConfig,proto3" json:"app-config"`
}

func (m *MatchSelector) Reset()                    { *m = MatchSelector{} }
func (m *MatchSelector) String() string            { return proto.CompactTextString(m) }
func (*MatchSelector) ProtoMessage()               {}
func (*MatchSelector) Descriptor() ([]byte, []int) { return fileDescriptorMatch, []int{0} }

func (m *MatchSelector) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MatchSelector) GetSecurityGroup() string {
	if m != nil {
		return m.SecurityGroup
	}
	return ""
}

func (m *MatchSelector) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *MatchSelector) GetAppConfig() string {
	if m != nil {
		return m.AppConfig
	}
	return ""
}

func init() {
	proto.RegisterType((*MatchSelector)(nil), "netproto.MatchSelector")
	proto.RegisterEnum("netproto.MatchSelector_Type", MatchSelector_Type_name, MatchSelector_Type_value)
}
func (m *MatchSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MatchSelector) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMatch(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.SecurityGroup) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMatch(dAtA, i, uint64(len(m.SecurityGroup)))
		i += copy(dAtA[i:], m.SecurityGroup)
	}
	if len(m.App) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMatch(dAtA, i, uint64(len(m.App)))
		i += copy(dAtA[i:], m.App)
	}
	if len(m.AppConfig) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMatch(dAtA, i, uint64(len(m.AppConfig)))
		i += copy(dAtA[i:], m.AppConfig)
	}
	return i, nil
}

func encodeVarintMatch(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MatchSelector) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	l = len(m.SecurityGroup)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	l = len(m.App)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	l = len(m.AppConfig)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	return n
}

func sovMatch(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMatch(x uint64) (n int) {
	return sovMatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MatchSelector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
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
			return fmt.Errorf("proto: MatchSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MatchSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityGroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityGroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.App = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppConfig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppConfig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMatch
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
func skipMatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
				return 0, ErrInvalidLengthMatch
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMatch
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
				next, err := skipMatch(dAtA[start:])
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
	ErrInvalidLengthMatch = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMatch   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("match.proto", fileDescriptorMatch) }

var fileDescriptorMatch = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x5f, 0x6a, 0xea, 0x40,
	0x14, 0x87, 0x8d, 0x8a, 0x57, 0xe7, 0x22, 0x78, 0x03, 0x17, 0xa2, 0x95, 0x58, 0x84, 0x16, 0x0b,
	0xd5, 0x29, 0xb4, 0x0f, 0x7d, 0x8d, 0x45, 0xfa, 0xd2, 0xaa, 0x44, 0x37, 0x30, 0x26, 0xe3, 0x38,
	0x90, 0xcc, 0x1c, 0x92, 0x49, 0x8b, 0x1b, 0x70, 0x0d, 0x5d, 0x83, 0x74, 0x21, 0x7d, 0xec, 0x0a,
	0xa4, 0xd8, 0x37, 0x57, 0x51, 0x32, 0x2a, 0x55, 0x68, 0xdf, 0xce, 0x9f, 0xef, 0xfb, 0xe5, 0x64,
	0xd0, 0xdf, 0x90, 0x28, 0x6f, 0xd6, 0x81, 0x48, 0x2a, 0x69, 0x16, 0x05, 0x55, 0xba, 0xaa, 0xd5,
	0x99, 0x94, 0x2c, 0xa0, 0x98, 0x00, 0xc7, 0x44, 0x08, 0xa9, 0x88, 0xe2, 0x52, 0xc4, 0x5b, 0xae,
	0xd6, 0x63, 0x5c, 0xcd, 0x92, 0x49, 0xc7, 0x93, 0x21, 0x06, 0x2a, 0x62, 0x22, 0x7c, 0x89, 0xe3,
	0x67, 0xfc, 0x44, 0x05, 0xf7, 0x28, 0x4e, 0x14, 0x0f, 0xe2, 0x54, 0x65, 0x54, 0x1c, 0xda, 0x98,
	0x0b, 0x2f, 0x48, 0x7c, 0xba, 0x8f, 0x69, 0x1f, 0xc4, 0x30, 0xc9, 0x24, 0xd6, 0xe3, 0x49, 0x32,
	0xd5, 0x9d, 0x6e, 0x74, 0xb5, 0xc3, 0xcf, 0x7e, 0xf9, 0x6a, 0x7a, 0x63, 0x48, 0x15, 0xd9, 0x62,
	0xcd, 0xd7, 0x2c, 0x2a, 0x3f, 0xa6, 0x3f, 0x35, 0xa2, 0x01, 0xf5, 0x94, 0x8c, 0x4c, 0x8c, 0xfe,
	0x38, 0xbe, 0x1f, 0xd1, 0x38, 0xb6, 0x8c, 0x53, 0xa3, 0x55, 0xea, 0xfe, 0xdf, 0xac, 0x1a, 0xff,
	0xc8, 0x76, 0x74, 0x29, 0x43, 0xae, 0x68, 0x08, 0x6a, 0xee, 0xee, 0x29, 0xb3, 0x8b, 0xca, 0x23,
	0xea, 0x25, 0x11, 0x57, 0xf3, 0xfb, 0x48, 0x26, 0x60, 0x65, 0xb5, 0x56, 0xdf, 0xac, 0x1a, 0x56,
	0xbc, 0x5b, 0xb4, 0x59, 0xba, 0x39, 0xb0, 0x8f, 0x15, 0xf3, 0x16, 0xe5, 0x1c, 0x00, 0x2b, 0xa7,
	0xcd, 0xf3, 0xe5, 0xa2, 0x7a, 0x32, 0x52, 0x51, 0x4f, 0x24, 0x61, 0xeb, 0xe8, 0xb8, 0xce, 0x78,
	0x0e, 0xf4, 0x62, 0xb3, 0x6a, 0xe4, 0x08, 0x80, 0x9b, 0x2a, 0x66, 0x1f, 0x95, 0x1c, 0x80, 0x3b,
	0x29, 0xa6, 0x9c, 0x59, 0x79, 0xed, 0x5f, 0x2d, 0x17, 0x55, 0xfb, 0x67, 0xdf, 0x01, 0xd8, 0x47,
	0x20, 0x02, 0xd0, 0xf6, 0xb4, 0xe7, 0x7e, 0x47, 0x34, 0xeb, 0x28, 0x9f, 0x12, 0x66, 0x11, 0xe5,
	0xfb, 0x83, 0x7e, 0xaf, 0x92, 0x31, 0x11, 0x2a, 0x3c, 0xdc, 0x0c, 0x07, 0xee, 0xb8, 0x62, 0x74,
	0x2b, 0x6f, 0x6b, 0xdb, 0x78, 0x5f, 0xdb, 0xc6, 0xc7, 0xda, 0x36, 0x5e, 0x3e, 0xed, 0xcc, 0xd0,
	0x98, 0x14, 0xf4, 0x4b, 0x5e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x19, 0xa0, 0x20, 0x1d,
	0x02, 0x00, 0x00,
}
