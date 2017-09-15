#include <hal_lock.hpp>
#include <pd_api.hpp>
#include <interface_api.hpp>
#include <if_pd.hpp>
#include <uplinkif_pd.hpp>
#include <uplinkpc_pd.hpp>
#include <enicif_pd.hpp>
#include <cpuif_pd.hpp>
#include <tunnelif_pd.hpp>

namespace hal {
namespace pd {

hal_ret_t
pd_if_create (pd_if_args_t *args)
{
    hal_ret_t       ret = HAL_RET_OK;
    intf::IfType    if_type;

    HAL_TRACE_DEBUG("PD-If:{}: If Create ", __FUNCTION__);


    if_type = hal::intf_get_if_type(args->intf);
    switch(if_type) {
        case intf::IF_TYPE_ENIC:
            ret = pd_enicif_create(args);
            break;
        case intf::IF_TYPE_UPLINK:
            ret = pd_uplinkif_create(args);
            break;
        case intf::IF_TYPE_UPLINK_PC:
            ret = pd_uplinkpc_create(args);
            break;
        case intf::IF_TYPE_TUNNEL:
            ret = pd_tunnelif_create(args);
            break;
        case intf::IF_TYPE_CPU:
            ret = pd_cpuif_create(args);
            break;
        default:
            HAL_ASSERT(0);


    }
    // Branch out for different interface types
    return ret;
}

hal_ret_t
pd_if_update (pd_if_args_t *args)
{
    hal_ret_t       ret = HAL_RET_OK;
    intf::IfType    if_type;

    HAL_TRACE_DEBUG("PD-If:{}: If Update", __FUNCTION__);


    if_type = hal::intf_get_if_type(args->intf);
    switch(if_type) {
        case intf::IF_TYPE_ENIC:
            break;
        case intf::IF_TYPE_UPLINK:
            break;
        case intf::IF_TYPE_UPLINK_PC:
            break;
        case intf::IF_TYPE_TUNNEL:
            break;
        case intf::IF_TYPE_CPU:
            break;
        default:
            HAL_ASSERT(0);


    }


    return ret;
}

// ----------------------------------------------------------------------------
// Returns the encap data and rewrite idx used for l2seg on an if. This is to be called from pd side
// Assumption: Ingress & Egress are same.
// ----------------------------------------------------------------------------
hal_ret_t
if_l2seg_get_encap_rewrite(if_t *pi_if, l2seg_t *pi_l2seg, uint32_t *encap_data,
                           uint32_t *rewrite_idx, uint32_t *tnnl_rewrite_idx)
{
    hal_ret_t ret = HAL_RET_OK;

    HAL_ASSERT_RETURN(pi_if && pi_l2seg && encap_data, HAL_RET_INVALID_ARG);
    HAL_ASSERT_RETURN(rewrite_idx && tnnl_rewrite_idx, HAL_RET_INVALID_ARG);

    switch(hal::intf_get_if_type(pi_if)) {
        case intf::IF_TYPE_ENIC:
        case intf::IF_TYPE_UPLINK:
        case intf::IF_TYPE_UPLINK_PC:
        {
                uint8_t is_tagged;
                uint16_t vlan_id;

                ret = if_l2seg_get_encap(pi_if, pi_l2seg, &is_tagged, &vlan_id);

                if (ret != HAL_RET_OK) {
                    break;
                }

                (*encap_data) = vlan_id;

                if (is_tagged) {
                    *tnnl_rewrite_idx = TUNNEL_REWRITE_ENCAP_VLAN_ID;
                } else {
                    *tnnl_rewrite_idx = TUNNEL_REWRITE_NOP_ID;
                }

                *rewrite_idx = REWRITE_NOP_ID;
                break;
        }
        case intf::IF_TYPE_TUNNEL:
            // TODO: Handle for Tunnel case
            break;
        default:
            HAL_ASSERT(0);
    }

    return ret;
}

}    // namespace pd
}    // namespace hal
