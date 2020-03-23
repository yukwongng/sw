#! /usr/bin/python3
import iota.harness.api as api
import iota.test.apulu.config.api as config_api
import iota.test.apulu.utils.pdsctl as pdsctl
import iota.test.apulu.utils.config.utils as conf_utils
import iota.test.apulu.utils.misc as misc_utils
import time

def __learn_endpoints(args):
    # Enabling Max age for all endpoints
    if not conf_utils.SetDeviceLearnTimeout(86400):
        return api.types.status.FAILURE

    if not conf_utils.ClearLearnData():
        return api.types.status.FAILURE

    if not api.IsSimulation():
        req = api.Trigger_CreateAllParallelCommandsRequest()
    else:
        req = api.Trigger_CreateExecuteCommandsRequest(serial = False)
    
    for wl in api.GetWorkloads():
        api.Trigger_AddCommand(req, wl.node_name, wl.workload_name,
                               "arping -c  5 -U %s -I %s" % (wl.ip_address, wl.interface))
        api.Logger.debug("ArPing from %s %s %s %s" % (wl.node_name, wl.workload_name, wl.interface, wl.ip_address))
        # send arping with secondary IPs too
        for sec_ip_addr in wl.sec_ip_addresses:
            api.Trigger_AddCommand(req, wl.node_name, wl.workload_name,
                                   "arping -c  5 -U %s -I %s" % (sec_ip_addr, wl.interface))
            api.Logger.debug("ArPing from %s %s %s %s" % (wl.node_name, wl.workload_name, wl.interface, sec_ip_addr))

    resp = api.Trigger(req)
    if resp is None:
        return api.types.status.FAILURE

    return api.types.status.SUCCESS

def __verify_learn():
    api.Logger.verbose("Verifying if all VNIC and Mappings are learnt")
    #TODO to invoke pdsctl show learn verification once it is done
    nodes = api.GetNaplesHostnames()
    for node in nodes:
        ret, resp = pdsctl.ExecutePdsctlShowCommand(node, "learn mac", yaml=False)
        if ret != True:
            api.Logger.error("Failed to execute show learn mac at node %s : %s" %(node, resp))
            return api.types.status.FAILURE

        ret, resp = pdsctl.ExecutePdsctlShowCommand(node, "learn ip", yaml=False)
        if ret != True:
            api.Logger.error("Failed to execute show learn ip at node %s : %s" %(node, resp))
            return api.types.status.FAILURE

        ret, resp = pdsctl.ExecutePdsctlShowCommand(node, "learn statistics", yaml=False)
        if ret != True:
            api.Logger.error("Failed to execute show learn statistics at node %s : %s" %(node, resp))
            return api.types.status.FAILURE

    return api.types.status.SUCCESS

def Main(args):
    api.Logger.verbose("Learn VNIC and Mappings")
    result = __learn_endpoints(args)
    if result != api.types.status.SUCCESS:
        return result
    result = __verify_learn()
    # sleep for some time to let metaswitch advertise these local mappings to other naples.
    # TODO: have to find out if there is any event to wait on
    api.Logger.debug("Sleeping for sometime letting remote mappings to get programmed")
    misc_utils.Sleep(40)
    return result

if __name__ == '__main__':
    Main(None)
