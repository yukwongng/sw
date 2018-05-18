// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringApiServer is a auto generated package.
Input file: mirror.proto
*/
package monitoringApiServer

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/pensando/sw/api"
	monitoring "github.com/pensando/sw/api/generated/monitoring"
	"github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
	"github.com/pensando/sw/venice/utils/runtime"
)

// dummy vars to suppress unused errors
var _ api.ObjectMeta
var _ listerwatcher.WatcherClient
var _ fmt.Stringer

type smonitoringMirrorBackend struct {
	Services map[string]apiserver.Service
	Messages map[string]apiserver.Message
	logger   log.Logger
	scheme   *runtime.Scheme
}

func (s *smonitoringMirrorBackend) regMsgsFunc(l log.Logger, scheme *runtime.Scheme) {
	l.Infof("registering message for smonitoringMirrorBackend")
	s.Messages = map[string]apiserver.Message{

		"monitoring.AppProtoSelector": apisrvpkg.NewMessage("monitoring.AppProtoSelector"),
		"monitoring.MatchRule":        apisrvpkg.NewMessage("monitoring.MatchRule"),
		"monitoring.MatchSelector":    apisrvpkg.NewMessage("monitoring.MatchSelector"),
		"monitoring.MirrorCollector":  apisrvpkg.NewMessage("monitoring.MirrorCollector"),
		"monitoring.MirrorSession": apisrvpkg.NewMessage("monitoring.MirrorSession").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := monitoring.MirrorSession{}
				return r.MakeKey(prefix)
			}
			r := i.(monitoring.MirrorSession)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(monitoring.MirrorSession)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(monitoring.MirrorSession)
			key := r.MakeKey(prefix)
			r.Kind = "MirrorSession"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				err = errors.Wrap(err, "KV create failed")
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*monitoring.MirrorSession)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &monitoring.MirrorSession{}
					err = kvs.ConsistentUpdate(ctx, key, into, updateFunc)
				} else {
					if r.ResourceVersion != "" {
						l.Infof("resource version is specified %s\n", r.ResourceVersion)
						err = kvs.Update(ctx, key, &r, kvstore.Compare(kvstore.WithVersion(key), "=", r.ResourceVersion))
					} else {
						err = kvs.Update(ctx, key, &r)
					}
					err = errors.Wrap(err, "KV update failed")
				}
			}
			return r, err
		}).WithKvTxnUpdater(func(ctx context.Context, txn kvstore.Txn, i interface{}, prefix string, create bool) error {
			r := i.(monitoring.MirrorSession)
			key := r.MakeKey(prefix)
			var err error
			if create {
				err = txn.Create(key, &r)
				err = errors.Wrap(err, "KV transaction create failed")
			} else {
				err = txn.Update(key, &r)
				err = errors.Wrap(err, "KV transaction update failed")
			}
			return err
		}).WithUUIDWriter(func(i interface{}) (interface{}, error) {
			r := i.(monitoring.MirrorSession)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(monitoring.MirrorSession)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(monitoring.MirrorSession)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path, ver, prefix string, i interface{}) (interface{}, error) {
			r := i.(monitoring.MirrorSession)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := monitoring.MirrorSession{}
			err := kvs.Get(ctx, key, &r)
			err = errors.Wrap(err, "KV get failed")
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := monitoring.MirrorSession{}
			err := kvs.Delete(ctx, key, &r)
			return r, err
		}).WithKvTxnDelFunc(func(ctx context.Context, txn kvstore.Txn, key string) error {
			return txn.Delete(key)
		}).WithValidate(func(i interface{}, ver string, ignoreStatus bool) error {
			r := i.(monitoring.MirrorSession)
			if !r.Validate(ver, ignoreStatus) {
				return fmt.Errorf("Default Validation failed")
			}
			return nil
		}),

		"monitoring.MirrorSessionSpec":           apisrvpkg.NewMessage("monitoring.MirrorSessionSpec"),
		"monitoring.MirrorSessionStatus":         apisrvpkg.NewMessage("monitoring.MirrorSessionStatus"),
		"monitoring.MirrorStartConditions":       apisrvpkg.NewMessage("monitoring.MirrorStartConditions"),
		"monitoring.MirrorStopConditions":        apisrvpkg.NewMessage("monitoring.MirrorStopConditions"),
		"monitoring.SmartNICMirrorSessionStatus": apisrvpkg.NewMessage("monitoring.SmartNICMirrorSessionStatus"),
		// Add a message handler for ListWatch options
		"api.ListWatchOptions": apisrvpkg.NewMessage("api.ListWatchOptions"),
	}

	apisrv.RegisterMessages("monitoring", s.Messages)
	// add messages to package.
	if pkgMessages == nil {
		pkgMessages = make(map[string]apiserver.Message)
	}
	for k, v := range s.Messages {
		pkgMessages[k] = v
	}
}

func (s *smonitoringMirrorBackend) regSvcsFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

}

func (s *smonitoringMirrorBackend) regWatchersFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

}

func (s *smonitoringMirrorBackend) CompleteRegistration(ctx context.Context, logger log.Logger,
	grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) error {
	// register all messages in the package if not done already
	s.logger = logger
	s.scheme = scheme
	registerMessages(logger, scheme)
	registerServices(ctx, logger, grpcserver, scheme)
	registerWatchers(ctx, logger, grpcserver, scheme)
	return nil
}

func (s *smonitoringMirrorBackend) Reset() {
	cleanupRegistration()
}

func init() {
	apisrv = apisrvpkg.MustGetAPIServer()

	svc := smonitoringMirrorBackend{}
	addMsgRegFunc(svc.regMsgsFunc)
	addSvcRegFunc(svc.regSvcsFunc)
	addWatcherRegFunc(svc.regWatchersFunc)

	apisrv.Register("monitoring.mirror.proto", &svc)
}
