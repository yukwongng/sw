#include "egress.h"
#include "EGRESS_p.h"
#include "athena.h"
#include "EGRESS_p4e_redir_k.h"

struct p4e_redir_k_     k;
struct phv_             p;

%%

p4e_to_rxdma:
    /*
    phvwr           p.p4e_to_p4plus_classic_nic_ip_valid, TRUE
    phvwr           p.p4e_to_p4plus_classic_nic_valid, TRUE
    phvwr           p.capri_rxdma_intrinsic_valid, TRUE
    phvwr           p.p4i_to_p4e_header_valid, FALSE
    phvwr           p.capri_p4_intrinsic_valid, TRUE
    phvwr           p.capri_intrinsic_valid, TRUE
    */
    phvwr           p.{ p4e_to_p4plus_classic_nic_ip_valid, \
                        p4e_to_p4plus_classic_nic_valid, \
                        capri_rxdma_intrinsic_valid, \
                        p4i_to_p4e_header_valid, \
                        capri_p4_intrinsic_valid, \
                        capri_intrinsic_valid }, 0x3b


    phvwr           p.capri_intrinsic_tm_oport, TM_PORT_DMA
    phvwr           p.capri_intrinsic_lif, k.control_metadata_redir_lif
    phvwr           p.capri_rxdma_intrinsic_qtype, k.control_metadata_redir_qtype
    phvwr           p.capri_rxdma_intrinsic_qid, k.control_metadata_redir_qid
    phvwr           p.capri_rxdma_intrinsic_rx_splitter_offset, \
                    (ASICPD_GLOBAL_INTRINSIC_HDR_SZ + CAPRI_P4_INTRINSIC_HDR_SZ +\
                    /*ASICPD_RXDMA_INTRINSIC_HDR_SZ*/ 10 + P4PLUS_CLASSIC_NIC_HDR_SZ)
    phvwr           p.p4e_to_p4plus_classic_nic_p4plus_app_id, k.control_metadata_redir_app_id
    phvwr           p.p4e_to_p4plus_classic_nic_packet_len, \
                    k.p4i_to_p4e_header_packet_len
    seq             c1, k.ipv4_1_valid, TRUE
    bcf             [c1], p4e_to_rxdma_ipv4
    seq             c2, k.tcp_valid, TRUE
    seq             c1, k.ipv6_1_valid, TRUE
    nop.!c1.e

p4e_to_rxdma_ipv6:
    phvwr           p.p4e_to_p4plus_classic_nic_ip_ip_sa, k.ipv6_1_srcAddr
    phvwr           p.p4e_to_p4plus_classic_nic_ip_ip_da, k.ipv6_1_dstAddr
    seq             c3, k.udp_1_valid, TRUE
    phvwr.c2.e      p.p4e_to_p4plus_classic_nic_pkt_type, \
                        CLASSIC_NIC_PKT_TYPE_IPV6_TCP
    phvwr.c3.e      p.p4e_to_p4plus_classic_nic_pkt_type, \
                        CLASSIC_NIC_PKT_TYPE_IPV6_UDP
    nop.e
    phvwr           p.p4e_to_p4plus_classic_nic_pkt_type, \
                        CLASSIC_NIC_PKT_TYPE_IPV6
p4e_to_rxdma_ipv4:
    phvwr           p.p4e_to_p4plus_classic_nic_ip_ip_sa, k.ipv4_1_srcAddr
    phvwr           p.p4e_to_p4plus_classic_nic_ip_ip_da, k.ipv4_1_dstAddr
    phvwr.c2.e      p.p4e_to_p4plus_classic_nic_pkt_type, \
                        CLASSIC_NIC_PKT_TYPE_IPV4_TCP
    seq             c2, k.udp_1_valid, TRUE
    phvwr.c2.e      p.p4e_to_p4plus_classic_nic_pkt_type, \
                        CLASSIC_NIC_PKT_TYPE_IPV4_UDP
    nop.e
    phvwr           p.p4e_to_p4plus_classic_nic_pkt_type, \
                        CLASSIC_NIC_PKT_TYPE_IPV4

.align
p4e_to_uplink:
    /*
    phvwr           p.p4e_to_p4plus_classic_nic_ip_valid, FALSE
    phvwr           p.p4e_to_p4plus_classic_nic_valid, FALSE
    phvwr           p.capri_rxdma_intrinsic_valid, FALSE
    phvwr           p.p4i_to_p4e_header_valid, FALSE
    phvwr           p.capri_p4_intrinsic_valid, TRUE
    phvwr           p.capri_intrinsic_valid, TRUE
    */
    phvwr.e         p.{ p4e_to_p4plus_classic_nic_ip_valid, \
                        p4e_to_p4plus_classic_nic_valid, \
                        capri_rxdma_intrinsic_valid, \
                        p4i_to_p4e_header_valid, \
                        capri_p4_intrinsic_valid, \
                        capri_intrinsic_valid }, 0x03


    phvwr           p.capri_intrinsic_tm_oport, k.control_metadata_redir_oport
