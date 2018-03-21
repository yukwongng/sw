//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//
// sw PHV injection
//-----------------------------------------------------------------------------

#include "nic/include/hal_lock.hpp"
#include "nic/include/pd_api.hpp"
#include "nic/hal/src/proxy.hpp"
#include "nic/hal/pd/iris/swphv_pd.hpp"
#include "nic/include/pd.hpp"
#include "nic/hal/pd/iris/hal_state_pd.hpp"
#include "nic/hal/pd/p4pd_api.hpp"
#include "nic/hal/pd/asicpd/asic_pd_common.hpp"
#include "nic/gen/tcp_proxy_txdma/include/ingress_phv.h"

namespace hal {
namespace pd {

// Iris PHV size in bytes
#define PD_IRIS_PHV_SIZE (4096 / 8)

// pd_swphv_inject injects software PHV into a pipeline 
hal_ret_t
pd_swphv_inject (pd_swphv_inject_args_t *args)
{
    uint8_t data[PD_IRIS_PHV_SIZE];
    hal_ret_t ret = HAL_RET_OK;
    bzero(data, PD_IRIS_PHV_SIZE);
    uint32_t   lif_id = 0;

    // switch based on pipeline type
    switch(args->type) {
    case PD_SWPHV_TYPE_RXDMA:
        lif_id = SERVICE_LIF_TCP_PROXY;
	break;
    case PD_SWPHV_TYPE_TXDMA:
        lif_id = SERVICE_LIF_TCP_PROXY;
	break;
    case PD_SWPHV_TYPE_INGRESS:
    case PD_SWPHV_TYPE_EGRESS:
	break;
    }

// define PHV_GEN_HDR when NCC generated header files are ready
#ifndef PHV_GEN_HDR
    sw_phv_intr_t *phv = (sw_phv_intr_t *)data;

    // initialize PHV intrinsic fields
    phv->p4_intr_global.drop = 1;
    phv->p4_intr_global.bypass = 1;
    phv->p4_intr.no_data = 1;
    phv->p4_intr_global.lif = lif_id;
    phv->dma.p4_intr_txdma.intr_qid = 0;
    phv->dma.p4_intr_txdma.intr_qtype = 0;

    HAL_TRACE_DEBUG("pd_swphv_inject: Injecting Software PHV type {}. size: {}", args->type, sizeof(sw_phv_intr_t));

    // Inject the phv
    ret = asicpd_sw_phv_inject((asicpd_swphv_type_t)args->type, 0, 0, 1, data);
#else
    ingress_phv_t *phv = (ingress_phv_t *)data;

    // initialize PHV intrinsic fields
    phv->p4_intr_global_drop = 1;
    phv->p4_intr_global_bypass = 0;
    phv->p4_intr_no_data = 1;
    phv->p4_intr_global_lif = lif_id;
    phv->p4_txdma_intr_qid = 0;
    phv->p4_txdma_intr_qtype = 0;

    HAL_TRACE_DEBUG("pd_swphv_inject: Injecting Software PHV type {}. size of phv: {}", args->type, sizeof(ingress_phv_t));

    // Inject the phv
    ret = asicpd_sw_phv_inject((asicpd_swphv_type_t)args->type, 0, 0, 1, data);
#endif
    return ret;
}

// pd_swphv_get_state
// get the current state of SW phv
hal_ret_t
pd_swphv_get_state (pd_swphv_get_state_args_t *state)
{
    asicpd_sw_phv_state_t    hw_state;
    hal_ret_t   ret = HAL_RET_OK;

    HAL_TRACE_DEBUG("pd_swphv_get_state: getting Software PHV state for type {}", state->type);

    // get it from capri PD
    ret = asicpd_sw_phv_get((asicpd_swphv_type_t)state->type, 0, &hw_state);
    if (ret != HAL_RET_OK) {
	return ret;
    }

    // translate state
    state->enabled           = hw_state.enabled;
    state->done              = hw_state.done;
    state->current_cntr      = hw_state.current_cntr;
    state->no_data_cntr      = hw_state.no_data_cntr;
    state->drop_no_data_cntr = hw_state.drop_no_data_cntr;

    return HAL_RET_OK;
}

}    // namespace pd
}    // namespace hal
