#! /usr/bin/python3
import pdb

import iota.harness.infra.types as types
import iota.harness.infra.utils.parser as parser
import iota.harness.infra.svc as svc

import iota.protos.pygen.topo_svc_pb2 as topo_pb2
import iota.protos.pygen.types_pb2 as types_pb2

from iota.harness.infra.utils.logger import Logger as Logger

class _Testbed:
    def __init__(self):
        self.curr_ts = None     # Current Testsuite
        self.prev_ts = None     # Previous Testsute

        self.__read_warmd_json()
        self.__node_ips = []
        return

    def __read_warmd_json(self):
        self.tbspec = parser.JsonParse("/tmp/warmd.json")
        return
    
    def __prepare_TestBedMsg(self, ts):
        msg = topo_pb2.TestBedMsg()
        if not ts:
            return msg

        msg.switch_port_id = int(self.tbspec.n3k.port)
        msg.naples_image = ts.GetImages().naples
        msg.venice_image = ts.GetImages().venice
        msg.driver_sources = ts.GetImages().drivers
        msg.iota_agent_image = ts.GetImages().iagent

        for k,v in self.tbspec.Instances.__dict__.items():
            msg.ip_address.append(v)
            self.__node_ips.append(v)
        return msg

    def __cleanup_testbed(self):
        msg = self.__prepare_TestBedMsg(self.prev_ts)
        resp = svc.CleanupTestbed(msg)
        if resp.api_response.api_status != types_pb2.API_STATUS_OK:
            Logger.error("Failed to initialize testbed: ",
                         types_pb2.APIResponseType.Name(resp.api_response.api_status))
            return types.status.FAILURE
        self.data_vlans = None
        return types.status.SUCCESS

    def __init_testbed(self):
        msg = self.__prepare_TestBedMsg(self.curr_ts)
        resp = svc.InitTestbed(msg)
        if resp.api_response.api_status != types_pb2.API_STATUS_OK:
            Logger.error("Failed to initialize testbed: ",
                         types_pb2.APIResponseType.Name(resp.api_response.api_status))
            return types.status.FAILURE
        self.data_vlans = resp.allocated_vlans
        return types.status.SUCCESS

    def InitForTestsuite(self, ts):
        self.prev_ts = self.curr_ts
        self.curr_ts = ts
        status = self.__cleanup_testbed()
        if status != types.status.SUCCESS:
            return status

        status = self.__init_testbed()
        return status

    def ReserveNodeIpAddress(self):
        if len(self.__node_ips):
            node_ip = self.__node_ips[0]
            del self.__node_ips[0]
        else:
            Logger.error("No Nodes available in Testbed.")
            assert(0)
        return node_ip

Testbed = _Testbed()
