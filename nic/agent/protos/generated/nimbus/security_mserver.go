// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: security.proto
*/

package nimbus

import (
	"context"
	"errors"
	"io"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/memdb"
	"github.com/pensando/sw/venice/utils/netutils"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// FindSecurityGroup finds an SecurityGroup by object meta
func (ms *MbusServer) FindSecurityGroup(objmeta *api.ObjectMeta) (*netproto.SecurityGroup, error) {
	// find the object
	obj, err := ms.memDB.FindObject("SecurityGroup", objmeta)
	if err != nil {
		return nil, err
	}

	return SecurityGroupFromObj(obj)
}

// ListSecurityGroups lists all SecurityGroups in the mbus
func (ms *MbusServer) ListSecurityGroups(ctx context.Context) ([]*netproto.SecurityGroup, error) {
	var objlist []*netproto.SecurityGroup

	// walk all objects
	objs := ms.memDB.ListObjects("SecurityGroup")
	for _, oo := range objs {
		obj, err := SecurityGroupFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// SecurityGroupStatusReactor is the reactor interface implemented by controllers
type SecurityGroupStatusReactor interface {
	OnSecurityGroupCreateReq(nodeID string, objinfo *netproto.SecurityGroup) error
	OnSecurityGroupUpdateReq(nodeID string, objinfo *netproto.SecurityGroup) error
	OnSecurityGroupDeleteReq(nodeID string, objinfo *netproto.SecurityGroup) error
	OnSecurityGroupOperUpdate(nodeID string, objinfo *netproto.SecurityGroup) error
	OnSecurityGroupOperDelete(nodeID string, objinfo *netproto.SecurityGroup) error
}

// SecurityGroupTopic is the SecurityGroup topic on message bus
type SecurityGroupTopic struct {
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor SecurityGroupStatusReactor // status event reactor
}

// AddSecurityGroupTopic returns a network RPC server
func AddSecurityGroupTopic(server *MbusServer, reactor SecurityGroupStatusReactor) (*SecurityGroupTopic, error) {
	// RPC handler instance
	handler := SecurityGroupTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterSecurityGroupApiServer(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

// CreateSecurityGroup creates SecurityGroup
func (eh *SecurityGroupTopic) CreateSecurityGroup(ctx context.Context, objinfo *netproto.SecurityGroup) (*netproto.SecurityGroup, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received CreateSecurityGroup from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnSecurityGroupCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("SecurityGroup", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateSecurityGroup updates SecurityGroup
func (eh *SecurityGroupTopic) UpdateSecurityGroup(ctx context.Context, objinfo *netproto.SecurityGroup) (*netproto.SecurityGroup, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateSecurityGroup from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("SecurityGroup", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnSecurityGroupUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteSecurityGroup deletes an SecurityGroup
func (eh *SecurityGroupTopic) DeleteSecurityGroup(ctx context.Context, objinfo *netproto.SecurityGroup) (*netproto.SecurityGroup, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteSecurityGroup from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("SecurityGroup", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnSecurityGroupDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// SecurityGroupFromObj converts memdb object to SecurityGroup
func SecurityGroupFromObj(obj memdb.Object) (*netproto.SecurityGroup, error) {
	switch obj.(type) {
	case *netproto.SecurityGroup:
		eobj := obj.(*netproto.SecurityGroup)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetSecurityGroup returns a specific SecurityGroup
func (eh *SecurityGroupTopic) GetSecurityGroup(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.SecurityGroup, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("SecurityGroup", objmeta)
	if err != nil {
		return nil, err
	}

	return SecurityGroupFromObj(obj)
}

// ListSecurityGroups lists all SecurityGroups matching object selector
func (eh *SecurityGroupTopic) ListSecurityGroups(ctx context.Context, objsel *api.ObjectMeta) (*netproto.SecurityGroupList, error) {
	var objlist netproto.SecurityGroupList

	// walk all objects
	objs := eh.server.memDB.ListObjects("SecurityGroup")
	for _, oo := range objs {
		obj, err := SecurityGroupFromObj(oo)
		if err == nil {
			objlist.SecurityGroups = append(objlist.SecurityGroups, obj)
		}
	}

	return &objlist, nil
}

// WatchSecurityGroups watches SecurityGroups and sends streaming resp
func (eh *SecurityGroupTopic) WatchSecurityGroups(ometa *api.ObjectMeta, stream netproto.SecurityGroupApi_WatchSecurityGroupsServer) error {
	// watch for changes
	watchChan := make(chan memdb.Event, memdb.WatchLen)
	defer close(watchChan)
	eh.server.memDB.WatchObjects("SecurityGroup", watchChan)
	defer eh.server.memDB.StopWatchObjects("SecurityGroup", watchChan)

	// get a list of all SecurityGroups
	objlist, err := eh.ListSecurityGroups(context.Background(), ometa)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	// increment stats
	eh.server.Stats("SecurityGroup", "ActiveWatch").Inc()
	eh.server.Stats("SecurityGroup", "WatchConnect").Inc()
	defer eh.server.Stats("SecurityGroup", "ActiveWatch").Dec()
	defer eh.server.Stats("SecurityGroup", "WatchDisconnect").Inc()

	ctx := stream.Context()

	// walk all SecurityGroups and send it out
	for _, obj := range objlist.SecurityGroups {
		watchEvt := netproto.SecurityGroupEvent{
			EventType:     api.EventType_CreateEvent,
			SecurityGroup: *obj,
		}
		err = stream.Send(&watchEvt)
		if err != nil {
			log.Errorf("Error sending SecurityGroup to stream. Err: %v", err)
			return err
		}
	}

	// loop forever on watch channel
	for {
		select {
		// read from channel
		case evt, ok := <-watchChan:
			if !ok {
				log.Errorf("Error reading from channel. Closing watch")
				return errors.New("Error reading from channel")
			}

			// convert the events
			var etype api.EventType
			switch evt.EventType {
			case memdb.CreateEvent:
				etype = api.EventType_CreateEvent
			case memdb.UpdateEvent:
				etype = api.EventType_UpdateEvent
			case memdb.DeleteEvent:
				etype = api.EventType_DeleteEvent
			}

			// get the object
			obj, err := SecurityGroupFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.SecurityGroupEvent{
				EventType:     etype,
				SecurityGroup: *obj,
			}

			// streaming send
			err = stream.Send(&watchEvt)
			if err != nil {
				log.Errorf("Error sending SecurityGroup to stream. Err: %v", err)
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// done
}

// updateSecurityGroupOper triggers oper update callbacks
func (eh *SecurityGroupTopic) updateSecurityGroupOper(oper *netproto.SecurityGroupEvent, nodeID string) error {
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("SecurityGroup", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			return eh.statusReactor.OnSecurityGroupOperUpdate(nodeID, &oper.SecurityGroup)
		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("SecurityGroup", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {
			eh.statusReactor.OnSecurityGroupOperDelete(nodeID, &oper.SecurityGroup)
		}
	}

	return nil
}

func (eh *SecurityGroupTopic) SecurityGroupOperUpdate(stream netproto.SecurityGroupApi_SecurityGroupOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("SecurityGroupOperUpdate stream ended. closing..")
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from SecurityGroupOperUpdate stream. Err: %v", err)
			return err
		}

		err = eh.updateSecurityGroupOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating SecurityGroup oper state. Err: %v", err)
		}
	}
}
