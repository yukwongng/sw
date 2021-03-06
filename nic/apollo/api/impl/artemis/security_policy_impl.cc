//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// datapath implementation of security policy
///
//----------------------------------------------------------------------------

#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/core/mem.hpp"
#include "nic/apollo/framework/api_engine.hpp"
#include "nic/apollo/framework/api_params.hpp"
#include "nic/apollo/api/policy.hpp"
#include "nic/apollo/api/impl/artemis/security_policy_impl.hpp"
#include "nic/apollo/api/impl/artemis/pds_impl_state.hpp"
#include "nic/apollo/api/impl/rfc/rfc.hpp"

namespace api {
namespace impl {

/// \defgroup PDS_SECURITY_POLICY_IMPL - security policy datapath implementation
/// \ingroup PDS_SECURITY_POLICY
/// @{

security_policy_impl *
security_policy_impl::factory(pds_policy_spec_t *spec) {
    security_policy_impl    *impl;

    // TODO: move to slab later
    impl = (security_policy_impl *)
               SDK_CALLOC(SDK_MEM_ALLOC_PDS_SECURITY_POLICY_IMPL,
                          sizeof(security_policy_impl));
    new (impl) security_policy_impl();
    return impl;
}

void
security_policy_impl::destroy(security_policy_impl *impl) {
    impl->~security_policy_impl();
    SDK_FREE(SDK_MEM_ALLOC_PDS_SECURITY_POLICY_IMPL, impl);
}

sdk_ret_t
security_policy_impl::reserve_resources(api_base *api_obj, api_base *orig_obj,
                                        api_obj_ctxt_t *obj_ctxt) {
    uint32_t             policy_block_id;
    pds_policy_spec_t    *spec;

    spec = &obj_ctxt->api_params->policy_spec;
    // allocate available block for this security policy
    if (security_policy_impl_db()->security_policy_idxr(spec->rule_info->af)->alloc(
            &policy_block_id) != SDK_RET_OK) {
        return sdk::SDK_RET_NO_RESOURCE;
    }
    security_policy_root_addr_ =
        security_policy_impl_db()->security_policy_region_addr(spec->rule_info->af) +
            (security_policy_impl_db()->security_policy_table_size(spec->rule_info->af) *
                 policy_block_id);
    return SDK_RET_OK;
}

sdk_ret_t
security_policy_impl::release_resources(api_base *api_obj) {
    uint32_t    policy_block_id;
    policy      *security_policy;

    security_policy = (policy *)api_obj;
    if (security_policy_root_addr_ != 0xFFFFFFFFFFFFFFFFUL) {
        policy_block_id =
            (security_policy_root_addr_ -
                 security_policy_impl_db()->security_policy_region_addr(security_policy->af()))/
                security_policy_impl_db()->security_policy_table_size(security_policy->af());
        security_policy_impl_db()->security_policy_idxr(security_policy->af())->free(policy_block_id);
    }
    return SDK_RET_OK;
}

sdk_ret_t
security_policy_impl::nuke_resources(api_base *api_obj) {
    return this->release_resources(api_obj);
}

sdk_ret_t
security_policy_impl::program_hw(api_base *api_obj, api_obj_ctxt_t *obj_ctxt) {
    sdk_ret_t            ret;
    pds_policy_spec_t    *spec;
    rfc::policy_params_t policy_params;

    PDS_TRACE_DEBUG("Processing security policy %s", spec->key.str());
    spec = &obj_ctxt->api_params->policy_spec;
    memset(&policy_params, 0, sizeof(policy_params));
    policy_params.policy.af = spec->rule_info->af;
    policy_params.policy.max_rules =
        (policy_params.policy.af ==IP_AF_IPV4) ?
            PDS_MAX_RULES_PER_IPV4_SECURITY_POLICY:
            PDS_MAX_RULES_PER_IPV6_SECURITY_POLICY;
    policy_params.policy.num_rules = spec->rule_info->num_rules;
    policy_params.policy.rules = &spec->rule_info->rules[0];
    policy_params.rfc_tree_root_addr = security_policy_root_addr_;
    policy_params.rfc_mem_size =
        security_policy_impl_db()->security_policy_table_size(spec->rule_info->af);
    ret = rfc_policy_create(&policy_params);
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to build RFC policy %s table, err %u",
                      spec->key.str(), ret);
    }
    return ret;
}

sdk_ret_t
security_policy_impl::update_hw(api_base *orig_obj, api_base *curr_obj,
                                api_obj_ctxt_t *obj_ctxt) {
    return this->program_hw(curr_obj, obj_ctxt);
}

sdk_ret_t
security_policy_impl::activate_hw(api_base *api_obj, api_base *orig_obj,
                                  pds_epoch_t epoch, api_op_t api_op,
                                  api_obj_ctxt_t *obj_ctxt)
{
    switch (api_op) {
    case API_OP_CREATE:
        // for security policy create, there is no stage 0 programming
        break;

    case API_OP_UPDATE:
        // need to walk all vnics AND subnets to see which of them are using
        // this policy table and then walk all the vnics that are part of the
        // vpcs and subnets and write new epoch data
        return SDK_RET_ERR;
        break;

    case API_OP_DELETE:
        // same as update but every entry written will have invalid bit set
        //return SDK_RET_ERR;
        break;

    default:
        break;
    }
    return SDK_RET_OK;
}

void
security_policy_impl::fill_status_(pds_policy_status_t *status) {
    status->policy_base_addr = security_policy_root_addr_;
}

sdk_ret_t
security_policy_impl::read_hw(api_base *api_obj, obj_key_t *key,
                              obj_info_t *info) {
    pds_policy_info_t *policy_info = (pds_policy_info_t *)info;

    fill_status_(&policy_info->status);

    return SDK_RET_OK;
}
/// \@}

}    // namespace impl
}    // namespace api
