/**
 * Copyright (c) 2019 Pensando Systems, Inc.
 *
 * @file    debug.cc
 *
 * @brief   This file has helper routines for troubleshooting the system
 */

#include "nic/sdk/linkmgr/port_mac.hpp"
#include "nic/sdk/linkmgr/linkmgr.hpp"
#include "nic/sdk/include/sdk/fd.hpp"
#include "nic/sdk/platform/asicerror/interrupts.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/framework/impl_base.hpp"
#include "nic/apollo/framework/api_engine.hpp"
#include "nic/apollo/api/port.hpp"
#include "nic/apollo/api/debug.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/apollo/core/core.hpp"

namespace debug {

/**
 * @brief        set clock frequency
 * @param[in]    freq clock frequency to be set
 * @return       #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_clock_frequency_update (pds_clock_freq_t freq)
{
    return impl_base::asic_impl()->set_frequency(freq);
}

/**
 * @brief        set arm clock frequency
 * @param[in]    freq clock frequency to be set
 * @return       #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_arm_clock_frequency_update (pds_clock_freq_t freq)
{
    return impl_base::asic_impl()->set_arm_frequency(freq);
}

/**
 * @brief        get system temperature
 * @param[out]   Temperate to be read
 * @return       #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_get_system_temperature (pds_system_temperature_t *temp)
{
    return impl_base::asic_impl()->get_system_temperature(temp);
}

/**
 * @brief        get system power
 * @param[out]   Power to be read
 * @return       #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_get_system_power (pds_system_power_t *pow)
{
    return impl_base::asic_impl()->get_system_power(pow);
}

/**
 * @brief       Gets API stats for different tables
 * @param[in]   CB to fill API stats & ctxt
 * @return      #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_table_stats_get (debug::table_stats_get_cb_t cb, void *ctxt)
{
    return impl_base::pipeline_impl()->table_stats(cb, ctxt);
}

sdk_ret_t
dump_interrupts (int fd)
{
    dprintf(fd, "%s\n", std::string(80, '-').c_str());
    dprintf(fd, "%-50s%-10s%-9s%-11s\n",
            "Name", "Count", "Severity", "Description");
    dprintf(fd, "%s\n", std::string(80, '-').c_str());
    ::walk_interrupts(intr_dump_cb, &fd);
    return SDK_RET_OK;
}

sdk_ret_t
clear_interrupts (int fd)
{
    ::clear_interrupts();
    dprintf(fd, "Interrupts cleared\n");
    return SDK_RET_OK;
}

bool
state_base_counters_dump (void *obj, void *ctxt)
{
    api::state_walk_ctxt_t *walk_ctxt = (api::state_walk_ctxt_t *)obj;
    api::state_base *state = (api::state_base *)walk_ctxt->state;
    auto ctrs = state->counters();
    int fd = *(int *)ctxt;

    dprintf(fd, "%-25s%-10u%-10u%-10u%-10u%-10u%-10u%-10u\n",
            walk_ctxt->obj_state.c_str(),
            ctrs.num_elems, ctrs.insert_ok, ctrs.insert_err,
            ctrs.remove_ok, ctrs.remove_err,
            ctrs.update_ok, ctrs.update_err);
    return false;
}

sdk_ret_t
dump_state_base_stats (int fd)
{
    dprintf(fd, "%s\n", std::string(95, '-').c_str());
    dprintf(fd, "%-25s%-10s%-10s%-10s%-10s%-10s%-10s%-10s\n",
            "Store", "NumElem", "InsertOK", "InsertErr",
            "RemoveOK", "RemoveErr", "UpdateOK", "UpdateErr");
    dprintf(fd, "%s\n", std::string(95, '-').c_str());
    api::g_pds_state.walk(state_base_counters_dump, (void *)&fd);
    return SDK_RET_OK;
}

/**
 * @brief       Handles command based on ctxt
 * @param[in]   ctxt  Context for CLI handler
 * @return      #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_handle_cmd (cmd_ctxt_t *ctxt)
{
    switch (ctxt->cmd) {
    case CLI_CMD_MAPPING_DUMP:
    case CLI_CMD_NACL_DUMP:
        return impl_base::pipeline_impl()->handle_cmd(ctxt);
    case CLI_CMD_INTR_DUMP:
        dump_interrupts(ctxt->fd);
        break;
    case CLI_CMD_INTR_CLEAR:
        clear_interrupts(ctxt->fd);
        break;
    case CLI_CMD_API_ENGINE_STATS_DUMP:
        api::api_engine_get()->dump_api_counters(ctxt->fd);
        break;
    case CLI_CMD_STORE_STATS_DUMP:
        dump_state_base_stats(ctxt->fd);
        break;
    default:
        return SDK_RET_INVALID_ARG;
    }
    return SDK_RET_OK;
}

/**
 * @brief       Setup LLC
 * @param[in]   llc_counters_t with type set
 * @return      #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_llc_setup (llc_counters_t *llc_args)
{
    return impl_base::asic_impl()->llc_setup(llc_args);
}

/**
 * @brief       LLC Stats Get
 * @param[out]  llc_counters_t to be filled
 * @return      #SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_llc_get (llc_counters_t *llc_args)
{
    return impl_base::asic_impl()->llc_get(llc_args);
}

sdk_ret_t
pds_pb_stats_get (debug::pb_stats_get_cb_t cb, void *ctxt)
{
    return impl_base::asic_impl()->pb_stats(cb, ctxt);
}

sdk_ret_t
pds_meter_stats_get (debug::meter_stats_get_cb_t cb, uint32_t lowidx,
                     uint32_t highidx, void *ctxt)
{
    return impl_base::pipeline_impl()->meter_stats(cb, lowidx, highidx, ctxt);
}

sdk_ret_t
pds_session_stats_get (debug::session_stats_get_cb_t cb, uint32_t lowidx,
                       uint32_t highidx, void *ctxt)
{
    return impl_base::pipeline_impl()->session_stats(cb, lowidx, highidx, ctxt);
}

sdk_ret_t
pds_fte_api_stats_get (void)
{
    return SDK_RET_OK;
}

sdk_ret_t
pds_fte_table_stats_get (void)
{
    return SDK_RET_OK;
}

sdk_ret_t
pds_fte_api_stats_clear (void)
{
    return SDK_RET_OK;
}

sdk_ret_t
pds_fte_table_stats_clear (void)
{
    return SDK_RET_OK;
}

sdk_ret_t
pds_session_get (debug::session_get_cb_t cb, void *ctxt)
{
    return impl_base::pipeline_impl()->session(cb, ctxt);
}

sdk_ret_t
pds_flow_get (debug::flow_get_cb_t cb, void *ctxt)
{
    return impl_base::pipeline_impl()->flow(cb, ctxt);
}

sdk_ret_t
pds_session_clear (uint32_t idx)
{
    return impl_base::pipeline_impl()->session_clear(idx);
}

void
flow_clear_resp_cb(sdk::ipc::ipc_msg_ptr msg, const void *ret)
{
    *(sdk_ret_t *)ret = *(sdk_ret_t *)msg->data();
}

sdk_ret_t
pds_flow_clear (pds_flow_key_t key)
{
    sdk_ret_t ret;
    pds_msg_t request;

    // send an IPC msg to VPP
    request.id = PDS_CMD_MSG_FLOW_CLEAR;
    memset(&request.cmd_msg.flow_clear, 0, sizeof(pds_flow_clear_cmd_msg_t));
    request.cmd_msg.flow_clear.key = key;

    // send a msg to VPP to clear the flows
    sdk::ipc::request(PDS_IPC_ID_VPP, PDS_MSG_TYPE_CMD, &request,
                      sizeof(pds_msg_t), flow_clear_resp_cb,
                      &ret);

    return ret;
}

/**
 * @brief    start AACS server
 * @param[in]    aacs_server_port     AACS server port
 * @return   SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
start_aacs_server (uint32_t aacs_server_port)
{
    sdk::linkmgr::start_aacs_server(aacs_server_port);
    return SDK_RET_OK;
}

/**
 * @brief    stop AACS server
 * @return   SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
stop_aacs_server (void)
{
    sdk::linkmgr::stop_aacs_server();
    return SDK_RET_OK;
}

/**
 * @brief    get slab information
 * @param[in]   cb   callback function to be called
 * @paraam[in]  ctxt context for callback function
 * @return   SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
pds_slab_get (api::state_walk_cb_t cb, void *ctxt)
{
    api::g_pds_state.slab_walk(cb, ctxt);
    // TODO: Revisit logic to access pds_impl_state through pds_state
    impl_base::pipeline_impl()->impl_state_slab_walk(cb, ctxt);
    return SDK_RET_OK;
}

}    // namespace debug
