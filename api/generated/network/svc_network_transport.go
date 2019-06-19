// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package network is a auto generated package.
Input file: svc_network.proto
*/
package network

import (
	"context"
	"encoding/json"
	"net/http"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/trace"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

type grpcServerNetworkV1 struct {
	Endpoints EndpointsNetworkV1Server

	AutoAddLbPolicyHdlr            grpctransport.Handler
	AutoAddNetworkHdlr             grpctransport.Handler
	AutoAddNetworkInterfaceHdlr    grpctransport.Handler
	AutoAddServiceHdlr             grpctransport.Handler
	AutoAddVirtualRouterHdlr       grpctransport.Handler
	AutoDeleteLbPolicyHdlr         grpctransport.Handler
	AutoDeleteNetworkHdlr          grpctransport.Handler
	AutoDeleteNetworkInterfaceHdlr grpctransport.Handler
	AutoDeleteServiceHdlr          grpctransport.Handler
	AutoDeleteVirtualRouterHdlr    grpctransport.Handler
	AutoGetLbPolicyHdlr            grpctransport.Handler
	AutoGetNetworkHdlr             grpctransport.Handler
	AutoGetNetworkInterfaceHdlr    grpctransport.Handler
	AutoGetServiceHdlr             grpctransport.Handler
	AutoGetVirtualRouterHdlr       grpctransport.Handler
	AutoListLbPolicyHdlr           grpctransport.Handler
	AutoListNetworkHdlr            grpctransport.Handler
	AutoListNetworkInterfaceHdlr   grpctransport.Handler
	AutoListServiceHdlr            grpctransport.Handler
	AutoListVirtualRouterHdlr      grpctransport.Handler
	AutoUpdateLbPolicyHdlr         grpctransport.Handler
	AutoUpdateNetworkHdlr          grpctransport.Handler
	AutoUpdateNetworkInterfaceHdlr grpctransport.Handler
	AutoUpdateServiceHdlr          grpctransport.Handler
	AutoUpdateVirtualRouterHdlr    grpctransport.Handler
}

// MakeGRPCServerNetworkV1 creates a GRPC server for NetworkV1 service
func MakeGRPCServerNetworkV1(ctx context.Context, endpoints EndpointsNetworkV1Server, logger log.Logger) NetworkV1Server {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(recoverVersion),
	}
	return &grpcServerNetworkV1{
		Endpoints: endpoints,
		AutoAddLbPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoAddLbPolicyEndpoint,
			DecodeGrpcReqLbPolicy,
			EncodeGrpcRespLbPolicy,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddLbPolicy", logger)))...,
		),

		AutoAddNetworkHdlr: grpctransport.NewServer(
			endpoints.AutoAddNetworkEndpoint,
			DecodeGrpcReqNetwork,
			EncodeGrpcRespNetwork,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddNetwork", logger)))...,
		),

		AutoAddNetworkInterfaceHdlr: grpctransport.NewServer(
			endpoints.AutoAddNetworkInterfaceEndpoint,
			DecodeGrpcReqNetworkInterface,
			EncodeGrpcRespNetworkInterface,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddNetworkInterface", logger)))...,
		),

		AutoAddServiceHdlr: grpctransport.NewServer(
			endpoints.AutoAddServiceEndpoint,
			DecodeGrpcReqService,
			EncodeGrpcRespService,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddService", logger)))...,
		),

		AutoAddVirtualRouterHdlr: grpctransport.NewServer(
			endpoints.AutoAddVirtualRouterEndpoint,
			DecodeGrpcReqVirtualRouter,
			EncodeGrpcRespVirtualRouter,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddVirtualRouter", logger)))...,
		),

		AutoDeleteLbPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteLbPolicyEndpoint,
			DecodeGrpcReqLbPolicy,
			EncodeGrpcRespLbPolicy,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteLbPolicy", logger)))...,
		),

		AutoDeleteNetworkHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteNetworkEndpoint,
			DecodeGrpcReqNetwork,
			EncodeGrpcRespNetwork,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteNetwork", logger)))...,
		),

		AutoDeleteNetworkInterfaceHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteNetworkInterfaceEndpoint,
			DecodeGrpcReqNetworkInterface,
			EncodeGrpcRespNetworkInterface,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteNetworkInterface", logger)))...,
		),

		AutoDeleteServiceHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteServiceEndpoint,
			DecodeGrpcReqService,
			EncodeGrpcRespService,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteService", logger)))...,
		),

		AutoDeleteVirtualRouterHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteVirtualRouterEndpoint,
			DecodeGrpcReqVirtualRouter,
			EncodeGrpcRespVirtualRouter,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteVirtualRouter", logger)))...,
		),

		AutoGetLbPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoGetLbPolicyEndpoint,
			DecodeGrpcReqLbPolicy,
			EncodeGrpcRespLbPolicy,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetLbPolicy", logger)))...,
		),

		AutoGetNetworkHdlr: grpctransport.NewServer(
			endpoints.AutoGetNetworkEndpoint,
			DecodeGrpcReqNetwork,
			EncodeGrpcRespNetwork,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetNetwork", logger)))...,
		),

		AutoGetNetworkInterfaceHdlr: grpctransport.NewServer(
			endpoints.AutoGetNetworkInterfaceEndpoint,
			DecodeGrpcReqNetworkInterface,
			EncodeGrpcRespNetworkInterface,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetNetworkInterface", logger)))...,
		),

		AutoGetServiceHdlr: grpctransport.NewServer(
			endpoints.AutoGetServiceEndpoint,
			DecodeGrpcReqService,
			EncodeGrpcRespService,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetService", logger)))...,
		),

		AutoGetVirtualRouterHdlr: grpctransport.NewServer(
			endpoints.AutoGetVirtualRouterEndpoint,
			DecodeGrpcReqVirtualRouter,
			EncodeGrpcRespVirtualRouter,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetVirtualRouter", logger)))...,
		),

		AutoListLbPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoListLbPolicyEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespLbPolicyList,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListLbPolicy", logger)))...,
		),

		AutoListNetworkHdlr: grpctransport.NewServer(
			endpoints.AutoListNetworkEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespNetworkList,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListNetwork", logger)))...,
		),

		AutoListNetworkInterfaceHdlr: grpctransport.NewServer(
			endpoints.AutoListNetworkInterfaceEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespNetworkInterfaceList,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListNetworkInterface", logger)))...,
		),

		AutoListServiceHdlr: grpctransport.NewServer(
			endpoints.AutoListServiceEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespServiceList,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListService", logger)))...,
		),

		AutoListVirtualRouterHdlr: grpctransport.NewServer(
			endpoints.AutoListVirtualRouterEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespVirtualRouterList,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListVirtualRouter", logger)))...,
		),

		AutoUpdateLbPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateLbPolicyEndpoint,
			DecodeGrpcReqLbPolicy,
			EncodeGrpcRespLbPolicy,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateLbPolicy", logger)))...,
		),

		AutoUpdateNetworkHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateNetworkEndpoint,
			DecodeGrpcReqNetwork,
			EncodeGrpcRespNetwork,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateNetwork", logger)))...,
		),

		AutoUpdateNetworkInterfaceHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateNetworkInterfaceEndpoint,
			DecodeGrpcReqNetworkInterface,
			EncodeGrpcRespNetworkInterface,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateNetworkInterface", logger)))...,
		),

		AutoUpdateServiceHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateServiceEndpoint,
			DecodeGrpcReqService,
			EncodeGrpcRespService,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateService", logger)))...,
		),

		AutoUpdateVirtualRouterHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateVirtualRouterEndpoint,
			DecodeGrpcReqVirtualRouter,
			EncodeGrpcRespVirtualRouter,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateVirtualRouter", logger)))...,
		),
	}
}

func (s *grpcServerNetworkV1) AutoAddLbPolicy(ctx oldcontext.Context, req *LbPolicy) (*LbPolicy, error) {
	_, resp, err := s.AutoAddLbPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoAddLbPolicy).V
	return &r, resp.(respNetworkV1AutoAddLbPolicy).Err
}

func decodeHTTPrespNetworkV1AutoAddLbPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp LbPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoAddNetwork(ctx oldcontext.Context, req *Network) (*Network, error) {
	_, resp, err := s.AutoAddNetworkHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoAddNetwork).V
	return &r, resp.(respNetworkV1AutoAddNetwork).Err
}

func decodeHTTPrespNetworkV1AutoAddNetwork(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Network
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoAddNetworkInterface(ctx oldcontext.Context, req *NetworkInterface) (*NetworkInterface, error) {
	_, resp, err := s.AutoAddNetworkInterfaceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoAddNetworkInterface).V
	return &r, resp.(respNetworkV1AutoAddNetworkInterface).Err
}

func decodeHTTPrespNetworkV1AutoAddNetworkInterface(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp NetworkInterface
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoAddService(ctx oldcontext.Context, req *Service) (*Service, error) {
	_, resp, err := s.AutoAddServiceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoAddService).V
	return &r, resp.(respNetworkV1AutoAddService).Err
}

func decodeHTTPrespNetworkV1AutoAddService(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Service
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoAddVirtualRouter(ctx oldcontext.Context, req *VirtualRouter) (*VirtualRouter, error) {
	_, resp, err := s.AutoAddVirtualRouterHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoAddVirtualRouter).V
	return &r, resp.(respNetworkV1AutoAddVirtualRouter).Err
}

func decodeHTTPrespNetworkV1AutoAddVirtualRouter(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp VirtualRouter
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoDeleteLbPolicy(ctx oldcontext.Context, req *LbPolicy) (*LbPolicy, error) {
	_, resp, err := s.AutoDeleteLbPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoDeleteLbPolicy).V
	return &r, resp.(respNetworkV1AutoDeleteLbPolicy).Err
}

func decodeHTTPrespNetworkV1AutoDeleteLbPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp LbPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoDeleteNetwork(ctx oldcontext.Context, req *Network) (*Network, error) {
	_, resp, err := s.AutoDeleteNetworkHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoDeleteNetwork).V
	return &r, resp.(respNetworkV1AutoDeleteNetwork).Err
}

func decodeHTTPrespNetworkV1AutoDeleteNetwork(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Network
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoDeleteNetworkInterface(ctx oldcontext.Context, req *NetworkInterface) (*NetworkInterface, error) {
	_, resp, err := s.AutoDeleteNetworkInterfaceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoDeleteNetworkInterface).V
	return &r, resp.(respNetworkV1AutoDeleteNetworkInterface).Err
}

func decodeHTTPrespNetworkV1AutoDeleteNetworkInterface(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp NetworkInterface
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoDeleteService(ctx oldcontext.Context, req *Service) (*Service, error) {
	_, resp, err := s.AutoDeleteServiceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoDeleteService).V
	return &r, resp.(respNetworkV1AutoDeleteService).Err
}

func decodeHTTPrespNetworkV1AutoDeleteService(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Service
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoDeleteVirtualRouter(ctx oldcontext.Context, req *VirtualRouter) (*VirtualRouter, error) {
	_, resp, err := s.AutoDeleteVirtualRouterHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoDeleteVirtualRouter).V
	return &r, resp.(respNetworkV1AutoDeleteVirtualRouter).Err
}

func decodeHTTPrespNetworkV1AutoDeleteVirtualRouter(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp VirtualRouter
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoGetLbPolicy(ctx oldcontext.Context, req *LbPolicy) (*LbPolicy, error) {
	_, resp, err := s.AutoGetLbPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoGetLbPolicy).V
	return &r, resp.(respNetworkV1AutoGetLbPolicy).Err
}

func decodeHTTPrespNetworkV1AutoGetLbPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp LbPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoGetNetwork(ctx oldcontext.Context, req *Network) (*Network, error) {
	_, resp, err := s.AutoGetNetworkHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoGetNetwork).V
	return &r, resp.(respNetworkV1AutoGetNetwork).Err
}

func decodeHTTPrespNetworkV1AutoGetNetwork(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Network
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoGetNetworkInterface(ctx oldcontext.Context, req *NetworkInterface) (*NetworkInterface, error) {
	_, resp, err := s.AutoGetNetworkInterfaceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoGetNetworkInterface).V
	return &r, resp.(respNetworkV1AutoGetNetworkInterface).Err
}

func decodeHTTPrespNetworkV1AutoGetNetworkInterface(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp NetworkInterface
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoGetService(ctx oldcontext.Context, req *Service) (*Service, error) {
	_, resp, err := s.AutoGetServiceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoGetService).V
	return &r, resp.(respNetworkV1AutoGetService).Err
}

func decodeHTTPrespNetworkV1AutoGetService(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Service
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoGetVirtualRouter(ctx oldcontext.Context, req *VirtualRouter) (*VirtualRouter, error) {
	_, resp, err := s.AutoGetVirtualRouterHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoGetVirtualRouter).V
	return &r, resp.(respNetworkV1AutoGetVirtualRouter).Err
}

func decodeHTTPrespNetworkV1AutoGetVirtualRouter(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp VirtualRouter
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoListLbPolicy(ctx oldcontext.Context, req *api.ListWatchOptions) (*LbPolicyList, error) {
	_, resp, err := s.AutoListLbPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoListLbPolicy).V
	return &r, resp.(respNetworkV1AutoListLbPolicy).Err
}

func decodeHTTPrespNetworkV1AutoListLbPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp LbPolicyList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoListNetwork(ctx oldcontext.Context, req *api.ListWatchOptions) (*NetworkList, error) {
	_, resp, err := s.AutoListNetworkHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoListNetwork).V
	return &r, resp.(respNetworkV1AutoListNetwork).Err
}

func decodeHTTPrespNetworkV1AutoListNetwork(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp NetworkList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoListNetworkInterface(ctx oldcontext.Context, req *api.ListWatchOptions) (*NetworkInterfaceList, error) {
	_, resp, err := s.AutoListNetworkInterfaceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoListNetworkInterface).V
	return &r, resp.(respNetworkV1AutoListNetworkInterface).Err
}

func decodeHTTPrespNetworkV1AutoListNetworkInterface(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp NetworkInterfaceList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoListService(ctx oldcontext.Context, req *api.ListWatchOptions) (*ServiceList, error) {
	_, resp, err := s.AutoListServiceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoListService).V
	return &r, resp.(respNetworkV1AutoListService).Err
}

func decodeHTTPrespNetworkV1AutoListService(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp ServiceList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoListVirtualRouter(ctx oldcontext.Context, req *api.ListWatchOptions) (*VirtualRouterList, error) {
	_, resp, err := s.AutoListVirtualRouterHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoListVirtualRouter).V
	return &r, resp.(respNetworkV1AutoListVirtualRouter).Err
}

func decodeHTTPrespNetworkV1AutoListVirtualRouter(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp VirtualRouterList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoUpdateLbPolicy(ctx oldcontext.Context, req *LbPolicy) (*LbPolicy, error) {
	_, resp, err := s.AutoUpdateLbPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoUpdateLbPolicy).V
	return &r, resp.(respNetworkV1AutoUpdateLbPolicy).Err
}

func decodeHTTPrespNetworkV1AutoUpdateLbPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp LbPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoUpdateNetwork(ctx oldcontext.Context, req *Network) (*Network, error) {
	_, resp, err := s.AutoUpdateNetworkHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoUpdateNetwork).V
	return &r, resp.(respNetworkV1AutoUpdateNetwork).Err
}

func decodeHTTPrespNetworkV1AutoUpdateNetwork(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Network
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoUpdateNetworkInterface(ctx oldcontext.Context, req *NetworkInterface) (*NetworkInterface, error) {
	_, resp, err := s.AutoUpdateNetworkInterfaceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoUpdateNetworkInterface).V
	return &r, resp.(respNetworkV1AutoUpdateNetworkInterface).Err
}

func decodeHTTPrespNetworkV1AutoUpdateNetworkInterface(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp NetworkInterface
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoUpdateService(ctx oldcontext.Context, req *Service) (*Service, error) {
	_, resp, err := s.AutoUpdateServiceHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoUpdateService).V
	return &r, resp.(respNetworkV1AutoUpdateService).Err
}

func decodeHTTPrespNetworkV1AutoUpdateService(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Service
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoUpdateVirtualRouter(ctx oldcontext.Context, req *VirtualRouter) (*VirtualRouter, error) {
	_, resp, err := s.AutoUpdateVirtualRouterHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respNetworkV1AutoUpdateVirtualRouter).V
	return &r, resp.(respNetworkV1AutoUpdateVirtualRouter).Err
}

func decodeHTTPrespNetworkV1AutoUpdateVirtualRouter(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp VirtualRouter
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerNetworkV1) AutoWatchSvcNetworkV1(in *api.ListWatchOptions, stream NetworkV1_AutoWatchSvcNetworkV1Server) error {
	return s.Endpoints.AutoWatchSvcNetworkV1(in, stream)
}

func (s *grpcServerNetworkV1) AutoWatchNetwork(in *api.ListWatchOptions, stream NetworkV1_AutoWatchNetworkServer) error {
	return s.Endpoints.AutoWatchNetwork(in, stream)
}

func (s *grpcServerNetworkV1) AutoWatchService(in *api.ListWatchOptions, stream NetworkV1_AutoWatchServiceServer) error {
	return s.Endpoints.AutoWatchService(in, stream)
}

func (s *grpcServerNetworkV1) AutoWatchLbPolicy(in *api.ListWatchOptions, stream NetworkV1_AutoWatchLbPolicyServer) error {
	return s.Endpoints.AutoWatchLbPolicy(in, stream)
}

func (s *grpcServerNetworkV1) AutoWatchVirtualRouter(in *api.ListWatchOptions, stream NetworkV1_AutoWatchVirtualRouterServer) error {
	return s.Endpoints.AutoWatchVirtualRouter(in, stream)
}

func (s *grpcServerNetworkV1) AutoWatchNetworkInterface(in *api.ListWatchOptions, stream NetworkV1_AutoWatchNetworkInterfaceServer) error {
	return s.Endpoints.AutoWatchNetworkInterface(in, stream)
}

func encodeHTTPLbPolicyList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPLbPolicyList(_ context.Context, r *http.Request) (interface{}, error) {
	var req LbPolicyList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqLbPolicyList encodes GRPC request
func EncodeGrpcReqLbPolicyList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*LbPolicyList)
	return req, nil
}

// DecodeGrpcReqLbPolicyList decodes GRPC request
func DecodeGrpcReqLbPolicyList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*LbPolicyList)
	return req, nil
}

// EncodeGrpcRespLbPolicyList endodes the GRPC response
func EncodeGrpcRespLbPolicyList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespLbPolicyList decodes the GRPC response
func DecodeGrpcRespLbPolicyList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPNetworkInterfaceList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPNetworkInterfaceList(_ context.Context, r *http.Request) (interface{}, error) {
	var req NetworkInterfaceList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqNetworkInterfaceList encodes GRPC request
func EncodeGrpcReqNetworkInterfaceList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*NetworkInterfaceList)
	return req, nil
}

// DecodeGrpcReqNetworkInterfaceList decodes GRPC request
func DecodeGrpcReqNetworkInterfaceList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*NetworkInterfaceList)
	return req, nil
}

// EncodeGrpcRespNetworkInterfaceList endodes the GRPC response
func EncodeGrpcRespNetworkInterfaceList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespNetworkInterfaceList decodes the GRPC response
func DecodeGrpcRespNetworkInterfaceList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPNetworkList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPNetworkList(_ context.Context, r *http.Request) (interface{}, error) {
	var req NetworkList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqNetworkList encodes GRPC request
func EncodeGrpcReqNetworkList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*NetworkList)
	return req, nil
}

// DecodeGrpcReqNetworkList decodes GRPC request
func DecodeGrpcReqNetworkList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*NetworkList)
	return req, nil
}

// EncodeGrpcRespNetworkList endodes the GRPC response
func EncodeGrpcRespNetworkList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespNetworkList decodes the GRPC response
func DecodeGrpcRespNetworkList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPServiceList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPServiceList(_ context.Context, r *http.Request) (interface{}, error) {
	var req ServiceList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqServiceList encodes GRPC request
func EncodeGrpcReqServiceList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ServiceList)
	return req, nil
}

// DecodeGrpcReqServiceList decodes GRPC request
func DecodeGrpcReqServiceList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ServiceList)
	return req, nil
}

// EncodeGrpcRespServiceList endodes the GRPC response
func EncodeGrpcRespServiceList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespServiceList decodes the GRPC response
func DecodeGrpcRespServiceList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPVirtualRouterList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPVirtualRouterList(_ context.Context, r *http.Request) (interface{}, error) {
	var req VirtualRouterList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqVirtualRouterList encodes GRPC request
func EncodeGrpcReqVirtualRouterList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*VirtualRouterList)
	return req, nil
}

// DecodeGrpcReqVirtualRouterList decodes GRPC request
func DecodeGrpcReqVirtualRouterList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*VirtualRouterList)
	return req, nil
}

// EncodeGrpcRespVirtualRouterList endodes the GRPC response
func EncodeGrpcRespVirtualRouterList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespVirtualRouterList decodes the GRPC response
func DecodeGrpcRespVirtualRouterList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
