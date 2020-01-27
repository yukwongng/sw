//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// athena flow cache implementation
///
//----------------------------------------------------------------------------

#include "nic/sdk/include/sdk/base.hpp"
#include "nic/sdk/include/sdk/platform.hpp"
#include "nic/sdk/include/sdk/table.hpp"
#include "nic/sdk/platform/capri/capri_tbl_rw.hpp"
#include "nic/sdk/lib/p4/p4_api.hpp"
#include "nic/sdk/lib/p4/p4_utils.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/api/include/athena/pds_flow_cache.h"
#include "nic/apollo/api/include/athena/pds_flow_session.h"
#include "ftl_wrapper.h"
#include "gen/p4gen/athena/include/p4pd.h"

using namespace sdk;
using namespace sdk::table;

extern "C" {

// Per thread address of table object
static thread_local ftl_base *ftl_table;

static char *
pds_flow6_key2str (void *key)
{
    static char str[256] = { 0 };
    flow_swkey_t *k = (flow_swkey_t *)key;
    char srcstr[INET6_ADDRSTRLEN + 1];
    char dststr[INET6_ADDRSTRLEN + 1];
    uint16_t    vnic_id;

    inet_ntop(AF_INET6, k->key_metadata_src, srcstr, INET6_ADDRSTRLEN);
    inet_ntop(AF_INET6, k->key_metadata_dst, dststr, INET6_ADDRSTRLEN);
    sprintf(str, "Src:%s Dst:%s Dport:%u Sport:%u Proto:%u "
            "Ktype:%u VNICID:%hu",
            srcstr, dststr,
            k->key_metadata_dport, k->key_metadata_sport,
            k->key_metadata_proto, k->key_metadata_ktype,
            k->key_metadata_vnic_id);
    return str;
}

static char *
pds_flow6_appdata2str (void *appdata)
{
    static char str[512] = { 0 };
    flow_appdata_t *d = (flow_appdata_t *)appdata;
    //sprintf(str, "session_id:%u", d->session_index);
    return str;
}

static sdk_ret_t
flow_cache_setup_entry_key (flow_hash_entry_t *entry,
                            pds_flow_key_t *key)
{
    if (!entry) {
        PDS_TRACE_ERR("entry is null");
        return SDK_RET_INVALID_ARG;
    }

    if (key->ip_addr_family == IP_AF_IPV4) { 
        ftlv6_set_key_ktype(entry, IP_AF_IPV4);
    } else {
        ftlv6_set_key_ktype(entry, IP_AF_IPV6);
    }
    ftlv6_set_key_src_ip(entry, key->ip_saddr);
    ftlv6_set_key_dst_ip(entry, key->ip_daddr);
    ftlv6_set_key_sport(entry, key->l4_sport);
    ftlv6_set_key_dport(entry, key->l4_dport);
    ftlv6_set_key_proto(entry, key->ip_proto);

    return SDK_RET_OK;
}

sdk_ret_t
pds_flow_cache_create (uint32_t core_id)
{
    sdk_table_factory_params_t factory_params = { 0 };

    factory_params.table_id = P4TBL_ID_FLOW;
    factory_params.num_hints = 4;
    factory_params.max_recircs = 8;
    factory_params.key2str = pds_flow6_key2str;
    factory_params.appdata2str = pds_flow6_appdata2str;
    factory_params.thread_id = core_id;
    factory_params.entry_alloc_cb = flow_hash_entry_t::alloc;

    if ((ftl_table = flow_hash::factory(&factory_params)) == NULL) {
        PDS_TRACE_ERR("Table creation failed in thread %u", core_id);
        return SDK_RET_OOM;
    }
    return SDK_RET_OK;
}

sdk_ret_t
pds_flow_cache_entry_create (pds_flow_spec_t *spec)
{
    sdk_ret_t ret;
    sdk_table_api_params_t params = { 0 };
    flow_hash_entry_t entry;
    uint32_t session_info_id;

    if (!spec) {
        PDS_TRACE_ERR("spec is null");
        return SDK_RET_INVALID_ARG;
    }
    session_info_id = spec->data.session_info_id;
    if (session_info_id > PDS_FLOW_SESSION_INFO_ID_MAX) {
        PDS_TRACE_ERR("session id %u is invalid", session_info_id);
        return SDK_RET_INVALID_ARG;
    }

    entry.clear();
    if ((ret = flow_cache_setup_entry_key(&entry, &spec->key))
             != SDK_RET_OK)
         return ret;
    ftlv6_set_session_index(&entry, session_info_id); 
    params.entry = &entry;
    return ftl_table->insert(&params);
}

sdk_ret_t
pds_flow_cache_entry_read (pds_flow_key_t *key,
                     pds_flow_info_t *info)
{
    sdk_ret_t ret;
    sdk_table_api_params_t params = { 0 };
    flow_hash_entry_t entry;

    if (!key || !info) {
        PDS_TRACE_ERR("key/info is null");
        return SDK_RET_INVALID_ARG;
    }

    entry.clear();
    if ((ret = flow_cache_setup_entry_key(&entry, key))
             != SDK_RET_OK)
         return ret;
    params.entry = &entry;
    ret = ftl_table->get(&params);
    if (ret != SDK_RET_OK) {
        return ret;
    } else {
        info->spec.data.session_info_id = ftlv6_get_session_id(&entry);
        return SDK_RET_OK;
    }
}

sdk_ret_t
pds_flow_cache_entry_update (pds_flow_spec_t *spec)
{
    sdk_ret_t ret;
    sdk_table_api_params_t params = { 0 };
    flow_hash_entry_t entry;

    if (!spec) {
        PDS_TRACE_ERR("spec is null");
        return SDK_RET_INVALID_ARG;
    }

    entry.clear();
    if ((ret = flow_cache_setup_entry_key(&entry, &spec->key))
             != SDK_RET_OK)
         return ret;
    params.entry = &entry;
    return ftl_table->update(&params);
}

sdk_ret_t
pds_flow_cache_entry_delete (pds_flow_key_t *key)
{
    sdk_ret_t ret;
    sdk_table_api_params_t params = { 0 };
    flow_hash_entry_t entry;

    if (!key) {
        PDS_TRACE_ERR("key is null");
        return SDK_RET_INVALID_ARG;
    }

    entry.clear();
    if ((ret = flow_cache_setup_entry_key(&entry, key))
             != SDK_RET_OK)
         return ret;
    params.entry = &entry;
    return ftl_table->remove(&params);
}

}
