#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_asm_defines.h"

struct tx_table_s0_t0_k k;
struct tx_table_s0_t0_esp_v4_tunnel_n2h_txdma1_initial_table_d d;
struct phv_ p;

%%
        .param esp_v4_tunnel_n2h_txdma1_allocate_barco_req_pindex 
        .param esp_v4_tunnel_n2h_get_in_desc_from_cb_cindex
        .param IPSEC_CB_BASE
        .align
ipsec_encap_txdma_initial_table:
    phvwr p.p4_intr_global_lif, k.{p4_intr_global_lif_sbit0_ebit2...p4_intr_global_lif_sbit3_ebit10}
    phvwr p.p4_intr_global_tm_iq, k.p4_intr_global_tm_iq
    phvwr p.p4_txdma_intr_qtype, k.p4_txdma_intr_qtype
    phvwr p.p4_txdma_intr_qid, k.p4_txdma_intr_qid
    phvwr p.p4_txdma_intr_qstate_addr, k.{p4_txdma_intr_qstate_addr_sbit0_ebit1...p4_txdma_intr_qstate_addr_sbit2_ebit33}

    add r1, r0, d.ipsec_cb_index
    sll r1, r1, IPSEC_CB_SHIFT_SIZE
    addi r1, r1, IPSEC_CB_BASE 
    phvwr p.txdma1_global_ipsec_cb_addr, r1 
    
    phvwr p.barco_req_command, d.barco_enc_cmd
    addi r2, r1, IPSEC_CB_IV_OFFSET
    phvwr p.barco_req_iv_address, r2
    phvwr p.barco_req_key_desc_index, d.key_index
    
    phvwri p.app_header_table0_valid, 1 
    phvwri p.common_te0_phv_table_lock_en, 1 
    phvwri p.common_te0_phv_table_pc, esp_v4_tunnel_n2h_get_in_desc_from_cb_cindex[33:6] 
    phvwri p.common_te0_phv_table_raw_table_size, 3
    add r3, r0, d.cb_cindex
    sll r3, r3, 3 
    add r3, r3, d.cb_ring_base_addr 
    phvwr p.common_te0_phv_table_addr, r3

    phvwri p.app_header_table1_valid, 1 
    phvwri p.common_te1_phv_table_lock_en, 1 
    phvwri p.common_te1_phv_table_pc, esp_v4_tunnel_n2h_txdma1_allocate_barco_req_pindex[33:6] 
    phvwri p.common_te1_phv_table_raw_table_size, 2
    phvwri p.common_te1_phv_table_addr, BRQ_REQ_SEMAPHORE_ADDR 
    nop.e
    nop
    
