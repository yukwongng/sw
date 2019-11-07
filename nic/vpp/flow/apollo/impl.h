//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//

#ifndef __VPP_FLOW_APOLLO_IMPL_H__
#define __VPP_FLOW_APOLLO_IMPL_H__

#include <gen/p4gen/apollo/include/p4pd.h>
#include <nic/apollo/p4/include/defines.h>

always_inline int
pds_session_get_advance_offset (void)
{
    return 0;
}

always_inline void
pds_session_prog_x2 (vlib_buffer_t **b, u32 session_id0,
                     u32 session_id1, u16 *next, u32 *counter)
{
    return;
}

always_inline void
pds_session_prog_x1 (vlib_buffer_t *b, u32 session_id,
                     u16 *next, u32 *counter)
{
    return;
}

always_inline int
pds_flow_prog_get_next_offset (vlib_buffer_t *p0)
{
    return -(APOLLO_PREDICATE_HDR_SZ +
            (vnet_buffer(p0)->l3_hdr_offset - vnet_buffer(p0)->l2_hdr_offset));
}

always_inline int
pds_flow_prog_get_next_node (void)
{
    return FLOW_PROG_NEXT_FWD_FLOW;
}

always_inline void
pds_flow_add_tx_hdrs_x2 (vlib_buffer_t *b0, vlib_buffer_t *b1)
{
    p4_tx_cpu_hdr_t *tx0, *tx1;
    u8 flag0, flag1;

    tx0 = vlib_buffer_get_current(b0);
    tx1 = vlib_buffer_get_current(b1);
    tx0->flags_octet = 0;
    tx1->flags_octet = 0;
    flag0 = vnet_buffer (b0)->pds_data.flags;
    flag1 = vnet_buffer (b1)->pds_data.flags;
    if (PREDICT_TRUE(flag0 & VPP_CPU_FLAGS_DIRECTION)) {
        tx0->direction = 1;
    }
    if (PREDICT_TRUE(flag1 & VPP_CPU_FLAGS_DIRECTION)) {
        tx1->direction = 1;
    }
}

always_inline void
pds_flow_add_tx_hdrs_x1 (vlib_buffer_t *b0)
{
    p4_tx_cpu_hdr_t *tx0;
    u8 flag0;

    tx0 = vlib_buffer_get_current(b0);
    tx0->flags_octet = 0;
    flag0 = vnet_buffer (b0)->pds_data.flags;
    if (PREDICT_TRUE(flag0 & VPP_CPU_FLAGS_DIRECTION)) {
        tx0->direction = 1;
    }
}

static char *
pds_flow4_key2str (void *key)
{
    static char str[256];
    flow_swkey_t *k = (flow_swkey_t *)key;
    char srcstr[INET_ADDRSTRLEN + 1];
    char dststr[INET_ADDRSTRLEN + 1];

    inet_ntop(AF_INET, k->key_metadata_src, srcstr, INET_ADDRSTRLEN);
    inet_ntop(AF_INET, k->key_metadata_dst, dststr, INET_ADDRSTRLEN);
    sprintf(str, "Src:%s Dst:%s Dport:%u Sport:%u Proto:%u VNIC:%u",
            srcstr, dststr,
            k->key_metadata_dport, k->key_metadata_sport,
            k->key_metadata_proto, k->key_metadata_lkp_id);
    return str;
}

static char *
pds_flow6_key2str (void *key)
{
    static char str[256];
    flow_swkey_t *k = (flow_swkey_t *)key;
    char srcstr[INET6_ADDRSTRLEN + 1];
    char dststr[INET6_ADDRSTRLEN + 1];

    inet_ntop(AF_INET6, k->key_metadata_src, srcstr, INET6_ADDRSTRLEN);
    inet_ntop(AF_INET6, k->key_metadata_dst, dststr, INET6_ADDRSTRLEN);
    sprintf(str, "Src:%s Dst:%s Dport:%u Sport:%u Proto:%u VNIC:%u",
            srcstr, dststr,
            k->key_metadata_dport, k->key_metadata_sport,
            k->key_metadata_proto, k->key_metadata_lkp_id);
    return str;
}

static char *
pds_flow_appdata2str (void *appdata)
{
    static char str[512];
    flow_appdata_t *d = (flow_appdata_t *)appdata;
    sprintf(str, "session_index:%d flow_role:%d",
            d->session_index, d->flow_role);
    return str;
}

always_inline void
pds_flow_extract_nexthop_info(void * local_entry,
                              void * remote_entry,
                              vlib_buffer_t *p0,
                              u8 is_ip4)
{
    return;
}

#endif    // __VPP_FLOW_APOLLO_IMPL_H__

