package vcprobe

import (
	"errors"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

// PenDVSPortSettings represents a group of DVS port settings
// The key of this map represents the key of port on a DVS
// The value of this map represents the new setting of this port
type PenDVSPortSettings map[string]*types.VMwareDVSPortSetting

// AddPenDVS adds a new PenDVS to the given vcprobe instance
func (v *VCProbe) AddPenDVS(dcName string, dvsCreateSpec *types.DVSCreateSpec) error {
	dvsName := dvsCreateSpec.ConfigSpec.GetDVSConfigSpec().Name
	_, finder, _ := v.GetClientWithRLock()
	defer v.ReleaseClientRLock()

	// Check if it exists first
	if _, err := v.getPenDVS(dcName, dvsName, finder); err == nil {
		// PenDVS exists
		// TODO: check error isn't intermittent
		return nil
	}

	dc, err := finder.Datacenter(v.ClientCtx, dcName)
	if err != nil {
		v.Log.Errorf("Datacenter: %s doesn't exist, err: %s", dcName, err)
		return err
	}

	finder.SetDatacenter(dc)

	folders, err := dc.Folders(v.ClientCtx)
	if err != nil {
		return err
	}

	task, err := folders.NetworkFolder.CreateDVS(v.ClientCtx, *dvsCreateSpec)
	if err != nil {
		v.Log.Errorf("Failed at creating dvs: %s, err: %s", dvsName, err)
		return err
	}

	_, err = task.WaitForResult(v.ClientCtx, nil)
	if err != nil {
		v.Log.Errorf("Failed at waiting results of creating dvs: %s, err: %s", dvsName, err)
		return err
	}

	return nil
}

// GetPenDVS returns the DVS if it exists or an error
func (v *VCProbe) GetPenDVS(dcName, dvsName string) (*object.DistributedVirtualSwitch, error) {
	_, finder, _ := v.GetClientWithRLock()
	defer v.ReleaseClientRLock()
	return v.getPenDVS(dcName, dvsName, finder)
}

func (v *VCProbe) getPenDVS(dcName, dvsName string, finder *find.Finder) (*object.DistributedVirtualSwitch, error) {
	dc, err := finder.Datacenter(v.ClientCtx, dcName)
	if err != nil {
		v.Log.Errorf("Datacenter: %s doesn't exist, err: %s", dcName, err)
		return nil, err
	}

	finder.SetDatacenter(dc)

	net, err := finder.Network(v.ClientCtx, dvsName)
	if err != nil {
		v.Log.Errorf("Failed at finding network: %s, err: %s", dvsName, err)
		return nil, err
	}

	objDvs, ok := net.(*object.DistributedVirtualSwitch)
	if !ok {
		v.Log.Errorf("Failed at getting DVS object")
		return nil, errors.New("Failed at getting DVS object")
	}
	return objDvs, nil
}

// RemovePenDVS removes the DVS
func (v *VCProbe) RemovePenDVS(dcName, dvsName string) error {
	_, finder, _ := v.GetClientWithRLock()
	defer v.ReleaseClientRLock()

	objDvs, err := v.getPenDVS(dcName, dvsName, finder)
	if err != nil {
		return err
	}

	task, err := objDvs.Destroy(v.ClientCtx)
	if err != nil {
		v.Log.Errorf("Failed at destroying DVS: %s from datacenter: %s, err: %s", dvsName, dcName, err)
		return err
	}

	err = task.Wait(v.ClientCtx)
	if err != nil {
		v.Log.Errorf("Failed at destroying DVS: %s from datacenter: %s, err: %s", dvsName, dcName, err)
		return err
	}

	return nil
}

/*
// TODO: Need to re-enable it.
// This test works fine with real VC. However, in vcsim environment,
// after adding portgroups to DVS, pg.PortKeys[i] (pg is mo.DistributedVirtualPortgroup)
// is not initialized properly. I already filed a ticket to VMware for future assistant.
// Since this is unit testing and can't expect my real VC stand up forever,
// we temporarily skip it. Once we hear from VMware or we figure out why this happens
// by ourselves, we can unlock it.

// UpdatePorts updates port(s) on a given DVS based on the input mapping(portsSetting)
func (d *PenDVS) UpdatePorts(portsSetting *PenDVSPortSettings) ([]types.DistributedVirtualPort, error) {
	numPorts := len(*portsSetting)

	portKeys := make([]string, numPorts)
	portSpecs := make([]types.DVPortConfigSpec, numPorts)

	for k := range *portsSetting {
		portKeys = append(portKeys, k)
	}

	criteria := types.DistributedVirtualSwitchPortCriteria{
		PortKey: portKeys,
	}

	d.DvsMutex.Lock()
	defer d.DvsMutex.Unlock()

	ports, err := d.ObjDvs.FetchDVPorts(d.ctx, &criteria)
	if err != nil {
		d.log.Errorf("Can't find ports, err: %s", err)
		return nil, err
	}

	for i := 0; i < numPorts; i++ {
		portSpecs[i].ConfigVersion = ports[i].Config.ConfigVersion
		portSpecs[i].Key = ports[i].Key
		portSpecs[i].Operation = string("edit")
		portSpecs[i].Scope = ports[i].Config.Scope
		portSpecs[i].Setting = (*portsSetting)[ports[i].Key]
	}

	task, err := d.ObjDvs.ReconfigureDVPort(d.ctx, portSpecs)
	if err != nil {
		d.log.Errorf("Failed at reconfig DVS ports, err: %s", err)
		return ports, err
	}

	_, err = task.WaitForResult(d.ctx, nil)
	if err != nil {
		d.log.Errorf("Failed at modifying DVS ports, err: %s", err)
		return ports, err
	}

	return ports, nil
}
*/

// getMoDVSRef converts object.DistributedVirtualSwitch to mo.DistributedVirtualSwitch
// func (d *PenDVS) getMoDVSRef() (*mo.DistributedVirtualSwitch, error) {
// 	d.DvsMutex.Lock()
// 	defer d.DvsMutex.Unlock()

// 	var dvs mo.DistributedVirtualSwitch
// 	err := d.ObjDvs.Properties(d.ctx, d.ObjDvs.Reference(), nil, &dvs)
// 	if err != nil {
// 		d.log.Errorf("Failed at getting DVS properties, err: %s", err)
// 		return nil, err
// 	}

// 	return &dvs, nil
// }
