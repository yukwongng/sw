// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package bookstoreGwService is a auto generated package.
Input file: protos/example.proto
*/
package bookstoreGwService

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pensando/grpc-gateway/runtime"
	"github.com/pensando/sw/api"
	bookstore "github.com/pensando/sw/api/generated/bookstore"
	"github.com/pensando/sw/api/generated/bookstore/grpc/client"
	"github.com/pensando/sw/venice/apigw/pkg"
	"github.com/pensando/sw/venice/utils/balancer"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta

type sBookstoreV1GwService struct {
	logger log.Logger
}

type adapterBookstoreV1 struct {
	service bookstore.ServiceBookstoreV1Client
}

func (a adapterBookstoreV1) AutoAddBook(oldctx oldcontext.Context, t *bookstore.Book, options ...grpc.CallOption) (*bookstore.Book, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoAddBook(ctx, t)
}

func (a adapterBookstoreV1) AutoAddOrder(oldctx oldcontext.Context, t *bookstore.Order, options ...grpc.CallOption) (*bookstore.Order, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoAddOrder(ctx, t)
}

func (a adapterBookstoreV1) AutoAddPublisher(oldctx oldcontext.Context, t *bookstore.Publisher, options ...grpc.CallOption) (*bookstore.Publisher, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoAddPublisher(ctx, t)
}

func (a adapterBookstoreV1) AutoDeleteBook(oldctx oldcontext.Context, t *bookstore.Book, options ...grpc.CallOption) (*bookstore.Book, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoDeleteBook(ctx, t)
}

func (a adapterBookstoreV1) AutoDeleteOrder(oldctx oldcontext.Context, t *bookstore.Order, options ...grpc.CallOption) (*bookstore.Order, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoDeleteOrder(ctx, t)
}

func (a adapterBookstoreV1) AutoDeletePublisher(oldctx oldcontext.Context, t *bookstore.Publisher, options ...grpc.CallOption) (*bookstore.Publisher, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoDeletePublisher(ctx, t)
}

func (a adapterBookstoreV1) AutoGetBook(oldctx oldcontext.Context, t *bookstore.Book, options ...grpc.CallOption) (*bookstore.Book, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoGetBook(ctx, t)
}

func (a adapterBookstoreV1) AutoGetOrder(oldctx oldcontext.Context, t *bookstore.Order, options ...grpc.CallOption) (*bookstore.Order, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoGetOrder(ctx, t)
}

func (a adapterBookstoreV1) AutoGetPublisher(oldctx oldcontext.Context, t *bookstore.Publisher, options ...grpc.CallOption) (*bookstore.Publisher, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoGetPublisher(ctx, t)
}

func (a adapterBookstoreV1) AutoListBook(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*bookstore.BookList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoListBook(ctx, t)
}

func (a adapterBookstoreV1) AutoListOrder(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*bookstore.OrderList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoListOrder(ctx, t)
}

func (a adapterBookstoreV1) AutoListPublisher(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*bookstore.PublisherList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoListPublisher(ctx, t)
}

func (a adapterBookstoreV1) AutoUpdateBook(oldctx oldcontext.Context, t *bookstore.Book, options ...grpc.CallOption) (*bookstore.Book, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoUpdateBook(ctx, t)
}

func (a adapterBookstoreV1) AutoUpdateOrder(oldctx oldcontext.Context, t *bookstore.Order, options ...grpc.CallOption) (*bookstore.Order, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoUpdateOrder(ctx, t)
}

func (a adapterBookstoreV1) AutoUpdatePublisher(oldctx oldcontext.Context, t *bookstore.Publisher, options ...grpc.CallOption) (*bookstore.Publisher, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	return a.service.AutoUpdatePublisher(ctx, t)
}

func (a adapterBookstoreV1) AutoWatchOrder(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (bookstore.BookstoreV1_AutoWatchOrderClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchOrder(ctx, in)
}

func (a adapterBookstoreV1) AutoWatchBook(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (bookstore.BookstoreV1_AutoWatchBookClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchBook(ctx, in)
}

func (a adapterBookstoreV1) AutoWatchPublisher(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (bookstore.BookstoreV1_AutoWatchPublisherClient, error) {
	ctx := context.Context(oldctx)
	return a.service.AutoWatchPublisher(ctx, in)
}

func (e *sBookstoreV1GwService) CompleteRegistration(ctx context.Context,
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
	err = bookstore.RegisterBookstoreV1HandlerWithClient(ctx, mux, cl)
	if err != nil {
		err = errors.Wrap(err, "service registration failed")
		return err
	}
	logger.InfoLog("msg", "registered service bookstore.BookstoreV1")

	m.Handle("/v1/bookstore/", http.StripPrefix("/v1/bookstore", mux))
	if fileCount == 1 {
		err = registerSwaggerDef(m, logger)
	}
	return err
}

func (e *sBookstoreV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface) (bookstore.BookstoreV1Client, error) {
	var opts []rpckit.Option
	if rslvr != nil {
		opts = append(opts, rpckit.WithBalancer(balancer.New(rslvr)))
	}
	client, err := rpckit.NewRPCClient("BookstoreV1GwService", grpcAddr, opts...)
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

	cl := adapterBookstoreV1{grpcclient.NewBookstoreV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {
	apigw := apigwpkg.MustGetAPIGateway()

	svcBookstoreV1 := sBookstoreV1GwService{}
	apigw.Register("bookstore.BookstoreV1", "bookstore/", &svcBookstoreV1)
}
