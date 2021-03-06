// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package statemgr

import (
	"context"
	"fmt"
	"strings"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/api/generated/ctkit"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/memdb/objReceiver"
	"github.com/pensando/sw/venice/utils/runtime"
)

// DistributedServiceCardState is a wrapper for smartNic object
type DistributedServiceCardState struct {
	DistributedServiceCard *ctkit.DistributedServiceCard `json:"-"` // smartNic object
	stateMgr               *Statemgr                     // pointer to state manager
	recvHandle             objReceiver.Receiver
	profileVersion         dscProfileVersion
	decommissioned         bool
	// ToRevisit: Should DSC be updated with correct profile, then we need
	//  below code
	//NodeVersion            cluster.DSCProfileVersion
}

// GetKey returns the key of DSCProfile
func (sns *DistributedServiceCardState) GetKey() string {
	return sns.DistributedServiceCard.GetKey()
}

// Write write the obect to api server
func (sns *DistributedServiceCardState) Write() error {
	var err error
	/* ToDo: Until we figure out  a way to write the status from NPM
	   without conflicting with CMD
	log.Infof("Push DSC card status")

	sns.DistributedServiceCard.Lock()
	defer sns.DistributedServiceCard.Unlock()
	prop := sns.DistributedServiceCard.Status.ProfileVersion
	newProp := &sns.NodeVersion
	if prop != newProp {
		sns.DistributedServiceCard.Status.ProfileVersion = newProp
		sns.DistributedServiceCard.Write()
		if err != nil {
			sns.DistributedServiceCard.Status.ProfileVersion = prop
		}
	}*/
	return err

}

func (sns *DistributedServiceCardState) isOrchestratorCompatible() bool {
	dscProfile := sns.DistributedServiceCard.Spec.DSCProfile
	dscProfileState, _ := sns.stateMgr.FindDSCProfile("default", dscProfile)
	if dscProfileState == nil {
		log.Errorf("Failed to get DSC profile [%v]", dscProfile)
		return false
	}

	if dscProfileState.DSCProfile.Spec.DeploymentTarget != cluster.DSCProfileSpec_VIRTUALIZED.String() || dscProfileState.DSCProfile.Spec.FeatureSet != cluster.DSCProfileSpec_FLOWAWARE_FIREWALL.String() {
		return false
	}

	return true
}

//GetDistributedServiceCardWatchOptions gets options
func (sm *Statemgr) GetDistributedServiceCardWatchOptions() *api.ListWatchOptions {
	opts := api.ListWatchOptions{}
	opts.FieldChangeSelector = []string{}
	return &opts
}

// DistributedServiceCardStateFromObj converts from memdb object to smartNic state
func DistributedServiceCardStateFromObj(obj runtime.Object) (*DistributedServiceCardState, error) {
	switch obj.(type) {
	case *ctkit.DistributedServiceCard:
		sobj := obj.(*ctkit.DistributedServiceCard)
		switch sobj.HandlerCtx.(type) {
		case *DistributedServiceCardState:
			nsobj := sobj.HandlerCtx.(*DistributedServiceCardState)
			return nsobj, nil
		default:
			return nil, ErrIncorrectObjectType
		}
	default:
		return nil, ErrIncorrectObjectType
	}
}

// NewDistributedServiceCardState creates new smartNic state object
func NewDistributedServiceCardState(smartNic *ctkit.DistributedServiceCard, stateMgr *Statemgr) (*DistributedServiceCardState, error) {
	recvHandle, err := stateMgr.mbus.AddReceiver(smartNic.Status.PrimaryMAC)
	if err != nil {
		return nil, fmt.Errorf("Error add dsc %v", err)
	}
	log.Infof("Added DSC %v as a receiver", smartNic.Status.PrimaryMAC)
	hs := &DistributedServiceCardState{
		DistributedServiceCard: smartNic,
		stateMgr:               stateMgr,
		recvHandle:             recvHandle,
	}
	smartNic.HandlerCtx = hs

	return hs, nil
}

// OnDistributedServiceCardCreate handles smartNic creation
func (sm *Statemgr) OnDistributedServiceCardCreate(smartNic *ctkit.DistributedServiceCard) error {
	defer sm.ProcessDSCEvent(CreateEvent, &smartNic.DistributedServiceCard)
	defer sm.sendDscUpdateNotification(&smartNic.DistributedServiceCard)
	sns, err := sm.dscCreate(smartNic)
	if err != nil {
		return err
	}
	sm.addDSCRelatedobjects(smartNic, sns, true)
	sm.PeriodicUpdaterPush(sns)
	return nil
}

func (sm *Statemgr) dscCreate(smartNic *ctkit.DistributedServiceCard) (*DistributedServiceCardState, error) {
	log.Infof("Creating smart nic: %+v", smartNic)

	// create new smartNic object
	sns, err := NewDistributedServiceCardState(smartNic, sm)
	if err != nil {
		log.Errorf("Error creating smartNic %+v. Err: %v", smartNic, err)
		return nil, err
	}

	log.Infof("Profile %s", smartNic.DistributedServiceCard.Spec.DSCProfile)
	profName := smartNic.DistributedServiceCard.Spec.DSCProfile
	if profName != "" {
		profileState, err := sm.FindDSCProfile("", profName)
		if err == nil {
			log.Infof("Found a profile send it to the DSC")
			profileState.DSCProfile.Lock()
			profileVersion := dscProfileVersion{
				profileState.DSCProfile.Name,
				profileState.DSCProfile.GenerationID,
			}
			profileState.DscList[smartNic.ObjectMeta.Name] = profileVersion
			sns.profileVersion = profileVersion
			//sns.NodeVersion = cluster.DSCProfileVersion{}
			ret := profileState.PushObj.AddObjReceivers([]objReceiver.Receiver{sns.recvHandle})
			if ret != nil {
				log.Infof("Add receiver failed %v", ret)
			} else {
				log.Infof("Added the dsc: %s to profile: %s", smartNic.Name, profName)
			}
			profileState.DSCProfile.Unlock()
		} else {
			//Retry as this might be timing
			return nil, kvstore.NewKeyNotFoundError(profName, 0)
		}
	}

	return sns, nil
}

func (sm *Statemgr) addDSCRelatedobjects(smartNic *ctkit.DistributedServiceCard, sns *DistributedServiceCardState, sendSgPolicies bool) {
	// see if smartnic is admitted
	/*if sm.isDscAdmitted(&smartNic.DistributedServiceCard) {
		if sendSgPolicies {
			// Update SGPolicies
			policies, _ := sm.ListSgpolicies()
			for _, policy := range policies {
				policy.NetworkSecurityPolicy.Lock()
				policy.processDSCUpdate(&smartNic.DistributedServiceCard)
				policy.NetworkSecurityPolicy.Unlock()
			}
		}
		// Update FirewallProfiles
		fwprofiles, _ := sm.ListFirewallProfiles()
		for _, fwprofile := range fwprofiles {
			fwprofile.FirewallProfile.Lock()
			fwprofile.processDSCUpdate(&smartNic.DistributedServiceCard)
			fwprofile.FirewallProfile.Unlock()
		}
	}*/

	hosts, err := sm.ctrler.Host().List(context.Background(), &api.ListWatchOptions{})
	if err != nil {
		log.Errorf("failed to get list of all hosts. Err : %v", err)
	}

	// walk all hosts and ee if they need to be associated to this snic
	for _, host := range hosts {
		associated := false
		for _, snid := range host.Spec.DSCs {
			if (snid.ID == smartNic.DistributedServiceCard.Spec.ID) || (snid.MACAddress == smartNic.DistributedServiceCard.Status.PrimaryMAC) {
				associated = true
			}
		}

		// if this host and smartnic are associated, trigger workload reconciliation
		if associated {
			host.Lock()
			hs, err := sm.FindHost(host.Tenant, host.Name)
			if err == nil {
				hs.workloads.Range(func(key, value interface{}) bool {
					wmeta := value.(api.ObjectMeta)
					wrk, err := sm.FindWorkload(wmeta.Tenant, wmeta.Name)
					if err == nil {
						sm.reconcileWorkload(wrk.Workload, hs, sns)
					} else {
						log.Errorf("Error finding workload. Err: %v", err)
					}
					return true
				})
			} else {
				log.Errorf("Error finding host %v. Err: %v", host.Name, err)
			}
			host.Unlock()
		}
	}
}

// OnDistributedServiceCardUpdate handles update event on smartnic
func (sm *Statemgr) OnDistributedServiceCardUpdate(smartNic *ctkit.DistributedServiceCard, nsnic *cluster.DistributedServiceCard) error {
	defer sm.sendDscUpdateNotification(nsnic)
	defer sm.ProcessDSCEvent(UpdateEvent, &smartNic.DistributedServiceCard)

	sns, err := DistributedServiceCardStateFromObj(smartNic)
	if err != nil {
		log.Errorf("Error finding smartnic. Err: %v", err)
		return err
	}

	if !sns.decommissioned {
		if nsnic.Status.AdmissionPhase == cluster.DistributedServiceCardStatus_DECOMMISSIONED.String() {
			sns.decommissioned = true
			sm.deleteDsc(smartNic)
		}
	} else {
		if nsnic.Status.AdmissionPhase == cluster.DistributedServiceCardStatus_ADMITTED.String() {
			sns.decommissioned = false
			_, err = sm.AddDsc(smartNic)
			log.Errorf("Error Adding smartnic. Err: %v", err)
			return err
		}
	}

	if !sns.decommissioned {
		sns, err = sm.updateDSC(smartNic, nsnic)
		if err != nil {
			return err
		}
	}

	sm.PeriodicUpdaterPush(sns)
	return nil
}

func (sm *Statemgr) updateDSC(smartNic *ctkit.DistributedServiceCard, nsnic *cluster.DistributedServiceCard) (*DistributedServiceCardState, error) {
	// see if we already have the smartNic
	log.Infof("Update of DistributedServiceCard")
	defer sm.sendDscUpdateNotification(nsnic)
	sns, err := DistributedServiceCardStateFromObj(smartNic)
	if err != nil {
		log.Errorf("Error finding smartnic. Err: %v", err)
		return nil, err
	}

	newProfile := nsnic.Spec.DSCProfile
	oldProfile := smartNic.DistributedServiceCard.Spec.DSCProfile
	log.Infof("Old:%s new: %s", oldProfile, newProfile)

	if newProfile != oldProfile {
		//Find the newProfile
		if newProfile != "" {
			profileState, err := sm.FindDSCProfile("", newProfile)
			if err == nil {
				profileState.DSCProfile.Lock()
				profileVersion := dscProfileVersion{
					profileState.DSCProfile.Name,
					profileState.DSCProfile.GenerationID,
				}
				profileState.DscList[smartNic.ObjectMeta.Name] = profileVersion
				sns.profileVersion = profileVersion
				//sns.NodeVersion = cluster.DSCProfileVersion{}
				ret := profileState.PushObj.AddObjReceivers([]objReceiver.Receiver{sns.recvHandle})
				if ret != nil {
					log.Infof("Add receiver failed %v", ret)
				} else {
					log.Infof("Added the dsc: %s to profile: %s", smartNic.Name, newProfile)
				}

				profileState.DSCProfile.Unlock()
			} else {
				//Retry as this might be timing
				return nil, kvstore.NewKeyNotFoundError(newProfile, 0)
			}
		}

		if oldProfile != "" {
			oldProfileState, _ := sm.FindDSCProfile("", oldProfile)
			oldProfileState.DSCProfile.Lock()
			if _, ok := oldProfileState.DscList[nsnic.ObjectMeta.Name]; ok {
				delete(oldProfileState.DscList, nsnic.ObjectMeta.Name)
				ret := oldProfileState.PushObj.RemoveObjReceivers([]objReceiver.Receiver{sns.recvHandle})
				if ret != nil {
					log.Infof("Remove receiver failed %v", ret)
				} else {
					log.Infof("removed the dsc: %s to profile: %s", smartNic.Name, oldProfile)
				}

			}
			oldProfileState.DSCProfile.Unlock()

		}
	} else {
		log.Infof("No change in profile")
	}

	sns.DistributedServiceCard.DistributedServiceCard = *nsnic

	return sns, nil
}

// OnDistributedServiceCardDelete handles smartNic deletion
func (sm *Statemgr) OnDistributedServiceCardDelete(smartNic *ctkit.DistributedServiceCard) error {
	defer sm.ProcessDSCEvent(DeleteEvent, &smartNic.DistributedServiceCard)
	defer sm.sendDscDeleteNotification(&smartNic.DistributedServiceCard)
	hs, err := sm.deleteDsc(smartNic)
	if err != nil {
		return err
	}

	return sm.deleteDscRelatedObjects(smartNic, hs, true)
}

func (sm *Statemgr) deleteDsc(smartNic *ctkit.DistributedServiceCard) (*DistributedServiceCardState, error) {
	// see if we have the smartNic
	hs, err := DistributedServiceCardStateFromObj(smartNic)
	if err != nil {
		log.Errorf("Could not find the smartNic %v. Err: %v", smartNic, err)
		return nil, err
	}

	log.Infof("Deleting smart nic: %+v", smartNic)
	defer func() {
		err := sm.mbus.DeleteReceiver(hs.recvHandle)
		if err != nil {
			log.Errorf("Error deleting receiver %v", err.Error())
		}
		log.Infof("Removed DSC %v as a receiver", hs.DistributedServiceCard.Status.PrimaryMAC)
	}()

	oldProfile := smartNic.Spec.DSCProfile
	if oldProfile != "" {
		oldProfileState, _ := sm.FindDSCProfile("", oldProfile)
		oldProfileState.DSCProfile.Lock()
		if _, ok := oldProfileState.DscList[smartNic.ObjectMeta.Name]; ok {
			delete(oldProfileState.DscList, smartNic.ObjectMeta.Name)

			oldProfileState.PushObj.RemoveObjReceivers([]objReceiver.Receiver{hs.recvHandle})
		}
		oldProfileState.DSCProfile.Unlock()

	}
	return hs, nil
}

//AddDsc add DSC as a receiver
func (sm *Statemgr) AddDsc(smartNic *ctkit.DistributedServiceCard) (*DistributedServiceCardState, error) {
	// see if we have the smartNic
	hs, err := DistributedServiceCardStateFromObj(smartNic)
	if err != nil {
		log.Errorf("Could not find the smartNic %v. Err: %v", smartNic, err)
		return nil, err
	}

	log.Infof("Add smart nic: %+v", smartNic)
	hs.recvHandle, err = sm.mbus.AddReceiver(smartNic.Status.PrimaryMAC)
	if err != nil {
		log.Errorf("Error adding receiver %v", err.Error())
		return hs, err
	}
	log.Infof("Add DSC %v as a receiver", hs.DistributedServiceCard.Status.PrimaryMAC)

	return hs, nil
}

func (sm *Statemgr) deleteDscRelatedObjects(smartNic *ctkit.DistributedServiceCard, hs *DistributedServiceCardState, sgPolicyDelete bool) error {

	hosts, err := sm.ctrler.Host().List(context.Background(), &api.ListWatchOptions{})
	if err != nil {
		log.Errorf("Failed to get list of all hosts. Err : %v", err)
		return err
	}

	// walk all hosts and see if they need to be dis-associated to this snic
	for _, host := range hosts {
		associated := false
		for _, snid := range host.Spec.DSCs {
			if (snid.ID == smartNic.DistributedServiceCard.Spec.ID) || (snid.MACAddress == smartNic.DistributedServiceCard.Status.PrimaryMAC) {
				associated = true
			}
		}

		// if this host and smartnic are associated, trigger workload reconciliation
		if associated {
			hs, err := sm.FindHost(host.Tenant, host.Name)
			if err == nil {
				hs.workloads.Range(func(key, value interface{}) bool {
					wmeta := value.(api.ObjectMeta)
					wrk, err := sm.FindWorkload(wmeta.Tenant, wmeta.Name)
					if err == nil {
						sm.reconcileWorkload(wrk.Workload, hs, nil)
					} else {
						log.Errorf("Error finding workload. Err: %v", err)
					}
					return true
				})
			} else {
				log.Errorf("Error finding host %v. Err: %v", host.Name, err)
				return err
			}
		}
	}
	return nil
}

// OnDistributedServiceCardReconnect is called when ctkit reconnects to apiserver
func (sm *Statemgr) OnDistributedServiceCardReconnect() {
	return
}

// FindDistributedServiceCard finds a smartNic
func (sm *Statemgr) FindDistributedServiceCard(tenant, name string) (*DistributedServiceCardState, error) {
	// find the object
	obj, err := sm.FindObject("DistributedServiceCard", "", "", name)
	if err != nil {
		return nil, err
	}

	return DistributedServiceCardStateFromObj(obj)
}

// FindDistributedServiceCardByMacAddr finds the smart nic by mac addr
func (sm *Statemgr) FindDistributedServiceCardByMacAddr(macAddr string) (*DistributedServiceCardState, error) {
	objs := sm.ListObjects("DistributedServiceCard")

	for _, obj := range objs {
		snic, err := DistributedServiceCardStateFromObj(obj)
		if err != nil {
			return nil, err
		}

		if snic.DistributedServiceCard.Status.PrimaryMAC == macAddr {
			return snic, nil
		}
	}

	return nil, fmt.Errorf("Smartnic not found for mac addr %v", macAddr)
}

// FindDistributedServiceCardByHname finds smart nic by used given name
func (sm *Statemgr) FindDistributedServiceCardByHname(hname string) (*DistributedServiceCardState, error) {
	objs := sm.ListObjects("DistributedServiceCard")

	for _, obj := range objs {
		snic, err := DistributedServiceCardStateFromObj(obj)
		if err != nil {
			return nil, err
		}

		if snic.DistributedServiceCard.Spec.ID == hname {
			return snic, nil
		}
	}

	//Update the DSC object with the profile, genID

	return nil, fmt.Errorf("Smartnic not found for name %v", hname)
}

// ListDistributedServiceCards lists all smart nics
func (sm *Statemgr) ListDistributedServiceCards() ([]*DistributedServiceCardState, error) {
	objs := sm.ListObjects("DistributedServiceCard")

	var dscs []*DistributedServiceCardState
	for _, obj := range objs {
		dsc, err := DistributedServiceCardStateFromObj(obj)
		if err != nil {
			return dscs, err
		}

		dscs = append(dscs, dsc)
	}

	return dscs, nil
}

// isDscHealthy returns true if smartnic is in healthry condition
func (sm *Statemgr) isDscHealthy(nsnic *cluster.DistributedServiceCard) bool {
	isHealthy := false
	log.Infof("DSC STATSUS %v", nsnic.Status.Conditions)
	if len(nsnic.Status.Conditions) > 0 {
		for _, cond := range nsnic.Status.Conditions {
			if cond.Type == cluster.DSCCondition_HEALTHY.String() && cond.Status == cluster.ConditionStatus_TRUE.String() {
				isHealthy = true
			}
		}
	}

	return isHealthy
}

// isDscAdmitted returns true if the DSC is admited into the cluster
func (sm *Statemgr) isDscAdmitted(nsnic *cluster.DistributedServiceCard) bool {
	return nsnic.Status.AdmissionPhase == cluster.DistributedServiceCardStatus_ADMITTED.String()
}

// isDscEnforcednMode returns true if the DSC in insertion mode cluster
func (sm *Statemgr) isDscEnforcednMode(nsnic *cluster.DistributedServiceCard) bool {

	if !sm.isDscAdmitted(nsnic) {
		return false
	}

	profileState, err := sm.FindDSCProfile("", nsnic.Spec.DSCProfile)
	if err != nil {
		return false
	}

	return (strings.ToLower(profileState.DSCProfile.DSCProfile.Spec.DeploymentTarget) == strings.ToLower(cluster.DSCProfileSpec_VIRTUALIZED.String()) && strings.ToLower(profileState.DSCProfile.DSCProfile.Spec.FeatureSet) == strings.ToLower(cluster.DSCProfileSpec_FLOWAWARE_FIREWALL.String()))
}

// isDscEnforcednMode returns true if the DSC in insertion mode cluster
func (sm *Statemgr) isDscFlowawareMode(nsnic *cluster.DistributedServiceCard) bool {

	if !sm.isDscAdmitted(nsnic) {
		return false
	}

	profileState, err := sm.FindDSCProfile("", nsnic.Spec.DSCProfile)
	if err != nil {
		return false
	}

	return (strings.ToLower(profileState.DSCProfile.DSCProfile.Spec.DeploymentTarget) == strings.ToLower(cluster.DSCProfileSpec_HOST.String()) && strings.ToLower(profileState.DSCProfile.DSCProfile.Spec.FeatureSet) == strings.ToLower(cluster.DSCProfileSpec_FLOWAWARE.String()))

}
