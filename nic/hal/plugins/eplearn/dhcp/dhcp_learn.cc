/*
 * dhcp.cc
 *
 *  Created on: Aug 9, 2017
 *      Author: sudhiaithal
 */

#include <netinet/in.h>
// clang-format off
#include "dhcp_trans.hpp"
#include "dhcp_packet.hpp"
// clang-format on

// using namespace DHCP_EP_LEARN;

namespace hal {
namespace eplearn {

static hal_ret_t get_dhcp_status(vrf_id_t vrf_id,
        ip_addr_t *ip_addr, DhcpStatus *dhcp_status);

void dhcp_init()
{
    dhcp_lib_init();
    register_dhcp_ep_status_callback(get_dhcp_status);
}

static hal_ret_t dhcp_process_request_internal(struct packet *decoded_packet,
                                               fte::ctx_t &ctx,
                                               dhcp_fsm_event_t event,
                                               bool add_entry) {
    struct dhcp_packet *raw_pkt = decoded_packet->raw;
    uint32_t xid = ntohl(raw_pkt->xid);
    dhcp_trans_t *trans;
    dhcp_trans_key_t trans_key;
    dhcp_event_data event_data;

    HAL_TRACE_DEBUG("Processing DHCP event {} : ", event);
    dhcp_trans_t::init_dhcp_trans_key(raw_pkt->chaddr, xid,
            ctx.sep(), &trans_key);
    trans = dhcp_trans_t::find_dhcptrans_by_id(trans_key);
    if (trans == NULL) {
        if (add_entry) {
            HAL_TRACE_DEBUG("Creating new DHCP transaction for event {}", event);
            trans = new dhcp_trans_t(decoded_packet, ctx);
        } else {
            HAL_TRACE_ERR("No existing transaction found for event {}", event);
            return HAL_RET_ENTRY_NOT_FOUND;
        }
    } else {
        trans->log_info("Found an existing transaction");
    }

    event_data.decoded_packet = decoded_packet;
    event_data.fte_ctx = &ctx;
    event_data.in_fte_pipeline = true;
    event_data.event = event;
    dhcp_trans_t::process_transaction(trans, event,
                                         (fsm_event_data)(&event_data));
    return HAL_RET_OK;
}

bool is_dhcp_flow(const hal::flow_key_t *key) {
    return (key->sport == 68 || key->dport == 68);
}

typedef std::function<void(struct packet *decoded_packet, fte::ctx_t &)>
    process_func;

process_func proc_function_handler(dhcp_fsm_event_t event, bool add_on_entry) {
    return [=](struct packet *pkt, fte::ctx_t &ctx) {
        dhcp_process_request_internal(pkt, ctx, event, add_on_entry);
    };
}

static std::map<int, process_func> process_func_map = {
    {DHCPDISCOVER, proc_function_handler(DHCP_DISCOVER, true)},
    {DHCPOFFER, proc_function_handler(DHCP_OFFER, false)},
    {DHCPREQUEST, proc_function_handler(DHCP_REQUEST, true)},
    {DHCPINFORM, proc_function_handler(DHCP_INFORM, true)},
    {DHCPRELEASE, proc_function_handler(DHCP_RELEASE, true)},
    {DHCPACK, proc_function_handler(DHCP_ACK, false)},
    {DHCPNAK, proc_function_handler(DHCP_NACK, false)},
    {DHCPDECLINE, proc_function_handler(DHCP_DECLINE, true)},
};

hal_ret_t dhcp_process_packet(fte::ctx_t &ctx) {
    struct packet *decoded_packet = NULL;
    uint32_t len = ctx.pkt_len();
    const unsigned char *ether_pkt = ctx.pkt();
    hal_ret_t ret = HAL_RET_OK;
    const fte::cpu_rxhdr_t* cpu_hdr = ctx.cpu_rxhdr();

    if (ether_pkt == nullptr) {
        /* Skipping as not packet to process */
        return ret;
    }

    ret = parse_dhcp_packet(ether_pkt + cpu_hdr->l3_offset,
            len - cpu_hdr->l3_offset, &decoded_packet);

    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Error parsing DHCP packet");
        return ret;
    }

    struct option_data data;
    ret = dhcp_lookup_option(decoded_packet, DHO_DHCP_MESSAGE_TYPE, &data);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Error looking up DHCP message type");
        return ret;
    }

    auto it = process_func_map.find(data.data[0]);
    if (it != process_func_map.end()) {
        it->second(decoded_packet, ctx);
    } else {
        HAL_TRACE_ERR("Cannot process option {} ",
                data.data[0]);
    }

    if (decoded_packet != NULL) {
        free_dhcp_packet(&decoded_packet);
    }

    return ret;
}

static hal_ret_t
get_dhcp_status(vrf_id_t vrf_id, ip_addr_t *ip_addr, DhcpStatus *dhcp_status)
{
    trans_ip_entry_key_t ip_entry_key;
    dhcp_trans_t *dhcp_trans;
    hal_ret_t ret;
    const dhcp_ctx *dhcp_ctx;
    ip_addr_t ipaddr = { 0 };
    std::map<dhcp_fsm_state_t, DhcpTransactionState>::iterator it;

    trans_t::init_ip_entry_key(ip_addr, vrf_id, &ip_entry_key);

    dhcp_trans = hal::eplearn::dhcp_trans_t::find_dhcp_trans_by_key(&ip_entry_key);
    if (dhcp_trans == nullptr) {
        ret = HAL_RET_ENTRY_NOT_FOUND;
        goto out;
    }

    dhcp_ctx = dhcp_trans->get_ctx();
    if (dhcp_trans->sm_->get_state() == DHCP_BOUND) {
        dhcp_status->set_renewal_time(dhcp_trans->get_timeout_remaining());
        dhcp_status->set_rebinding_time(dhcp_ctx->rebinding_time_);
    }
    dhcp_status->set_xid(dhcp_ctx->xid_);

    it = dhcp_trans_state_map.find(
            (dhcp_fsm_state_t)(dhcp_trans->sm_->get_state()));
    if (it != dhcp_trans_state_map.end()) {
        dhcp_status->set_state(it->second);
    }

    ipaddr.addr.v4_addr = dhcp_ctx->default_gateway_.s_addr;
    ip_addr_to_spec(dhcp_status->mutable_gateway_ip(), &ipaddr);

    ipaddr.addr.v4_addr = dhcp_ctx->yiaddr_.s_addr;
    ip_addr_to_spec(dhcp_status->mutable_ip_addr(), &ipaddr);

    ipaddr.addr.v4_addr = dhcp_ctx->subnet_mask_.s_addr;
    ip_addr_to_spec(dhcp_status->mutable_subnet_mask(), &ipaddr);

out:
    return ret;
}

}  // namespace eplearn
}  // namespace hal
