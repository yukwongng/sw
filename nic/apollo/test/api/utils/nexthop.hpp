//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------

#ifndef __TEST_API_UTILS_NH_HPP__
#define __TEST_API_UTILS_NH_HPP__

#include "nic/apollo/api/include/pds_nexthop.hpp"
#include "nic/apollo/test/api/utils/api_base.hpp"
#include "nic/apollo/test/api/utils/batch.hpp"
#include "nic/apollo/test/api/utils/feeder.hpp"
#include "nic/apollo/test/api/utils/if.hpp"

namespace test {
namespace api {

// globals
extern const pds_nh_type_t k_nh_type;
extern const uint32_t k_max_nh;

enum nexthop_attrs {
    NEXTHOP_ATTR_TYPE                   =  bit(0),
    NEXTHOP_ATTR_NH_INFO_OVERLAY        =  bit(1),
    NEXTHOP_ATTR_NH_INFO_UNDERLAY       =  bit(2),
    NEXTHOP_ATTR_NH_INFO_NH_IP          =  bit(3),
};

// NH test feeder class
class nexthop_feeder : public feeder {
public:
    pds_nexthop_spec_t spec;
    vector<pds_nexthop_info_t *> vec;

    // Constructor
    nexthop_feeder() { };
    nexthop_feeder(const nexthop_feeder& feeder) {
        init(ipaddr2str(&feeder.spec.ip), MAC_TO_UINT64(feeder.spec.mac),
             feeder.num_obj, feeder.spec.key, feeder.spec.type,
             feeder.spec.vlan, feeder.spec.vpc, feeder.spec.l3_if,
             feeder.spec.tep, feeder.stash());
    }

    // Initialize feeder with the base set of values
    void init(std::string ip_str="0.0.0.0",
              uint64_t mac=0x0E0D0A0B0200,
              uint32_t num_obj=PDS_MAX_NEXTHOP,
              pds_obj_key_t key=int2pdsobjkey(1),
              pds_nh_type_t type=k_nh_type,
              uint16_t vlan=1, pds_obj_key_t vpc=int2pdsobjkey(1),
              pds_obj_key_t l3_intf = k_l3_if_key,
              pds_obj_key_t tep = int2pdsobjkey(1),
              bool stash_ = false);

    // Iterate helper routines
    void iter_next(int width = 1);

    // Build routines
    void key_build(pds_obj_key_t *key) const;
    void spec_build(pds_nexthop_spec_t *spec) const;

    // Compare routines
    bool key_compare(const pds_obj_key_t *key) const;
    bool spec_compare(const pds_nexthop_spec_t *spec) const;
    bool status_compare(const pds_nexthop_status_t *status1,
                        const pds_nexthop_status_t *status2) const;
};

// Dump prototypes
inline std::ostream&
operator<<(std::ostream& os, const pds_nexthop_spec_t *spec) {
    os << &spec->key
       << " type: " << spec->type;
    if (spec->type == PDS_NH_TYPE_IP) {
        os << " ip: " << spec->ip
           << " mac: " << macaddr2str(spec->mac)
           << " vlan: " << spec->vlan
           << " vpc: " << std::string(spec->vpc.str());
    } else if (spec->type == PDS_NH_TYPE_UNDERLAY) {
        os << " underlay mac:" << macaddr2str(spec->underlay_mac)
           << " l3 if: " << std::string(spec->l3_if.str());
    } else if (spec->type == PDS_NH_TYPE_OVERLAY) {
        os << " tep: " << std::string(spec->tep.str());
    }
    return os;
}

inline std::ostream&
operator<<(std::ostream& os, const pds_nexthop_status_t *status) {
    os << " HW id: " << status->hw_id;
    return os;
}

inline std::ostream&
operator<<(std::ostream& os, const pds_nexthop_info_t *obj) {
    os << " NH info =>"
       << &obj->spec
       << &obj->status
       << &obj->stats
       << std::endl;
    return os;
}

inline std::ostream&
operator<<(std::ostream& os, const nexthop_feeder& obj) {
    os << "NH feeder =>" << &obj.spec << " ";
    return os;
}

// CRUD prototypes
API_CREATE(nexthop);
API_READ(nexthop);
API_READ_CMP(nexthop);
API_UPDATE(nexthop);
API_DELETE(nexthop);

// Nexthop crud helper prototypes
void nexthop_create(nexthop_feeder& feeder);
void nexthop_read(nexthop_feeder& feeder, sdk_ret_t exp_result = SDK_RET_OK);
void nexthop_update(nexthop_feeder& feeder, pds_nexthop_spec_t *spec,
                    uint64_t chg_bmap, sdk_ret_t exp_result = SDK_RET_OK);
void nexthop_delete(nexthop_feeder& feeder);

// Misc function prototypes
void sample_nexthop_setup(pds_batch_ctxt_t);
void sample_underlay_nexthop_setup(pds_batch_ctxt_t);

void sample_nexthop_teardown(pds_batch_ctxt_t);
void sample_underlay_nexthop_teardown(pds_batch_ctxt_t);

}    // namespace api
}    // namespace test

#endif    // __TEST_API_UTILS_NH_HPP__
