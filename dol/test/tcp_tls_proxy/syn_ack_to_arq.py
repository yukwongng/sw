# Testcase definition file.

import pdb
import copy

from config.store               import Store
from config.objects.proxycb_service    import ProxyCbServiceHelper
from config.objects.tcp_proxy_cb        import TcpCbHelper
from infra.common.objects import ObjectDatabase as ObjectDatabase
import test.callbacks.networking.modcbs as modcbs
from infra.common.logging import logger
from config.objects.cpucb        import CpuCbHelper
import test.tcp_tls_proxy.tcp_proxy as tcp_proxy
import test.tcp_tls_proxy.tcp_tls_proxy as tcp_tls_proxy
from infra.common.glopts import GlobalOptions

def Setup(infra, module):
    print("Setup(): Sample Implementation")
    modcbs.Setup(infra, module)
    return

def Teardown(infra, module):
    print("Teardown(): Sample Implementation.")
    return

def TestCaseSetup(tc):
    tc.pvtdata = ObjectDatabase(logger)
    tcp_proxy.SetupProxyArgs(tc)

    id = ProxyCbServiceHelper.GetFlowInfo(tc.config.flow._FlowObject__session)
    TcpCbHelper.main(id)
    tcbid = "TcpCb%04d" % id
    # 1. Configure TCB in HBM before packet injection
    tcb = tc.infra_data.ConfigStore.objects.db[tcbid]
    tcp_proxy.init_tcb_inorder(tc, tcb)
    tcb.debug_dol |= tcp_proxy.tcp_debug_dol_leave_in_arq
    # set tcb state to SYN_SENT(2)
    tcb.state = tcp_proxy.tcp_state_SYN_SENT
    tcb.SetObjValPd()

    # 2. Clone objects that are needed for verification
    rnmdr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["RNMDR"])
    rnmdr.GetMeta()
    rnmdr.GetRingEntries([rnmdr.pi])
    rnmpr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["RNMPR"])
    rnmpr.GetMeta()
    rnmpr.GetRingEntries([0])
    tnmdr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["TNMDR"])
    tnmdr.GetMeta()
    tnmpr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["TNMPR"])
    tnmpr.GetMeta()
    arq = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["CPU0000_ARQ"])
    arq.GetMeta()

    sesqid = "TCPCB%04d_SESQ" % id
    sesq = copy.deepcopy(tc.infra_data.ConfigStore.objects.db[sesqid])
    tlscbid = "TlsCb%04d" % id
    tlscb_cur = tc.infra_data.ConfigStore.objects.db[tlscbid]
    tlscb_cur.debug_dol = tcp_tls_proxy.tls_debug_dol_bypass_barco
    tlscb_cur.SetObjValPd()
    tlscb = copy.deepcopy(tlscb_cur)
    tcpcb = copy.deepcopy(tcb)

    tc.pvtdata.Add(tlscb)
    tc.pvtdata.Add(rnmdr)
    tc.pvtdata.Add(rnmpr)
    tc.pvtdata.Add(tnmdr)
    tc.pvtdata.Add(tnmpr)
    tc.pvtdata.Add(tcpcb)
    tc.pvtdata.Add(sesq)
    tc.pvtdata.Add(arq)
    return

def TestCaseVerify(tc):
    if GlobalOptions.dryrun:
        return True

    rnmdr = tc.pvtdata.db["RNMDR"]
    rnmpr = tc.pvtdata.db["RNMPR"]
    arq = tc.pvtdata.db["CPU0000_ARQ"]

    # 1. Fetch current values from Platform
    rnmdr_cur = tc.infra_data.ConfigStore.objects.db["RNMDR"]
    rnmdr_cur.GetMeta()
    rnmpr_cur = tc.infra_data.ConfigStore.objects.db["RNMPR"]
    rnmpr_cur.GetMeta()
    arq_cur = tc.infra_data.ConfigStore.objects.db["CPU0000_ARQ"]
    arq_cur.GetMeta()
    arq_cur.GetRingEntries([arq.pi])
    arq_cur.GetRingEntryAOL([0])

    print("arq pi 0x%x" % (arq.pi))
    print("arq cur pi 0x%x" % (arq_cur.pi))
    # 2. Verify descriptor
    if rnmdr.ringentries[rnmdr.pi].handle != arq_cur.ringentries[arq.pi].handle:
    #if rnmdr.ringentries[rnmdr.pi].handle != arq_cur.ringentries[0].handle:
        print("Descriptor handle not as expected in ringentries 0x%x 0x%x" % (rnmdr.ringentries[rnmdr.pi].handle, arq_cur.ringentries[arq.pi].handle))
        #print("Descriptor handle not as expected in ringentries 0x%x 0x%x" % (rnmdr.ringentries[rnmdr.pi].handle, arq_cur.ringentries[arq.pi].handle))
        return False

    if ((rnmdr.ringentries[rnmdr.pi].handle is not None) and (arq_cur.ringentries[arq.pi].handle is not None)):
        print("Descriptor handle as expected in ringentries 0x%x 0x%x" % (rnmdr.ringentries[rnmdr.pi].handle, arq_cur.ringentries[arq.pi].handle))
    # 3. Verify page
    if rnmpr.ringentries[0].handle != arq_cur.swdre_list[0].Addr1:
        print("Page handle not as expected in arq_cur.swdre_list")
        #return False

    return True


def TestCaseTeardown(tc):
    print("TestCaseTeardown(): Sample Implementation.")
    return
