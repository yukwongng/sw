//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// lif implementation state maintenance
///
//----------------------------------------------------------------------------

#include "nic/sdk/lib/p4/p4_api.hpp"
// #include "gen/p4gen/apollo_txdma/include/apollo_txdma_p4pd.h"
#include "nic/apollo/api/impl/lif_impl_state.hpp"

/// \defgroup PDS_LIF_IMPL_STATE - lif state functionality
/// \ingroup PDS_LIF
/// \@{

#define PDS_MAX_LIFS    512

namespace api {
namespace impl {

lif_impl_state::lif_impl_state(pds_state *state) {
    p4pd_table_properties_t    tinfo;

    // maintain all lifs in the system in a hash table
    lif_ht_ = ht::factory(PDS_MAX_LIFS >> 2,
                          lif_impl::lif_key_func_get,
                          lif_impl::lif_hash_func_compute,
                          lif_impl::lif_key_func_compare);
    SDK_ASSERT(lif_ht_ != NULL);

#if 0
    p4pluspd_txdma_table_properties_get(
                P4_APOLLO_TXDMA_TBL_ID_TX_TABLE_S5_T4_LIF_RATE_LIMITER_TABLE,
                &tinfo);
    tx_rate_limiter_tbl_ =
        directmap::factory(tinfo.tablename,
                           P4_APOLLO_TXDMA_TBL_ID_TX_TABLE_S5_T4_LIF_RATE_LIMITER_TABLE,
                           tinfo.tabledepth, tinfo.actiondata_struct_size,
                           false, true, NULL);
    SDK_ASSERT(tx_rate_limiter_tbl_ != NULL);
#endif
}

lif_impl_state::~lif_impl_state() {
    ht::destroy(lif_ht_);
    // TODO: uncomment later when rate limter table comes
    //directmap::destroy(tx_rate_limiter_tbl_);
}

/// \@}

}    // namespace impl
}    // namespace api
