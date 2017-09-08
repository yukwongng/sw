#include "ingress.h"
#include "INGRESS_p.h"
#include "ipsec_asm_defines.h"

struct rx_table_s1_t0_k k;
struct rx_table_s1_t0_allocate_input_desc_semaphore_d d;
struct phv_ p;

%%
        .param          esp_ipv4_tunnel_h2n_allocate_input_desc_index 
        .param          RNMDR_TABLE_BASE
        .align

esp_ipv4_tunnel_h2n_allocate_input_desc_semaphore:
    phvwri p.p42p4plus_hdr_table0_valid, 1
    //revisit when hi-word, low-word comes.
    addi r1, r0, esp_ipv4_tunnel_h2n_allocate_input_desc_index 
    srl r1, r1, 6
    phvwr p.common_te0_phv_table_pc, r1 
    phvwri p.common_te0_phv_table_raw_table_size, 3
    phvwri p.common_te0_phv_table_lock_en, 0
    sll r1, d.in_desc_ring_index, 3 
    addi r1, r1, RNMDR_TABLE_BASE 
    phvwr p.common_te0_phv_table_addr, r1
    nop.e
    nop 

