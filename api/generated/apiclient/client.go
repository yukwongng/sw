// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package apiclient

import (
	gogocodec "github.com/gogo/protobuf/codec"
	"google.golang.org/grpc"

	bookstore "github.com/pensando/sw/api/generated/bookstore"
	bookstoreClient "github.com/pensando/sw/api/generated/bookstore/grpc/client"
	cmd "github.com/pensando/sw/api/generated/cmd"
	cmdClient "github.com/pensando/sw/api/generated/cmd/grpc/client"
	network "github.com/pensando/sw/api/generated/network"
	networkClient "github.com/pensando/sw/api/generated/network/grpc/client"
	networkencryption "github.com/pensando/sw/api/generated/networkencryption"
	networkencryptionClient "github.com/pensando/sw/api/generated/networkencryption/grpc/client"
	"github.com/pensando/sw/utils/log"
	"github.com/pensando/sw/utils/rpckit"
)

const codecSize = 1024 * 1024

// Services is list of all services exposed by the client ---
type Services interface {
	Close() error

	// Package is bookstore and len of messages is 3
	BookstoreV1() bookstore.BookstoreV1Interface
	// Package is cmd and len of messages is 3
	CmdV1() cmd.CmdV1Interface
	// Package is network and len of messages is 1
	EndpointV1() network.EndpointV1Interface
	// Package is network and len of messages is 1
	LbPolicyV1() network.LbPolicyV1Interface
	// Package is network and len of messages is 1
	NetworkV1() network.NetworkV1Interface
	// Package is network and len of messages is 1
	SecurityGroupV1() network.SecurityGroupV1Interface
	// Package is network and len of messages is 1
	ServiceV1() network.ServiceV1Interface
	// Package is network and len of messages is 1
	SgpolicyV1() network.SgpolicyV1Interface
	// Package is network and len of messages is 1
	TenantV1() network.TenantV1Interface
	// Package is networkencryption and len of messages is 1
	TrafficEncryptionPolicyV1() networkencryption.TrafficEncryptionPolicyV1Interface
}

type apiGrpcServerClient struct {
	url    string
	logger log.Logger
	client *rpckit.RPCClient

	aBookstoreV1               bookstore.BookstoreV1Interface
	aCmdV1                     cmd.CmdV1Interface
	aEndpointV1                network.EndpointV1Interface
	aLbPolicyV1                network.LbPolicyV1Interface
	aNetworkV1                 network.NetworkV1Interface
	aSecurityGroupV1           network.SecurityGroupV1Interface
	aServiceV1                 network.ServiceV1Interface
	aSgpolicyV1                network.SgpolicyV1Interface
	aTenantV1                  network.TenantV1Interface
	aTrafficEncryptionPolicyV1 networkencryption.TrafficEncryptionPolicyV1Interface
}

// Close closes the client
func (a *apiGrpcServerClient) Close() error {
	return a.client.Close()
}

func (a *apiGrpcServerClient) BookstoreV1() bookstore.BookstoreV1Interface {
	return a.aBookstoreV1
}

func (a *apiGrpcServerClient) CmdV1() cmd.CmdV1Interface {
	return a.aCmdV1
}

func (a *apiGrpcServerClient) EndpointV1() network.EndpointV1Interface {
	return a.aEndpointV1
}

func (a *apiGrpcServerClient) LbPolicyV1() network.LbPolicyV1Interface {
	return a.aLbPolicyV1
}

func (a *apiGrpcServerClient) NetworkV1() network.NetworkV1Interface {
	return a.aNetworkV1
}

func (a *apiGrpcServerClient) SecurityGroupV1() network.SecurityGroupV1Interface {
	return a.aSecurityGroupV1
}

func (a *apiGrpcServerClient) ServiceV1() network.ServiceV1Interface {
	return a.aServiceV1
}

func (a *apiGrpcServerClient) SgpolicyV1() network.SgpolicyV1Interface {
	return a.aSgpolicyV1
}

func (a *apiGrpcServerClient) TenantV1() network.TenantV1Interface {
	return a.aTenantV1
}

func (a *apiGrpcServerClient) TrafficEncryptionPolicyV1() networkencryption.TrafficEncryptionPolicyV1Interface {
	return a.aTrafficEncryptionPolicyV1
}

// NewGrpcAPIClient returns a gRPC client
func NewGrpcAPIClient(url string, logger log.Logger, opts ...grpc.DialOption) (Services, error) {
	client, err := rpckit.NewRPCClient("ApiClient", url, rpckit.WithCodec(gogocodec.New(codecSize)))
	if err != nil {
		logger.ErrorLog("msg", "Failed to connect to gRPC server", "URL", url, "error", err)
		return nil, err
	}
	return &apiGrpcServerClient{
		url:    url,
		client: client,
		logger: logger,

		aBookstoreV1:               bookstoreClient.NewGrpcCrudClientBookstoreV1(client.ClientConn, logger),
		aCmdV1:                     cmdClient.NewGrpcCrudClientCmdV1(client.ClientConn, logger),
		aEndpointV1:                networkClient.NewGrpcCrudClientEndpointV1(client.ClientConn, logger),
		aLbPolicyV1:                networkClient.NewGrpcCrudClientLbPolicyV1(client.ClientConn, logger),
		aNetworkV1:                 networkClient.NewGrpcCrudClientNetworkV1(client.ClientConn, logger),
		aSecurityGroupV1:           networkClient.NewGrpcCrudClientSecurityGroupV1(client.ClientConn, logger),
		aServiceV1:                 networkClient.NewGrpcCrudClientServiceV1(client.ClientConn, logger),
		aSgpolicyV1:                networkClient.NewGrpcCrudClientSgpolicyV1(client.ClientConn, logger),
		aTenantV1:                  networkClient.NewGrpcCrudClientTenantV1(client.ClientConn, logger),
		aTrafficEncryptionPolicyV1: networkencryptionClient.NewGrpcCrudClientTrafficEncryptionPolicyV1(client.ClientConn, logger),
	}, nil
}

type apiRestServerClient struct {
	url string

	aBookstoreV1               bookstore.BookstoreV1Interface
	aCmdV1                     cmd.CmdV1Interface
	aEndpointV1                network.EndpointV1Interface
	aLbPolicyV1                network.LbPolicyV1Interface
	aNetworkV1                 network.NetworkV1Interface
	aSecurityGroupV1           network.SecurityGroupV1Interface
	aServiceV1                 network.ServiceV1Interface
	aSgpolicyV1                network.SgpolicyV1Interface
	aTenantV1                  network.TenantV1Interface
	aTrafficEncryptionPolicyV1 networkencryption.TrafficEncryptionPolicyV1Interface
}

// Close sloses the client
func (a *apiRestServerClient) Close() error {
	return nil
}

func (a *apiRestServerClient) BookstoreV1() bookstore.BookstoreV1Interface {
	return a.aBookstoreV1
}

func (a *apiRestServerClient) CmdV1() cmd.CmdV1Interface {
	return a.aCmdV1
}

func (a *apiRestServerClient) EndpointV1() network.EndpointV1Interface {
	return a.aEndpointV1
}

func (a *apiRestServerClient) LbPolicyV1() network.LbPolicyV1Interface {
	return a.aLbPolicyV1
}

func (a *apiRestServerClient) NetworkV1() network.NetworkV1Interface {
	return a.aNetworkV1
}

func (a *apiRestServerClient) SecurityGroupV1() network.SecurityGroupV1Interface {
	return a.aSecurityGroupV1
}

func (a *apiRestServerClient) ServiceV1() network.ServiceV1Interface {
	return a.aServiceV1
}

func (a *apiRestServerClient) SgpolicyV1() network.SgpolicyV1Interface {
	return a.aSgpolicyV1
}

func (a *apiRestServerClient) TenantV1() network.TenantV1Interface {
	return a.aTenantV1
}

func (a *apiRestServerClient) TrafficEncryptionPolicyV1() networkencryption.TrafficEncryptionPolicyV1Interface {
	return a.aTrafficEncryptionPolicyV1
}

// NewRestAPIClient returns a REST client
func NewRestAPIClient(url string) (Services, error) {
	return &apiRestServerClient{
		url: url,

		aBookstoreV1:               bookstoreClient.NewRestCrudClientBookstoreV1(url),
		aCmdV1:                     cmdClient.NewRestCrudClientCmdV1(url),
		aEndpointV1:                networkClient.NewRestCrudClientEndpointV1(url),
		aLbPolicyV1:                networkClient.NewRestCrudClientLbPolicyV1(url),
		aNetworkV1:                 networkClient.NewRestCrudClientNetworkV1(url),
		aSecurityGroupV1:           networkClient.NewRestCrudClientSecurityGroupV1(url),
		aServiceV1:                 networkClient.NewRestCrudClientServiceV1(url),
		aSgpolicyV1:                networkClient.NewRestCrudClientSgpolicyV1(url),
		aTenantV1:                  networkClient.NewRestCrudClientTenantV1(url),
		aTrafficEncryptionPolicyV1: networkencryptionClient.NewRestCrudClientTrafficEncryptionPolicyV1(url),
	}, nil
}
