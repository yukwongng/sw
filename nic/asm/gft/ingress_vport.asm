#include "ingress.h"
#include "INGRESS_p.h"
#include "../../p4/gft/include/defines.h"

struct ingress_vport_k k;
struct ingress_vport_d d;
struct phv_ p;

%%

ingress_vport:
    sub         r1, k.{capri_p4_intrinsic_frame_size_sbit0_ebit5,\
                       capri_p4_intrinsic_frame_size_sbit6_ebit13}, \
                    CAPRI_GLOBAL_INTRINSIC_HDR_SZ
    phvwr       p.capri_p4_intrinsic_packet_len, r1
    phvwr       p.capri_intrinsic_tm_oport, TM_PORT_UPLINK_0
    cmov.e      r1, c1, d.ingress_vport_d.vport, EXCEPTION_VPORT
    phvwr       p.capri_intrinsic_lif, r1
