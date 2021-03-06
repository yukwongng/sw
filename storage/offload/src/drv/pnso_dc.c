/*
 * {C} Copyright 2018 Pensando Systems Inc.
 * All rights reserved.
 *
 */
#include "osal.h"
#include "pnso_api.h"

#include "sonic_api_int.h"

#include "pnso_mpool.h"
#include "pnso_pbuf.h"
#include "pnso_svc.h"
#include "pnso_chain.h"
#include "pnso_cpdc.h"
#include "pnso_cpdc_cmn.h"
#include "pnso_seq.h"
#include "pnso_utils.h"

static inline bool
is_chksum_present_in_header_format(struct service_info *svc_info)
{
	struct cp_header_format *hdr_fmt;

	hdr_fmt = lookup_hdr_format(svc_info->hdr_fmt_idx, false);
	OSAL_LOG_DEBUG("checksum is present in header format for decompression");

	return hdr_fmt ? (hdr_fmt->chksum_len > 0) : false;
}

static int
fill_dc_desc(struct service_info *svc_info, struct cpdc_desc *desc)
{
	uint32_t src_buf_len = svc_info->si_src_blist.len;
	uint32_t dst_buf_len = svc_info->si_dst_blist.len;
	uint64_t aligned_addr;
	struct cpdc_sgl *sgl;
	uint32_t scratch_buf_len = 0;
	pnso_error_t err;

	memset(desc, 0, sizeof(*desc));

	desc->cd_src = (uint64_t) sonic_virt_to_phy(svc_info->si_src_sgl.sgl);
	desc->cd_dst = (uint64_t) sonic_virt_to_phy(svc_info->si_dst_sgl.sgl);

	desc->u.cd_bits.cc_enabled = 1;
	desc->u.cd_bits.cc_integrity_type = CPDC_INTEGRITY_TYPE_DFLT;
	desc->u.cd_bits.cc_integrity_src = CPDC_INTEGRITY_SRC_DFLT;
	desc->u.cd_bits.cc_chksum_adler = CPDC_CHKSUM_TYPE_DFLT;

	if (svc_info->si_desc_flags & PNSO_DC_DFLAG_HEADER_PRESENT) {
		desc->u.cd_bits.cc_header_present = 1;
		if (is_chksum_present_in_header_format(svc_info))
			desc->u.cd_bits.cc_chksum_verify_enabled = 1;
	}

	desc->u.cd_bits.cc_src_is_list = 1;
	desc->u.cd_bits.cc_dst_is_list = 1;

	sgl = svc_info->si_dst_sgl.sgl;
	if (sgl->cs_len_0 == CPDC_SCRATCH_BUFFER_LEN)
		scratch_buf_len = CPDC_SCRATCH_BUFFER_LEN;

	desc->cd_datain_len = cpdc_desc_data_len_set_eval(svc_info->si_type,
							  src_buf_len);
	desc->cd_threshold_len =
		cpdc_desc_data_len_set_eval(svc_info->si_type,
				dst_buf_len + scratch_buf_len);

	err = svc_status_desc_addr_get(&svc_info->si_status_desc, 0,
			&aligned_addr, CPDC_STATUS_INTEG_CLEAR_SZ);
	desc->cd_status_addr = aligned_addr;
	if (err)
		goto out;

	if (chn_service_has_interm_status(svc_info)) {
		err = svc_status_desc_addr_get(&svc_info->si_istatus_desc, 0,
			&aligned_addr, CPDC_STATUS_MIN_CLEAR_SZ);
		desc->cd_status_addr = aligned_addr;
	}

	desc->cd_status_data = CPDC_DC_STATUS_DATA;

	CPDC_PPRINT_DESC(desc);
out:
	return err;
}

static pnso_error_t
decompress_setup(struct service_info *svc_info,
		const struct service_params *svc_params)
{
	pnso_error_t err;
	struct cpdc_desc *dc_desc;

	OSAL_LOG_DEBUG("enter ...");

	dc_desc = cpdc_get_desc(svc_info, false);
	if (!dc_desc) {
		err = EAGAIN;
		OSAL_LOG_DEBUG("cannot obtain dc desc from pool! err: %d", err);
		goto out;
	}
	svc_info->si_desc = dc_desc;

	err = cpdc_setup_status_desc(svc_info, false);
	if (err) {
		OSAL_LOG_DEBUG("cannot obtain dc status desc from pool! err: %d",
				err);
		goto out;
	}

	err = cpdc_setup_rmem_status_desc(svc_info, false);
	if (err) {
		OSAL_LOG_DEBUG("intermediate status not available, "
				"using host status");
	}

	err = cpdc_setup_rmem_dst_blist(svc_info, svc_params);
	if (err) {
		OSAL_LOG_ERROR("failed to setup dc rmem dst buffer list! err: %d",
				err);
		goto out;
	}

	err = cpdc_update_service_info_sgls(svc_info);
	if (err) {
		OSAL_LOG_ERROR("cannot obtain dc src/dst sgl from pool! err: %d",
				err);
		goto out;
	}

	err = fill_dc_desc(svc_info, dc_desc);
	if (err) {
		OSAL_LOG_ERROR("cannot fill_dc_desc! err: %d", err);
		goto out;
	}

	err = svc_seq_desc_setup(svc_info, dc_desc, sizeof(*dc_desc), 0);
	if (err) {
		OSAL_LOG_DEBUG("failed to setup sequencer desc! err: %d", err);
		goto out;
	}

	PAS_INC_NUM_DC_REQUESTS(svc_info->si_pcr);
	PAS_INC_NUM_DC_BYTES_IN(svc_info->si_pcr, svc_info->si_src_blist.len);

	err = PNSO_OK;
	OSAL_LOG_DEBUG("exit! service initialized!");
	return err;

out:
	OSAL_LOG_SPECIAL_ERROR("exit! err: %d", err);
	return err;
}

static pnso_error_t
decompress_chain(struct chain_entry *centry)
{
	struct service_info		*svc_info = &centry->ce_svc_info;
	struct cpdc_chain_params	*cpdc_chain = &svc_info->si_cpdc_chain;
	struct interm_buf_list		*iblist;
	struct service_info		*svc_next;
	pnso_error_t			err = PNSO_OK;

	if (chn_service_has_sub_chain(svc_info)) {

		/*
		 * Length must always be indicated to P4+ chainer (for rate
		 * limiting) even when PDMA is not applicable.
		 */
		cpdc_chain->ccp_data_len = svc_info->si_dst_blist.len;
		if (chn_service_has_interm_blist(svc_info)) {
			iblist = &svc_info->si_iblist;
			cpdc_chain->ccp_comp_buf_addr =
				iblist->blist.buffers[0].buf;
			if (chn_service_has_sgl_pdma(svc_info)) {
				cpdc_chain->ccp_aol_dst_vec_addr =
					sonic_virt_to_phy(svc_info->si_sgl_pdma);
				cpdc_chain->ccp_cmd.ccpc_sgl_pdma_en = true;
				SGL_PDMA_PPRINT(cpdc_chain->ccp_aol_dst_vec_addr);
			}
		}

		err = cpdc_setup_status_chain_dma(svc_info, cpdc_chain);
		if (err)
			goto out;

		cpdc_chain->ccp_cmd.status_len_no_hdr = true;
		cpdc_chain->ccp_cmd.ccpc_stop_chain_on_error = true;

		/*
		 * Set ccp_pad_boundary_shift whether or not 
		 * padding is applicable.
		 */
		cpdc_chain->ccp_pad_boundary_shift =
			(uint8_t)ilog2(svc_info->si_block_size);
		svc_next = chn_service_next_svc_get(svc_info);
		OSAL_ASSERT(svc_next);

		err = svc_next->si_ops.sub_chain_from_cpdc(svc_next,
				cpdc_chain);
		if (!err)
			err = seq_setup_cpdc_chain(svc_info,
					svc_info->si_desc);
	}
out:
	if (err)
		OSAL_LOG_SPECIAL_ERROR("failed seq_setup_cpdc_chain: err %d",
				err);

	return err;
}

static pnso_error_t
decompress_sub_chain_from_cpdc(struct service_info *svc_info,
			       struct cpdc_chain_params *cpdc_chain)
{
	struct service_info	*svc_prev = chn_service_prev_svc_get(svc_info);

	cpdc_chain->ccp_sgl_vec_addr =
		sonic_virt_to_phy(svc_info->si_src_sgl.sgl);
	cpdc_chain->ccp_cmd.ccpc_sgl_update_en =
		chn_service_type_is_cp(svc_prev);
	cpdc_chain->ccp_cmd.desc_dlen_update_en =
		cpdc_chain->ccp_cmd.ccpc_sgl_update_en;

        svc_rate_limit_control_eval(svc_info, &cpdc_chain->ccp_rl_control);
	cpdc_chain->ccp_cmd.ccpc_next_doorbell_en = true;
	cpdc_chain->ccp_cmd.ccpc_next_db_action_ring_push = true;
	cpdc_chain->ccp_num_alt_descs = 1;

	return ring_spec_info_fill(svc_info->si_seq_info.sqi_ring,
				   &cpdc_chain->ccp_ring_spec,
				   svc_info->si_desc, 1);
}

static pnso_error_t
decompress_sub_chain_from_crypto(struct service_info *svc_info,
				 struct crypto_chain_params *crypto_chain)
{
	struct cpdc_desc *dc_desc;

	if (crypto_chain->ccp_cmd.ccpc_comp_sgl_src_en) {
		dc_desc = svc_info->si_desc;
		dc_desc->cd_src = crypto_chain->ccp_comp_sgl_src_addr;
		dc_desc->u.cd_bits.cc_src_is_list = true;
		CPDC_PPRINT_DESC(dc_desc);
	}
        svc_rate_limit_control_eval(svc_info, &crypto_chain->ccp_rl_control);
	crypto_chain->ccp_cmd.ccpc_next_doorbell_en = true;
	crypto_chain->ccp_cmd.ccpc_next_db_action_ring_push = true;
	return ring_spec_info_fill(svc_info->si_seq_info.sqi_ring,
				   &crypto_chain->ccp_ring_spec,
				   svc_info->si_desc, 1);
}

static pnso_error_t
decompress_enable_interrupt(struct service_info *svc_info, void *poll_ctx)
{
	return cpdc_setup_interrupt_params(svc_info, poll_ctx);
}

static pnso_error_t
decompress_ring_db(struct service_info *svc_info)
{
	pnso_error_t err = EINVAL;
	const struct sequencer_info *seq_info;
	bool ring_db;

	OSAL_LOG_DEBUG("enter ...");

	OSAL_ASSERT(svc_info);

	ring_db = (svc_info->si_flags & CHAIN_SFLAG_LONE_SERVICE) ||
		(svc_info->si_flags & CHAIN_SFLAG_FIRST_SERVICE);
	if (ring_db) {
		OSAL_LOG_INFO("ring door bell <===");

		seq_info = &svc_info->si_seq_info;
		err = seq_ring_db(svc_info);
	}

	OSAL_LOG_DEBUG("exit!");
	return err;
}

static pnso_error_t
decompress_poll(struct service_info *svc_info)
{
	return cpdc_poll(svc_info, NULL);
}

static pnso_error_t
decompress_write_result(struct service_info *svc_info)
{
	pnso_error_t err = EINVAL;
	struct pnso_service_status *svc_status;
	struct cpdc_desc *dc_desc;
	struct cpdc_status_desc *status_desc;

	struct cpdc_sgl *sgl;
	uint32_t scratch_buf_len = 0;

	OSAL_LOG_DEBUG("enter ...");

	OSAL_ASSERT(svc_info);

	svc_status = svc_info->si_svc_status;
	if (svc_status->svc_type != svc_info->si_type) {
		svc_status->err = err;
		OSAL_LOG_ERROR("service type mismatch! svc_type: %d si_type: %d err: %d",
			svc_status->svc_type,  svc_info->si_type, err);
		goto out;
	}

	dc_desc = (struct cpdc_desc *) svc_info->si_desc;
	status_desc = (struct cpdc_status_desc *) svc_info->si_status_desc.desc;

	if (!status_desc->csd_valid) {
		svc_status->err = err;
		OSAL_LOG_ERROR("valid bit not set! err: %d", err);
		goto out;
	}

	if (status_desc->csd_partial_data != dc_desc->cd_status_data)
		OSAL_LOG_DEBUG("partial data mismatch, expected %u received: %u err: %d",
				dc_desc->cd_status_data,
				status_desc->csd_partial_data, err);

	if (status_desc->csd_err) {
		svc_status->err = cpdc_convert_desc_error(status_desc->csd_err);
		OSAL_LOG_DEBUG("hw error reported! csd_err: %d err: %d",
				status_desc->csd_err, svc_status->err);
		goto pass_err;
	}

	sgl = svc_info->si_dst_sgl.sgl;
	if (sgl->cs_len_0 == CPDC_SCRATCH_BUFFER_LEN)
		scratch_buf_len = CPDC_SCRATCH_BUFFER_LEN;

	svc_status->u.dst.data_len = scratch_buf_len ?
		cpdc_desc_data_len_get_eval(svc_info->si_type,
				status_desc->csd_output_data_len) -
				scratch_buf_len :
		cpdc_desc_data_len_get_eval(svc_info->si_type,
				status_desc->csd_output_data_len);
	OSAL_LOG_DEBUG("scratch_buf_len: %u csd_output_data_len: %u svc_status->u.dst.data_len: %u",
			scratch_buf_len, status_desc->csd_output_data_len,
			svc_status->u.dst.data_len);

	chn_service_deps_data_len_set(svc_info, svc_status->u.dst.data_len);

	PAS_INC_NUM_DC_BYTES_OUT(svc_info->si_pcr,
			svc_status->u.dst.data_len);

	err = PNSO_OK;
	OSAL_LOG_DEBUG("exit! status/result update success!");
	return err;
pass_err:
	OSAL_LOG_DEBUG("exit!");
	return err;
out:
	OSAL_LOG_ERROR("exit! err: %d", err);
	return err;
}

static void
decompress_teardown(struct service_info *svc_info)
{
	struct cpdc_desc *dc_desc;

	OSAL_LOG_DEBUG("enter ...");

	OSAL_ASSERT(svc_info);

	/*
	 * Trace the desc/SGLs once more to verify any padding applied
	 * by sequencer.
	 */
	CPDC_PPRINT_DESC(svc_info->si_desc);

	seq_cleanup_cpdc_chain(svc_info);
	seq_cleanup_desc(svc_info);

	pc_res_sgl_put(svc_info->si_pcr, &svc_info->si_dst_sgl);
	pc_res_sgl_put(svc_info->si_pcr, &svc_info->si_src_sgl);

	cpdc_teardown_rmem_dst_blist(svc_info);
	cpdc_teardown_rmem_status_desc(svc_info);

	cpdc_teardown_status_desc(svc_info);

	dc_desc = (struct cpdc_desc *) svc_info->si_desc;
	cpdc_put_desc(svc_info, false, dc_desc);

	OSAL_LOG_DEBUG("exit!");
}

struct service_ops dc_ops = {
	.setup = decompress_setup,
	.chain = decompress_chain,
	.sub_chain_from_cpdc = decompress_sub_chain_from_cpdc,
	.sub_chain_from_crypto = decompress_sub_chain_from_crypto,
	.enable_interrupt = decompress_enable_interrupt,
	.ring_db = decompress_ring_db,
	.poll = decompress_poll,
	.write_result = decompress_write_result,
	.teardown = decompress_teardown
};
