#include <base.h>
#include <arpa/inet.h>
#include <hal_lock.hpp>
#include <pd_api.hpp>
#include <ipseccb_pd.hpp>
#include <capri_loader.h>
#include <capri_hbm.hpp>
#include <wring_pd.hpp>
#include <proxy.hpp>
#include <hal.hpp>
#include <lif_manager.hpp>
#include <ipsec_rxdma_actions_p4plus_ingress.h>
#include <p4plus_pd_api.h>

namespace hal {
namespace pd {

void *
ipseccb_pd_get_hw_key_func (void *entry)
{
    HAL_ASSERT(entry != NULL);
    return (void *)&(((pd_ipseccb_t *)entry)->hw_id);
}

uint32_t
ipseccb_pd_compute_hw_hash_func (void *key, uint32_t ht_size)
{
    return hal::utils::hash_algo::fnv_hash(key, sizeof(ipseccb_hw_id_t)) % ht_size;
}

bool
ipseccb_pd_compare_hw_key_func (void *key1, void *key2)
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
p4pd_get_ipsec_rx_stage0_prog_addr(uint64_t* offset)
{
    char progname[] = "rxdma_stage0.bin";
    char labelname[]= "ipsec_rx_stage0";

    int ret = capri_program_label_to_offset("p4plus",
                                            progname,
                                            labelname,
                                            offset);
    if(ret < 0) {
        return HAL_RET_HW_FAIL;
    }
    return HAL_RET_OK;
}

static hal_ret_t 
p4pd_add_or_del_ipsec_rx_stage0_entry(pd_ipseccb_t* ipseccb_pd, bool del)
{
    common_p4plus_stage0_app_header_table_d     data = {0};
    hal_ret_t                                   ret = HAL_RET_OK;
    uint64_t                                    pc_offset = 0;

    // hardware index for this entry
    ipseccb_hw_id_t hwid = ipseccb_pd->hw_id + 
        (P4PD_IPSECCB_STAGE_ENTRY_OFFSET * P4PD_HWID_IPSEC_RX_STAGE0);

    if(!del) {
        // get pc address
        if(p4pd_get_ipsec_rx_stage0_prog_addr(&pc_offset) != HAL_RET_OK) {
            HAL_TRACE_ERR("Failed to get pc address");
            ret = HAL_RET_HW_FAIL;
        }
        HAL_TRACE_DEBUG("Received pc address", pc_offset);
        data.action_id = pc_offset;
        //data.u.ipsec_encap_rxdma_initial_table_d.xxx = FFFF
    }
    HAL_TRACE_DEBUG("Programming stage0 at hw-id: 0x{0:x}", hwid); 
    if(!p4plus_hbm_write(hwid,  (uint8_t *)&data, sizeof(data))){
        HAL_TRACE_ERR("Failed to create rx: stage0 entry for IPSECCB");
        ret = HAL_RET_HW_FAIL;
    }
    return ret;
}

hal_ret_t 
p4pd_add_or_del_ipseccb_rxdma_entry(pd_ipseccb_t* ipseccb_pd, bool del)
{
    hal_ret_t   ret = HAL_RET_OK;

    ret = p4pd_add_or_del_ipsec_rx_stage0_entry(ipseccb_pd, del);
    if(ret != HAL_RET_OK) {
        goto cleanup;
    }

    return HAL_RET_OK;
cleanup:
    /* TODO: CLEANUP */
    return ret;
}

hal_ret_t 
p4pd_get_ipsec_rx_stage0_entry(pd_ipseccb_t* ipseccb_pd)
{
    common_p4plus_stage0_app_header_table_d data = {0};

    // hardware index for this entry
    ipseccb_hw_id_t hwid = ipseccb_pd->hw_id + 
        (P4PD_IPSECCB_STAGE_ENTRY_OFFSET * P4PD_HWID_IPSEC_RX_STAGE0);

    if(!p4plus_hbm_read(hwid,  (uint8_t *)&data, sizeof(data))){
        HAL_TRACE_ERR("Failed to get rx: stage0 entry for IPSEC CB");
        return HAL_RET_HW_FAIL;
    }
    return HAL_RET_OK;
}

hal_ret_t 
p4pd_get_ipseccb_rxdma_entry(pd_ipseccb_t* ipseccb_pd)
{
    hal_ret_t   ret = HAL_RET_OK;
    
    ret = p4pd_get_ipsec_rx_stage0_entry(ipseccb_pd);
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
p4pd_add_or_del_ipseccb_txdma_entry(pd_ipseccb_t* ipseccb_pd, bool del)
{
    return HAL_RET_OK;
}

hal_ret_t 
p4pd_get_ipseccb_txdma_entry(pd_ipseccb_t* ipseccb_pd)
{
    /* TODO */
    return HAL_RET_OK;
}

/**************************/

ipseccb_hw_id_t
pd_ipseccb_get_base_hw_index(pd_ipseccb_t* ipseccb_pd)
{
    HAL_ASSERT(NULL != ipseccb_pd);
    HAL_ASSERT(NULL != ipseccb_pd->ipseccb);
    
    // Get the base address of IPSEC CB from LIF Manager.
    // Set qtype and qid as 0 to get the start offset. 
    uint64_t offset = g_lif_manager->GetLIFQStateAddr(SERVICE_LIF_IPSEC_ESP, 0, 0);
    HAL_TRACE_DEBUG("received offset 0x{0:x}", offset);
    return offset + \
        (ipseccb_pd->ipseccb->cb_id * P4PD_HBM_IPSEC_CB_ENTRY_SIZE);
}

hal_ret_t
p4pd_add_or_del_ipseccb_entry(pd_ipseccb_t* ipseccb_pd, bool del) 
{
    hal_ret_t                   ret = HAL_RET_OK;
 
    ret = p4pd_add_or_del_ipseccb_rxdma_entry(ipseccb_pd, del);
    if(ret != HAL_RET_OK) {
        goto err;    
    }
   
    ret = p4pd_add_or_del_ipseccb_txdma_entry(ipseccb_pd, del);
    if(ret != HAL_RET_OK) {
        goto err;    
    }

err:
    return ret;
}

static
hal_ret_t
p4pd_get_ipseccb_entry(pd_ipseccb_t* ipseccb_pd) 
{
    hal_ret_t                   ret = HAL_RET_OK;
    
    ret = p4pd_get_ipseccb_rxdma_entry(ipseccb_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to get rxdma entry for ipseccb");
        goto err;    
    }
   
    ret = p4pd_get_ipseccb_txdma_entry(ipseccb_pd);
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
pd_ipseccb_create (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    pd_ipseccb_s              *ipseccb_pd;

    HAL_TRACE_DEBUG("Creating pd state for IPSEC CB.");

    // allocate PD ipseccb state
    ipseccb_pd = ipseccb_pd_alloc_init();
    if (ipseccb_pd == NULL) {
        return HAL_RET_OOM;
    }
    HAL_TRACE_DEBUG("Alloc done");
    ipseccb_pd->ipseccb = args->ipseccb;
    // get hw-id for this IPSECCB
    ipseccb_pd->hw_id = pd_ipseccb_get_base_hw_index(ipseccb_pd);
    printf("Received hw-id: 0x%lx ", ipseccb_pd->hw_id);
    
    // program ipseccb
    ret = p4pd_add_or_del_ipseccb_entry(ipseccb_pd, false);
    if(ret != HAL_RET_OK) {
        goto cleanup;    
    }
    // add to db
    ret = add_ipseccb_pd_to_db(ipseccb_pd);
    if (ret != HAL_RET_OK) {
       goto cleanup;
    }
    args->ipseccb->pd = ipseccb_pd;

    return HAL_RET_OK;

cleanup:

    if (ipseccb_pd) {
        ipseccb_pd_free(ipseccb_pd);
    }
    return ret;
}

hal_ret_t
pd_ipseccb_update (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    
    if(!args) {
       return HAL_RET_INVALID_ARG; 
    }

    ipseccb_t*                ipseccb = args->ipseccb;
    pd_ipseccb_t*             ipseccb_pd = (pd_ipseccb_t*)ipseccb->pd;

    HAL_TRACE_DEBUG("IPSECCB pd update");
    
    // program ipseccb
    ret = p4pd_add_or_del_ipseccb_entry(ipseccb_pd, false);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to update ipseccb");
    }
    return ret;
}

hal_ret_t
pd_ipseccb_delete (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    
    if(!args) {
       return HAL_RET_INVALID_ARG; 
    }

    ipseccb_t*                ipseccb = args->ipseccb;
    pd_ipseccb_t*             ipseccb_pd = (pd_ipseccb_t*)ipseccb->pd;

    HAL_TRACE_DEBUG("IPSECCB pd delete");
    
    // program ipseccb
    ret = p4pd_add_or_del_ipseccb_entry(ipseccb_pd, true);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to delete ipseccb entry"); 
    }
    
    del_ipseccb_pd_from_db(ipseccb_pd);

    ipseccb_pd_free(ipseccb_pd);

    return ret;
}

hal_ret_t
pd_ipseccb_get (pd_ipseccb_args_t *args)
{
    hal_ret_t               ret;
    pd_ipseccb_t              ipseccb_pd;

    HAL_TRACE_DEBUG("IPSECCB pd get for id: {}", args->ipseccb->cb_id);

    // allocate PD ipseccb state
    ipseccb_pd_init(&ipseccb_pd);
    ipseccb_pd.ipseccb = args->ipseccb;
    
    // get hw-id for this IPSECCB
    ipseccb_pd.hw_id = pd_ipseccb_get_base_hw_index(&ipseccb_pd);
    HAL_TRACE_DEBUG("Received hw-id 0x{0:x}", ipseccb_pd.hw_id);

    // get hw ipseccb entry
    ret = p4pd_get_ipseccb_entry(&ipseccb_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Get request failed for id: 0x{0:x}", ipseccb_pd.ipseccb->cb_id);
    }
    return ret;
}

}    // namespace pd
}    // namespace hal
