//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//

#include "includes.h"

// *INDENT-OFF*
VLIB_PLUGIN_REGISTER () = {
    .description = "Arp-proxy Plugin",
};
// *INDENT-ON*

vlib_node_registration_t arp_proxy_node;

arp_proxy_main_t arp_proxy_main;

always_inline void
arp_proxy_trace_fill (arp_proxy_trace_t *trace, ethernet_arp_header_t *arp,
                      mac_addr_t mac, u32 id)
{
    if (arp) {
        clib_memcpy(&trace->src, &arp->ip4_over_ethernet[0].ip4,
                sizeof(ip4_address_t));
        clib_memcpy(&trace->dst, &arp->ip4_over_ethernet[1].ip4,
                sizeof(ip4_address_t));
        clib_memcpy(trace->smac, &arp->ip4_over_ethernet[1].mac,
                ETH_ADDR_LEN);
    }
    clib_memcpy(trace->vr_mac, mac, ETH_ADDR_LEN);
    trace->bd = id;
}

static int
arp_proxy_dst_mac_get(void *hdr, mac_addr_t mac_addr, uint32_t dst_addr)
{
    u16 vpc_id;
    u16 bd_id;
    int ret = 0;

    vpc_id = ((p4_rx_cpu_hdr_t *)hdr)->vpc_id;
    bd_id = ((p4_rx_cpu_hdr_t *)hdr)->ingress_bd_id;

    if (arp_proxy_main.dst_mac_get_cb &&
        pds_impl_db_vpc_is_control_vpc(vpc_id)) {
        ret = arp_proxy_main.dst_mac_get_cb(vpc_id, bd_id, mac_addr, dst_addr);
    } else {
        ret = pds_dst_mac_get(vpc_id, bd_id, mac_addr, dst_addr);
    }

    return ret;
}

void
arp_proxy_register_vendor_dst_mac_get_cb(arp_proxy_vendor_dst_mac_get_cb cb)
{
    arp_proxy_main.dst_mac_get_cb = cb;
}

always_inline void
arp_proxy_internal (vlib_buffer_t *p0, u16 *next0, u32 *counter,
                    vlib_node_runtime_t *node, vlib_main_t *vm)
{
    ethernet_header_t *e0;
    ethernet_arp_header_t *arp = NULL;
    arp_proxy_trace_t *trace;
    void *p4_rx_meta = NULL;
    mac_addr_t mac;
    u32 vrip = 0, bd_id = 0;
    u32 dst;
    u32 offset = 0;
    u16 vnic_nh_hw_id;
    u8 *vrmac;

    p4_rx_meta = (void*) (vlib_buffer_get_current(p0));
    bd_id = pds_ingress_bd_id_get(p4_rx_meta);
    if (PREDICT_FALSE(pds_vnic_nexthop_get(p4_rx_meta, &vnic_nh_hw_id,
                                           &offset))) {
        counter[ARP_PROXY_COUNTER_VNIC_MISSING]++;
        goto error;
    }
    vnet_buffer(p0)->pds_tx_data.nh_id = vnic_nh_hw_id;

    arp = (ethernet_arp_header_t*) (vlib_buffer_get_current(p0) + offset);
    dst = clib_net_to_host_u32(arp->ip4_over_ethernet[1].ip4.data_u32);
    if (PREDICT_FALSE(pds_subnet_check(bd_id, dst))) {
        counter[ARP_PROXY_COUNTER_SUBNET_CHECK_FAIL]++;
        goto error;
    }

    if (PREDICT_TRUE(
            arp->opcode ==
            clib_host_to_net_u16 (ETHERNET_ARP_OPCODE_request))) {
        if (0 != arp_proxy_dst_mac_get(p4_rx_meta, mac, dst)) {
            counter[ARP_PROXY_COUNTER_NO_MAC]++;
            goto error;
        }

        e0 = vlib_buffer_get_current(p0) + VPP_P4_TO_ARM_HDR_SZ;
        if (0 == clib_memcmp(mac, e0->src_address, ETH_ADDR_LEN)) {
            // this is arp probe sent by host to detect DAD, so drop this.
            counter[ARP_PROXY_COUNTER_DAD_DROP]++;
            goto error;
        }
        // Ethernet
        if (0 != pds_impl_db_vr_ip_mac_get(bd_id, &vrip, &vrmac)) {
            counter[ARP_PROXY_COUNTER_SUBNET_CHECK_FAIL]++;
            goto error;
        }
        clib_memcpy(&e0->dst_address, &e0->src_address, ETH_ADDR_LEN);
        clib_memcpy(&e0->src_address, vrmac, ETH_ADDR_LEN);

        // ARP Reply
        arp->opcode = clib_host_to_net_u16 (ETHERNET_ARP_OPCODE_reply);
        clib_memswap(&arp->ip4_over_ethernet[1].ip4.data_u32,
                     &arp->ip4_over_ethernet[0].ip4.data_u32,
                     sizeof(u32));
        clib_memcpy(&arp->ip4_over_ethernet[1].mac,
                    &arp->ip4_over_ethernet[0].mac, ETH_ADDR_LEN);
        clib_memcpy(&arp->ip4_over_ethernet[0].mac,
                    mac, ETH_ADDR_LEN);
        *next0 = ARP_PROXY_NEXT_EXIT;
        counter[ARP_PROXY_COUNTER_REPLY_SUCCESS]++;
        if (PREDICT_FALSE(node->flags & VLIB_NODE_FLAG_TRACE &&
                          p0->flags & VLIB_BUFFER_IS_TRACED)) {
            trace = vlib_add_trace (vm, node, p0, sizeof (trace[0]));
            arp_proxy_trace_fill(trace, arp, mac, bd_id);
        }
    } else {
        // TODO
        counter[ARP_PROXY_COUNTER_NOT_ARP_REQUEST]++;
        goto error;
    }
    return;

error:
    *next0 = ARP_PROXY_NEXT_DROP;
    if (PREDICT_FALSE(node->flags & VLIB_NODE_FLAG_TRACE &&
                      p0->flags & VLIB_BUFFER_IS_TRACED)) {
        trace = vlib_add_trace (vm, node, p0, sizeof (trace[0]));
        arp_proxy_trace_fill(trace, arp, mac, bd_id);
    }
    return;
}

static uword
arp_proxy (vlib_main_t *vm,
           vlib_node_runtime_t *node,
           vlib_frame_t *from_frame)
{
    u32 counter[ARP_PROXY_COUNTER_LAST] = {0};

    PDS_PACKET_LOOP_START {

        PDS_PACKET_DUAL_LOOP_START (WRITE, READ) {
            vlib_buffer_t *p0, *p1;

            p0 = PDS_PACKET_BUFFER(0);
            p1 = PDS_PACKET_BUFFER(1);
            vnet_buffer(p0)->sw_if_index[VLIB_TX] = vnet_buffer(p0)->sw_if_index[VLIB_RX];
            vnet_buffer(p1)->sw_if_index[VLIB_TX] = vnet_buffer(p1)->sw_if_index[VLIB_RX];

            arp_proxy_internal(p0, PDS_PACKET_NEXT_NODE_PTR(0), counter, node, vm);
            arp_proxy_internal(p1, PDS_PACKET_NEXT_NODE_PTR(1), counter, node, vm);
            vlib_buffer_advance(p0,
                                VPP_P4_TO_ARM_HDR_SZ - VPP_ARM_TO_P4_HDR_SZ );
            vlib_buffer_advance(p1,
                                VPP_P4_TO_ARM_HDR_SZ - VPP_ARM_TO_P4_HDR_SZ );

        } PDS_PACKET_DUAL_LOOP_END;

        PDS_PACKET_SINGLE_LOOP_START {
            vlib_buffer_t *p0;

            p0 = PDS_PACKET_BUFFER(0);

            vnet_buffer(p0)->sw_if_index[VLIB_TX] = vnet_buffer(p0)->sw_if_index[VLIB_RX];

            arp_proxy_internal(p0, PDS_PACKET_NEXT_NODE_PTR(0), counter, node, vm);
            vlib_buffer_advance(p0,
                                VPP_P4_TO_ARM_HDR_SZ - VPP_ARM_TO_P4_HDR_SZ );

        } PDS_PACKET_SINGLE_LOOP_END;

    } PDS_PACKET_LOOP_END;

#define _(n, s) \
    vlib_node_increment_counter (vm, node->node_index,           \
            ARP_PROXY_COUNTER_##n,                               \
            counter[ARP_PROXY_COUNTER_##n]);
    foreach_arp_proxy_counter
#undef _

    return from_frame->n_vectors;
}

static u8 *
arp_proxy_trace (u8 * s, va_list * args)
{
    CLIB_UNUSED (vlib_main_t * vm) = va_arg (*args, vlib_main_t *);
    CLIB_UNUSED (vlib_node_t * node) = va_arg (*args, vlib_node_t *);
    arp_proxy_trace_t *t = va_arg (*args, arp_proxy_trace_t *);

    s = format(s, "%U %U -> %U %U %d",
               format_ip4_address, &t->src,
               format_mac_address_t, &t->smac,
               format_ip4_address, &t->dst,
               format_mac_address_t, &t->vr_mac, t->bd);
    return s;
}

static char * arp_proxy_error_strings[] = {
#define _(n,s) s,
    foreach_arp_proxy_counter
#undef _
};

VLIB_REGISTER_NODE (arp_proxy_node) = {
    .function = arp_proxy,
    .name = "pds-arp-proxy",
    /* Takes a vector of packets. */
    .vector_size = sizeof (u32),

    .n_errors = ARP_PROXY_COUNTER_LAST,
    .error_strings = arp_proxy_error_strings,

    .n_next_nodes = ARP_PROXY_N_NEXT,
    .next_nodes = {
#define _(s,n) [ARP_PROXY_NEXT_##s] = n,
    foreach_arp_proxy_next
#undef _
    },

    .format_trace = arp_proxy_trace,
};

static clib_error_t *
arp_proxy_init (vlib_main_t * vm)
{
    pds_arp_proxy_pipeline_init();
    pds_mapping_table_init();
    return 0;
}

VLIB_INIT_FUNCTION (arp_proxy_init) =
{
    .runs_after = VLIB_INITS("pds_infra_init"),
};

