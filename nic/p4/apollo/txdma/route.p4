/*****************************************************************************/
/* Route LPM lookup                                                          */
/*****************************************************************************/

#define LPM_KEY_SIZE        4
#define LPM_DATA_SIZE       2
#define LPM_S0_ENTRY_PAD    4
#define LPM_S1_ENTRY_PAD    4
#define LPM_S2_ENTRY_PAD    20
#define s0_stage            2
#define s1_stage            3
#define s2_stage            4
#define s0_name             route_lpm_s0
#define s1_name             route_lpm_s1
#define s2_name             route_lpm_s2
#define s2_name_ext         route_lpm_s2_ext
#define key_field           scratch_metadata.v4_addr
#define res_field           scratch_metadata.nh_index
#define key                 p4_to_txdma_header.lpm_dst
#define next_addr           p4_to_txdma_header.lpm_addr
#define s2_offset           txdma_control.lpm_s2_offset
#define base_addr           p4_to_txdma_header.lpm_addr
#define result              txdma_to_p4e_header.nexthop_index

action route_lpm_s2_ext() {
}

#include "../include/lpm.h"

control route {
    if (predicate_header.lpm_bypass == FALSE) {
        apply(route_lpm_s0);
        apply(route_lpm_s1);
        apply(route_lpm_s2);
    }
}
