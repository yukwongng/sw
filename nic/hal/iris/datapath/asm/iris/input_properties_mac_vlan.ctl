#include "ingress.h"
#include "INGRESS_p.h"
#include "nic/hal/iris/datapath/p4/include/defines.h"

struct input_properties_mac_vlan_k k;
struct input_properties_mac_vlan_d d;
struct phv_                        p;

d = {
  input_properties_mac_vlan_d.vrf = 0x1;
  input_properties_mac_vlan_d.dir = 0x1;
  input_properties_mac_vlan_d.mdest_flow_miss_action = FLOW_MISS_ACTION_CPU;
};

k = {
  control_metadata_tm_iport = TM_PORT_UPLINK_0;
  capri_p4_intrinsic_frame_size_sbit0_ebit5 = 0;
  capri_p4_intrinsic_frame_size_sbit6_ebit13 = 250;
};

c1 = 1;
