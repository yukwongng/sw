//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// VNIC implementation state
///
//----------------------------------------------------------------------------

#ifndef __VNIC_IMPL_STATE_HPP__
#define __VNIC_IMPL_STATEHPP__

#include "nic/sdk/lib/table/directmap/directmap.hpp"
#include "nic/sdk/lib/table/hash/hash.hpp"
#include "nic/sdk/lib/table/slhash/slhash.hpp"
#include "nic/sdk/lib/table/sltcam/sltcam.hpp"
#include "nic/apollo/framework/api_base.hpp"
#include "nic/apollo/framework/state_base.hpp"
#include "nic/apollo/api/pds_state.hpp"

using sdk::table::sdk_table_factory_params_t;
using sdk::table::slhash;

namespace api {
namespace impl {


/// \defgroup PDS_VNIC_IMPL_STATE - VNIC state functionality
/// \ingroup PDS_VNIC
/// @{

// forward declaration
class vnic_impl;

/// \brief state maintained for VNICs
class vnic_impl_state : public state_base {
public:
    /// \brief constructor
    vnic_impl_state(pds_state *state);

    /// \brief destructor
    ~vnic_impl_state();

    /// \brief  API to initiate transaction over all the table manamgement
    ///         library instances
    /// \return SDK_RET_OK on success, failure status code on error
    sdk_ret_t table_transaction_begin(void);

    /// \brief  API to end transaction over all the table manamgement
    ///         library instances
    /// \return SDK_RET_OK on success, failure status code on error
    sdk_ret_t table_transaction_end(void);

    /// \brief     API to get table stats
    /// \param[in]  cb      callback to be called on stats
    ///             ctxt    opaque ctxt passed to the callback
    /// \return     SDK_RET_OK on success, failure status code on error
    sdk_ret_t table_stats(debug::table_stats_get_cb_t cb, void *ctxt);

private:
    indexer *vnic_idxr(void) { return vnic_idxr_; }
    slab *vnic_impl_slab(void) { return vnic_impl_slab_; }
    mem_hash *local_mapping_tbl(void) { return local_mapping_tbl_; }
    ///< vnic_impl class is friend of vnic_impl_state
    friend class vnic_impl;
    ///< mapping_impl class is friend of vnic_impl_state
    friend class mapping_impl;

private:
    // NOTE: there is no explicit table mgmt for rx and tx stats, we directly
    //       index using hw_id_ of vnic and and bzero out when we create vnic
    slab     *vnic_impl_slab_;
    ///< indexer to allocate hw vnic id
    indexer  *vnic_idxr_;
    mem_hash *local_mapping_tbl_;
};

/// \@}

}    // namespace impl
}    // namespace api

#endif    // __VNIC_IMPL_STATE_HPP__
