#! /usr/bin/python3
import iota.harness.api as api
import iota.protos.pygen.topo_svc_pb2 as topo_svc_pb2
import iota.test.iris.config.api as cfg_api
import iota.test.iris.config.netagent.api as agent_api
import iota.test.iris.testcases.telemetry.utils as utils
import iota.test.iris.testcases.expanded_telemetry.utils as eutils
from iota.test.iris.testcases.aging.aging_utils import *
import pdb
import json
import ipaddress
from scapy.utils import *
from scapy.utils import rdpcap
from scapy.utils import wrpcap
from scapy import packet
from scapy.all import Ether

def Setup(tc):
    #
    # Set up global variables
    #
    tc.feature = 'lif-erspan'
    tc.protocol = 'all'
    tc.udp_count = 1
    tc.icmp_count = 1

    #
    # Establish Workloads
    #
    tc.classic_mode = None
    tc.error = False
    result = eutils.establishWorkloads(tc)
    if result != api.types.status.SUCCESS:
        tc.error = True
        return result

    #
    # Check to see if Naples-hosted node is present
    #
    result = eutils.establishNaplesWorkload(tc)
    if result != api.types.status.SUCCESS:
        tc.error = True
        return result

    tc.classic_mode = False
    if api.IsBareMetalWorkloadType(tc.naples.node_name):
        tc.classic_mode = True

    #
    # Figure out (Collector) Workloads that are remote to Naples node and
    # identify remote Naples-peer, if applicable
    #
    result = eutils.establishRemoteWorkloads(tc)
    if result != api.types.status.SUCCESS:
        tc.error = True
        return result

    #
    # Figure out Workloads that are local to Naples node and identify
    # local Naples-peer, if applicable
    #
    result = eutils.establishLocalWorkloads(tc)
    if result != api.types.status.SUCCESS:
        tc.error = True
        return result

    if tc.classic_mode == True:
        result = eutils.establishCollectorWorkloadInClassicMode(tc, 
                                                                '10.255.0.2')
        if result != api.types.status.SUCCESS:
            tc.error = True
            return result

    #
    # Allocate Secondary-IPs for collectors as needed in order to test up to
    # Eight collectors
    #
    result = eutils.establishCollectorSecondaryIPs(tc)
    if result != api.types.status.SUCCESS:
        tc.error = True
        return result

    eutils.debugWorkLoadTraces(tc)

    #
    # Retrieve relevant Table-Id's
    #
    eutils.retrieveTableIds(tc)

    #
    # Ignore non-applicable test-options in Sanity mode
    # - [vlan = native] is not supported until P4 issue on
    #   reciepient Naples-node is fixed for ERSPAN on
    #   untagged pkts
    # - [collector = local] is not supported until code has
    #   support to enable the same
    # - Multi-collector testing is not enabled in freebsd
    #   environment until tcpdump capture on secondary-IP is resolved
    #
    tc.ignore = False
    if tc.args.type == 'sanity':
        for c in range(0, len(tc.collector)):
            if tc.iterators.collector == 'local' or\
               (api.GetNodeOs(tc.naples.node_name) == 'freebsd' and\
                tc.iterators.ccount > 1) or\
               (tc.collector[c].IsNaples() and tc.iterators.vlan == 'native'):
                tc.ignore = True
                return api.types.status.SUCCESS

    #
    # Set up runtime validation knobs
    # - dupcheck:   To circumvent duplicate-pkts checks in case of LIF-ERSPAN
    #               where multiple workloads map to the same LIF
    #
    tc.dupcheck = 'enable'
    if tc.iterators.peer == 'local' or tc.iterators.iftype != 'uplink':
        tc.dupcheck = 'disable'

    #
    # Preserve current TCP-Close configs and shrink it to 1-second
    #
    if tc.classic_mode == False:
        tc.tcp_close_val = get_timeout_val('tcp-close')
        update_timeout('tcp-close', "1s")

    return api.types.status.SUCCESS

def Trigger(tc):
    if tc.ignore == True:
        return api.types.status.SUCCESS

    if tc.error == True:
        return api.types.status.FAILURE

    tc.resp_tcpdump = None
    tc.resp_cleanup = None

    protoDir = api.GetTopologyDirectory() +\
               "/gen/telemetry/{}/{}".format('mirror', 'lif')
    api.Logger.error("Template Config files location: ", protoDir)

    policies = utils.GetTargetJsons('mirror', 'lif')
    for policy_json in policies:
        #
        # Get template-Mirror Config
        #
        newObjects = agent_api.AddOneConfig(policy_json)
        if len (newObjects) == 0:
            api.Logger.error("Adding new objects to store failed")
            tc.error = True
            return api.types.status.FAILURE

        #
        # Modify template-Mirror Config to make sure that Naples-node
        # act as either source or destination
        #
        # Set up Collector in the remote node
        #
        if newObjects[0].kind == 'Collector':
            tc.collector_objects = newObjects
            agent_api.RemoveConfigObjects(tc.collector_objects)
        elif newObjects[0].kind == 'Interface':
            tc.interface_objects = newObjects
            agent_api.RemoveConfigObjects(tc.interface_objects)

    for i in range(0, len(policies)):
        #
        # Push Collector object
        #
        if i == 0:
            colObjects = tc.collector_objects
            ret = eutils.generateLifCollectorConfig(tc, colObjects)
            if ret != api.types.status.SUCCESS:
                api.Logger.error("Unable to identify Collector Workload")
                tc.error = True
                return api.types.status.FAILURE

            ret = agent_api.PushConfigObjects(colObjects, [tc.naples.node_name])
            if ret != api.types.status.SUCCESS:
                api.Logger.error("Unable to push collector objects")
                tc.error = True
                return api.types.status.FAILURE
            continue

        #
        # Establish Forwarding set up between Naples-peer and Collectors
        #
        eutils.establishForwardingSetup(tc)

        #
        # Update Interface objects
        #
        ifObjects = tc.interface_objects
        ret = eutils.generateLifInterfaceConfig(tc, ifObjects, colObjects)
        if ret != api.types.status.SUCCESS:
            agent_api.DeleteConfigObjects(tc.collector_objects,
                                         [tc.naples.node_name])
            api.Logger.error("Unable to identify Uplink/LIF Interfaces")
            tc.error = True
            return api.types.status.FAILURE

        ret = agent_api.UpdateConfigObjects(ifObjects, [tc.naples.node_name])
        if ret != api.types.status.SUCCESS:
            agent_api.DeleteConfigObjects(tc.collector_objects,
                                         [tc.naples.node_name])
            api.Logger.error("Unable to update interface objects")
            tc.error = True
            return api.types.status.FAILURE

        req_tcpdump = api.Trigger_CreateExecuteCommandsRequest(serial = True)
        for c in range(0, len(tc.collector)):
            #
            # Set up TCPDUMP's on the collector
            #
            if tc.collector[c].IsNaples():
                cmd = "tcpdump -c 200 -XX -vv -nni %s ip proto gre and dst %s"%\
                      (tc.collector[c].interface, tc.collector_ip_address[c])
            else:
                cmd = "tcpdump -p -c 200 -XX -vv -nni %s ip proto gre\
                       and dst %s" %\
                      (tc.collector[c].interface, tc.collector_ip_address[c])
            eutils.add_command(req_tcpdump, tc.collector[c], cmd, True)

        resp_tcpdump = api.Trigger(req_tcpdump)
        for cmd in resp_tcpdump.commands:
            api.PrintCommandResults(cmd)

        #
        # Classic mode requires a delay to make sure that TCPDUMP background
        # process is fully up
        #
        if tc.classic_mode == True:
            time.sleep(2)

        #
        # Trigger packets for ERSPAN to take effect
        #
        tc.dest_port = '120'
        if api.GetNodeOs(tc.naples.node_name) == 'linux':
            eutils.triggerTrafficInClassicModeLinux(tc, tc.protocol, 
                                           tc.udp_count, tc.icmp_count)
        else:
            eutils.triggerTrafficInHostPinModeOrFreeBSD(tc, tc.protocol, 
                                           tc.udp_count, tc.icmp_count)

        #
        # Introduce a delay to make sure that TCPDUMP background
        # process is terminated only after all applicable packets are captured
        #
        if tc.classic_mode == True:
            time.sleep(2)

        #
        # Terminate TCPDUMP background process
        #
        term_resp_tcpdump = api.Trigger_TerminateAllCommands(resp_tcpdump)
        tc.resp_tcpdump = api.Trigger_AggregateCommandsResponse(resp_tcpdump,
                                                              term_resp_tcpdump)

        #
        # Dump sessions/flows/P4-tables for debug purposes
        #
        eutils.showSessionAndP4TablesForDebug(tc)

        # Delete the objects
        eutils.deGenerateLifInterfaceConfig(tc, tc.interface_objects, 
                                            tc.collector_objects)
        agent_api.UpdateConfigObjects(tc.interface_objects, 
                                     [tc.naples.node_name])

        agent_api.DeleteConfigObjects(tc.collector_objects, 
                                     [tc.naples.node_name])

        #
        # Make sure that Mirror-config has been removed
        #
        tc.resp_cleanup = eutils.showP4TablesForValidation(tc)

    return api.types.status.SUCCESS

def Verify(tc):
    if tc.ignore == True:
        return api.types.status.SUCCESS

    if tc.error == True:
        return api.types.status.FAILURE

    if tc.resp_tcpdump is None or tc.resp_cleanup is None:
        return api.types.status.FAILURE

    #
    # Validate ERSPAN packets reception
    #
    res_1 = eutils.validateErspanPackets(tc, tc.protocol, tc.feature, 
                                         tc.iterators.direction)

    #
    # Validate Config-cleanup
    #
    res_2 = eutils.validateConfigCleanup(tc)

    if res_1 == api.types.status.FAILURE or res_2 == api.types.status.FAILURE:
        return api.types.status.FAILURE

    return api.types.status.SUCCESS

def Teardown(tc):
    if tc.ignore == True:
        return api.types.status.SUCCESS

    #
    # Restore current Time-Close configs
    #
    if tc.classic_mode is not None and tc.classic_mode == False:
        update_timeout('tcp-close', tc.tcp_close_val)

    if tc.error == True:
        return api.types.status.FAILURE

    eutils.deEstablishCollectorSecondaryIPs(tc)

    return api.types.status.SUCCESS

