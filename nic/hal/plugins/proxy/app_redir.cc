#include "nic/hal/src/proxy.hpp"
#include "nic/hal/plugins/proxy/proxy_plugin.hpp"
#include "nic/hal/pd/common/cpupkt_api.hpp"
#include "nic/p4/nw/include/defines.h"

namespace hal {
namespace proxy {

hal::pd::cpupkt_ctxt_t* app_chain_ctx = NULL;

static inline hal_ret_t
app_redir_init(void)
{
    hal_ret_t   ret = HAL_RET_OK;

    if (app_chain_ctx == NULL)   {
        app_chain_ctx = hal::pd::cpupkt_ctxt_alloc_init();
        if (app_chain_ctx == NULL)   {
            ret = HAL_RET_NO_RESOURCE;
        }
    }

    return ret;
}

static inline hal_ret_t
app_redir_pkt_send(fte::ctx_t& ctx)
{
    uint8_t                         *pkt;
    const fte::cpu_rxhdr_t          *cpu_rxhdr;
    size_t                          pkt_len;
    hal::pd::cpu_to_p4plus_header_t cpu_header;
    hal::pd::p4plus_to_p4_header_t  p4_header;
    hal_ret_t                       ret;

    /*
     * Register current ctx chain queue ID with cpupkt API
     * for sending packet.
     */
    cpu_rxhdr = ctx.cpu_rxhdr();
    ret = hal::pd::cpupkt_register_tx_queue(app_chain_ctx, 
                                            types::WRING_TYPE_APP_REDIR_RAWC,
                                            cpu_rxhdr->qid);
    if (ret == HAL_RET_OK) {
        pkt = ctx.pkt();
        pkt_len = ctx.pkt_len();
        memset(&cpu_header, 0, sizeof(cpu_header)); 
        memset(&p4_header, 0, sizeof(p4_header)); 
        cpu_header.src_lif = cpu_rxhdr->src_lif;
        p4_header.flags |= P4PLUS_TO_P4_FLAGS_LKP_INST;

        HAL_TRACE_DEBUG("{} pkt_len {} src_lif {}", __FUNCTION__,
                        pkt_len, cpu_header.src_lif);

        ret = hal::pd::cpupkt_send(app_chain_ctx,
                                   types::WRING_TYPE_APP_REDIR_RAWC,
                                   cpu_rxhdr->qid,
                                   &cpu_header,
                                   &p4_header,
                                   pkt,
                                   pkt_len,
                                   SERVICE_LIF_APP_REDIR,
                                   APP_REDIR_RAWC_QTYPE,
                                   cpu_rxhdr->qid,
                                   0);
    }

    return ret;
}

static inline hal_ret_t
app_redir_flow_fwding_update(fte::ctx_t&ctx,
                             flow_key_t &flow_key,
                             proxy_flow_info_t* pfi)
{
    fte::flow_update_t flowupd = {type: fte::FLOWUPD_FWDING_INFO};

    HAL_TRACE_DEBUG("app_redir flow forwarding role: {} qid1: {} qid2: {}",
                    ctx.role(), pfi->qid1, pfi->qid2);
    HAL_TRACE_DEBUG("app_redir updating lport = {} for sport = {} dport = {}",
                    pfi->proxy->meta->lif_info[0].lport_id,
                    flow_key.sport, flow_key.dport);

    // update fwding info
    flowupd.fwding.lport = pfi->proxy->meta->lif_info[0].lport_id;
    flowupd.fwding.qid_en = true;
    flowupd.fwding.qtype = RAWR_REDIR_QTYPE;
    flowupd.fwding.qid = ctx.role() ==  hal::FLOW_ROLE_INITIATOR ?
                         pfi->qid1 : pfi->qid2;
    return ctx.update_flow(flowupd);
}


fte::pipeline_action_t
app_redir_miss_exec(fte::ctx_t& ctx)
{
    proxy_flow_info_t       *pfi;
    hal_ret_t               ret = HAL_RET_OK;
    flow_key_t              flow_key = ctx.key();

    ret = app_redir_init();
    if (ret != HAL_RET_OK) {
        goto error;
    }

    /*
     * Ignore direction. Always set it to 0
     */
    flow_key.dir = 0;

    /*
     * get the flow info for the Application (L7) Redirect proxy service 
     */
    pfi = proxy_get_flow_info(types::PROXY_TYPE_APP_REDIR,
                              &flow_key);
    if (pfi) {

        /*
         * Update flow
         */
        ret = app_redir_flow_fwding_update(ctx, flow_key, pfi);
        if (ret != HAL_RET_OK) {
            goto error;
        }
    } else {
        HAL_TRACE_DEBUG("app_redir is not enabled for the flow");
    } 

    return fte::PIPELINE_CONTINUE;

error:
    ctx.set_feature_status(ret);
    return fte::PIPELINE_CONTINUE;
}

fte::pipeline_action_t
app_redir_exec(fte::ctx_t& ctx)
{
    hal_ret_t   ret;

    ret = app_redir_init();
    if (ret != HAL_RET_OK) {
        goto error;
    }

    ret = app_redir_pkt_send(ctx);
    if (ret != HAL_RET_OK) {
        goto error;
    }
    return fte::PIPELINE_CONTINUE;

error:
    ctx.set_feature_status(ret);
    return fte::PIPELINE_CONTINUE;
}

} // namespace hal
} // namespace net
