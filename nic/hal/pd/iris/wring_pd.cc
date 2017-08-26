#include <base.h>
#include <hal_lock.hpp>
#include <pd_api.hpp>
#include <wring_pd.hpp>
#include <capri_loader.h>
#include <capri_hbm.hpp>
#include <capri_common.h>
#include <p4plus_pd_api.h>

namespace hal {
namespace pd {

static pd_wring_meta_t g_meta[types::WRingType_MAX];

hal_ret_t  
wring_pd_meta_init() {
    /*
     * Add meta info for each ring
     * Fomat: is Global Ring, region name, # of slots in the ring, slot size in bytes
     *         associate object region name, associated size of object
     */

    g_meta[types::WRING_TYPE_SERQ] = 
        (pd_wring_meta_t) {false, "serq", 64, DEFAULT_WRING_SLOT_SIZE, "", 0, 0};
 
    g_meta[types::WRING_TYPE_NMDR_TX] = 
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_NMDR_TX, 16384, DEFAULT_WRING_SLOT_SIZE,
                                        CAPRI_HBM_REG_DESCRIPTOR_TX, 128,
                                        CAPRI_SEM_TNMDR_RAW_ADDR};
 
    g_meta[types::WRING_TYPE_NMDR_RX] = 
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_NMDR_RX, 16384, DEFAULT_WRING_SLOT_SIZE,
                                        CAPRI_HBM_REG_DESCRIPTOR_RX, 128,
                                        CAPRI_SEM_RNMDR_RAW_ADDR};
    
    g_meta[types::WRING_TYPE_NMPR_SMALL_TX] = 
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_NMPR_SMALL_TX, 16384, DEFAULT_WRING_SLOT_SIZE,
                                        CAPRI_HBM_REG_PAGE_SMALL_TX, 2048,
                                        CAPRI_SEM_TNMPR_SMALL_RAW_ADDR};

    g_meta[types::WRING_TYPE_NMPR_SMALL_RX] = 
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_NMPR_SMALL_RX, 16384, DEFAULT_WRING_SLOT_SIZE,
                                        CAPRI_HBM_REG_PAGE_SMALL_RX, 2048,
                                        CAPRI_SEM_RNMPR_SMALL_RAW_ADDR};

    g_meta[types::WRING_TYPE_NMPR_BIG_TX] = 
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_NMPR_BIG_TX, 16384, DEFAULT_WRING_SLOT_SIZE,
                                        CAPRI_HBM_REG_PAGE_BIG_TX, 9216,
                                        CAPRI_SEM_TNMPR_RAW_ADDR};

    g_meta[types::WRING_TYPE_NMPR_BIG_RX] = 
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_NMPR_BIG_RX, 16384, DEFAULT_WRING_SLOT_SIZE,
                                        CAPRI_HBM_REG_PAGE_BIG_RX, 9216,
                                        CAPRI_SEM_RNMPR_RAW_ADDR};

    g_meta[types::WRING_TYPE_BSQ] =
        (pd_wring_meta_t) {false, "bsq", 64, DEFAULT_WRING_SLOT_SIZE, "", 0, 0};

    g_meta[types::WRING_TYPE_BRQ] =
        (pd_wring_meta_t) {true, CAPRI_HBM_REG_BRQ, 1024, 128, "", 0, 0};

    return HAL_RET_OK;
}

hal_ret_t
get_default_slot_value(types::WRingType type, uint32_t index, uint8_t* value)
{
    pd_wring_meta_t     *meta = &g_meta[type];

    if(strlen(meta->obj_hbm_reg_name) <= 0) {
        memset(value, 0, meta->slot_size_in_bytes);
    } else {
        wring_hw_id_t obj_addr_base = get_start_offset(meta->obj_hbm_reg_name);
        if(obj_addr_base == 0) {
            HAL_TRACE_ERR("Failed to get the addr for the object");
            memset(value, 0, meta->slot_size_in_bytes);
            return HAL_RET_ENTRY_NOT_FOUND;
        }
        // get the addr of the object
        uint64_t obj_addr = obj_addr_base + (index * meta->obj_size);
        obj_addr = htonll(obj_addr);
        memcpy(value, &obj_addr, sizeof(obj_addr));
    }
    return HAL_RET_OK;
}

hal_ret_t
wring_pd_get_base_addr(types::WRingType type, uint32_t wring_id, wring_hw_id_t* wring_base)
{
    pd_wring_meta_t     *meta = &g_meta[type];
    wring_hw_id_t       addr = 0;

    addr = get_start_offset(meta->hbm_reg_name); 
    if(!addr) {
        HAL_TRACE_ERR("Could not find addr for {} region", meta->hbm_reg_name);
        return HAL_RET_ERR;
    }

    if(meta->is_global) {
        // Global rings start at the base of hbm_region
        *wring_base = addr;    
    } else {
        // Flow local ring. Get the offset based on wring id
        uint32_t wring_size = meta->num_slots * meta->slot_size_in_bytes; 
        *wring_base = addr + (wring_id * wring_size);
    }

    return HAL_RET_OK;
}

hal_ret_t
wring_pd_table_init(types::WRingType type, uint32_t wring_id) 
{
    hal_ret_t           ret = HAL_RET_OK;
    pd_wring_meta_t     *meta = &g_meta[type];
    wring_hw_id_t       wring_base;

    ret = wring_pd_get_base_addr(type, wring_id, &wring_base);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Could not find the wring base addr");
        return ret;
    }
    
    uint32_t reg_size = get_size_kb(meta->hbm_reg_name);
    if(!reg_size) {
        HAL_TRACE_ERR("Could not find size for {} region", meta->hbm_reg_name);
        return HAL_RET_ERR;
    }

    uint32_t required_size = meta->num_slots * meta->slot_size_in_bytes;
    HAL_ASSERT(reg_size * 1024 >= required_size);

    // Allocate memory for storing value for a slot
    uint8_t value[meta->slot_size_in_bytes];
    for(uint32_t  index = 0; index<meta->num_slots; index++) {

        uint64_t slot_addr = wring_base + (index * meta->slot_size_in_bytes);
        
        get_default_slot_value(type, index, value);
        p4plus_hbm_write(slot_addr, value, meta->slot_size_in_bytes);
    }
    if (meta->semaphore_addr) {
        // Initialize CI to ring size
        uint32_t val32;
        // HW full condition is PI + 1 == CI, so for a ring size of N,
        // need to set CI to N + 1
        val32 = htonl(meta->num_slots + 1);
        HAL_TRACE_DEBUG("writing {0} {1} to 0x{2:x}\n", meta->num_slots,
                                val32, meta->semaphore_addr +
                                CAPRI_SEM_INC_NOT_FULL_CI_OFFSET);
        p4plus_hbm_write(meta->semaphore_addr + CAPRI_SEM_INC_NOT_FULL_CI_OFFSET,
                                        (uint8_t *)&val32, sizeof (uint32_t));
    }
    return HAL_RET_OK;
}

static
hal_ret_t
p4pd_wring_get_entry(pd_wring_t* wring_pd) 
{
    hal_ret_t           ret = HAL_RET_OK;
    wring_t*            wring = wring_pd->wring;
    pd_wring_meta_t     *meta = &g_meta[wring->wring_type];
    wring_hw_id_t       wring_base;
   
    ret = wring_pd_get_base_addr(wring->wring_type, 
                                 wring->wring_id, 
                                 &wring_base);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Could not find the wring base addr");
        return ret;
    }
    
    // Normalize
    uint32_t slot_index = (wring_pd->wring->slot_index % meta->num_slots);
    wring_hw_id_t slot_addr = wring_base + (slot_index * meta->slot_size_in_bytes);

    HAL_TRACE_DEBUG("Reading from the addr: {}", slot_addr);

    uint8_t value[meta->slot_size_in_bytes];
    if(!p4plus_hbm_read(slot_addr, 
                        value, 
                        meta->slot_size_in_bytes)) {
        HAL_TRACE_ERR("Failed to read the data from the hw)");    
    }
    if(meta->slot_size_in_bytes == sizeof(uint64_t)) {
        wring->slot_value = ntohll(*(uint64_t *)value);
    } else {
        // TODO: Handle swizzle for arbiterary size
        memcpy(&wring->slot_value, value, sizeof(wring->slot_value));
    }
    return ret;
}

hal_ret_t
wring_pd_init_global_rings() 
{
    hal_ret_t   ret = HAL_RET_OK;
   
    ret = wring_pd_meta_init();
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to initialize meta");
        return HAL_RET_ERR; 
    }

    for(int i = 1; i< types::WRingType_ARRAYSIZE; i++) {
        if(g_meta[i].is_global) {
            ret = wring_pd_table_init(types::WRingType(i), 0);
            if(ret != HAL_RET_OK) {
                HAL_TRACE_ERR("Failed to initialize p4plus wring of type: {}", i); 
                // continue with the initialization of the remaining tables
                continue;
            } 
            HAL_TRACE_DEBUG("Initialized p4plus ring: {}", 
                types::WRingType_Name(types::WRingType(i)));
        }
    }
    return ret;
}

void *
wring_pd_get_hw_key_func (void *entry)
{
    HAL_ASSERT(entry != NULL);
    return (void *)&(((pd_wring_t *)entry)->hw_id);
}

uint32_t
wring_pd_compute_hw_hash_func (void *key, uint32_t ht_size)
{
    return hal::utils::hash_algo::fnv_hash(key, sizeof(wring_hw_id_t)) % ht_size;
}

bool
wring_pd_compare_hw_key_func (void *key1, void *key2)
{
    HAL_ASSERT((key1 != NULL) && (key2 != NULL));
    if (*(wring_hw_id_t *)key1 == *(wring_hw_id_t *)key2) {
        return true;
    }
    return false;
}

hal_ret_t
pd_wring_create (pd_wring_args_t *args)
{
    hal_ret_t               ret;
    pd_wring_s              *wring_pd;

    HAL_TRACE_DEBUG("WRING pd create");

    // allocate PD wring state
    wring_pd = wring_pd_alloc_init();
    if (wring_pd == NULL) {
        return HAL_RET_OOM;
    }
    HAL_TRACE_DEBUG("Alloc done");
    wring_pd->wring = args->wring;

    // TODO: get hw base  for this WRING
    //wring_pd->hw_id = pd_wring_get_base_hw_index(wring_pd);
    ret = wring_pd_table_init(wring_pd->wring->wring_type,
                              wring_pd->wring->wring_id);
    if(ret != HAL_RET_OK) {
        goto cleanup;    
    }

    // add to db
    ret = add_wring_pd_to_db(wring_pd);
    if (ret != HAL_RET_OK) {
       goto cleanup;
    }
    HAL_TRACE_DEBUG("DB add done");
    args->wring->pd = wring_pd;

    return HAL_RET_OK;

cleanup:
    if (wring_pd) {
        wring_pd_free(wring_pd);
    }
    return ret;
}

hal_ret_t
pd_wring_get (pd_wring_args_t *args)
{
    hal_ret_t               ret;
    pd_wring_t              wring_pd;

    HAL_TRACE_DEBUG("Wring pd get");

    // allocate PD wring state
    wring_pd_init(&wring_pd);

    wring_pd.wring = args->wring;

    // get hw wring entry
    ret = p4pd_wring_get_entry(&wring_pd);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to get wring entry"); 
        goto cleanup;    
    }
    HAL_TRACE_DEBUG("Get done");
cleanup:
    return ret;
}

}    // namespace pd
}    // namespace hal
