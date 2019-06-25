control iflow {}
control read_qstate {}
control mapping {}
control vnic_info_txdma {}
control remote_46_mapping {}
control route_lookup {}
control rfc {}
control pkt_dma {}

#define txdma_predicate             app_header
#define flow_enable                 table0_valid
#define pass_two                    table1_valid
#define lpm1_enable                 table2_valid
#define cps_path_en                 table3_valid
