#ifndef _LKLSHIM_HPP_
#define _LKLSHIM_HPP_

#include "nic/include/base.h"
#include "nic/utils/slab/slab.hpp"
#include "nic/include/list.hpp"
#include "nic/hal/hal.hpp"
#include "nic/include/hal_state.hpp"
//#include "nic/hal/src/tcpcb.hpp"
//#include "nic/hal/src/tlscb.hpp"
#include "nic/utils/ht/ht.hpp"
#include "nic/include/cpupkt_headers.hpp"
#include "nic/hal/src/session.hpp"

#include <netinet/in.h>
#include <memory>
#include <map>
//#include <openssl/crypto.h>
//#include <openssl/x509.h>
//#include <openssl/pem.h>
//#include <openssl/ssl.h>
//#include <openssl/err.h>

extern "C" {
#include <netinet/in.h>
}
using hal::utils::ht_ctxt_t;
using hal::utils::slab;

namespace hal {

#define HAL_MAX_LKLSHIM_FLOWS  (1024 * 8)
#define PACKED __attribute__((__packed__))

typedef enum {
    FLOW_STATE_INVALID = 0,
    FLOW_STATE_SYN_RCVD = 1,
    FLOW_STATE_CONNECT_PENDING = 2,
    FLOW_STATE_ESTABLISHED = 3
} lklshim_flow_state_e;

typedef struct lklshim_flow_key_t_ {
    in_addr_t src_ip;
    in_addr_t dst_ip;
    uint16_t    src_port;
    uint16_t    dst_port;
    //uint16_t    protocol;
} PACKED lklshim_flow_key_t;

#define MAC_SIZE 6

typedef struct lklshim_flow_ns_t_ {
    int                  sockfd;
    lklshim_flow_state_e state;
    void                 *skbuff;
    //tcpcb_t              *tcpcb;
    //tlscb_t              *tlscb;
    char                 dev[16];
    uint8_t   src_mac[MAC_SIZE];
    uint8_t   dst_mac[MAC_SIZE];
} lklshim_flow_ns_t;

typedef struct lklshim_flow_t_ {
    lklshim_flow_key_t      key;
    lklshim_flow_ns_t       hostns;
    lklshim_flow_ns_t       netns;
    /*
     * TLS related parameters for the N-flow.
     */
//    SSL_CTX             *ssl_ctx;
//    SSL                 *ssl;
//    SSL_METHOD          *ssl_meth;

    ht_ctxt_t               ht_ctxt;
    hal::flow_direction_t   itor_dir;
    uint16_t                dst_lif;
    uint32_t                iqid;
    uint32_t                rqid;
} lklshim_flow_t;

typedef struct lklshim_listen_sockets_t_ {
    int       tcp_portnum; // Key
    int       ipv4_sockfd;
    int       ipv6_sockfd;
    ht_ctxt_t ht_ctxt;
} PACKED lklshim_listen_sockets_t;

/*
 * Extern definitions.
 */
extern ht *lklshim_flow_db;

static inline void
lklshim_make_flow_v4key (lklshim_flow_key_t *key, in_addr_t src, in_addr_t dst,
			 uint16_t sport, uint16_t dport)
{
    key->src_ip = src;
    key->dst_ip = dst;
    key->src_port = sport;
    key->dst_port = dport;
}

static inline lklshim_flow_t *
lklshim_flow_entry_lookup (lklshim_flow_key_t *key)
{
    lklshim_flow_t *lklshim_flow;

    HAL_ASSERT(key != NULL);
    lklshim_flow =
        (lklshim_flow_t *)lklshim_flow_db->lookup(key);
    return(lklshim_flow);
}

bool lklshim_process_flow_miss_rx_packet (void *pkt_skb, hal::flow_direction_t dir, uint32_t iqid, uint32_t rqid);
bool lklshim_process_flow_hit_rx_packet (void *pkt_skb, hal::flow_direction_t dir, const hal::pd::p4_to_p4plus_cpu_pkt_t* rxhdr);
void lklshim_process_tx_packet(unsigned char* pkt, unsigned int len, void* flow, bool is_connect_req, void *tcpcb);

void lklshim_flowdb_init(void);
} //namespace hal

#endif // _LKLSHIM_HPP_
