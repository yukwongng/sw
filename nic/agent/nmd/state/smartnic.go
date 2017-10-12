// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package state

import (
	"github.com/pensando/sw/api/generated/cmd"
	"github.com/pensando/sw/venice/cmd/grpc"
	"github.com/pensando/sw/venice/utils/log"
)

// RegisterSmartNICReq registers a NIC with CMD
func (n *NMD) RegisterSmartNICReq(nic *cmd.SmartNIC) (grpc.RegisterNICResponse, error) {

	resp, err := n.cmd.RegisterSmartNICReq(nic)
	if err != nil {
		log.Errorf("Failed to register NIC, mac: %s err: %v", nic.ObjectMeta.Name, err)
		return resp, err
	}

	log.Infof("Register NIC response mac: %s response: %v", nic.ObjectMeta.Name, resp)
	return resp, nil
}

// UpdateSmartNICReq registers a NIC with CMD/Venice cluster
func (n *NMD) UpdateSmartNICReq(nic *cmd.SmartNIC) (*cmd.SmartNIC, error) {

	nic, err := n.cmd.UpdateSmartNICReq(nic)
	if err != nil {
		log.Errorf("Failed to update NIC, mac: %s err: %v", nic.ObjectMeta.Name, err)
		return nil, err
	}

	log.Infof("Update NIC response mac: %s nic: %+v", nic.ObjectMeta.Name, nic)
	return nic, nil
}

// CreateSmartNIC creates a local smartNIC object
func (n *NMD) CreateSmartNIC(nic *cmd.SmartNIC) error {

	log.Infof("SmartNIC create, mac: %s", nic.ObjectMeta.Name)

	// add the nic to database
	n.SetSmartNIC(nic)
	err := n.store.Write(nic)

	return err
}

// UpdateSmartNIC updates the local smartNIC object
func (n *NMD) UpdateSmartNIC(nic *cmd.SmartNIC) error {

	log.Infof("SmartNIC update, mac: %s", nic.ObjectMeta.Name)

	// update nic in the DB
	n.SetSmartNIC(nic)
	err := n.store.Write(nic)

	return err
}

// DeleteSmartNIC deletes the local smartNIC object
func (n *NMD) DeleteSmartNIC(nic *cmd.SmartNIC) error {

	log.Infof("SmartNIC delete, mac: %s", nic.ObjectMeta.Name)

	// remove nic from DB
	n.SetSmartNIC(nil)
	err := n.store.Delete(nic)

	return err
}
