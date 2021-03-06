#include "nic/include/base.hpp"
#include "nic/hal/hal.hpp"
#include "nic/sdk/include/sdk/lock.hpp"
#include "nic/hal/plugins/cfg/nw/session.hpp"
#include "nic/include/fte.hpp"
#include "nic/hal/iris/include/hal_state.hpp"
#include "nic/hal/plugins/cfg/ipsec/ipsec.hpp"
#include "nic/hal/plugins/cfg/nw/vrf.hpp"
#include "nic/include/pd_api.hpp"
#include "nic/hal/plugins/cfg/nw/vrf_api.hpp"
#include "nic/hal/src/utils/utils.hpp"
#include <google/protobuf/util/json_util.h>

namespace hal {
void *
ipsec_sa_get_key_func (void *entry)
{
    SDK_ASSERT(entry != NULL);
    HAL_TRACE_DEBUG("Got id as {}\n", (((ipsec_sa_t *)entry)->sa_id));
    return (void *)&(((ipsec_sa_t *)entry)->sa_id);
}

uint32_t
ipsec_sa_key_size ()
{
    return sizeof(ipsec_sa_id_t);
}

void *
ipsec_sa_get_handle_key_func (void *entry)
{
    SDK_ASSERT(entry != NULL);
    return (void *)&(((ipsec_sa_t *)entry)->hal_handle);
}

uint32_t
ipsec_sa_handle_key_size ()
{
    return sizeof(hal_handle_t);
}

//------------------------------------------------------------------------------
// validate an incoming IPSECCB create request
// TODO:
// 1. check if IPSECCB exists already
//------------------------------------------------------------------------------
static hal_ret_t
validate_ipsec_sa_encrypt_create (IpsecSAEncrypt& spec, IpsecSAEncryptResponse *rsp)
{
    // must have key-handle set
    if (!spec.has_key_or_handle()) {
        rsp->set_api_status(types::API_STATUS_IPSEC_CB_ID_INVALID);
        return HAL_RET_INVALID_ARG;
    }
    auto kh = spec.key_or_handle();
    // must have key in the key-handle
    if (spec.key_or_handle().key_or_handle_case() !=
            IpsecSAEncryptKeyHandle::kCbId) {
        rsp->set_api_status(types::API_STATUS_IPSEC_CB_ID_INVALID);
        return HAL_RET_INVALID_ARG;
    }
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// insert this IPSEC CB in all meta data structures
//------------------------------------------------------------------------------
static inline hal_ret_t
add_ipsec_sa_to_db (ipsec_sa_t *ipsec)
{
    g_hal_state->ipsec_sa_id_ht()->insert(ipsec, &ipsec->ht_ctxt);
    return HAL_RET_OK;
}

static inline hal_ret_t
del_ipsec_sa_from_db(ipsec_sa_t *ipsec)
{
    g_hal_state->ipsec_sa_id_ht()->remove(&ipsec->ht_ctxt);
    return HAL_RET_OK;
}


static inline void
ipsec_sa_encrypt_spec_dump (IpsecSAEncrypt& spec)
{
    std::string    ipsec_sa_encrypt_cfg_str;

    if (hal::utils::hal_trace_level() < ::utils::trace_debug)  {
        return;
    }

    google::protobuf::util::MessageToJsonString(spec, &ipsec_sa_encrypt_cfg_str);
    HAL_TRACE_DEBUG("IPSec SA Encrypt Config:");
    HAL_TRACE_DEBUG("{}", ipsec_sa_encrypt_cfg_str.c_str());
    return;
}

//------------------------------------------------------------------------------
// process a IPSEC CB create request
// TODO: if IPSEC CB exists, treat this as modify (vrf id in the meta must
// match though)
//------------------------------------------------------------------------------
hal_ret_t
ipsec_saencrypt_create (IpsecSAEncrypt& spec, IpsecSAEncryptResponse *rsp)
{
    hal_ret_t              ret = HAL_RET_OK;
    ipsec_sa_t                *ipsec;
    pd::pd_ipsec_encrypt_create_args_t    pd_ipsec_encrypt_args;
    pd::pd_func_args_t pd_func_args = {0};
    ep_t *sep, *dep;
    mac_addr_t *smac = NULL, *dmac = NULL;
    vrf_t   *vrf;
    mac_addr_t smac1 = {0x0, 0xee, 0xff, 0x0, 0x0, 0x02};
    mac_addr_t dmac1 = {0x0, 0xee, 0xff, 0x0, 0x0, 0x03};

    ipsec_sa_encrypt_spec_dump(spec);
    // validate the request message
    ret = validate_ipsec_sa_encrypt_create(spec, rsp);

    if ((spec.encryption_algorithm() != ipsec::ENCRYPTION_ALGORITHM_AES_GCM_256) ||
        (spec.authentication_algorithm() != ipsec::AUTHENTICATION_AES_GCM)) {
        HAL_TRACE_DEBUG("Unsupported Encyption or Authentication Algo. EncAlgo {} AuthAlgo{}", spec.encryption_algorithm(), spec.authentication_algorithm());
        return HAL_RET_IPSEC_ALGO_NOT_SUPPORTED;
    }
    vrf = vrf_lookup_by_id(spec.tep_vrf().vrf_id());
    if (vrf) {
        HAL_TRACE_DEBUG("vrf success id = {} ", vrf->vrf_id);
    } else {
        HAL_TRACE_ERR("Vrf not found for vrf-id {}", spec.tep_vrf().vrf_id());
        return HAL_RET_VRF_ID_INVALID; 
    }

    if (spec.key_or_handle().cb_id() > HAL_MAX_IPSEC_SUPP_SA) {
        return HAL_RET_INVALID_ARG;
    }

    ipsec = ipsec_sa_alloc_init();
    if (ipsec == NULL) {
        rsp->set_api_status(types::API_STATUS_OUT_OF_MEM);
        return HAL_RET_OOM;
    }

    ipsec->sa_id = spec.key_or_handle().cb_id();
    ipsec->vrf = vrf->vrf_id;

    HAL_TRACE_DEBUG("Got with SA ID {} vrf id {}", ipsec->sa_id, ipsec->vrf);


    ipsec->iv_size = IPSEC_DEF_IV_SIZE;
    ipsec->block_size = IPSEC_AES_GCM_DEF_BLOCK_SIZE;
    ipsec->icv_size = IPSEC_AES_GCM_DEF_ICV_SIZE;
    ipsec->esn_hi = ipsec->esn_lo = 0;

    ipsec->barco_enc_cmd = IPSEC_BARCO_ENCRYPT_AES_GCM_256;
    ipsec->key_size = IPSEC_AES_GCM_DEF_KEY_SIZE;
    ipsec->key_type = types::CryptoKeyType::CRYPTO_KEY_TYPE_AES256;

    ipsec->iv = spec.iv();
    ipsec->iv_salt = spec.salt();
    ipsec->spi = spec.spi();

    memcpy((uint8_t*)ipsec->key, (uint8_t*)spec.encryption_key().key().c_str(), 32);

    ip_addr_spec_to_ip_addr(&ipsec->tunnel_sip4, spec.local_gateway_ip());
    ip_addr_spec_to_ip_addr(&ipsec->tunnel_dip4, spec.remote_gateway_ip());

    sep = find_ep_by_v4_key(ipsec->vrf, ipsec->tunnel_sip4.addr.v4_addr);
    if (sep) {
        smac = ep_get_mac_addr(sep);
        if (smac) {
            memcpy(ipsec->smac, smac, ETH_ADDR_LEN);
        }
    } else {
        memcpy(ipsec->smac, smac1, ETH_ADDR_LEN);
        HAL_TRACE_DEBUG("Src EP Lookup failed in vrf {} for IP {}", ipsec->vrf, ipaddr2str(&ipsec->tunnel_sip4));
    }
    dep = find_ep_by_v4_key(ipsec->vrf, ipsec->tunnel_dip4.addr.v4_addr);
    if (dep) {
        dmac = ep_get_mac_addr(dep);
        if (dmac) {
            memcpy(ipsec->dmac, dmac, ETH_ADDR_LEN);
        }
    } else {
        memcpy(ipsec->dmac, dmac1, ETH_ADDR_LEN);
        HAL_TRACE_DEBUG("Dest EP Lookup failed in vrf {} for IP {}", ipsec->vrf, ipaddr2str(&ipsec->tunnel_dip4));
    }

    ipsec->hal_handle = hal_alloc_handle();

    // allocate all PD resources and finish programming
    pd::pd_ipsec_encrypt_create_args_init(&pd_ipsec_encrypt_args);
    pd_ipsec_encrypt_args.ipsec_sa = ipsec;
    pd_func_args.pd_ipsec_encrypt_create = &pd_ipsec_encrypt_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_IPSEC_ENCRYPT_CREATE, &pd_func_args);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD IPSEC CB create failure, err : {}", ret);
        rsp->set_api_status(types::API_STATUS_HW_PROG_ERR);
        goto cleanup;
    }
    // add this L2 segment to our db
    ret = add_ipsec_sa_to_db(ipsec);
    SDK_ASSERT(ret == HAL_RET_OK);

    // prepare the response
    rsp->set_api_status(types::API_STATUS_OK);
    rsp->mutable_ipsec_sa_status()->set_ipsec_sa_handle(ipsec->hal_handle);

    HAL_TRACE_DEBUG("Returning Success for SA ID {}", ipsec->sa_id);
    return HAL_RET_OK;

cleanup:

    ipsec_sa_free(ipsec);
    return ret;
}

//------------------------------------------------------------------------------
// process a IPSEC CB update request
//------------------------------------------------------------------------------
hal_ret_t
ipsec_saencrypt_update (IpsecSAEncrypt& spec, IpsecSAEncryptResponse *rsp)
{
    hal_ret_t              ret = HAL_RET_OK;
    ipsec_sa_t*               ipsec;
    pd::pd_ipsec_encrypt_update_args_t    pd_ipsec_encrypt_args;
    pd::pd_func_args_t pd_func_args = {0};
    ep_t *sep, *dep;
    mac_addr_t *smac = NULL, *dmac = NULL;
    vrf_t   *vrf;

    ipsec_sa_encrypt_spec_dump(spec);
    auto kh = spec.key_or_handle();
    HAL_TRACE_DEBUG("Got with SA ID {}", kh.cb_id());

    ipsec = find_ipsec_sa_by_id(kh.cb_id());
    if (ipsec == NULL) {
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_IPSEC_CB_NOT_FOUND;
    }
    if ((spec.encryption_algorithm() != ipsec::ENCRYPTION_ALGORITHM_AES_GCM_256) ||
        (spec.authentication_algorithm() != ipsec::AUTHENTICATION_AES_GCM)) {
        HAL_TRACE_DEBUG("Unsupported Encyption or Authentication Algo. EncAlgo {} AuthAlgo{}", spec.encryption_algorithm(), spec.authentication_algorithm());
        return HAL_RET_IPSEC_ALGO_NOT_SUPPORTED;
    }
    vrf = vrf_lookup_by_id(spec.tep_vrf().vrf_id());
    if (vrf) {
        ipsec->vrf = vrf->vrf_id;
        HAL_TRACE_DEBUG("vrf success id = {}", ipsec->vrf);
    } else {
        HAL_TRACE_ERR("Vrf not found for vrf-id {}", spec.tep_vrf().vrf_id());
        return HAL_RET_VRF_ID_INVALID; 
    }


    ipsec->iv_size = IPSEC_DEF_IV_SIZE;
    ipsec->block_size = IPSEC_AES_GCM_DEF_BLOCK_SIZE;
    ipsec->icv_size = IPSEC_AES_GCM_DEF_ICV_SIZE;
    ipsec->esn_hi = ipsec->esn_lo = 0;

    ipsec->barco_enc_cmd = IPSEC_BARCO_ENCRYPT_AES_GCM_256;

    pd::pd_ipsec_encrypt_update_args_init(&pd_ipsec_encrypt_args);
    pd_ipsec_encrypt_args.ipsec_sa = ipsec;

    ipsec->iv = spec.iv();
    ipsec->iv_salt = spec.salt();
    ipsec->spi = spec.spi();

    ip_addr_spec_to_ip_addr(&ipsec->tunnel_sip4, spec.local_gateway_ip());
    ip_addr_spec_to_ip_addr(&ipsec->tunnel_dip4, spec.remote_gateway_ip());

    sep = find_ep_by_v4_key(ipsec->vrf, ipsec->tunnel_sip4.addr.v4_addr);
    if (sep) {
        smac = ep_get_mac_addr(sep);
        if (smac) {
            memcpy(ipsec->smac, smac, ETH_ADDR_LEN);
        }
    } else {
        HAL_TRACE_DEBUG("Src EP Lookup failed in vrf {} for IP {}", ipsec->vrf, ipaddr2str(&ipsec->tunnel_sip4));
    }
    dep = find_ep_by_v4_key(ipsec->vrf, ipsec->tunnel_dip4.addr.v4_addr);
    if (dep) {
        dmac = ep_get_mac_addr(dep);
        if (dmac) {
            memcpy(ipsec->dmac, dmac, ETH_ADDR_LEN);
        }
    } else {
        HAL_TRACE_DEBUG("Dest EP Lookup failed in vrf {} for IP {}", ipsec->vrf, ipaddr2str(&ipsec->tunnel_dip4));
    }
    pd_func_args.pd_ipsec_encrypt_update = &pd_ipsec_encrypt_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_IPSEC_ENCRYPT_UPDATE, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD IPSECCB: Update Failed, err: {}", ret);
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }
    rsp->set_api_status(types::API_STATUS_OK);

    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a IPSEC CB get request
//------------------------------------------------------------------------------
hal_ret_t
ipsec_saencrypt_get (IpsecSAEncryptGetRequest& req, IpsecSAEncryptGetResponseMsg *resp)
{
    hal_ret_t              ret = HAL_RET_OK;
    ipsec_sa_t                ripsec;
    ipsec_sa_t*               ipsec;
    pd::pd_ipsec_encrypt_get_args_t    pd_ipsec_encrypt_args;
    pd::pd_func_args_t pd_func_args = {0};
    IpsecSAEncryptGetResponse *rsp = resp->add_response();

    auto kh = req.key_or_handle();

    HAL_TRACE_DEBUG("Got with SA ID {}\n", kh.cb_id());
    ipsec = find_ipsec_sa_by_id(kh.cb_id());
    if (ipsec == NULL) {
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_IPSEC_CB_NOT_FOUND;
    }

    ipsec_sa_init(&ripsec);
    ripsec.sa_id = ipsec->sa_id;
    pd::pd_ipsec_encrypt_get_args_init(&pd_ipsec_encrypt_args);
    pd_ipsec_encrypt_args.ipsec_sa = &ripsec;

    pd_func_args.pd_ipsec_encrypt_get = &pd_ipsec_encrypt_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_IPSEC_ENCRYPT_GET, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD IPSECCB: Failed to get, err: {}", ret);
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }

    // fill config spec of this IPSEC CB
    rsp->mutable_spec()->mutable_key_or_handle()->set_cb_id(ripsec.sa_id);
    if (ripsec.barco_enc_cmd == IPSEC_BARCO_ENCRYPT_AES_GCM_256) {
        rsp->mutable_spec()->set_authentication_algorithm(ipsec::AUTHENTICATION_AES_GCM);
        rsp->mutable_spec()->set_encryption_algorithm(ipsec::ENCRYPTION_ALGORITHM_AES_GCM_256);
    }
    rsp->mutable_spec()->set_salt(ripsec.iv_salt);
    rsp->mutable_spec()->set_iv(ripsec.iv);
    rsp->mutable_spec()->set_spi(ripsec.spi);
    rsp->mutable_spec()->set_iv_size(ripsec.iv_size);
    rsp->mutable_spec()->set_icv_size(ripsec.icv_size);
    uint64_t seq_no  = ripsec.esn_hi;
    seq_no = seq_no << 32;
    seq_no = seq_no | ripsec.esn_lo;
    rsp->mutable_spec()->set_seq_no(seq_no);
    rsp->mutable_spec()->set_iv(ripsec.iv);
    rsp->mutable_spec()->set_key_index(ripsec.key_index);
    rsp->mutable_spec()->mutable_encryption_key()->set_key(ripsec.key, IPSEC_AES_GCM_DEF_KEY_SIZE);
    rsp->mutable_spec()->mutable_authentication_key()->set_key(ripsec.key, IPSEC_AES_GCM_DEF_KEY_SIZE);
    rsp->mutable_spec()->mutable_tep_vrf()->set_vrf_id(ipsec->vrf);
    HAL_TRACE_DEBUG("tep_vrf {}", ipsec->vrf);
    rsp->mutable_spec()->set_total_pkts(ripsec.total_pkts);
    rsp->mutable_spec()->set_total_bytes(ripsec.total_bytes);
    rsp->mutable_spec()->set_total_drops(ripsec.total_drops);
    rsp->mutable_spec()->set_total_rx_pkts(ripsec.total_rx_pkts);
    rsp->mutable_spec()->set_total_rx_bytes(ripsec.total_rx_bytes);
    rsp->mutable_spec()->set_total_rx_drops(ripsec.total_rx_drops);

    ripsec.tunnel_sip4.af = IP_AF_IPV4;
    ip_addr_to_spec(rsp->mutable_spec()->mutable_local_gateway_ip(), &ripsec.tunnel_sip4);
    ripsec.tunnel_dip4.af = IP_AF_IPV4;
    ip_addr_to_spec(rsp->mutable_spec()->mutable_remote_gateway_ip(), &ripsec.tunnel_dip4);

    // fill operational state of this IPSEC CB
    rsp->mutable_status()->set_ipsec_sa_handle(ipsec->hal_handle);

    // fill stats of this IPSEC CB
    rsp->set_api_status(types::API_STATUS_OK);

    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a IPSEC CB delete request
//------------------------------------------------------------------------------
hal_ret_t
ipsec_saencrypt_delete (ipsec::IpsecSAEncryptDeleteRequest& req, ipsec::IpsecSAEncryptDeleteResponse *rsp)
{
    hal_ret_t              ret = HAL_RET_OK;
    ipsec_sa_t*               ipsec;
    pd::pd_ipsec_encrypt_delete_args_t    pd_ipsec_encrypt_args;
    pd::pd_func_args_t pd_func_args = {0};

    auto kh = req.key_or_handle();
    ipsec = find_ipsec_sa_by_id(kh.cb_id());
    if (ipsec == NULL) {
        rsp->set_api_status(types::API_STATUS_OK);
        return HAL_RET_OK;
    }

    pd::pd_ipsec_encrypt_delete_args_init(&pd_ipsec_encrypt_args);
    pd_ipsec_encrypt_args.ipsec_sa = ipsec;

    pd_func_args.pd_ipsec_encrypt_delete = &pd_ipsec_encrypt_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_IPSEC_ENCRYPT_DELETE, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD IPSECCB: delete Failed, err: {}", ret);
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }

    // fill stats of this IPSEC CB
    rsp->set_api_status(types::API_STATUS_OK);
    del_ipsec_sa_from_db(ipsec);
    ipsec_sa_free(ipsec);

    return HAL_RET_OK;
}

hal_ret_t
ipsec_global_statistics_get (ipsec::IpsecGlobalStatisticsGetRequest& req, 
                             ipsec::IpsecGlobalStatisticsGetResponseMsg *resp)
{
    hal_ret_t              ret = HAL_RET_OK;
    IpsecGlobalStatisticsGetResponse *rsp = resp->add_response();

    IpsecGlobalStatisticsGetResponse spec;
    ipsec_global_stats_cb_t stats_cb;

    pd::pd_func_args_t pd_func_args = {0};
    pd::pd_ipsec_global_stats_get_args_t pd_ipsec_stats_get_args;  

    auto clear_on_read = req.clear_on_read();
    pd_ipsec_stats_get_args.clear_on_read = clear_on_read;
    pd_ipsec_stats_get_args.stats_cb = &stats_cb;
    pd_func_args.pd_ipsec_global_stats_get = &pd_ipsec_stats_get_args;

    if (clear_on_read != 0) {

    }

    ret = pd::hal_pd_call(pd::PD_FUNC_ID_IPSEC_GLOBAL_STATS_GET, &pd_func_args);
    rsp->set_api_status(types::API_STATUS_OK);
    rsp->mutable_spec()->set_encrypt_input_desc_errors(stats_cb.encrypt_input_desc_errors);
    rsp->mutable_spec()->set_encrypt_output_desc_errors(stats_cb.encrypt_output_desc_errors);
    rsp->mutable_spec()->set_encrypt_cb_ring_base_errors(stats_cb.encrypt_cb_ring_base_errors);
    rsp->mutable_spec()->set_encrypt_input_page_errors(stats_cb.encrypt_input_page_errors);
    rsp->mutable_spec()->set_encrypt_barco_req_addr_errors(stats_cb.encrypt_barco_req_addr_errors);
    rsp->mutable_spec()->set_encrypt_barco_cb_base_errors(stats_cb.encrypt_barco_cb_base_errors);
    rsp->mutable_spec()->set_encrypt_pad_addr_errors(stats_cb.encrypt_pad_addr_errors);
    rsp->mutable_spec()->set_encrypt_tail_bytes_errors(stats_cb.encrypt_tail_bytes_errors);
    rsp->mutable_spec()->set_encrypt_output_page_errors(stats_cb.encrypt_output_page_errors);
    rsp->mutable_spec()->set_encrypt_stage4_inpage_errors(stats_cb.encrypt_stage4_inpage_errors);
    rsp->mutable_spec()->set_encrypt_table0_inpage_errors(stats_cb.encrypt_table0_inpage_errors);
    rsp->mutable_spec()->set_encrypt_table2_inpage_errors(stats_cb.encrypt_table2_inpage_errors);
    rsp->mutable_spec()->set_encrypt_table3_inpage_errors(stats_cb.encrypt_table3_inpage_errors);
    rsp->mutable_spec()->set_encrypt_bad_barco_addr_errors(stats_cb.encrypt_bad_barco_addr_errors);
    rsp->mutable_spec()->set_encrypt_barco_full_errors(stats_cb.encrypt_barco_full_errors);
    rsp->mutable_spec()->set_encrypt_cb_ring_dma_errors(stats_cb.encrypt_cb_ring_dma_errors);
    rsp->mutable_spec()->set_encrypt_desc_exhaust_errors(stats_cb.encrypt_desc_exhaust_errors);
    rsp->mutable_spec()->set_encrypt_txdma1_enter_counters(stats_cb.encrypt_txdma1_enter_counters);
    rsp->mutable_spec()->set_encrypt_txdma2_enter_counters(stats_cb.encrypt_txdma2_enter_counters);
    rsp->mutable_spec()->set_encrypt_txdma1_dummy_errors(stats_cb.encrypt_txdma1_dummy_errors);
    rsp->mutable_spec()->set_encrypt_rxdma_dummy_desc_errors(stats_cb.encrypt_rxdma_dummy_desc_errors);
    rsp->mutable_spec()->set_encrypt_rxdma_enter_counters(stats_cb.encrypt_rxdma_enter_counters);
    rsp->mutable_spec()->set_encrypt_barco_bad_indesc_errors(stats_cb.encrypt_barco_bad_indesc_errors);
    rsp->mutable_spec()->set_encrypt_barco_bad_outdesc_errors(stats_cb.encrypt_barco_bad_outdesc_errors);
    rsp->mutable_spec()->set_encrypt_txdma2_bad_indesc_free_errors(stats_cb.encrypt_txdma2_bad_indesc_free_errors);
    rsp->mutable_spec()->set_encrypt_txdma2_bad_outdesc_free_errors(stats_cb.encrypt_txdma2_bad_outdesc_free_errors);
    rsp->mutable_spec()->set_encrypt_txdma1_bad_indesc_free_errors(stats_cb.encrypt_txdma1_bad_indesc_free_errors);
    rsp->mutable_spec()->set_encrypt_txdma1_bad_outdesc_free_errors(stats_cb.encrypt_txdma1_bad_outdesc_free_errors);
    rsp->mutable_spec()->set_encrypt_txdma1_sem_free_errors(stats_cb.encrypt_txdma1_sem_free_errors);
    rsp->mutable_spec()->set_encrypt_txdma2_sem_free_errors(stats_cb.encrypt_txdma2_sem_free_errors);
    rsp->mutable_spec()->set_encrypt_txdma1_barco_ring_full_errors(stats_cb.encrypt_txdma1_barco_ring_full_errors);
    rsp->mutable_spec()->set_encrypt_rxdma_cb_ring_full_errors(stats_cb.encrypt_rxdma_cb_ring_full_errors);
    rsp->mutable_spec()->set_encrypt_txdma2_barco_req_errors(stats_cb.encrypt_txdma2_barco_req_errors);
    
    rsp->mutable_spec()->set_decrypt_input_desc_errors(stats_cb.decrypt_input_desc_errors);
    rsp->mutable_spec()->set_decrypt_output_desc_errors(stats_cb.decrypt_output_desc_errors);
    rsp->mutable_spec()->set_decrypt_cb_ring_base_errors(stats_cb.decrypt_cb_ring_base_errors);
    rsp->mutable_spec()->set_decrypt_input_page_errors(stats_cb.decrypt_input_page_errors);
    rsp->mutable_spec()->set_decrypt_barco_req_addr_errors(stats_cb.decrypt_barco_req_addr_errors);
    rsp->mutable_spec()->set_decrypt_barco_cb_addr_errors(stats_cb.decrypt_barco_cb_addr_errors);
    rsp->mutable_spec()->set_decrypt_stage4_inpage_errors(stats_cb.decrypt_stage4_inpage_errors);
    rsp->mutable_spec()->set_decrypt_output_page_errors(stats_cb.decrypt_output_page_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_enter_counters(stats_cb.decrypt_txdma1_enter_counters);
    rsp->mutable_spec()->set_decrypt_txdma2_enter_counters(stats_cb.decrypt_txdma2_enter_counters);
    rsp->mutable_spec()->set_decrypt_desc_exhaust_errors(stats_cb.decrypt_desc_exhaust_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_drop_counters(stats_cb.decrypt_txdma1_drop_counters);
    rsp->mutable_spec()->set_decrypt_txdma1_dummy_errors(stats_cb.decrypt_txdma1_dummy_errors);
    rsp->mutable_spec()->set_decrypt_rxdma_dummy_desc_errors(stats_cb.decrypt_rxdma_dummy_desc_errors);
    rsp->mutable_spec()->set_decrypt_load_ipsec_int_errors(stats_cb.decrypt_load_ipsec_int_errors);
    rsp->mutable_spec()->set_decrypt_txdma2_dummy_free(stats_cb.decrypt_txdma2_dummy_free);
    rsp->mutable_spec()->set_decrypt_rxdma_enter_counters(stats_cb.decrypt_rxdma_enter_counters);
    rsp->mutable_spec()->set_decrypt_txdma2_barco_bad_indesc_errors(stats_cb.decrypt_txdma2_barco_bad_indesc_errors);
    rsp->mutable_spec()->set_decrypt_txdma2_barco_bad_outdesc_errors(stats_cb.decrypt_txdma2_barco_bad_outdesc_errors);
    rsp->mutable_spec()->set_decrypt_txdma2_bad_indesc_free_errors(stats_cb.decrypt_txdma2_bad_indesc_free_errors);
    rsp->mutable_spec()->set_decrypt_txdma2_bad_outdesc_free_errors(stats_cb.decrypt_txdma2_bad_outdesc_free_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_bad_indesc_free_errors(stats_cb.decrypt_txdma1_bad_indesc_free_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_bad_outdesc_free_errors(stats_cb.decrypt_txdma1_bad_outdesc_free_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_sem_free_errors(stats_cb.decrypt_txdma1_sem_free_errors);
    rsp->mutable_spec()->set_decrypt_txdma2_sem_free_errors(stats_cb.decrypt_txdma2_sem_free_errors);
    rsp->mutable_spec()->set_decrypt_rxdma_cb_ring_full_errors(stats_cb.decrypt_rxdma_cb_ring_full_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_barco_ring_full_errors(stats_cb.decrypt_txdma1_barco_ring_full_errors);
    rsp->mutable_spec()->set_decrypt_txdma1_barco_full_errors(stats_cb.decrypt_txdma1_barco_full_errors);
    rsp->mutable_spec()->set_decrypt_txdma2_invalid_barco_req_errors(stats_cb.decrypt_txdma2_invalid_barco_req_errors);


    rsp->mutable_spec()->set_enc_rnmdpr_sem_pindex(stats_cb.enc_rnmdpr_pi_counters);
    rsp->mutable_spec()->set_enc_rnmdpr_sem_cindex(stats_cb.enc_rnmdpr_ci_counters);
    rsp->mutable_spec()->set_dec_rnmdpr_sem_pindex(stats_cb.dec_rnmdpr_pi_counters);
    rsp->mutable_spec()->set_dec_rnmdpr_sem_cindex(stats_cb.dec_rnmdpr_ci_counters);
    rsp->mutable_spec()->set_gcm0_barco_full_errors(stats_cb.gcm0_full_counters);
    rsp->mutable_spec()->set_gcm1_barco_full_errors(stats_cb.gcm1_full_counters);
    rsp->mutable_spec()->set_enc_global_barco_pi(stats_cb.enc_global_barco_pi);
    rsp->mutable_spec()->set_enc_global_barco_ci(stats_cb.enc_global_barco_ci);
    rsp->mutable_spec()->set_dec_global_barco_pi(stats_cb.dec_global_barco_pi);
    rsp->mutable_spec()->set_dec_global_barco_ci(stats_cb.dec_global_barco_ci);
    return ret;
}

}    // namespace hal
