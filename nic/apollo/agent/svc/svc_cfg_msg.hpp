//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// svc config message information
///
//----------------------------------------------------------------------------

#ifndef __SVC_CFG_MSG_HPP__
#define __SVC_CFG_MSG_HPP__

#include "nic/apollo/api/include/pds_debug.hpp"

typedef enum cfg_msg_e {
    CFG_MSG_NONE = 0,
    CFG_MSG_VPC_CREATE,
    CFG_MSG_VPC_UPDATE,
    CFG_MSG_VPC_DELETE,
    CFG_MSG_VPC_GET,
    CFG_MSG_VPC_PEER_CREATE,
    CFG_MSG_VPC_PEER_DELETE,
    CFG_MSG_VPC_PEER_GET,
    CFG_MSG_VNIC_CREATE,
    CFG_MSG_VNIC_UPDATE,
    CFG_MSG_VNIC_DELETE,
    CFG_MSG_VNIC_GET,
    CFG_MSG_SUBNET_CREATE,
    CFG_MSG_SUBNET_UPDATE,
    CFG_MSG_SUBNET_DELETE,
    CFG_MSG_SUBNET_GET,
    CFG_MSG_SECURITY_POLICY_CREATE,
    CFG_MSG_SECURITY_POLICY_UPDATE,
    CFG_MSG_SECURITY_POLICY_DELETE,
    CFG_MSG_SECURITY_POLICY_GET,
    CFG_MSG_SECURITY_PROFILE_CREATE,
    CFG_MSG_SECURITY_PROFILE_UPDATE,
    CFG_MSG_SECURITY_PROFILE_DELETE,
    CFG_MSG_SECURITY_PROFILE_GET,
    CFG_MSG_SECURITY_RULE_CREATE,
    CFG_MSG_SECURITY_RULE_UPDATE,
    CFG_MSG_SECURITY_RULE_DELETE,
    CFG_MSG_SECURITY_RULE_GET,
    CFG_MSG_POLICER_CREATE,
    CFG_MSG_POLICER_UPDATE,
    CFG_MSG_POLICER_DELETE,
    CFG_MSG_POLICER_GET,
    CFG_MSG_MAPPING_CREATE,
    CFG_MSG_MAPPING_UPDATE,
    CFG_MSG_MAPPING_DELETE,
    CFG_MSG_MAPPING_GET,
    CFG_MSG_INTERFACE_CREATE,
    CFG_MSG_INTERFACE_UPDATE,
    CFG_MSG_INTERFACE_DELETE,
    CFG_MSG_INTERFACE_GET,
    CFG_MSG_LIF_GET,
    CFG_MSG_DHCP_POLICY_CREATE,
    CFG_MSG_DHCP_POLICY_UPDATE,
    CFG_MSG_DHCP_POLICY_DELETE,
    CFG_MSG_DHCP_POLICY_GET,
    CFG_MSG_NAT_PORT_BLOCK_CREATE,
    CFG_MSG_NAT_PORT_BLOCK_DELETE,
    CFG_MSG_NAT_PORT_BLOCK_GET,
    CFG_MSG_NEXTHOP_CREATE,
    CFG_MSG_NEXTHOP_UPDATE,
    CFG_MSG_NEXTHOP_DELETE,
    CFG_MSG_NEXTHOP_GET,
    CFG_MSG_NHGROUP_CREATE,
    CFG_MSG_NHGROUP_UPDATE,
    CFG_MSG_NHGROUP_DELETE,
    CFG_MSG_NHGROUP_GET,
    CFG_MSG_METER_CREATE,
    CFG_MSG_METER_UPDATE,
    CFG_MSG_METER_DELETE,
    CFG_MSG_METER_GET,
    CFG_MSG_ROUTE_TABLE_CREATE,
    CFG_MSG_ROUTE_TABLE_UPDATE,
    CFG_MSG_ROUTE_TABLE_DELETE,
    CFG_MSG_ROUTE_TABLE_GET,
    CFG_MSG_ROUTE_CREATE,
    CFG_MSG_ROUTE_UPDATE,
    CFG_MSG_ROUTE_DELETE,
    CFG_MSG_ROUTE_GET,
    CFG_MSG_SVC_MAPPING_CREATE,
    CFG_MSG_SVC_MAPPING_UPDATE,
    CFG_MSG_SVC_MAPPING_DELETE,
    CFG_MSG_SVC_MAPPING_GET,
    CFG_MSG_TUNNEL_CREATE,
    CFG_MSG_TUNNEL_UPDATE,
    CFG_MSG_TUNNEL_DELETE,
    CFG_MSG_TUNNEL_GET,
    CFG_MSG_DEVICE_CREATE,
    CFG_MSG_DEVICE_UPDATE,
    CFG_MSG_DEVICE_DELETE,
    CFG_MSG_DEVICE_GET,
    CFG_MSG_COMMAND,
} cfg_msg_t;

typedef struct cfg_ctxt_s {
    cfg_msg_t cfg;    // CLI config command
    void *req;        // config request
} cfg_ctxt_t;

typedef enum svc_req_type_e {
    SVC_REQ_TYPE_NONE = 0,
    SVC_REQ_TYPE_CFG,
    SVC_REQ_TYPE_CMD,
} svc_req_type_t;

typedef struct svc_req_ctxt_s {
    svc_req_type_t type;
    union {
        cfg_ctxt_t cfg_ctxt;
        cmd_ctxt_t cmd_ctxt;
    };
} svc_req_ctxt_t;

#endif    // __SVC_CFG_MSG_HPP__
