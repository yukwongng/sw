header_type capri_deparser_len_t {
    fields {
        trunc_pkt_len       : 16;
        ipv4_0_hdr_len      : 16;
        ipv4_1_hdr_len      : 16;
        ipv4_2_hdr_len      : 16;
        l4_payload_len      : 16;
    }
}

header_type key_metadata_t {
    fields {
        ktype               : 2;
        vnic_id             : 10;
        src                 : 128;
        dst                 : 128;
        smac                : 48;
        dmac                : 48;
        proto               : 8;
        sport               : 16;
        dport               : 16;
        ipv4_src            : 32;
        ipv4_dst            : 32;
        tcp_flags           : 6;
        ingress_port        : 1;
    }
}

header_type control_metadata_t {
    fields {
        forward_to_uplink                   : 1;
        redir_to_rxdma                      : 1;
        skip_flow_lkp                       : 1;
        flow_ohash_lkp                      : 1;
        direction                           : 1;
        parse_tcp_option_error              : 1;
        flow_miss                           : 1;
        session_index_valid                 : 1;
        conntrack_index_valid               : 1;
        epoch1_id_valid                     : 1;
        epoch2_id_valid                     : 1;
        throttle_pps_valid                  : 1;
        throttle_bw_valid                   : 1;
        statistics_id_valid                 : 1;
        histogram_id_valid                  : 1;
        update_checksum                     : 1;
        launch_v4                           : 1; // Dummy - never set
        strip_outer_encap_flag              : 1;
        strip_l2_header_flag                : 1;
        strip_vlan_tag_flag                 : 1;
        add_vlan_tag_flag                   : 1;
        nat_type                            : 2;
        encap_type                          : 2;
        nat_address                         : 128;
        dmac                                : 48;
        smac                                : 48;
        vlan                                : 12;
        ipv4_sa                             : 32;
        ipv4_da                             : 32;
        udp_sport                           : 16;
        udp_dport                           : 16;
        mpls_label1                         : 20;
        mpls_label2                         : 20;
        mpls_label3                         : 20;
        egress_action                       : 3;
        p4plus_app_id                       : 8;
        statistics_id                       : 9;
        histogram_id                        : 9;
        allowed_flow_state_bitmask          : 10;
        statistics_mask                     : 16;
        index                               : 22;
        session_index                       : 22;
        conntrack_index                     : 22;
        epoch1_value                        : 16;
        epoch2_value                        : 16;
        epoch1_id                           : 20;
        epoch2_id                           : 20;
        throttle_pps                        : 13;
        throttle_bw                         : 13;
        rx_packet_len                       : 14;
        tx_packet_len                       : 14;
        p4i_drop_reason                     : 32;
        p4e_drop_reason                     : 32;
    }
}

header_type scratch_metadata_t {
    fields {
        flag                : 1;
        ipv4_src            : 32;
        flow_hash           : 10;
        flow_hint           : 20;
        class_id            : 8;
        addr                : 32;
        local_vnic_tag      : 10;
        vpc_id              : 10;
        drop                : 1;
        tcp_seq_num         : 32;
        tcp_ack_num         : 32;
        tcp_win_sz          : 16;
        tcp_win_scale       : 4;
        last_seen_timestamp : 48;
        tcp_flags           : 8;
        session_stats_addr  : 34;
        hint_valid          : 1;
        cpu_flags           : 16;
        nexthop_index       : 12;
        num_nexthops        : 4;
        pad6                : 6;
        update_ip_chksum    : 1;
        update_l4_chksum    : 1;
        index_type          : 1;
        index               : 22;


        //common types
        mac                 : 48;
        ipv4                : 32;

        flow_data_pad       : 4;

        // Session info
        timestamp           : 16;
        config_epoch        : 32;
        config_substrate_src_ip : 32;
        pop_hdr_flag        : 1;
        user_pkt_rewrite_type   : 2;
        user_pkt_rewrite_ip : 128;

        // Session info - substrate encap to switch
        encap_type          : 3;
        smac                : 48;
        dmac                : 48;
        vlan                : 12;
        ip_ttl              : 8;
        ip_saddr            : 32;
        ip_daddr            : 32;
        udp_sport           : 16;
        udp_dport           : 16;
        mpls1_label         : 20;
        mpls2_label         : 20;

        // Counters
        counter_rx          : 64;
        counter_tx          : 64;

        // policer
        policer_valid       : 1;
        policer_pkt_rate    : 1;
        policer_rlimit_en   : 1;
        policer_rlimit_prof : 2;
        policer_color_aware : 1;
        policer_rsvd        : 1;
        policer_axi_wr_pend : 1;
        policer_burst       : 40;
        policer_rate        : 40;
        policer_tbkt        : 40;
        packet_len          : 16;
    }
}

header_type offset_metadata_t {
    fields {
        l2_1                : 8;
        l2_2                : 8;
        l3_1                : 8;
        l3_2                : 8;
        l4_1                : 8;
        l4_2                : 8;
        user_packet_offset  : 8;
        payload_offset      : 16;
    }
}

header_type capri_gso_csum_phv_loc_t {
    fields {
        gso_checksum            : 16;
    }
}

metadata key_metadata_t         key_metadata;
metadata control_metadata_t     control_metadata;


metadata offset_metadata_t      offset_metadata;

@pragma deparser_variable_length_header
@pragma dont_trim
metadata capri_deparser_len_t   capri_deparser_len;

@pragma gso_csum_header
@pragma dont_trim
metadata capri_gso_csum_phv_loc_t   capri_gso_csum;

@pragma scratch_metadata
metadata scratch_metadata_t     scratch_metadata;
