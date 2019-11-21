// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delphi/ipam.proto

package dnetproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import clientApi "github.com/pensando/sw/nic/delphi/gosdk/client_api"
import netproto3 "github.com/pensando/sw/nic/agent/protos/netproto"
import delphi "github.com/pensando/sw/nic/delphi/proto/delphi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Tenant object
type IPAMPolicy struct {
	Meta       *delphi.ObjectMeta    `protobuf:"bytes,1,opt,name=Meta" json:"Meta,omitempty"`
	Key        string                `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	IPAMPolicy *netproto3.IPAMPolicy `protobuf:"bytes,3,opt,name=IPAMPolicy" json:"IPAMPolicy,omitempty"`
}

func (m *IPAMPolicy) GetDelphiMessage() proto.Message {
	return m
}

func (m *IPAMPolicy) GetDelphiMeta() *delphi.ObjectMeta {
	return m.Meta
}

func (m *IPAMPolicy) SetDelphiMeta(meta *delphi.ObjectMeta) {
	m.Meta = meta
}

func (m *IPAMPolicy) GetDelphiKey() string {
	return fmt.Sprintf("%v", m.Key)
}

func (m *IPAMPolicy) GetDelphiKind() string {
	return "IPAMPolicy"
}

func (m *IPAMPolicy) GetDelphiPath() string {
	return fmt.Sprintf("%s|%s", m.GetDelphiKind(), m.GetDelphiKey())
}

func (m *IPAMPolicy) DelphiClone() clientApi.BaseObject {
	obj, _ := proto.Clone(m).(*IPAMPolicy)
	return obj
}

func IPAMPolicyMount(client clientApi.Client, mode delphi.MountMode) {
	client.MountKind("IPAMPolicy", mode)
}

func IPAMPolicyMountKey(client clientApi.Client, key string, mode delphi.MountMode) {
	client.MountKindKey("IPAMPolicy", fmt.Sprintf("%v", key), mode)
}

func GetIPAMPolicy(client clientApi.Client, key string) *IPAMPolicy {
	o := client.GetObject("IPAMPolicy", fmt.Sprintf("%v", key))
	if o == nil {
		return nil
	}
	obj, ok := o.(*IPAMPolicy)
	if ok != true {
		panic("Cast failed")
	}
	return obj
}

func (m *IPAMPolicy) IsPersistent() bool {
	return false
}
func IPAMPolicyFactory(sdkClient clientApi.Client, data []byte) (clientApi.BaseObject, error) {
	var msg IPAMPolicy
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func IPAMPolicyWatch(client clientApi.Client, reactor IPAMPolicyReactor) {
	client.WatchKind("IPAMPolicy", reactor)
}
func IPAMPolicyList(client clientApi.Client) []*IPAMPolicy {
	bobjs := client.List("IPAMPolicy")
	objs := make([]*IPAMPolicy, 0)
	for _, bobj := range bobjs {
		obj, _ := bobj.(*IPAMPolicy)
		objs = append(objs, obj)
	}
	return objs
}
func (m *IPAMPolicy) TriggerEvent(sdkClient clientApi.Client, old clientApi.BaseObject, op delphi.ObjectOperation, rl []clientApi.BaseReactor) {
	for _, r := range rl {
		rctr, ok := r.(IPAMPolicyReactor)
		if ok == false {
			panic("Not a Reactor")
		}
		if op == delphi.ObjectOperation_SetOp {
			if old == nil {
				rctr.OnIPAMPolicyCreate(m)
			} else {
				oldObj, ok := old.(*IPAMPolicy)
				if ok == false {
					panic("Not an IPAMPolicy object")
				}
				rctr.OnIPAMPolicyUpdate(oldObj, m)
			}
		} else {
			rctr.OnIPAMPolicyDelete(m)
		}
	}
}

type IPAMPolicyReactor interface {
	OnIPAMPolicyCreate(obj *IPAMPolicy)
	OnIPAMPolicyUpdate(old *IPAMPolicy, obj *IPAMPolicy)
	OnIPAMPolicyDelete(obj *IPAMPolicy)
}

func (m *IPAMPolicy) Reset()                    { *m = IPAMPolicy{} }
func (m *IPAMPolicy) String() string            { return proto.CompactTextString(m) }
func (*IPAMPolicy) ProtoMessage()               {}
func (*IPAMPolicy) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *IPAMPolicy) GetMeta() *delphi.ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *IPAMPolicy) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *IPAMPolicy) GetIPAMPolicy() *netproto3.IPAMPolicy {
	if m != nil {
		return m.IPAMPolicy
	}
	return nil
}

func init() {
	clientApi.RegisterFactory("IPAMPolicy", IPAMPolicyFactory)
	proto.RegisterType((*IPAMPolicy)(nil), "dnetproto.IPAMPolicy")
}

func init() { proto.RegisterFile("delphi/ipam.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x49, 0xcd, 0x29,
	0xc8, 0xc8, 0xd4, 0xcf, 0x2c, 0x48, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c,
	0xc9, 0x4b, 0x2d, 0x01, 0x33, 0xa5, 0xac, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x0b, 0x52, 0xf3, 0x8a, 0x13, 0xf3, 0x52, 0xf2, 0xf5, 0x8b, 0xcb, 0xf5, 0xf3, 0x32,
	0x93, 0xf5, 0x13, 0xd3, 0x53, 0xf3, 0x4a, 0xf4, 0xc1, 0xea, 0x8a, 0xf5, 0x61, 0x3a, 0x90, 0xcc,
	0x91, 0xe2, 0x81, 0x18, 0x0d, 0xe1, 0x29, 0x35, 0x31, 0x72, 0x71, 0x79, 0x06, 0x38, 0xfa, 0x06,
	0xe4, 0xe7, 0x64, 0x26, 0x57, 0x0a, 0xa9, 0x71, 0xb1, 0xf8, 0xa6, 0x96, 0x24, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x1b, 0x09, 0xe9, 0x41, 0xd5, 0xfa, 0x27, 0x65, 0xa5, 0x26, 0x97, 0x80, 0x64,
	0x82, 0xc0, 0xf2, 0x42, 0x02, 0x5c, 0xcc, 0xde, 0xa9, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x20, 0xa6, 0x90, 0x09, 0xb2, 0x39, 0x12, 0xcc, 0x60, 0xfd, 0x22, 0x7a, 0x30, 0x07, 0xe8,
	0x21, 0xe4, 0x82, 0x90, 0xd4, 0x59, 0xb1, 0x34, 0x4c, 0x57, 0x62, 0x4c, 0x62, 0x03, 0xab, 0x31,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xcf, 0x9c, 0x57, 0xf6, 0x00, 0x00, 0x00,
}
