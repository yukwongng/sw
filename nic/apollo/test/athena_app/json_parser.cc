//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//

#include <iostream>
#include <stdio.h>
#include "nic/sdk/include/sdk/base.hpp"
#include "nic/apollo/core/trace.hpp"
#include "fte_athena.hpp"
#include "json_parser.hpp"

using std::string;
namespace pt = boost::property_tree;

namespace fte_ath {

uint16_t g_vlan_to_vnic[MAX_VLAN_ID];
uint16_t g_mpls_label_to_vnic[MAX_MPLS_LABEL];
flow_cache_policy_info_t g_flow_cache_policy[MAX_VNIC_ID];
uint16_t g_vnic_id_list[MAX_VNIC_ID];
uint32_t g_num_policies;

static int
str2mac(const char* mac, uint8_t *out)
{
    if(ETH_ADDR_LEN == sscanf(mac, "%hhx:%hhx:%hhx:%hhx:%hhx:%hhx",
                              out, out + 1, out + 2,
                              out + 3, out + 4, out + 5)) {
        return 0;
    } else {
        return 1;
    }
}

int
parse_flow_cache_policy_cfg (const char *cfgfile)
{
    pt::ptree json_pt;
    std::string cfg_file;
    flow_cache_policy_info_t *policy;
    uint16_t vnic_id;

    cfg_file = std::string(cfgfile);
    // make sure cfg file exists
    if (access(cfg_file.c_str(), R_OK) < 0) {
        fprintf(stderr, "Policy json file %s doesn't exist or not accessible\n",
                cfg_file.c_str());
        return -1;
    }

    std::ifstream json_cfg(cfg_file.c_str());
    read_json(json_cfg, json_pt);
    printf("Parsing file %s for policies...\n", cfg_file.c_str());
    try {
        std::string mode = json_pt.get<std::string>("app_mode");
        if (mode == "l2_fwd") {
            fte_ath::g_athena_app_mode = ATHENA_APP_MODE_L2_FWD;
            return 0;
        } else if (mode == "no_dpdk") {
            fte_ath::g_athena_app_mode = ATHENA_APP_MODE_NO_DPDK;
            return 0;
        } else if (mode == "cpp") {
            fte_ath::g_athena_app_mode = ATHENA_APP_MODE_CPP;
        }

        BOOST_FOREACH (pt::ptree::value_type &vnic,
                       json_pt.get_child("vnic")) {
            vnic_id = vnic.second.get<uint32_t>("vnic_id");
            policy = &g_flow_cache_policy[vnic_id];

            policy->vnic_id = vnic_id;
            policy->vlan_id = vnic.second.get<uint16_t>("vlan_id");
            g_vlan_to_vnic[policy->vlan_id] = policy->vnic_id;
            policy->src_slot_id = vnic.second.get<uint32_t>("slot_id");
            g_mpls_label_to_vnic[policy->src_slot_id] = policy->vnic_id;

            pt::ptree& session = vnic.second.get_child("session");
            policy->epoch = session.get<uint16_t>("epoch");
            policy->skip_flow_log = session.get<bool>("skip_flow_log");
            pt::ptree& session_info_s2h = session.get_child("to_host");
            policy->to_host.tcp_flags =
                session_info_s2h.get<uint8_t>("tcp_flags");
            policy->to_host.policer_bw1 =
                std::strtoull(session_info_s2h.get<std::string>("policer_bw1").c_str(), NULL, 10);
            pt::ptree& session_info_h2s = session.get_child("to_switch");
            policy->to_switch.tcp_flags =
                session_info_h2s.get<uint8_t>("tcp_flags");
            policy->to_switch.policer_bw1 =
                std::strtoull(session_info_h2s.get<std::string>("policer_bw1").c_str(), NULL, 10);
            str2mac(session_info_h2s.get<std::string>("host_mac").c_str(),
                    policy->to_switch.host_mac);

            pt::ptree& rewrite_underlay = vnic.second.get_child("rewrite_underlay");
            std::string encap_type = rewrite_underlay.get<std::string>("type");
            if (encap_type == "mplsoudp")
                policy->rewrite_underlay.encap_type =  ENCAP_MPLSOUDP;
            else if (encap_type == "geneve")
                policy->rewrite_underlay.encap_type =  ENCAP_GENEVE;
            str2mac(rewrite_underlay.get<std::string>("smac").c_str(),
                    policy->rewrite_underlay.substrate_smac);
            str2mac(rewrite_underlay.get<std::string>("dmac").c_str(),
                    policy->rewrite_underlay.substrate_dmac);
            policy->rewrite_underlay.substrate_vlan =
                rewrite_underlay.get<uint16_t>("vlan_id");
            str2ipv4addr(rewrite_underlay.get<std::string>("ipv4_sip").c_str(),
                         &policy->rewrite_underlay.substrate_sip);
            str2ipv4addr(rewrite_underlay.get<std::string>("ipv4_dip").c_str(),
                         &policy->rewrite_underlay.substrate_dip);
            policy->rewrite_underlay.mpls_label1 =
                rewrite_underlay.get<uint32_t>("mpls_label1");
            policy->rewrite_underlay.mpls_label2 =
                rewrite_underlay.get<uint32_t>("mpls_label2");

            pt::ptree& rewrite_host = vnic.second.get_child("rewrite_host");
            str2mac(rewrite_host.get<std::string>("smac").c_str(),
                    policy->rewrite_host.ep_smac);
            str2mac(rewrite_host.get<std::string>("dmac").c_str(),
                    policy->rewrite_host.ep_dmac);

            pt::ptree& v4_flows = vnic.second.get_child("v4_flows");
            str2ipv4addr(v4_flows.get<std::string>("sip").c_str(),
                         &policy->v4_flows.sip);
            str2ipv4addr(v4_flows.get<std::string>("dip").c_str(),
                         &policy->v4_flows.dip);
            policy->v4_flows.proto =
                v4_flows.get<uint8_t>("proto");
            policy->v4_flows.sport =
                v4_flows.get<uint16_t>("sport");
            policy->v4_flows.dport =
                v4_flows.get<uint16_t>("dport");
            policy->v4_flows.num_flows =
                v4_flows.get<uint32_t>("num_flows");
            policy->v4_flows.inc_type =
                v4_flows.get<uint8_t>("inc_type");

            g_vnic_id_list[g_num_policies++] = vnic_id;
        }
    } catch (std::exception const& e) {
        std::cerr << e.what() << std::endl;
        return -1;
    }
    printf("POLICIES PARSED %u\n", g_num_policies);
    return 0;
}

}
