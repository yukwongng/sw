// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: nat.proto
*/
package restapi

import (
	"testing"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/protos/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func TestNatBindingList(t *testing.T) {
	t.Parallel()
	var ok bool
	var natbindingList []*netproto.NatBinding

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/bindings/", &natbindingList)

	AssertOk(t, err, "Error getting natbindings from the REST Server")
	for _, o := range natbindingList {
		if o.Name == "preCreatedNatBinding" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedNatBinding in Response: %v", natbindingList)
	}

}

func TestNatBindingPost(t *testing.T) {
	t.Parallel()
	var resp Response
	var ok bool
	var natbindingList []*netproto.NatBinding

	postData := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testPostNatBinding",
		},
		Spec: netproto.NatBindingSpec{
			NatPoolName: "preCreatedNatPool",
			IPAddress:   "10.1.1.2",
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/bindings/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/bindings/", &natbindingList)

	AssertOk(t, err, "Error posting natbinding to REST Server")
	AssertOk(t, getErr, "Error getting natbindings from the REST Server")
	for _, o := range natbindingList {
		if o.Name == "testPostNatBinding" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostNatBinding in Response: %v", natbindingList)
	}

}

func TestNatBindingUpdate(t *testing.T) {
	t.Parallel()
	var resp Response
	var natbindingList []*netproto.NatBinding

	var actualNatBindingSpec netproto.NatBindingSpec
	updatedNatBindingSpec := netproto.NatBindingSpec{
		NatPoolName: "preCreatedNatPool",
		IPAddress:   "192.168.1.2",
	}
	putData := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Name:      "preCreatedNatBinding",
			Namespace: "default",
		},
		Spec: updatedNatBindingSpec,
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/nat/bindings/default/default/preCreatedNatBinding", &putData, &resp)
	AssertOk(t, err, "Error updating natbinding to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/bindings/", &natbindingList)
	AssertOk(t, getErr, "Error getting natbindings from the REST Server")
	for _, o := range natbindingList {
		if o.Name == "preCreatedNatBinding" {
			actualNatBindingSpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedNatBindingSpec, actualNatBindingSpec, "Could not validate updated spec.")

}

func TestNatBindingDelete(t *testing.T) {
	t.Parallel()
	var resp Response
	var found bool
	var natbindingList []*netproto.NatBinding

	deleteData := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testDeleteNatBinding",
		},
		Spec: netproto.NatBindingSpec{
			NatPoolName: "preCreatedNatPool",
			IPAddress:   "10.1.1.1",
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/bindings/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/nat/bindings/default/default/testDeleteNatBinding", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/bindings/", &natbindingList)

	AssertOk(t, postErr, "Error posting natbinding to REST Server")
	AssertOk(t, err, "Error deleting natbinding from REST Server")
	AssertOk(t, getErr, "Error getting natbindings from the REST Server")
	for _, o := range natbindingList {
		if o.Name == "testDeleteNatBinding" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteNatBinding in Response after deleting: %v", natbindingList)
	}

}

func TestNatBindingCreateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badPostData := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{
			Name: "",
		},
	}

	err := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/bindings/", &badPostData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatBindingDeleteErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/nat/bindings/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatBindingUpdateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPPut("http://"+agentRestURL+"/api/nat/bindings/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatPolicyList(t *testing.T) {
	t.Parallel()
	var ok bool
	var natpolicyList []*netproto.NatPolicy

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/policies/", &natpolicyList)

	AssertOk(t, err, "Error getting natpolicys from the REST Server")
	for _, o := range natpolicyList {
		if o.Name == "preCreatedNatPolicy" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedNatPolicy in Response: %v", natpolicyList)
	}

}

func TestNatPolicyPost(t *testing.T) {
	t.Parallel()
	var resp Response
	var ok bool
	var natpolicyList []*netproto.NatPolicy

	postData := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testPostNatPolicy",
		},
		Spec: netproto.NatPolicySpec{
			Rules: []netproto.NatRule{
				{
					Src: &netproto.MatchSelector{
						Addresses: []string{"10.0.0.0 - 10.0.1.0", "172.17.0.0/24", "4.4.4.4"},
					},
					Dst: &netproto.MatchSelector{
						Addresses: []string{"192.168.0.1 - 192.168.1.0", "8.8.8.8/8"},
					},
					NatPool: "preCreatedNatPool",
					Action:  "SNAT",
				},
			},
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/policies/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/policies/", &natpolicyList)

	AssertOk(t, err, "Error posting natpolicy to REST Server")
	AssertOk(t, getErr, "Error getting natpolicys from the REST Server")
	for _, o := range natpolicyList {
		if o.Name == "testPostNatPolicy" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostNatPolicy in Response: %v", natpolicyList)
	}

}

func TestNatPolicyUpdate(t *testing.T) {
	t.Parallel()
	var resp Response
	var natpolicyList []*netproto.NatPolicy

	var actualNatPolicySpec netproto.NatPolicySpec
	updatedNatPolicySpec := netproto.NatPolicySpec{
		Rules: []netproto.NatRule{
			{
				NatPool: "updatedNatPool",
			},
		},
	}
	putData := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Name:      "preCreatedNatPolicy",
			Namespace: "default",
		},
		Spec: updatedNatPolicySpec,
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/nat/policies/default/default/preCreatedNatPolicy", &putData, &resp)
	AssertOk(t, err, "Error updating natpolicy to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/policies/", &natpolicyList)
	AssertOk(t, getErr, "Error getting natpolicys from the REST Server")
	for _, o := range natpolicyList {
		if o.Name == "preCreatedNatPolicy" {
			actualNatPolicySpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedNatPolicySpec, actualNatPolicySpec, "Could not validate updated spec.")

}

func TestNatPolicyDelete(t *testing.T) {
	t.Parallel()
	var resp Response
	var found bool
	var natpolicyList []*netproto.NatPolicy

	deleteData := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testDeleteNatPolicy",
		},
		Spec: netproto.NatPolicySpec{
			Rules: []netproto.NatRule{
				{
					Src: &netproto.MatchSelector{
						Addresses: []string{"10.0.0.0 - 10.0.1.0", "172.17.0.0/24", "4.4.4.4"},
					},
					Dst: &netproto.MatchSelector{
						Addresses: []string{"192.168.0.1 - 192.168.1.0", "8.8.8.8"},
					},
					NatPool: "preCreatedNatPool",
					Action:  "SNAT",
				},
			},
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/policies/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/nat/policies/default/default/testDeleteNatPolicy", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/policies/", &natpolicyList)

	AssertOk(t, postErr, "Error posting natpolicy to REST Server")
	AssertOk(t, err, "Error deleting natpolicy from REST Server")
	AssertOk(t, getErr, "Error getting natpolicys from the REST Server")
	for _, o := range natpolicyList {
		if o.Name == "testDeleteNatPolicy" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteNatPolicy in Response after deleting: %v", natpolicyList)
	}

}

func TestNatPolicyCreateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badPostData := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{
			Name: "",
		},
	}

	err := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/policies/", &badPostData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatPolicyDeleteErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/nat/policies/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatPolicyUpdateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPPut("http://"+agentRestURL+"/api/nat/policies/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatPoolList(t *testing.T) {
	t.Parallel()
	var ok bool
	var natpoolList []*netproto.NatPool

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/pools/", &natpoolList)

	AssertOk(t, err, "Error getting natpools from the REST Server")
	for _, o := range natpoolList {
		if o.Name == "preCreatedNatPool" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedNatPool in Response: %v", natpoolList)
	}

}

func TestNatPoolPost(t *testing.T) {
	t.Parallel()
	var resp Response
	var ok bool
	var natpoolList []*netproto.NatPool

	postData := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testPostNatPool",
		},
		Spec: netproto.NatPoolSpec{
			IPRange: "10.1.2.1-10.1.2.200",
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/pools/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/pools/", &natpoolList)

	AssertOk(t, err, "Error posting natpool to REST Server")
	AssertOk(t, getErr, "Error getting natpools from the REST Server")
	for _, o := range natpoolList {
		if o.Name == "testPostNatPool" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostNatPool in Response: %v", natpoolList)
	}

}

func TestNatPoolUpdate(t *testing.T) {
	t.Parallel()
	var resp Response
	var natpoolList []*netproto.NatPool

	var actualNatPoolSpec netproto.NatPoolSpec
	updatedNatPoolSpec := netproto.NatPoolSpec{
		IPRange: "192.168.1.1-198.168.1.200",
	}
	putData := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Name:      "preCreatedNatPool",
			Namespace: "default",
		},
		Spec: updatedNatPoolSpec,
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/nat/pools/default/default/preCreatedNatPool", &putData, &resp)
	AssertOk(t, err, "Error updating natpool to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/pools/", &natpoolList)
	AssertOk(t, getErr, "Error getting natpools from the REST Server")
	for _, o := range natpoolList {
		if o.Name == "preCreatedNatPool" {
			actualNatPoolSpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedNatPoolSpec, actualNatPoolSpec, "Could not validate updated spec.")

}

func TestNatPoolDelete(t *testing.T) {
	t.Parallel()
	var resp Response
	var found bool
	var natpoolList []*netproto.NatPool

	deleteData := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testDeleteNatPool",
		},
		Spec: netproto.NatPoolSpec{
			IPRange: "10.1.2.1-10.1.2.200",
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/pools/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/nat/pools/default/default/testDeleteNatPool", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/nat/pools/", &natpoolList)

	AssertOk(t, postErr, "Error posting natpool to REST Server")
	AssertOk(t, err, "Error deleting natpool from REST Server")
	AssertOk(t, getErr, "Error getting natpools from the REST Server")
	for _, o := range natpoolList {
		if o.Name == "testDeleteNatPool" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteNatPool in Response after deleting: %v", natpoolList)
	}

}

func TestNatPoolCreateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badPostData := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{
			Name: "",
		},
	}

	err := netutils.HTTPPost("http://"+agentRestURL+"/api/nat/pools/", &badPostData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatPoolDeleteErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/nat/pools/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestNatPoolUpdateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPPut("http://"+agentRestURL+"/api/nat/pools/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}
