// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package rpckit

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	hdr "github.com/pensando/sw/venice/utils/histogram"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"
)

// middleware callback roles
const (
	RoleClient = "Client"
	RoleServer = "Server"
)

func addMeta(ctx context.Context, key, value string) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok == false {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}
	newmd := metadata.Join(md, metadata.Pairs(key, value))
	return metadata.NewOutgoingContext(ctx, newmd)
}

// rpcServerUnaryInterceptor returns an intercept handler for unary rpc calls
func rpcServerUnaryInterceptor(rpcServer *RPCServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		name := fmt.Sprintf("%s/%s", rpcServer.mysvcName, info.FullMethod)
		rpcstrt := time.Now()
		// call all the request middlewares
		for _, m := range rpcServer.middlewares {
			ctx = m.ReqInterceptor(ctx, RoleServer, rpcServer.mysvcName, info.FullMethod, req)
		}
		hdlrstrt := time.Now()
		// finally call the handler
		resp, err := handler(ctx, req)
		hdr.Record("handler:"+name, time.Since(hdlrstrt))
		// call all response middlewares
		for _, m := range rpcServer.middlewares {
			ctx = m.RespInterceptor(ctx, RoleServer, rpcServer.mysvcName, info.FullMethod, req, resp, err)
		}
		hdr.Record("rpc:"+name, time.Since(rpcstrt))
		return resp, err
	}
}

// rpcServerStreamInterceptor  is the server intercept handler for stream RPCs
func rpcServerStreamInterceptor(rpcServer *RPCServer) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// call all the request middlewares
		ctx := ss.Context()
		for _, m := range rpcServer.middlewares {
			ctx = m.ReqInterceptor(ctx, RoleServer, rpcServer.mysvcName, info.FullMethod, struct{}{})
		}

		return handler(srv, ss)
	}
}

// rpcClientUnaryInterceptor intercepts the client rpc calls
func rpcClientUnaryInterceptor(rpcClient *RPCClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// call all the request middlewares
		for _, m := range rpcClient.middlewares {
			ctx = m.ReqInterceptor(ctx, RoleClient, rpcClient.mysvcName, method, req)
		}

		ctx = addMeta(ctx, "nodeuuid", rpcClient.nodeuuid)

		// finally, call the invoker
		err := invoker(ctx, method, req, reply, cc, opts...)

		// call all the response middlewares
		for _, m := range rpcClient.middlewares {
			ctx = m.RespInterceptor(ctx, RoleClient, rpcClient.mysvcName, method, req, reply, err)
		}

		return err
	}
}

// rpcClientStreamInterceptor intercepts client rpc calls
// FIXME: to be implemented
func rpcClientStreamInterceptor(rpcClient *RPCClient) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

		ctx = addMeta(ctx, "nodeuuid", rpcClient.nodeuuid)

		return streamer(ctx, desc, cc, method, opts...)
	}
}

// Logging middleware
type logMiddleware struct {
	// empty struct
}

// newLogMiddleware returns a Logging middleware
func newLogMiddleware() *logMiddleware {
	return &logMiddleware{}
}

// ReqInterceptor implements request interception
func (l *logMiddleware) ReqInterceptor(ctx context.Context, role, mysvcName, method string, req interface{}) context.Context {
	var reqStr string

	meta, err := runtime.GetObjectMeta(req)
	if err != nil {
		reqStr = fmt.Sprintf("{%+v}", meta)
	} else {
		reqStr = ""
	}

	switch role {
	case RoleClient:
		log.Debugf("Client %s Making RPC request: %s() Req: %s", mysvcName, method, reqStr)
	case RoleServer:
		log.Debugf("Server %s received RPC: %s() Req: %s", mysvcName, method, reqStr)
	}

	return ctx
}

// RespInterceptor implements response interception
func (l *logMiddleware) RespInterceptor(ctx context.Context, role, mysvcName, method string, req, reply interface{}, err error) context.Context {
	var replyStr string

	meta, err2 := runtime.GetObjectMeta(reply)
	if err2 != nil {
		replyStr = fmt.Sprintf("{%+v}", meta)
	} else {
		replyStr = ""
	}

	switch role {
	case RoleClient:
		if err != nil {
			log.Errorf("Client %s received RPC response: %s() Resp: %s, error: %v", mysvcName, method, replyStr, err)
		} else {
			log.Debugf("Client %s received RPC response: %s() Resp: %s, error: %v", mysvcName, method, replyStr, err)
		}
	case RoleServer:
		if err != nil {
			log.Errorf("Server %s returning RPC response: %s() Resp: %s, error: %v", mysvcName, method, replyStr, err)
		} else {
			log.Debugf("Server %s returning RPC response: %s() Resp: %s, error: %v", mysvcName, method, replyStr, err)
		}
	}

	return ctx
}

// Stats middleware
type statsMiddleware struct {
}

// newStatsMiddleware returns a new stats middleware
func newStatsMiddleware(svcName, role string) *statsMiddleware {
	return &statsMiddleware{}
}

// ReqInterceptor implements request interception
func (s *statsMiddleware) ReqInterceptor(ctx context.Context, role, mysvcName, method string, req interface{}) context.Context {
	switch role {
	case RoleClient:
		// get existing metadata
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}
		newMd := metadata.Join(md, metadata.Pairs("cname", mysvcName))
		ctx = metadata.NewOutgoingContext(ctx, newMd)
		// increment request stats
		singletonMap.Add(fmt.Sprintf("%v/client/requests%v", mysvcName, method), 1)
	case RoleServer:
		// increment total request stats
		singletonMap.Add(fmt.Sprintf("%v/server/requests%v", mysvcName, method), 1)

		// get grpc metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			break
		}
		clientNames, ok := md["cname"]
		if !ok {
			break
		}
		singletonMap.Add(fmt.Sprintf("%v/server/requests%v/client/%v", mysvcName, method, clientNames[0]), 1)
	}

	return ctx
}

// RespInterceptor handles responses
func (s *statsMiddleware) RespInterceptor(ctx context.Context, role, mysvcName, method string, req, reply interface{}, err error) context.Context {

	switch role {
	case RoleClient:
		// increment response and error stats
		singletonMap.Add(fmt.Sprintf("%v/client/responses%v", mysvcName, method), 1)
		if err != nil {
			singletonMap.Add(fmt.Sprintf("%v/client/errors%v", mysvcName, method), 1)
		}
	case RoleServer:
		// increment total response and error stats
		singletonMap.Add(fmt.Sprintf("%v/server/responses%v", mysvcName, method), 1)
		if err != nil {
			singletonMap.Add(fmt.Sprintf("%v/server/errors%v", mysvcName, method), 1)
		}

		// increment per client response and error stats
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			break
		}
		clientNames, ok := md["cname"]
		if !ok {
			break
		}
		singletonMap.Add(fmt.Sprintf("%v/server/responses%v/client/%v", mysvcName, method, clientNames[0]), 1)
		if err != nil {
			singletonMap.Add(fmt.Sprintf("%v/server/errors%v/client/%v", mysvcName, method, clientNames[0]), 1)
		}
	}

	return ctx
}
