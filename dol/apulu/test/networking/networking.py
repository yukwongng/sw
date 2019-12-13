#! /usr/bin/python3
# Networking Module
import pdb

import apollo.config.utils as utils
import apollo.config.topo as topo
import apollo.config.generator as generator
import apollo.test.callbacks.common.modcbs as modcbs

def Setup(infra, module):
    if 'WORKFLOW_START' in module.name:
        topo.CachedObjs.select_objs = True
        topo.CachedObjs.setMaxLimits(module.testspec.selectors.maxlimits)
    modcbs.Setup(infra, module)
    return True

def TestCaseSetup(tc):
    tc.AddIgnorePacketField('UDP', 'sport')
    tc.AddIgnorePacketField('UDP', 'chksum')
    tc.AddIgnorePacketField('IP', 'chksum') #Needed to pass NAT testcase
    return True

def TestCaseTeardown(tc):
    return True

def TestCasePreTrigger(tc):
    return True

def TestCaseStepSetup(tc, step):
    return True

def TestCaseStepTrigger(tc, step):
    return True

def TestCaseStepVerify(tc, step):
    return True

def TestCaseStepTeardown(tc, step):
    return True

def TestCaseVerify(tc):
    if tc.module.name.find("WORKFLOW"):
        generator.__read();
    if 'WORKFLOW_START' in tc.module.name:
        topo.CachedObjs.select_objs = False
        topo.CachedObjs.use_selected_objs = True
    elif 'WORKFLOW_END' in tc.module.name:
        topo.CachedObjs.reset()
    return True
