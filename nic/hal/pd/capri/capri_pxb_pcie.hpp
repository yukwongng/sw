/*
 * capri_pxb_pcie.hpp
 * Madhava Cheethirala (Pensando Systems)
 */

#ifndef __CAPRI_PXB_PCIE_HPP__
#define __CAPRI_PXB_PCIE_HPP__

#include <stdio.h>
#include <string>
#include <errno.h>
#include <stdlib.h>
#include <stdint.h>
#include "nic/include/base.h"

#define CAPRI_PCIE_MAX_LIFS 1024

#define CAPRI_PCIE_ATOMIC_PAGE_ID 0x3ffUL
#define CAPRI_PCIE_ATOMIC_REGION_ID 0xfUL
#define CAPRI_PCIE_ATOMIC_BASE_ADDR_GET(_region_id, _page_id) \
            ((1UL << 33) | (_region_id << 29) | (_page_id << 12) | 0x000)
#define CAPRI_PCIE_ATOMIC_BASE_ADDR \
            CAPRI_PCIE_ATOMIC_BASE_ADDR_GET(CAPRI_PCIE_ATOMIC_REGION_ID, CAPRI_PCIE_ATOMIC_PAGE_ID)


/** capri_pxb_pcie_init
 * API to init the pxb pcie module
 *
 * @return hal_ret_t: Status of the operation
 */
hal_ret_t capri_pxb_pcie_init();

#endif //__CAPRI_PCB_PCIE_HPP__
