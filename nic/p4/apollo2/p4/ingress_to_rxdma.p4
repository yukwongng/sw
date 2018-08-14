/*****************************************************************************/
/* Ingress pipeline to RxDMA pipeline                                        */
/*****************************************************************************/
action ingress_to_rxdma() {
    if ((service_header.local_ip_mapping_done == FALSE) or
        (service_header.flow_done == FALSE)) {
        add_header(service_header);
        modify_field(capri_intrinsic.tm_oport, TM_PORT_INGRESS);
    } else {
        modify_field(capri_intrinsic.tm_oport, TM_PORT_DMA);
        add_header(capri_p4_intrinsic);
        add_header(capri_rxdma_intrinsic);
        add_header(p4_to_arm_header);
        add_header(p4_to_rxdma_header);
        // Splitter offset should point to here
        modify_field(capri_rxdma_intrinsic.rx_splitter_offset,
                     (CAPRI_GLOBAL_INTRINSIC_HDR_SZ +
                      CAPRI_RXDMA_INTRINSIC_HDR_SZ +
                      APOLLO_P4_TO_RXDMA_HDR_SZ));
        add_header(predicate_header);
        add_header(p4_to_txdma_header);
        add_header(apollo_i2e_metadata);
        remove_header(service_header);

        modify_field(predicate_header.direction, control_metadata.direction);
        if (control_metadata.direction == RX_FROM_SWITCH) {
            modify_field(predicate_header.lpm_bypass, TRUE);
        }
    }
}

@pragma stage 5
table ingress_to_rxdma {
    actions {
        ingress_to_rxdma;
    }
}

control ingress_to_rxdma {
    apply(ingress_to_rxdma);
}
