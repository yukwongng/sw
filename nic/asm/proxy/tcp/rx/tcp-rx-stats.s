/*
 *	Implements the stats stage of the TCP RxDMA P4+ pipeline
 */


#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "tcp-constants.h"	
#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_s7_t0_tcp_rx_k.h"

#define DEBUG_DOL_ATOMIC_STATS_DELTA    23
	
struct phv_ p;
struct s7_t0_tcp_rx_k_ k;
	
%%
	.align	
tcp_rx_stats_stage_start:
    CAPRI_CLEAR_TABLE_VALID(0)

// TODO: Move to multi stats update

// **NOTE: Offsets need to match definition in __tcp_rx_stats_t in 
// p4pd_tcp_proxy_api.h
bytes_rcvd_atomic_stats_update_start:
    CAPRI_ATOMIC_STATS_INCR1(bytes_rcvd, k.common_phv_qstate_addr,
                TCP_TCB_RX_STATS_OFFSET + 0 * 8, k.to_s7_bytes_rcvd)
bytes_rcvd_atomic_stats_update_done:

pkts_rcvd_atomic_stats_update_start:
    CAPRI_ATOMIC_STATS_INCR1(pkts_rcvd, k.common_phv_qstate_addr,
                TCP_TCB_RX_STATS_OFFSET + 1 * 8, k.to_s7_pkts_rcvd)
pkts_rcvd_atomic_stats_update_done:

bytes_acked_atomic_stats_update_start:
    CAPRI_ATOMIC_STATS_INCR1(bytes_acked, k.common_phv_qstate_addr,
                TCP_TCB_RX_STATS_OFFSET + 2 * 8, k.to_s7_bytes_acked)
bytes_acked_atomic_stats_update_done:


tcp_rx_stats_stage_done:
	nop.e
	nop

