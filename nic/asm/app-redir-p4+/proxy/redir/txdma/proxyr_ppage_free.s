#include "app_redir_common.h"

struct phv_                             p;
struct proxyr_ppage_free_k              k;
struct proxyr_ppage_free_ppage_free_d   d;

/*
 * Registers usage; must match same definitions in proxyr_mpage_free.s
 */
#define r_page_addr                 r1  // page address to free
#define r_page_is_small             r2  // small page indicator
#define r_table_base                r3  // RNMPR_TABLE_BASE or RNMPR_SMALL_TABLE_BASE
#define r_table_idx                 r4  // PI index
#define r_return                    r5  // return address
#define r_scratch                   r6

%%
    .param      proxyr_page_free
    
    .align

proxyr_s6_ppage_free:

    CAPRI_CLEAR_TABLE1_VALID

    /*
     * ppage free semaphore should never be full.
     * but abort if so... proxyr_s6_desc_free will launch stats collect
     */
    sne         c1, d.pindex_full, r0
    phvwri.c1.e p.t3_s2s_inc_stat_sem_free_full, 1
    
    addi        r_page_is_small, r0, FALSE   // delay slot
    add         r_page_addr, r0, k.{to_s6_ppage_sbit0_ebit5...\
                                    to_s6_ppage_sbit30_ebit33}
    jal         r_return, proxyr_page_free
    add         r_table_idx, r0, d.pindex                   // delay slot
    nop.e
    nop
    
