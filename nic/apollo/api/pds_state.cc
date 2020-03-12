/**
 * Copyright (c) 2018 Pensando Systems, Inc.
 *
 * @file    pds_state.cc
 *
 * @brief   This file contains implementation of pds state class
 */

#include "nic/apollo/api/pds_state.hpp"

namespace api {

/**< (singleton) instance of all PDS state in one place */
pds_state g_pds_state;

/**
 * @defgroup PDS_STATE - Internal state
 * @{
 */

/**< @brief    constructor */
pds_state::pds_state() {
    catalog_ = NULL;
    mpartition_ = NULL;
    memset(state_, 0, sizeof(state_));
    event_cb_ = nullptr;
    memset(&system_mac_, 0, sizeof(system_mac_));
    // set vpp mock mode as needed
    if (getenv("VPP_IPC_MOCK_MODE")) {
        vpp_ipc_mock_ = true;
    }
    // set the config path
    SDK_ASSERT(std::getenv("CONFIG_PATH"));
    cfg_path_ = std::string(std::getenv("CONFIG_PATH"));
    if (cfg_path_.empty()) {
        cfg_path_ = std::string("./");
    } else {
        cfg_path_ += std::string("/");
    }
}

/**< @brief    destructor */
pds_state::~pds_state() {
}

sdk_ret_t
pds_state::parse_global_config_(string pipeline, string cfg_file)
{
    ptree     pt;

    cfg_file = cfg_path_ + pipeline + "/" + cfg_file;
    // make sure global config file exists
    if (access(cfg_file.c_str(), R_OK) < 0) {
        fprintf(stderr, "Config file %s doesn't exist or not accessible\n",
                cfg_file.c_str());
        return SDK_RET_ERR;
    }
    // parse the config now
    std::ifstream json_cfg(cfg_file.c_str());
    read_json(json_cfg, pt);
    try {
        std::string mode = pt.get<std::string>("mode");
        if (mode == "sim") {
            platform_type_ = platform_type_t::PLATFORM_TYPE_SIM;
        } else if (mode == "hw") {
            platform_type_ = platform_type_t::PLATFORM_TYPE_HW;
        } else if (mode == "rtl") {
            platform_type_ = platform_type_t::PLATFORM_TYPE_RTL;
        } else if (mode == "haps") {
            platform_type_ = platform_type_t::PLATFORM_TYPE_HAPS;
        } else if (mode == "mock") {
            platform_type_ = platform_type_t::PLATFORM_TYPE_MOCK;
        }
    } catch (std::exception const& e) {
        std::cerr << e.what() << std::endl;
        return sdk::SDK_RET_INVALID_ARG;
    }
    return SDK_RET_OK;
}

sdk_ret_t
pds_state::init(string pipeline, string cfg_file) {
    string path("");

    // parse and store global configuration
    pipeline_ = pipeline;
    SDK_ASSERT(parse_global_config_(pipeline, cfg_file) == SDK_RET_OK);

    // create persistent store
    if (platform_type_ == platform_type_t::PLATFORM_TYPE_HW) {
        kvstore_ = sdk::lib::kvstore::factory("/data/pdsagent.db", (1UL << 31));
    } else {
        path = std::string(getenv("PDSPKG_TOPDIR"));
        if (path.empty()) {
            path = "./";
        } else {
            path += "/";
        }
        kvstore_ = sdk::lib::kvstore::factory((path + "pdsagent.db").c_str(),
                                              (1UL << 31));
    }
    if (kvstore_ == NULL) {
        return SDK_RET_ERR;
    }
    state_[PDS_STATE_DEVICE] = new device_state();
    state_[PDS_STATE_LIF] = new lif_state();
    state_[PDS_STATE_IF] = new if_state();
    state_[PDS_STATE_TEP] = new tep_state();
    state_[PDS_STATE_VPC] = new vpc_state();
    state_[PDS_STATE_SUBNET] = new subnet_state();
    state_[PDS_STATE_VNIC] = new vnic_state();
    state_[PDS_STATE_MAPPING] = new mapping_state(kvstore_);
    state_[PDS_STATE_ROUTE_TABLE] = new route_table_state();
    state_[PDS_STATE_POLICY] = new policy_state();
    state_[PDS_STATE_MIRROR] = new mirror_session_state();
    state_[PDS_STATE_METER] = new meter_state();
    state_[PDS_STATE_TAG] = new tag_state();
    state_[PDS_STATE_SVC_MAPPING] = new svc_mapping_state();
    state_[PDS_STATE_VPC_PEER] = new vpc_peer_state();
    state_[PDS_STATE_NEXTHOP] = new nexthop_state();
    state_[PDS_STATE_NEXTHOP_GROUP] = new nexthop_group_state();
    state_[PDS_STATE_POLICER] = new policer_state();
    state_[PDS_STATE_NAT] = new nat_state();
    state_[PDS_STATE_DHCP] = new dhcp_state();
    state_[PDS_STATE_LEARN] = new learn_state();
    return SDK_RET_OK;
}

void
pds_state::destroy(pds_state *ps) {
    if (ps->catalog_) {
        catalog::destroy(ps->catalog_);
    }
    if (ps->mpartition_) {
        sdk::platform::utils::mpartition::destroy(ps->mpartition_);
    }
    if (ps->pginfo_) {
        sdk::platform::utils::program_info::destroy(ps->pginfo_);
    }
    for (uint32_t i = PDS_STATE_MIN + 1; i < PDS_STATE_MAX; i++) {
        if (ps->state_[i]) {
            delete ps->state_[i];
        }
    }
}

sdk_ret_t
pds_state::slab_walk(state_walk_cb_t walk_cb, void *ctxt) {
    for (uint32_t i = PDS_STATE_MIN + 1; i < PDS_STATE_MAX; i ++) {
        state_[i]->slab_walk(walk_cb, ctxt);
    }
    return SDK_RET_OK;
}

sdk_ret_t
pds_state::walk(state_walk_cb_t walk_cb, void *ctxt) {
    state_walk_ctxt_t walk_ctxt;

    for (uint32_t i = PDS_STATE_MIN + 1; i < PDS_STATE_MAX; i ++) {
        if (state_[i]) {
            walk_ctxt.obj_state.assign(PDS_STATE_str(pds_state_t(i)));
            walk_ctxt.state = state_[i];
            walk_cb(&walk_ctxt, ctxt);
        }
    }
    return SDK_RET_OK;
}

sdk_ret_t
pds_state::transaction_begin(void) {
    //return mapping_db()->transaction_begin();
    return SDK_RET_OK;
}

sdk_ret_t
pds_state::transaction_end(bool abort) {
    //return mapping_db()->transaction_end(abort);
    return SDK_RET_OK;
}

/** * @} */    // end of PDS_STATE

}    // namespace api
