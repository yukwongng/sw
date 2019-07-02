// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

// Package npm contains network policy manager code
/************ Architecture of NetCtrler ******************
 *
 *    +-----------------------+
 *    |   API Server          |
 *    +-----------------------+
 *        |             ^
 *        V             |
 *    +-----------------------+
 *    |     Controller kit    |
 *  +-+-----------------------+-------------+
 *  |        Network State Manager          |
 *  | +---------+ +----------+ +----------+ |
 *  | | Network | | Endpoint | | SecGroup | |
 *  | +---------+ +----------+ +----------+ |
 *  +---------------------------------------+
 *  |            Nimbus Server              |
 *  +---------------------------------------+
 *       ^           ^            ^
 *       |           |            |
 *       V           V            V
 *  +--------+   +--------+   +--------+
 *  | Agent  |   | Agent  |   | Agent  |
 *  | Node 1 |   | Node 2 |   | Node 3 |
 *  +--------+   +--------+   +--------+
 *
 *********************************************************/
package npm

import (
	"context"
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	diagapi "github.com/pensando/sw/api/generated/diagnostics"
	"github.com/pensando/sw/nic/agent/protos/generated/nimbus"
	"github.com/pensando/sw/venice/ctrler/npm/statemgr"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils"
	debugStats "github.com/pensando/sw/venice/utils/debug/stats"
	diagsvc "github.com/pensando/sw/venice/utils/diagnostics/service"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"
	"github.com/pensando/sw/venice/utils/rpckit"
	"github.com/pensando/sw/venice/utils/tsdb"
)

// Netctrler is a netctrler instance
type Netctrler struct {
	StateMgr   *statemgr.Statemgr // state manager
	mserver    *nimbus.MbusServer // nimbus server
	debugStats *debugStats.Stats
	rpcServer  *rpckit.RPCServer
}

// NewNetctrler returns a controller instance
func NewNetctrler(serverURL, restURL, apisrvURL, vmmURL string, resolver resolver.Interface, logger log.Logger, enableDiagnostics bool) (*Netctrler, error) {

	// init tsdb client
	tsdbOpts := &tsdb.Opts{
		ClientName:              globals.Npm,
		ResolverClient:          resolver,
		Collector:               globals.Collector,
		DBName:                  "default",
		SendInterval:            time.Duration(30) * time.Second,
		ConnectionRetryInterval: 100 * time.Millisecond,
		StartLocalServer:        true,
	}
	tsdb.Init(context.Background(), tsdbOpts)

	// create an RPC server
	rpcServer, err := rpckit.NewRPCServer(globals.Npm, serverURL)
	if err != nil {
		logger.Fatalf("Error creating rpc server. Err; %v", err)
	}
	// create nimbus server
	msrv := nimbus.NewMbusServer(globals.Npm, rpcServer)

	var options []statemgr.Option
	if enableDiagnostics {
		options = append(options,
			statemgr.WithDiagnosticsHandler("Debug", diagapi.DiagnosticsRequest_Log.String(), diagsvc.NewElasticLogsHandler(globals.Npm, utils.GetHostname(), diagapi.ModuleStatus_Venice, resolver, logger)),
			statemgr.WithDiagnosticsHandler("Debug", diagapi.DiagnosticsRequest_Stats.String(), diagsvc.NewExpVarHandler(globals.Npm, utils.GetHostname(), diagapi.ModuleStatus_Venice, logger)))
	}
	// create network state manager
	stateMgr, err := statemgr.NewStatemgr(rpcServer, apisrvURL, resolver, msrv, logger, options...)
	if err != nil {
		logger.Errorf("Could not create network manager. Err: %v", err)
		return nil, err
	}

	logger.Infof("RPC server is running at %v", serverURL)

	// create the controller instance
	ctrler := Netctrler{
		StateMgr:  stateMgr,
		mserver:   msrv,
		rpcServer: rpcServer,
	}

	// start the RPC server
	rpcServer.Start()

	// start debug REST server
	if restURL != "" {
		ctrler.runDebugRESTServer(restURL)
	}

	return &ctrler, nil
}

// Stop server and release resources
func (c *Netctrler) Stop() error {
	if c.rpcServer != nil {
		c.rpcServer.Stop()
	}

	c.StateMgr.Stop()
	return nil
}

// runDebugRESTServer starts REST server for debug purposes
func (c *Netctrler) runDebugRESTServer(restURL string) error {
	// start a debug http server
	router := mux.NewRouter()
	router.Methods("GET").Subrouter().Handle("/debug/vars", expvar.Handler())
	router.Methods("GET").Subrouter().Handle("/nimbus/db/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jData, err := c.mserver.DumpDatabase()
		if err != nil {
			log.Errorf("Error dumping nimbus server database. Err: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}))
	router.Methods("GET").Subrouter().Handle("/ctkit/db/{kind}/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		kind, ok := vars["kind"]
		if !ok {
			http.Error(w, fmt.Sprintf("kind not specified"), http.StatusInternalServerError)
			return
		}
		objlist := c.StateMgr.ListObjects(kind)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(objlist)
	}))

	// start the server
	go http.ListenAndServe(restURL, router)

	return nil
}
