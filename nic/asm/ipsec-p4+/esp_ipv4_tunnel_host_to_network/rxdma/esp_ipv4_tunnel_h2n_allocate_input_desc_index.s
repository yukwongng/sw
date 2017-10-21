#include "ingress.h"
#include "INGRESS_p.h"
#include "ipsec_asm_defines.h"

struct rx_table_s2_t0_k k;
struct rx_table_s2_t0_allocate_input_desc_index_d d;
struct phv_ p;

%%
        .param          esp_ipv4_tunnel_h2n_update_input_desc_aol
        .align

esp_ipv4_tunnel_h2n_allocate_input_desc_index:
    phvwri p.app_header_table0_valid, 1
    phvwri p.common_te0_phv_table_pc, esp_ipv4_tunnel_h2n_update_input_desc_aol[33:6] 
    phvwri p.common_te0_phv_table_raw_table_size, 6
    phvwri p.common_te0_phv_table_lock_en, 1
    //sll r1, d.in_desc_index, DESC_SHIFT_WIDTH
    //addi r1, r1, IN_DESC_ADDR_BASE
    phvwr p.common_te0_phv_table_addr, d.in_desc_index 
    phvwr p.ipsec_int_header_in_desc, d.in_desc_index 
    phvwr p.t0_s2s_in_desc_addr, d.in_desc_index 
    nop.e
    nop  

