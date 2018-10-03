/*
 * {C} Copyright 2018 Pensando Systems Inc.
 * All rights reserved.
 *
 */
#include <netdevice.h>
#include <net.h>
#include <kernel.h>

#include "osal.h"
#include "sonic_dev.h"
#include "sonic_lif.h"
#include "pnso_api.h"

#include "sonic_api_int.h"

#include "pnso_mpool.h"
#include "pnso_pbuf.h"
#include "pnso_chain.h"
#include "pnso_crypto.h"
#include "pnso_cpdc.h"
#include "pnso_crypto_cmn.h"
#include "pnso_seq.h"
#include "pnso_utils.h"
#include "sonic_api_int.h"

static enum crypto_algo_cmd_hi   crypto_algo_cmd_hi_tbl[PNSO_CRYPTO_TYPE_MAX] = {
	[PNSO_CRYPTO_TYPE_XTS] = CRYPTO_ALGO_CMD_HI_AES_XTS,
	[PNSO_CRYPTO_TYPE_GCM] = CRYPTO_ALGO_CMD_HI_AES_GCM,
};

static enum crypto_algo_cmd_lo   crypto_algo_cmd_lo_tbl[PNSO_CRYPTO_TYPE_MAX] = {
	[PNSO_CRYPTO_TYPE_XTS] = CRYPTO_ALGO_CMD_LO_AES_XTS,
	[PNSO_CRYPTO_TYPE_GCM] = CRYPTO_ALGO_CMD_LO_AES_GCM,
};

static pnso_error_t
crypto_validate_input(struct service_info *svc_info,
		      const struct service_params *svc_params)
{
	struct pnso_crypto_desc	*pnso_crypto_desc;

	if (!svc_info || !svc_params) {
		OSAL_LOG_ERROR("invalid svc_info 0x%llx or svc_params 0x%llx!",
				(uint64_t)svc_info, (uint64_t)svc_params);
		return EINVAL;
	}

	if (!svc_params->sp_src_blist || !svc_params->sp_dst_blist) {
		OSAL_LOG_ERROR("invalid sp_src_blist 0x%llx or "
				"sp_dst_blist 0x%llx!",
				(uint64_t)svc_params->sp_src_blist,
				(uint64_t)svc_params->sp_dst_blist);
		return EINVAL;
	}

	pnso_crypto_desc = svc_params->u.sp_crypto_desc;
	if (!pnso_crypto_desc) {
		OSAL_LOG_ERROR("invalid desc 0x%llx specified!",
				(uint64_t)pnso_crypto_desc);
		return EINVAL;
	}

	if ((pnso_crypto_desc->algo_type == PNSO_CRYPTO_TYPE_NONE) ||
	    (pnso_crypto_desc->algo_type >= PNSO_CRYPTO_TYPE_MAX)) {
		OSAL_LOG_ERROR("invalid algo_type %u specified!",
				pnso_crypto_desc->algo_type);
		return EINVAL;
	}

	if (!pnso_crypto_desc->iv_addr) {
		OSAL_LOG_ERROR("invalid iv_addr 0x%llx specified!",
				pnso_crypto_desc->iv_addr);
		return EINVAL;
	}

	return PNSO_OK;
}

static inline void
crypto_desc_fill(struct service_info *svc_info,
		 struct pnso_crypto_desc *pnso_crypto_desc)
{
	struct crypto_desc *crypto_desc = svc_info->si_desc;
	struct crypto_status_desc *status_desc = svc_info->si_status_desc;

	/*
	 * Intermediate status is never directy "polled" so no need to clear it.
	 */
	memset(crypto_desc, 0, sizeof(*crypto_desc));
        memset(status_desc, 0, sizeof(*status_desc));
	if (svc_info->si_istatus_desc)
		crypto_desc->cd_status_addr = 
		    mpool_get_object_phy_addr(MPOOL_TYPE_RMEM_INTERM_CRYPTO_STATUS,
					      svc_info->si_istatus_desc);
        else
		crypto_desc->cd_status_addr = osal_virt_to_phy(status_desc);

	crypto_desc->cd_in_aol = sonic_virt_to_phy(svc_info->si_src_aol);
	crypto_desc->cd_out_aol = sonic_virt_to_phy(svc_info->si_dst_aol);

	OSAL_ASSERT(pnso_crypto_desc->algo_type < PNSO_CRYPTO_TYPE_MAX);
	crypto_desc->cd_db_data = CRYPTO_ENCRYPT_CPL_DATA;
	if (svc_info->si_type == PNSO_SVC_TYPE_DECRYPT) {
		crypto_desc->cd_cmd.cc_is_decrypt = true;
		crypto_desc->cd_db_data = CRYPTO_DECRYPT_CPL_DATA;
	}
	crypto_desc->cd_cmd.cc_token_3 =
			    crypto_algo_cmd_lo_tbl[pnso_crypto_desc->algo_type];
	crypto_desc->cd_cmd.cc_token_4 =
			    crypto_algo_cmd_hi_tbl[pnso_crypto_desc->algo_type];
	crypto_desc->cd_key_desc_idx = 
		     sonic_get_crypto_key_idx(pnso_crypto_desc->key_desc_idx);

	crypto_desc->cd_iv_addr = sonic_hostpa_to_devpa(pnso_crypto_desc->iv_addr);
	crypto_desc->cd_db_addr = crypto_desc->cd_status_addr +
				  sizeof(status_desc->csd_err);

	CRYPTO_PPRINT_DESC(crypto_desc);
}

static inline pnso_error_t
crypto_dst_blist_setup(struct service_info *svc_info,
		       const struct service_params *svc_params)
{
	const struct per_core_resource	*pc_res = svc_info->si_pc_res;

	/*
	 * Produce output to intermediate buffers if there is a chain subordinate.
	 */
	if (chn_service_has_sub_chain(svc_info)) {
		svc_info->si_iblist =
			pc_res_interm_buf_list_get(pc_res, MPOOL_TYPE_INTERM_BUF_LIST,
						   MPOOL_TYPE_RMEM_INTERM_BUF);
		if (!svc_info->si_iblist) {
			OSAL_LOG_ERROR("failed to obtain intermediate buffers");
			return ENOMEM;
		}
		svc_info->si_dst_blist = &svc_info->si_iblist->blist;

		svc_info->si_istatus_desc =
		    pc_res_mpool_object_get(pc_res,
					    MPOOL_TYPE_RMEM_INTERM_CRYPTO_STATUS);
		if (!svc_info->si_istatus_desc) {
			OSAL_LOG_ERROR("failed to obtain intermediate status_desc");
			return ENOMEM;
		}

		svc_info->si_sgl_pdma =
		    pc_res_sgl_pdma_packed_get(pc_res, svc_params->sp_dst_blist);
		if (!svc_info->si_sgl_pdma) {
			OSAL_LOG_ERROR("failed to obtain chain SGL for PDMA");
			return ENOMEM;
		}
	} else {
		svc_info->si_dst_blist = svc_params->sp_dst_blist;
	}
	return PNSO_OK;
}

static inline pnso_error_t
crypto_src_dst_aol_fill(struct service_info *svc_info,
			const struct service_params *svc_params)
{
	const struct per_core_resource	*pc_res = svc_info->si_pc_res;

	/*
	 * 1) First in chain or the only service: enter the src/dst buf into AOLs
	 *    in packed format (i.e. more efficient usage of the AOLs).
	 * 2) Part of a chain (and is not first): use a sparse vector of AOLs
	 *    for the src/dst buf info to facilitate sequencer padding.
	 */
	if (!chn_service_is_in_chain(svc_info) ||
	     chn_service_is_first(svc_info)) {
		svc_info->si_src_aol = 
			crypto_aol_packed_get(pc_res, svc_params->sp_src_blist,
					      &svc_info->si_src_len);
		svc_info->si_dst_aol = 
			crypto_aol_packed_get(pc_res, svc_info->si_dst_blist,
					      &svc_info->si_dst_len);
	} else {
		svc_info->si_src_aol = 
			crypto_aol_sparse_get(pc_res, svc_info->si_block_size,
				svc_params->sp_src_blist, &svc_info->si_src_len);
		svc_info->si_dst_aol = 
			crypto_aol_sparse_get(pc_res, svc_info->si_block_size,
				svc_info->si_dst_blist, &svc_info->si_dst_len);
	}

	if (!svc_info->si_src_aol || !svc_info->si_dst_aol) {
		return ENOMEM;
	}

	OSAL_LOG_INFO("src_total_len %u dst_total_len %u",
		      svc_info->si_src_len, svc_info->si_dst_len);
	if ((svc_info->si_src_len == 0) ||
	    (svc_info->si_dst_len < svc_info->si_src_len)) {
		OSAL_LOG_ERROR("length error: src_len %u dst_len %u",
				svc_info->si_src_len, svc_info->si_dst_len);
		return EINVAL;
	}

	/*
	 * Fix up intermediate buffer list length.
	 */
	if (svc_info->si_iblist) {
                svc_info->si_iblist->blist.buffers[0].len = svc_info->si_src_len;
                svc_info->si_dst_len = svc_info->si_src_len;
	}

	return PNSO_OK;
}

static pnso_error_t
crypto_setup(struct service_info *svc_info,
	     const struct service_params *svc_params)
{
	const struct per_core_resource	*pc_res;
	pnso_error_t			err;

	err = crypto_validate_input(svc_info, svc_params);
	if (err)
		return err;

	pc_res = svc_info->si_pc_res;
	svc_info->si_desc = pc_res_mpool_object_get(pc_res,
						    MPOOL_TYPE_CRYPTO_DESC);
	if (!svc_info->si_desc)
		return ENOMEM;

	svc_info->si_status_desc = pc_res_mpool_object_get(pc_res,
					  MPOOL_TYPE_CRYPTO_STATUS_DESC);
	if (!svc_info->si_status_desc)
		return ENOMEM;

	err = crypto_dst_blist_setup(svc_info, svc_params);
	if (err)
		return err;

	err = crypto_src_dst_aol_fill(svc_info, svc_params);
	if (err)
		return err;

	svc_info->si_desc_flags = 0;
	crypto_desc_fill(svc_info, svc_params->u.sp_crypto_desc);

	if (!chn_service_is_in_chain(svc_info) ||
	     chn_service_is_first(svc_info)) {
		svc_info->si_seq_info.sqi_desc = seq_setup_desc(svc_info,
				svc_info->si_desc, sizeof(struct crypto_desc));
		if (!svc_info->si_seq_info.sqi_desc) {
			OSAL_LOG_ERROR("failed to setup sequencer desc");
			return EINVAL;
		}
	}

	return PNSO_OK;
}

static pnso_error_t
crypto_encrypt_setup(struct service_info *svc_info,
		     const struct service_params *svc_params)
{
	svc_info->si_type = PNSO_SVC_TYPE_ENCRYPT;
	return crypto_setup(svc_info, svc_params);
}

static pnso_error_t
crypto_decrypt_setup(struct service_info *svc_info,
		     const struct service_params *svc_params)
{
	svc_info->si_type = PNSO_SVC_TYPE_DECRYPT;
	return crypto_setup(svc_info, svc_params);
}

static pnso_error_t
crypto_chain(struct chain_entry *centry)
{
	struct service_info		*svc_info = &centry->ce_svc_info;
	struct crypto_chain_params	*crypto_chain = &svc_info->si_crypto_chain;
	struct interm_buf_list		*iblist;
	struct service_info		*svc_next;
	pnso_error_t			err;

	if (chn_service_has_sub_chain(svc_info)) {
		iblist = svc_info->si_iblist;
		crypto_chain->ccp_crypto_buf_addr = iblist->blist.buffers[0].buf;
		crypto_chain->ccp_data_len = iblist->blist.buffers[0].len;
		crypto_chain->ccp_sgl_pdma_dst_addr =
				osal_virt_to_phy(svc_info->si_sgl_pdma);
		crypto_chain->ccp_cmd.ccpc_sgl_pdma_en = true;
		crypto_chain->ccp_cmd.ccpc_sgl_pdma_len_from_desc = true;

		crypto_chain->ccp_status_addr_0 =
		  mpool_get_object_phy_addr(MPOOL_TYPE_RMEM_INTERM_CRYPTO_STATUS,
					    svc_info->si_istatus_desc);
		crypto_chain->ccp_status_addr_1 =
			osal_virt_to_phy(svc_info->si_status_desc);
		crypto_chain->ccp_status_len = sizeof(struct crypto_status_desc);
		crypto_chain->ccp_cmd.ccpc_status_dma_en = true;
		crypto_chain->ccp_cmd.ccpc_stop_chain_on_error = true;

		OSAL_ASSERT(centry->ce_next);
		svc_next = &centry->ce_next->ce_svc_info;
		err = svc_next->si_ops.sub_chain_from_crypto(svc_next,
							     crypto_chain);
		if (!err) {
			svc_info->si_seq_info.sqi_desc = 
				seq_setup_crypto_chain(svc_info, svc_info->si_desc);
			if (!svc_info->si_seq_info.sqi_desc) {
				OSAL_LOG_ERROR("failed seq_setup_crypto_chain");
				return EINVAL;
			}
		}
	}
	return PNSO_OK;
}

static pnso_error_t
crypto_sub_chain_from_cpdc(struct service_info *svc_info,
			   struct cpdc_chain_params *cpdc_chain)
{
	cpdc_chain->ccp_aol_src_vec_addr = osal_virt_to_phy(svc_info->si_src_aol);
	cpdc_chain->ccp_aol_dst_vec_addr = osal_virt_to_phy(svc_info->si_dst_aol);
	cpdc_chain->ccp_cmd.ccpc_aol_pad_en = !!cpdc_chain->ccp_pad_buf_addr;
	cpdc_chain->ccp_cmd.ccpc_next_doorbell_en = true;
	cpdc_chain->ccp_cmd.ccpc_next_db_action_ring_push = true;
	return ring_spec_info_fill(svc_info->si_seq_info.sqi_ring_id,
				   &cpdc_chain->ccp_ring_spec,
				   svc_info->si_desc, 1);
}

static pnso_error_t
crypto_sub_chain_from_crypto(struct service_info *svc_info,
			     struct crypto_chain_params *crypto_chain)
{
	/*
	 * This is supportable when there's a valid use case.
	 * For testing purposes, it is possible to chain encrypt to decrypt.
	 */
	crypto_chain->ccp_cmd.ccpc_next_doorbell_en = true;
	crypto_chain->ccp_cmd.ccpc_next_db_action_ring_push = true;
	return ring_spec_info_fill(svc_info->si_seq_info.sqi_ring_id,
				   &crypto_chain->ccp_ring_spec,
				   svc_info->si_desc, 1);
}

static pnso_error_t
crypto_schedule(const struct service_info *svc_info)
{
	pnso_error_t err = EINVAL;
	const struct sequencer_info *seq_info;
	bool ring_db;

	ring_db = svc_info->si_flags & (CHAIN_SFLAG_LONE_SERVICE |
					CHAIN_SFLAG_FIRST_SERVICE);
	if (ring_db) {
		seq_info = &svc_info->si_seq_info;
		seq_ring_db(svc_info);
		err = PNSO_OK;
	}
	return err;
}

static pnso_error_t
crypto_poll(const struct service_info *svc_info)
{
	volatile struct crypto_status_desc *status_desc;
	uint64_t cpl_data;

	status_desc = svc_info->si_status_desc;
	cpl_data = svc_info->si_type == PNSO_SVC_TYPE_DECRYPT ?
		   CRYPTO_DECRYPT_CPL_DATA : CRYPTO_ENCRYPT_CPL_DATA;
	while (status_desc->csd_cpl_data != cpl_data)
		osal_yield();

	return PNSO_OK;
}

static pnso_error_t
crypto_read_status(const struct service_info *svc_info)
{
	struct crypto_status_desc *status_desc;
	pnso_error_t err;

	status_desc = svc_info->si_status_desc;
	err = status_desc->csd_err;
	if (err) {
		OSAL_LOG_ERROR("hw error reported: %d", err);
	}
	return err;
}

static pnso_error_t
crypto_write_result(struct service_info *svc_info)
{
	struct pnso_service_status *svc_status;
	struct crypto_status_desc *status_desc;

	svc_status = svc_info->si_svc_status;
	status_desc = svc_info->si_status_desc;
	if (status_desc->csd_err) {
		svc_status->err = crypto_desc_status_convert(status_desc->csd_err);
		OSAL_LOG_ERROR("service failed: %d", svc_status->err);
		return EINVAL;
	}

	return PNSO_OK;
}

static void
crypto_teardown(const struct service_info *svc_info)
{
	const struct crypto_chain_params *crypto_chain = &svc_info->si_crypto_chain;
	enum mem_pool_type aol_type;

	aol_type = !chn_service_is_in_chain(svc_info) ||
		   chn_service_is_first(svc_info) ? MPOOL_TYPE_CRYPTO_AOL :
						    MPOOL_TYPE_CRYPTO_AOL_VECTOR;
	crypto_aol_put(svc_info->si_pc_res, aol_type, svc_info->si_dst_aol);
	crypto_aol_put(svc_info->si_pc_res, aol_type, svc_info->si_src_aol);
	pc_res_mpool_object_put(svc_info->si_pc_res, MPOOL_TYPE_CRYPTO_STATUS_DESC,
				svc_info->si_status_desc);
	pc_res_mpool_object_put(svc_info->si_pc_res, MPOOL_TYPE_CRYPTO_DESC,
				svc_info->si_desc);
	pc_res_interm_buf_list_put(svc_info->si_pc_res, svc_info->si_iblist);
	pc_res_mpool_object_put(svc_info->si_pc_res,
				MPOOL_TYPE_RMEM_INTERM_CRYPTO_STATUS,
				svc_info->si_istatus_desc);
	pc_res_sgl_pdma_put(svc_info->si_pc_res,
			    svc_info->si_sgl_pdma);

	if (crypto_chain->ccp_seq_spec.sqs_seq_status_q)
		sonic_put_seq_statusq(crypto_chain->ccp_seq_spec.sqs_seq_status_q);
}

struct service_ops encrypt_ops = {
	.setup = crypto_encrypt_setup,
	.chain = crypto_chain,
	.sub_chain_from_cpdc = crypto_sub_chain_from_cpdc,
	.sub_chain_from_crypto = crypto_sub_chain_from_crypto,
	.schedule = crypto_schedule,
	.poll = crypto_poll,
	.read_status = crypto_read_status,
	.write_result = crypto_write_result,
	.teardown = crypto_teardown,
};

struct service_ops decrypt_ops = {
	.setup = crypto_decrypt_setup,
	.chain = crypto_chain,
	.sub_chain_from_cpdc = crypto_sub_chain_from_cpdc,
	.sub_chain_from_crypto = crypto_sub_chain_from_crypto,
	.schedule = crypto_schedule,
	.poll = crypto_poll,
	.read_status = crypto_read_status,
	.write_result = crypto_write_result,
	.teardown = crypto_teardown,
};
