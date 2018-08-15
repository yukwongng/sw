#include "egress.h"
#include "EGRESS_p.h"
#include "nic/hal/iris/datapath/p4/include/defines.h"

struct ddos_src_vf_policer_action_k     k;
struct ddos_src_vf_policer_action_d     d;
struct phv_             p;

d = {
    u.ddos_src_vf_policer_action_d.ddos_src_vf_policer_saved_color     = 0x1;
    u.ddos_src_vf_policer_action_d.ddos_src_vf_policer_dropped_packets = 0x1;
    };

k = {
    ddos_metadata_ddos_src_vf_policer_color = 0x1;
    };

