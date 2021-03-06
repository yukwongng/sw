//
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// interface entry handling
///
//----------------------------------------------------------------------------

#ifndef __API_IF_HPP__
#define __API_IF_HPP__

#include <string.h>
#include "nic/sdk/include/sdk/if.hpp"
#include "nic/sdk/lib/ht/ht.hpp"
#include "nic/apollo/api/include/pds.hpp"
#include "nic/apollo/framework/api_base.hpp"
#include "nic/apollo/framework/impl_base.hpp"
#include "nic/apollo/api/include/pds_if.hpp"

#define PDS_MAX_IF    256             ///< Max interfaces

// attribute update bits for interface object
#define PDS_IF_UPD_ADMIN_STATE         0x1

namespace api {

// forward declaration
class if_state;

/// \defgroup PDS_IF_ENTRY - interface entry functionality
/// \ingroup PDS_IF
/// @{

/// \brief    interface entry
class if_entry : public api_base {
public:
    /// \brief    factory method to allocate & initialize a interface entry
    /// \param[in] key        key of the interface
    /// \param[in] ifindex    encoded interface index
    /// \return    new instance of interface or NULL, in case of error
    static if_entry *factory(pds_obj_key_t& key, if_index_t ifindex);

    /// \brief    factory method to allocate & initialize a interface entry
    /// \param[in] spec    interface specification
    /// \return    new instance of interface or NULL, in case of error
    static if_entry *factory(pds_if_spec_t *spec);

    /// \brief          release all the s/w state associate with the given
    ///                 interface, if any, and free the memory
    /// \param[in]      intf    interface to be freed
    /// \NOTE: h/w entries should have been cleaned up (by calling
    ///        impl->cleanup_hw() before calling this
    static void destroy(if_entry *intf);

    /// \brief    clone this object and return cloned object
    /// \param[in]    api_ctxt API context carrying object related configuration
    /// \return       new object instance of current object
    virtual api_base *clone(api_ctxt_t *api_ctxt) override;

    /// \brief    free all the memory associated with this object without
    ///           touching any of the databases or h/w etc.
    /// \param[in] if    interface to be freed
    /// \return   sdk_ret_ok or error code
    static sdk_ret_t free(if_entry *intf);

    /// \brief     allocate h/w resources for this object
    /// \param[in] orig_obj old version of the unmodified object
    /// \param[in] obj_ctxt transient state associated with this API
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t reserve_resources(api_base *orig_obj,
                                        api_obj_ctxt_t *obj_ctxt) override;

    /// \brief  free h/w resources used by this object, if any
    /// \return SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t release_resources(void) override;

    /// \brief     initialize an interface entry with the given config
    /// \param[in] api_ctxt API context carrying the configuration
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t init_config(api_ctxt_t *api_ctxt) override;

    /// \brief     program all h/w tables relevant to this object except
    ///            stage 0 table(s), if any
    /// \param[in] obj_ctxt transient state associated with this API
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t program_create(api_obj_ctxt_t *obj_ctxt) override;

    /// \brief     cleanup all h/w tables relevant to this object except
    ///            stage 0 table(s), if any, by updating packed entries
    ///            with latest epoch#
    /// \param[in] obj_ctxt transient state associated with this API
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t cleanup_config(api_obj_ctxt_t *obj_ctxt) override;

    /// \brief    compute the object diff during update operation compare the
    ///           attributes of the object on which this API is invoked and the
    ///           attrs provided in the update API call passed in the object
    ///           context (as cloned object + api_params) and compute the upd
    ///           bitmap (and stash in the object context for later use)
    /// \param[in] obj_ctxt    transient state associated with this API
    /// \return #SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t compute_update(api_obj_ctxt_t *obj_ctxt) override;

    /// \brief          add all objects that may be affected if this object is
    ///                 updated to framework's object dependency list
    /// \param[in]      obj_ctxt    transient state associated with this API
    ///                             processing
    /// \return         SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t add_deps(api_obj_ctxt_t *obj_ctxt) override {
        return SDK_RET_OK;
    }

    /// \brief          reprogram all h/w tables relevant to this object and
    ///                 dependent on other objects except stage 0 table(s),
    ///                 if any
    /// \param[in]      obj_ctxt    transient state associated with this API
    ///                             processing
    /// \return         SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t reprogram_config(api_obj_ctxt_t *obj_ctxt) override {
        return SDK_RET_INVALID_OP;
    }

    /// \brief     update all h/w tables relevant to this object except
    ///            stage 0 table(s), if any, by updating packed entries
    ///            with latest epoch#
    /// \param[in] orig_obj old version of the unmodified object
    /// \param[in] obj_ctxt transient state associated with this API
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t program_update(api_base *orig_obj,
                                    api_obj_ctxt_t *obj_ctxt) override;

    /// \brief     activate the epoch in the dataplane by programming
    ///            stage 0 tables, if any
    /// \param[in] epoch    epoch being activated
    /// \param[in] api_op   api operation
    /// \param[in] orig_obj old/original version of the unmodified object
    /// \param[in] obj_ctxt transient state associated with this API
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t activate_config(pds_epoch_t epoch, api_op_t api_op,
                                      api_base *orig_obj,
                                      api_obj_ctxt_t *obj_ctxt) override;

    /// \brief  add given interface to the database
    /// \return SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t add_to_db(void) override;

    /// \brief  delete given interface from the database
    /// \return SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t del_from_db(void) override;

    /// \brief     this method is called on new object that needs to
    ///            replace the old version of the object in the DBs
    /// \param[in] orig_obj old version of the object being swapped out
    /// \param[in] obj_ctxt transient state associated with this API
    /// \return    SDK_RET_OK on success, failure status code on error
    virtual sdk_ret_t update_db(api_base *orig_obj,
                                api_obj_ctxt_t *obj_ctxt) override;

    /// \brief          initiate delay deletion of this object
    virtual sdk_ret_t delay_delete(void) override;

    /// \brief          read config
    /// \param[out]     info pointer to the info object
    /// \return         SDK_RET_OK on success, failure status code on error
    sdk_ret_t read(pds_if_info_t *info);

    /// \brief          return stringified key of the object (for debugging)
    virtual string key2str(void) const override {
        return "if-" + std::string(key_.str());
    }

    /// \brief          helper function to get key given interface entry
    /// \param          entry    pointer to interface instance
    /// \return         pointer to the interface instance's key
    static void *if_key_func_get(void *entry) {
        if_entry *intf = (if_entry *)entry;
        return (void *)&(intf->key_);
    }

    /// \brief          helper function to get ifindex given interface entry
    /// \param          entry    pointer to interface instance
    /// \return         pointer to the interface instance's key
    static void *ifindex_func_get(void *entry) {
        if_entry *intf = (if_entry *)entry;
        return (void *)&(intf->ifindex_);
    }

    /// \brief     return impl instance of this interface object
    /// \return    impl instance of the interface object
    impl_base *impl(void) { return impl_; }

    /// \brief    return the interface index
    /// \return interface index
    if_index_t ifindex(void) const { return ifindex_; }

    /// \brief    return the interface key
    /// \return interface index
    pds_obj_key_t key(void) const { return key_; }

    /// \brief    return admin state
    /// \return interface admin state
    pds_if_state_t admin_state(void) const { return admin_state_; }

    /// \brief    return the interface type
    /// \return interface type
    if_type_t type(void) const { return type_; }

    /// \brief    return the wire encap of this (L3) interface
    /// \return   wire encap of this L3 interface
    pds_encap_t l3_encap(void) { return if_info_.l3_.encap_; }

    /// \brief    return the vpc of this (L3) interface
    /// \return   vpc of this L3 interface
    pds_obj_key_t l3_vpc(void) { return if_info_.l3_.vpc_; }

    /// \brief    return the MAC address of this interface
    /// \return   MAC address of this interface
    mac_addr_t& l3_mac(void) { return if_info_.l3_.mac_; }

    /// \brief    return IP subnet of this interface
    /// \return   IP prefix configured on this interface
    ip_prefix_t& l3_ip_prefix(void) { return if_info_.l3_.ip_pfx_; }

    /// \brief    return gateway for control interface
    /// \return   gateway ip for control interface
    ip_addr_t& control_gateway(void) { return if_info_.control_.gateway_; }

    /// \brief    set the opaque port info for this port
    /// \param[in] port_info    pointer to the port specific information
    void set_port_info(void *port_info) {
        if_info_.port_.port_info_ = port_info;
    }

    /// \brief    set the interface type for this interface
    /// \param[in] if_type    interface type
    void set_type(if_type_t if_type) {
        type_ = if_type;
    }

    /// \brief    return port specific information
    /// \return return pointer to the port specific information
    void *port_info(void) { return if_info_.port_.port_info_; }

    /// \brief    return oper status
    /// \return interface operational state
    pds_if_state_t state(void);

    /// \brief    return port type
    /// \return interface port type
    port_type_t port_type(void);

    /// \brief    return interface name
    /// \return user consumable interface name with system mac
    std::string name(void);

    /// \brief    return eth interface entry corresponding to the given
    ///           interface
    /// \param[in] intf    interface instance
    /// \return eth interface corresponding to the given interface or NULL
    static if_entry *eth_if(if_entry *intf);

private:
    /// \brief constructor
    if_entry();

    /// \brief destructor
    ~if_entry();

    /// \brief      fill the interface sw spec
    /// \param[out] spec specification
    /// \return     SDK_RET_OK on success, failure status code on error
    sdk_ret_t fill_spec_(pds_if_spec_t *spec);

    /// \brief    free h/w resources used by this object, if any
    ///           (this API is invoked during object deletes)
    /// \return    SDK_RET_OK on success, failure status code on error
    sdk_ret_t nuke_resources_(void);

private:
    pds_obj_key_t  key_;           ///< interface key
    if_index_t     ifindex_;       ///< interface index
    if_type_t      type_;          ///< interface type
    pds_if_state_t admin_state_;   ///< admin state
    union {
        ///< physical port specific information
        struct {
            void *port_info_;      ///< interface specific information
        } port_;
        ///< uplink specific information
        struct {
            pds_obj_key_t port_;
        } uplink_;
        ///< L3 interface specific information
        union {
            struct {
                pds_obj_key_t vpc_;    ///< vpc of this L3 interface
                ip_prefix_t ip_pfx_;   ///< IP subnet of this L3 interface
                pds_obj_key_t port_;   ///< physical port of this L3 interface
                pds_encap_t encap_;    ///< wire encap, if any
                mac_addr_t mac_;       ///< MAC address of this L3 interface
            } l3_;
            struct {
                ip_prefix_t ip_pfx_;   ///< IP subnet of this interface
                mac_addr_t mac_;       ///< MAC address of this interface
                ip_addr_t gateway_;
            } control_;
        };
    } if_info_;

    ht_ctxt_t ht_ctxt_;            ///< hash table context
    ht_ctxt_t ifindex_ht_ctxt_;    ///< hash table context
    impl_base *impl_;              ///< impl object instance
    friend class if_state;         ///< if_state is friend of if_entry
};

/// \@}    // end of PDS_IF_ENTRY

}    // namespace api

using api::if_entry;

#endif    // __API_IF_HPP__
