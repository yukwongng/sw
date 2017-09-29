#include "nic/include/base.h"
#include <arpa/inet.h>
#include "nic/include/hal_lock.hpp"
#include "nic/include/pd_api.hpp"
#include "nic/hal/pd/iris/ipseccb_pd.hpp"
#include "nic/hal/pd/capri/capri_loader.h"
#include "nic/hal/pd/capri/capri_hbm.hpp"
#include "nic/hal/pd/iris/wring_pd.hpp"
#include "nic/hal/src/proxy.hpp"
#include "nic/hal/hal.hpp"
#include "nic/hal/src/lif_manager.hpp"
#include "nic/gen/esp_v4_tunnel_n2h_rxdma/include/esp_v4_tunnel_n2h_rxdma_p4plus_ingress.h"
#include "nic/hal/pd/iris/p4plus_pd_api.h"

namespace hal {
namespace pd {

hal_ret_t p4pd_get_ipsec_decrypt_tx_stage0_prog_addr(uint64_t* offset);

void *
ipseccb_pd_decrypt_get_hw_key_func (void *entry)
{
    HAL_ASSERT(entry != NULL);
    return (void *)&(((pd_ipseccb_decrypt_t *)entry)->hw_id);
}

uint32_t
ipseccb_pd_decrypt_compute_hw_hash_func (void *key, uint32_t ht_size)
{
    return hal::utils::hash_algo::fnv_hash(key, sizeof(ipseccb_hw_id_t)) % ht_size;
}

bool
ipseccb_pd_decrypt_compare_hw_key_func (void *key1, void *key2)
{
    HAL_ASSERT((key1 != NULL) && (key2 != NULL));
    if (*(ipseccb_hw_id_t *)key1 == *(ipseccb_hw_id_t *)key2) {
        return true;
    }
    return false;
}

/********************************************
 * RxDMA
 * ******************************************/

hal_ret_t
p4pd_get_ipsec_decrypt_rx_stage0_prog_addr(uint64_t* offset)
{
    char progname[] = "rxdma_stage0.bin";
    char labelname[]= "ipsec_rx_n2h_stage0";

    int ret = capri_program_label_to_offset("p4plus",
                                            progname,
                                            labelname,
                                            offset);
    if(ret < 0) {
        return HAL_RET_HW_FAIL;
    }
    *offset >>= MPU_PC_ADDR_SHIFT;
    return HAL_RET_OK;
}

static hal_ret_t 
p4pd_add_or_del_ipsec_decrypt_rx_stage0_entry(pd_ipseccb_decrypt_t* ipseccb_pd, bool del)
{
    common_p4plus_stage0_app_header_table_d     data = {0};
    hal_ret_t                                   ret = HAL_RET_OK;
    uint64_t                                    pc_offset = 0;
    uint64_t                                    ipsec_cb_ring_base;
    uint64_t                                    ipsec_cb_ring_addr;

    // hardware index for this entry
    ipseccb_hw_id_t hwid = ipseccb_pd->hw_id + 
        (P4PD_IPSECCB_STAGE_ENTRY_OFFSET * P4PD_HWID_IPSEC_RX_STAGE0);

    if(!del) {
        // get pc address
        if(p4pd_get_ipsec_decrypt_rx_stage0_prog_addr(&pc_offset) != HAL_RET_OK) {
            HAL_TRACE_ERR("Failed to get pc address");
            ret = HAL_RET_HW_FAIL;
        }
        data.action_id = pc_offset;
        HAL_TRACE_DEBUG("Received Rx Decrypt pc address {}", pc_offset);
        if (p4pd_get_ipsec_decrypt_tx_stage0_prog_addr(&pc_offset) != HAL_RET_OK) {
            HAL_TRACE_ERR("Failed to get pc address");
        }
        HAL_TRACE_DEBUG("Received Decrypt TX pc address {}", pc_offset);
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.total = 2;
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.iv_salt = ipseccb_pd->ipseccb->iv_salt;
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.iv_size = ipseccb_pd->ipseccb->iv_size;
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.block_size = ipseccb_pd->ipseccb->block_size;
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.icv_size = ipseccb_pd->ipseccb->icv_size;
        //data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.barco_enc_cmd = ipseccb_pd->ipseccb->barco_enc_cmd;
        // for now aes-decrypt-encoding hard-coded.
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.barco_enc_cmd = htonl(0x30100000);
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.expected_seq_no = ipseccb_pd->ipseccb->esn_lo;
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.key_index = ipseccb_pd->ipseccb->key_index;
  
        // the below may have to use a different range for the reverse direction 
        ipsec_cb_ring_base = get_start_offset(CAPRI_HBM_REG_IPSECCB);
        ipsec_cb_ring_addr = ipsec_cb_ring_base+(ipseccb_pd->ipseccb->cb_id * IPSEC_CB_RING_ENTRY_SIZE);
        HAL_TRACE_DEBUG("Ring base in Decrypt 0x{0:x} CB Ring Addr 0x{0:x}", ipsec_cb_ring_base, ipsec_cb_ring_addr);

        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.cb_ring_base_addr = htonll(ipsec_cb_ring_addr);
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.cb_cindex = 0;
        data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.cb_pindex = 0;
         
    }
    HAL_TRACE_DEBUG("Programming Decrypt stage0 at hw-id: 0x{0:x}", hwid); 
    if(!p4plus_hbm_write(hwid,  (uint8_t *)&data, sizeof(data))){
        HAL_TRACE_ERR("Failed to create rx: stage0 entry for IPSECCB");
        ret = HAL_RET_HW_FAIL;
    }
    return ret;
}

hal_ret_t 
p4pd_add_or_del_ipseccb_decrypt_rxdma_entry(pd_ipseccb_decrypt_t* ipseccb_pd, bool del)
{
    hal_ret_t   ret = HAL_RET_OK;

    ret = p4pd_add_or_del_ipsec_decrypt_rx_stage0_entry(ipseccb_pd, del);
    if(ret != HAL_RET_OK) {
        goto cleanup;
    }

    return HAL_RET_OK;
cleanup:
    /* TODO: CLEANUP */
    return ret;
}

hal_ret_t 
p4pd_get_ipsec_decrypt_rx_stage0_entry(pd_ipseccb_decrypt_t* ipseccb_pd)
{
    common_p4plus_stage0_app_header_table_d data = {0};
    uint64_t ipsec_cb_ring_addr;
    uint8_t cb_cindex, cb_pindex;
    uint64_t replay_seq_no_bmp = 0;
    uint32_t expected_seq_no, last_replay_seq_no;

    // hardware index for this entry
    ipseccb_hw_id_t hwid = ipseccb_pd->hw_id + 
        (P4PD_IPSECCB_STAGE_ENTRY_OFFSET * P4PD_HWID_IPSEC_RX_STAGE0);

    if(!p4plus_hbm_read(hwid,  (uint8_t *)&data, sizeof(data))){
        HAL_TRACE_ERR("Failed to get rx: stage0 entry for IPSEC CB");
        return HAL_RET_HW_FAIL;
    }
    
    ipseccb_pd->ipseccb->iv_salt = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.iv_salt;
    HAL_TRACE_DEBUG("Got salt {}", ipseccb_pd->ipseccb->iv_salt);
    ipseccb_pd->ipseccb->iv_size = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.iv_size;
    HAL_TRACE_DEBUG("Got iv_size {}", ipseccb_pd->ipseccb->iv_size);
    ipseccb_pd->ipseccb->block_size = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.block_size;
    ipseccb_pd->ipseccb->icv_size = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.icv_size;
    ipseccb_pd->ipseccb->barco_enc_cmd = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.barco_enc_cmd;
    ipseccb_pd->ipseccb->key_index = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.key_index;
   
    ipsec_cb_ring_addr = ntohll(data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.cb_ring_base_addr);
    cb_cindex = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.cb_cindex;
    cb_pindex = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.cb_pindex;
    HAL_TRACE_DEBUG("CB Ring Addr 0x{0:x} Pindex {} CIndex {}", ipsec_cb_ring_addr, cb_pindex, cb_cindex);

    expected_seq_no = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.expected_seq_no;
    last_replay_seq_no = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.last_replay_seq_no;
    replay_seq_no_bmp = data.u.esp_v4_tunnel_n2h_rxdma_initial_table_d.replay_seq_no_bmp;
    HAL_TRACE_DEBUG("expected_seq_no: 0x{0:x} last_replay_seq_no: 0x{0:x} replay_seq_no_bmp: 0x{0:x}", 
                    expected_seq_no, last_replay_seq_no, replay_seq_no_bmp); 
    return HAL_RET_OK;
}

hal_ret_t 
p4pd_get_ipseccb_decrypt_rxdma_entry(pd_ipseccb_decrypt_t* ipseccb_pd)
{
    hal_ret_t   ret = HAL_RET_OK;
    
    ret = p4pd_get_ipsec_decrypt_rx_stage0_entry(ipseccb_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to get ipsec_rx entry");
        goto cleanup;
    }
    return HAL_RET_OK;
cleanup:
    /* TODO: CLEANUP */
    return ret;
}

/********************************************
 * TxDMA
 * ******************************************/
hal_ret_t
p4pd_get_ipsec_decrypt_tx_stage0_prog_addr(uint64_t* offset)
{
    char progname[] = "txdma_stage0.bin";
    char labelname[]= "ipsec_tx_n2h_stage0";

    int ret = capri_program_label_to_offset("p4plus",
                                            progname,
                                            labelname,
                                            offset);
    if(ret < 0) {
        return HAL_RET_HW_FAIL;
    }
    *offset >>= MPU_PC_ADDR_SHIFT;
    return HAL_RET_OK;
}

hal_ret_t 
p4pd_add_or_del_ipseccb_decrypt_txdma_entry(pd_ipseccb_decrypt_t* ipseccb_pd, bool del)
{
    return HAL_RET_OK;
}

hal_ret_t 
p4pd_get_ipseccb_decrypt_txdma_entry(pd_ipseccb_decrypt_t* ipseccb_pd)
{
    /* TODO */
    return HAL_RET_OK;
}

/**************************/

ipseccb_hw_id_t
pd_ipseccb_decrypt_get_base_hw_index(pd_ipseccb_decrypt_t* ipseccb_pd)
{
    HAL_ASSERT(NULL != ipseccb_pd);
    HAL_ASSERT(NULL != ipseccb_pd->ipseccb);
    
    // Get the base address of IPSEC CB from LIF Manager.
    // Set qtype and qid as 0 to get the start offset. 
    uint64_t offset = g_lif_manager->GetLIFQStateAddr(SERVICE_LIF_IPSEC_ESP, 1, 0);
    HAL_TRACE_DEBUG("received decrypt offset 0x{0:x}", offset);
    return offset + \
        (ipseccb_pd->ipseccb->cb_id * P4PD_HBM_IPSEC_CB_ENTRY_SIZE);
}

hal_ret_t
p4pd_add_or_del_ipseccb_decrypt_entry(pd_ipseccb_decrypt_t* ipseccb_pd, bool del) 
{
    hal_ret_t                   ret = HAL_RET_OK;
 
    ret = p4pd_add_or_del_ipseccb_decrypt_rxdma_entry(ipseccb_pd, del);
    if(ret != HAL_RET_OK) {
        goto err;    
    }
   
    ret = p4pd_add_or_del_ipseccb_decrypt_txdma_entry(ipseccb_pd, del);
    if(ret != HAL_RET_OK) {
        goto err;    
    }

err:
    return ret;
}

static
hal_ret_t
p4pd_get_ipseccb_decrypt_entry(pd_ipseccb_decrypt_t* ipseccb_pd) 
{
    hal_ret_t                   ret = HAL_RET_OK;
    
    ret = p4pd_get_ipseccb_decrypt_rxdma_entry(ipseccb_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to get rxdma entry for ipseccb");
        goto err;    
    }
   
    ret = p4pd_get_ipseccb_decrypt_txdma_entry(ipseccb_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to get txdma entry for ipseccb");
        goto err;    
    }

err:
    /*TODO: cleanup */
    return ret;
}

/********************************************
 * APIs
 *******************************************/

hal_ret_t
pd_ipseccb_decrypt_create (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    pd_ipseccb_decrypt_s              *ipseccb_pd;

    HAL_TRACE_DEBUG("Creating pd state for IPSEC CB.");

    // allocate PD ipseccb state
    ipseccb_pd = ipseccb_pd_decrypt_alloc_init();
    if (ipseccb_pd == NULL) {
        return HAL_RET_OOM;
    }
    HAL_TRACE_DEBUG("Alloc done");
    ipseccb_pd->ipseccb = args->ipseccb;
    // get hw-id for this IPSECCB
    ipseccb_pd->hw_id = pd_ipseccb_decrypt_get_base_hw_index(ipseccb_pd);
    printf("Received decrypt hw-id: 0x%lx ", ipseccb_pd->hw_id);
    
    // program ipseccb
    ret = p4pd_add_or_del_ipseccb_decrypt_entry(ipseccb_pd, false);
    if(ret != HAL_RET_OK) {
        goto cleanup;    
    }
    // add to db
    ret = add_ipseccb_pd_decrypt_to_db(ipseccb_pd);
    if (ret != HAL_RET_OK) {
       goto cleanup;
    }
    args->ipseccb->pd_decrypt = ipseccb_pd;

    return HAL_RET_OK;

cleanup:

    if (ipseccb_pd) {
        ipseccb_pd_decrypt_free(ipseccb_pd);
    }
    return ret;
}

hal_ret_t
pd_ipseccb_decrypt_update (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    
    if(!args) {
       return HAL_RET_INVALID_ARG; 
    }

    ipseccb_t*                ipseccb = args->ipseccb;
    pd_ipseccb_decrypt_t*             ipseccb_pd = (pd_ipseccb_decrypt_t*)ipseccb->pd_decrypt;

    HAL_TRACE_DEBUG("IPSECCB pd update");
    
    // program ipseccb
    ret = p4pd_add_or_del_ipseccb_decrypt_entry(ipseccb_pd, false);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to update ipseccb");
    }
    return ret;
}

hal_ret_t
pd_ipseccb_decrypt_delete (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    
    if(!args) {
       return HAL_RET_INVALID_ARG; 
    }

    ipseccb_t*                ipseccb = args->ipseccb;
    pd_ipseccb_decrypt_t*             ipseccb_pd = (pd_ipseccb_decrypt_t*)ipseccb->pd_decrypt;

    HAL_TRACE_DEBUG("IPSECCB pd delete");
    
    // program ipseccb
    ret = p4pd_add_or_del_ipseccb_decrypt_entry(ipseccb_pd, true);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to delete ipseccb entry"); 
    }
    
    del_ipseccb_pd_decrypt_from_db(ipseccb_pd);

    ipseccb_pd_decrypt_free(ipseccb_pd);

    return ret;
}

hal_ret_t
pd_ipseccb_decrypt_get (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    pd_ipseccb_decrypt_t              ipseccb_pd;

    HAL_TRACE_DEBUG("IPSECCB pd get for id: {}", args->ipseccb->cb_id);

    // allocate PD ipseccb state
    ipseccb_pd_decrypt_init(&ipseccb_pd);
    ipseccb_pd.ipseccb = args->ipseccb;
    
    // get hw-id for this IPSECCB
    ipseccb_pd.hw_id = pd_ipseccb_decrypt_get_base_hw_index(&ipseccb_pd);
    HAL_TRACE_DEBUG("Received hw-id 0x{0:x}", ipseccb_pd.hw_id);

    // get hw ipseccb entry
    ret = p4pd_get_ipseccb_decrypt_entry(&ipseccb_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Get request failed for id: 0x{0:x}", ipseccb_pd.ipseccb->cb_id);
    }
    return ret;
}

}    // namespace pd
}    // namespace hal
