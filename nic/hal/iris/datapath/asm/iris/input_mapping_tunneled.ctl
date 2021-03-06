#include "ingress.h"
#include "INGRESS_p.h"
#include "nic/hal/iris/datapath/p4/include/defines.h"

struct input_mapping_tunneled_k k;
struct phv_                   p;

k = {
  inner_ethernet_srcAddr = 0x00aabbccddee;
  inner_ethernet_dstAddr = 0x00aabbccddff;
  inner_ipv4_protocol = 17;
  inner_ipv4_srcAddr = 0xddee;
  inner_ipv4_dstAddr = 0xbbccddff;
};
