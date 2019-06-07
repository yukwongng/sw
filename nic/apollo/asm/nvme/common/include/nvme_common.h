#ifndef __NVME_COMMON_H
#define __NVME_COMMON_H

#include "capri.h"

#define LOG_NUM_PRP_BYTES       3   //2^3 = 8
#define LOG_NUM_PAGE_PTR_BYTES  3   //2^3 = 8

//ring pindex/cindex
#define SQ_P_INDEX              d.{pi_0}.hx
#define SQ_C_INDEX              d.{ci_0}.hx
#define CMDID_RING_PROXY_CI     d.{cmdid_ring_proxy_ci}.hx
#define CMDID_RING_PROXY_CI_LE  d.{cmdid_ring_proxy_ci}
#define CMDID_RING_PI           d.{cmdid_ring_pi}.hx
#define CMDID_RING_CI           d.{cmdid_ring_ci}.hx
#define CMDID                   d.{cmdid}
#define TX_PDUID_RING_PROXY_CI     d.{tx_pduid_ring_proxy_ci}.hx
#define TX_PDUID_RING_PROXY_CI_LE  d.{tx_pduid_ring_proxy_ci}
#define TX_PDUID_RING_PI           d.{tx_pduid_ring_pi}.hx
#define TX_PDUID_RING_CI           d.{tx_pduid_ring_ci}.hx
#define RX_PDUID_RING_PROXY_CI     d.{rx_pduid_ring_proxy_ci}.hx
#define RX_PDUID_RING_PROXY_CI_LE  d.{rx_pduid_ring_proxy_ci}
#define RX_PDUID_RING_PI           d.{rx_pduid_ring_pi}.hx
#define RX_PDUID_RING_CI           d.{rx_pduid_ring_ci}.hx
#define PDUID                   d.{pduid}
#define XTS_Q_PI                d.{xts_q_pi}.hx
#define XTS_Q_CI                d.{xts_q_ci}.hx
#define DGST_Q_PI               d.{dgst_q_pi}.hx
#define DGST_Q_CI               d.{dgst_q_ci}.hx
#define TCP_Q_PI                d.{tcp_q_pi}.hx
#define TCP_Q_CI                d.{tcp_q_ci}.hx
#define SESSPREXTSTX_P_INDEX    d.{pi_0}.hx
#define SESSPREXTSTX_C_INDEX    d.{ci_0}.hx
#define SESSPOSTXTSTX_P_INDEX   d.{pi_1}.hx
#define SESSPOSTXTSTX_C_INDEX   d.{ci_1}.hx
#define NMDPR_IDX               d.{idx}.wx
#define SESSPREDGSTTX_P_INDEX   d.{pi_0}.hx
#define SESSPREDGSTTX_C_INDEX   d.{ci_0}.hx
#define SESSPOSTDGSTTX_P_INDEX  d.{pi_1}.hx
#define SESSPOSTDGSTTX_C_INDEX  d.{ci_1}.hx

//cb, ring entry sizes
#define LOG_SQE_SIZE            6   //2^6 = 64
#define LOG_CQE_SIZE            4   //2^4 = 16
#define LOG_NSCB_SIZE           6   //2^6 = 64
#define LOG_TXSESSPRODCB_SIZE   6   //2^6 = 64

#define LOG_CMDID_SIZE              1   //2^1 = 2
#define LOG_CMDID_RING_ENTRY_SIZE   LOG_CMDID_SIZE

#define LOG_PDUID_SIZE              1   //2^1 = 2
#define LOG_PDUID_RING_ENTRY_SIZE   LOG_PDUID_SIZE
#define TX_LOG_PDUID_RING_ENTRY_SIZE    LOG_PDUID_RING_ENTRY_SIZE
#define RX_LOG_PDUID_RING_ENTRY_SIZE    LOG_PDUID_RING_ENTRY_SIZE

#define LOG_SESS_Q_ENTRY_SIZE   (LOG_CMDID_SIZE+LOG_PDUID_SIZE)
#define LOG_XTS_DESC_SIZE       7   //2^7 = 128
#define LOG_AOL_DESC_SIZE       6   //2^6 = 64
#define AOL_DESC_SIZE           (1 << (LOG_AOL_DESC_SIZE))
#define LOG_AOL_PAIR_SIZE       7   //2^7 = 128
#define LOG_IV_SIZE             4   //2^4 = 16
#define LOG_ONE_AOL_SIZE        4   //2^4 = 16
#define LOG_DGST_DESC_SIZE      6   //2^6 = 64

#define LOG_NMDPR_RING_SIZE     (CAPRI_TNMDPR_BIG_RING_SHIFT)
#define NMDPR_RING_SIZE         (1 << LOG_NMDPR_RING_SIZE)
#define NMDPR_RING_SIZE_MASK    (NMDPR_RING_SIZE-1)
#define LOG_NMDPR_RING_ENTRY_SIZE   3
#define NMDPR_RING_ENTRY_SIZE   (1 << LOG_NMDPR_RING_ENTRY_SIZE)
#define LOG_HBM_AL_RING_ENTRY_SIZE  3 // 2^3 = 8

//rings, their ids and priorities
//sq
#define MAX_REQ_TX_RINGS        1

#define SQ_PRI                  0

#define SQ_RING_ID              0

//sessxtstx
#define MAX_SESSXTS_TX_RINGS    2

#define SESSPREXTS_TX_PRI       1
#define SESSPOSTXTS_TX_PRI      0

#define SESSPREXTS_TX_RING_ID   0
#define SESSPOSTXTS_TX_RING_ID  1

//sessdgsttx
#define MAX_SESSDGST_TX_RINGS   2

#define SESSPREDGST_TX_PRI      1
#define SESSPOSTDGST_TX_PRI     0

#define SESSPREDGST_TX_RING_ID   0
#define SESSPOSTDGST_TX_RING_ID  1

//globals
#define K_GLOBAL_LIF            k.phv_global_common_lif
#define K_GLOBAL_QTYPE          k.phv_global_common_qtype
#define K_GLOBAL_QID            k.phv_global_common_qid
#define K_GLOBAL_CB_ADDR        k.phv_global_cb_addr
#define K_GLOBAL_LIF_QID        k.{phv_global_common_lif...phv_global_common_qid}

//qtypes
#define NVME_QTYPE_SQ           0
#define NVME_QTYPE_CQ           1
#define NVME_QTYPE_ARMQ         2
#define NVME_QTYPE_SESSXTSTX    3
#define NVME_QTYPE_SESSDGSTTX   4
#define NVME_QTYPE_SESSXTSRX    5
#define NVME_QTYPE_SESSDGSTRX   6

//dma cmd ptrs
#define NVME_REQ_TX_DMA_CMD_PTR (PHV_FIELD_START_OFFSET(cmd_ctxt_dma_dma_cmd_type)/16)
#define NVME_SESSPREXTSTX_DMA_CMD_PTR (PHV_FIELD_START_OFFSET(pkt_desc_dma_dma_cmd_type)/16)
#define NVME_SESSPOSTXTSTX_DMA_CMD_PTR (PHV_FIELD_START_OFFSET(last_ddgst_aol_desc_dma_dma_cmd_type)/16)
#define NVME_SESS_POST_DGST_TX_DMA_CMD_PTR (PHV_FIELD_START_OFFSET(tcp_wqe_dma0_dma_cmd_type)/16)
#define NVME_SESS_PRE_DGST_TX_DMA_CMD_PTR (PHV_FIELD_START_OFFSET(ddgst_desc_dma_dma_cmd_type)/16)

//Digest related defines
#define HW_DGST_STATUS_SIZE                  8
#define HW_DGST_INTG_TAG_SIZE                8

#define HW_DGST_DB_ADDR                      (CAP_ADDR_BASE_MD_HENS_OFFSET + CAP_HENS_CSR_DHS_CRYPTO_CTL_CP_CFG_Q_PD_IDX_BYTE_ADDRESS)
#define HW_CMD_DGST_DB_ON                    (1<<3)
#define HW_CMD_DGST_OPAQUE_TAG_ON            (1<<4)
#define HW_CMD_DGST_SRC_LIST_ADDR            (1<<5)
#define HW_CMD_DGST_INTG_TAG_TYPE_CRC32C     (1<<13)
#define HW_CMD_HDGST                         (HW_CMD_DGST_OPAQUE_TAG_ON | HW_CMD_DGST_DB_ON | HW_CMD_DGST_INTG_TAG_TYPE_CRC32C)
#define HW_CMD_DDGST                         (HW_CMD_DGST_SRC_LIST_ADDR | HW_CMD_DGST_INTG_TAG_TYPE_CRC32C)
#define HW_DGST_NEXT_PTR_OFFSET              48

#define HW_DGST_STATUS_VALID_BIT             15
#define HW_DGST_STATUS_ENC_ERR_START_BIT     12

#define HW_DGST_STATUS_VALID                 (1<<3)
#define HW_DGST_STATUS_NO_ERR                0
#define HW_DGST_STATUS_AXI_TIMEOUT           1
#define HW_DGST_STATUS_AXI_DATA_ERR          2
#define HW_DGST_STATUS_AXI_ADDR_ERR          3
#define HW_DGST_STATUS_COMP_FAIL             4
#define HW_DGST_STATUS_DATA_TOO_LONG         5
#define HW_DGST_STATUS_CKSUM_FAIL            6
#define HW_DGST_STATUS_SGL_DESC_ERR          7

#define HW_DGST_STATUS_SUCCESS               (HW_DGST_STATUS_VALID | HW_DGST_STATUS_NO_ERR)



//command context offsets
#define LOG_CMD_CTXT_SIZE               6   //2^6 = 64
#define NVME_CMD_CTXT_SIZE              (1 << LOG_CMD_CTXT_SIZE)

//pdu context offsets
#define LOG_PDU_CTXT_SIZE               10  //2^10 = 1024
#define NVME_PDU_CTXT_SIZE              (1 << LOG_PDU_CTXT_SIZE)
#define NVME_PDU_CTXT0_SIZE             64
#define NVME_PDU_CTXT1_SIZE             64
#define NVME_PDU_CTXT_HDR_SIZE          (NVME_PDU_CTXT0_SIZE + NVME_PDU_CTXT1_SIZE)

#define NVME_PDU_CTXT_PRP_LIST_OFFSET   NVME_PDU_CTXT_HDR_SIZE
#define NVME_PDU_CTXT_MAX_PRPS          40  // ideally we need 33 prps, but aligning
#define NVME_PDU_CTXT_PRP_LIST_SIZE     (NVME_PDU_CTXT_MAX_PRPS << LOG_NUM_PRP_BYTES)

#define NVME_PDU_CTXT_PAGE_LIST_OFFSET  (NVME_PDU_CTXT_HDR_SIZE + NVME_PDU_CTXT_PRP_LIST_SIZE)
#define NVME_PDU_CTXT_MAX_PAGE_PTRS     16
#define NVME_PDU_CTXT_PAGE_LIST_SIZE    (NVME_PDU_CTXT_MAX_PAGE_PTRS << LOG_NUM_PAGE_PTR_BYTES)

#define NVME_PDU_CTXT_AOL_DESC_LIST_OFFSET  (NVME_PDU_CTXT_HDR_SIZE + NVME_PDU_CTXT_PRP_LIST_SIZE + NVME_PDU_CTXT_PAGE_LIST_SIZE)
#define NVME_PDU_CTXT_MAX_AOL_DESCS         7 //7*64B
#define NVME_PDU_CTXT_AOL_DESC_LIST_SIZE    (NVME_PDU_CTXT_MAX_AOL_DESCS  << LOG_AOL_DESC_SIZE)

#define NVME_PDU_CTXT_HDGST_STATUS_OFFSET  (NVME_PDU_CTXT0_SIZE + 32)
#define NVME_PDU_CTXT_DDGST_STATUS_OFFSET  (NVME_PDU_CTXT_HDGST_STATUS_OFFSET + HW_DGST_STATUS_SIZE + HW_DGST_INTG_TAG_SIZE)

struct nvme_pdu_ctxt_page_entry_t {
    len:16;
    page_addr:48;
};

#define NVME_PDU_CTXT_PAGE_ENTRY_SIZE sizeof(struct nvme_pdu_ctxt_page_entry_t)

struct hbm_al_ring_entry_t {
    pad : 16;
    len : 14;
    descr_addr : 34;
};

//XTS related defines
#define HW_XTS_TX_DB_ADDR        CAPRI_BARCO_MD_HENS_REG_XTS0_PRODUCER_IDX
#define HW_XTS_RX_DB_ADDR        CAPRI_BARCO_MD_HENS_REG_XTS1_PRODUCER_IDX

#define HW_XTS_DESC_FLAGS_OPAQUE_TAG_WR_EN 1

#define HW_CMD_XTS_ENCRYPT      0x40000000
#define HW_CMD_XTS_DECRYPT      0x40100000

//Semaphores
#define NVME_TX_NMDPR_ALLOC_IDX CAPRI_SEM_TNMDPR_BIG_ALLOC_INF_ADDR
#define NVME_TX_NMDPR_FREE_IDX  CAPRI_SEM_TNMDPR_BIG_FREE_INC_ADDR

#define NVME_RX_NMDPR_ALLOC_IDX CAPRI_SEM_RNMDPR_BIG_ALLOC_INF_ADDR
#define NVME_RX_NMDPR_FREE_IDX  CAPRI_SEM_RNMDPR_BIG_FREE_INC_ADDR

#define NVME_LOG_MAX_DATA_SIZE_IN_PAGE  13  // 2^13 = 8K
#define NVME_MAX_DATA_SIZE_IN_PAGE      (1 << NVME_LOG_MAX_DATA_SIZE_IN_PAGE)
#define NVME_MAX_DATA_SIZE_IN_PAGE_MASK (NVME_MAX_DATA_SIZE_IN_PAGE - 1)

#define NVME_PAGE_DATA_LEN_SHIFT    48

//NVME-over-TCP related defines

#define PKT_DESC_SCRATCH_OVERHEAD           64
#define PKT_DESC_AOL_OVERHEAD               64
#define PKT_DESC_OVERHEAD                   (PKT_DESC_SCRATCH_OVERHEAD + PKT_DESC_AOL_OVERHEAD) //  (scratch + AOL)
#define NVME_O_TCP_CH_SIZE                  8
#define NVME_O_TCP_PSH_CMD_CAPSULE_SIZE     64
#define NVME_O_TCP_HDGST_SIZE               4
#define NVME_O_TCP_CMD_CAPSULE_HEADER_SIZE  \
    (NVME_O_TCP_CH_SIZE + NVME_O_TCP_PSH_CMD_CAPSULE_SIZE + NVME_O_TCP_HDGST_SIZE)
#define NVME_O_TCP_CMD_CAPSULE_CH_PSH_SIZE (NVME_O_TCP_CH_SIZE + NVME_O_TCP_PSH_CMD_CAPSULE_SIZE)
#define NVME_O_TCP_DDGST_SIZE               4
#define TCP_PAGE_NVME_O_TCP_CH_OFFSET       128

//PDU types
#define NVME_O_TCP_PDU_TYPE_ICREQ           0
#define NVME_O_TCP_PDU_TYPE_ICRESP          1
#define NVME_O_TCP_PDU_TYPE_H2CTERMREQ      2
#define NVME_O_TCP_PDU_TYPE_C2HTERMREQ      3
#define NVME_O_TCP_PDU_TYPE_CAPSULECMD      4
#define NVME_O_TCP_PDU_TYPE_CAPSULERESP     5
#define NVME_O_TCP_PDU_TYPE_H2CDATA         6
#define NVME_O_TCP_PDU_TYPE_C2HDATA         7
#define NVME_O_TCP_PDU_TYPE_R2T             9

#endif //__NVME_COMMON_H
