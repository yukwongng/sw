/*
 * Implements the TSO stage of the TxDMA P4+ pipeline
 */

#include "tcp-constants.h"
#include "tcp-phv.h"
#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "ingress.h"
#include "INGRESS_p.h"
#include "defines.h"
#include "INGRESS_s6_t0_tcp_tx_k.h"

struct phv_ p    ;
struct s6_t0_tcp_tx_k_ k    ;
struct s6_t0_tcp_tx_tso_d d    ;


%%
    .align
    .param          tcp_tx_stats_stage5_start

tcp_tso_process_start:
    /* check SESQ for pending data to be transmitted */
    sne             c6, k.common_phv_debug_dol_dont_tx, r0
    bcf             [c6], flow_tso_process_done
    phvwri.c6       p.p4_intr_global_drop, 1
    or              r1, k.to_s6_pending_tso_data, k.to_s6_pending_tso_retx
    or              r1, r1, k.common_phv_pending_ack_send
    sne             c1, r1, r0
    bal.c1          r7, tcp_write_xmit
    nop
    b               flow_tso_process_done
    nop

tcp_write_xmit:
    seq             c1, k.common_phv_pending_ack_send, 1
    bcf             [c1], dma_cmd_intrinsic

    seq             c1, k.to_s6_xmit_cursor_addr, r0
    bcf             [c1], flow_tso_process_done
    nop

dma_cmd_intrinsic:
    phvwri          p.{p4_intr_global_tm_iport...p4_intr_global_tm_oport}, \
                        (9 << 4) | 11

    // We rang the doorbell with TCP proxy service lif, but the P4
    // pipeline needs the original source_lif of the packet to derive
    // the input properties, as well as for spoofing checks
    phvwr           p.p4_intr_global_lif, d.source_lif
    phvwri          p.p4_txdma_intr_dma_cmd_ptr, TCP_PHV_TXDMA_COMMANDS_START

    // app header
#ifdef HW
    phvwr           p.{tcp_app_header_p4plus_app_id...tcp_app_header_flags}, \
                        ((P4PLUS_APPTYPE_TCPTLS << 12) | \
                            P4PLUS_TO_P4_FLAGS_LKP_INST | \
                            P4PLUS_TO_P4_FLAGS_UPDATE_IP_LEN | \
                            P4PLUS_TO_P4_FLAGS_COMPUTE_L4_CSUM | \
                            P4PLUS_TO_P4_FLAGS_UPDATE_IP_ID)

    phvwr           p.tcp_app_header_ip_id_delta, d.ip_id
    tbladd          d.ip_id, 1
#else
    phvwr           p.{tcp_app_header_p4plus_app_id...tcp_app_header_flags}, \
                        ((P4PLUS_APPTYPE_TCPTLS << 12) | \
                            P4PLUS_TO_P4_FLAGS_LKP_INST | \
                            P4PLUS_TO_P4_FLAGS_UPDATE_IP_LEN | \
                            P4PLUS_TO_P4_FLAGS_COMPUTE_L4_CSUM)

    phvwr           p.tcp_app_header_ip_id_delta, d.ip_id
#endif
    CAPRI_DMA_CMD_PHV2PKT_SETUP2(intrinsic_dma_dma_cmd, p4_intr_global_tm_iport,
                                p4_intr_global_tm_instance_type,
                                p4_txdma_intr_qid, tcp_app_header_vlan_tag)

dma_cmd_hdr:
    add             r5, k.common_phv_qstate_addr, TCP_TCB_HEADER_TEMPLATE_OFFSET

    CAPRI_OPERAND_DEBUG(d.header_len)
    CAPRI_DMA_CMD_MEM2PKT_SETUP(l2l3_header_dma_dma_cmd, r5, d.header_len[13:0])
dma_cmd_tcp_header:
    tbladd          d.quick_acks_decr, 1
    phvwr           p.{tcp_header_source_port...tcp_header_dest_port}, \
                        d.{source_port...dest_port}
    phvwrpair       p.tcp_header_seq_no, k.t0_s2s_snd_nxt, \
                        p.tcp_header_ack_no, k.to_s6_rcv_nxt

    //phvwr           p.tcp_header_data_ofs, 8		// 32 bytes
    phvwrpair       p.tcp_header_data_ofs, 5, \
                        p.tcp_header_window, k.t0_s2s_snd_wnd
	phvwr           p.{tcp_nop_opt1_kind...tcp_nop_opt2_kind}, \
                        (TCPOPT_NOP << 8 | TCPOPT_NOP)

    phvwr           p.tx2rx_quick_acks_decr, d.quick_acks_decr

    // Disable timestamps for now
    CAPRI_DMA_CMD_PHV2PKT_SETUP(tcp_header_dma_dma_cmd, tcp_header_source_port, tcp_header_urg)
    //CAPRI_DMA_CMD_PHV2PKT_SETUP(tcp_header_dma_dma_cmd, tcp_header_source_port, tcp_ts_opts_ts_ecr)

dma_cmd_data:
    seq             c2, k.to_s6_xmit_cursor_addr, r0
    /* r6 has tcp data len being sent */
    addi            r6, r0, 0
    /* We can end up taking this branch if we ended up here
     * to send pure ack and there is really no data in retx queue
     * to send
     */
    phvwri.c2       p.{tcp_header_dma_dma_pkt_eop...tcp_header_dma_dma_cmd_eop}, 3
    bcf             [c2], flow_tso_process_done
    nop

    /* Write A = xmit_cursor_addr + xmit_cursor_offset */

    add             r2, k.to_s6_xmit_cursor_addr, k.to_s6_xmit_cursor_offset

    /* Write L = min(mss, descriptor entry len) */
    seq             c1, k.to_s6_xmit_cursor_len, 0
    b.c1            pkts_sent_stats_update_start
    phvwri.c1       p.tcp_header_dma_dma_pkt_eop, 1
    slt             c1, k.to_s6_rcv_mss, k.to_s6_xmit_cursor_len
    add.c1          r6, k.to_s6_rcv_mss, r0
    add.!c1         r6, k.to_s6_xmit_cursor_len, r0

    phvwrpair       p.data_dma_dma_cmd_size, r6, \
                        p.data_dma_dma_cmd_addr, r2
    phvwri          p.{data_dma_dma_pkt_eop...data_dma_dma_cmd_type}, \
                        1 << 4 | CAPRI_DMA_COMMAND_MEM_TO_PKT
        
bytes_sent_stats_update_start:
    CAPRI_STATS_INC(bytes_sent, r6, d.bytes_sent, p.to_s7_bytes_sent)
bytes_sent_stats_update_end:

pkts_sent_stats_update_start:
    CAPRI_STATS_INC(pkts_sent, 1, d.pkts_sent, p.to_s7_pkts_sent)
pkts_sent_stats_update_end:

debug_num_phv_to_pkt_stats_update_start:
    CAPRI_STATS_INC(debug_num_phv_to_pkt, 2, d.debug_num_phv_to_pkt, p.to_s7_debug_num_phv_to_pkt)
debug_num_phv_to_pkt_stats_update_end:

debug_num_mem_to_pkt_stats_update_start:
    CAPRI_STATS_INC(debug_num_mem_to_pkt, 2, d.debug_num_mem_to_pkt, p.to_s7_debug_num_mem_to_pkt)
debug_num_mem_to_pkt_stats_update_end:

dma_cmd_write_tx2rx_shared:
    /* Set the DMA_WRITE CMD for copying tx2rx shared data from phv to mem */
    add             r5, k.common_phv_qstate_addr, TCP_TCB_TX2RX_SHARED_WRITE_OFFSET

    CAPRI_DMA_CMD_PHV2MEM_SETUP_STOP(tx2rx_dma_dma_cmd, r5, tx2rx_prr_out, tx2rx_pad1_tx2rx)
     

tcp_retx:
tcp_retx_done:

flow_tso_process_done:
    CAPRI_NEXT_TABLE0_READ_NO_TABLE_LKUP(tcp_tx_stats_stage5_start)
    nop.e
    nop

