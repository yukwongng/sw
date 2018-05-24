#include "tyche.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct rx_flow_hash_k k;
struct rx_flow_hash_d d;
struct phv_ p;

%%

rx_flow_hash_info:
    seq.c1      c1, d.rx_flow_hash_info_d.entry_valid, TRUE
    bcf         [!c1], rx_flow_hash_miss
    phvwr.c1    p.flow_action_metadata_flow_index, \
                    d.rx_flow_hash_info_d.flow_index
    phvwr       p.flow_action_metadata_prio, d.rx_flow_hash_info_d.prio
    phvwrpair   p.flow_action_metadata_parent_policer_index, \
                    d.rx_flow_hash_info_d.parent_policer_index, \
                    p.flow_action_metadata_child_policer_index, \
                    d.rx_flow_hash_info_d.child_policer_index
    phvwrpair   p.capri_intrinsic_lif[10:4], \
                    d.rx_flow_hash_info_d.lif_sbit0_ebit6, \
                    p.capri_intrinsic_lif[3:0], \
                    d.rx_flow_hash_info_d.lif_sbit7_ebit10
    phvwr       p.capri_intrinsic_tm_oport, d.rx_flow_hash_info_d.oport

    // copy hint only if is non-zero
    sne         c1, d.rx_flow_hash_info_d.hint9, r0
    phvwr.c1    p.flow_action_metadata_overflow_lkp, TRUE
    or.e        r1, d.rx_flow_hash_info_d.hint9, 1, 31
    phvwr.c1    p.flow_lkp_metadata_overflow_hash, r1

rx_flow_hash_miss:
    nop.e
    nop

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
rx_flow_hash_error:
    nop.e
    nop
