// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package network is a auto generated package.
Input file: protos/network.proto
*/
package network

import (
	fmt "fmt"
	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

// Dummy definitions to suppress nonused warnings
var _ kvstore.Interface
var _ log.Logger
var _ listerwatcher.WatcherClient

// MakeKey generates a KV store key for the object
func (m *Endpoint) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "endpoint/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *LbPolicy) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "lb-policy/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Network) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "network/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *SecurityGroup) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "security-group/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Service) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "service/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Sgpolicy) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "sgpolicy/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *Tenant) MakeKey(prefix string) string {
	return fmt.Sprint("/venice/", prefix, "/", "tenant/", m.Name)
}

// MakeKey generates a KV store key for the object
func (m *EndpointList) MakeKey(prefix string) string {
	obj := Endpoint{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *LbPolicyList) MakeKey(prefix string) string {
	obj := LbPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *NetworkList) MakeKey(prefix string) string {
	obj := Network{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *SecurityGroupList) MakeKey(prefix string) string {
	obj := SecurityGroup{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *ServiceList) MakeKey(prefix string) string {
	obj := Service{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *SgpolicyList) MakeKey(prefix string) string {
	obj := Sgpolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *TenantList) MakeKey(prefix string) string {
	obj := Tenant{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgEndpointWatchHelper) MakeKey(prefix string) string {
	obj := Endpoint{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgLbPolicyWatchHelper) MakeKey(prefix string) string {
	obj := LbPolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgNetworkWatchHelper) MakeKey(prefix string) string {
	obj := Network{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgSecurityGroupWatchHelper) MakeKey(prefix string) string {
	obj := SecurityGroup{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgServiceWatchHelper) MakeKey(prefix string) string {
	obj := Service{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgSgpolicyWatchHelper) MakeKey(prefix string) string {
	obj := Sgpolicy{}
	return obj.MakeKey(prefix)
}

// MakeKey generates a KV store key for the object
func (m *AutoMsgTenantWatchHelper) MakeKey(prefix string) string {
	obj := Tenant{}
	return obj.MakeKey(prefix)
}
