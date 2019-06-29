#include <gtest/gtest.h>
#include "nic/hal/plugins/alg_ftp/core.hpp"
#include "nic/hal/iris/datapath/p4/include/defines.h"
#include "ftp_test.hpp"

#include <tins/tins.h>

using namespace hal::plugins::alg_ftp;
using namespace session;
using namespace fte;

hal_handle_t ftp_test::client_eph, ftp_test::server_eph, ftp_test::client_eph1, ftp_test::server_eph1;
hal_handle_t ftp_test::client_eph2, ftp_test::server_eph2;

TEST_F(ftp_test, ftp_session)
{
    SessionGetRequest       req;
    SessionGetResponseMsg   rsp;
    hal_ret_t          ret;
    uint8_t            ftp_port[] = {0x50, 0x4f, 0x52, 0x54, 0x20, 0x31, 0x30, 0x2c, 
                                   0x30, 0x2c, 0x30, 0x2c, 0x31, 0x2c, 0x31, 0x34,
                                   0x34, 0x2c, 0x32, 0x31, 0x31, 0x0d, 0x0a};
    uint8_t            ftp_port_rsp[] = {0x32, 0x30, 0x30, 0x20, 0x50, 0x4f, 0x52, 0x54,
                                       0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
                                       0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
                                       0x66, 0x75, 0x6c, 0x2e, 0x20, 0x43, 0x6f, 0x6e,
                                       0x73, 0x69, 0x64, 0x65, 0x72, 0x20, 0x75, 0x73,
                                       0x69, 0x6e, 0x67, 0x20, 0x50, 0x41, 0x53, 0x56,
                                       0x2e, 0x0d, 0x0a};

    // Create TCP control session
    // TCP SYN
    Tins::TCP tcp = Tins::TCP(FTP_PORT, 2000);
    tcp.flags(Tins::TCP::SYN);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(1200);
    ret = inject_ipv4_pkt(fte::FLOW_MISS_LIFQ, server_eph, client_eph, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_FALSE(ctx_.drop_flow());
    EXPECT_TRUE(ctx_.session()->iflow->pgm_attrs.mcast_en);
    EXPECT_TRUE(ctx_.session()->rflow->pgm_attrs.mcast_en);
    EXPECT_EQ(ctx_.session()->iflow->pgm_attrs.mcast_ptr, P4_NW_MCAST_INDEX_FLOW_REL_COPY);
    EXPECT_EQ(ctx_.session()->rflow->pgm_attrs.mcast_ptr, P4_NW_MCAST_INDEX_FLOW_REL_COPY);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_FTP);
    EXPECT_NE(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->rule_id, 0);
    hal::session_t *session = ctx_.session();
    EXPECT_NE(session->sfw_rule_id, 0);
    EXPECT_EQ(session->skip_sfw_reval, 0);
    EXPECT_EQ(session->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_EQ(ctx_.session()->idle_timeout, 30);

    // TCP SYN/ACK on ALG_CFLOW_LIFQ
    tcp = Tins::TCP(2000, FTP_PORT);
    tcp.flags(Tins::TCP::SYN | Tins::TCP::ACK);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, client_eph, server_eph, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(ctx_.session(), session);
 
    //TCP ACK on ALG_CFLOW_LIFQ
    tcp = Tins::TCP(FTP_PORT, 2000);
    tcp.flags(Tins::TCP::ACK);
    tcp.seq(1);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, server_eph, client_eph, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(ctx_.session(), session);

    //FTP PORT Command
    tcp = Tins::TCP(FTP_PORT, 2000) /
          Tins::RawPDU(ftp_port, sizeof(ftp_port));
    tcp.seq(1);
    //tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    //tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    //tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, server_eph, client_eph, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_EQ(ctx_.session(), session);

    //FTP PORT Response
    tcp = Tins::TCP(2000, FTP_PORT) /
         Tins::RawPDU(ftp_port_rsp, sizeof(ftp_port_rsp));
    tcp.seq(1);
    //tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    //tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    //tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, client_eph, server_eph, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_EQ(ctx_.session(), session);

    CHECK_ALLOW_TCP(client_eph, server_eph, 37075, FTP_DATA_PORT, "c:20 -> s:37075");
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_FTP);
    EXPECT_NE(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->rule_id, 0);
    EXPECT_EQ(ctx_.session()->skip_sfw_reval, 1);
    EXPECT_EQ(ctx_.session()->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_NE(ctx_.session()->sfw_rule_id, 0);
    EXPECT_EQ(ctx_.session()->idle_timeout, 30);
    CHECK_DENY_TCP(client_eph, server_eph, 37075, 2000, "c:2000 -> s:37075");
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_DENY);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_NONE);
    EXPECT_EQ(ctx_.session()->skip_sfw_reval, 0);
    EXPECT_EQ(ctx_.session()->sfw_action, nwsec::SECURITY_RULE_ACTION_DENY);
    EXPECT_NE(ctx_.session()->sfw_rule_id, 0);
    CHECK_DENY_TCP(server_eph, client_eph,  37075, 2000, "c:2000 -> s:37075"); 
    CHECK_DENY_TCP(client_eph, server_eph, 37075, 2001, "c:2001 -> s:37075");

    req.set_session_handle(session->hal_handle);
    ret = hal::session_get(req, &rsp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(rsp.response(0).status().alg(), nwsec::APP_SVC_FTP);
    EXPECT_TRUE(rsp.response(0).status().has_ftp_info());
    EXPECT_EQ(rsp.response(0).status().ftp_info().parse_error(), 0);
}

TEST_F(ftp_test, ftp_session_allow_mismatch)
{
    SessionGetRequest       req;
    SessionGetResponseMsg   rsp;
    hal_ret_t          ret;
    uint8_t            ftp_port[] = {0x50, 0x4f, 0x52, 0x54, 0x20, 0x31, 0x30, 0x2c,
                                   0x30, 0x2c, 0x30, 0x2c, 0x31, 0x2c, 0x31, 0x34,
                                   0x34, 0x2c, 0x32, 0x31, 0x31, 0x0d, 0x0a};
    uint8_t            ftp_port_rsp[] = {0x32, 0x30, 0x30, 0x20, 0x50, 0x4f, 0x52, 0x54,
                                       0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
                                       0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
                                       0x66, 0x75, 0x6c, 0x2e, 0x20, 0x43, 0x6f, 0x6e,
                                       0x73, 0x69, 0x64, 0x65, 0x72, 0x20, 0x75, 0x73,
                                       0x69, 0x6e, 0x67, 0x20, 0x50, 0x41, 0x53, 0x56,
                                       0x2e, 0x0d, 0x0a};

    // Create TCP control session
    // TCP SYN
    Tins::TCP tcp = Tins::TCP(FTP_PORT, 2000);
    tcp.flags(Tins::TCP::SYN);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(1200);
    ret = inject_ipv4_pkt(fte::FLOW_MISS_LIFQ, server_eph1, client_eph1, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_FALSE(ctx_.drop_flow());
    EXPECT_TRUE(ctx_.session()->iflow->pgm_attrs.mcast_en);
    EXPECT_TRUE(ctx_.session()->rflow->pgm_attrs.mcast_en);
    EXPECT_EQ(ctx_.session()->iflow->pgm_attrs.mcast_ptr, P4_NW_MCAST_INDEX_FLOW_REL_COPY);
    EXPECT_EQ(ctx_.session()->rflow->pgm_attrs.mcast_ptr, P4_NW_MCAST_INDEX_FLOW_REL_COPY);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_FTP);
    hal::session_t *session = ctx_.session();

    // TCP SYN/ACK on ALG_CFLOW_LIFQ
    tcp = Tins::TCP(2000, FTP_PORT);
    tcp.flags(Tins::TCP::SYN | Tins::TCP::ACK);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, client_eph1, server_eph1, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(ctx_.session(), session);

    //TCP ACK on ALG_CFLOW_LIFQ
    tcp = Tins::TCP(FTP_PORT, 2000);
    tcp.flags(Tins::TCP::ACK);
    tcp.seq(1);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, server_eph1, client_eph1, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(ctx_.session(), session);

    //FTP PORT Command
    tcp = Tins::TCP(FTP_PORT, 2000) /
          Tins::RawPDU(ftp_port, sizeof(ftp_port));
    tcp.seq(1);
    //tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    //tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    //tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, server_eph1, client_eph1, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_EQ(ctx_.session(), session);

    //FTP PORT Response
    tcp = Tins::TCP(2000, FTP_PORT) /
         Tins::RawPDU(ftp_port_rsp, sizeof(ftp_port_rsp));
    tcp.seq(1);
    //tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    //tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    //tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, client_eph1, server_eph1, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_EQ(ctx_.session(), session);

    CHECK_ALLOW_TCP(client_eph, server_eph1, 37075, FTP_DATA_PORT, "c:20 -> s:37075");
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_FTP);
    EXPECT_NE(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->rule_id, 0);
    CHECK_DENY_TCP(client_eph, server_eph1, 37075, 2000, "c:2000 -> s:37075");
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_DENY);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_NONE);
    CHECK_DENY_TCP(server_eph1, client_eph,  37075, 2000, "c:2000 -> s:37075");
    CHECK_DENY_TCP(client_eph, server_eph1, 37075, 2001, "c:2001 -> s:37075");

    req.set_session_handle(session->hal_handle);
    ret = hal::session_get(req, &rsp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(rsp.response(0).status().alg(), nwsec::APP_SVC_FTP);
    EXPECT_TRUE(rsp.response(0).status().has_ftp_info());
    EXPECT_EQ(rsp.response(0).status().ftp_info().parse_error(), 0);
}

TEST_F(ftp_test, ftp_session_not_allow_mismatch)
{
    SessionGetRequest  req;
    SessionGetResponse rsp;
    hal_ret_t          ret;
    uint8_t            ftp_port[] = {0x50, 0x4f, 0x52, 0x54, 0x20, 0x31, 0x30, 0x2c,
                                   0x30, 0x2c, 0x30, 0x2c, 0x31, 0x2c, 0x31, 0x34,
                                   0x34, 0x2c, 0x32, 0x31, 0x31, 0x0d, 0x0a};
    uint8_t            ftp_port_rsp[] = {0x32, 0x30, 0x30, 0x20, 0x50, 0x4f, 0x52, 0x54,
                                       0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
                                       0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
                                       0x66, 0x75, 0x6c, 0x2e, 0x20, 0x43, 0x6f, 0x6e,
                                       0x73, 0x69, 0x64, 0x65, 0x72, 0x20, 0x75, 0x73,
                                       0x69, 0x6e, 0x67, 0x20, 0x50, 0x41, 0x53, 0x56,
                                       0x2e, 0x0d, 0x0a};

    // Create TCP control session
    // TCP SYN
    Tins::TCP tcp = Tins::TCP(FTP_PORT, 2000);
    tcp.flags(Tins::TCP::SYN);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(1200);
    ret = inject_ipv4_pkt(fte::FLOW_MISS_LIFQ, server_eph2, client_eph2, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_FALSE(ctx_.drop_flow());
    EXPECT_TRUE(ctx_.session()->iflow->pgm_attrs.mcast_en);
    EXPECT_TRUE(ctx_.session()->rflow->pgm_attrs.mcast_en);
    EXPECT_EQ(ctx_.session()->iflow->pgm_attrs.mcast_ptr, P4_NW_MCAST_INDEX_FLOW_REL_COPY);
    EXPECT_EQ(ctx_.session()->rflow->pgm_attrs.mcast_ptr, P4_NW_MCAST_INDEX_FLOW_REL_COPY);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->sfw_action, nwsec::SECURITY_RULE_ACTION_ALLOW);
    EXPECT_EQ(ctx_.flow_log(hal::FLOW_ROLE_INITIATOR)->alg, nwsec::APP_SVC_FTP);
    hal::session_t *session = ctx_.session();

    // TCP SYN/ACK on ALG_CFLOW_LIFQ
    tcp = Tins::TCP(2000, FTP_PORT);
    tcp.flags(Tins::TCP::SYN | Tins::TCP::ACK);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, client_eph2, server_eph2, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(ctx_.session(), session);

    //TCP ACK on ALG_CFLOW_LIFQ
    tcp = Tins::TCP(FTP_PORT, 2000);
    tcp.flags(Tins::TCP::ACK);
    tcp.seq(1);
    tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, server_eph2, client_eph2, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_EQ(ctx_.session(), session);

    //FTP PORT Command
    tcp = Tins::TCP(FTP_PORT, 2000) /
          Tins::RawPDU(ftp_port, sizeof(ftp_port));
    tcp.seq(1);
    //tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    //tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    //tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, server_eph2, client_eph2, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_EQ(ctx_.session(), session);

    //FTP PORT Response
    tcp = Tins::TCP(2000, FTP_PORT) /
         Tins::RawPDU(ftp_port_rsp, sizeof(ftp_port_rsp));
    tcp.seq(1);
    //tcp.add_option(Tins::TCP::option(Tins::TCP::SACK_OK));
    //tcp.add_option(Tins::TCP::option(Tins::TCP::NOP));
    //tcp.mss(200);
    ret = inject_ipv4_pkt(fte::ALG_CFLOW_LIFQ, client_eph2, server_eph2, tcp);
    EXPECT_EQ(ret, HAL_RET_OK);
    EXPECT_FALSE(ctx_.drop());
    EXPECT_EQ(ctx_.session(), session);

    CHECK_DENY_TCP(client_eph, server_eph2, 37075, FTP_DATA_PORT, "c:20 -> s:37075");
    CHECK_DENY_TCP(server_eph2, client_eph,  37075, 2000, "c:2000 -> s:37075");
    CHECK_DENY_TCP(client_eph, server_eph2, 37075, 2001, "c:2001 -> s:37075");

}

TEST_F(ftp_test, app_sess_force_delete) {
    SessionGetResponseMsg  resp;
    hal::session_t        *session = NULL;
    hal_ret_t              ret;
    hal_handle_t           sess_hdl = 0;

    ret = hal::session_get_all(&resp);
    EXPECT_EQ(ret, HAL_RET_OK);
    for (int idx=0; idx<resp.response_size(); idx++) {
        SessionGetResponse rsp = resp.response(idx);
        if (rsp.status().has_ftp_info() && 
            rsp.status().ftp_info().iscontrol()) {
            sess_hdl = rsp.status().session_handle();
        }
    }
    
    // Invoke delete callback
    session = hal::find_session_by_handle(sess_hdl);
    ASSERT_TRUE(session != NULL);
    ret = session_delete(session, true);
    ASSERT_EQ(ret, HAL_RET_OK);
}
