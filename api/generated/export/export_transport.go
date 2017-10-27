// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package export is a auto generated package.
Input file: protos/export.proto
*/
package export

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/log"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

type grpcServerExportPolicyV1 struct {
	Endpoints EndpointsExportPolicyV1Server

	AutoAddExportPolicyHdlr    grpctransport.Handler
	AutoDeleteExportPolicyHdlr grpctransport.Handler
	AutoGetExportPolicyHdlr    grpctransport.Handler
	AutoListExportPolicyHdlr   grpctransport.Handler
	AutoUpdateExportPolicyHdlr grpctransport.Handler
}

// MakeGRPCServerExportPolicyV1 creates a GRPC server for ExportPolicyV1 service
func MakeGRPCServerExportPolicyV1(ctx context.Context, endpoints EndpointsExportPolicyV1Server, logger log.Logger) ExportPolicyV1Server {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(recoverVersion),
	}
	return &grpcServerExportPolicyV1{
		Endpoints: endpoints,
		AutoAddExportPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoAddExportPolicyEndpoint,
			DecodeGrpcReqExportPolicy,
			EncodeGrpcRespExportPolicy,
			append(options, grpctransport.ServerBefore(opentracing.FromGRPCRequest(stdopentracing.GlobalTracer(), "AutoAddExportPolicy", logger)))...,
		),

		AutoDeleteExportPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteExportPolicyEndpoint,
			DecodeGrpcReqExportPolicy,
			EncodeGrpcRespExportPolicy,
			append(options, grpctransport.ServerBefore(opentracing.FromGRPCRequest(stdopentracing.GlobalTracer(), "AutoDeleteExportPolicy", logger)))...,
		),

		AutoGetExportPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoGetExportPolicyEndpoint,
			DecodeGrpcReqExportPolicy,
			EncodeGrpcRespExportPolicy,
			append(options, grpctransport.ServerBefore(opentracing.FromGRPCRequest(stdopentracing.GlobalTracer(), "AutoGetExportPolicy", logger)))...,
		),

		AutoListExportPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoListExportPolicyEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespExportPolicyList,
			append(options, grpctransport.ServerBefore(opentracing.FromGRPCRequest(stdopentracing.GlobalTracer(), "AutoListExportPolicy", logger)))...,
		),

		AutoUpdateExportPolicyHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateExportPolicyEndpoint,
			DecodeGrpcReqExportPolicy,
			EncodeGrpcRespExportPolicy,
			append(options, grpctransport.ServerBefore(opentracing.FromGRPCRequest(stdopentracing.GlobalTracer(), "AutoUpdateExportPolicy", logger)))...,
		),
	}
}

func (s *grpcServerExportPolicyV1) AutoAddExportPolicy(ctx oldcontext.Context, req *ExportPolicy) (*ExportPolicy, error) {
	_, resp, err := s.AutoAddExportPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respExportPolicyV1AutoAddExportPolicy).V
	return &r, resp.(respExportPolicyV1AutoAddExportPolicy).Err
}

func decodeHTTPrespExportPolicyV1AutoAddExportPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp ExportPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerExportPolicyV1) AutoDeleteExportPolicy(ctx oldcontext.Context, req *ExportPolicy) (*ExportPolicy, error) {
	_, resp, err := s.AutoDeleteExportPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respExportPolicyV1AutoDeleteExportPolicy).V
	return &r, resp.(respExportPolicyV1AutoDeleteExportPolicy).Err
}

func decodeHTTPrespExportPolicyV1AutoDeleteExportPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp ExportPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerExportPolicyV1) AutoGetExportPolicy(ctx oldcontext.Context, req *ExportPolicy) (*ExportPolicy, error) {
	_, resp, err := s.AutoGetExportPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respExportPolicyV1AutoGetExportPolicy).V
	return &r, resp.(respExportPolicyV1AutoGetExportPolicy).Err
}

func decodeHTTPrespExportPolicyV1AutoGetExportPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp ExportPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerExportPolicyV1) AutoListExportPolicy(ctx oldcontext.Context, req *api.ListWatchOptions) (*ExportPolicyList, error) {
	_, resp, err := s.AutoListExportPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respExportPolicyV1AutoListExportPolicy).V
	return &r, resp.(respExportPolicyV1AutoListExportPolicy).Err
}

func decodeHTTPrespExportPolicyV1AutoListExportPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp ExportPolicyList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerExportPolicyV1) AutoUpdateExportPolicy(ctx oldcontext.Context, req *ExportPolicy) (*ExportPolicy, error) {
	_, resp, err := s.AutoUpdateExportPolicyHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respExportPolicyV1AutoUpdateExportPolicy).V
	return &r, resp.(respExportPolicyV1AutoUpdateExportPolicy).Err
}

func decodeHTTPrespExportPolicyV1AutoUpdateExportPolicy(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp ExportPolicy
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerExportPolicyV1) AutoWatchExportPolicy(in *api.ListWatchOptions, stream ExportPolicyV1_AutoWatchExportPolicyServer) error {
	return s.Endpoints.AutoWatchExportPolicy(in, stream)
}

func encodeHTTPExportPolicy(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPExportPolicy(_ context.Context, r *http.Request) (interface{}, error) {
	var req ExportPolicy
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqExportPolicy encodes GRPC request
func EncodeGrpcReqExportPolicy(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicy)
	return req, nil
}

// DecodeGrpcReqExportPolicy decodes GRPC request
func DecodeGrpcReqExportPolicy(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicy)
	return req, nil
}

// EncodeGrpcRespExportPolicy encodes GRC response
func EncodeGrpcRespExportPolicy(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespExportPolicy decodes GRPC response
func DecodeGrpcRespExportPolicy(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPExportPolicyList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPExportPolicyList(_ context.Context, r *http.Request) (interface{}, error) {
	var req ExportPolicyList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqExportPolicyList encodes GRPC request
func EncodeGrpcReqExportPolicyList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicyList)
	return req, nil
}

// DecodeGrpcReqExportPolicyList decodes GRPC request
func DecodeGrpcReqExportPolicyList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicyList)
	return req, nil
}

// EncodeGrpcRespExportPolicyList endodes the GRPC response
func EncodeGrpcRespExportPolicyList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespExportPolicyList decodes the GRPC response
func DecodeGrpcRespExportPolicyList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPExportPolicySpec(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPExportPolicySpec(_ context.Context, r *http.Request) (interface{}, error) {
	var req ExportPolicySpec
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqExportPolicySpec encodes GRPC request
func EncodeGrpcReqExportPolicySpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicySpec)
	return req, nil
}

// DecodeGrpcReqExportPolicySpec decodes GRPC request
func DecodeGrpcReqExportPolicySpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicySpec)
	return req, nil
}

// EncodeGrpcRespExportPolicySpec encodes GRC response
func EncodeGrpcRespExportPolicySpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespExportPolicySpec decodes GRPC response
func DecodeGrpcRespExportPolicySpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPExportPolicyStatus(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPExportPolicyStatus(_ context.Context, r *http.Request) (interface{}, error) {
	var req ExportPolicyStatus
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqExportPolicyStatus encodes GRPC request
func EncodeGrpcReqExportPolicyStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicyStatus)
	return req, nil
}

// DecodeGrpcReqExportPolicyStatus decodes GRPC request
func DecodeGrpcReqExportPolicyStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExportPolicyStatus)
	return req, nil
}

// EncodeGrpcRespExportPolicyStatus encodes GRC response
func EncodeGrpcRespExportPolicyStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespExportPolicyStatus decodes GRPC response
func DecodeGrpcRespExportPolicyStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPExternalCred(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPExternalCred(_ context.Context, r *http.Request) (interface{}, error) {
	var req ExternalCred
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqExternalCred encodes GRPC request
func EncodeGrpcReqExternalCred(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExternalCred)
	return req, nil
}

// DecodeGrpcReqExternalCred decodes GRPC request
func DecodeGrpcReqExternalCred(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*ExternalCred)
	return req, nil
}

// EncodeGrpcRespExternalCred encodes GRC response
func EncodeGrpcRespExternalCred(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespExternalCred decodes GRPC response
func DecodeGrpcRespExternalCred(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
