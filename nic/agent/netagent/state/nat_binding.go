// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package state

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/protos/netproto"
	"github.com/pensando/sw/nic/agent/netagent/state/types"
	"github.com/pensando/sw/venice/utils/log"
)

// CreateNatBinding creates a nat binding
func (na *Nagent) CreateNatBinding(nb *netproto.NatBinding) error {
	err := na.validateMeta(nb.Kind, nb.ObjectMeta)
	if err != nil {
		return err
	}
	oldNp, err := na.FindNatBinding(nb.ObjectMeta)
	if err == nil {
		// check if the contents are same
		if !proto.Equal(oldNp, nb) {
			log.Errorf("NatBinding %+v already exists", oldNp)
			return errors.New("nat binding already exists")
		}

		log.Infof("Received duplicate nat binding create {%+v}", nb)
		return nil
	}

	// find the corresponding namespace
	ns, err := na.FindNamespace(nb.Tenant, nb.Namespace)
	if err != nil {
		return err
	}

	// find the corresponding natpool
	np, err := na.findNatPool(nb.ObjectMeta, nb.Spec.NatPoolName)
	if err != nil {
		log.Infof("Could not find the specified NatPool. %v", nb.Spec.NatPoolName)
		return err
	}

	// find the corresponding natpool's namespace
	natPoolNS, err := na.FindNamespace(nb.Tenant, np.Namespace)
	if err != nil {
		log.Errorf("Could not find nat pool's namespace. NatPool : {%v}", np)
	}

	nb.Status.NatBindingID, err = na.Store.GetNextID(types.NatBindingID)

	if err != nil {
		log.Errorf("Could not allocate nat binding id. {%+v}", err)
		return err
	}

	// create it in datapath
	nb, err = na.Datapath.CreateNatBinding(nb, np, natPoolNS.Status.NamespaceID, ns)
	if err != nil {
		log.Errorf("Error creating nat binding in datapath. NatBinding {%+v}. Err: %v", nb, err)
		return err
	}

	// Add the current Nat Binding as a dependency to the Nat Pool.
	err = na.Solver.Add(np, nb)
	if err != nil {
		log.Errorf("Could not add dependency. Parent: %v. Child: %v", np, nb)
		return err
	}

	// Since Nat Binding can refer to a nat pool outside of its own namespace. We need to add Namespace <- NatBinding dependency
	// Add the current Nat Binding to its namespace only if it is not the same ns as the nat pool it is referring to.
	if ns.Status.NamespaceID != natPoolNS.Status.NamespaceID {
		err = na.Solver.Add(ns, nb)
		if err != nil {
			log.Errorf("Could not add dependency. Parent: %v. Child: %v", ns, nb)
			return err
		}
	}

	// save it in db
	key := na.Solver.ObjectKey(nb.ObjectMeta, nb.TypeMeta)
	na.Lock()
	na.NatBindingDB[key] = nb
	na.Unlock()
	err = na.Store.Write(nb)

	return err
}

// FindNatBinding finds a nat binding in local db
func (na *Nagent) FindNatBinding(meta api.ObjectMeta) (*netproto.NatBinding, error) {
	typeMeta := api.TypeMeta{
		Kind: "NatBinding",
	}
	// lock the db
	na.Lock()
	defer na.Unlock()

	// lookup the database
	key := na.Solver.ObjectKey(meta, typeMeta)
	nb, ok := na.NatBindingDB[key]
	if !ok {
		return nil, fmt.Errorf("nat binding not found %v", nb)
	}

	return nb, nil
}

// ListNatBinding returns the list of nat bindings
func (na *Nagent) ListNatBinding() []*netproto.NatBinding {
	var natBindingList []*netproto.NatBinding
	// lock the db
	na.Lock()
	defer na.Unlock()

	for _, nb := range na.NatBindingDB {
		natBindingList = append(natBindingList, nb)
	}

	return natBindingList
}

// UpdateNatBinding updates a nat binding
func (na *Nagent) UpdateNatBinding(nb *netproto.NatBinding) error {
	// find the corresponding namespace
	ns, err := na.FindNamespace(nb.Tenant, nb.Namespace)
	if err != nil {
		return err
	}
	existingNp, err := na.FindNatBinding(nb.ObjectMeta)
	if err != nil {
		log.Errorf("NatBinding %v not found", nb.ObjectMeta)
		return err
	}

	if proto.Equal(nb, existingNp) {
		log.Infof("Nothing to update.")
		return nil
	}

	err = na.Datapath.UpdateNatBinding(nb, ns)
	key := na.Solver.ObjectKey(nb.ObjectMeta, nb.TypeMeta)
	na.Lock()
	na.NatBindingDB[key] = nb
	na.Unlock()
	err = na.Store.Write(nb)
	return err
}

// DeleteNatBinding deletes a nat binding
func (na *Nagent) DeleteNatBinding(tn, namespace, name string) error {
	nb := &netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    tn,
			Namespace: namespace,
			Name:      name,
		},
	}
	err := na.validateMeta(nb.Kind, nb.ObjectMeta)
	if err != nil {
		return err
	}
	// find the corresponding namespace
	ns, err := na.FindNamespace(nb.Tenant, nb.Namespace)
	if err != nil {
		return err
	}

	existingNatBinding, err := na.FindNatBinding(nb.ObjectMeta)
	if err != nil {
		log.Errorf("NatBinding %+v not found", nb.ObjectMeta)
		return errors.New("nat binding not found")
	}

	// find the corresponding natpool
	np, err := na.findNatPool(nb.ObjectMeta, existingNatBinding.Spec.NatPoolName)
	if err != nil {
		log.Infof("Could not find the specified NatPool. %v", existingNatBinding.Spec.NatPoolName)
		return err
	}

	// find the corresponding natpool's namespace
	natPoolNS, err := na.FindNamespace(nb.Tenant, np.Namespace)
	if err != nil {
		log.Errorf("Could not find nat pool's namespace. NatPool : {%v}", np)
	}

	// check if the current nat binding has any objects referring to it
	err = na.Solver.Solve(existingNatBinding)
	if err != nil {
		log.Errorf("Found active references to %v. Err: %v", existingNatBinding.Name, err)
		return err
	}

	// delete it in the datapath
	err = na.Datapath.DeleteNatBinding(existingNatBinding, ns)
	if err != nil {
		log.Errorf("Error deleting nat binding {%+v}. Err: %v", nb, err)
	}

	// remove the reference to the nat pool
	err = na.Solver.Remove(np, existingNatBinding)
	if err != nil {
		log.Errorf("Could not remove the reference to the nat pool: %v. Err: %v", np.Name, err)
		return err
	}

	// remove the reference to its namespace. If we have added it.
	if ns.Status.NamespaceID != natPoolNS.Status.NamespaceID {
		err = na.Solver.Remove(ns, nb)
		if err != nil {
			log.Errorf("Could not remove the reference to the namespace: %v. Err: %v", ns.Name, nb)
			return err
		}
	}

	// delete from db
	key := na.Solver.ObjectKey(nb.ObjectMeta, nb.TypeMeta)
	na.Lock()
	delete(na.NatBindingDB, key)
	na.Unlock()
	err = na.Store.Delete(nb)

	return err
}

// findNatPool finds the associated NatPool with the NatBinding.
// The binding can refer to the nat pool outside its own namespace.
// In such cases, we expect the natpool name to be written as <remote namespace>/<natpoolname>
// These are expected to be tenant scoped
func (na *Nagent) findNatPool(natBindingMeta api.ObjectMeta, natPool string) (*netproto.NatPool, error) {
	var poolMeta api.ObjectMeta

	np := strings.Split(natPool, "/")
	switch len(np) {
	// NatPool in local namespace
	case 1:
		poolMeta.Tenant = natBindingMeta.Tenant
		poolMeta.Namespace = natBindingMeta.Namespace
		poolMeta.Name = natPool
		return na.FindNatPool(poolMeta)
		// NatPool in remote namespace
	case 2:
		poolMeta.Tenant = natBindingMeta.Tenant
		poolMeta.Namespace = np[0]
		poolMeta.Name = np[1]
		return na.FindNatPool(poolMeta)
	default:
		return nil, fmt.Errorf("nat pool {%v} not found", natPool)
	}
}
