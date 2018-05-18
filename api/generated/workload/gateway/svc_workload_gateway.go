// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package workloadGwService is a auto generated package.
Input file: svc_workload.proto
*/
package workloadGwService

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pensando/grpc-gateway/runtime"

	"github.com/pensando/sw/api"
	workload "github.com/pensando/sw/api/generated/workload"
	grpcclient "github.com/pensando/sw/api/generated/workload/grpc/client"
	"github.com/pensando/sw/venice/apigw"
	"github.com/pensando/sw/venice/apigw/pkg"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/balancer"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta

type sWorkloadV1GwService struct {
	logger     log.Logger
	defSvcProf apigw.ServiceProfile
	svcProf    map[string]apigw.ServiceProfile
}

type adapterWorkloadV1 struct {
	conn    *rpckit.RPCClient
	service workload.ServiceWorkloadV1Client
	gwSvc   *sWorkloadV1GwService
	gw      apigw.APIGateway
}

func (a adapterWorkloadV1) AutoAddEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoAddEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoDeleteEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoDeleteEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoGetEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoGetEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoListEndpoint(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*workload.EndpointList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.EndpointList), err
}

func (a adapterWorkloadV1) AutoUpdateEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoUpdateEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoWatchEndpoint(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (workload.WorkloadV1_AutoWatchEndpointClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchEndpoint(ctx, in)
}

func (e *sWorkloadV1GwService) setupSvcProfile() {
	e.defSvcProf = apigwpkg.NewServiceProfile(nil)
	e.defSvcProf.SetDefaults()
	e.svcProf = make(map[string]apigw.ServiceProfile)

	e.svcProf["AutoAddEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf)
	e.svcProf["AutoDeleteEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf)
	e.svcProf["AutoGetEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf)
	e.svcProf["AutoListEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf)
	e.svcProf["AutoUpdateEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf)
}

// GetDefaultServiceProfile returns the default fallback service profile for this service
func (e *sWorkloadV1GwService) GetDefaultServiceProfile() (apigw.ServiceProfile, error) {
	if e.defSvcProf == nil {
		return nil, errors.New("not found")
	}
	return e.defSvcProf, nil
}

// GetServiceProfile returns the service profile for a given method in this service
func (e *sWorkloadV1GwService) GetServiceProfile(method string) (apigw.ServiceProfile, error) {
	if ret, ok := e.svcProf[method]; ok {
		return ret, nil
	}
	return nil, errors.New("not found")
}

// GetCrudServiceProfile returns the service profile for a auto generated crud operation
func (e *sWorkloadV1GwService) GetCrudServiceProfile(obj string, oper apiserver.APIOperType) (apigw.ServiceProfile, error) {
	name := apiserver.GetCrudServiceName(obj, oper)
	if name != "" {
		return e.GetServiceProfile(name)
	}
	return nil, errors.New("not found")
}

func (e *sWorkloadV1GwService) CompleteRegistration(ctx context.Context,
	logger log.Logger,
	grpcserver *grpc.Server,
	m *http.ServeMux,
	rslvr resolver.Interface,
	wg *sync.WaitGroup) error {
	apigw := apigwpkg.MustGetAPIGateway()
	// IP:port destination or service discovery key.
	grpcaddr := "pen-apiserver"
	grpcaddr = apigw.GetAPIServerAddr(grpcaddr)
	e.logger = logger

	marshaller := runtime.JSONBuiltin{}
	opts := runtime.WithMarshalerOption("*", &marshaller)
	muxMutex.Lock()
	if mux == nil {
		mux = runtime.NewServeMux(opts)
	}
	muxMutex.Unlock()
	e.setupSvcProfile()

	err := registerSwaggerDef(m, logger)
	if err != nil {
		logger.ErrorLog("msg", "failed to register swagger spec", "service", "workload.WorkloadV1", "error", err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			nctx, cancel := context.WithCancel(ctx)
			cl, err := e.newClient(nctx, grpcaddr, rslvr, apigw.GetDevMode())
			if err == nil {
				muxMutex.Lock()
				err = workload.RegisterWorkloadV1HandlerWithClient(ctx, mux, cl)
				muxMutex.Unlock()
				if err == nil {
					logger.InfoLog("msg", "registered service workload.WorkloadV1")
					m.Handle("/v1/workload/", http.StripPrefix("/v1/workload", mux))
					return
				} else {
					err = errors.Wrap(err, "failed to register")
				}
			} else {
				err = errors.Wrap(err, "failed to create client")
			}
			cancel()
			logger.ErrorLog("msg", "failed to register", "service", "workload.WorkloadV1", "error", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return nil
}

func (e *sWorkloadV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface, devmode bool) (*adapterWorkloadV1, error) {
	var opts []rpckit.Option
	if rslvr != nil {
		opts = append(opts, rpckit.WithBalancer(balancer.New(rslvr)))
	} else {
		opts = append(opts, rpckit.WithRemoteServerName("pen-apiserver"))
	}

	if !devmode {
		opts = append(opts, rpckit.WithTracerEnabled(false))
		opts = append(opts, rpckit.WithLoggerEnabled(false))
		opts = append(opts, rpckit.WithStatsEnabled(false))
	}

	client, err := rpckit.NewRPCClient(globals.APIGw, grpcAddr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "create rpc client")
	}

	e.logger.Infof("Connected to GRPC Server %s", grpcAddr)
	defer func() {
		go func() {
			<-ctx.Done()
			if cerr := client.Close(); cerr != nil {
				e.logger.ErrorLog("msg", "Failed to close conn on Done()", "addr", grpcAddr, "error", cerr)
			}
		}()
	}()

	cl := &adapterWorkloadV1{conn: client, gw: apigwpkg.MustGetAPIGateway(), gwSvc: e, service: grpcclient.NewWorkloadV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {

	apigw := apigwpkg.MustGetAPIGateway()

	svcWorkloadV1 := sWorkloadV1GwService{}
	apigw.Register("workload.WorkloadV1", "workload/", &svcWorkloadV1)
}
