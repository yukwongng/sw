// {C} Copyright 2018 Pensando Systems Inc. All rights reserved

#include "nic/include/base.hpp"
#include "nic/include/hal_mem.hpp"
#include "nic/hal/iris/include/hal_state.hpp"
#include "nic/hal/pd/capri/capri.hpp"
#include "nic/hal/pd/capri/capri_hbm.hpp"
#include "nic/hal/pd/capri/capri_config.hpp"
#include "nic/hal/pd/capri/capri_tbl_rw.hpp"
#include "nic/hal/pd/capri/capri_tm_rw.hpp"
#include "nic/hal/pd/capri/capri_txs_scheduler.hpp"
#include "nic/include/hal.hpp"
#include "nic/hal/svc/hal_ext.hpp"
#include "nic/include/asic_pd.hpp"
#include "nic/sdk/lib/p4/p4_api.hpp"
#include "nic/hal/pd/capri/capri_pxb_pcie.hpp"
#include "nic/hal/pd/capri/capri_state.hpp"
#include "nic/hal/pd/capri/capri_sw_phv.hpp"
#include "nic/hal/pd/capri/capri_barco_crypto.hpp"
#include "nic/include/capri_common.h"
#include "nic/asic/capri/verif/apis/cap_npv_api.h"
#include "nic/asic/capri/verif/apis/cap_dpa_api.h"
#include "nic/asic/capri/verif/apis/cap_pics_api.h"
#include "nic/asic/capri/verif/apis/cap_pict_api.h"
#include "nic/asic/capri/verif/apis/cap_ppa_api.h"
#include "nic/asic/capri/verif/apis/cap_prd_api.h"
#include "nic/asic/capri/verif/apis/cap_psp_api.h"
#include "nic/asic/capri/verif/apis/cap_ptd_api.h"
#include "nic/asic/capri/verif/apis/cap_stg_api.h"
#include "nic/asic/capri/verif/apis/cap_wa_api.h"
#include "nic/hal/pd/capri/capri_quiesce.hpp"
#include "nic/asic/capri/model/cap_top/cap_top_csr.h"
#include "nic/asic/capri/model/cap_prd/cap_prd_csr.h"
#include "nic/asic/capri/model/utils/cap_csr_py_if.h"

class capri_state_pd *g_capri_state_pd;
uint64_t capri_hbm_base;
uint64_t hbm_repl_table_offset;
uint32_t capri_coreclk_freq; //Mhz

/* capri_default_config_init
 * Load any bin files needed for initializing default configs
 */
static hal_ret_t
capri_default_config_init (capri_cfg_t *cfg)
{
    hal_ret_t   ret = HAL_RET_OK;
    std::string hbm_full_path;
    std::string full_path;
    int         num_phases = 2;
    int         i;

    for (i = 0; i < num_phases; i++) {
        full_path =  std::string(cfg->cfg_path) + "/init_bins/" + cfg->default_config_dir + "/init_" +
                                            std::to_string(i) + "_bin";

        HAL_TRACE_DEBUG("Init phase {} Binaries dir: {}", i, full_path.c_str());

        // Check if directory is present
        if (access(full_path.c_str(), R_OK) < 0) {
            HAL_TRACE_DEBUG("Skipping init binaries");
            return HAL_RET_OK;
        }

        ret = capri_load_config((char *)full_path.c_str());
        HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                                "Error loading init phase {} binaries ret {}",
                                i, ret);

        // Now do any polling for init complete for this phase
        capri_tm_hw_config_load_poll(i);
    }

    return ret;
}

static hal_ret_t
capri_timer_hbm_init (void)
{
    hal_ret_t ret = HAL_RET_OK;
    uint64_t timer_key_hbm_base_addr;
    uint64_t timer_key_hbm_addr;
    uint64_t zero_data[8] = { 0 };

    timer_key_hbm_base_addr = (uint64_t)get_start_offset((char *)JTIMERS);
    HAL_TRACE_DEBUG("HBM timer key base addr {:#x}", timer_key_hbm_base_addr);
    timer_key_hbm_addr = timer_key_hbm_base_addr;
    while (timer_key_hbm_addr < timer_key_hbm_base_addr +
                                CAPRI_TIMER_HBM_KEY_SPACE) {
        capri_hbm_write_mem(timer_key_hbm_addr, (uint8_t *)zero_data, sizeof(zero_data));
        timer_key_hbm_addr += sizeof(zero_data);
    }

    return ret;
}

static hal_ret_t
capri_pgm_init (capri_cfg_t *cfg)
{
    hal_ret_t      ret;
    std::string    full_path;

    for (uint8_t i = 0; i < cfg->num_pgm_cfgs; i++) {
        full_path =  std::string(cfg->cfg_path) + "/" + cfg->pgm_name +
            "/" + cfg->pgm_cfg[i].path;
        HAL_TRACE_DEBUG("Loading programs from dir {}", full_path.c_str());

        // check if directory is present
        if (access(full_path.c_str(), R_OK) < 0) {
            HAL_TRACE_ERR("{} not present/no read permissions",
                          full_path.c_str());
            return HAL_RET_ERR;
        }
        ret = capri_load_config((char *)full_path.c_str());
        if (ret != HAL_RET_OK) {
            HAL_TRACE_ERR("Failed to load config {}", full_path);
            return ret;
        }
    }

    return HAL_RET_OK;
}

static hal_ret_t
capri_asm_init (capri_cfg_t *cfg)
{
    int             iret = 0;
    uint64_t        base_addr;
    hal_ret_t       ret = HAL_RET_OK;
    std::string     full_path;
    uint32_t num_symbols = 0;
    sdk::platform::p4_prog_param_info_t *symbols = NULL;

    for (uint8_t i = 0; i < cfg->num_asm_cfgs; i++) {
        full_path =  std::string(cfg->cfg_path) + "/" + cfg->pgm_name +
            "/" + cfg->asm_cfg[i].path;
        HAL_TRACE_DEBUG("Loading ASM binaries from dir {}", full_path.c_str());

	// Check if directory is present
	if (access(full_path.c_str(), R_OK) < 0) {
            HAL_TRACE_ERR("{} not_present/no_read_permissions",
                full_path.c_str());
            return HAL_RET_ERR;
        }

        symbols = NULL;
        if(cfg->asm_cfg[i].symbols_func) {
            num_symbols = cfg->asm_cfg[i].symbols_func((void **)&symbols, cfg->platform);
        }

        base_addr = get_start_offset(cfg->asm_cfg[i].base_addr.c_str());
        HAL_TRACE_DEBUG("base addr {:#x}", base_addr);
        iret = sdk::platform::p4_load_mpu_programs(cfg->asm_cfg[i].name.c_str(),
           (char *)full_path.c_str(),
           base_addr,
           symbols,
           num_symbols,
           cfg->asm_cfg[i].sort_func);

       if(symbols)
           HAL_FREE(hal::HAL_MEM_ALLOC_PD, symbols);

       if (iret != 0) {
          HAL_TRACE_ERR("Failed to load program {}", full_path);
          return HAL_RET_ERR;
       }
   }
   return HAL_RET_OK;
}

static hal_ret_t
capri_hbm_regions_init (capri_cfg_t *cfg)
{
    hal_ret_t           ret = HAL_RET_OK;

    ret = capri_asm_init(cfg);
    if (ret != HAL_RET_OK) {
        return ret;
    }

    ret = capri_pgm_init(cfg);
    if (ret != HAL_RET_OK) {
        return ret;
    }

    ret = capri_timer_hbm_init();
    if (ret != HAL_RET_OK) {
        return ret;
    }

    // reset all the HBM regions that are marked for reset
    reset_hbm_regions(cfg);

    return ret;
}

static hal_ret_t
capri_cache_init (capri_cfg_t *cfg)
{
    hal_ret_t   ret = HAL_RET_OK;

    // Program Global parameters of the cache.
    ret = capri_hbm_cache_init(cfg);
    if (ret != HAL_RET_OK) {
        return ret;
    }

    // Process all the regions.
    ret = capri_hbm_cache_regions_init();
    if (ret != HAL_RET_OK) {
        return ret;
    }
    return ret;
}


/*
 * ASIC teams wants to disable Error recovery of Seq ID check pRDMA,
 * as this recovery path is not tested thoroughly, and we might be
 * runing into for an outstanding issue is suspicion.
 * Disabling from HAL for now, but Helen will commit this
 * into prd_asic_init api called from HAL, then this will be removed
 */
static hal_ret_t
capri_prd_init()
{
    cap_top_csr_t & cap0 = CAP_BLK_REG_MODEL_ACCESS(cap_top_csr_t, 0, 0);
    cap_pr_csr_t &pr_csr = cap0.pr.pr;

    pr_csr.prd.cfg_ctrl.read();
    pr_csr.prd.cfg_ctrl.pkt_phv_sync_err_recovery_en(0);
    pr_csr.prd.cfg_ctrl.write();
    HAL_TRACE_DEBUG("Disabled pkt_phv_sync_err_recovery_en in pr_prd_cfg_ctrl");
    return HAL_RET_OK;
}


static hal_ret_t
capri_repl_init (capri_cfg_t *cfg)
{
    capri_hbm_base = get_hbm_base();
    hbm_repl_table_offset = get_hbm_offset(JP4_REPL);
    if (hbm_repl_table_offset != INVALID_MEM_ADDRESS) {
        capri_tm_repl_table_base_addr_set(hbm_repl_table_offset / CAPRI_REPL_ENTRY_WIDTH);
        capri_tm_repl_table_token_size_set(cfg->repl_entry_width * 8);
    }

    return HAL_RET_OK;
}

typedef struct block_info_s {
    int  inst_count;
    void (*soft_reset)(int chip_id, int inst);
    void (*init_start)(int chip_id, int inst);
    void (*init_done) (int chip_id, int inst);
} block_info_t;

#define MAX_INIT_BLOCKS 10

block_info_t blocks_info[MAX_INIT_BLOCKS];

static void
capri_block_info_init(void)
{
    blocks_info[0].inst_count = 1;
    blocks_info[0].soft_reset = cap_npv_soft_reset;
    blocks_info[0].init_start = cap_npv_init_start;
    blocks_info[0].init_done  = cap_npv_init_done;

    blocks_info[1].inst_count = 2;
    blocks_info[1].soft_reset = cap_dpa_soft_reset;
    blocks_info[1].init_start = cap_dpa_init_start;
    blocks_info[1].init_done  = cap_dpa_init_done;

    blocks_info[2].inst_count = 4;
    blocks_info[2].soft_reset = cap_pics_soft_reset;
    blocks_info[2].init_start = cap_pics_init_start;
    blocks_info[2].init_done  = cap_pics_init_done;

    blocks_info[3].inst_count = 2;
    blocks_info[3].soft_reset = cap_pict_soft_reset;
    blocks_info[3].init_start = cap_pict_init_start;
    blocks_info[3].init_done  = cap_pict_init_done;

    blocks_info[4].inst_count = 2;
    blocks_info[4].soft_reset = cap_ppa_soft_reset;
    blocks_info[4].init_start = cap_ppa_init_start;
    blocks_info[4].init_done  = cap_ppa_init_done;

    blocks_info[5].inst_count = 1;
    blocks_info[5].soft_reset = cap_prd_soft_reset;
    blocks_info[5].init_start = cap_prd_init_start;
    blocks_info[5].init_done  = cap_prd_init_done;

    blocks_info[6].inst_count = 1;
    blocks_info[6].soft_reset = cap_psp_soft_reset;
    blocks_info[6].init_start = cap_psp_init_start;
    blocks_info[6].init_done  = cap_psp_init_done;

    blocks_info[7].inst_count = 1;
    blocks_info[7].soft_reset = cap_ptd_soft_reset;
    blocks_info[7].init_start = cap_ptd_init_start;
    blocks_info[7].init_done  = cap_ptd_init_done;

    blocks_info[8].inst_count = 28;
    blocks_info[8].soft_reset = cap_stg_soft_reset;
    blocks_info[8].init_start = cap_stg_init_start;
    blocks_info[8].init_done  = cap_stg_init_done;

    blocks_info[9].inst_count = 1;
    blocks_info[9].soft_reset = cap_wa_soft_reset;
    blocks_info[9].init_start = cap_wa_init_start;
    blocks_info[9].init_done  = cap_wa_init_done;

    return;
}

//------------------------------------------------------------------------------
// Reset all the capri blocks
//------------------------------------------------------------------------------
static hal_ret_t
capri_block_reset(capri_cfg_t *cfg)
{
    hal_ret_t    ret         = HAL_RET_OK;
    int          chip_id     = 0;
    block_info_t *block_info = NULL;

    for (int block = 0; block < MAX_INIT_BLOCKS; ++block) {
        block_info = &blocks_info[block];

        for(int inst = 0; inst < block_info->inst_count; ++inst) {
            block_info->soft_reset(chip_id, inst);
        }
    }

    return ret;
}

//------------------------------------------------------------------------------
// Start the init of blocks
//------------------------------------------------------------------------------
static hal_ret_t
capri_block_init_start(capri_cfg_t *cfg)
{
    hal_ret_t    ret         = HAL_RET_OK;
    int          chip_id     = 0;
    block_info_t *block_info = NULL;

    for (int block = 0; block < MAX_INIT_BLOCKS; ++block) {
        block_info = &blocks_info[block];

        for(int inst = 0; inst < block_info->inst_count; ++inst) {
            block_info->init_start(chip_id, inst);
        }
    }

    return ret;
}

//------------------------------------------------------------------------------
// Wait for the blocks to initialize
//------------------------------------------------------------------------------
static hal_ret_t
capri_block_init_done(capri_cfg_t *cfg)
{
    hal_ret_t    ret         = HAL_RET_OK;
    int          chip_id     = 0;
    block_info_t *block_info = NULL;

    for (int block = 0; block < MAX_INIT_BLOCKS; ++block) {
        block_info = &blocks_info[block];

        for(int inst = 0; inst < block_info->inst_count; ++inst) {
            block_info->init_done(chip_id, inst);
        }
    }

    return ret;
}

//------------------------------------------------------------------------------
// Init all the capri blocks owned by HAL
//------------------------------------------------------------------------------
hal_ret_t
capri_block_init(capri_cfg_t *cfg)
{
    hal_ret_t           ret = HAL_RET_OK;

    // initialize block info
    capri_block_info_init();

    // soft reset
    ret = capri_block_reset(cfg);
    if (ret != HAL_RET_OK) {
        return ret;
    }

    // init blocks
    ret = capri_block_init_start(cfg);
    if (ret != HAL_RET_OK) {
        return ret;
    }

    // wait for blocks to be inited
    ret = capri_block_init_done(cfg);
    if (ret != HAL_RET_OK) {
        return ret;
    }

    return ret;
}

//------------------------------------------------------------------------------
// perform all the CAPRI specific initialization
// - link all the P4 programs, by resolving symbols, labels etc.
// - load the P4/P4+ programs in HBM
// - do all the parser/deparser related register programming
// - do all the table configuration related register programming
//------------------------------------------------------------------------------
static hal_ret_t
capri_init (capri_cfg_t *cfg)
{
    hal_ret_t    ret;

    HAL_ASSERT_TRACE_RETURN((cfg != NULL), HAL_RET_INVALID_ARG, "Invalid cfg");
    HAL_TRACE_DEBUG("Initializing Capri");

    g_capri_state_pd = capri_state_pd::factory();
    HAL_ASSERT_TRACE_RETURN((g_capri_state_pd != NULL), HAL_RET_INVALID_ARG,
                            "Failed to instantiate Capri PD");

    g_capri_state_pd->set_cfg_path(cfg->cfg_path);
    if (capri_table_rw_init(cfg) != CAPRI_OK) {
        return HAL_RET_ERR;
    }

    ret = capri_hbm_regions_init(cfg);
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri HBM region init failure, err : {}", ret);

    ret = capri_block_init(cfg);
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri block region init failure, err : {}", ret);

    ret = capri_cache_init(cfg);
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri cache init failure, err : {}", ret);

    // do asic init before overwriting with the default configs
    ret = capri_tm_asic_init();
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri TM ASIC init failure, err : {}", ret);

    if (!cfg->catalog->qos_sw_init_enabled()) {
        ret = capri_default_config_init(cfg);
        HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                                "Capri default config init failure, err : {}",
                                ret);
    }

    ret = capri_txs_scheduler_init(cfg->admin_cos, cfg);
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri scheduler init failure, err : {}", ret);

    // Call PXB/PCIE init only in MODEL and RTL simulation
    // This will be done by PCIe manager for the actual chip
    if (cfg->platform == platform_type_t::PLATFORM_TYPE_SIM ||
        cfg->platform == platform_type_t::PLATFORM_TYPE_RTL) {
        ret = capri_pxb_pcie_init();
        HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                                "PXB/PCIE init failure, err : {}", ret);
    }

    ret = capri_tm_init(cfg->catalog);
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri TM init failure, err : {}", ret);

    ret = capri_repl_init(cfg);
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri replication init failure, err : {}", ret);

    ret = hal::pd::capri_sw_phv_init();
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri s/w phv init failure, err : {}", ret);

    if (!cfg->loader_info_file.empty()) {
        sdk::platform::p4_list_program_addr(
            cfg->cfg_path.c_str(), cfg->loader_info_file.c_str());
    }

    ret = hal::pd::capri_barco_crypto_init(cfg);
    if (ret != HAL_RET_OK) {
        // GFT always fails here !!!
        HAL_TRACE_ERR("Failed to inic barco_crypto. err: {}", ret);
        return ret;
    }

    ret = hal::pd::capri_quiesce_init();
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri quiesce init failure, err : {}", ret);

    ret = capri_prd_init();
    HAL_ASSERT_TRACE_RETURN((ret == HAL_RET_OK), ret,
                            "Capri PRD init failure, err : {}", ret);
    hal::svc::set_hal_status(hal::HAL_STATUS_ASIC_INIT_DONE);

    return ret;
}

//------------------------------------------------------------------------------
// perform all the CAPRI specific initialization
//------------------------------------------------------------------------------
hal_ret_t
hal::pd::asic_hbm_parse (asic_cfg_t *cfg)
{
    return hal_sdk_ret_to_hal_ret(capri_hbm_parse(cfg->cfg_path, cfg->pgm_name));
}

hal_ret_t
hal::pd::asic_init (asic_cfg_t *cfg)
{
    capri_cfg_t    capri_cfg;

    HAL_ASSERT(cfg != NULL);
    capri_cfg.loader_info_file = cfg->loader_info_file;
    capri_cfg.default_config_dir = cfg->default_config_dir;
    capri_cfg.cfg_path = cfg->cfg_path;
    capri_cfg.admin_cos = cfg->admin_cos;
    capri_cfg.repl_entry_width = cfg->repl_entry_width;
    capri_cfg.catalog = cfg->catalog;
    capri_cfg.p4_cache = true;
    capri_cfg.p4plus_cache = true;
    capri_cfg.llc_cache = true;
    capri_cfg.platform = cfg->platform;
    capri_cfg.num_pgm_cfgs = cfg->num_pgm_cfgs;
    capri_cfg.pgm_name = cfg->pgm_name;
    for (int i = 0; i < cfg->num_pgm_cfgs; i++) {
        capri_cfg.pgm_cfg[i].path = cfg->pgm_cfg[i].path;
    }
    capri_cfg.num_asm_cfgs = cfg->num_asm_cfgs;
    for (int i = 0; i < cfg->num_asm_cfgs; i++) {
        capri_cfg.asm_cfg[i].name = cfg->asm_cfg[i].name;
        capri_cfg.asm_cfg[i].path = cfg->asm_cfg[i].path;
        capri_cfg.asm_cfg[i].symbols_func = cfg->asm_cfg[i].symbols_func;
        capri_cfg.asm_cfg[i].base_addr = cfg->asm_cfg[i].base_addr;
        capri_cfg.asm_cfg[i].sort_func =
            cfg->asm_cfg[i].sort_func;
    }
    return capri_init(&capri_cfg);
}
