// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package aggwatch is a auto generated package.
Input file: svc_aggwatch.proto
*/
package aggwatch

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/listerwatcher"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/venice/utils/log"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta
var _ grpc.ServerStream
var _ fmt.Formatter
var _ *listerwatcher.WatcherClient

// MiddlewareAggWatchV1Client add middleware to the client
type MiddlewareAggWatchV1Client func(ServiceAggWatchV1Client) ServiceAggWatchV1Client

// EndpointsAggWatchV1Client is the endpoints for the client
type EndpointsAggWatchV1Client struct {
	Client                         AggWatchV1Client
	AutoWatchSvcAggWatchV1Endpoint endpoint.Endpoint
}

// EndpointsAggWatchV1RestClient is the REST client
type EndpointsAggWatchV1RestClient struct {
	logger   log.Logger
	client   *http.Client
	instance string
	bufferId string

	AutoWatchSvcAggWatchV1Endpoint endpoint.Endpoint
}

// MiddlewareAggWatchV1Server adds middle ware to the server
type MiddlewareAggWatchV1Server func(ServiceAggWatchV1Server) ServiceAggWatchV1Server

// EndpointsAggWatchV1Server is the server endpoints
type EndpointsAggWatchV1Server struct {
	svcWatchHandlerAggWatchV1 func(options *api.AggWatchOptions, stream grpc.ServerStream) error
}

func (e EndpointsAggWatchV1Client) AutoWatchSvcAggWatchV1(ctx context.Context, in *api.AggWatchOptions) (AggWatchV1_AutoWatchSvcAggWatchV1Client, error) {
	return e.Client.AutoWatchSvcAggWatchV1(ctx, in)
}

func (e EndpointsAggWatchV1Server) AutoWatchSvcAggWatchV1(in *api.AggWatchOptions, stream AggWatchV1_AutoWatchSvcAggWatchV1Server) error {
	return e.svcWatchHandlerAggWatchV1(in, stream)
}

// MakeAutoWatchSvcAggWatchV1Endpoint creates the Watch endpoint for the service
func MakeAutoWatchSvcAggWatchV1Endpoint(s ServiceAggWatchV1Server, logger log.Logger) func(options *api.AggWatchOptions, stream grpc.ServerStream) error {
	return func(options *api.AggWatchOptions, stream grpc.ServerStream) error {
		wstream := stream.(AggWatchV1_AutoWatchSvcAggWatchV1Server)
		return s.AutoWatchSvcAggWatchV1(options, wstream)
	}
}

// MakeAggWatchV1ServerEndpoints creates server endpoints
func MakeAggWatchV1ServerEndpoints(s ServiceAggWatchV1Server, logger log.Logger) EndpointsAggWatchV1Server {
	return EndpointsAggWatchV1Server{
		svcWatchHandlerAggWatchV1: MakeAutoWatchSvcAggWatchV1Endpoint(s, logger),
	}
}

// LoggingAggWatchV1MiddlewareClient adds middleware for the client
func LoggingAggWatchV1MiddlewareClient(logger log.Logger) MiddlewareAggWatchV1Client {
	return func(next ServiceAggWatchV1Client) ServiceAggWatchV1Client {
		return loggingAggWatchV1MiddlewareClient{
			logger: logger,
			next:   next,
		}
	}
}

type loggingAggWatchV1MiddlewareClient struct {
	logger log.Logger
	next   ServiceAggWatchV1Client
}

// LoggingAggWatchV1MiddlewareServer adds middleware for the client
func LoggingAggWatchV1MiddlewareServer(logger log.Logger) MiddlewareAggWatchV1Server {
	return func(next ServiceAggWatchV1Server) ServiceAggWatchV1Server {
		return loggingAggWatchV1MiddlewareServer{
			logger: logger,
			next:   next,
		}
	}
}

type loggingAggWatchV1MiddlewareServer struct {
	logger log.Logger
	next   ServiceAggWatchV1Server
}

func (m loggingAggWatchV1MiddlewareClient) AutoWatchSvcAggWatchV1(ctx context.Context, in *api.AggWatchOptions) (resp AggWatchV1_AutoWatchSvcAggWatchV1Client, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "AggWatchV1", "method", "AutoWatchSvcAggWatchV1", "result", rslt, "duration", time.Since(begin), "error", err)
	}(time.Now())
	resp, err = m.next.AutoWatchSvcAggWatchV1(ctx, in)
	return
}

func (m loggingAggWatchV1MiddlewareServer) AutoWatchSvcAggWatchV1(in *api.AggWatchOptions, stream AggWatchV1_AutoWatchSvcAggWatchV1Server) (err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(stream.Context(), "service", "AggWatchV1", "method", "AutoWatchSvcAggWatchV1", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	err = m.next.AutoWatchSvcAggWatchV1(in, stream)
	return
}

func (r *EndpointsAggWatchV1RestClient) updateHTTPHeader(ctx context.Context, header *http.Header) {
	val, ok := loginctx.AuthzHeaderFromContext(ctx)
	if ok {
		header.Add("Authorization", val)
	}
	val, ok = loginctx.ExtRequestIDHeaderFromContext(ctx)
	if ok {
		header.Add("Pensando-Psm-External-Request-Id", val)
	}
}
func (r *EndpointsAggWatchV1RestClient) getHTTPRequest(ctx context.Context, in interface{}, method, path string) (*http.Request, error) {
	target, err := url.Parse(r.instance)
	if err != nil {
		return nil, fmt.Errorf("invalid instance %s", r.instance)
	}
	target.Path = path
	req, err := http.NewRequest(method, target.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request (%s)", err)
	}
	r.updateHTTPHeader(ctx, &req.Header)
	if err = encodeHTTPRequest(ctx, req, in); err != nil {
		return nil, fmt.Errorf("could not encode request (%s)", err)
	}
	return req, nil
}

//
func makeURIAggWatchV1AutoWatchSvcAggWatchV1WatchOper(in *api.AggWatchOptions) string {
	return ""

}

// MakeAggWatchV1RestClientEndpoints make REST client endpoints
func MakeAggWatchV1RestClientEndpoints(instance string, httpClient *http.Client) (EndpointsAggWatchV1RestClient, error) {
	if !strings.HasPrefix(instance, "https") {
		instance = "https://" + instance
	}

	return EndpointsAggWatchV1RestClient{
		instance: instance,
		client:   httpClient,
	}, nil

}

// MakeAggWatchV1StagedRestClientEndpoints makes staged REST client endpoints
func MakeAggWatchV1StagedRestClientEndpoints(instance string, bufferId string, httpClient *http.Client) (EndpointsAggWatchV1RestClient, error) {
	if !strings.HasPrefix(instance, "https") {
		instance = "https://" + instance
	}

	return EndpointsAggWatchV1RestClient{
		instance: instance,
		bufferId: bufferId,
		client:   httpClient,
	}, nil
}
