#include "net_plugin.hpp"
#include <session.hpp>

namespace hal {
namespace net {

static inline session::FlowAction
pol_check_sg_policy(fte::ctx_t &ctx)
{
    if (ctx.protobuf_request()) {
        return ctx.sess_spec()->initiator_flow().flow_data().flow_info().flow_action();
    }

    // TODO(goli) check Security policy
    return session::FLOW_ACTION_ALLOW;
}

//------------------------------------------------------------------------------
// extract all the TCP related state from session spec
//------------------------------------------------------------------------------
static hal_ret_t
extract_session_state_from_spec (fte::flow_state_t *flow_state,
                                 const session::FlowData& flow_data)
{
    auto conn_track_info = flow_data.conn_track_info();
    flow_state->state = flow_data.flow_info().tcp_state();
    flow_state->tcp_seq_num = conn_track_info.tcp_seq_num();
    flow_state->tcp_ack_num = conn_track_info.tcp_ack_num();
    flow_state->tcp_win_sz = conn_track_info.tcp_win_sz();
    flow_state->tcp_win_scale = conn_track_info.tcp_win_scale();
    flow_state->tcp_mss = conn_track_info.tcp_mss();
    flow_state->create_ts = conn_track_info.flow_create_ts();
    flow_state->last_pkt_ts = flow_state->create_ts;
    flow_state->packets = conn_track_info.flow_packets();
    flow_state->bytes = conn_track_info.flow_bytes();
    flow_state->exception_bmap = conn_track_info.exception_bits();

    return HAL_RET_OK;
}


static inline bool
conn_tracking_configured(fte::ctx_t &ctx)
{
    if (ctx.protobuf_request()) {
        return ctx.sess_spec()->conn_track_en();
    }

    if (ctx.key().proto != types::IPPROTO_TCP) {
        return false;
    }

    // TODO(goli) check Security profile
    return false;
}

fte::pipeline_action_t
dfw_exec(fte::ctx_t& ctx)
{
    HAL_TRACE_DEBUG("Invoking firewall feature");

    fte::flow_update_t flowupd = {type: fte::FLOWUPD_ACTION};

    flowupd.action = pol_check_sg_policy(ctx);
    ctx.update_flow(hal::FLOW_ROLE_INITIATOR, flowupd);
    ctx.update_flow(hal::FLOW_ROLE_RESPONDER, flowupd);

    // connection tracking
    if (!conn_tracking_configured(ctx)) {
        return fte::PIPELINE_CONTINUE;
    }

    //iflow
    flowupd = {type: fte::FLOWUPD_FLOW_STATE};
    if (ctx.protobuf_request()) {
        extract_session_state_from_spec(&flowupd.flow_state,
                                        ctx.sess_spec()->initiator_flow().flow_data());
        flowupd.flow_state.syn_ack_delta = ctx.sess_spec()->iflow_syn_ack_delta();
    } else {
        flowupd.flow_state.state = session::FLOW_TCP_STATE_TCP_SYN_RCVD;
    }

    ctx.update_flow(hal::FLOW_ROLE_INITIATOR, flowupd); 

    //rflow
    flowupd = {type: fte::FLOWUPD_FLOW_STATE};
    if (ctx.protobuf_request()) {
        extract_session_state_from_spec(&flowupd.flow_state,
                                        ctx.sess_spec()->responder_flow().flow_data());
    } else {
        flowupd.flow_state.state = session::FLOW_TCP_STATE_INIT;
    }
    ctx.update_flow(hal::FLOW_ROLE_RESPONDER, flowupd);

    return fte::PIPELINE_CONTINUE;
}

} // namespace net
} // namespace hal
