#include "../../p4/include/apulu_sacl_defines.h"
#include "nic/apollo/asm/include/apulu.h"
#include "INGRESS_p.h"
#include "ingress.h"

struct read_pktdesc2_k   k;
struct read_pktdesc2_d   d;
struct phv_              p;

%%

read_pktdesc2:
    /* Initialize rx_to_tx_hdr */
    phvwr      p.{rx_to_tx_hdr_sacl_base_addr4, \
               rx_to_tx_hdr_sip_classid4, \
               rx_to_tx_hdr_dip_classid4, \
               rx_to_tx_hdr_pad4, \
               rx_to_tx_hdr_sport_classid4, \
               rx_to_tx_hdr_dport_classid4, \
               rx_to_tx_hdr_sacl_base_addr5, \
               rx_to_tx_hdr_sip_classid5, \
               rx_to_tx_hdr_dip_classid5, \
               rx_to_tx_hdr_pad5, \
               rx_to_tx_hdr_sport_classid5, \
               rx_to_tx_hdr_dport_classid5, \
               rx_to_tx_hdr_vpc_id, \
               rx_to_tx_hdr_vnic_id, \
               rx_to_tx_hdr_iptype, \
               rx_to_tx_hdr_rx_packet, \
               rx_to_tx_hdr_payload_len, \
               rx_to_tx_hdr_stag0_classid, \
               rx_to_tx_hdr_stag1_classid, \
               rx_to_tx_hdr_stag2_classid, \
               rx_to_tx_hdr_stag3_classid, \
               rx_to_tx_hdr_stag4_classid, \
               rx_to_tx_hdr_dtag0_classid, \
               rx_to_tx_hdr_dtag1_classid, \
               rx_to_tx_hdr_dtag2_classid, \
               rx_to_tx_hdr_dtag3_classid, \
               rx_to_tx_hdr_dtag4_classid, \
               rx_to_tx_hdr_pad6, \
               rx_to_tx_hdr_pad8}, \
               d[511:(512-(offsetof(p,rx_to_tx_hdr_sacl_base_addr4) + \
                   sizeof(p.rx_to_tx_hdr_sacl_base_addr4) - \
                   offsetof(p, rx_to_tx_hdr_pad8)))]

    // Copy the first set of TAGs
    phvwr.e    p.txdma_control_stag_classid, d.read_pktdesc2_d.stag0_classid
    phvwr      p.txdma_control_dtag_classid, d.read_pktdesc2_d.dtag0_classid

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
read_pktdesc2_error:
    phvwr.e         p.capri_intr_drop, 1
    nop
