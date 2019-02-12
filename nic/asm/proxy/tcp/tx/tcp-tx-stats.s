/*
 *	Implements the stats stage of the TCP TxDMA P4+ pipeline
 */


#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "tcp-constants.h"	
#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_s6_t0_tcp_tx_k.h"
	
struct phv_ p;
struct s6_t0_tcp_tx_k_ k;
	
%%
	.align	
tcp_tx_stats_stage5_start:
        CAPRI_CLEAR_TABLE_VALID(0)

// TODO: Move to multi stats update

// **NOTE: Offsets need to match definition in __tcp_tx_stats_t in 
// p4pd_tcp_proxy_api.h
bytes_sent_atomic_stats_update_start:
        CAPRI_ATOMIC_STATS_INCR1(bytes_sent, k.common_phv_qstate_addr,
                TCP_TCB_TX_STATS_OFFSET + 0 * 8, k.to_s6_bytes_sent)
bytes_sent_atomic_stats_update_done:

pkts_sent_atomic_stats_update_start:
        CAPRI_ATOMIC_STATS_INCR1(pkts_sent, k.common_phv_qstate_addr,
                TCP_TCB_TX_STATS_OFFSET + 1 * 8, k.to_s6_pkts_sent)
pkts_sent_atomic_stats_update_done:

tcp_rx_stats_stage7_done:
	nop.e
	nop

