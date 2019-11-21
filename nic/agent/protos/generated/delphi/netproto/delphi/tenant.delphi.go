// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delphi/tenant.proto

package dnetproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import clientApi "github.com/pensando/sw/nic/delphi/gosdk/client_api"
import netproto10 "github.com/pensando/sw/nic/agent/protos/netproto"
import delphi "github.com/pensando/sw/nic/delphi/proto/delphi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Tenant object
type Tenant struct {
	Meta   *delphi.ObjectMeta `protobuf:"bytes,1,opt,name=Meta" json:"Meta,omitempty"`
	Key    string             `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	Tenant *netproto10.Tenant `protobuf:"bytes,3,opt,name=Tenant" json:"Tenant,omitempty"`
}

func (m *Tenant) GetDelphiMessage() proto.Message {
	return m
}

func (m *Tenant) GetDelphiMeta() *delphi.ObjectMeta {
	return m.Meta
}

func (m *Tenant) SetDelphiMeta(meta *delphi.ObjectMeta) {
	m.Meta = meta
}

func (m *Tenant) GetDelphiKey() string {
	return fmt.Sprintf("%v", m.Key)
}

func (m *Tenant) GetDelphiKind() string {
	return "Tenant"
}

func (m *Tenant) GetDelphiPath() string {
	return fmt.Sprintf("%s|%s", m.GetDelphiKind(), m.GetDelphiKey())
}

func (m *Tenant) DelphiClone() clientApi.BaseObject {
	obj, _ := proto.Clone(m).(*Tenant)
	return obj
}

func TenantMount(client clientApi.Client, mode delphi.MountMode) {
	client.MountKind("Tenant", mode)
}

func TenantMountKey(client clientApi.Client, key string, mode delphi.MountMode) {
	client.MountKindKey("Tenant", fmt.Sprintf("%v", key), mode)
}

func GetTenant(client clientApi.Client, key string) *Tenant {
	o := client.GetObject("Tenant", fmt.Sprintf("%v", key))
	if o == nil {
		return nil
	}
	obj, ok := o.(*Tenant)
	if ok != true {
		panic("Cast failed")
	}
	return obj
}

func (m *Tenant) IsPersistent() bool {
	return false
}
func TenantFactory(sdkClient clientApi.Client, data []byte) (clientApi.BaseObject, error) {
	var msg Tenant
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func TenantWatch(client clientApi.Client, reactor TenantReactor) {
	client.WatchKind("Tenant", reactor)
}
func TenantList(client clientApi.Client) []*Tenant {
	bobjs := client.List("Tenant")
	objs := make([]*Tenant, 0)
	for _, bobj := range bobjs {
		obj, _ := bobj.(*Tenant)
		objs = append(objs, obj)
	}
	return objs
}
func (m *Tenant) TriggerEvent(sdkClient clientApi.Client, old clientApi.BaseObject, op delphi.ObjectOperation, rl []clientApi.BaseReactor) {
	for _, r := range rl {
		rctr, ok := r.(TenantReactor)
		if ok == false {
			panic("Not a Reactor")
		}
		if op == delphi.ObjectOperation_SetOp {
			if old == nil {
				rctr.OnTenantCreate(m)
			} else {
				oldObj, ok := old.(*Tenant)
				if ok == false {
					panic("Not an Tenant object")
				}
				rctr.OnTenantUpdate(oldObj, m)
			}
		} else {
			rctr.OnTenantDelete(m)
		}
	}
}

type TenantReactor interface {
	OnTenantCreate(obj *Tenant)
	OnTenantUpdate(old *Tenant, obj *Tenant)
	OnTenantDelete(obj *Tenant)
}

func (m *Tenant) Reset()                    { *m = Tenant{} }
func (m *Tenant) String() string            { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()               {}
func (*Tenant) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Tenant) GetMeta() *delphi.ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Tenant) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tenant) GetTenant() *netproto10.Tenant {
	if m != nil {
		return m.Tenant
	}
	return nil
}

func init() {
	clientApi.RegisterFactory("Tenant", TenantFactory)
	proto.RegisterType((*Tenant)(nil), "dnetproto.Tenant")
}

func init() { proto.RegisterFile("delphi/tenant.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0xcd, 0x29,
	0xc8, 0xc8, 0xd4, 0x2f, 0x49, 0xcd, 0x4b, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4c, 0xc9, 0x4b, 0x2d, 0x01, 0x33, 0xa5, 0x6c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0x0b, 0x52, 0xf3, 0x8a, 0x13, 0xf3, 0x52, 0xf2, 0xf5, 0x8b, 0xcb, 0xf5,
	0xf3, 0x32, 0x93, 0xf5, 0x13, 0xd3, 0x53, 0xf3, 0x4a, 0xf4, 0xc1, 0xea, 0x8a, 0xf5, 0x61, 0x3a,
	0x50, 0x4c, 0x92, 0xe2, 0x81, 0x18, 0x0f, 0xe1, 0x29, 0x15, 0x71, 0xb1, 0x85, 0x80, 0x65, 0x85,
	0xd4, 0xb8, 0x58, 0x7c, 0x53, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf4,
	0xa0, 0xca, 0xfc, 0x93, 0xb2, 0x52, 0x93, 0x4b, 0x40, 0x32, 0x41, 0x60, 0x79, 0x21, 0x01, 0x2e,
	0x66, 0xef, 0xd4, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x03, 0x66,
	0x86, 0x04, 0x33, 0x58, 0xaf, 0x80, 0x1e, 0xcc, 0x66, 0x3d, 0x88, 0x78, 0x10, 0x54, 0xde, 0x8a,
	0xa5, 0x61, 0xba, 0x12, 0x63, 0x12, 0x1b, 0x58, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1e,
	0x21, 0x6d, 0xd8, 0xe9, 0x00, 0x00, 0x00,
}
