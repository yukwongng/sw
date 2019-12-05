//---------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
// LI VXLAN Tunnel HAL integration
//---------------------------------------------------------------

#include "nic/metaswitch/stubs/hals/pds_ms_li_intf.hpp"
#include "nic/metaswitch/stubs/common/pdsa_linux_util.hpp"
#include "nic/metaswitch/stubs/common/pdsa_state.hpp"
#include "nic/metaswitch/stubs/hals/pds_ms_ifindex.hpp"
#include "nic/sdk/lib/logger/logger.hpp"
#include <li_fte.hpp>
#include <li_lipi_slave_join.hpp>
#include <li_port.hpp>

extern NBB_ULONG li_proc_id;

namespace pds_ms {

using pdsa_stub::Error;
using pdsa_stub::get_linux_intf_params;

void li_intf_t::parse_ips_info_(ATG_LIPI_PORT_ADD_UPDATE* port_add_upd_ips) {
    ips_info_.ifindex = port_add_upd_ips->id.if_index;
    ips_info_.if_name = port_add_upd_ips->id.if_name;
    ips_info_.admin_state = 
        (port_add_upd_ips->port_settings.port_enabled == ATG_YES);
    ips_info_.admin_state_valid = 
        (port_add_upd_ips->port_settings.port_enabled_updated == ATG_YES);
}

void li_intf_t::fetch_store_info_(pdsa_stub::state_t* state) {
    store_info_.phy_port_if_obj = state->if_store().get(ips_info_.ifindex);
    if (op_delete_) {
        if (unlikely(store_info_.phy_port_if_obj == nullptr)) {
            throw Error(std::string("Port delete for unknown IfIndex ")
                        .append(std::to_string(ips_info_.ifindex)));
        }
        return;
    }
    if (store_info_.phy_port_if_obj == nullptr) { 
        op_create_ = true;
        return; 
    }
    if (!store_info_.phy_port_if_obj->phy_port_properties().hal_created) { 
        // If Obj was saved in store to cache the linux parameters 
        // but not yet created in HAL - issue HAL PDS create now
        op_create_ = true; 
    }
}

bool li_intf_t::cache_new_obj_in_cookie_(void) {
    std::unique_ptr<if_obj_t> new_if_obj; 
    if (store_info_.phy_port_if_obj == nullptr) {
        // This is the first time we are seeing this uplink interface 
        auto port_prop = if_obj_t::phy_port_properties_t {0};
        port_prop.ifindex = ips_info_.ifindex;
        new_if_obj.reset(new if_obj_t(port_prop));

        // TODO: Move Linux Intf param fetch to the Mgmt Stub
        auto& phy_port_prop = new_if_obj->phy_port_properties();
        if (!get_linux_intf_params(ips_info_.if_name, 
                                   &phy_port_prop.lnx_ifindex,
                                   phy_port_prop.mac_addr)) {
            throw Error (std::string("Could not fetch Linux params for ")
                         .append(ips_info_.if_name));
        }
    } else {
        // Create a new object in order to store the updated fields
        // but do not save it in the Global State yet
        new_if_obj.reset(new if_obj_t(*(store_info_.phy_port_if_obj)));
    }
    auto& phy_port_prop = new_if_obj->phy_port_properties();
    if (op_create_) {
        phy_port_prop.hal_created = true;
        phy_port_prop.admin_state = ips_info_.admin_state;
        SDK_TRACE_DEBUG ("MS If 0x%lx: Admin State %d", 
                         ips_info_.ifindex, ips_info_.admin_state);
    } else if (ips_info_.admin_state_valid) {
        SDK_TRACE_DEBUG ("MS If 0x%lx: Admin State change to %d", 
                         ips_info_.ifindex, ips_info_.admin_state);
        // Update the new admin state in the new If object
        phy_port_prop.admin_state = ips_info_.admin_state;
    } else {
        // Update request but no change in the fields we are 
        // interested in
        SDK_TRACE_VERBOSE ("MS If 0x%lx: No-op update", ips_info_.ifindex);
        return false;
    }
    // Update the local store info context so that the make_pds_spec 
    // refers to the latest fields
    store_info_.phy_port_if_obj = new_if_obj.get(); 
    // Cache the new object in the cookie to revisit asynchronously
    // when the PDS API response is received
    cookie_uptr_->objs.push_back(std::move(new_if_obj));
    return true;
}

pds_if_key_t li_intf_t::make_pds_if_key_(void) {
    pds_if_key_t key; 
    key.id = ms_to_pds_ifindex (ips_info_.ifindex);
    return key;
}

pds_if_spec_t li_intf_t::make_pds_if_spec_(void) {
    pds_if_spec_t spec = {0};
    spec.key = make_pds_if_key_();
    spec.type = PDS_IF_TYPE_L3;
    auto& port_prop = store_info_.phy_port_if_obj->phy_port_properties();
    spec.admin_state = 
        (port_prop.admin_state) ? PDS_IF_STATE_UP : PDS_IF_STATE_DOWN;
    // TODO: Change this to eth IfIndex when HAL supports it
    auto ifindex = ms_to_pds_eth_ifindex (ips_info_.ifindex);
    spec.l3_if_info.port_num = ETH_IFINDEX_TO_PARENT_PORT(ifindex)-1;
    memcpy (spec.l3_if_info.mac_addr, port_prop.mac_addr, ETH_ADDR_LEN);
    return spec;
}

pds_batch_ctxt_guard_t li_intf_t::make_batch_pds_spec_(void) {
    pds_batch_ctxt_guard_t bctxt_guard_;

    SDK_ASSERT(cookie_uptr_); // Cookie should not be empty
    pds_batch_params_t bp {0, true, (uint64_t) cookie_uptr_.get()};
    auto bctxt = pds_batch_start(&bp);

    if (unlikely (!bctxt)) {
        throw Error(std::string("PDS Batch Start failed for MS If ")
                    .append(std::to_string(ips_info_.ifindex)));
    }
    bctxt_guard_.set (bctxt);

    if (op_delete_) { // Delete
        auto if_key = make_pds_if_key_();
        pds_if_delete(&if_key, bctxt);

    } else { // Add or update
        auto if_spec = make_pds_if_spec_();
        sdk_ret_t ret;
        if (op_create_) {
            ret = pds_if_create(&if_spec, bctxt);
        } else {
            ret = pds_if_update(&if_spec, bctxt);
        }
        if (unlikely (ret != SDK_RET_OK)) {
            throw Error(std::string("PDS If Create or Update failed for MS If ")
                        .append(std::to_string(ips_info_.ifindex)));
        }
    }
    return bctxt_guard_;
}

void li_intf_t::handle_add_upd_ips(ATG_LIPI_PORT_ADD_UPDATE* port_add_upd_ips) {
    pds_batch_ctxt_guard_t  pds_bctxt_guard;
    port_add_upd_ips->return_code = ATG_OK;

    parse_ips_info_(port_add_upd_ips);

    if (ms_ifindex_to_pds_type (ips_info_.ifindex) != IF_TYPE_L3) {
        // Nothing to do for non-L3 interfaces
        return;
    }
    if (port_add_upd_ips->port_settings.no_switch_port == ATG_NO) {
        // Not an IP interface
        // IP to non-IP transition is not supported
        return;
    }

    // Alloc new cookie and cache IPS
    cookie_uptr_.reset (new cookie_t);

    { // Enter thread-safe context to access/modify global state
        auto state_ctxt = pdsa_stub::state_t::thread_context();
        fetch_store_info_(state_ctxt.state());

        if (op_create_) {
            SDK_TRACE_INFO ("MS If 0x%lx: Create IPS", ips_info_.ifindex);
        } else {
            SDK_TRACE_INFO ("MS If 0x%lx: Update IPS", ips_info_.ifindex);
        }
        if (unlikely(!cache_new_obj_in_cookie_())) {
            // No change
            SDK_TRACE_DEBUG ("MS If 0x%lx: No-op IPS", ips_info_.ifindex);
            return;
        } 

        pds_bctxt_guard = make_batch_pds_spec_(); 
        // If we have batched multiple IPS earlier flush it now
        // Cannot add a Intf create to an existing batch
        state_ctxt.state()->flush_outstanding_pds_batch();

    } // End of state thread_context
      // Do Not access/modify global state after this

    cookie_uptr_->send_ips_reply = 
        [port_add_upd_ips] (bool pds_status) -> void {
            // ----------------------------------------------------------------
            // This block is executed asynchronously when PDS response is rcvd
            // ----------------------------------------------------------------
            NBB_CREATE_THREAD_CONTEXT
            NBS_ENTER_SHARED_CONTEXT(li_proc_id);
            NBS_GET_SHARED_DATA();

            auto key = li::Port::get_key(*port_add_upd_ips);
            auto& port_store = li::Fte::get().get_lipi_join()->get_port_store();
            auto it = port_store.find(key);
            if (it == port_store.end()) {
                auto send_response = 
                    li::Port::set_ips_rc(&port_add_upd_ips->ips_hdr,
                                         (pds_status) ? ATG_OK : ATG_UNSUCCESSFUL);
                SDK_ASSERT(send_response);
                SDK_TRACE_DEBUG("Phy port 0x%x: Send Async IPS reply %s stateless mode",
                                key.index, (pds_status) ? "Success" : "Failure");
                li::Fte::get().get_lipi_join()->
                    send_ips_reply(&port_add_upd_ips->ips_hdr);
            } else {
                if (pds_status) {
                    SDK_TRACE_DEBUG("Phy Port 0x%x: Send Async IPS Reply success stateful mode",
                                    key.index);
                    (*it)->update_complete(ATG_OK);
                } else {
                    SDK_TRACE_DEBUG("Phy Port 0x%x: Send Async IPS Reply failure stateful mode",
                                    key.index);
                    (*it)->update_failed(ATG_UNSUCCESSFUL);
                }
            }
            NBS_RELEASE_SHARED_DATA();
            NBS_EXIT_SHARED_CONTEXT();
            NBB_DESTROY_THREAD_CONTEXT    
        };
    // All processing complete, only batch commit remains - 
    // safe to release the cookie_uptr_ unique_ptr
    auto cookie = cookie_uptr_.release();
    if (unlikely (pds_batch_commit(pds_bctxt_guard.release()) != SDK_RET_OK)) {
        delete cookie;
        throw Error(std::string("Batch commit failed for Add-Update Port MS If ")
                    .append(std::to_string(ips_info_.ifindex)));
    }
    port_add_upd_ips->return_code = ATG_ASYNC_COMPLETION;
    SDK_TRACE_DEBUG ("MS If 0x%lx: Add/Upd PDS Batch commit successful", 
                     ips_info_.ifindex);
}

void li_intf_t::handle_delete(NBB_ULONG ifindex) {
    pds_batch_ctxt_guard_t  pds_bctxt_guard;
    op_delete_ = true;

    // TODO: Current API signature does not allow Async callback to MS
    // However since we should not block the MS NBase main thread
    // the HAL proessing is always asynchrnous. 
    // Assuming that Deletes never fail the Store is also updated
    // in a synchronous fashion for deletes so that it is in sync with MS.

    ips_info_.ifindex = ifindex;
    SDK_TRACE_INFO ("MS If 0x%lx: Delete IPS", ips_info_.ifindex);

    // Empty cookie to force async PDS.
    // But it does not have any cached objects for now
    cookie_uptr_.reset (new cookie_t);
    pds_bctxt_guard = make_batch_pds_spec_ (); 

    { // Enter thread-safe context to access/modify global state
        auto state_ctxt = pdsa_stub::state_t::thread_context();
        // If we have batched multiple IPS earlier flush it now
        // Cannot add a Tunnel create to an existing batch
        state_ctxt.state()->flush_outstanding_pds_batch();
    } // End of state thread_context
      // Do Not access/modify global state after this

    cookie_uptr_->send_ips_reply = 
        [ifindex] (bool pds_status) -> void {
            // ----------------------------------------------------------------
            // This block is executed asynchronously when PDS response is rcvd
            // ----------------------------------------------------------------
            SDK_TRACE_DEBUG("Phy port 0x%x Delete: Rcvd Async PDS response %s",
                            ifindex, (pds_status) ? "Success" : "Failure");

        };
    // All processing complete, only batch commit remains - 
    // safe to release the cookie_uptr_ unique_ptr
    auto cookie = cookie_uptr_.release();
    if (pds_batch_commit(pds_bctxt_guard.release()) != SDK_RET_OK) {
        delete cookie;
        throw Error(std::string("Batch commit failed for delete MS If ")
                    .append(std::to_string(ips_info_.ifindex)));
    }
    SDK_TRACE_DEBUG ("MS If 0x%lx: Delete PDS Batch commit successful", 
                     ips_info_.ifindex);

    { // Enter thread-safe context to access/modify global state
        auto state_ctxt = pdsa_stub::state_t::thread_context();
        // Deletes are synchronous - Delete the store entry immediately 
        state_ctxt.state()->if_store().erase(ifindex);
    }
}

} // End namespace