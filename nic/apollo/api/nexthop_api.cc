//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file implements nexthop CRUD APIs
///
//----------------------------------------------------------------------------

#include "nic/apollo/framework/api_base.hpp"
#include "nic/apollo/framework/api_msg.hpp"
#include "nic/apollo/framework/api_engine.hpp"
#include "nic/apollo/framework/api_params.hpp"
#include "nic/apollo/api/obj_api.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/apollo/api/nexthop.hpp"
#include "nic/apollo/api/nexthop_state.hpp"

static sdk_ret_t
pds_nexthop_api_handle (pds_batch_ctxt_t bctxt, api_op_t op,
                        pds_obj_key_t *key, pds_nexthop_spec_t *spec)
{
    sdk_ret_t rv;
    api_ctxt_t *api_ctxt;

    if ((rv = pds_obj_api_validate(op, key, spec)) != SDK_RET_OK) {
        return rv;
    }

    api_ctxt = api::api_ctxt_alloc(OBJ_ID_NEXTHOP, op);
    if (likely(api_ctxt != NULL)) {
        if (op == API_OP_DELETE) {
            api_ctxt->api_params->key = *key;
        } else {
            api_ctxt->api_params->nexthop_spec = *spec;
        }
        return process_api(bctxt, api_ctxt);
    }
    return SDK_RET_OOM;
}

static inline nexthop *
pds_nexthop_find (pds_obj_key_t *key)
{
    return (nexthop_db()->find(key));
}

//----------------------------------------------------------------------------
// nexthop API entry point implementation
//----------------------------------------------------------------------------

sdk_ret_t
pds_nexthop_create (_In_ pds_nexthop_spec_t *spec, _In_ pds_batch_ctxt_t bctxt)
{
    return pds_nexthop_api_handle(bctxt, API_OP_CREATE, NULL, spec);
}

sdk_ret_t
pds_nexthop_read (_In_ pds_obj_key_t *key, _Out_ pds_nexthop_info_t *info)
{
    nexthop *entry;

    if (key == NULL || info == NULL) {
        return SDK_RET_INVALID_ARG;
    }

    if ((entry = pds_nexthop_find(key)) == NULL) {
        return SDK_RET_ENTRY_NOT_FOUND;
    }

    return entry->read(info);
}

typedef struct pds_nexthop_read_args_s {
    nexthop_read_cb_t cb;
    void *ctxt;
} pds_nexthop_read_args_t;

bool
pds_nexthop_info_from_entry (void *entry, void *ctxt)
{
    nexthop *nh = (nexthop *)entry;
    pds_nexthop_read_args_t *args = (pds_nexthop_read_args_t *)ctxt;
    pds_nexthop_info_t info;

    memset(&info, 0, sizeof(pds_nexthop_info_t));

    // call entry read
    nh->read(&info);

    // call cb on info
    args->cb(&info, args->ctxt);

    return false;
}

sdk_ret_t
pds_nexthop_read_all (nexthop_read_cb_t nexthop_read_cb, void *ctxt)
{
    pds_nexthop_read_args_t args = {0};

    args.ctxt = ctxt;
    args.cb = nexthop_read_cb;

    return nexthop_db()->walk(pds_nexthop_info_from_entry, &args);
}

sdk_ret_t
pds_nexthop_update (_In_ pds_nexthop_spec_t *spec, _In_ pds_batch_ctxt_t bctxt)
{
    return pds_nexthop_api_handle(bctxt, API_OP_UPDATE, NULL, spec);
}

sdk_ret_t
pds_nexthop_delete (_In_ pds_obj_key_t *key, _In_ pds_batch_ctxt_t bctxt)
{
    return pds_nexthop_api_handle(bctxt, API_OP_DELETE, key, NULL);
}
