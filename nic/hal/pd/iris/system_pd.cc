#include "nic/include/hal_lock.hpp"
#include "nic/p4/nw/include/table_sizes.h"
#include "nic/hal/pd/utils/tcam/tcam.hpp"
#include "nic/hal/pd/iris/hal_state_pd.hpp"
#include "nic/hal/pd/iris/system_pd.hpp"
#include "nic/p4/nw/include/defines.h"
#include "nic/hal/pd/p4pd_api.hpp"
#include "nic/hal/pd/capri/capri_hbm.hpp"
#include "nic/hal/pd/iris/p4plus_pd_api.h"

namespace hal {
namespace pd {

hal_ret_t
pd_system_populate_drop_stats(DropStatsEntry *stats_entry, 
                              uint8_t idx);

hal_ret_t
pd_drop_stats_get(pd_system_args_t *pd_sys_args)
{
    hal_ret_t               ret = HAL_RET_OK;
    SystemResponse          *rsp = pd_sys_args->rsp;
    DropStatsEntry          *stats_entry = NULL;

    HAL_TRACE_DEBUG("PD-System: Querying drop stats");

    for (int i = 0; i < DROP_STATS_TABLE_SIZE; i++) {
        stats_entry = rsp->mutable_stats()->mutable_drop_stats()->
            add_drop_entries();
        pd_system_populate_drop_stats(stats_entry, i);
    }

    return ret;
}

inline hbm_addr_t
hbm_get_addr_for_stat_index(p4pd_table_id table_id,
                            uint8_t idx)
{
    hbm_addr_t  stats_base_addr;
    p4pd_table_properties_t  tbl_ctx;

    stats_base_addr =  get_start_offset(JP4_ATOMIC_STATS);
    stats_base_addr &= 0x7FFFFFFF;

    switch (table_id) {
    case P4TBL_ID_INGRESS_TX_STATS:
        p4pd_table_properties_get(P4TBL_ID_TX_STATS, &tbl_ctx);
        stats_base_addr += ((tbl_ctx.tabledepth << 5) + (tbl_ctx.tabledepth << 4) + (tbl_ctx.tabledepth << 3));
        // Fall through
    case P4TBL_ID_TX_STATS:
        p4pd_table_properties_get(P4TBL_ID_DROP_STATS, &tbl_ctx);
        stats_base_addr += (tbl_ctx.tabledepth << 3);
        // Fall through
    case P4TBL_ID_DROP_STATS:
        p4pd_table_properties_get(P4TBL_ID_COPP_ACTION, &tbl_ctx);
        stats_base_addr += (tbl_ctx.tabledepth << 5);
        // Fall through
    case P4TBL_ID_COPP_ACTION:
        p4pd_table_properties_get(P4TBL_ID_RX_POLICER_ACTION, &tbl_ctx);
        stats_base_addr += (tbl_ctx.tabledepth << 5);
        // Fall through
    case P4TBL_ID_RX_POLICER_ACTION:
        p4pd_table_properties_get(P4TBL_ID_FLOW_STATS, &tbl_ctx);
        stats_base_addr += (tbl_ctx.tabledepth << 5);
        // Fall through
    case P4TBL_ID_FLOW_STATS:
    default:
        break;
    }

    switch (table_id) {
    case P4TBL_ID_FLOW_STATS:
    case P4TBL_ID_RX_POLICER_ACTION:
    case P4TBL_ID_COPP_ACTION:
        stats_base_addr += (idx << 5);
        break;
    case P4TBL_ID_DROP_STATS:
    case P4TBL_ID_INGRESS_TX_STATS:
        stats_base_addr += (idx << 3);
        break;
    case P4TBL_ID_TX_STATS:
        stats_base_addr += ((idx << 5) + (idx << 4) + (idx << 3));
        break;
    default:
        break;
    }

    return stats_base_addr;
}

hal_ret_t
pd_system_populate_drop_stats(DropStatsEntry *stats_entry, 
                            uint8_t idx)
{
    hal_ret_t               ret = HAL_RET_OK;
    Tcam                    *tcam;
    drop_stats_swkey         key = { 0 };
    drop_stats_swkey_mask    key_mask = { 0 };
    drop_stats_actiondata    data = { 0 };
    uint64_t                 hbm_counter = 0, 
                             hbm_counter1  = 0, hbm_counter2 = 0;
    hbm_addr_t               stats_base_addr;

    // Find the base addr for specific Drop Stats index
    stats_base_addr = hbm_get_addr_for_stat_index(P4TBL_ID_DROP_STATS, idx);

    tcam = g_hal_state_pd->tcam_table(P4TBL_ID_DROP_STATS);
    HAL_ASSERT(tcam != NULL);

    // Retrieve from SW to get the stats idx
    ret = tcam->retrieve(idx, &key, &key_mask, &data);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Unable to retrieve stats idx for: {}",
                idx);
        goto end;
    }

	// Read count from HBM with stats idx.
    if (!p4plus_hbm_read(stats_base_addr, (uint8_t *)&hbm_counter1,
                sizeof(hbm_counter1))) {
        HAL_TRACE_ERR("Unable to retrieve HBM stats for: {}", idx);
        goto end;
    }

    // Read from drop stats table
    ret = tcam->retrieve_from_hw(idx, &key, &key_mask, &data);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Unable to retrieve drop stats for entry: {}",
                idx);
        goto end;
    }

    if (key.entry_inactive_drop_stats) {
        goto end;
    }

	// Read count from HBM with stats idx.
    if (!p4plus_hbm_read(stats_base_addr, (uint8_t *)&hbm_counter2,
                sizeof(hbm_counter2))) {
    	HAL_TRACE_ERR("Unable to retrieve HBM stats for entry: {}", idx);
        goto end;  
    }

    // TODO: Do we need a check. Why cant we read drop_stats, hbm 
    if (hbm_counter2 > hbm_counter1) {
        hbm_counter = hbm_counter2;
    } else {
        hbm_counter = hbm_counter1;
    }

    pd_system_decode(&key, &key_mask, &data, stats_entry, hbm_counter);

end:
    return ret;
}

hal_ret_t
pd_system_decode(drop_stats_swkey *key, drop_stats_swkey_mask *key_mask,
        drop_stats_actiondata *data, DropStatsEntry *stats_entry,
        uint64_t hbm_counter)
{
    hal_ret_t   ret = HAL_RET_OK;
    uint64_t drop_reason, drop_reason_mask;

    memcpy(&drop_reason, key->control_metadata_drop_reason,
           sizeof(key->control_metadata_drop_reason));
    memcpy(&drop_reason_mask, key_mask->control_metadata_drop_reason_mask,
           sizeof(key_mask->control_metadata_drop_reason_mask));
    drop_reason &= drop_reason_mask;
    memcpy(key->control_metadata_drop_reason, &drop_reason,
           sizeof(key->control_metadata_drop_reason));

    stats_entry->set_drop_input_mapping(
            drop_reason & (1 << DROP_INPUT_MAPPING));
    stats_entry->set_drop_input_mapping_dejavu(
            drop_reason & (1 << DROP_INPUT_MAPPING_DEJAVU));
    stats_entry->set_drop_flow_hit(
            drop_reason & (1 << DROP_FLOW_HIT));
    stats_entry->set_drop_flow_miss(
            drop_reason & (1 << DROP_FLOW_MISS));
    stats_entry->set_drop_ipsg(
            drop_reason & (1 << DROP_IPSG));
    stats_entry->set_drop_ingress_policer(
            drop_reason & (1 << DROP_INGRESS_POLICER));
    stats_entry->set_drop_egress_policer(
            drop_reason & (1 << DROP_RX_POLICER));
    stats_entry->set_drop_nacl(
            drop_reason & (1 << DROP_NACL));
    stats_entry->set_drop_malformed_pkt(
            drop_reason & (1 << DROP_MALFORMED_PKT));
    stats_entry->set_drop_ping_of_death(
            drop_reason & (1 << DROP_PING_OF_DEATH));
    stats_entry->set_drop_fragment_too_small(
            drop_reason & (1 << DROP_FRAGMENT_TOO_SMALL));
    stats_entry->set_drop_ip_normalization(
            drop_reason & (1 << DROP_IP_NORMALIZATION));
    stats_entry->set_drop_tcp_normalization(
            drop_reason & (1 << DROP_TCP_NORMALIZATION));
    stats_entry->set_drop_tcp_non_syn_first_pkt(
            drop_reason & (1 << DROP_TCP_NON_SYN_FIRST_PKT));
    stats_entry->set_drop_icmp_normalization(
            drop_reason & (1 << DROP_ICMP_NORMALIZATION));
    stats_entry->set_drop_icmp_src_quench_msg(
            drop_reason & (1 << DROP_ICMP_SRC_QUENCH_MSG));
    stats_entry->set_drop_icmp_redirect_msg(
            drop_reason & (1 << DROP_ICMP_REDIRECT_MSG));
    stats_entry->set_drop_icmp_info_req_msg(
            drop_reason & (1 << DROP_ICMP_INFO_REQ_MSG));
    stats_entry->set_drop_icmp_addr_req_msg(
            drop_reason & (1 << DROP_ICMP_ADDR_REQ_MSG));
    stats_entry->set_drop_icmp_traceroute_msg(
            drop_reason & (1 << DROP_ICMP_TRACEROUTE_MSG));
    stats_entry->set_drop_icmp_rsvd_type_msg(
            drop_reason & (1 << DROP_ICMP_RSVD_TYPE_MSG));
    stats_entry->set_drop_input_properties_miss(
            drop_reason & (1 << DROP_INPUT_PROPERTIES_MISS));
    stats_entry->set_drop_tcp_out_of_window(
            drop_reason & (1 << DROP_TCP_OUT_OF_WINDOW));
    stats_entry->set_drop_tcp_split_handshake(
            drop_reason & (1 << DROP_TCP_SPLIT_HANDSHAKE));
    stats_entry->set_drop_tcp_win_zero_drop(
            drop_reason & (1 << DROP_TCP_WIN_ZERO_DROP));
    stats_entry->set_drop_tcp_ack_err(
            drop_reason & (1 << DROP_TCP_ACK_ERR));
    stats_entry->set_drop_tcp_data_after_fin(
            drop_reason & (1 << DROP_TCP_DATA_AFTER_FIN));
    stats_entry->set_drop_tcp_non_rst_pkt_after_rst(
            drop_reason & (1 << DROP_TCP_NON_RST_PKT_AFTER_RST));
    stats_entry->set_drop_tcp_invalid_responder_first_pkt(
            drop_reason & (1 << DROP_TCP_INVALID_RESPONDER_FIRST_PKT));
    stats_entry->set_drop_tcp_unexpected_syn(
            drop_reason & (1 << DROP_TCP_UNEXPECTED_PKT));

    stats_entry->set_drop_count(
            hbm_counter +
            data->drop_stats_action_u.drop_stats_drop_stats.drop_pkts);

    return ret;
}

hal_ret_t
pd_system_populate_index_table_stats(sys::TableStatsEntry *stats_entry,
        p4pd_table_id id)
{
    hal_ret_t               ret = HAL_RET_OK;
    DirectMap               *dm;

    dm = g_hal_state_pd->dm_table(id);
    if (!dm) {
        return ret;
    }

    stats_entry->set_table_type (sys::TABLE_TYPE_INDEX);
    stats_entry->set_table_name(dm->table_name());
    stats_entry->set_table_size(dm->table_capacity());
    stats_entry->set_overflow_table_size(0);
    stats_entry->set_entries_in_use(dm->table_num_entries_in_use());
    stats_entry->set_overflow_entries_in_use(0);
    stats_entry->set_num_inserts(dm->table_num_inserts());
    stats_entry->set_num_insert_errors(dm->table_num_insert_errors());
    stats_entry->set_num_deletes(dm->table_num_deletes());
    stats_entry->set_num_delete_errors(dm->table_num_delete_errors());

    return ret;
}

hal_ret_t
pd_system_populate_flow_table_stats(sys::TableStatsEntry *stats_entry,
        p4pd_table_id id)
{
    hal_ret_t               ret = HAL_RET_OK;
    Flow                    *fl;

    fl = g_hal_state_pd->flow_table();
    if (!fl) {
        return ret;
    }

    stats_entry->set_table_type (sys::TABLE_TYPE_HASH);
    stats_entry->set_table_name(fl->table_name());
    stats_entry->set_table_size(fl->table_capacity());
    stats_entry->set_overflow_table_size(fl->oflow_table_capacity());
    stats_entry->set_entries_in_use(fl->table_num_entries_in_use());
    stats_entry->set_overflow_entries_in_use(fl->oflow_table_num_entries_in_use());
    stats_entry->set_num_inserts(fl->table_num_inserts());
    stats_entry->set_num_insert_errors(fl->table_num_insert_errors());
    stats_entry->set_num_deletes(fl->table_num_deletes());
    stats_entry->set_num_delete_errors(fl->table_num_delete_errors());

    return ret;
}

hal_ret_t
pd_system_populate_acl_tcam_table_stats(sys::TableStatsEntry *stats_entry,
        p4pd_table_id id)
{
    hal_ret_t               ret = HAL_RET_OK;
    acl_tcam                *acl_tcam;

    acl_tcam = g_hal_state_pd->acl_table();
    if (!acl_tcam) {
        return ret;
    }

    stats_entry->set_table_type (sys::TABLE_TYPE_TCAM);
    stats_entry->set_table_name(acl_tcam->table_name());
    stats_entry->set_table_size(acl_tcam->table_capacity());
    stats_entry->set_overflow_table_size(0);
    stats_entry->set_entries_in_use(acl_tcam->table_num_entries_in_use());
    stats_entry->set_overflow_entries_in_use(0);
    stats_entry->set_num_inserts(acl_tcam->table_num_inserts());
    stats_entry->set_num_insert_errors(acl_tcam->table_num_insert_errors());
    stats_entry->set_num_deletes(acl_tcam->table_num_deletes());
    stats_entry->set_num_delete_errors(acl_tcam->table_num_delete_errors());

    return ret;
}

hal_ret_t
pd_system_populate_tcam_table_stats(sys::TableStatsEntry *stats_entry,
        p4pd_table_id id)
{
    hal_ret_t               ret = HAL_RET_OK;
    Tcam                    *tcam;

    tcam = g_hal_state_pd->tcam_table(id);
    if (!tcam) {
        return ret;
    }

    stats_entry->set_table_type (sys::TABLE_TYPE_TCAM);
    stats_entry->set_table_name(tcam->table_name());
    stats_entry->set_table_size(tcam->table_capacity());
    stats_entry->set_overflow_table_size(0);
    stats_entry->set_entries_in_use(tcam->table_num_entries_in_use());
    stats_entry->set_overflow_entries_in_use(0);
    stats_entry->set_num_inserts(tcam->table_num_inserts());
    stats_entry->set_num_insert_errors(tcam->table_num_insert_errors());
    stats_entry->set_num_deletes(tcam->table_num_deletes());
    stats_entry->set_num_delete_errors(tcam->table_num_delete_errors());

    return ret;
}

hal_ret_t
pd_system_populate_hash_tcam_table_stats(sys::TableStatsEntry *stats_entry,
        p4pd_table_id id)
{
    hal_ret_t               ret = HAL_RET_OK;
    Hash                    *hash;

    hash = g_hal_state_pd->hash_tcam_table(id);

    if (!hash) {
        return ret;
    }

    stats_entry->set_table_type (sys::TABLE_TYPE_HASH_TCAM);
    stats_entry->set_table_name(hash->table_name());
    stats_entry->set_table_size(hash->table_capacity());
    stats_entry->set_overflow_table_size(hash->oflow_table_capacity());
    stats_entry->set_entries_in_use(hash->table_num_entries_in_use());
    stats_entry->set_overflow_entries_in_use(hash->oflow_table_num_entries_in_use());
    stats_entry->set_num_inserts(hash->table_num_inserts());
    stats_entry->set_num_insert_errors(hash->table_num_insert_errors());
    stats_entry->set_num_deletes(hash->table_num_deletes());
    stats_entry->set_num_delete_errors(hash->table_num_delete_errors());

    return ret;
}

static hal_ret_t
pd_system_populate_table_stats(sys::TableStatsEntry *stats_entry, 
        p4pd_table_id id)
{
    hal_ret_t               ret = HAL_RET_OK;

    HAL_TRACE_DEBUG("PD-System: Populating table stats");

    if (id >= P4TBL_ID_INDEX_MIN && id <= P4TBL_ID_INDEX_MAX) {
        return pd_system_populate_index_table_stats(stats_entry, id);
    } else if (id == P4TBL_ID_NACL) {
        return pd_system_populate_acl_tcam_table_stats(stats_entry, id);
    } else if (id >= P4TBL_ID_TCAM_MIN && id <= P4TBL_ID_TCAM_MAX) {
        return pd_system_populate_tcam_table_stats(stats_entry, id);
    } else if (id >= P4TBL_ID_HASH_OTCAM_MIN && id <= P4TBL_ID_HASH_OTCAM_MAX) {
        return pd_system_populate_hash_tcam_table_stats(stats_entry, id);
    } else if (id >= P4TBL_ID_HASH_MIN && id <= P4TBL_ID_HASH_MAX) {
        return pd_system_populate_flow_table_stats(stats_entry, id);
    } else {
        stats_entry->set_table_type (sys::TABLE_TYPE_NONE);
    }

    return ret;
}

hal_ret_t
pd_table_stats_get(pd_system_args_t *pd_sys_args)
{
    hal_ret_t               ret = HAL_RET_OK;
    SystemResponse          *rsp = pd_sys_args->rsp;
    sys::TableStatsEntry    *stats_entry = NULL;
    int                     i;

    HAL_TRACE_DEBUG("PD-System: Querying table stats");

    for (i = (int)P4TBL_ID_TBLMIN; i <= (int)P4TBL_ID_TBLMAX; i++) {
        if (i >= (int)P4TBL_ID_MPU_MIN && i <= (int)P4TBL_ID_MPU_MAX) {
            continue;
        }
        stats_entry = rsp->mutable_stats()->mutable_table_stats()->
            add_table_stats();
     	pd_system_populate_table_stats(stats_entry, (p4pd_table_id)i);
    }

    return ret;
}

}    // namespace pd
}    // namespace hal
