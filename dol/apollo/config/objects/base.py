#! /usr/bin/python3

import copy
from collections import defaultdict

from infra.common.logging import logger
import infra.config.base as base

import apollo.config.agent.api as api
import apollo.config.utils as utils
import apollo.config.topo as topo

import apollo.test.utils.pdsctl as pdsctl

import types_pb2 as types_pb2

class StatusObjectBase(base.StatusObjectBase):
    def __init__(self, objtype):
        super().__init__()
        self.ObjType = objtype
        self.HwId = -1
        return

    def GetHwId(self):
        return self.HwId

class ConfigObjectBase(base.ConfigObjectBase):
    def __init__(self, objtype, node):
        super().__init__()
        self.Origin = topo.OriginTypes.FIXED
        self.HwHabitant = True
        self.Singleton = False
        self.ObjType = objtype
        self.Parent = None
        self.Children = []
        self.Deps = defaultdict(list)
        self.Precedent = None
        self.Mutable = False
        self.deleted = False
        self.Dirty = False
        self.Node = node
        return

    def __get_GrpcMsg(self, op):
        grpcreq = api.client[self.Node].GetGRPCMsgReq(self.ObjType, op)
        if grpcreq is None:
            logger.error("GRPC req method not added for obj:%s op:%s" %(self.ObjType, op))
            assert(0)
        return grpcreq()

    def __populate_BatchContext(self, grpcmsg, cookie):
        grpcmsg.BatchCtxt.BatchCookie = cookie
        return

    def IsV4(self):
        af = getattr(self, 'AddrFamily', None)
        if af == 'IPV4':
            return True
        return False

    def IsV6(self):
        af = getattr(self, 'AddrFamily', None)
        if af == 'IPV6':
            return True
        return False

    def AddChild(self, child):
        child.Parent = self
        self.Children.append(child)

    def GetDependees(self, node):
        # returns the list of dependees
        dependees = []
        return dependees

    def BuildDependency(self):
        dependees = self.GetDependees(self.Node)
        for dependee in dependees:
            # add ourself as an dependent to dependee
            dependee.AddDependent(self)
        return

    def AddDependent(self, dep):
        self.Deps[dep.ObjType].append(dep)

    def DeriveOperInfo(self):
        self.BuildDependency()
        return

    def SetHwHabitant(self, value):
        self.HwHabitant = value
        if self.HwHabitant == True and self.HasPrecedent() == True:
             self.Precedent.HwHabitant = False

    def IsHwHabitant(self):
        return self.HwHabitant

    def SetSingleton(self, value):
        self.Singleton = value

    def IsSingleton(self):
        return self.Singleton

    def SetDirty(self, value):
        self.Dirty = value

    def SetOrigin(self, origintype):
        self.Origin = origintype

    def IsOriginFixed(self):
        return True if (self.Origin == topo.OriginTypes.FIXED) else False

    def HasPrecedent(self):
         return False if (self.Precedent == None) else True

    def GetPrecedent(self):
         return self.Precedent

    def Create(self, spec=None):
        utils.CreateObject(self)
        return

    def Read(self, spec=None, expRetCode=types_pb2.API_STATUS_OK):
        if self.Dirty == False:
            return utils.ReadObject(self, expRetCode)
        else:
            logger.info("Not reading object from Hw since it is marked Dirty")
            return True

    def ReadAfterDelete(self, spec=None):
        return self.Read(types_pb2.API_STATUS_NOT_FOUND)

    def Delete(self, spec=None):
        utils.DeleteObject(self)
        return

    def UpdateNotify(self, dObj):
        return

    def RollbackMany(self, attrlist):
        if self.HasPrecedent():
            for attr in attrlist:
                setattr(self, attr, getattr(self.Precedent, attr))
        return

    def CopyObject(self):
        clone = copy.copy(self)
        return clone

    def Update(self, spec=None):
        if self.Mutable:
            logger.info("Update Obj %s" % repr(self))
            clone = self.CopyObject()
            clone.Precedent = None
            self.Precedent = clone
            self.HwHabitant = False
            self.UpdateAttributes()
            logger.info("Updated values - Obj %s" % repr(self))
            self.CommitUpdate()
        return

    def RollbackUpdate(self, spec=None):
        self.PrepareRollbackUpdate(spec)
        self.CommitUpdate(spec)
        return

    def PrepareRollbackUpdate(self, spec=None):
        if self.HasPrecedent():
            self.RollbackAttributes()
            self.Precedent = None
            self.HwHabitant = False
            self.SetDirty(True)
        return

    def CommitUpdate(self, spec=None):
        self.SetDirty(False)
        return utils.UpdateObject(self)

    def ValidateSpec(self, spec):
        logger.error("Method not implemented by class: %s" % self.__class__)
        assert(0)
        return False

    def ValidateStats(self, stats):
        return True

    def ValidateStatus(self, status):
        return True

    def ValidateYamlSpec(self, spec):
        logger.error("Method not implemented by class: %s" % self.__class__)
        assert(0)
        return False

    def ValidateYamlStats(self, stats):
        return True

    def ValidateYamlStatus(self, status):
        return True

    def Equals(self, obj, spec):
        return True

    def MarkDeleted(self, flag=True):
        self.deleted = flag
        return

    def IsDeleted(self):
        return self.deleted

    def SetBaseClassAttr(self):
        logger.error("Method not implemented by class: %s" % self.__class__)
        assert(0)
        return

    def PopulateKey(self, grpcmsg):
        logger.error("Method not implemented by class: %s" % self.__class__)
        assert(0)
        return

    def PopulateSpec(self, grpcmsg):
        logger.error("Method not implemented by class: %s" % self.__class__)
        assert(0)
        return

    def GetGrpcCreateMessage(self, cookie=0):
        grpcmsg = self.__get_GrpcMsg(api.ApiOps.CREATE)
        self.__populate_BatchContext(grpcmsg, cookie)
        self.PopulateSpec(grpcmsg)
        return grpcmsg

    def GetGrpcReadMessage(self):
        grpcmsg = self.__get_GrpcMsg(api.ApiOps.GET)
        self.PopulateKey(grpcmsg)
        return grpcmsg

    def GetGrpcUpdateMessage(self, cookie=0):
        grpcmsg = self.__get_GrpcMsg(api.ApiOps.UPDATE)
        self.__populate_BatchContext(grpcmsg, cookie)
        self.PopulateSpec(grpcmsg)
        return grpcmsg

    def GetGrpcDeleteMessage(self, cookie=0):
        grpcmsg = self.__get_GrpcMsg(api.ApiOps.DELETE)
        self.__populate_BatchContext(grpcmsg, cookie)
        self.PopulateKey(grpcmsg)
        return grpcmsg

class ConfigClientBase(base.ConfigClientBase):
    def __init__(self, objtype, maxlimit=0):
        super().__init__()
        self.Objs = defaultdict(dict)
        self.ObjType = objtype
        self.Maxlimit = maxlimit
        return

    def IsValidConfig(self, node):
        count = self.GetNumObjects(node)
        if count > self.Maxlimit:
            return False, "%s count %d exceeds allowed limit of %d" % \
                          (self.ObjType, count, self.Maxlimit)
        return True, ""

    def GetKeyfromSpec(self, spec, yaml=False):
        if yaml: return spec['id']
        return spec.Id

    def Objects(self, node):
        if self.Objs.get(node, None):
	        return self.Objs[node].values()
        return []

    def GetNumHwObjects(self, node):
        count = len(self.Objects(node))
        # TODO can be improved, if object has a reference to gen object
        for obj in self.Objects(node):
            if (obj.HwHabitant == False):
                count = count - 1
        return count

    def GetNumObjects(self, node):
        return len(self.Objects(node))

    def GetObjectByKey(self, node, key):
        return self.Objs[node].get(key, None)

    def GetObjectsByKeys(self, node, keys, filterfn=None):
        return list(filter(filterfn, map(lambda key: self.GetObjectByKey(node, key), keys)))

    def GetObjectType(self):
        return self.ObjType

    def ShowObjects(self, node):
        for obj in self.Objects(node):
            obj.Show()
        return

    def __get_GrpcMsg(self, node, op):
        grpcreq = api.client[node].GetGRPCMsgReq(self.ObjType, op)
        if grpcreq is None:
            logger.error("GRPC req method not added for obj:%s op:%s" %(self.ObjType, op))
            assert(0)
        return grpcreq()

    def GetGrpcReadAllMessage(self, node):
        grpcmsg = self.__get_GrpcMsg(node, api.ApiOps.GET)
        return grpcmsg

    def ValidateGrpcRead(self, node, getResp):
        if utils.IsDryRun(): return True
        numObjs = 0
        for obj in getResp:
            if not utils.ValidateGrpcResponse(obj):
                logger.error("GRPC get request failed for ", obj)
                return False
            for resp in obj.Response:
                numObjs += 1
                key = self.GetKeyfromSpec(resp.Spec)
                cfgObj = self.GetObjectByKey(node, key)
                if not utils.ValidateObject(cfgObj, resp):
                    logger.error("GRPC read validation failed for ", obj)
                    cfgObj.Show()
                    return False
                if hasattr(cfgObj, 'Status'):
                    cfgObj.Status.Update(resp.Status)

        assert(numObjs == self.GetNumHwObjects(node))
        return True

    def GrpcRead(self, node):
        # read all via grpc
        msg = self.GetGrpcReadAllMessage(node)
        resp = api.client[node].Get(self.ObjType, [msg])
        if not self.ValidateGrpcRead(node, resp):
            logger.critical("Object validation failed for %s" % (self.ObjType))
            assert(0)
        return

    def ValidatePdsctlRead(self, node, ret, stdout):
        if utils.IsDryRun(): return True
        if not ret:
            logger.error("pdsctl show cmd failed for ", self.ObjType)
            return False
        # split output per object
        cmdop = stdout.split("---")
        assert((len(cmdop) - 1) == self.GetNumHwObjects(node))
        for op in cmdop:
            yamlOp = utils.LoadYaml(op)
            if not yamlOp:
                continue
            key = self.GetKeyfromSpec(yamlOp['spec'], yaml=True)
            cfgObj = self.GetObjectByKey(node, key)
            if not utils.ValidateObject(cfgObj, yamlOp, yaml=True):
                logger.error("pdsctl read validation failed for ", op)
                cfgObj.Show()
                return False
        return True

    def PdsctlRead(self, node):
        # read all via pdsctl
        ret, op = pdsctl.GetObjects(self.ObjType)
        if not self.ValidatePdsctlRead(node, ret, op):
            logger.critical("Object validation failed for ", self.ObjType, ret, op)
            assert(0)
        return

    def ReadObjects(self, node):
        logger.info("Reading %s Objects" % (self.ObjType.name))
        self.GrpcRead(node)
        self.PdsctlRead(node)
        return

    def CreateObjects(self, node):
        fixed, discovered = [], []
        for obj in self.Objects(node):
            (fixed if obj.IsOriginFixed() else discovered).append(obj)

        logger.info("%s objects: fixed: %d discovered %d" %(self.ObjType.name, len(fixed), len(discovered)))
        # set HwHabitant to false for discovered objects
        for obj in discovered:
            obj.SetHwHabitant(False)

        # return if there is no fixed object
        if len(fixed) == 0:
            logger.info(f"Skip Creating {self.ObjType.name} Objects in agent")
            return

        self.ShowObjects(node)
        logger.info(f"Creating {self.ObjType.name} Objects in agent")
        cookie = utils.GetBatchCookie(node)
        msgs = list(map(lambda x: x.GetGrpcCreateMessage(cookie), fixed))
        api.client[node].Create(self.ObjType, msgs)
        #TODO: Add validation for create
        return
