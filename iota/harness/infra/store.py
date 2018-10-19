#! /usr/bin/python3
import pdb
import copy

__gl_testbed = None
__gl_workloads = {}

def SetTestbed(tb):
    global __gl_testbed
    assert(tb)
    __gl_testbed = tb
    return

def GetTestbed():
    global __gl_testbed
    assert(__gl_testbed)
    return __gl_testbed

class Workload:
    def __init__(self, msg):
        self.workload_name = msg.workload_name
        self.node_name = msg.node_name
        self.encap_vlan = msg.encap_vlan
        self.ip_prefix = msg.ip_prefix
        self.ip_address = msg.ip_prefix.split('/')[0]
        self.mac_address = msg.mac_address
        self.interface = msg.interface
        self.interface_type = msg.interface_type
        self.pinned_port = msg.pinned_port
        self.uplink_vlan = msg.pinned_port
        return

def AddWorkloads(req):
    global __gl_workloads
    for wlmsg in req.workloads:
        wl = Workload(wlmsg)
        __gl_workloads[wl.workload_name] = wl
    return

def GetWorkloads():
    global __gl_workloads
    return __gl_workloads.values()
    
def DeleteWorkloads(req):
    global __gl_workloads
    for wl in req.workloads:
        del __gl_workloads[wl.workload_name]
    return

def GetLocalWorkloadPairs():
    pairs = []
    for w1 in GetWorkloads():
        for w2 in GetWorkloads():
            if id(w1) == id(w2): continue
            if w1.node_name == w2.node_name:
                pairs.append((w1, w2))
    return pairs

def GetRemoteWorkloadPairs():
    pairs = []
    for w1 in GetWorkloads():
        for w2 in GetWorkloads():
            if id(w1) == id(w2): continue
            if w1.node_name != w2.node_name:
                pairs.append((w1, w2))
    return pairs
