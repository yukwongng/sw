#! /usr/bin/python3
import grpc

from iota.harness.infra.utils.logger import Logger as Logger

import iota.protos.pygen.types_pb2 as types_pb2
import iota.protos.pygen.cfg_svc_pb2 as cfg_svc
import iota.protos.pygen.topo_svc_pb2 as topo_svc

import iota.harness.infra.store as store
import iota.harness.infra.types as types
import iota.harness.infra.utils.utils as utils
import iota.harness.infra.utils.parser as parser

from iota.harness.infra.glopts import GlobalOptions

IotaSvcChannel = None
TopoSvcStub = None
CfgSvcStub = None

def Init():
    server = 'localhost:' + str(GlobalOptions.svcport)
    Logger.info("Creating GRPC Channel to IOTA Service %s" % server)
    IotaSvcChannel = grpc.insecure_channel(server)
    Logger.debug("Waiting for IOTA Service to be UP")
    grpc.channel_ready_future(IotaSvcChannel).result()
    Logger.info("Connected to IOTA Service")

    global TopoSvcStub
    TopoSvcStub = topo_svc.TopologyApiStub(IotaSvcChannel)

    global CfgSvcStub
    CfgSvcStub = cfg_svc.ConfigMgmtApiStub(IotaSvcChannel)
    return

def __rpc(req, rpcfn):
    utils.LogMessageContents("Request", req, Logger.debug)
    resp = rpcfn(req)
    if resp.api_response.api_status != types_pb2.API_STATUS_OK:
        Logger.error("Error: ",
                     types_pb2.APIResponseType.Name(resp.api_response.api_status),
                     resp.api_response.error_msg)
        return None
    utils.LogMessageContents("Response", resp, Logger.debug)
    return resp

def CleanupTestbed(req):
    global TopoSvcStub
    Logger.info("Cleaning up Testbed:")
    return __rpc(req, TopoSvcStub.CleanUpTestBed)

def InitTestbed(req):
    global TopoSvcStub
    Logger.info("Initializing Testbed:")
    return __rpc(req, TopoSvcStub.InitTestBed)

def AddNodes(req):
    global TopoSvcStub
    Logger.info("Add Nodes:")
    return __rpc(req, TopoSvcStub.AddNodes)

def AddWorkloads(req):
    global TopoSvcStub
    Logger.info("Add Workloads:")
    return __rpc(req, TopoSvcStub.AddWorkloads)

def Trigger(req):
    global TopoSvcStub
    Logger.info("Trigger Message:")
    return __rpc(req, TopoSvcStub.Trigger)

def PushConfig(req):
    global CfgSvcStub
    Logger.info("Push Config:")
    return __rpc(req, CfgSvcStub.PushConfig)

def QueryConfig(req):
    global CfgSvcStub
    Logger.info("Query Config:")
    return __rpc(req, CfgSvcStub.QueryConfig)

def InitCfgService(req):
    global CfgSvcStub
    Logger.info("Init Config Service:")
    return __rpc(req, CfgSvcStub.InitCfgService)

def GenerateConfigs(req):
    global CfgSvcStub
    Logger.info("Generate Configs:")
    return __rpc(req, CfgSvcStub.GenerateConfigs)

def MakeCluster(req):
    global CfgSvcStub
    Logger.info("Make Cluster:")
    return __rpc(req, CfgSvcStub.MakeCluster)

def GetVeniceMgmtIpAddresses():
    return store.GetTestbed().GetCurrentTestsuite().GetTopology().GetVeniceMgmtIpAddresses()

def GetDataVlans():
    return store.GetTestbed().GetDataVlans()

def GetVeniceHostnames():
    return store.GetTestbed().GetCurrentTestsuite().GetTopology().GetVeniceHostnames()
