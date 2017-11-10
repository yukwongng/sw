#ifndef __HAL_STATE_PD_HPP__
#define __HAL_STATE_PD_HPP__

#include "nic/utils/indexer/indexer.hpp"
#include "nic/utils/slab/slab.hpp"
#include "nic/utils/ht/ht.hpp"
#include "nic/hal/pd/utils/directmap/directmap.hpp"
#include "nic/hal/pd/utils/hash/hash.hpp"
#include "nic/hal/pd/utils/tcam/tcam.hpp"
#include "nic/hal/pd/utils/flow/flow.hpp"
#include "nic/hal/pd/utils/met/met.hpp"
#include "nic/hal/pd/utils/acl_tcam/acl_tcam.hpp"
#include "nic/gen/iris/include/p4pd.h"
#include "nic/gen/common_txdma_actions/include/common_txdma_actions_p4pd.h"
#include "nic/gen/common_rxdma_actions/include/common_rxdma_actions_p4pd.h"
#include "nic/utils/bm_allocator/bm_allocator.hpp"


using hal::utils::indexer;
using hal::utils::slab;
using hal::utils::ht;
using hal::BMAllocator;
using hal::pd::utils::DirectMap;
using hal::pd::utils::Hash;
using hal::pd::utils::Tcam;
using hal::pd::utils::Flow;
using hal::pd::utils::Met;
using hal::pd::utils::acl_tcam;

namespace hal {
namespace pd {

// LIF HW ID Space for SB LIFs, Uplink Ifs/PCs
#define HAL_MAX_HW_LIFS         1025        
#define HAL_MAX_UPLINK_IFS      16
#define HAL_MAX_UPLINK_IF_PCS   32      // Both Uplink IFs and PCs combined

#define HAL_RW_TABLE_SIZE               4096
#define HAL_TUNNEL_RW_TABLE_SIZE        1024

#define HAL_MAX_HW_TM_PORTS             12
#define HAL_MAX_HW_BUF_POOLS_PER_PORT   32
#define HAL_MAX_HW_OQUEUES_PER_PORT     32
#define HAL_HW_OQUEUE_NODE_TYPES        3
#define HAL_MAX_HW_INGRESS_POLICERS     2048
#define HAL_MAX_HW_EGRESS_POLICERS      2048

#define HAL_MAX_HW_ACLS                 512 

//-----------------------------------------------------------------------------
// class hal_state_pd
//  
//  - slab memory instances for all pd objects.
//  - indexers needed for pd
//-----------------------------------------------------------------------------
class hal_state_pd {
public:
    static hal_state_pd *factory(void);
    ~hal_state_pd();

    slab *get_slab(hal_slab_t slab_id);

    // get APIs for global variables
    uint32_t rwr_tbl_decap_vlan_idx(void) const { 
        return rewr_tbl_decap_vlan_idx_; }
    void set_rwr_tbl_decap_vlan_idx(uint32_t idx) { 
        rewr_tbl_decap_vlan_idx_ = idx; }

    uint32_t tnnl_rwr_tbl_encap_vlan_idx(void) const { 
        return tnnl_rewr_tbl_enc_vlan_idx_; }
    void set_tnnl_rwr_tbl_encap_vlan_idx(uint32_t idx) { 
        tnnl_rewr_tbl_enc_vlan_idx_ = idx; }

    // get APIs for vrf related state
    slab *vrf_slab(void) const { return vrf_slab_; }
    indexer *vrf_hwid_idxr(void) const { return vrf_hwid_idxr_; }

    // get APIs for port related state
    slab *port_slab(void) const { return port_slab_; }

    // get APIs for security related state
    indexer *nwsec_profile_hwid_idxr(void) const { return nwsec_profile_hwid_idxr_; }

    // get APIs for L2 segment related state
    slab *l2seg_slab(void) const { return l2seg_slab_; }
    ht *l2seg_hwid_ht(void) const { return l2seg_hwid_ht_; }
    indexer *l2seg_cpu_idxr(void) const { return l2seg_cpu_idxr_; }

    // get API for lport indexer
    indexer *lport_idxr(void) const { return lport_idxr_; }

    // get APIs for LIF related state
    slab *lif_pd_slab(void) const { return lif_pd_slab_; }
    indexer *lif_hwid_idxr(void) const { return lif_hwid_idxr_; }

    // get APIs for Uplinkif  related state
    slab *uplinkif_pd_slab(void) const { return uplinkif_pd_slab_; }
    indexer *uplinkifpc_hwid_idxr(void) const { return uplinkifpc_idxr_; }

    // get APIs for uplinkpc related state
    slab *uplinkpc_pd_slab(void) const { return uplinkpc_pd_slab_; }

    // get APIs for enicif  related state
    slab *enicif_pd_slab(void) const { return enicif_pd_slab_; }
    slab *if_l2seg_entry_slab(void) const { return if_l2seg_entry_slab_; }

    // get APIs for cpuif  related state
    slab *cpuif_pd_slab(void) const { return cpuif_pd_slab_; }

    // get APIs for TunnelIf related state
    slab *tunnelif_pd_slab(void) const { return tunnelif_pd_slab_; }

    // get APIs for LIF related state
    slab *ep_pd_slab(void) const { return ep_pd_slab_; }
    slab *ep_pd_ip_entry_slab(void) const { return ep_pd_ip_entry_slab_; }
    // indexer *rw_table_idxr(void) const { return rw_table_idxr_; }

    // get APIs for nwsec related state
    slab *nwsec_pd_slab(void) const { return nwsec_pd_slab_; }

    // get APIs for dos policy related state
    slab *dos_pd_slab(void) const { return dos_pd_slab_; }

    // get APIs for session related state
    slab *session_slab(void) const { return session_slab_; }

    // get APIs for TLS CB related state
    slab *tlscb_slab(void) const { return tlscb_slab_; }
    ht *tlscb_hwid_ht(void) const { return tlscb_hwid_ht_; }

    // get APIs for TCP CB related state
    slab *tcpcb_slab(void) const { return tcpcb_slab_; }
    ht *tcpcb_hwid_ht(void) const { return tcpcb_hwid_ht_; }

    // get APIs for buf-pool related state
    slab *buf_pool_pd_slab(void) const { return buf_pool_pd_slab_; }
    indexer *buf_pool_hwid_idxr(uint32_t port_num) const { return buf_pool_hwid_idxr_[port_num]; }

    // get APIs for Queue related state
    slab *queue_pd_slab(void) const { return queue_pd_slab_; }
    indexer *queue_hwid_idxr(uint32_t port_num, uint32_t node_type) const { return queue_hwid_idxr_[port_num][node_type]; }
 
    // get APIs for Policer related state
    slab *policer_pd_slab(void) const { return policer_pd_slab_; }
    indexer *ingress_policer_hwid_idxr(void) const { return ingress_policer_hwid_idxr_; }
    indexer *egress_policer_hwid_idxr(void) const { return egress_policer_hwid_idxr_; }

    // get APIs for Acl related state
    slab *acl_pd_slab(void) const { return acl_pd_slab_; }

    // get APIs for WRING related state
    slab *wring_slab(void) const { return wring_slab_; }
    ht *wring_hwid_ht(void) const { return wring_hwid_ht_; }

    // get APIs for Crypto state
    indexer *crypto_pd_keys_idxr(void) { return session_keys_idxr_; }
    indexer *crypto_asym_dma_descr_idxr(void) { return crypto_asym_dma_descr_idxr_; }
    indexer *crypto_sym_msg_descr_idxr(void) { return crypto_sym_msg_descr_idxr_; }
    indexer *hbm_mem_idxr(void) { return hbm_mem_idxr_; }
    indexer *crypto_asym_key_descr_idxr(void) { return crypto_asym_key_descr_idxr_; }
    
    // get APIs for IPSEC CB related state
    slab *ipseccb_slab(void) const { return ipseccb_slab_; }
    ht *ipseccb_hwid_ht(void) const { return ipseccb_hwid_ht_; }

    slab *ipseccb_decrypt_slab(void) const { return ipseccb_decrypt_slab_; }
    ht *ipseccb_decrypt_hwid_ht(void) const { return ipseccb_decrypt_hwid_ht_; }

    // get APIs for L4LB related state
    slab *l4lb_pd_slab(void) const { return l4lb_pd_slab_; }
    
    // get APIs for TCP CB related state
    slab *cpucb_slab(void) const { return cpucb_slab_; }
    ht *cpucb_hwid_ht(void) const { return cpucb_hwid_ht_; }

    // get APIs for Raw Redirect CB related state
    slab *rawrcb_slab(void) const { return rawrcb_slab_; }
    ht *rawrcb_hwid_ht(void) const { return rawrcb_hwid_ht_; }

    // get APIs for Raw Chain CB related state
    slab *rawccb_slab(void) const { return rawccb_slab_; }
    ht *rawccb_hwid_ht(void) const { return rawccb_hwid_ht_; }

    // get APIs for RW table related state
    slab *rw_entry_slab(void) const { return rw_entry_slab_; }
    ht *rw_table_ht(void) const { return rw_table_ht_; }
    indexer *rw_tbl_idxr(void) { return rw_tbl_idxr_; }

    // get APIs for Tunnel RW table related state
    slab *tnnl_rw_entry_slab(void) const { return tnnl_rw_entry_slab_; }
    ht *tnnl_rw_table_ht(void) const { return tnnl_rw_table_ht_; }
    indexer *tnnl_rw_tbl_idxr(void) { return tnnl_rw_tbl_idxr_; }

    // get APIs for utils slabs
    slab *directmap_entry_slab(void) const { return directmap_entry_slab_; }

    // get APIs for CPU PKT related state
    slab *cpupkt_slab(void) const { return cpupkt_slab_; }
    slab *cpupkt_qinst_info_slab(void) const {return cpupkt_qinst_info_slab_; }
    indexer *cpupkt_descr_hw_id_idxr(void) {return cpupkt_descr_hwid_idxr_; }
    indexer *cpupkt_page_hw_id_idxr(void) {return cpupkt_page_hwid_idxr_; }

    // get APIs for TXS scheduler related state
    BMAllocator *txs_scheduler_map_idxr(void) { return txs_scheduler_map_idxr_; }    

    hal_ret_t init_tables(void);
    hal_ret_t p4plus_rxdma_init_tables(void);
    hal_ret_t p4plus_txdma_init_tables(void);
    DirectMap *dm_table(p4pd_table_id tid) const {
        if ((tid < P4TBL_ID_INDEX_MIN) || (tid > P4TBL_ID_INDEX_MAX)) {
            return NULL;
        }
        return dm_tables_[tid - P4TBL_ID_INDEX_MIN];
    }

    Hash *hash_tcam_table(p4pd_table_id tid) const {
        if ((tid < P4TBL_ID_HASH_OTCAM_MIN) ||
            (tid > P4TBL_ID_HASH_OTCAM_MAX)) {
            return NULL;
        }
        return hash_tcam_tables_[tid - P4TBL_ID_HASH_OTCAM_MIN];
    }

    Tcam *tcam_table(p4pd_table_id tid) const {
        if ((tid < P4TBL_ID_TCAM_MIN) || (tid > P4TBL_ID_TCAM_MAX)) {
            return NULL;
        }
        return tcam_tables_[tid - P4TBL_ID_TCAM_MIN];
    }

    Flow *flow_table(void) const {
        return flow_table_;
    }

    Met *met_table(void) const {
        return met_table_;
    }

    acl_tcam *acl_table(void) const {
        return acl_table_;
    }

    DirectMap *p4plus_rxdma_dm_table(p4pd_common_rxdma_actions_table_id tid) const {
        if ((tid < P4_COMMON_RXDMA_ACTIONS_TBL_ID_INDEX_MIN) ||
            (tid > P4_COMMON_RXDMA_ACTIONS_TBL_ID_INDEX_MAX)) {
            return NULL;
        }
        return p4plus_rxdma_dm_tables_[tid - P4_COMMON_RXDMA_ACTIONS_TBL_ID_INDEX_MIN];
    }

    DirectMap *p4plus_txdma_dm_table(p4pd_common_txdma_actions_table_id tid) const {
        if ((tid < P4_COMMON_TXDMA_ACTIONS_TBL_ID_INDEX_MIN) ||
            (tid > P4_COMMON_TXDMA_ACTIONS_TBL_ID_INDEX_MAX)) {
            return NULL;
        }
        return p4plus_txdma_dm_tables_[tid - P4_COMMON_TXDMA_ACTIONS_TBL_ID_INDEX_MIN];
    }

private:
    bool init(void);
    hal_state_pd();

    // global variables
    struct {
        uint32_t  rewr_tbl_decap_vlan_idx_;
        uint32_t  tnnl_rewr_tbl_enc_vlan_idx_;
    } __PACK__;

    // vrf related state
    struct {
        slab       *vrf_slab_;
        indexer    *vrf_hwid_idxr_;
    } __PACK__;

    // port related state
    struct {
        slab       *port_slab_;
    } __PACK__;

    struct {
        indexer    *nwsec_profile_hwid_idxr_;
    } __PACK__;

    // l2seg related state
    struct {
        slab       *l2seg_slab_;
        ht         *l2seg_hwid_ht_;
        indexer    *l2seg_cpu_idxr_;
    } __PACK__;

    // Lport Indexer
    struct {
        indexer    *lport_idxr_;
    } __PACK__;

    // LIF related state
    struct {
        slab       *lif_pd_slab_;
        indexer    *lif_hwid_idxr_;         // Used even by Uplink IF/PCs
    } __PACK__;

    // Uplink IF related state
    struct {
        slab       *uplinkif_pd_slab_;
        slab       *uplinkpc_pd_slab_;
        indexer    *uplinkifpc_idxr_;       // User for both Uplink IF/PCs
    } __PACK__;

    // Enic IF related state
    struct {
        slab       *enicif_pd_slab_;
        slab       *if_l2seg_entry_slab_;
    } __PACK__;

    // CPU IF related state
    struct {
        slab       *cpuif_pd_slab_;
    } __PACK__;

    // Tunnel IF related state
    struct {
        slab       *tunnelif_pd_slab_;
    } __PACK__;

    // EP related state
    struct {
        slab       *ep_pd_slab_;
        slab       *ep_pd_ip_entry_slab_;
        // indexer    *rw_table_idxr_;
    } __PACK__;

    // nwsec related state
    struct {
        slab       *nwsec_pd_slab_;
    } __PACK__;

    // dos policy related state
    struct {
        slab       *dos_pd_slab_;
    } __PACK__;

    // session related state
    struct {
        slab       *session_slab_;
    } __PACK__;

    // tlscb related state
    struct {
        slab       *tlscb_slab_;
        ht         *tlscb_hwid_ht_;
    } __PACK__;

    // tcpcb related state
    struct {
        slab       *tcpcb_slab_;
        ht         *tcpcb_hwid_ht_;
    } __PACK__;

    // Buf-Pool related state
    struct {
        slab       *buf_pool_pd_slab_;
        indexer    *buf_pool_hwid_idxr_[HAL_MAX_HW_TM_PORTS];
    } __PACK__;

    // Queue related state
    struct {
        slab       *queue_pd_slab_;
        indexer    *queue_hwid_idxr_[HAL_MAX_HW_TM_PORTS][HAL_HW_OQUEUE_NODE_TYPES];
    } __PACK__;

    // Policer related state
    struct {
        slab       *policer_pd_slab_;
        indexer    *ingress_policer_hwid_idxr_;
        indexer    *egress_policer_hwid_idxr_;
    } __PACK__;

    // Acl related state
    struct {
        slab       *acl_pd_slab_;
    } __PACK__;

    // wring related state
    struct {
        slab       *wring_slab_;
        ht         *wring_hwid_ht_;
    } __PACK__;

    // Crypto related state
    struct {
        indexer    *session_keys_idxr_;
        indexer    *crypto_asym_dma_descr_idxr_;
        indexer    *crypto_sym_msg_descr_idxr_;
        indexer    *hbm_mem_idxr_;
        indexer    *crypto_asym_key_descr_idxr_;
    } __PACK__;
    
    // ipseccb related state
    struct {
        slab       *ipseccb_slab_;
        ht         *ipseccb_hwid_ht_;
        slab       *ipseccb_decrypt_slab_;
        ht         *ipseccb_decrypt_hwid_ht_;
    } __PACK__;

    // l4lb related state
    struct {
        slab       *l4lb_pd_slab_;
    } __PACK__;

    // rw table management
    struct {
        slab      *rw_entry_slab_;
        ht        *rw_table_ht_;   
        indexer   *rw_tbl_idxr_;
    } __PACK__;

    // tnnl rw table management
    struct {
        slab      *tnnl_rw_entry_slab_;
        ht        *tnnl_rw_table_ht_;   
        indexer   *tnnl_rw_tbl_idxr_;
    } __PACK__;
    
    // cpucb related state
    struct {
        slab       *cpucb_slab_;
        ht         *cpucb_hwid_ht_;
    } __PACK__;

    // cpupkt related state
    struct {
        slab       *cpupkt_slab_;
        slab       *cpupkt_qinst_info_slab_;
        indexer    *cpupkt_descr_hwid_idxr_;
        indexer    *cpupkt_page_hwid_idxr_;
    } __PACK__;

    // Raw Redirect CB related state
    struct {
        slab       *rawrcb_slab_;
        ht         *rawrcb_hwid_ht_;
    } __PACK__;

    // TXS scheduler related state
    struct {
        BMAllocator    *txs_scheduler_map_idxr_;
    } __PACK__;

    // Raw Chain CB related state
    struct {
        slab       *rawccb_slab_;
        ht         *rawccb_hwid_ht_;
    } __PACK__;

    // utils slabs
    struct {
        slab       *directmap_entry_slab_;
    } __PACK__;

    DirectMap    **dm_tables_;
    Hash         **hash_tcam_tables_;
    Tcam         **tcam_tables_;
    Flow         *flow_table_;
    Met          *met_table_;
    acl_tcam     *acl_table_;
    DirectMap    **p4plus_rxdma_dm_tables_;
    DirectMap    **p4plus_txdma_dm_tables_;
};

extern class hal_state_pd    *g_hal_state_pd;

} // namespace pd
} // namespace hal

#endif    // __HAL_STATE_PD_HPP__

