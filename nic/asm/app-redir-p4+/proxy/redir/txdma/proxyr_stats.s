#include "app_redir_common.h"

struct phv_                             p;
struct proxyr_stats_k                   k;
struct proxyr_stats_err_stat_inc_d      d;

/*
 * Registers usage
 */
#define r_discard                   r1

%%
    
    .align

/*
 * Non-atomic increment of a 32-bit saturating error counters.
 * In addition, a 64-bit packet discard counter is also incremented.
 *
 * Caution: function is stage agnostic, but must not be launched from stage 7!
 */
proxyr_err_stats_inc:

    CAPRI_CLEAR_TABLE3_VALID
    
    add         r_discard, r0, r0
    or          r_discard, r_discard, k.t3_s2s_inc_stat_cb_not_ready
    tblsadd     d.stat_cb_not_ready,  k.t3_s2s_inc_stat_cb_not_ready
    or          r_discard, r_discard,          k.t3_s2s_inc_stat_null_ring_indices_addr
    tblsadd     d.stat_null_ring_indices_addr, k.t3_s2s_inc_stat_null_ring_indices_addr
    or          r_discard, r_discard, k.t3_s2s_inc_stat_aol_err
    tblsadd     d.stat_aol_err,       k.t3_s2s_inc_stat_aol_err
    or          r_discard, r_discard, k.t3_s2s_inc_stat_rxq_full
    tblsadd     d.stat_rxq_full,      k.t3_s2s_inc_stat_rxq_full
    or          r_discard, r_discard, k.t3_s2s_inc_stat_txq_empty
    tblsadd     d.stat_txq_empty,     k.t3_s2s_inc_stat_txq_empty
    or          r_discard, r_discard,  k.t3_s2s_inc_stat_sem_alloc_full
    tblsadd     d.stat_sem_alloc_full, k.t3_s2s_inc_stat_sem_alloc_full
    or          r_discard, r_discard, k.t3_s2s_inc_stat_sem_free_full
    tblsadd     d.stat_sem_free_full, k.t3_s2s_inc_stat_sem_free_full
    
    tbladd.e    d.stat_pkts_discard, r_discard
    phvwri      p.{t3_s2s_inc_stat_begin...t3_s2s_inc_stat_end}, 0 // delay slot
    
    .align

/*
 * Non-atomic increment of normal (good) counters.
 *
 * Caution: function is stage agnostic, but must not be launched from stage 7!
 */
proxyr_normal_stats_inc:

    CAPRI_CLEAR_TABLE3_VALID
    
    tbladd.e    d.stat_pkts_redir, k.t3_s2s_inc_stat_pkts_redir
    phvwri      p.{t3_s2s_inc_stat_begin...t3_s2s_inc_stat_end}, 0 // delay slot

