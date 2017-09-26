#pragma once

#include <base.h>
#include <fte_ctx.hpp>

namespace fte {

// FTE features
typedef uint16_t feature_id_t;

const feature_id_t FTE_FEATURE_NULL          = 0;
const feature_id_t FTE_FEATURE_STAGE_MARKER  = 1;
const feature_id_t FTE_FEATURE_FWDING        = 2;
const feature_id_t FTE_FEATURE_TUNNEL        = 3;
const feature_id_t FTE_FEATURE_DFW           = 4;
const feature_id_t FTE_FEATURE_LB            = 5;
const feature_id_t FTE_FEATURE_TCP_PROXY     = 6;
const feature_id_t FTE_FEATURE_QOS           = 7;
const feature_id_t FTE_FEATURE_DOL_TEST_ONLY = 8;

// FTE pipeline progress action codes
//  continue - continue the pipeline
//  end - end processing
//  restart - restart the processing from the begining
//            (may select a diffrent pipeline)
#define FTE_PIPELINE_ACTION_ENTRIES(ENTRY)       \
    ENTRY(PIPELINE_CONTINUE, 0, "continue")      \
    ENTRY(PIPELINE_END,      1, "end")           \
    ENTRY(PIPELINE_RESTART,  2, "restart")

DEFINE_ENUM(pipeline_action_t, FTE_PIPELINE_ACTION_ENTRIES)
#undef FTE_PIPELINE_ACTION_ENTRIES

typedef std::function<pipeline_action_t(ctx_t &ctx)> exec_handler_t;
typedef std::function<void(ctx_t &ctx)> commit_handler_t;
hal_ret_t register_feature(const feature_id_t& fid,
                           const std::string& name,
                           const exec_handler_t &exec_handler,
                           const commit_handler_t &commit_handler = nullptr);

//  FTE Pipeline 
hal_ret_t register_pipeline(const std::string& name, const lifqid_t &lifq,
                            feature_id_t features_outbound[],
                            uint16_t num_features_outbound,
                            feature_id_t features_inbound[] = {},
                            uint16_t num_features_inbound = 0,
                            const lifqid_t &lifq_mask = lifqid_t{0x7FF, 0x7, 0xFFFFFF});

// Selects a pipeline and invokes features in that pipeline
// Doesn't update P4 tables and write packet to send queue
hal_ret_t execute_pipeline(ctx_t &ctx);

// for unit test code only
void unregister_features_and_pipelines();

// grpc session request
hal_ret_t session_create(SessionSpec& spec, SessionResponse *rsp);

// FTE pkt loop (infinite loop)
void pkt_loop(uint8_t fte_id);


} // namespace fte
