#define TRUE                           1
#define FALSE                          0
#define ASM_INSTRUCTION_OFFSET_MAX     (8 * 256)

#define EXCEPTION_VPORT                0x3FF

/*****************************************************************************/
/* Header sizes                                                              */
/*****************************************************************************/
#define CAPRI_GLOBAL_INTRINSIC_HDR_SZ  17
#define CAPRI_P4_INTRINSIC_HDR_SZ      5
#define CAPRI_RXDMA_INTRINSIC_HDR_SZ   15
#define P4PLUS_ROCE_HDR_SZ             20
#define P4PLUS_CLASSIC_NIC_HDR_SZ      48
#define P4PLUS_CPU_HDR_SZ              40
#define P4PLUS_CPU_PKT_SZ              37

/*****************************************************************************/
/* Ether types                                                               */
/*****************************************************************************/
#define ETHERTYPE_STAG                 0x88a8
#define ETHERTYPE_CTAG                 0x8100
#define ETHERTYPE_QINQ                 0x9100
#define ETHERTYPE_MPLS                 0x8847
#define ETHERTYPE_IPV4                 0x0800
#define ETHERTYPE_IPV6                 0x86dd
#define ETHERTYPE_ARP                  0x0806
#define ETHERTYPE_RARP                 0x8035
#define ETHERTYPE_ETHERNET             0x6558

/*****************************************************************************/
/* Tunnel types                                                              */
/*****************************************************************************/
#define INGRESS_TUNNEL_TYPE_NONE       0
#define INGRESS_TUNNEL_TYPE_VXLAN      1
#define INGRESS_TUNNEL_TYPE_GRE        2
#define INGRESS_TUNNEL_TYPE_IP_IN_IP   3
#define INGRESS_TUNNEL_TYPE_GENEVE     4
#define INGRESS_TUNNEL_TYPE_NVGRE      5
#define INGRESS_TUNNEL_TYPE_MPLS_L2VPN 6
#define INGRESS_TUNNEL_TYPE_MPLS_L3VPN 7
#define INGRESS_TUNNEL_TYPE_VXLAN_GPE  8

/*****************************************************************************/
/* IP protocol types                                                         */
/*****************************************************************************/
#define IP_PROTO_ICMP                  1
#define IP_PROTO_IGMP                  2
#define IP_PROTO_IPV4                  4
#define IP_PROTO_TCP                   6
#define IP_PROTO_UDP                   17
#define IP_PROTO_IPV6                  41
#define IP_PROTO_GRE                   47
#define IP_PROTO_IPSEC_ESP             50
#define IP_PROTO_IPSEC_AH              51
#define IP_PROTO_ICMPV6                58
#define IP_PROTO_EIGRP                 88
#define IP_PROTO_OSPF                  89
#define IP_PROTO_PIM                   103
#define IP_PROTO_VRRP                  112

/*****************************************************************************/
/* UDP services                                                              */
/*****************************************************************************/
#define UDP_PORT_VXLAN                 4789
#define UDP_PORT_VXLAN_GPE             4790
#define UDP_PORT_ROCE_V2               4791
#define UDP_PORT_GENV                  6081

/*****************************************************************************/
/* UDP option types (kind)                                                   */
/*****************************************************************************/
#define UDP_KIND_EOL                    0
#define UDP_KIND_NOP                    1
#define UDP_KIND_OCS                    2
#define UDP_KIND_MSS                    5
#define UDP_KIND_TIMESTAMP              6

/*****************************************************************************/
/* GRE services                                                              */
/*****************************************************************************/
#define GRE_PROTO_NVGRE                0x20006558
#define GRE_PROTO_ERSPAN_T3            0x22EB   /* Type III version 2 */

/*****************************************************************************/
/* TM port values                                                            */
/*****************************************************************************/
#define TM_PORT_UPLINK_0               0
#define TM_PORT_UPLINK_1               1
#define TM_PORT_UPLINK_2               2
#define TM_PORT_UPLINK_3               3
#define TM_PORT_UPLINK_4               4
#define TM_PORT_UPLINK_5               5
#define TM_PORT_UPLINK_6               6
#define TM_PORT_UPLINK_7               7
#define TM_PORT_NCSI                   8
#define TM_PORT_DMA                    9
#define TM_PORT_EGRESS                 10
#define TM_PORT_INGRESS                11

/*****************************************************************************/
/* TCP flags                                                                 */
/*****************************************************************************/
#define TCP_FLAG_CWR                   0x80
#define TCP_FLAG_ECE                   0x40
#define TCP_FLAG_URG                   0x20
#define TCP_FLAG_ACK                   0x10
#define TCP_FLAG_PSH                   0x08
#define TCP_FLAG_RST                   0x04
#define TCP_FLAG_SYN                   0x02
#define TCP_FLAG_FIN                   0x01

/*****************************************************************************/
/* Policer color                                                             */
/*****************************************************************************/
#define POLICER_COLOR_GREEN            0
#define POLICER_COLOR_YELLOW           1
#define POLICER_COLOR_RED              2

/*****************************************************************************/ 
/* Checksum flags from parser (for compiling P4 code only, don't use in ASM  */ 
/*****************************************************************************/ 
#define CSUM_HDR_UDP                   1
#define CSUM_HDR_UDP_OPT_OCS           0

/*****************************************************************************/
/* P4+ app types                                                             */
/*****************************************************************************/
#define P4PLUS_APPTYPE_DEFAULT         0
#define P4PLUS_APPTYPE_CLASSIC_NIC     1
#define P4PLUS_APPTYPE_RDMA            2
#define P4PLUS_APPTYPE_TCPTLS          3
#define P4PLUS_APPTYPE_IPSEC           4
#define P4PLUS_APPTYPE_NDE             5
#define P4PLUS_APPTYPE_STORAGE         6
#define P4PLUS_APPTYPE_TELEMETRY       7
#define P4PLUS_APPTYPE_CPU             8
#define P4PLUS_APPTYPE_RAW_REDIR       9
#define P4PLUS_APPTYPE_P4PT            10
#define P4PLUS_APP_TYPE_MIN            P4PLUS_APPTYPE_DEFAULT
#define P4PLUS_APP_TYPE_MAX            P4PLUS_APPTYPE_P4PT

/*****************************************************************************/
/* Match types                                                               */
/*****************************************************************************/
#define MATCH_ETHERNET_DST             0x00000001
#define MATCH_ETHERNET_SRC             0x00000002
#define MATCH_ETHERNET_TYPE            0x00000004
#define MATCH_CUSTOMER_VLAN_ID         0x00000008
#define MATCH_IP_SRC                   0x00000010
#define MATCH_IP_DST                   0x00000020
#define MATCH_IP_DSCP                  0x00000040
#define MATCH_IP_PROTO                 0x00000080
#define MATCH_IP_TTL                   0x00000100

#define MATCH_TRANSPORT_SRC_PORT_1     0x00000001
#define MATCH_TRANSPORT_DST_PORT_1     0x00000002
#define MATCH_TCP_FLAGS_1              0x00000004
#define MATCH_ICMP_TYPE_1              0x00000008
#define MATCH_ICMP_CODE_1              0x00000010
#define MATCH_TENANT_ID_1              0x00000020
#define MATCH_GRE_PROTO_1              0x00000040
#define MATCH_TRANSPORT_SRC_PORT_2     0x00000080
#define MATCH_TRANSPORT_DST_PORT_2     0x00000100
#define MATCH_TCP_FLAGS_2              0x00000200
#define MATCH_ICMP_TYPE_2              0x00000400
#define MATCH_ICMP_CODE_2              0x00000800
#define MATCH_TENANT_ID_2              0x00001000
#define MATCH_GRE_PROTO_2              0x00002000
#define MATCH_TRANSPORT_SRC_PORT_3     0x00004000
#define MATCH_TRANSPORT_DST_PORT_3     0x00008000
#define MATCH_TCP_FLAGS_3              0x00010000
#define MATCH_ICMP_TYPE_3              0x00020000
#define MATCH_ICMP_CODE_3              0x00040000
#define MATCH_TENANT_ID_3              0x00080000
#define MATCH_GRE_PROTO_3              0x00100000

/*****************************************************************************/
/* Match types - bit positions                                               */
/*****************************************************************************/
#define MATCH_ETHERNET_DST_BIT_POS              0
#define MATCH_ETHERNET_SRC_BIT_POS              1
#define MATCH_ETHERNET_TYPE_BIT_POS             2
#define MATCH_CUSTOMER_VLAN_ID_BIT_POS          3
#define MATCH_IP_SRC_BIT_POS                    4
#define MATCH_IP_DST_BIT_POS                    5
#define MATCH_IP_DSCP_BIT_POS                   6
#define MATCH_IP_PROTO_BIT_POS                  7
#define MATCH_IP_TTL_BIT_POS                    8

#define MATCH_TRANSPORT_SRC_PORT_1_BIT_POS      0
#define MATCH_TRANSPORT_DST_PORT_1_BIT_POS      1
#define MATCH_TCP_FLAGS_1_BIT_POS               2
#define MATCH_ICMP_TYPE_1_BIT_POS               3
#define MATCH_ICMP_CODE_1_BIT_POS               4
#define MATCH_TENANT_ID_1_BIT_POS               5
#define MATCH_GRE_PROTO_1_BIT_POS               6
#define MATCH_TRANSPORT_SRC_PORT_2_BIT_POS      7
#define MATCH_TRANSPORT_DST_PORT_2_BIT_POS      8
#define MATCH_TCP_FLAGS_2_BIT_POS               9
#define MATCH_ICMP_TYPE_2_BIT_POS               10
#define MATCH_ICMP_CODE_2_BIT_POS               11
#define MATCH_TENANT_ID_2_BIT_POS               12
#define MATCH_GRE_PROTO_2_BIT_POS               13
#define MATCH_TRANSPORT_SRC_PORT_3_BIT_POS      14
#define MATCH_TRANSPORT_DST_PORT_3_BIT_POS      15
#define MATCH_TCP_FLAGS_3_BIT_POS               16
#define MATCH_ICMP_TYPE_3_BIT_POS               17
#define MATCH_ICMP_CODE_3_BIT_POS               18
#define MATCH_TENANT_ID_3_BIT_POS               19
#define MATCH_GRE_PROTO_3_BIT_POS               20

/*****************************************************************************/
/* Transposition types                                                       */
/*****************************************************************************/
#define TRANSPOSITIONS_PUSH_ETHERNET_00            0x00000001
#define TRANSPOSITIONS_PUSH_IPV4_00                0x00000020
#define TRANSPOSITIONS_PUSH_IPV6_00                0x00000040
#define TRANSPOSITIONS_PUSH_CTAG_00                0x00000080
#define TRANSPOSITIONS_PUSH_ETHERNET_01            0x00000010
#define TRANSPOSITIONS_PUSH_IPV4_01                0x00000020
#define TRANSPOSITIONS_PUSH_IPV6_01                0x00000040
#define TRANSPOSITIONS_PUSH_CTAG_01                0x00000080

#define TRANSPOSITIONS_ENCAP_LEN_MASK_00           0x00FC0000
#define TRANSPOSITIONS_ENCAP_LEN_MASK_01           0x0003F000
#define TRANSPOSITIONS_ENCAP_LEN_SHIFT_00          18
#define TRANSPOSITIONS_ENCAP_LEN_SHIFT_01          12

#define TRANSPOSITIONS_POP_ETHERNET                0x00000001
#define TRANSPOSITIONS_POP_CTAG                    0x00000002
#define TRANSPOSITIONS_POP_IPV4                    0x00000004
#define TRANSPOSITIONS_POP_IPV6                    0x00000008
#define TRANSPOSITIONS_MODIFY_ETHERNET_SRC         0x00000010
#define TRANSPOSITIONS_MODIFY_ETHERNET_DST         0x00000020
#define TRANSPOSITIONS_MODIFY_ETHERNET_TYPE        0x00000040
#define TRANSPOSITIONS_MODIFY_CTAG                 0x00000080
#define TRANSPOSITIONS_MODIFY_IPV4_SRC             0x00000100
#define TRANSPOSITIONS_MODIFY_IPV4_DST             0x00000200
#define TRANSPOSITIONS_MODIFY_IPV4_DSCP            0x00000400
#define TRANSPOSITIONS_MODIFY_IPV4_PROTO           0x00000800
#define TRANSPOSITIONS_MODIFY_IPV4_TTL             0x00001000
#define TRANSPOSITIONS_MODIFY_IPV6_SRC             0x00002000
#define TRANSPOSITIONS_MODIFY_IPV6_DST             0x00004000
#define TRANSPOSITIONS_MODIFY_IPV6_DSCP            0x00008000
#define TRANSPOSITIONS_MODIFY_IPV6_PROTO           0x00010000
#define TRANSPOSITIONS_MODIFY_IPV6_TTL             0x00020000

#define TRANSPOSITIONS_PUSH_VXLAN_00               0x00000001
#define TRANSPOSITIONS_PUSH_UDP_00                 0x00000002
#define TRANSPOSITIONS_PUSH_VXLAN_01               0x00000004
#define TRANSPOSITIONS_PUSH_UDP_01                 0x00000008

#define TRANSPOSITIONS_POP_VXLAN_1                 0x00000001
#define TRANSPOSITIONS_POP_UDP_1                   0x00000002
#define TRANSPOSITIONS_MODIFY_UDP_SPORT_1          0x00000004
#define TRANSPOSITIONS_MODIFY_UDP_DPORT_1          0x00000008
#define TRANSPOSITIONS_MODIFY_TCP_SPORT_1          0x00000010
#define TRANSPOSITIONS_MODIFY_TCP_DPORT_1          0x00000020
#define TRANSPOSITIONS_MODIFY_ICMP_TYPE_1          0x00000040
#define TRANSPOSITIONS_MODIFY_ICMP_CODE_1          0x00000080
#define TRANSPOSITIONS_POP_VXLAN_2                 0x00000100
#define TRANSPOSITIONS_POP_UDP_2                   0x00000200
#define TRANSPOSITIONS_MODIFY_UDP_SPORT_2          0x00000400
#define TRANSPOSITIONS_MODIFY_UDP_DPORT_2          0x00000800
#define TRANSPOSITIONS_MODIFY_TCP_SPORT_2          0x00001000
#define TRANSPOSITIONS_MODIFY_TCP_DPORT_2          0x00002000
#define TRANSPOSITIONS_MODIFY_ICMP_TYPE_2          0x00004000
#define TRANSPOSITIONS_MODIFY_ICMP_CODE_2          0x00008000
#define TRANSPOSITIONS_POP_VXLAN_3                 0x00010000
#define TRANSPOSITIONS_POP_UDP_3                   0x00020000
#define TRANSPOSITIONS_MODIFY_UDP_SPORT_3          0x00040000
#define TRANSPOSITIONS_MODIFY_UDP_DPORT_3          0x00080000
#define TRANSPOSITIONS_MODIFY_TCP_SPORT_3          0x00100000
#define TRANSPOSITIONS_MODIFY_TCP_DPORT_3          0x00200000
#define TRANSPOSITIONS_MODIFY_ICMP_TYPE_3          0x00400000
#define TRANSPOSITIONS_MODIFY_ICMP_CODE_3          0x00800000
