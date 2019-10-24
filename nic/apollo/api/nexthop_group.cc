//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// nexthop group handling
///
//----------------------------------------------------------------------------

#include "nic/sdk/include/sdk/base.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/core/mem.hpp"
#include "nic/apollo/api/nexthop.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/apollo/framework/api_ctxt.hpp"
#include "nic/apollo/framework/api_engine.hpp"

namespace api {

/// \defgroup PDS_NEXTHOP_GROUP - nexthop group functionality
/// \ingroup PDS_NEXTHOP
/// @{

nexthop_group::nexthop_group() {
    //SDK_SPINLOCK_INIT(&slock_, PTHREAD_PROCESS_PRIVATE);
    type_ = PDS_NHGROUP_TYPE_NONE;
    ht_ctxt_.reset();
    impl_ = NULL;
}

nexthop_group *
nexthop_group::factory(pds_nexthop_group_spec_t *spec) {
    nexthop_group *nh_group;

    ///< create nexthop group with defaults, if any
    nh_group = nexthop_group_db()->alloc();
    if (nh_group) {
        new (nh_group) nexthop_group();
        nh_group->impl_ = impl_base::factory(impl::IMPL_OBJ_ID_NEXTHOP_GROUP,
                                             spec);
        if (nh_group->impl_ == NULL) {
            nexthop_group::destroy(nh_group);
            return NULL;
        }
    }
    return nh_group;
}

nexthop_group::~nexthop_group() {
    // TODO: fix me
    //SDK_SPINLOCK_DESTROY(&slock_);
}

void
nexthop_group::destroy(nexthop_group *nh_group) {
    nh_group->nuke_resources_();
    impl_base::destroy(impl::IMPL_OBJ_ID_NEXTHOP_GROUP, nh_group->impl_);
    nh_group->~nexthop_group();
    nexthop_group_db()->free(nh_group);
}

sdk_ret_t
nexthop_group::init_config(api_ctxt_t *api_ctxt) {
    pds_nexthop_group_spec_t *spec = &api_ctxt->api_params->nexthop_group_spec;

    memcpy(&this->key_, &spec->key, sizeof(key_));
    type_ = spec->type;
    num_nexthops_ = spec->num_nexthops;
    return SDK_RET_OK;
}

sdk_ret_t
nexthop_group::reserve_resources(api_base *orig_obj, obj_ctxt_t *obj_ctxt) {
    return impl_->reserve_resources(this, obj_ctxt);
}

sdk_ret_t
nexthop_group::program_config(obj_ctxt_t *obj_ctxt) {
    PDS_TRACE_DEBUG("Programming nexthop group %u, type %u", key_, type_);
    return impl_->program_hw(this, obj_ctxt);
}

sdk_ret_t
nexthop_group::reprogram_config(api_op_t api_op) {
    return SDK_RET_ERR;
}

sdk_ret_t
nexthop_group::release_resources(void) {
    return impl_->release_resources(this);
}

sdk_ret_t
nexthop_group::nuke_resources_(void) {
    return impl_->nuke_resources(this);
}

sdk_ret_t
nexthop_group::cleanup_config(obj_ctxt_t *obj_ctxt) {
    return impl_->cleanup_hw(this, obj_ctxt);
}

sdk_ret_t
nexthop_group::update_config(api_base *orig_obj, obj_ctxt_t *obj_ctxt) {
    //return impl_->update_hw();
    return sdk::SDK_RET_INVALID_OP;
}

sdk_ret_t
nexthop_group::activate_config(pds_epoch_t epoch, api_op_t api_op,
                         obj_ctxt_t *obj_ctxt) {
    PDS_TRACE_DEBUG("Activating nexthop group %u", key_);
    return impl_->activate_hw(this, epoch, api_op, obj_ctxt);
}

sdk_ret_t
nexthop_group::reactivate_config(pds_epoch_t epoch, api_op_t api_op) {
    return sdk::SDK_RET_INVALID_OP;
}

sdk_ret_t
nexthop_group::read(pds_nexthop_group_key_t *key,
                    pds_nexthop_group_info_t *info) {
    return impl_->read_hw(this, (impl::obj_key_t *)key,
                          (impl::obj_info_t *)info);
    return sdk::SDK_RET_INVALID_OP;
}

sdk_ret_t
nexthop_group::update_db(api_base *orig_obj, obj_ctxt_t *obj_ctxt) {
    return sdk::SDK_RET_INVALID_OP;
}

sdk_ret_t
nexthop_group::add_to_db(void) {
    return nexthop_group_db()->insert(this);
}

sdk_ret_t
nexthop_group::del_from_db(void) {
    if (nexthop_group_db()->remove(this)) {
        return SDK_RET_OK;
    }
    return SDK_RET_ENTRY_NOT_FOUND;
}

sdk_ret_t
nexthop_group::delay_delete(void) {
    return delay_delete_to_slab(PDS_SLAB_ID_NEXTHOP_GROUP, this);
}

/// @}     // end of PDS_NEXTHOP

}    // namespace api