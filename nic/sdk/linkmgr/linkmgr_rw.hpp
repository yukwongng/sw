//------------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
// 
// read-write APIs to interact with the ASIC
//------------------------------------------------------------------------------

#ifndef __LINKMGR_RW_HPP__
#define __LINKMGR_RW_HPP__

#include "sdk/base.hpp"
#include "sdk/pal.hpp"
#include "linkmgr_pd.hpp"

namespace sdk {
namespace linkmgr {
namespace pd {

extern uint32_t read_reg_base (uint32_t chip, uint64_t addr);
extern void write_reg_base(uint32_t chip, uint64_t addr, uint32_t  data);

// wrapper macros for read/write operations
#define READ_REG_BASE(chip, addr, data) {                                    \
    SDK_TRACE_DEBUG("PORT: %s: chip %d addr 0x%x",                           \
                    __FUNCTION__, chip, addr);                               \
    if (g_linkmgr_cfg.mock_hw_access == false) {                             \
        *data = read_reg_base(chip, addr);                                   \
    }                                                                        \
}

#define WRITE_REG_BASE(chip, addr, data) {                                   \
    SDK_TRACE_DEBUG("PORT: %s: chip %d addr 0x%x data 0x%x",                 \
                    __FUNCTION__, chip, addr, data);                         \
    if (g_linkmgr_cfg.mock_hw_access == false) {                             \
        write_reg_base(chip, addr, data);                                    \
    }                                                                        \
}

}    // namespace pd
}    // namespace linkmgr
}    // namespace sdk

#endif  // __LINKMGR_RW_HPP__

