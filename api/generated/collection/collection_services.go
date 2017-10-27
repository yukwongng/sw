// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package collection is a auto generated package.
Input file: protos/collection.proto
*/
package collection

import (
	"context"
	"github.com/pensando/sw/api"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

// ServiceCollectionPolicyV1Client  is the client interface for the service.
type ServiceCollectionPolicyV1Client interface {
	AutoAddCollectionPolicy(ctx context.Context, t *CollectionPolicy) (*CollectionPolicy, error)
	AutoDeleteCollectionPolicy(ctx context.Context, t *CollectionPolicy) (*CollectionPolicy, error)
	AutoGetCollectionPolicy(ctx context.Context, t *CollectionPolicy) (*CollectionPolicy, error)
	AutoListCollectionPolicy(ctx context.Context, t *api.ListWatchOptions) (*CollectionPolicyList, error)
	AutoUpdateCollectionPolicy(ctx context.Context, t *CollectionPolicy) (*CollectionPolicy, error)

	AutoWatchCollectionPolicy(ctx context.Context, in *api.ListWatchOptions) (CollectionPolicyV1_AutoWatchCollectionPolicyClient, error)
}

// ServiceCollectionPolicyV1Server is the server interface for the service.
type ServiceCollectionPolicyV1Server interface {
	AutoAddCollectionPolicy(ctx context.Context, t CollectionPolicy) (CollectionPolicy, error)
	AutoDeleteCollectionPolicy(ctx context.Context, t CollectionPolicy) (CollectionPolicy, error)
	AutoGetCollectionPolicy(ctx context.Context, t CollectionPolicy) (CollectionPolicy, error)
	AutoListCollectionPolicy(ctx context.Context, t api.ListWatchOptions) (CollectionPolicyList, error)
	AutoUpdateCollectionPolicy(ctx context.Context, t CollectionPolicy) (CollectionPolicy, error)

	AutoWatchCollectionPolicy(in *api.ListWatchOptions, stream CollectionPolicyV1_AutoWatchCollectionPolicyServer) error
}
