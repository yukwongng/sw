# Testcase definition file.

import pdb
import copy

from iris.config.store               import Store
from iris.config.objects.proxycb_service    import ProxyCbServiceHelper
from iris.config.objects.tcp_proxy_cb        import TcpCbHelper
from infra.common.objects import ObjectDatabase as ObjectDatabase
import iris.test.callbacks.networking.modcbs as modcbs
from infra.common.logging import logger
from iris.config.objects.cpucb        import CpuCbHelper
import iris.test.tcp_tls_proxy.tcp_tls_proxy as tcp_tls_proxy
import iris.test.tcp_tls_proxy.tcp_proxy as tcp_proxy
from infra.common.glopts import GlobalOptions


def Setup(infra, module):
    print("Setup(): Sample Implementation")
    modcbs.Setup(infra, module)
    return

def Teardown(infra, module):
    print("Teardown(): Sample Implementation.")
    return

def TestCaseSetup(tc):
    tc.pvtdata = ObjectDatabase()

    tcp_proxy.SetupProxyArgs(tc)

    id = ProxyCbServiceHelper.GetFlowInfo(tc.config.flow._FlowObject__session)
    TcpCbHelper.main(id)
    tcbid = "TcpCb%04d" % id
    # 1. Configure TCB in HBM before packet injection
    tcb = tc.infra_data.ConfigStore.objects.db[tcbid]
    tcp_proxy.init_tcb_inorder(tc, tcb)
    # set tcb state to ESTABLISHED(1)
    tcb.state = tcp_proxy.tcp_state_ESTABLISHED
    tcb.debug_dol_tx |= tcp_proxy.tcp_tx_debug_dol_dont_tx
    tcb.SetObjValPd()

    # 2. Clone objects that are needed for verification
    rnmdpr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["RNMDPR_BIG"])
    rnmdpr.GetMeta()
    tnmdpr = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["TNMDPR_BIG"])
    tnmdpr.GetMeta()
    arq = copy.deepcopy(tc.infra_data.ConfigStore.objects.db["CPU0000_ARQ"])

    sesqid = "TCPCB%04d_SESQ" % id
    sesq = copy.deepcopy(tc.infra_data.ConfigStore.objects.db[sesqid])
    sesq.GetMeta()
    tlscbid = "TlsCb%04d" % id
    tlscb_cur = tc.infra_data.ConfigStore.objects.db[tlscbid]
    #tlscb_cur.debug_dol = (tcp_tls_proxy.tls_debug_dol_leave_in_arq | tcp_tls_proxy.tls_debug_dol_bypass_proxy | tcp_tls_proxy.tls_debug_dol_bypass_barco | tcp_tls_proxy.tls_debug_dol_fake_handshake_msg)
    tlscb_cur.debug_dol = (tcp_tls_proxy.tls_debug_dol_bypass_proxy | tcp_tls_proxy.tls_debug_dol_bypass_barco | tcp_tls_proxy.tls_debug_dol_arm_loop_ctlr_pkts)
    tlscb_cur.other_fid = 0xffff
    tlscb_cur.is_decrypt_flow = 1
    tlscb_cur.serq_pi = 0
    tlscb_cur.serq_ci = 0
    tlscb_cur.SetObjValPd()
    tlscb = copy.deepcopy(tlscb_cur)
    tlscb.GetObjValPd()
    tcpcb = copy.deepcopy(tcb)
    tcpcb.GetObjValPd()

    tc.pvtdata.Add(tlscb)
    tc.pvtdata.Add(rnmdpr)
    tc.pvtdata.Add(tnmdpr)
    tc.pvtdata.Add(tcpcb)
    tc.pvtdata.Add(sesq)
    tc.pvtdata.Add(arq)
    return

def TestCaseVerify(tc):
    if GlobalOptions.dryrun:
        return True

    id = ProxyCbServiceHelper.GetFlowInfo(tc.config.flow._FlowObject__session)
    tlscbid = "TlsCb%04d" % id
    tlscb = tc.pvtdata.db[tlscbid]
    tlscb_cur = tc.infra_data.ConfigStore.objects.db[tlscbid]
    print("pre-sync: tnmdpr_alloc %d enc_requests %d" % (tlscb_cur.tnmdpr_alloc, tlscb_cur.enc_requests))
    print("pre-sync: rnmdpr_free %d enc_completions %d" % (tlscb_cur.rnmdpr_free, tlscb_cur.enc_completions))
    print("pre-sync: pre_debug_stage0_7_thread 0x%x post_debug_stage0_7_thread 0x%x" % (tlscb_cur.pre_debug_stage0_7_thread, tlscb_cur.post_debug_stage0_7_thread))
    tlscb_cur.GetObjValPd()
    print("post-sync: tnmdpr_alloc %d enc_requests %d" % (tlscb_cur.tnmdpr_alloc, tlscb_cur.enc_requests))
    print("post-sync: rnmdpr_free %d enc_completions %d" % (tlscb_cur.rnmdpr_free, tlscb_cur.enc_completions))
    print("post-sync: pre_debug_stage0_7_thread 0x%x post_debug_stage0_7_thread 0x%x" % (tlscb_cur.pre_debug_stage0_7_thread, tlscb_cur.post_debug_stage0_7_thread))



        
    tcbid = "TcpCb%04d" % id
    tcb = tc.pvtdata.db[tcbid]
    # 2. Verify pi/ci got update got updated for ARQ
    tcb_cur = tc.infra_data.ConfigStore.objects.db[tcbid]
    print("pre-sync: tcb_cur.sesq_pi %d tcb_cur.sesq_ci %d" % (tcb_cur.sesq_pi, tcb_cur.sesq_ci))
    tcb_cur.GetObjValPd()
    print("post-sync: tcb_cur.sesq_pi %d tcb_cur.sesq_ci %d" % (tcb_cur.sesq_pi, tcb_cur.sesq_ci))
    if (tcb_cur.sesq_pi != (tcb.sesq_pi+1) or tcb_cur.sesq_ci != (tcb.sesq_ci+1)):
        print("sesq pi/ci not as expected")
        #VijayV to enable this test after ci is fixed in p4+
        #return False


    # 3. Fetch current values from Platform
    rnmdpr_cur = tc.infra_data.ConfigStore.objects.db["RNMDPR_BIG"]
    rnmdpr_cur.GetMeta()
    rnmdpr = tc.pvtdata.db["RNMDPR_BIG"]

    tnmdpr_cur = tc.infra_data.ConfigStore.objects.db["TNMDPR_BIG"]
    tnmdpr_cur.GetMeta()
    tnmdpr = tc.pvtdata.db["TNMDPR_BIG"]
    arq = tc.pvtdata.db["CPU0000_ARQ"]

    arq_cur = tc.infra_data.ConfigStore.objects.db["CPU0000_ARQ"]
    arq_cur.GetMeta()

    print("bytes_rcvd = %d:" % tcb.bytes_rcvd)
    print("pkts_rcvd = %d:" % tcb.pkts_rcvd)


    print("bytes_rcvd = %d:" % tcb_cur.bytes_rcvd)
    print("pkts_rcvd = %d:" % tcb_cur.pkts_rcvd)

    print("bytes_sent = %d:" % tcb.bytes_sent)
    print("pkts_sent = %d:" % tcb.pkts_sent)

    print("bytes_sent = %d:" % tcb_cur.bytes_sent)
    print("pkts_sent = %d:" % tcb_cur.pkts_sent)

    # 4. Verify pkt rx stats
    if tcb_cur.pkts_rcvd != tcb.pkts_rcvd + 1:
        print("pkt rx stats not as expected")
        return False

    if ((tcb_cur.bytes_rcvd - tcb.bytes_rcvd) != tc.packets.Get('PKT1').payloadsize):
        print("Warning! pkt rx byte stats not as expected %d %d" % (tcb_cur.bytes_rcvd, tcb.bytes_rcvd))
        return False
    
    # 5. Verify pkt tx stats
    # If set tcp_tx_debug_dol_dont_tx, then pkt does not tx and stats should not
    # increment either
    #if tcb_cur.pkts_sent != tcb.pkts_sent + 1:
        #print("pkt tx stats (%d) not as expected (%d)" % (tcb_cur.pkts_sent, tcb.pkts_sent))
        #return False

    #if ((tcb_cur.bytes_sent - tcb.bytes_sent) != 71): # TODO: remove hardcoding
        #print("Warning! pkt tx byte stats not as expected %d %d" % (tcb_cur.bytes_sent, tcb.bytes_sent))
        #return False


    return True

def TestCaseTeardown(tc):
    print("TestCaseTeardown(): Sample Implementation.")
    return
