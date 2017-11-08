#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_asm_defines.h"

struct tx_table_s1_t0_k k;
struct tx_table_s1_t0_ipsec_encap_txdma2_load_barco_req_d d;
struct phv_ p;

%%
        .param esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_in_desc 
        .param esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_out_desc
        .param esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_ipsec_int
        .align
esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_barco_req:
    add r1, r0, d.input_list_address
    phvwr p.txdma2_global_in_desc_addr, r1.dx
    phvwri p.app_header_table0_valid, 1
    phvwri p.common_te0_phv_table_lock_en, 1
    phvwri p.common_te0_phv_table_raw_table_size, 6 
    phvwri p.common_te0_phv_table_pc, esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_in_desc[33:6] 
    add r1, r0, d.input_list_address
    phvwr  p.common_te0_phv_table_addr, r1.dx 
    phvwri p.app_header_table1_valid, 1
    phvwri p.common_te1_phv_table_lock_en, 1
    phvwri p.common_te1_phv_table_raw_table_size, 6 
    phvwri p.common_te1_phv_table_pc, esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_out_desc[33:6] 
    add r1, r0, d.output_list_address
    phvwr  p.common_te1_phv_table_addr, r1.dx
    phvwri p.app_header_table2_valid, 1
    phvwri p.common_te2_phv_table_lock_en, 1
    phvwri p.common_te2_phv_table_raw_table_size, 6 
    phvwri p.common_te2_phv_table_pc, esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_ipsec_int[33:6] 
    add r1, r0, d.input_list_address
    add r2, r0, r1.dx
    subi r2, r2, 64 
    phvwr  p.common_te2_phv_table_addr, r2 
    nop.e
    nop
