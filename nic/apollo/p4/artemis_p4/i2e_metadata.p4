header_type artemis_i2e_metadata_t {
    fields {
        entropy_hash        : 32;
        pa_or_ca_xlate_idx  : 16;
        public_xlate_idx    : 16;
        ca6_xlate_idx       : 16;
        service_xlate_idx   : 16;
        service_xlate_port  : 16;
        session_index       : 24;
        vnic_id             : 8;
    }
}
