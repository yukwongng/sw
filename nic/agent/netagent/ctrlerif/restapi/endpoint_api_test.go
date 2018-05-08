// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: endpoint.proto
*/
package restapi

import (
	"testing"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func TestEndpointList(t *testing.T) {
	t.Parallel()
	var ok bool
	var endpointList []*netproto.Endpoint

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/endpoints/", &endpointList)

	AssertOk(t, err, "Error getting endpoints from the REST Server")
	for _, o := range endpointList {
		if o.Name == "preCreatedEndpoint" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedEndpoint in Response: %v", endpointList)
	}

}

func TestEndpointPost(t *testing.T) {
	t.Parallel()
	var resp error
	var ok bool
	var endpointList []*netproto.Endpoint

	postData := netproto.Endpoint{
		TypeMeta: api.TypeMeta{Kind: "Endpoint"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testPostEndpoint",
		},
		Spec: netproto.EndpointSpec{
			EndpointUUID: "testEndpointUUID",
			WorkloadUUID: "testWorkloadUUID",
			NetworkName:  "preCreatedNetwork",
		},
		Status: netproto.EndpointStatus{
			NodeUUID:    "dummy-node-uuid",
			IPv4Address: "10.1.1.0/24",
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/endpoints/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/endpoints/", &endpointList)

	AssertOk(t, err, "Error posting endpoint to REST Server")
	AssertOk(t, getErr, "Error getting endpoints from the REST Server")
	for _, o := range endpointList {
		if o.Name == "testPostEndpoint" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostEndpoint in Response: %v", endpointList)
	}

}

func TestEndpointUpdate(t *testing.T) {
	t.Parallel()
	var resp error
	var endpointList []*netproto.Endpoint

	var actualEndpointSpec netproto.EndpointSpec
	updatedEndpointSpec := netproto.EndpointSpec{
		EndpointUUID: "testEndpointUUID",
		WorkloadUUID: "updatedWorkloadUUID",
		NetworkName:  "preCreatedNetwork",
	}
	putData := netproto.Endpoint{
		TypeMeta: api.TypeMeta{Kind: "Endpoint"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedEndpoint",
		},
		Spec: updatedEndpointSpec,
		Status: netproto.EndpointStatus{
			NodeUUID:    "dummy-node-uuid",
			IPv4Address: "10.1.1.0/24",
		},
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/endpoints/default/default/preCreatedEndpoint", &putData, &resp)
	AssertOk(t, err, "Error updating endpoint to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/endpoints/", &endpointList)
	AssertOk(t, getErr, "Error getting endpoints from the REST Server")

	for _, o := range endpointList {
		if o.Name == "preCreatedEndpoint" {
			actualEndpointSpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedEndpointSpec, actualEndpointSpec, "Could not validate updated spec.")

}

func TestEndpointDelete(t *testing.T) {
	t.Parallel()
	var resp error
	var found bool
	var endpointList []*netproto.Endpoint

	deleteData := netproto.Endpoint{
		TypeMeta: api.TypeMeta{Kind: "Endpoint"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testDeleteEndpoint",
		},
		Spec: netproto.EndpointSpec{
			EndpointUUID: "testEndpointUUID",
			WorkloadUUID: "testWorkloadUUID",
			NetworkName:  "preCreatedNetwork",
		},
		Status: netproto.EndpointStatus{
			NodeUUID:    "dummy-node-uuid",
			IPv4Address: "10.1.1.0/24",
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/endpoints/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/endpoints/default/default/testDeleteEndpoint", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/endpoints/", &endpointList)

	AssertOk(t, postErr, "Error posting endpoint to REST Server")
	AssertOk(t, err, "Error deleting endpoint from REST Server")
	AssertOk(t, getErr, "Error getting endpoints from the REST Server")
	for _, o := range endpointList {
		if o.Name == "testDeleteEndpoint" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteEndpoint in Response after deleting: %v", endpointList)
	}

}
