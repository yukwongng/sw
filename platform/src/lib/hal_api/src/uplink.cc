
#include <iostream>
#include <grpc++/grpc++.h>

#include "types.grpc.pb.h"

#include "uplink.hpp"
#include "vrf.hpp"
#include "l2seg.hpp"
#include "print.hpp"

using namespace std;

std::map<uint64_t, Uplink*> Uplink::uplink_db;

Uplink *
Uplink::Factory(uplink_id_t id, bool is_oob)
{
    api_trace("Uplink Create");

    if (uplink_db.find(id) != uplink_db.end()) {
        NIC_LOG_WARN("Duplicate Create of Uplink with id: {}",
                       id);
        return NULL;
    }

    NIC_LOG_DEBUG("Id: {}, is_oob: {}", id, is_oob);


    Uplink *uplink = new Uplink(id, is_oob);

    // Store in DB for disruptive restart
    uplink_db[uplink->GetId()] = uplink;

    // Classic:
    if (hal->GetMode() == FWD_MODE_CLASSIC ||
        uplink->IsOOB()) {

        // Create VRF for each uplink
        uplink->vrf = HalVrf::Factory(uplink->IsOOB() ? types::VRF_TYPE_MANAGEMENT : types::VRF_TYPE_CUSTOMER);

#if 0
        // Create native L2seg
        uplink->native_l2seg = HalL2Segment::Factory(uplink->vrf, 0);

        // Update uplink with native l2seg
        uplink->UpdateHalWithNativeL2seg(uplink->native_l2seg->GetId());

        // Update l2seg with uplink
        uplink->native_l2seg->AddUplink(uplink);
#endif
    }

    return uplink;
}

void
Uplink::Destroy(Uplink *uplink)
{
    api_trace("Uplink Delete");

    if (hal->GetMode() == FWD_MODE_CLASSIC ||
        uplink->IsOOB()) {

#if 0
        // Remove uplink from l2seg
        uplink->native_l2seg->DelUplink(uplink);

        // Update uplink with native l2seg
        uplink->UpdateHalWithNativeL2seg(0);

        // Delete L2seg
        HalL2Segment::Destroy(uplink->GetNativeL2Seg());
#endif
        // Delete Vrf
        HalVrf::Destroy(uplink->GetVrf());
    }

    // Remove from DB
    uplink_db.erase(uplink->GetId());

    if (uplink) {
        uplink->~Uplink();
    }
}


Uplink::Uplink(uplink_id_t id, bool is_oob)
{
    NIC_LOG_DEBUG("Uplink Create: {}", id);
    this->id = id;
    this->is_oob = is_oob;
    this->num_lifs = 0;
}

Uplink::~Uplink()
{

}

int
Uplink::UpdateHalWithNativeL2seg(uint32_t native_l2seg_id)
{
    // grpc::ClientContext             context;
    grpc::Status                    status;
    InterfaceGetRequest             *req __attribute__((unused));
    InterfaceGetResponse            rsp;
    InterfaceGetRequestMsg          req_msg;
    InterfaceGetResponseMsg         rsp_msg;
    InterfaceSpec                   *if_spec;
    InterfaceResponse               if_rsp;
    InterfaceRequestMsg             if_req_msg;
    InterfaceResponseMsg            if_rsp_msg;

    req = req_msg.add_request();
    req->mutable_key_or_handle()->set_interface_id(id);
    status = hal->interface_get(req_msg, rsp_msg);
    if (status.ok()) {
        NIC_LOG_DEBUG("Updated Uplink:{} with native l2seg:{}", id, native_l2seg_id);
        for (int i = 0; i < rsp_msg.response().size(); i++) {
            rsp = rsp_msg.response(i);
            if (rsp.api_status() == types::API_STATUS_OK) {
                if (rsp.spec().type() == intf::IF_TYPE_UPLINK) {

                    if_spec = if_req_msg.add_request();
                    if_spec->CopyFrom(rsp.spec());
                    if_spec->mutable_if_uplink_info()->
                        set_native_l2segment_id(native_l2seg_id);
                    if_spec->mutable_if_uplink_info()->set_is_oob_management(is_oob);

                    status = hal->interface_update(if_req_msg, if_rsp_msg);
                    if (status.ok()) {
                        rsp = rsp_msg.response(0);
                        if (rsp.api_status() == types::API_STATUS_OK) {
                            NIC_LOG_DEBUG("Uplink {} updated with "
                                            "native_l2seg_id: {} succeeded",
                                            id, native_l2seg_id);
                        } else {
                            NIC_LOG_ERR("Failed to update uplink's:{} "
                                          "native_l2seg_id:{}: err: {}",
                                          id, native_l2seg_id,
                                          rsp.api_status());
                        }
                    } else {
                            NIC_LOG_ERR("Failed to update uplink's:{} "
                                          "native_l2seg_id:{}: err: {}, err_msg: {}",
                                          status.error_code(),
                                          status.error_message());
                    }
                }
            }
        }
    }

    return 0;
}

uint32_t
Uplink::GetId()
{
    return id;
}

uint64_t
Uplink::GetHandle()
{
    return handle;
}

uint32_t
Uplink::GetPortNum()
{
    return port_num;
}

uint32_t
Uplink::GetNumLifs()
{
    return num_lifs;
}

void
Uplink::IncNumLifs()
{
    num_lifs++;
}

void
Uplink::DecNumLifs()
{
    num_lifs--;
}

HalVrf *
Uplink::GetVrf()
{
    return vrf;
}

HalL2Segment *
Uplink::GetNativeL2Seg()
{
    return native_l2seg;
}

bool
Uplink::IsOOB()
{
    return is_oob;
}

void
Uplink::SetNativeL2Seg(HalL2Segment *l2seg)
{
    native_l2seg = l2seg;
}

void
Uplink::SetPortNum()
{
    this->port_num = port_num;
}


