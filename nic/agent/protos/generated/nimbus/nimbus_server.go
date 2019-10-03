// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package nimbus

import (
	"errors"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/log"
	memdb "github.com/pensando/sw/venice/utils/memdb"
	"github.com/pensando/sw/venice/utils/rpckit"
	"github.com/pensando/sw/venice/utils/tsdb"
)

var (
	// ErrIncorrectObjectType is returned when type deferencing from memdb.Object is of invalid type
	ErrIncorrectObjectType = errors.New("incorrect object type")
)

const (
	// DefaultWatchBatchSize is the max batch size for watch events
	DefaultWatchBatchSize = 100
	// DefaultWatchHoldInterval is the time that the batching logic waits to accumulate events
	DefaultWatchHoldInterval = time.Millisecond * 10
)

type EventStatus struct {
	LastObjectMeta *api.ObjectMeta
}

// MbusServer is the message bus server
type MbusServer struct {
	svcName    string              // service name
	grpcServer *rpckit.RPCServer   // gRPC server instance
	listenURL  string              // URL to listen on
	memDB      *memdb.Memdb        // database of all objects
	stats      map[string]tsdb.Obj // nimbus stats
}

// AddObject adds object to mbus
func (ms *MbusServer) AddObject(obj memdb.Object) error {
	ms.Stats(obj.GetObjectKind(), "AddEvent").Inc()
	ms.Stats(obj.GetObjectKind(), "ObjectCount").Inc()

	return ms.memDB.AddObject(obj)
}

// UpdateObject updates an object in mbus
func (ms *MbusServer) UpdateObject(obj memdb.Object) error {
	ms.Stats(obj.GetObjectKind(), "UpdateEvent").Inc()

	return ms.memDB.UpdateObject(obj)
}

// FindObject finds the memdb object
func (ms *MbusServer) FindObject(obj memdb.Object) (memdb.Object, error) {
	return ms.memDB.FindObject(obj.GetObjectKind(), obj.GetObjectMeta())
}

// DeleteObject deletes an object from mbus
func (ms *MbusServer) DeleteObject(obj memdb.Object) error {
	ms.Stats(obj.GetObjectKind(), "DeleteEvent").Inc()
	ms.Stats(obj.GetObjectKind(), "ObjectCount").Dec()

	return ms.memDB.DeleteObject(obj)
}

// AddNodeState adds node state to an object
func (ms *MbusServer) AddNodeState(nodeID string, obj memdb.Object) error {
	return ms.memDB.AddNodeState(nodeID, obj)
}

// DelNodeState deletes node state from an object
func (ms *MbusServer) DelNodeState(nodeID string, obj memdb.Object) error {
	return ms.memDB.DelNodeState(nodeID, obj)
}

// Stats returns a counter for stats
func (ms *MbusServer) Stats(kind, cname string) api.Counter {
	var err error
	tsdbObj, ok := ms.stats[kind]
	if !ok {
		keyTags := map[string]string{"node": "venice", "module": ms.svcName, "kind": kind}
		tsdbObj, err = tsdb.NewObj("NimbusStats", keyTags, nil, nil)
		if err != nil {
			log.Fatalf("unable to create tsdb object, keys %+v", keyTags)
			return nil
		}
	}

	return tsdbObj.Counter(cname)
}

// DumpDatabase dumps the entire database as json
func (ms *MbusServer) DumpDatabase() ([]byte, error) {
	return ms.memDB.MarshalJSON()
}

// NewMbusServer creates a new instance of message bus server
func NewMbusServer(svcName string, grpcServer *rpckit.RPCServer) *MbusServer {
	mbusServer := MbusServer{
		svcName:    svcName,
		memDB:      memdb.NewMemdb(),
		stats:      make(map[string]tsdb.Obj),
		grpcServer: grpcServer,
	}
	return &mbusServer
}
