// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package cmdGwService is a auto generated package.
Input file: protos/cmd.proto
*/
package cmdGwService

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pensando/grpc-gateway/runtime"
	"github.com/pensando/sw/api"
	cmd "github.com/pensando/sw/api/generated/cmd"
	"github.com/pensando/sw/api/generated/cmd/grpc/client"
	"github.com/pensando/sw/apigw/pkg"
	"github.com/pensando/sw/utils/balancer"
	"github.com/pensando/sw/utils/log"
	"github.com/pensando/sw/utils/resolver"
	"github.com/pensando/sw/utils/rpckit"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta

type sCmdV1GwService struct {
	logger log.Logger
}

type adapterCmdV1 struct {
	service cmd.ServiceCmdV1Client
}

func (a adapterCmdV1) AutoAddCluster(oldctx oldcontext.Context, t *cmd.Cluster, options ...grpc.CallOption) (*cmd.Cluster, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoAddCluster(ctx, t)
}

func (a adapterCmdV1) AutoAddNode(oldctx oldcontext.Context, t *cmd.Node, options ...grpc.CallOption) (*cmd.Node, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoAddNode(ctx, t)
}

func (a adapterCmdV1) AutoAddSmartNIC(oldctx oldcontext.Context, t *cmd.SmartNIC, options ...grpc.CallOption) (*cmd.SmartNIC, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoAddSmartNIC(ctx, t)
}

func (a adapterCmdV1) AutoDeleteCluster(oldctx oldcontext.Context, t *cmd.Cluster, options ...grpc.CallOption) (*cmd.Cluster, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoDeleteCluster(ctx, t)
}

func (a adapterCmdV1) AutoDeleteNode(oldctx oldcontext.Context, t *cmd.Node, options ...grpc.CallOption) (*cmd.Node, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoDeleteNode(ctx, t)
}

func (a adapterCmdV1) AutoDeleteSmartNIC(oldctx oldcontext.Context, t *cmd.SmartNIC, options ...grpc.CallOption) (*cmd.SmartNIC, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoDeleteSmartNIC(ctx, t)
}

func (a adapterCmdV1) AutoGetCluster(oldctx oldcontext.Context, t *cmd.Cluster, options ...grpc.CallOption) (*cmd.Cluster, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoGetCluster(ctx, t)
}

func (a adapterCmdV1) AutoGetNode(oldctx oldcontext.Context, t *cmd.Node, options ...grpc.CallOption) (*cmd.Node, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoGetNode(ctx, t)
}

func (a adapterCmdV1) AutoGetSmartNIC(oldctx oldcontext.Context, t *cmd.SmartNIC, options ...grpc.CallOption) (*cmd.SmartNIC, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoGetSmartNIC(ctx, t)
}

func (a adapterCmdV1) AutoListCluster(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*cmd.ClusterList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoListCluster(ctx, t)
}

func (a adapterCmdV1) AutoListNode(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*cmd.NodeList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoListNode(ctx, t)
}

func (a adapterCmdV1) AutoListSmartNIC(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*cmd.SmartNICList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoListSmartNIC(ctx, t)
}

func (a adapterCmdV1) AutoUpdateCluster(oldctx oldcontext.Context, t *cmd.Cluster, options ...grpc.CallOption) (*cmd.Cluster, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoUpdateCluster(ctx, t)
}

func (a adapterCmdV1) AutoUpdateNode(oldctx oldcontext.Context, t *cmd.Node, options ...grpc.CallOption) (*cmd.Node, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoUpdateNode(ctx, t)
}

func (a adapterCmdV1) AutoUpdateSmartNIC(oldctx oldcontext.Context, t *cmd.SmartNIC, options ...grpc.CallOption) (*cmd.SmartNIC, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoUpdateSmartNIC(ctx, t)
}

func (a adapterCmdV1) AutoWatchCluster(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (cmd.CmdV1_AutoWatchClusterClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchCluster(ctx, in)
}

func (a adapterCmdV1) AutoWatchNode(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (cmd.CmdV1_AutoWatchNodeClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchNode(ctx, in)
}

func (a adapterCmdV1) AutoWatchSmartNIC(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (cmd.CmdV1_AutoWatchSmartNICClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchSmartNIC(ctx, in)
}

func (e *sCmdV1GwService) CompleteRegistration(ctx context.Context,
	logger log.Logger,
	grpcserver *grpc.Server,
	m *http.ServeMux,
	rslvr resolver.Interface) error {
	apigw := apigwpkg.MustGetAPIGateway()
	// IP:port destination or service discovery key.
	grpcaddr := "pen-apiserver"
	grpcaddr = apigw.GetAPIServerAddr(grpcaddr)
	e.logger = logger
	cl, err := e.newClient(ctx, grpcaddr, rslvr)
	if cl == nil || err != nil {
		err = errors.Wrap(err, "could not create client")
		return err
	}
	marshaller := runtime.JSONBuiltin{}
	opts := runtime.WithMarshalerOption("*", &marshaller)
	if mux == nil {
		mux = runtime.NewServeMux(opts)
	}
	fileCount++
	err = cmd.RegisterCmdV1HandlerWithClient(ctx, mux, cl)
	if err != nil {
		err = errors.Wrap(err, "service registration failed")
		return err
	}
	logger.InfoLog("msg", "registered service cmd.CmdV1")

	m.Handle("/v1/cmd/", http.StripPrefix("/v1/cmd", mux))
	if fileCount == 1 {
		err = registerSwaggerDef(m, logger)
	}
	return err
}

func (e *sCmdV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface) (cmd.CmdV1Client, error) {
	var opts []rpckit.Option
	if rslvr != nil {
		opts = append(opts, rpckit.WithBalancer(balancer.New(rslvr)))
	}
	client, err := rpckit.NewRPCClient("CmdV1GwService", grpcAddr, opts...)
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

	cl := adapterCmdV1{grpcclient.NewCmdV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {
	apigw := apigwpkg.MustGetAPIGateway()

	svcCmdV1 := sCmdV1GwService{}
	apigw.Register("cmd.CmdV1", "cmd/", &svcCmdV1)
}
