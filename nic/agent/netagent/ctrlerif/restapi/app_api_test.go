// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: app.proto
*/
package restapi

import (
	"testing"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/protos/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func TestAppList(t *testing.T) {
	t.Parallel()
	var ok bool
	var appList []*netproto.App

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/apps/", &appList)

	AssertOk(t, err, "Error getting apps from the REST Server")
	for _, o := range appList {
		if o.Name == "preCreatedApp" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedApp in Response: %v", appList)
	}

}

func TestAppPost(t *testing.T) {
	t.Parallel()
	var resp Response
	var ok bool
	var appList []*netproto.App

	postData := netproto.App{
		TypeMeta: api.TypeMeta{Kind: "App"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testPostApp",
		},
		Spec: netproto.AppSpec{
			ALG: &netproto.ALG{
				DNS: &netproto.DNS{
					DropMultiZonePackets:   true,
					DropLargeDomainPackets: true,
					QueryResponseTimeout:   "30s",
				},
			},
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/apps/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/apps/", &appList)

	AssertOk(t, err, "Error posting app to REST Server")
	AssertOk(t, getErr, "Error getting apps from the REST Server")
	for _, o := range appList {
		if o.Name == "testPostApp" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostApp in Response: %v", appList)
	}

}

func TestAppUpdate(t *testing.T) {
	t.Parallel()
	var resp Response
	var appList []*netproto.App

	var actualAppSpec netproto.AppSpec
	updatedAppSpec := netproto.AppSpec{
		ALG: &netproto.ALG{
			DNS: &netproto.DNS{
				DropMultiZonePackets:   true,
				DropLargeDomainPackets: true,
				QueryResponseTimeout:   "30s",
			},
		},
	}
	putData := netproto.App{
		TypeMeta: api.TypeMeta{Kind: "App"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Name:      "preCreatedApp",
			Namespace: "default",
		},
		Spec: updatedAppSpec,
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/apps/default/default/preCreatedApp", &putData, &resp)
	AssertOk(t, err, "Error updating app to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/apps/", &appList)
	AssertOk(t, getErr, "Error getting apps from the REST Server")
	for _, o := range appList {
		if o.Name == "preCreatedApp" {
			actualAppSpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedAppSpec, actualAppSpec, "Could not validate updated spec.")

}

func TestAppDelete(t *testing.T) {
	t.Parallel()
	var resp Response
	var found bool
	var appList []*netproto.App

	deleteData := netproto.App{
		TypeMeta: api.TypeMeta{Kind: "App"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testDeleteApp",
		},
		Spec: netproto.AppSpec{
			ALG: &netproto.ALG{
				DNS: &netproto.DNS{
					DropMultiZonePackets:   true,
					DropLargeDomainPackets: true,
					QueryResponseTimeout:   "30s",
				},
			},
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/apps/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/apps/default/default/testDeleteApp", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/apps/", &appList)

	AssertOk(t, postErr, "Error posting app to REST Server")
	AssertOk(t, err, "Error deleting app from REST Server")
	AssertOk(t, getErr, "Error getting apps from the REST Server")
	for _, o := range appList {
		if o.Name == "testDeleteApp" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteApp in Response after deleting: %v", appList)
	}

}

func TestAppCreateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badPostData := netproto.App{
		TypeMeta: api.TypeMeta{Kind: "App"},
		ObjectMeta: api.ObjectMeta{
			Name: "",
		},
	}

	err := netutils.HTTPPost("http://"+agentRestURL+"/api/apps/", &badPostData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestAppDeleteErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.App{
		TypeMeta: api.TypeMeta{Kind: "App"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/apps/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestAppUpdateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.App{
		TypeMeta: api.TypeMeta{Kind: "App"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPPut("http://"+agentRestURL+"/api/apps/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}
