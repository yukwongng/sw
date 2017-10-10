#ifndef __IF_UTILS_HPP__
#define __IF_UTILS_HPP__

#include "nic/hal/src/tenant.hpp"
#include "nic/hal/src/l2segment.hpp"
#include "nic/hal/src/interface.hpp"
#include "nic/hal/src/lif.hpp"
#include "nic/proto/types.pb.h"
#include "nic/include/ip.h"

namespace hal {

hal_ret_t pltfm_get_port_from_front_port_num(uint32_t fp_num, 
                                             uint32_t *port_num);

// APIs
hal_ret_t tenant_add_l2seg (tenant_t *tenant, l2seg_t *l2seg);
hal_ret_t tenant_del_l2seg (tenant_t *tenant, l2seg_t *l2seg);

hal_ret_t l2seg_handle_nwsec_update (l2seg_t *l2seg, 
                                     nwsec_profile_t *nwsec_prof);

hal_ret_t if_handle_nwsec_update (l2seg_t *l2seg, if_t *hal_if, 
                                  nwsec_profile_t *nwsec_prof);
// Adding IFs to lif 
hal_ret_t lif_add_if (lif_t *lif, if_t *hal_if);
hal_ret_t lif_del_if (lif_t *lif, if_t *hal_if);

// Adding ifs to l2seg
hal_ret_t l2seg_add_if (l2seg_t *l2seg, if_t *hal_if);
hal_ret_t l2seg_del_if (l2seg_t *l2seg, if_t *hal_if);

// Adding l2segs to if
hal_ret_t if_add_l2seg (if_t *hal_if, l2seg_t *l2seg);
hal_ret_t if_del_l2seg (if_t *hal_if, l2seg_t *l2seg);

// Handle lif update in IF
hal_ret_t if_handle_lif_update (pd::pd_if_lif_upd_args_t *args);

}    // namespace hal

#endif    // __IF_UTILS_HPP__

