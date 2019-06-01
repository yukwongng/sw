//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#ifndef _NVME_PD_H
#define _NVME_PD_H

#include "nic/sdk/platform/capri/capri_common.hpp"
#include "p4pd_nvme_api.h"


#define HOSTMEM_PAGE_SIZE  (1 << 12)  //4096 Bytes
#define MAX_LIFS           2048

#define PACKED __attribute__((__packed__))

#define SESSTX_PREXTS_RING_ID      RING_ID_0
#define SESSTX_POSTXTS_RING_ID      RING_ID_1

#define MAX_SESSXTSTX_RINGS 2
#define MAX_SESSXTSTX_HOST_RINGS 0

#define MAX_SESSXTSRX_RINGS 2
#define MAX_SESSXTSRX_HOST_RINGS 0

#define SESSTX_PREDGST_RING_ID      RING_ID_0
#define SESSTX_POSTDGST_RING_ID     RING_ID_1

#define MAX_SESSDGSTTX_RINGS 2
#define MAX_SESSDGSTTX_HOST_RINGS 0

#define MAX_SESSDGSTRX_RINGS 2
#define MAX_SESSDGSTRX_HOST_RINGS 0

typedef struct nvme_cmd_context_s {
    uint8_t pad[64];
} nvme_cmd_context_t;

typedef struct nvme_pdu_context_s {
    uint8_t pad[1024];
} nvme_pdu_context_t;

typedef struct nvme_aol_s {
    uint8_t pad[128];
} nvme_aol_t;

typedef struct nvme_iv_s {
    uint8_t pad[16];
} nvme_iv_t;

#define NVME_TX_SESS_XTSQ_DEPTH 64
#define NVME_TX_SESS_XTSQ_ENTRY_SIZE sizeof(s1_t0_nvme_sessprexts_tx_sess_wqe_process_bitfield_t)
#define NVME_TX_SESS_XTSQ_SIZE NVME_TX_SESS_XTSQ_DEPTH * NVME_TX_SESS_XTSQ_ENTRY_SIZE
#define NVME_TX_SESS_DGSTQ_DEPTH 64
#define NVME_TX_SESS_DGSTQ_ENTRY_SIZE sizeof(s1_t0_nvme_sesspredgst_tx_sess_wqe_process_bitfield_t)
#define NVME_TX_SESS_DGSTQ_SIZE NVME_TX_SESS_DGSTQ_DEPTH * NVME_TX_SESS_DGSTQ_ENTRY_SIZE
#define NVME_RX_SESS_XTSQ_DEPTH 64
//#define NVME_RX_SESS_XTSQ_ENTRY_SIZE sizeof(s1_t0_nvme_sessprexts_rx_sess_wqe_process_bitfield_t)
#define NVME_RX_SESS_XTSQ_ENTRY_SIZE sizeof(s1_t0_nvme_sessprexts_tx_sess_wqe_process_bitfield_t)
#define NVME_RX_SESS_XTSQ_SIZE NVME_RX_SESS_XTSQ_DEPTH * NVME_RX_SESS_XTSQ_ENTRY_SIZE
#define NVME_RX_SESS_DGSTQ_DEPTH 64
//#define NVME_RX_SESS_DGSTQ_ENTRY_SIZE sizeof(s1_t0_nvme_sesspredgst_rx_sess_wqe_process_bitfield_t)
#define NVME_RX_SESS_DGSTQ_ENTRY_SIZE sizeof(s1_t0_nvme_sesspredgst_tx_sess_wqe_process_bitfield_t)
#define NVME_RX_SESS_DGSTQ_SIZE NVME_RX_SESS_DGSTQ_DEPTH * NVME_RX_SESS_DGSTQ_ENTRY_SIZE

#define nvme_nscb_t s2_t0_nvme_req_tx_nscb_process_bitfield_t
#define nvme_txsessprodcb_t s5_t0_nvme_req_tx_sessprodcb_process_bitfield_t
#define nvme_rxsessprodcb_t s5_t0_nvme_req_rx_sessprodcb_process_bitfield_t
#define nvme_resourcecb_t s4_t1_nvme_req_tx_resourcecb_process_bitfield_t
#define nvme_cmd_context_ring_entry_t s5_t1_nvme_req_tx_cmdid_fetch_process_bitfield_t
#define nvme_pdu_context_ring_entry_t s5_t2_nvme_req_tx_pduid_fetch_process_bitfield_t

//HBM organization of NVME data path tables

typedef enum nvme_dpath_ds_type_s {
    NVME_TYPE_NSCB = 0,
    NVME_TYPE_CMD_CONTEXT,
    NVME_TYPE_CMD_CONTEXT_RING,
    NVME_TYPE_RESOURCECB,
    NVME_TYPE_TX_SESSPRODCB,
    NVME_TYPE_RX_SESSPRODCB,
    NVME_TYPE_TX_PDU_CONTEXT,
    NVME_TYPE_TX_PDU_CONTEXT_RING,
    NVME_TYPE_RX_PDU_CONTEXT,
    NVME_TYPE_RX_PDU_CONTEXT_RING,
    NVME_TYPE_TX_SESS_XTSQ,
    NVME_TYPE_TX_SESS_DGSTQ,
    NVME_TYPE_RX_SESS_XTSQ,
    NVME_TYPE_RX_SESS_DGSTQ,
    NVME_TYPE_TX_XTS_AOL_ARRAY,
    NVME_TYPE_TX_XTS_IV_ARRAY,
    NVME_TYPE_RX_XTS_AOL_ARRAY,
    NVME_TYPE_RX_XTS_IV_ARRAY,
    NVME_TYPE_MAX
} nvme_dpath_ds_type_t;
     
typedef struct nvme_hbm_alloc_info_s {
    int type;
    int num_entries;
    int size;
} nvme_hbm_alloc_info_t;
    
static nvme_hbm_alloc_info_t nvme_hbm_alloc_table[] = {
    {NVME_TYPE_NSCB, 512, sizeof(nvme_nscb_t)},
    {NVME_TYPE_CMD_CONTEXT, 512, sizeof(nvme_cmd_context_t)},
    {NVME_TYPE_CMD_CONTEXT_RING, 512, sizeof(nvme_cmd_context_ring_entry_t)},
    {NVME_TYPE_RESOURCECB, 1, sizeof(nvme_resourcecb_t)},
    {NVME_TYPE_TX_SESSPRODCB, 512, sizeof(nvme_txsessprodcb_t)},
    {NVME_TYPE_RX_SESSPRODCB, 512, sizeof(nvme_rxsessprodcb_t)},
    {NVME_TYPE_TX_PDU_CONTEXT, 512, sizeof(nvme_pdu_context_t)},
    {NVME_TYPE_TX_PDU_CONTEXT_RING, 512, sizeof(nvme_pdu_context_ring_entry_t)},
    {NVME_TYPE_RX_PDU_CONTEXT, 512, sizeof(nvme_pdu_context_t)},
    {NVME_TYPE_RX_PDU_CONTEXT_RING, 512, sizeof(nvme_pdu_context_ring_entry_t)},
    {NVME_TYPE_TX_SESS_XTSQ, 512, NVME_TX_SESS_XTSQ_DEPTH * NVME_TX_SESS_XTSQ_ENTRY_SIZE},
    {NVME_TYPE_TX_SESS_DGSTQ, 512, NVME_TX_SESS_DGSTQ_DEPTH * NVME_TX_SESS_DGSTQ_ENTRY_SIZE},
    {NVME_TYPE_RX_SESS_XTSQ, 512, NVME_RX_SESS_XTSQ_DEPTH * NVME_RX_SESS_XTSQ_ENTRY_SIZE},
    {NVME_TYPE_RX_SESS_DGSTQ, 512, NVME_RX_SESS_DGSTQ_DEPTH * NVME_RX_SESS_DGSTQ_ENTRY_SIZE},
    //XXX: Dividing the resources by 16 for now. 
    {NVME_TYPE_TX_XTS_AOL_ARRAY, CAPRI_BARCO_XTS_RING_SLOTS/16, sizeof(nvme_aol_t)},
    {NVME_TYPE_TX_XTS_IV_ARRAY, CAPRI_BARCO_XTS_RING_SLOTS/16, sizeof(nvme_iv_t)},
    {NVME_TYPE_RX_XTS_AOL_ARRAY, CAPRI_BARCO_XTS_RING_SLOTS/16, sizeof(nvme_aol_t)},
    {NVME_TYPE_RX_XTS_IV_ARRAY, CAPRI_BARCO_XTS_RING_SLOTS/16, sizeof(nvme_iv_t)},
};

static inline int nvme_hbm_offset(int type) {
    int i;
    int offset = 0;
    for (i = 0; i < type; i++) {
        assert(i == nvme_hbm_alloc_table[i].type);
        offset += nvme_hbm_alloc_table[i].size * nvme_hbm_alloc_table[i].num_entries;
    }
    return (offset);
}

static inline int nvme_hbm_resource_max(int type) {
    assert(type < NVME_TYPE_MAX);
    return (nvme_hbm_alloc_table[type].num_entries);
}

static inline int nvme_hbm_resource_size(int type) {
    assert(type < NVME_TYPE_MAX);
    return (nvme_hbm_alloc_table[type].size);
}

static inline int nvme_hbm_resource_total_mem(int type) {
    assert(type < NVME_TYPE_MAX);
    return (nvme_hbm_resource_size(type) * nvme_hbm_resource_max(type));
}
    
static inline uint8_t *
memrev (uint8_t *block, size_t elnum)
{
     uint8_t *s, *t, tmp;

    for (s = block, t = s + (elnum - 1); s < t; s++, t--) {
        tmp = *s;
        *s = *t;
        *t = tmp;
    }
     return block;
}

#define MAX3(x, y, z) ((x) > (y) ? ((x) > (z) ? (x) : (z)) : ((y) > (z) ? (y) : (z)))

#endif // _NVME_PD_H
