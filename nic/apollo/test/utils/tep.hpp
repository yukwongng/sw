//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
#ifndef __TEST_UTILS_TEP_HPP__
#define __TEST_UTILS_TEP_HPP__

#include "nic/apollo/api/include/pds_tep.hpp"
#include "nic/apollo/api/encap_utils.hpp"
#include "nic/apollo/api/tep_utils.hpp"
#include "nic/apollo/test/utils/api_base.hpp"
#include "nic/apollo/test/utils/feeder.hpp"

namespace api_test {

extern const pds_encap_t k_default_tep_encap;
extern const pds_tep_type_t k_default_tep_type;
extern const pds_encap_t k_zero_encap;

// NH test feeder class
class tep_feeder : public feeder {
public:
    uint32_t id;
    ip_addr_t remote_ip;
    ip_addr_t dipi;
    uint64_t dmac;
    pds_tep_type_t type;
    pds_encap_t encap;
    bool nat;

    // Constructor
    tep_feeder() { };
    tep_feeder(const tep_feeder& feeder) {
        init(feeder.id, ipaddr2str(&feeder.remote_ip), feeder.num_obj, feeder.encap, feeder.nat,
             feeder.type, ipaddr2str(&feeder.dipi), feeder.dmac);
    }

    // Initialize feeder with the base set of values
    void init(uint32_t id, std::string ip_str, uint32_t num_tep=PDS_MAX_TEP,
              pds_encap_t encap=k_default_tep_encap, bool nat=FALSE,
              pds_tep_type_t type=k_default_tep_type,
              std::string dipi_str="0.0.0.0", uint64_t dmac=0);

    // Iterate helper routines
    void iter_next(int width = 1);

    // Build routines
    void key_build(pds_tep_key_t *key) const;
    void spec_build(pds_tep_spec_t *spec) const;

    // Compare routines
    bool key_compare(const pds_tep_key_t *key) const;
    bool spec_compare(const pds_tep_spec_t *spec) const;
};

// Dump prototypes
inline std::ostream&
operator<<(std::ostream& os, const tep_feeder& obj) {
    os << "TEP feeder =>"
       << " ID: " << obj.id
       << " IP: " << obj.remote_ip
       << " type: " << obj.type
       << " DIPi: " << obj.dipi
       << " dmac: " << mac2str(obj.dmac)
       << " nat: " << obj.nat
       << " encap: " << pds_encap2str(obj.encap) << " ";
    return os;
}

// CRUD prototypes
API_CREATE(tep);
API_READ(tep);
API_UPDATE(tep);
API_DELETE(tep);

// Misc function prototypes
void sample_tep_setup(uint32_t tep_id=2, std::string ip_str="30.30.30.1",
                      uint32_t num_tep=PDS_MAX_TEP);
void sample_tep_validate(uint32_t tep_id=2, std::string ip_str="30.30.30.1",
                         uint32_t num_tep=PDS_MAX_TEP);
void sample_tep_teardown(uint32_t tep_id=2, std::string ip_str="30.30.30.1",
                         uint32_t num_tep=PDS_MAX_TEP);

}    // namespace api_test

#endif    // __TEST_UTILS_TEP_HPP__
