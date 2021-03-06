package factory

import (
	"errors"

	cfgModel "github.com/pensando/sw/iota/test/venice/iotakit/cfg/enterprise"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/base"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/basenet"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/cloud"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/common"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/enterprise"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/vcenter"
	"github.com/pensando/sw/iota/test/venice/iotakit/testbed"
)

// NewDefaultSysModel creates a sysmodel for a testbed
func NewDefaultSysModel(tb *testbed.TestBed, cfgType cfgModel.CfgType, skipSetup bool) (*enterprise.SysModel, error) {

	sm := &enterprise.SysModel{SysModel: base.SysModel{Type: common.DefaultModel}}

	if err := sm.Init(tb, cfgType, skipSetup); err != nil {
		sm.CollectLogs()
		return nil, err
	}

	return sm, nil
}

// NewVcenterSysModel creates a sysmodel for a testbed
func NewVcenterSysModel(tb *testbed.TestBed, cfgType cfgModel.CfgType, skipSetup bool) (*vcenter.VcenterSysModel, error) {

	vsm := &vcenter.VcenterSysModel{SysModel: enterprise.SysModel{SysModel: base.SysModel{Type: common.VcenterModel}}}

	if err := vsm.Init(tb, cfgType, skipSetup); err != nil {
		vsm.CollectLogs()
		return nil, err
	}

	return vsm, nil
}

// NewCloudSysModel creates a sysmodel for a testbed
func NewCloudSysModel(tb *testbed.TestBed, cfgType cfgModel.CfgType, skipSetup bool) (*cloud.SysModel, error) {

	vsm := &cloud.SysModel{SysModel: base.SysModel{Type: common.CloudModel}}

	if err := vsm.Init(tb, cfgType, skipSetup); err != nil {
		vsm.CollectLogs()
		return nil, errors.New("could not initialize config objects")

	}

	return vsm, nil

}

// NewBasenetSysModel creates a sysmodel for a testbed
func NewBasenetSysModel(tb *testbed.TestBed, cfgType cfgModel.CfgType, skipSetup bool) (*basenet.SysModel, error) {

	vsm := &basenet.SysModel{SysModel: base.SysModel{Type: common.BaseNetModel}}

	if err := vsm.Init(tb, cfgType, skipSetup); err != nil {
		vsm.CollectLogs()
		return nil, err
	}

	return vsm, nil
}
