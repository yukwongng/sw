//
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// vpc state handling
///
//----------------------------------------------------------------------------

#include "nic/apollo/core/mem.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/api/vpc_state.hpp"

namespace api {

/// \defgroup PDS_VPC_STATE - vpc database functionality
/// \ingroup PDS_VPC
/// \@{

vpc_state::vpc_state() {
    // TODO: need to tune multi-threading related params later
    vpc_ht_ = ht::factory(PDS_MAX_VPC >> 1, vpc_entry::vpc_key_func_get,
                          vpc_entry::vpc_hash_func_compute,
                          vpc_entry::vpc_key_func_compare);
    SDK_ASSERT(vpc_ht_ != NULL);
    // we should accomodate one extra vpc of type provider/substrate/internet
    vpc_idxr_ = indexer::factory(PDS_MAX_VPC + 1);
    SDK_ASSERT(vpc_idxr_ != NULL);
    vpc_slab_ = slab::factory("vpc", PDS_SLAB_ID_VPC, sizeof(vpc_entry), 16,
                              true, true, NULL);
    SDK_ASSERT(vpc_slab_ != NULL);
}

vpc_state::~vpc_state() {
    ht::destroy(vpc_ht_);
    indexer::destroy(vpc_idxr_);
    slab::destroy(vpc_slab_);
}

vpc_entry *
vpc_state::alloc(void) {
    return ((vpc_entry *)vpc_slab_->alloc());
}

sdk_ret_t
vpc_state::insert(vpc_entry *vpc) {
    return vpc_ht_->insert_with_key(&vpc->key_, vpc, &vpc->ht_ctxt_);
}

vpc_entry *
vpc_state::remove(vpc_entry *vpc) {
    return (vpc_entry *)(vpc_ht_->remove(&vpc->key_));
}

void
vpc_state::free(vpc_entry *vpc) {
    vpc_slab_->free(vpc);
}

vpc_entry *
vpc_state::find(pds_vpc_key_t *key) const {
    return (vpc_entry *)(vpc_ht_->lookup(key));
}

/// \@}    // end of PDS_VPC_STATE

}    // namespace api
