// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package watcher

import (
	"context"
	"testing"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/cluster"
	_ "github.com/pensando/sw/api/generated/exports/apiserver"
	"github.com/pensando/sw/api/generated/network"
	"github.com/pensando/sw/api/generated/security"
	_ "github.com/pensando/sw/api/hooks/apiserver"
	"github.com/pensando/sw/venice/apiserver"
	apisrvpkg "github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/ctrler/npm/statemgr"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/orch"
	vchserver "github.com/pensando/sw/venice/orch/vchub/server"
	vchstore "github.com/pensando/sw/venice/orch/vchub/store"
	"github.com/pensando/sw/venice/utils/debug"
	"github.com/pensando/sw/venice/utils/kvstore/store"
	kvs "github.com/pensando/sw/venice/utils/kvstore/store"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/netutils"
	"github.com/pensando/sw/venice/utils/runtime"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func createAPIServer(url string) (apiserver.Server, apiserver.Config) {
	logger := log.GetNewLogger(log.GetDefaultConfig("api_watcher_test"))

	// api server config
	sch := runtime.GetDefaultScheme()
	apisrvConfig := apiserver.Config{
		GrpcServerPort: url,
		Logger:         logger,
		Version:        "v1",
		Scheme:         sch,
		Kvstore: store.Config{
			Type:    store.KVStoreTypeMemkv,
			Servers: []string{""},
			Codec:   runtime.NewJSONCodec(sch),
		},
	}
	// create api server
	apiSrv := apisrvpkg.MustGetAPIServer()
	go apiSrv.Run(apisrvConfig)
	time.Sleep(time.Millisecond * 100)

	return apiSrv, apisrvConfig
}

func TestApiWatcher(t *testing.T) {
	// create network state manager
	stateMgr, err := statemgr.NewStatemgr(&dummyWriter{})
	if err != nil {
		t.Fatalf("Could not create network manager. Err: %v", err)
		return
	}

	// generate an available port for api server to use.
	apiSrvListener := netutils.TestListenAddr{}
	err = apiSrvListener.GetAvailablePort()
	if err != nil {
		t.Errorf("could not find an available port for the api server")
	}
	apisrvURL := apiSrvListener.ListenURL.String()

	// create api server
	apiSrv, _ := createAPIServer(apisrvURL)
	Assert(t, (apiSrv != nil), "Error creating api server", apiSrv)

	// create memkvstore
	kvs, err := vchstore.Init("vchub", kvs.KVStoreTypeMemkv)
	AssertOk(t, err, "Error initing kvstore")
	Assert(t, (kvs != nil), "Error creating kvstore")

	// generate an available port for vchub to use.
	vchubSrvListener := netutils.TestListenAddr{}
	err = vchubSrvListener.GetAvailablePort()
	if err != nil {
		t.Errorf("could not find an available port for the vchub")
	}
	vmmURL := vchubSrvListener.ListenURL.String()

	// create vchub
	vcs, err := vchserver.NewVCHServer(vmmURL)
	AssertOk(t, err, "Error creating vchub")
	time.Sleep(time.Millisecond * 10)

	// create watcher on api server
	watcher, err := NewWatcher(stateMgr, apisrvURL, vmmURL, nil, debug.New(t.Name()).Build())
	AssertOk(t, err, "Error creating watchr")
	Assert(t, (watcher != nil), "Error creating watcher", watcher)
	time.Sleep(time.Millisecond * 10)

	// create an api server client
	l := log.GetNewLogger(log.GetDefaultConfig("NpmApiWatcher"))
	apicl, err := apiclient.NewGrpcAPIClient(globals.Npm, apisrvURL, l)
	AssertOk(t, err, "Error creating api server client")

	// create a default tenant
	tenant := cluster.Tenant{
		TypeMeta: api.TypeMeta{Kind: "Tenant"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "testTenant",
			Name:   "testTenant",
		},
	}

	tn, err := apicl.ClusterV1().Tenant().Create(context.Background(), &tenant)
	AssertOk(t, err, "failed to create tenant")
	AssertEquals(t, "testTenant", tn.Name, "tenant names did not match")

	// verify the tenant got created
	AssertEventually(t, func() (bool, interface{}) {
		_, terr := stateMgr.FindTenant("testTenant")
		return terr == nil, nil
	}, "Tenant not found in statemgr")
	ts, err := stateMgr.FindTenant("testTenant")
	AssertOk(t, err, "Could not find the tenant")
	AssertEquals(t, "testTenant", ts.Name, "tenant names did not match")

	//create a network in api server
	net := network.Network{
		TypeMeta: api.TypeMeta{Kind: "network"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "testTenant",
			Name:   "testNetwork",
		},
		Spec: network.NetworkSpec{
			IPv4Subnet:  "10.1.1.0/24",
			IPv4Gateway: "10.1.1.254",
		},
	}
	ns, err := apicl.NetworkV1().Network().Create(context.Background(), &net)
	AssertOk(t, err, "Error creating network")
	Assert(t, (ns.Spec.IPv4Subnet == "10.1.1.0/24"), "Got invalid network", ns)

	// verify network got created
	AssertEventually(t, func() (bool, interface{}) {
		_, nerr := stateMgr.FindNetwork("testTenant", "testNetwork")
		return (nerr == nil), nil
	}, "Network not found in statemgr")
	nw, err := stateMgr.FindNetwork("testTenant", "testNetwork")
	AssertOk(t, err, "Could not find the network")
	Assert(t, (nw.Spec.IPv4Subnet == "10.1.1.0/24"), "Got invalid network", nw)

	// create a vmm nwif
	nwif := orch.NwIF{
		ObjectKind: "NwIF",
		ObjectMeta: &api.ObjectMeta{
			Tenant: "testTenant",
			Name:   "test-nwif",
			UUID:   "test-nwif",
		},
		Config: &orch.NwIF_Config{
			LocalVLAN: 22,
		},
		Status: &orch.NwIF_Status{
			IpAddress:   "10.1.1.1",
			MacAddress:  "11:11:11:11:11:11",
			PortGroup:   "test",
			Network:     "testNetwork",
			SmartNIC_ID: "test-host",
		},
	}
	err = vchstore.NwIFCreate(context.Background(), "test-nwif", &nwif)
	AssertOk(t, err, "Error creating nw if")

	// verify endpoint got created
	AssertEventually(t, func() (bool, interface{}) {
		_, perr := stateMgr.FindEndpoint("testTenant", "test-nwif")
		return (perr == nil), nil
	}, "Endpoint not found in statemgr")
	ep, err := stateMgr.FindEndpoint("testTenant", "test-nwif")
	AssertOk(t, err, "Could not find the endpoint")
	Assert(t, (ep.Status.MicroSegmentVlan == 22), "Endpoint did not match", ep)

	// delete the nwif
	err = vchstore.NwIFDelete(context.Background(), "test-nwif")
	AssertOk(t, err, "Error deleting nw if")

	// verify endpoint is gone
	AssertEventually(t, func() (bool, interface{}) {
		_, perr := stateMgr.FindEndpoint("testTenant", "test-nwif")
		return (perr != nil), nil
	}, "Endpoint still found in statemgr")

	// delete the tenant
	_, err = apicl.ClusterV1().Tenant().Delete(context.Background(), &tenant.ObjectMeta)
	AssertOk(t, err, "Error deleting tenant")

	// delete the network
	_, err = apicl.NetworkV1().Network().Delete(context.Background(), &net.ObjectMeta)
	AssertOk(t, err, "Error deleting network")

	// verify network is gone
	AssertEventually(t, func() (bool, interface{}) {
		_, nerr := stateMgr.FindNetwork("testTenant", "test")
		return (nerr != nil), nil
	}, "Endpoint still found in statemgr")

	// stop the api server
	apicl.Close()
	watcher.Stop()
	apiSrv.Stop()
	vcs.StopServer()
	kvs.Close()
}

func TestAPIServerRestarts(t *testing.T) {
	// create network state manager
	stateMgr, err := statemgr.NewStatemgr(&dummyWriter{})
	if err != nil {
		t.Fatalf("Could not create network manager. Err: %v", err)
		return
	}

	// generate an available port for api server to use.
	apiSrvListener := netutils.TestListenAddr{}
	err = apiSrvListener.GetAvailablePort()
	if err != nil {
		t.Errorf("could not find an available port for the api server")
	}
	apisrvURL := apiSrvListener.ListenURL.String()

	// create api server
	apiSrv, _ := createAPIServer(apisrvURL)
	Assert(t, (apiSrv != nil), "Error creating api server", apiSrv)

	// create memkvstore
	kvs, err := vchstore.Init("vchub", kvs.KVStoreTypeMemkv)
	AssertOk(t, err, "Error initing kvstore")
	Assert(t, (kvs != nil), "Error creating kvstore")

	// generate an available port for vchub to use.
	vchubSrvListener := netutils.TestListenAddr{}
	err = vchubSrvListener.GetAvailablePort()
	if err != nil {
		t.Errorf("could not find an available port for the vchub")
	}
	vmmURL := vchubSrvListener.ListenURL.String()

	// create vchub
	vcs, err := vchserver.NewVCHServer(vmmURL)
	AssertOk(t, err, "Error creating vchub")
	time.Sleep(time.Millisecond * 10)

	// create watcher on api server
	watcher, err := NewWatcher(stateMgr, apisrvURL, vmmURL, nil, debug.New(t.Name()).Build())
	AssertOk(t, err, "Error creating watchr")
	Assert(t, (watcher != nil), "Error creating watcher", watcher)
	time.Sleep(time.Millisecond * 10)

	// create an api server client
	l := log.GetNewLogger(log.GetDefaultConfig("NpmApiWatcher"))
	apicl, err := apiclient.NewGrpcAPIClient(globals.Npm, apisrvURL, l)
	AssertOk(t, err, "Error creating api server client")

	// create a default tenant
	tenant := cluster.Tenant{
		TypeMeta: api.TypeMeta{Kind: "Tenant"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "testTenant",
			Name:   "testTenant",
		},
	}

	tn, err := apicl.ClusterV1().Tenant().Create(context.Background(), &tenant)
	AssertOk(t, err, "failed to create tenant")
	AssertEquals(t, "testTenant", tn.Name, "tenant names did not match")

	// verify the tenant got created
	AssertEventually(t, func() (bool, interface{}) {
		_, terr := stateMgr.FindTenant("testTenant")
		return terr == nil, nil
	}, "Tenant not found in statemgr")
	ts, err := stateMgr.FindTenant("testTenant")
	AssertOk(t, err, "Could not find the tenant")
	AssertEquals(t, "testTenant", ts.Name, "tenant names did not match")

	//create a network in api server
	net := network.Network{
		TypeMeta: api.TypeMeta{Kind: "network"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "testTenant",
			Name:   "testNetwork",
		},
		Spec: network.NetworkSpec{
			IPv4Subnet:  "10.1.1.0/24",
			IPv4Gateway: "10.1.1.254",
		},
	}
	ns, err := apicl.NetworkV1().Network().Create(context.Background(), &net)
	AssertOk(t, err, "Error creating network")
	Assert(t, (ns.Spec.IPv4Subnet == "10.1.1.0/24"), "Got invalid network", ns)

	// verify network got created
	AssertEventually(t, func() (bool, interface{}) {
		_, nerr := stateMgr.FindNetwork("testTenant", "testNetwork")
		return (nerr == nil), nil
	}, "Network not found in statemgr")
	nw, err := stateMgr.FindNetwork("testTenant", "testNetwork")
	AssertOk(t, err, "Could not find the network")
	Assert(t, (nw.Spec.IPv4Subnet == "10.1.1.0/24"), "Got invalid network", nw)
	//
	// create a vmm nwif
	nwif := orch.NwIF{
		ObjectKind: "NwIF",
		ObjectMeta: &api.ObjectMeta{
			Tenant: "testTenant",
			Name:   "test-nwif",
			UUID:   "test-nwif",
		},
		Config: &orch.NwIF_Config{
			LocalVLAN: 22,
		},
		Status: &orch.NwIF_Status{
			IpAddress:   "10.1.1.1",
			MacAddress:  "11:11:11:11:11:11",
			PortGroup:   "test",
			Network:     "testNetwork",
			SmartNIC_ID: "test-host",
		},
	}
	err = vchstore.NwIFCreate(context.Background(), "test-nwif", &nwif)
	AssertOk(t, err, "Error creating nw if")

	// verify endpoint got created
	AssertEventually(t, func() (bool, interface{}) {
		_, perr := stateMgr.FindEndpoint("testTenant", "test-nwif")
		return (perr == nil), nil
	}, "Endpoint not found in statemgr")
	ep, err := stateMgr.FindEndpoint("testTenant", "test-nwif")
	AssertOk(t, err, "Could not find the endpoint")
	Assert(t, (ep.Status.MicroSegmentVlan == 22), "Endpoint did not match", ep)

	// create sg policy
	inrules := []security.SGRule{
		{
			Ports:  "tcp/80",
			Action: "Allow",
		},
	}
	outrules := []security.SGRule{
		{
			Ports:  "tcp/80",
			Action: "Allow",
		},
	}
	sgp := security.Sgpolicy{
		TypeMeta: api.TypeMeta{Kind: "Sgpolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testpolicy",
		},
		Spec: security.SgpolicySpec{
			InRules:  inrules,
			OutRules: outrules,
		},
	}
	sgps, err := apicl.SecurityV1().Sgpolicy().Create(context.Background(), &sgp)
	AssertOk(t, err, "Error creating security policy")
	AssertEquals(t, inrules, sgps.Spec.InRules, "inrules did not match")
	AssertEquals(t, outrules, sgps.Spec.OutRules, "outrules did not match")

	// stop api server and watchers
	apiSrv.Stop()
	watcher.Stop()
	apicl.Close()

	// restart api server and watchers
	apiSrv, _ = createAPIServer(apisrvURL)
	Assert(t, (apiSrv != nil), "Error restarting api server", apiSrv)
	apicl, err = apiclient.NewGrpcAPIClient(globals.Npm, apisrvURL, l)
	AssertOk(t, err, "Error restarting api server client")
	watcher, err = NewWatcher(stateMgr, apisrvURL, vmmURL, nil, debug.New(t.Name()).Build())
	AssertOk(t, err, "Error restarting watcher")

	// delete all the objects
	_, err = apicl.ClusterV1().Tenant().Delete(context.Background(), &tenant.ObjectMeta)
	AssertOk(t, err, "could not delete tenant")

	_, err = apicl.NetworkV1().Network().Delete(context.Background(), &net.ObjectMeta)
	AssertOk(t, err, "could not delete network")

	err = vchstore.NwIFDelete(context.Background(), "test-nwif")
	AssertOk(t, err, "could not delete ep")

	_, err = apicl.SecurityV1().Sgpolicy().Delete(context.Background(), &sgp.ObjectMeta)
	AssertOk(t, err, "could not delete sg policy")

	// stop the api server
	apicl.Close()
	watcher.Stop()
	apiSrv.Stop()
	vcs.StopServer()
	kvs.Close()
}

func TestApiServerClient(t *testing.T) {
	testCount := 5
	for i := 0; i < testCount; i++ {
		// generate an available port for api server to use.
		apiSrvListener := netutils.TestListenAddr{}
		err := apiSrvListener.GetAvailablePort()
		if err != nil {
			t.Errorf("could not find an available port for the api server")
		}
		apisrvURL := apiSrvListener.ListenURL.String()
		// create api server
		apiSrv, _ := createAPIServer(apisrvURL)
		Assert(t, (apiSrv != nil), "Error creating api server", apiSrv)

		// create an api server client
		l := log.GetNewLogger(log.GetDefaultConfig("NpmApiWatcher"))
		apicl, err := apiclient.NewGrpcAPIClient(globals.Npm, apisrvURL, l)
		AssertOk(t, err, "Error creating api server client")

		// network watcher
		opts := api.ListWatchOptions{}
		watchCtx, watchCancel := context.WithCancel(context.Background())
		netWatcher, err := apicl.NetworkV1().Network().Watch(watchCtx, &opts)
		AssertOk(t, err, "Error creating api server watcher")
		Assert(t, netWatcher != nil, "Netwatcher is nil")

		// stop api client & server
		watchCancel()
		apicl.Close()
		apiSrv.Stop()
		netWatcher.Stop()
	}
}

func TestApiWatcherConnectDisconnect(t *testing.T) {
	// create network state manager
	stateMgr, err := statemgr.NewStatemgr(&dummyWriter{})
	if err != nil {
		t.Fatalf("Could not create network manager. Err: %v", err)
		return
	}

	// create memkvstore
	kvs, err := vchstore.Init("", kvs.KVStoreTypeMemkv)
	AssertOk(t, err, "Error initing kvstore")
	Assert(t, (kvs != nil), "Error creating kvstore")

	// generate an available port for vchub to use.
	vchubSrvListener := netutils.TestListenAddr{}
	err = vchubSrvListener.GetAvailablePort()
	if err != nil {
		t.Errorf("could not find an available port for the vchub")
	}
	vmmURL := vchubSrvListener.ListenURL.String()

	// create watcher on api server
	watcher, err := NewWatcher(stateMgr, "", vmmURL, nil, debug.New(t.Name()).Build())
	AssertOk(t, err, "Error creating watchr")
	Assert(t, (watcher != nil), "Error creating watcher", watcher)
	time.Sleep(time.Millisecond * 10)

	// create vchub
	vcs, err := vchserver.NewVCHServer(vmmURL)
	AssertOk(t, err, "Error creating vchub")

	// stop the server and restart it
	vcs.StopServer()
	time.Sleep(time.Second * 3)
	vcs, err = vchserver.NewVCHServer(vmmURL)
	AssertOk(t, err, "Error creating vchub")
	time.Sleep(time.Millisecond * 10)

	// stop the vchub
	vcs.StopServer()
	watcher.Stop()
	kvs.Close()
}
