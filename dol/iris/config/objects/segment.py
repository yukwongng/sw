#! /usr/bin/python3
import pdb
import json
import infra.common.defs        as defs
import infra.common.objects     as objects
import infra.config.base        as base

import iris.config.resmgr            as resmgr
import iris.config.objects.endpoint  as endpoint
import iris.config.objects.enic      as enic
import iris.config.objects.network   as nw
import iris.config.objects.oif_list  as oiflist
import iris.config.objects.multicast_group as multicast_group

import iris.config.hal.api            as halapi
import iris.config.agent.api          as agentapi
import iris.config.hal.defs           as haldefs

from infra.common.logging       import logger
from infra.common.glopts        import GlobalOptions
from iris.config.store               import Store

global gl_pinif_iter
gl_pinif_iter = 0

class AgentSegmentObjectSpec:
    def __init__(self, seg):
        ipobj = seg.ipv4_pool.get()
        self.IPv4Subnet = ipobj.get() + '/16'
        self.IPv4Gateway = ipobj.get()

        ip6obj = seg.ipv6_pool.get()
        self.IPv6Subnet = ip6obj.get() + '/96'
        self.IPv6Gateway = ip6obj.get()

        self.VlanID = seg.vlan_id
        self.VxlanVNI = seg.vxlan_id
        return

class AgentSegmentObject(base.AgentObjectBase):
    def __init__(self, seg):
        super().__init__("Network", seg.GID(), seg.tenant.GID())
        self.spec = AgentSegmentObjectSpec(seg)
        return
    def to_JSON(self):
        seg = dict()
        seg["kind"] = "Network"
        seg["meta"] = {"name":self.meta.Name, "tenant":self.meta.Tenant, "namespace": self.meta.Namespace}
        seg["spec"] = {\
            "ipv4-subnet": self.spec.IPv4Subnet,\
            "ipv4-gateway": self.spec.IPv4Gateway,\
            "ipv6-subnet": self.spec.IPv6Subnet,\
            "ipv6-gateway": self.spec.IPv6Gateway,\
            "vlan-id": self.spec.VlanID,\
            # "vxlan-vni": self.spec.VxlanVNI\
        }
        return json.dumps(seg)

class SegmentObject(base.ConfigObjectBase):
    def __init__(self):
        super().__init__()
        self.Clone(Store.templates.Get('SEGMENT'))
        return

    def Init(self, tenant, spec):
        self.id = resmgr.SegIdAllocator.get()
        self.GID("Seg%04d" % self.id)

        self.label      = getattr(spec, 'label', None)
        if self.label is not None:
            self.label      = self.label.upper()
        self.spec       = spec
        self.access     = getattr(spec.endpoints, 'access', False)
        self.tenant     = tenant
        self.type       = spec.type.upper()
        self.native     = spec.native
        self.l4lb       = getattr(spec, 'l4lb', False)
        self.blackhole  = getattr(spec, 'blackhole', False)
        self.fabencap   = getattr(spec, 'fabencap', 'vlan')
        self.fabencap   = self.fabencap.upper()
        self.floodlist  = None

        self.nw_sgenable = False
        nwspec = getattr(spec, 'networks', None)
        if nwspec is not None:
            self.nw_sgenable = getattr(nwspec, 'sgenable', False)

        self.ep_sgenable = getattr(spec.endpoints, 'sgenable', False)

        self.pinnedif = None
        self.__pin_interface()

        self.vrf_id = None
        self.nw_ids  = []
        self.mbrif_ids = []

        if self.blackhole:
            self.vlan_id    = resmgr.BlackHoleSegVlanAllocator.get()
            self.vxlan_id   = resmgr.BlackHoleSegVxlanAllocator.get()
        else:
            self.vlan_id = resmgr.SegVlanAllocator.get()
            self.provider_vlan_id = resmgr.SegProviderVlanAllocator.get()
            self.vxlan_id = resmgr.SegVxlanAllocator.get()

        self.macaddr    = resmgr.RouterMacAllocator.get()
        if self.IsInfraSegment():
            self.subnet     = resmgr.TepIpSubnetAllocator.get()
            self.subnet6    = resmgr.TepIpv6SubnetAllocator.get()
        else:
            self.subnet     = resmgr.IpSubnetAllocator.get()
            self.subnet6    = resmgr.Ipv6SubnetAllocator.get()

        self.gipo = resmgr.GIPoAddressAllocator.get()

        self.eplearn = getattr(spec, 'eplearn', False)

        self.ipv4_pool  = resmgr.CreateIpv4AddrPool(self.subnet.get())
        self.ipv6_pool  = resmgr.CreateIpv6AddrPool(self.subnet6.get())

        if GlobalOptions.classic is False:
            if self.l4lb:
                self.bend_subnet = resmgr.L4LbBackendIpSubnetAllocator.get()
                self.bend_subnet6 = resmgr.L4LbBackendIpv6SubnetAllocator.get()
                self.bend_ipv4_pool = resmgr.CreateIpv4AddrPool(self.bend_subnet.get())
                self.bend_ipv6_pool = resmgr.CreateIpv6AddrPool(self.bend_subnet6.get())

        policy = "BROADCAST_FWD_POLICY_" + spec.broadcast.upper()
        self.broadcast_policy = haldefs.segment.BroadcastFwdPolicy.Value(policy)

        policy = "MULTICAST_FWD_POLICY_" + spec.multicast.upper()
        self.multicast_policy = haldefs.segment.MulticastFwdPolicy.Value(policy)

        self.obj_helper_ep = endpoint.EndpointObjectHelper()
        self.obj_helper_enic = enic.EnicObjectHelper()
        self.obj_helper_nw = nw.NetworkObjectHelper()
        self.obj_helper_oiflist = oiflist.OifListObjectHelper()
        self.obj_helper_mcgroup = multicast_group.MulticastGroupObjectHelper()

        self.hal_handle = None

        self.Show()
        if GlobalOptions.classic is False:
            self.obj_helper_nw.Generate(self)
        self.obj_helper_enic.Generate(self, self.spec.endpoints)
        self.obj_helper_ep.Generate(self, self.spec.endpoints)
        self.obj_helper_oiflist.Generate(self)
        self.obj_helper_mcgroup.Generate(self)

        self.floodlist = self.obj_helper_oiflist.GetOifList()

        self.Show(detail = True)
        return

    def IsNetworkSgEnabled(self):
        return self.nw_sgenable
    def IsEndpointSgEnabled(self):
        return self.ep_sgenable

    def AllocIpv4Address(self, backend = False):
        if backend:
            assert(self.l4lb)
            return self.bend_ipv4_pool.get()
        return self.ipv4_pool.get()
    def AllocIpv6Address(self, backend = False):
        if backend:
            assert(self.l4lb)
            return self.bend_ipv6_pool.get()
        return self.ipv6_pool.get()

    def IsTenantSegment(self):
        return self.type == 'TENANT'
    def IsInfraSegment(self):
        return self.type == 'INFRA'
    def IsSpanSegment(self):
        return self.type == 'SPAN'

    def IsNative(self):
        return self.native

    def IsFabEncapVlan(self):
        return self.fabencap == 'VLAN'
    def IsFabEncapVxlan(self):
        return self.fabencap == 'VXLAN'

    def IsIPV4EpLearnEnabled(self, remote_ep=False):
        return self.eplearn and self.tenant.IsIPV4EpLearnEnabled() and \
                (self.eplearn.remote if remote_ep else True)

    def IsIPV6EpLearnEnabled(self, remote_ep=False):
        return self.eplearn and self.tenant.IsIPV6EpLearnEnabled() and \
                (self.eplearn.remote if remote_ep else True)

    def __pin_interface_for_hostpin_mode(self):
        trunks = Store.GetTrunkingUplinks()
        global gl_pinif_iter
        gl_pinif_iter += 1
        gl_pinif_iter %= len(trunks)
        self.pinnedif = trunks[gl_pinif_iter]
        return

    def __pin_interface_for_classic(self):
        self.pinnedif = self.tenant.GetPinIf()
        return

    def __pin_interface(self):
        if self.tenant.IsHostPinned():
            self.__pin_interface_for_hostpin_mode()
        elif GlobalOptions.classic:
            self.__pin_interface_for_classic()
        else:
            return
        logger.info("- %s: Pinning to Interface: %s" %\
                       (self.GID(), self.pinnedif))
        return

    def __copy__(self):
        seg = SegmentObject()
        seg.id = self.id
        seg.hal_handle = self.hal_handle
        seg.type = self.type
        seg.fabencap = self.fabencap
        seg.vlan_id = self.vlan_id
        seg.vxlan_id = self.vxlan_id
        seg.multicast_policy = self.multicast_policy
        seg.broadcast_policy = self.broadcast_policy
        seg.nw_ids = self.nw_ids[:]
        seg.mbrif_ids = self.mbrif_ids[:]
        return seg


    def Equals(self, other, lgh):
        if not isinstance(other, self.__class__):
            return False
        fields = ["id", "hal_handle", "type", "fabencap", "vlan_id", "vxlan_id",
                  "multicast_policy", "broadcast_policy"]
        if not self.CompareObjectFields(other, fields, lgh):
            return False

        if set(self.nw_ids) != set(other.nw_ids):
            lgh.error("Network ids don't match Expected : %s, actual : %s"
                      %(set(self.nw_ids), set(other.nw_ids)))
            return False

        if set(self.mbrif_ids) != set(other.mbrif_ids):
            lgh.error("Member IF ids don't match Expected : %s, actual : %s"
                      %(set(self.uplink_ids), set(other.uplink_ids)))
            return False

        return True

    def Show(self, detail = False):
        logger.info("Segment = %s(%d)" % (self.GID(), self.id))
        logger.info("- FabEncap   = %s" % self.fabencap)
        logger.info("- VLAN       = %d" % self.vlan_id)

        if GlobalOptions.classic is False:
            if not self.IsInfraSegment():
                logger.info("- VXLAN      = %s" % self.vxlan_id)
                logger.info("- GIPo       = %s" % self.gipo.get())
            logger.info("- Native     = %s" % self.native)
            if self.blackhole:
                logger.info("- BlackHole  = %s" % self.blackhole)
            logger.info("- MAC        = %s" % self.macaddr.get())
            logger.info("- Subnet     = %s" % self.subnet.get())
            logger.info("- Subnet6    = %s" % self.subnet6.get())
        logger.info("- Broadcast  = %s" % self.spec.broadcast)
        logger.info("- Multicast  = %s" % self.spec.multicast)
        if self.floodlist:
            self.floodlist.Show()

        if detail == False: return
        self.obj_helper_ep.Show()
        self.obj_helper_enic.Show()
        return

    def Summary(self):
        summary = ''
        summary += 'GID:%s' % self.GID()
        summary += '/ID:%s' % self.id
        summary += '/Native:%s' % (self.native)
        summary += '/Rmac:%s' % (self.macaddr.get())
        summary += '/AccEnc:%d' % (self.vlan_id)
        if self.IsFabEncapVxlan():
            summary += '/FabEnc:0x%x' % self.vxlan_id
            summary += '/GIPo:%s' % self.gipo.get()
        else:
            summary += '/FabEnc:%d' % self.vlan_id
        return summary

    def GetDirectEnics(self, backend = False):
        if backend:
            return self.obj_helper_enic.backend_direct
        return self.obj_helper_enic.direct
    def GetPvlanEnics(self, backend = False):
        if backend:
            return self.obj_helper_enic.backend_pvlan
        return self.obj_helper_enic.pvlan
    def GetUsegEnics(self, backend = False):
        if backend:
            return self.obj_helper_enic.backend_useg
        return self.obj_helper_enic.useg
    def GetClassicEnics(self, backend = False):
        if backend: assert(0)
        return self.obj_helper_enic.classic
    def GetEps(self, backend = False):
        if backend:
            return self.obj_helper_ep.backend_eps
        return self.obj_helper_ep.eps
    def GetLocalEps(self, backend = False):
        if backend:
            return self.obj_helper_ep.backend_local
        return self.obj_helper_ep.local
    def GetRemoteEps(self, backend = False):
        if backend:
            return self.obj_helper_ep.backend_remote
        return self.obj_helper_ep.remote

    def GetExmProfile(self, profstr):
        if self.native == False:
            profstr += '_QTAG'
        else:
            profstr += '_NATIVE'
        profstr = 'GFT_EXMP_' + profstr
        profile = Store.objects.Get(profstr)
        return profile

    def GetTrspnProfile(self, profstr):
        if self.native == False:
            profstr += '_QTAG'
        else:
            profstr += '_NATIVE'
        profstr = 'GFT_TRSPNP_' + profstr
        profile = Store.objects.Get(profstr)
        return profile

    def ConfigureEndpoints(self):
        return self.obj_helper_ep.Configure()

    def ConfigureEnics(self):
        return self.obj_helper_enic.Configure()

    def ConfigureNetworks(self):
        return self.obj_helper_nw.Configure()

    def ConfigureMulticastGroups(self):
        return self.obj_helper_mcgroup.Configure()

    def ConfigureChildren(self):
        self.ConfigureEnics()
        self.ConfigureEndpoints()
        self.ConfigureMulticastGroups()
        return

    def AddSecurityGroup(self, sg):
        if self.IsNetworkSgEnabled() == False:
            return
        logger.info("- Adding SecurityGroup:%s to Segment:%s" %\
                       (sg.GID(), self.GID()))
        self.obj_helper_nw.AddSecurityGroup(sg)
        return

    def PrepareHALRequestSpec(self, req_spec):
        req_spec.key_or_handle.segment_id = self.id
        req_spec.vrf_key_handle.vrf_id = self.tenant.id
        #req_spec.access_encap.encap_type    = haldefs.common.ENCAP_TYPE_DOT1Q
        #req_spec.access_encap.encap_value   = self.vlan_id
        if self.IsFabEncapVxlan():
            req_spec.wire_encap.encap_type    = haldefs.common.ENCAP_TYPE_VXLAN
            req_spec.wire_encap.encap_value   = self.vxlan_id
            req_spec.gipo.ip_af                 = haldefs.common.IP_AF_INET
            req_spec.gipo.v4_addr               = self.gipo.getnum()
        else:
            req_spec.wire_encap.encap_type    = haldefs.common.ENCAP_TYPE_DOT1Q
            req_spec.wire_encap.encap_value   = self.vlan_id
        req_spec.mcast_fwd_policy = self.multicast_policy
        req_spec.bcast_fwd_policy = self.broadcast_policy
        req_spec.proxy_arp_enabled = getattr(self, "proxy_arp_enabled", False)

        if self.eplearn:
            if self.eplearn.arp_entry_timeout:
                req_spec.eplearn_cfg.arp.entry_timeout =  int(self.eplearn.arp_entry_timeout)
            if self.eplearn.dhcp:
                dhcp_server = req_spec.eplearn_cfg.dhcp.trusted_servers.add()
                dhcp_server = self.ipv4_pool.GetLast()
            if hasattr(self.eplearn, "dpkt"):
                req_spec.eplearn_cfg.dpkt.enabled = True
            if hasattr(self.eplearn, "arp_probe"):
                req_spec.eplearn_cfg.arp.probe_enabled = self.eplearn.arp_probe

        if (self.pinnedif != None):
            req_spec.pinned_uplink_if_key_handle.if_handle = self.pinnedif.hal_handle

        for nw in self.obj_helper_nw.nws:
            if nw.ip_prefix:
                nkh = req_spec.network_key_handle.add()
                nkh.nw_key.ip_prefix.CopyFrom(nw.ip_prefix)
                nkh.nw_key.vrf_key_handle.vrf_id = self.tenant.id
                self.nw_ids.append(nw.ip_prefix)

        from iris.config.objects.uplink      import UplinkHelper
        from iris.config.objects.uplinkpc    import UplinkPcHelper
        for uplink in UplinkHelper.trunks:
            ifkh = req_spec.if_key_handle.add()
            ifkh.interface_id = uplink.id
            self.mbrif_ids.append(uplink.id)

        for uplinkpc in UplinkPcHelper.trunks:
            ifkh = req_spec.if_key_handle.add()
            ifkh.interface_id = uplinkpc.id
            self.mbrif_ids.append(uplinkpc.id)

        return

    def ProcessHALResponse(self, req_spec, resp_spec):
        logger.info("- Segment %s = %s, vrf_id 0x%x" %\
                       (self.GID(), haldefs.common.ApiStatus.Name(resp_spec.api_status),
                           resp_spec.l2segment_status.vrf_id))
        self.hal_handle = resp_spec.l2segment_status.key_or_handle.l2segment_handle
        self.vrf_id = resp_spec.l2segment_status.vrf_id
        return

    def PrepareHALGetRequestSpec(self, get_req_spec):
        get_req_spec.key_or_handle.l2segment_handle = self.hal_handle
        get_req_spec.vrf_key_handle.vrf_id = self.tenant.id
        return

    def ProcessHALGetResponse(self, get_req_spec, get_resp_spec):
        if get_resp_spec.api_status == haldefs.common.ApiStatus.Value('API_STATUS_OK'):
            self.id = get_resp_spec.spec.key_or_handle.segment_id
            #self.vlan_id = get_resp_spec.spec.access_encap.encap_value
            if get_resp_spec.spec.wire_encap.encap_type ==  haldefs.common.ENCAP_TYPE_VXLAN:
                self.fabencap = 'VXLAN'
                self.vxlan_id = get_resp_spec.spec.wire_encap.encap_value
                self.gipo     = get_resp_spec.spec.gipo
            elif get_resp_spec.spec.wire_encap.encap_type ==  haldefs.common.ENCAP_TYPE_DOT1Q:
                self.fabencap = 'VLAN'
                self.vlan_id = get_resp_spec.spec.wire_encap.encap_value
            else:
                self.fabencap = None

            self.multicast_policy = get_resp_spec.spec.mcast_fwd_policy
            self.broadcast_policy = get_resp_spec.spec.bcast_fwd_policy

            self.nw_ids = []
            for nw_id in get_req_spec.spec.network_key_handle.nw_key.ip_prefix:
                self.nw_ids.append(nw_id)

            self.mbrif_ids = []
            for ifkh in get_req_spec.spec.if_key_handle:
                self.mbrif_ids.append(ifkh.interface_id)

        else:
            logger.error("- Segment %s = %s is missing." %\
                       (self.GID(), haldefs.common.ApiStatus.Name(get_resp_spec.api_status)))
            self.id = None
        return

    def Get(self):
        halapi.GetSegments([self])

    def IsFilterMatch(self, spec):
        return super().IsFilterMatch(spec.filters)

    def PrepareAgentObject(self):
        return AgentSegmentObject(self)

# Helper Class to Generate/Configure/Manage Segment Objects.
class SegmentObjectHelper:
    def __init__(self):
        self.segs = []
        self.bhseg = None
        self.backend_eps = None
        self.backend_remote_eps = None
        self.backend_remote_eps_tunneled = None
        return

    def Configure(self):
        logger.info("Configuring %d Segments." % len(self.segs))
        for seg in self.segs:
            seg.ConfigureNetworks()
        if GlobalOptions.agent:
            agentapi.ConfigureSegments(self.segs)
        else:
            halapi.ConfigureSegments(self.segs)
        return

    def ConfigurePhase2(self):
        for seg in self.segs:
            seg.ConfigureChildren()
        return

    def Generate(self, tenant, spec, count):
        logger.info("Creating %d Segments for Tenant = %s" %\
                       (count, tenant.GID()))
        for s in range(count):
            seg = SegmentObject()
            seg.Init(tenant, spec)
            self.segs.append(seg)
            if seg.blackhole:
                assert(self.bhseg is None)
                self.bhseg = seg
        return

    def AddToStore(self):
        Store.objects.SetAll(self.segs)
        return

    def GetEps(self, backend = False):
        eps = []
        for seg in self.segs:
            eps += seg.GetEps(backend)
        return eps

    def GetLocalEps(self, backend = False):
        eps = []
        for seg in self.segs:
            eps += seg.GetLocalEps(backend)
        return eps

    def GetRemoteEps(self, backend = False):
        eps = []
        for seg in self.segs:
            eps += seg.GetRemoteEps(backend)
        return eps

    def __get_backend_eps(self, remote):
        if self.backend_remote_eps: return
        remote_eps = self.GetRemoteEps(backend = True)
        self.backend_remote_eps_tunneled = []
        self.backend_remote_eps = []
        for ep in remote_eps:
            if ep.IsBehindTunnel():
                self.backend_remote_eps_tunneled.append(ep)
            else:
                self.backend_remote_eps.append(ep)
        #self.backend_remote_eps = self.GetRemoteEps(backend = True)
        logger.info("Remote L4LB Backend Pool:")
        for bend in self.backend_remote_eps:
            logger.info("- Backend: %s" % bend.GID())
        logger.info("Remote L4LB Backend Tunneled Pool:")
        for bend in self.backend_remote_eps_tunneled:
            logger.info("- Tunneled Backend: %s" % bend.GID())

        self.backend_eps = self.GetLocalEps(backend = True)
        logger.info("L4LB Backend Pool:")
        for bend in self.backend_eps:
            logger.info("- Backend: %s" % bend.GID())
        return

    def AllocL4LbBackend(self, remote, tnnled, pick_ep_in_l2seg_vxlan):
        self.__get_backend_eps(remote)
        if remote:
            if tnnled:
                eplist = self.backend_remote_eps_tunneled
            else:
                eplist = self.backend_remote_eps
        else:
            eplist = self.backend_eps
            count = len(eplist)
            found = False
            for c in range(count):
                ep = eplist[c]
                if pick_ep_in_l2seg_vxlan and ep.IsInL2SegVxlan():
                    found = True
                    break;
            if found:
                obj = eplist[c]
                del eplist[c]
                return obj



        obj = eplist[0]
        del eplist[0]
        return obj

    def GetBlackholeSegment(self):
        return self.bhseg


def GetMatchingObjects(selectors):
    segments =  Store.objects.GetAllByClass(SegmentObject)
    return [seg for seg in segments if seg.IsFilterMatch(selectors.segment)]
