#include "tyche.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct rx_key1_k k;
struct rx_key1_d d;
struct phv_ p;

%%

rx_key1:
    nop.!c1.e
    crestore    [c4-c1], d.rx_key1_d.match_fields[3:0], 0xF
    phvwr.c1    p.flow_lkp_metadata_ethernet_dst_1, k.ethernet_1_dstAddr
    add         r7, r0, r0
    or.c2       r7, r7, k.ethernet_1_srcAddr, 16
    or.c3       r7, r7, k.ethernet_1_etherType
    phvwr       p.{flow_lkp_metadata_ethernet_src_1, \
                   flow_lkp_metadata_ethernet_type_1}, r7
    phvwr.c4    p.flow_lkp_metadata_ctag_1, \
                    k.{ctag_1_vid_sbit0_ebit3,ctag_1_vid_sbit4_ebit11}
    bbeq        k.ipv4_1_valid, TRUE, rx_key1_ipv4
    crestore    [c5-c1], d.rx_key1_d.match_fields[8:4], 0x1F
    bbeq        k.ipv6_1_valid, TRUE, rx_key1_ipv6
    nop
    nop.e
    nop

rx_key1_ipv4:
    phvwr.c1    p.flow_lkp_metadata_ip_src_1, \
                    k.{ipv4_1_srcAddr_sbit0_ebit7...ipv4_1_srcAddr_sbit16_ebit31}
    phvwr.c2    p.flow_lkp_metadata_ip_dst_1, k.ipv4_1_dstAddr
    add         r7, r0, r0
    or.c3       r7, r7, k.ipv4_1_diffserv, 16
    or.c4       r7, r7, k.ipv4_1_protocol, 8
    or.c5       r7, r7, k.ipv4_1_ttl
    phvwr.f.e   p.{flow_lkp_metadata_ip_dscp_1,flow_lkp_metadata_ip_proto_1, \
                   flow_lkp_metadata_ip_ttl_1}, r7
    nop

rx_key1_ipv6:
    phvwr.c1    p.flow_lkp_metadata_ip_src_1[127:72], \
                    k.{ipv6_1_srcAddr_sbit0_ebit15...ipv6_1_srcAddr_sbit48_ebit55}
    phvwr.c1    p.flow_lkp_metadata_ip_src_1[71:0], \
                    k.{ipv6_1_srcAddr_sbit56_ebit127}
    phvwr.c2    p.flow_lkp_metadata_ip_dst_1, k.ipv6_1_dstAddr
    add         r7, r0, r0
    or.c3       r7, r7, k.{ipv6_1_trafficClass_sbit0_ebit3,\
                           ipv6_1_trafficClass_sbit4_ebit7}, 16
    or.c4       r7, r7, k.ipv6_1_nextHdr, 8
    or.c5       r7, r7, k.ipv6_1_hopLimit
    phvwr.f.e   p.{flow_lkp_metadata_ip_dscp_1,flow_lkp_metadata_ip_proto_1, \
                   flow_lkp_metadata_ip_ttl_1}, r7
    nop

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
rx_key1_error:
    nop.e
    nop
