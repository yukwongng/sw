package vchub

import (
	"sync"

	"github.com/vmware/govmomi/vim25/types"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/defs"
	"github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/useg"
	"github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/vcprobe"
)

// PenDC represents an instance of a Datacenter
type PenDC struct {
	sync.Mutex
	*defs.State
	Name string
	// Map from dvs display name to PenDVS inside this DC
	DvsMap map[string]*PenDVS
	probe  *vcprobe.VCProbe
}

// NewPenDC creates a new penDC object
func (v *VCHub) NewPenDC(dcName string) (*PenDC, error) {
	v.DcMapLock.Lock()
	defer v.DcMapLock.Unlock()

	var dc *PenDC
	if dcExisting, ok := v.DcMap[dcName]; ok {
		dc = dcExisting
	} else {
		dc = &PenDC{
			State:  v.State,
			probe:  v.probe,
			Name:   dcName,
			DvsMap: map[string]*PenDVS{},
		}
		v.DcMap[dcName] = dc
	}

	dvsName := defs.DefaultDVSName

	// Create a pen dvs
	// Pvlan allocations on the dvs
	pvlanConfigSpecArray := []types.VMwareDVSPvlanConfigSpec{}

	// Setup all the pvlan allocations now
	for i := 1; i < useg.ReservedPGVlanCount; i += 2 {
		PvlanEntryProm := types.VMwareDVSPvlanMapEntry{
			PrimaryVlanId:   int32(i),
			PvlanType:       "promiscuous",
			SecondaryVlanId: int32(i),
		}
		pvlanMapEntry := types.VMwareDVSPvlanMapEntry{
			PrimaryVlanId:   int32(i),
			PvlanType:       "isolated",
			SecondaryVlanId: int32(i + 1),
		}
		pvlanSpecProm := types.VMwareDVSPvlanConfigSpec{
			PvlanEntry: PvlanEntryProm,
			Operation:  "add",
		}
		pvlanSpec := types.VMwareDVSPvlanConfigSpec{
			PvlanEntry: pvlanMapEntry,
			Operation:  "add",
		}
		pvlanConfigSpecArray = append(pvlanConfigSpecArray, pvlanSpecProm)
		pvlanConfigSpecArray = append(pvlanConfigSpecArray, pvlanSpec)
	}

	// TODO: Set number of uplinks
	var spec types.DVSCreateSpec
	spec.ConfigSpec = &types.VMwareDVSConfigSpec{
		PvlanConfigSpec: pvlanConfigSpecArray,
	}
	spec.ConfigSpec.GetDVSConfigSpec().Name = dvsName
	err := dc.AddPenDVS(&spec)
	if err != nil {
		v.Log.Errorf("Failed to create DVS %s in DC %s, err %s", dvsName, dcName, err)
		return nil, err
	}

	return dc, nil
}

// GetDC returns the DC by display name
func (v *VCHub) GetDC(name string) *PenDC {
	v.DcMapLock.Lock()
	defer v.DcMapLock.Unlock()
	dc, ok := v.DcMap[name]
	if !ok {
		return nil
	}
	return dc
}

// AddPG adds a PG to all DVS in this DC, unless dvsName is not blank
func (d *PenDC) AddPG(pgName string, networkMeta api.ObjectMeta, dvsName string) []error {
	d.Lock()
	defer d.Unlock()
	var errs []error
	for _, penDVS := range d.DvsMap {
		if dvsName == "" || dvsName == penDVS.DvsName {
			err := penDVS.AddPenPG(pgName, networkMeta)
			if err != nil {
				d.Log.Errorf("Create PG for dvs %s returned err %s", penDVS.DvsName, err)
				errs = append(errs, err)
			}
		}
	}
	return errs
}

// GetPG returns the pg with the matching name. Looks thorugh all
// DVS unless dvsName is supplied
func (d *PenDC) GetPG(pgName string, dvsName string) *PenPG {
	d.Lock()
	defer d.Unlock()
	for _, penDVS := range d.DvsMap {
		if dvsName == "" || dvsName == penDVS.DvsName {
			pg := penDVS.GetPenPG(pgName)
			if pg != nil {
				return pg
			}
		}
	}
	return nil
}

// RemovePG removes a PG from all DVS in this DC, unless dvsName is not blank
func (d *PenDC) RemovePG(pgName string, dvsName string) []error {
	d.Lock()
	defer d.Unlock()
	var errs []error
	for _, penDVS := range d.DvsMap {
		if dvsName == "" || dvsName == penDVS.DvsName {
			err := penDVS.RemovePenPG(pgName)
			if err != nil {
				d.Log.Errorf("Delete PG for dvs %s returned err %s", penDVS.DvsName, err)
				errs = append(errs, err)
				continue
			}
			penDVS.UsegMgr.ReleaseVlansForPG(pgName)
		}
	}
	return errs
}