#include "ingress.h"
#include "INGRESS_p.h"

struct ingress_tx_stats_k k;
struct ingress_tx_stats_d d;
struct phv_               p;

k = {
  capri_intrinsic_drop = 1;
  control_metadata_src_lif = 0xaa;
};

r5 = 0xdead0000;
