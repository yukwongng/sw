//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//

#ifndef __VPP_IMPL_APULU_MAPPING_H__
#define __VPP_IMPL_APULU_MAPPING_H__

#include <nic/vpp/infra/utils.h>
#include <nic/apollo/packet/apulu/p4_cpu_hdr.h>

#ifdef __cplusplus
extern "C" {
#endif

static inline int
pds_vnic_id_get (void *hdr)
{
    return (((p4_rx_cpu_hdr_t*)hdr)->vnic_id);
}

static inline int
pds_ingress_bd_id_get (void *hdr)
{
    return (((p4_rx_cpu_hdr_t*)hdr)->ingress_bd_id);
}

// get arp packet header offset
static inline int
pds_arp_pkt_offset_get (void *hdr)
{
    // TODO: figure out from vnic if there is a vlan tag
    return (sizeof(p4_rx_cpu_hdr_t) + ETH_HDR_LEN);
}

// Function prototypes
int pds_dst_mac_get(void *p4_rx_meta, mac_addr_t mac_addr, bool remote,
                    uint32_t dst_addr);
void pds_mapping_table_init();

#ifdef __cplusplus
}
#endif
#endif    // __VPP_IMPL_APULU_MAPPING_H__
