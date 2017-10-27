// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package grpcclient

import (
	"context"
	"errors"
	oldlog "log"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	api "github.com/pensando/sw/api"
	export "github.com/pensando/sw/api/generated/export"
	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	apiserver "github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta
var _ listerwatcher.WatcherClient
var _ kvstore.Interface

// NewExportPolicyV1 sets up a new client for ExportPolicyV1
func NewExportPolicyV1(conn *grpc.ClientConn, logger log.Logger) export.ServiceExportPolicyV1Client {

	var lAutoAddExportPolicyEndpoint endpoint.Endpoint
	{
		lAutoAddExportPolicyEndpoint = grpctransport.NewClient(
			conn,
			"export.ExportPolicyV1",
			"AutoAddExportPolicy",
			export.EncodeGrpcReqExportPolicy,
			export.DecodeGrpcRespExportPolicy,
			&export.ExportPolicy{},
			grpctransport.ClientBefore(opentracing.ToGRPCRequest(stdopentracing.GlobalTracer(), logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoAddExportPolicyEndpoint = opentracing.TraceClient(stdopentracing.GlobalTracer(), "ExportPolicyV1:AutoAddExportPolicy")(lAutoAddExportPolicyEndpoint)
	}
	var lAutoDeleteExportPolicyEndpoint endpoint.Endpoint
	{
		lAutoDeleteExportPolicyEndpoint = grpctransport.NewClient(
			conn,
			"export.ExportPolicyV1",
			"AutoDeleteExportPolicy",
			export.EncodeGrpcReqExportPolicy,
			export.DecodeGrpcRespExportPolicy,
			&export.ExportPolicy{},
			grpctransport.ClientBefore(opentracing.ToGRPCRequest(stdopentracing.GlobalTracer(), logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoDeleteExportPolicyEndpoint = opentracing.TraceClient(stdopentracing.GlobalTracer(), "ExportPolicyV1:AutoDeleteExportPolicy")(lAutoDeleteExportPolicyEndpoint)
	}
	var lAutoGetExportPolicyEndpoint endpoint.Endpoint
	{
		lAutoGetExportPolicyEndpoint = grpctransport.NewClient(
			conn,
			"export.ExportPolicyV1",
			"AutoGetExportPolicy",
			export.EncodeGrpcReqExportPolicy,
			export.DecodeGrpcRespExportPolicy,
			&export.ExportPolicy{},
			grpctransport.ClientBefore(opentracing.ToGRPCRequest(stdopentracing.GlobalTracer(), logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoGetExportPolicyEndpoint = opentracing.TraceClient(stdopentracing.GlobalTracer(), "ExportPolicyV1:AutoGetExportPolicy")(lAutoGetExportPolicyEndpoint)
	}
	var lAutoListExportPolicyEndpoint endpoint.Endpoint
	{
		lAutoListExportPolicyEndpoint = grpctransport.NewClient(
			conn,
			"export.ExportPolicyV1",
			"AutoListExportPolicy",
			export.EncodeGrpcReqListWatchOptions,
			export.DecodeGrpcRespExportPolicyList,
			&export.ExportPolicyList{},
			grpctransport.ClientBefore(opentracing.ToGRPCRequest(stdopentracing.GlobalTracer(), logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoListExportPolicyEndpoint = opentracing.TraceClient(stdopentracing.GlobalTracer(), "ExportPolicyV1:AutoListExportPolicy")(lAutoListExportPolicyEndpoint)
	}
	var lAutoUpdateExportPolicyEndpoint endpoint.Endpoint
	{
		lAutoUpdateExportPolicyEndpoint = grpctransport.NewClient(
			conn,
			"export.ExportPolicyV1",
			"AutoUpdateExportPolicy",
			export.EncodeGrpcReqExportPolicy,
			export.DecodeGrpcRespExportPolicy,
			&export.ExportPolicy{},
			grpctransport.ClientBefore(opentracing.ToGRPCRequest(stdopentracing.GlobalTracer(), logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoUpdateExportPolicyEndpoint = opentracing.TraceClient(stdopentracing.GlobalTracer(), "ExportPolicyV1:AutoUpdateExportPolicy")(lAutoUpdateExportPolicyEndpoint)
	}
	return export.EndpointsExportPolicyV1Client{
		Client: export.NewExportPolicyV1Client(conn),

		AutoAddExportPolicyEndpoint:    lAutoAddExportPolicyEndpoint,
		AutoDeleteExportPolicyEndpoint: lAutoDeleteExportPolicyEndpoint,
		AutoGetExportPolicyEndpoint:    lAutoGetExportPolicyEndpoint,
		AutoListExportPolicyEndpoint:   lAutoListExportPolicyEndpoint,
		AutoUpdateExportPolicyEndpoint: lAutoUpdateExportPolicyEndpoint,
	}
}

// NewExportPolicyV1Backend creates an instrumented client with middleware
func NewExportPolicyV1Backend(conn *grpc.ClientConn, logger log.Logger) export.ServiceExportPolicyV1Client {
	cl := NewExportPolicyV1(conn, logger)
	cl = export.LoggingExportPolicyV1MiddlewareClient(logger)(cl)
	return cl
}

type grpcObjExportPolicyV1ExportPolicy struct {
	logger log.Logger
	client export.ServiceExportPolicyV1Client
}

func (a *grpcObjExportPolicyV1ExportPolicy) Create(ctx context.Context, in *export.ExportPolicy) (*export.ExportPolicy, error) {
	a.logger.DebugLog("msg", "recieved call", "object", "ExportPolicy", "oper", "create")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoAddExportPolicy(nctx, in)
}

func (a *grpcObjExportPolicyV1ExportPolicy) Update(ctx context.Context, in *export.ExportPolicy) (*export.ExportPolicy, error) {
	a.logger.DebugLog("msg", "received call", "object", "ExportPolicy", "oper", "update")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoUpdateExportPolicy(nctx, in)
}

func (a *grpcObjExportPolicyV1ExportPolicy) Get(ctx context.Context, objMeta *api.ObjectMeta) (*export.ExportPolicy, error) {
	a.logger.DebugLog("msg", "received call", "object", "ExportPolicy", "oper", "get")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := export.ExportPolicy{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoGetExportPolicy(nctx, &in)
}

func (a *grpcObjExportPolicyV1ExportPolicy) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*export.ExportPolicy, error) {
	a.logger.DebugLog("msg", "received call", "object", "ExportPolicy", "oper", "delete")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := export.ExportPolicy{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoDeleteExportPolicy(nctx, &in)
}

func (a *grpcObjExportPolicyV1ExportPolicy) List(ctx context.Context, options *api.ListWatchOptions) ([]*export.ExportPolicy, error) {
	a.logger.DebugLog("msg", "received call", "object", "ExportPolicy", "oper", "list")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	r, err := a.client.AutoListExportPolicy(nctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *grpcObjExportPolicyV1ExportPolicy) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	a.logger.DebugLog("msg", "received call", "object", "ExportPolicy", "oper", "WatchOper")
	nctx := addVersion(ctx, "v1")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	stream, err := a.client.AutoWatchExportPolicy(nctx, options)
	if err != nil {
		return nil, err
	}
	wstream := stream.(export.ExportPolicyV1_AutoWatchExportPolicyClient)
	bridgefn := func(lw *listerwatcher.WatcherClient) {
		for {
			r, err := wstream.Recv()
			if err != nil {
				a.logger.ErrorLog("msg", "error on recieve", "error", err)
				close(lw.OutCh)
				return
			}
			ev := kvstore.WatchEvent{
				Type:   kvstore.WatchEventType(r.Type),
				Object: r.Object,
			}
			lw.OutCh <- &ev
		}
	}
	lw := listerwatcher.NewWatcherClient(wstream, bridgefn)
	lw.Run()
	return lw, nil
}

func (a *grpcObjExportPolicyV1ExportPolicy) Allowed(oper apiserver.APIOperType) bool {
	return true
}

type restObjExportPolicyV1ExportPolicy struct {
	endpoints export.EndpointsExportPolicyV1RestClient
	instance  string
}

func (a *restObjExportPolicyV1ExportPolicy) Create(ctx context.Context, in *export.ExportPolicy) (*export.ExportPolicy, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoAddExportPolicy(ctx, in)
}

func (a *restObjExportPolicyV1ExportPolicy) Update(ctx context.Context, in *export.ExportPolicy) (*export.ExportPolicy, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoUpdateExportPolicy(ctx, in)
}

func (a *restObjExportPolicyV1ExportPolicy) Get(ctx context.Context, objMeta *api.ObjectMeta) (*export.ExportPolicy, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := export.ExportPolicy{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoGetExportPolicy(ctx, &in)
}

func (a *restObjExportPolicyV1ExportPolicy) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*export.ExportPolicy, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := export.ExportPolicy{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoDeleteExportPolicy(ctx, &in)
}

func (a *restObjExportPolicyV1ExportPolicy) List(ctx context.Context, options *api.ListWatchOptions) ([]*export.ExportPolicy, error) {
	if options == nil {
		return nil, errors.New("invalid input")
	}
	r, err := a.endpoints.AutoListExportPolicy(ctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *restObjExportPolicyV1ExportPolicy) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	return nil, errors.New("not allowed")
}

func (a *restObjExportPolicyV1ExportPolicy) Allowed(oper apiserver.APIOperType) bool {
	switch oper {
	case apiserver.CreateOper:
		return true
	case apiserver.UpdateOper:
		return true
	case apiserver.GetOper:
		return true
	case apiserver.DeleteOper:
		return true
	case apiserver.ListOper:
		return false
	case apiserver.WatchOper:
		return false
	default:
		return false
	}
}

type crudClientExportPolicyV1 struct {
	grpcExportPolicy export.ExportPolicyInterface
}

// NewGrpcCrudClientExportPolicyV1 creates a GRPC client for the service
func NewGrpcCrudClientExportPolicyV1(conn *grpc.ClientConn, logger log.Logger) export.ExportPolicyV1Interface {
	client := NewExportPolicyV1Backend(conn, logger)
	return &crudClientExportPolicyV1{

		grpcExportPolicy: &grpcObjExportPolicyV1ExportPolicy{client: client, logger: logger},
	}
}

func (a *crudClientExportPolicyV1) ExportPolicy() export.ExportPolicyInterface {
	return a.grpcExportPolicy
}

type crudRestClientExportPolicyV1 struct {
	restExportPolicy export.ExportPolicyInterface
}

// NewRestCrudClientExportPolicyV1 creates a REST client for the service.
func NewRestCrudClientExportPolicyV1(url string) export.ExportPolicyV1Interface {
	endpoints, err := export.MakeExportPolicyV1RestClientEndpoints(url)
	if err != nil {
		oldlog.Fatal("failed to create client")
	}
	return &crudRestClientExportPolicyV1{

		restExportPolicy: &restObjExportPolicyV1ExportPolicy{endpoints: endpoints, instance: url},
	}
}

func (a *crudRestClientExportPolicyV1) ExportPolicy() export.ExportPolicyInterface {
	return a.restExportPolicy
}
