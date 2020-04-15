// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package staging

import (
	"context"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/kvstore"
)

// Dummy vars to suppress unused imports message
var _ context.Context
var _ api.ObjectMeta
var _ kvstore.Interface

const KindBuffer ObjKind = "Buffer"
const KindBulkEditAction ObjKind = "BulkEditAction"
const KindClearAction ObjKind = "ClearAction"
const KindCommitAction ObjKind = "CommitAction"
