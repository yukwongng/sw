//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// policer state handling
///
//----------------------------------------------------------------------------

#include "nic/apollo/core/mem.hpp"
#include "nic/apollo/api/policer_state.hpp"

namespace api {

/// \defgroup PDS_POLICER_STATE - policer database functionality
/// \ingroup PDS_POLICER
/// \@{

policer_state::policer_state() {
    policer_ht_ = ht::factory(PDS_MAX_POLICER >> 2,
                           policer::policer_key_func_get, policer::key_size());
    SDK_ASSERT(policer_ht_ != NULL);

    policer_slab_ = slab::factory("policer", PDS_SLAB_ID_POLICER,
                                  sizeof(policer), 16, true, true, true, NULL);
    SDK_ASSERT(policer_slab_ != NULL);
}

policer_state::~policer_state() {
    ht::destroy(policer_ht_);
    slab::destroy(policer_slab_);
}

policer *
policer_state::alloc(void) {
    return ((policer *)policer_slab_->alloc());
}

sdk_ret_t
policer_state::insert(policer *pol) {
    return policer_ht_->insert_with_key(&pol->key_, pol, &pol->ht_ctxt_);
}

policer *
policer_state::remove(policer *pol) {
    return (policer *)(policer_ht_->remove(&pol->key_));
}

void
policer_state::free(policer *pol) {
    policer_slab_->free(pol);
}

policer *
policer_state::find(pds_policer_key_t *key) const {
    return (policer *)(policer_ht_->lookup(key));
}

sdk_ret_t
policer_state::walk(state_walk_cb_t walk_cb, void *ctxt) {
    return policer_ht_->walk(walk_cb, ctxt);
}

sdk_ret_t
policer_state::slab_walk(state_walk_cb_t walk_cb, void *ctxt) {
    walk_cb(policer_slab_, ctxt);
    return SDK_RET_OK;
}

/// \@}

}    // namespace api