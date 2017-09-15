#include <sys/queue.h>

#include "fte.hpp"
#include "fte_flow.hpp"
#include <defines.h>

namespace fte {

// FTE features
typedef struct feature_s feature_t;
struct feature_s {
    STAILQ_ENTRY(feature_s) entries;
    feature_id_t id;
    std::string name;
    exec_handler_t exec_handler;
    commit_handler_t commit_handler;
};

static STAILQ_HEAD(feature_list_s_, feature_s) g_feature_list_ =
    STAILQ_HEAD_INITIALIZER(g_feature_list_);

static inline feature_t *feature_alloc_()
{
    return (feature_t *)HAL_CALLOC(feature_t, sizeof(feature_t));
}

static inline void feature_add_(feature_t *feature)
{
    STAILQ_INSERT_TAIL(&g_feature_list_, feature, entries);
}

static inline feature_t *feature_lookup_(feature_id_t fid)
{
    feature_t *feature;
    STAILQ_FOREACH(feature, &g_feature_list_, entries) {
        if (feature->id == fid) {
            return feature;
        }
    }
    return nullptr;
}

hal_ret_t register_feature(const feature_id_t& fid, const std::string& name,
                           const exec_handler_t &exec_handler,
                           const commit_handler_t &commit_handler)
{
    feature_t *feature;

    HAL_TRACE_DEBUG("fte::{}: id={}, name={}", __FUNCTION__, fid, name);

    if (!exec_handler) {
        HAL_TRACE_ERR("fte: skipping invalid feature id={} name={} - null exec_handler",
                      fid, name);
        return HAL_RET_INVALID_ARG;
    }

    if ((feature = feature_lookup_(fid)) != nullptr) {
        HAL_TRACE_ERR("fte: skipping duplicate feature id={} name={} old-name={}",
                      fid, name, feature->name);
        return HAL_RET_DUP_INS_FAIL;
    }

    feature = feature_alloc_();
    feature->id = fid;
    feature->name = name;
    feature->exec_handler = exec_handler;
    feature->commit_handler = commit_handler;

    feature_add_(feature);

    return HAL_RET_OK;
}

// FTE pipelines
// A list is used instead of a map for pipelines. It is most likely faster
// to scan a list as the no.o pipelines will be very small (single digits).
// Also we keep the most frequently used pipelines in the front (flow miss
// pipeline will be the first, with minimal lookup cost)
typedef struct pipeline_s pipeline_t;
struct pipeline_s
{
    STAILQ_ENTRY(pipeline_s) entries;
    lifqid_t lifq;
    lifqid_t lifq_mask;
    std::string name;
    uint16_t num_features_outbound;
    uint16_t num_features_inbound;
    feature_t *features[0];
};

static STAILQ_HEAD(pipeline_list_s_, pipeline_s) g_pipeline_list_ =
    STAILQ_HEAD_INITIALIZER(g_pipeline_list_);

static inline pipeline_t *
pipeline_alloc_(uint16_t num_features)
{
    return (pipeline_t *)HAL_CALLOC(pipeline_t, sizeof(pipeline_t) +
                                    num_features * sizeof(feature_t *));
}

static inline void
pipeline_add_(pipeline_t *pipeline)
{
    STAILQ_INSERT_TAIL(&g_pipeline_list_, pipeline, entries);
}

static inline pipeline_t *
pipeline_lookup_(const lifqid_t& lifq)
{
    pipeline_t *pipeline;

    STAILQ_FOREACH(pipeline, &g_pipeline_list_, entries) {
        if (pipeline->lifq.lif == (lifq.lif & pipeline->lifq_mask.lif) &&
            pipeline->lifq.qtype == (lifq.qtype & pipeline->lifq_mask.qtype) &&
            pipeline->lifq.qid == (lifq.qid & pipeline->lifq_mask.qid)) {
            return pipeline;
        }
    }
    return nullptr;
}

static inline pipeline_action_t
pipeline_invoke_exec_(pipeline_t *pipeline, ctx_t &ctx, uint8_t start, uint8_t end)
{
    pipeline_action_t rc;
    for (int i = start; i < end; i++) {
        feature_t *feature = pipeline->features[i];

        ctx.set_feature_name(feature->name.c_str());
        ctx.set_feature_status(HAL_RET_OK);
        rc = feature->exec_handler(ctx);
        HAL_TRACE_DEBUG("fte:exec_handler feature={} pipeline={} action={}", feature->name,
                        pipeline->name, rc);

        if (rc != PIPELINE_CONTINUE) {
            break;
        }
    }

    return rc;
}

hal_ret_t
register_pipeline(const std::string& name, const lifqid_t& lifq,
                  feature_id_t features_outbound[], uint16_t num_features_outbound,
                  feature_id_t features_inbound[], uint16_t num_features_inbound,
                  const lifqid_t& lifq_mask)
{
    pipeline_t *pipeline;

    HAL_TRACE_DEBUG("fte::{}: name={} lifq={}", __FUNCTION__, name, lifq);
    if ((pipeline = pipeline_lookup_(lifq)) != nullptr) {
        HAL_TRACE_ERR("fte: skipping duplicate pipline {} lifq={} old-name={}",
                      name, lifq, pipeline->name);
        return HAL_RET_DUP_INS_FAIL;
    }

    pipeline = pipeline_alloc_(num_features_outbound+num_features_inbound);
    pipeline->lifq = lifq;
    pipeline->lifq_mask = lifq_mask;
    pipeline->name = name;

    // Initialize feature blocks
    pipeline->num_features_outbound =num_features_outbound;
    for (int i = 0; i < num_features_outbound; i++) {
        feature_t *feature = feature_lookup_(features_outbound[i]);
        if (!feature) {
            HAL_TRACE_ERR("fte: unknown feature-id {} in outbound pipeline {} - skipping",
                          features_outbound[i], name);
            HAL_FREE(pipeline_t, pipeline);
            return HAL_RET_INVALID_ARG;
        }
        HAL_TRACE_DEBUG("fte: outbound pipeline feature {}/{}", name, feature->name);
        pipeline->features[i] = feature;
    }

    pipeline->num_features_inbound =num_features_inbound;
    for (int i = 0; i < num_features_inbound; i++) {
        feature_t *feature = feature_lookup_(features_inbound[i]);
        if (!feature) {
            HAL_TRACE_ERR("fte: unknown feature-id {} in inbound pipeline {} - skipping",
                          features_inbound[i], name);
            HAL_FREE(pipeline_t, pipeline);
            return HAL_RET_INVALID_ARG;
        }
        HAL_TRACE_DEBUG("fte: inbound pipeline feature {}/{}", name, feature->name);
        pipeline->features[i+num_features_outbound] = feature;
    }

    // Add to global pipline list
    pipeline_add_(pipeline);

    return HAL_RET_OK;
}

static inline hal_ret_t
pipeline_execute_(ctx_t &ctx)
{
    uint8_t iflow_start, iflow_end, rflow_start, rflow_end;
    pipeline_action_t rc;

    do {
        pipeline_t *pipeline = pipeline_lookup_(ctx.arm_lifq());
        if (!pipeline) {
            HAL_TRACE_ERR("fte: pipeline not registered for lifq {} - ignoring packet",
                          ctx.arm_lifq());
            return HAL_RET_INVALID_ARG;
        }
        
        HAL_TRACE_DEBUG("fte: executing pipeline {} lifq={} dir={}",
                        pipeline->name, pipeline->lifq, ctx.direction());

        iflow_start = 0;
        iflow_end = rflow_start = pipeline->num_features_outbound;
        rflow_end = pipeline->num_features_outbound + pipeline->num_features_inbound;
        // Flip the feature list for network initiated
        if (pipeline->num_features_inbound > 0 && ctx.direction() == FLOW_DIR_FROM_UPLINK) {
            rflow_start = 0;
            rflow_end = iflow_start = pipeline->num_features_outbound;
            iflow_end = pipeline->num_features_outbound + pipeline->num_features_inbound;
        }
        
        // Invoke all initiator feature handlers 
        ctx.set_role(hal::FLOW_ROLE_INITIATOR);
        rc = pipeline_invoke_exec_(pipeline, ctx, iflow_start, iflow_end);
        if (rc == PIPELINE_CONTINUE && ctx.valid_rflow()) {
            ctx.set_role(hal::FLOW_ROLE_RESPONDER);
            rc = pipeline_invoke_exec_(pipeline, ctx, rflow_start, rflow_end);
        }
    } while (rc == PIPELINE_RESTART);

    return ctx.feature_status();
}

hal_ret_t execute_pipeline(ctx_t &ctx)
{
    hal_ret_t rc = pipeline_execute_(ctx);

    // TODO - invoke commit handlers
    // need to track commit_handlers, we might have executed
    // multiple pipelines and skipped some features.

    return rc;
}

// for unit test code only
void unregister_features_and_pipelines() {
    while (!STAILQ_EMPTY(&g_pipeline_list_)) {
        pipeline_t *pipeline = STAILQ_FIRST(&g_pipeline_list_);
        STAILQ_REMOVE_HEAD(&g_pipeline_list_, entries);
        HAL_FREE(pipeline_t, pipeline);
    }
    while (!STAILQ_EMPTY(&g_feature_list_)) {
        feature_t *feature = STAILQ_FIRST(&g_feature_list_);
        STAILQ_REMOVE_HEAD(&g_feature_list_, entries);
        HAL_FREE(feature_t, feature);
    }
}

// Process grpc session_create
hal_ret_t
session_create (SessionSpec& spec, SessionResponse *rsp)
{
    hal_ret_t ret;
    ctx_t ctx = {};
    flow_t iflow[ctx_t::MAX_STAGES], rflow[ctx_t::MAX_STAGES];

    HAL_TRACE_DEBUG("--------------------- API Start ------------------------");
    HAL_TRACE_DEBUG("fte::{}: Session id {} Create in Tenant id {}", __FUNCTION__, 
                    spec.session_id(), spec.meta().tenant_id());

    //Init context
    ret = ctx.init(&spec, rsp,  iflow, rflow);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("fte: failied to init context, ret={}", ret);
        goto end;
    }

    // execute the pipeline
    ret = execute_pipeline(ctx);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("fte: failied to execute pipeline, ret={}", ret);
        goto end;
    }

    // update GFT
    ret = ctx.update_gft();
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("fte: failied to updated gft, ret={}", ret);
        goto end;
    }

 end:
    //update api status
    switch (ret) {
    case HAL_RET_OK:
        rsp->set_api_status(types::API_STATUS_OK);
        break;
    case HAL_RET_HW_PROG_ERR:
        rsp->set_api_status(types::API_STATUS_HW_PROG_ERR);
        break;
    case HAL_RET_TABLE_FULL:
    case HAL_RET_OTCAM_FULL:
        rsp->set_api_status(types::API_STATUS_OUT_OF_RESOURCE);
        break;
    case HAL_RET_OOM:
        rsp->set_api_status(types::API_STATUS_OUT_OF_MEM);
        break;
    case HAL_RET_INVALID_ARG:
        rsp->set_api_status(types::API_STATUS_INVALID_ARG);
        break;
    case HAL_RET_TENANT_NOT_FOUND:
        rsp->set_api_status(types::API_STATUS_TENANT_NOT_FOUND);
        break;
    case HAL_RET_L2SEG_NOT_FOUND:
        rsp->set_api_status(types::API_STATUS_L2_SEGMENT_NOT_FOUND);
        break;
    case HAL_RET_IF_NOT_FOUND:
        rsp->set_api_status(types::API_STATUS_INTERFACE_NOT_FOUND);
        break;
    case HAL_RET_SECURITY_PROFILE_NOT_FOUND:
        rsp->set_api_status(types::API_STATUS_NWSEC_PROFILE_NOT_FOUND);
        break;
    case HAL_RET_POLICER_NOT_FOUND:
        rsp->set_api_status(types::API_STATUS_POLICER_NOT_FOUND);
        break;
    case HAL_RET_HANDLE_INVALID:
        rsp->set_api_status(types::API_STATUS_HANDLE_INVALID);
        break;
    default:
        rsp->set_api_status(types::API_STATUS_ERR);
        break;
    }


    HAL_TRACE_DEBUG("----------------------- API End ------------------------");
    return ret;
}

// FTE main pkt loop
void
pkt_loop(arm_rx_t rx, arm_tx_t tx)
{
    hal_ret_t ret;
    phv_t *phv;
    uint8_t *pkt;
    size_t pkt_len;
    ctx_t ctx;
    flow_t iflow[ctx_t::MAX_STAGES], rflow[ctx_t::MAX_STAGES];

    while(true) {
        // read the packet
        ret = rx(&phv, &pkt, &pkt_len);
        if (ret != HAL_RET_OK) {
            HAL_TRACE_ERR("fte: arm rx failed, ret={}", ret);
            continue;
        }

        // Init ctx_t
        ret = ctx.init(phv, pkt, pkt_len, iflow, rflow);
        if (ret != HAL_RET_OK) {
            HAL_TRACE_ERR("fte: failied to init context, ret={}", ret);
            continue;
        }

        // execute the pipeline
        ret = execute_pipeline(ctx);
        if (ret != HAL_RET_OK) {
            HAL_TRACE_ERR("fte: failied to execute pipeline, ret={}", ret);
            continue;
        }

        // update GFT
        ret = ctx.update_gft();
        if (ret != HAL_RET_OK) {
            HAL_TRACE_ERR("fte: failied to updated gft, ret={}", ret);
            continue;
        }

        // write the packet
        if (ctx.pkt()) {
            ret = tx(ctx.phv(), ctx.pkt(), ctx.pkt_len());
            if (ret != HAL_RET_OK) {
                HAL_TRACE_ERR("fte: failied to transmit pkt, ret={}", ret);
                continue;
            }
        }
    }
}
} // namespace fte
