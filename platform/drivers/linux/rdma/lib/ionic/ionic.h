#ifndef IONIC_H
#define IONIC_H

#include <inttypes.h>
#include <stdbool.h>
#include <stddef.h>
#include <endian.h>
#include <pthread.h>

#include <infiniband/driver.h>
#include <util/udma_barrier.h>
#include <ccan/list.h>

#include "ionic-abi.h"

#include "ionic_dbg.h"
#include "ionic_memory.h"
#include "ionic_queue.h"
#include "table.h"

#define IONIC_MIN_RDMA_VERSION	0
#define IONIC_MAX_RDMA_VERSION	1

struct ionic_ctx {
	struct verbs_context	vctx;

	bool			fallback;
	uint32_t		pg_shift;

	int			version;
	uint8_t			compat;
	uint8_t			opcodes;

	uint8_t			sq_qtype;
	uint8_t			rq_qtype;
	uint8_t			cq_qtype;

	uint64_t		*dbpage;

	pthread_mutex_t		mut;
	struct tbl_root		qp_tbl;
};

struct ionic_cq {
	struct ibv_cq		ibcq;

	uint32_t		cqid;

	pthread_spinlock_t	lock;
	struct list_head	poll_sq;
	struct list_head	flush_sq;
	struct list_head	flush_rq;
	struct ionic_queue	q;
	uint16_t		arm_any_prod;
	uint16_t		arm_sol_prod;
};

struct ionic_sq_meta {
	uint64_t		wrid;
	uint32_t		len;
	uint16_t		seq;
	uint8_t			op;
	uint8_t			status;
	bool			remote;
	bool			signal;
};

struct ionic_rq_meta {
	uint64_t		wrid;
	uint32_t		len; /* XXX byte_len must come from cqe */
};

struct ionic_qp {
	union {
		struct verbs_qp		vqp;
		struct verbs_srq	vsrq;
	};

	uint32_t		qpid;
	bool			has_sq;
	bool			has_rq;
	bool			is_srq;

	bool			sig_all;

	struct list_node	cq_poll_sq;
	struct list_node	cq_flush_sq;
	struct list_node	cq_flush_rq;

	pthread_spinlock_t	sq_lock;
	bool			sq_flush;
	struct ionic_queue	sq;
	struct ionic_sq_meta	*sq_meta;
	uint16_t		*sq_msn_idx;
	uint16_t		sq_msn_prod;
	uint16_t		sq_msn_cons;
	uint16_t		sq_npg_cons;

	void			*sq_hbm_ptr;
	uint16_t		sq_hbm_prod;

	pthread_spinlock_t	rq_lock;
	bool			rq_flush;
	struct ionic_queue	rq;
	struct ionic_rq_meta	*rq_meta;
};

struct ionic_ah {
	struct ibv_ah		ibah;
	uint32_t		ahid;
};

struct ionic_dev {
	struct verbs_device	vdev;
	uint8_t			abi_version;
};

static inline struct ionic_dev *to_ionic_dev(struct ibv_device *ibdev)
{
	return container_of(ibdev, struct ionic_dev, vdev.device);
}

static inline struct ionic_ctx *to_ionic_ctx(struct ibv_context *ibctx)
{
	return container_of(ibctx, struct ionic_ctx, vctx.context);
}

static inline struct ionic_cq *to_ionic_cq(struct ibv_cq *ibcq)
{
	return container_of(ibcq, struct ionic_cq, ibcq);
}

static inline struct ionic_qp *to_ionic_qp(struct ibv_qp *ibqp)
{
	return container_of(ibqp, struct ionic_qp, vqp.qp);
}

static inline struct ionic_qp *to_ionic_srq(struct ibv_srq *ibsrq)
{
	return container_of(ibsrq, struct ionic_qp, vsrq.srq);
}

static inline struct ionic_ah *to_ionic_ah(struct ibv_ah *ibah)
{
        return container_of(ibah, struct ionic_ah, ibah);
}

static inline uint16_t ionic_get_sqe_size(uint16_t max_sge,
					  uint16_t max_inline)
{
	if (max_sge < 2)
		max_sge = 2;

	max_sge *= 16;

	if (max_sge < max_inline)
		max_sge = max_inline;

	return 32 + max_sge;
}

static inline uint32_t ionic_get_rqe_size(uint16_t max_sge)
{
	if (max_sge < 2)
		max_sge = 2;

	max_sge *= 16;

	return 32 + max_sge;
}

static inline uint8_t ibv_to_ionic_wr_opcd(uint8_t ibv_opcd)
{
	uint8_t bnxt_opcd;

	switch (ibv_opcd) {
	case IBV_WR_SEND:
		bnxt_opcd = IONIC_WR_OPCD_SEND;
		break;
	case IBV_WR_SEND_WITH_IMM:
		bnxt_opcd = IONIC_WR_OPCD_SEND_IMM;
		break;
	case IBV_WR_RDMA_WRITE:
		bnxt_opcd = IONIC_WR_OPCD_RDMA_WRITE;
		break;
	case IBV_WR_RDMA_WRITE_WITH_IMM:
		bnxt_opcd = IONIC_WR_OPCD_RDMA_WRITE_IMM;
		break;
	case IBV_WR_RDMA_READ:
		bnxt_opcd = IONIC_WR_OPCD_RDMA_READ;
		break;
	case IBV_WR_ATOMIC_CMP_AND_SWP:
		bnxt_opcd = IONIC_WR_OPCD_ATOMIC_CS;
		break;
	case IBV_WR_ATOMIC_FETCH_AND_ADD:
		bnxt_opcd = IONIC_WR_OPCD_ATOMIC_FA;
		break;
		/* TODO: Add other opcodes */
	default:
		bnxt_opcd = IONIC_WR_OPCD_INVAL;
		break;
	};

	return bnxt_opcd;
}

static inline void ionic_set_ibv_send_flags(int flags, struct sqwqe_t *wqe)
{
	if (flags & IBV_SEND_FENCE) {
		wqe->base.fence = 1;
	}
	if (flags & IBV_SEND_SOLICITED) {
		wqe->base.solicited_event = 1;
	}
	if (flags & IBV_SEND_INLINE) {
		wqe->base.inline_data_vld = 1;
	}
	if (flags & IBV_SEND_SIGNALED) {
		wqe->base.complete_notify = 1;
	}
}

static inline uint8_t ionic_ibv_wr_to_wc_opcd(uint8_t wr_opcd)
{
	uint8_t wc_opcd;

	switch (wr_opcd) {
	case IBV_WR_SEND_WITH_IMM:
	case IBV_WR_SEND:
		wc_opcd = IBV_WC_SEND;
		break;
	case IBV_WR_RDMA_WRITE_WITH_IMM:
	case IBV_WR_RDMA_WRITE:
		wc_opcd = IBV_WC_RDMA_WRITE;
		break;
	case IBV_WR_RDMA_READ:
		wc_opcd = IBV_WC_RDMA_READ;
		break;
	case IBV_WR_ATOMIC_CMP_AND_SWP:
		wc_opcd = IBV_WC_COMP_SWAP;
		break;
	case IBV_WR_ATOMIC_FETCH_AND_ADD:
		wc_opcd = IBV_WC_FETCH_ADD;
		break;
	default:
		wc_opcd = 0xFF;
		break;
	}

	return wc_opcd;
}

#define CQ_STATUS_SUCCESS		0
#define CQ_STATUS_LOCAL_LEN_ERR		1
#define CQ_STATUS_LOCAL_QP_OPER_ERR	2
#define CQ_STATUS_LOCAL_PROT_ERR	3
#define CQ_STATUS_WQE_FLUSHED_ERR	4
#define CQ_STATUS_MEM_MGMT_OPER_ERR	5
#define CQ_STATUS_BAD_RESP_ERR		6
#define CQ_STATUS_LOCAL_ACC_ERR		7
#define CQ_STATUS_REMOTE_INV_REQ_ERR	8
#define CQ_STATUS_REMOTE_ACC_ERR	9
#define CQ_STATUS_REMOTE_OPER_ERR	10
#define CQ_STATUS_RETRY_EXCEEDED	11
#define CQ_STATUS_RNR_RETRY_EXCEEDED	12
#define CQ_STATUS_XRC_VIO_ERR		13

#define OP_TYPE_SEND			0
#define OP_TYPE_SEND_INV		1
#define OP_TYPE_SEND_IMM		2
#define OP_TYPE_READ			3
#define OP_TYPE_WRITE			4
#define OP_TYPE_WRITE_IMM		5
#define OP_TYPE_CMP_N_SWAP		6
#define OP_TYPE_FETCH_N_ADD		7
#define OP_TYPE_FRPNR			8
#define OP_TYPE_LOCAL_INV		9
#define OP_TYPE_BIND_MW			10
#define OP_TYPE_SEND_INV_IMM		11 // vendor specific

#define OP_TYPE_RDMA_OPER_WITH_IMM	16
#define OP_TYPE_SEND_RCVD		17
#define OP_TYPE_INVALID			18

static inline enum ibv_wc_opcode ionic_to_ibv_wc_opcd(uint8_t ionic_opcd)
{
	enum ibv_wc_opcode ibv_opcd;

	/* XXX should this use ionic_wc_type instead? */
	switch (ionic_opcd) {
	case OP_TYPE_SEND:
	case OP_TYPE_SEND_INV:
	case OP_TYPE_SEND_IMM:
		ibv_opcd = IBV_WC_SEND;
		break;
	case OP_TYPE_READ:
		ibv_opcd = IBV_WC_RDMA_READ;
		break;
	case OP_TYPE_WRITE:
	case OP_TYPE_WRITE_IMM:
		ibv_opcd = IBV_WC_RDMA_WRITE;
		break;
	case OP_TYPE_CMP_N_SWAP:
		ibv_opcd = IBV_WC_COMP_SWAP;
		break;
	case OP_TYPE_FETCH_N_ADD:
		ibv_opcd = IBV_WC_FETCH_ADD;
		break;
	case OP_TYPE_LOCAL_INV:
		ibv_opcd = IBV_WC_LOCAL_INV;
		break;
	case OP_TYPE_BIND_MW:
		ibv_opcd = IBV_WC_BIND_MW;
		break;
	default:
		ibv_opcd = 0;
	}

	return ibv_opcd;
}

static inline uint8_t ionic_to_ibv_wc_status(uint8_t wcst)
{
	uint8_t ibv_wcst;

	/* XXX should this use ionic_{req,rsp}_wc_status instead?
	 * also, do we really need two different enums for wc status? */
	switch (wcst) {
	case 0:
		ibv_wcst = IBV_WC_SUCCESS;
		break;
	case CQ_STATUS_LOCAL_LEN_ERR:
		ibv_wcst = IBV_WC_LOC_LEN_ERR;
		break;
	case CQ_STATUS_LOCAL_QP_OPER_ERR:
		ibv_wcst = IBV_WC_LOC_QP_OP_ERR;
		break;
	case CQ_STATUS_LOCAL_PROT_ERR:
		ibv_wcst = IBV_WC_LOC_PROT_ERR;
		break;
	case CQ_STATUS_WQE_FLUSHED_ERR:
		ibv_wcst = IBV_WC_WR_FLUSH_ERR;
		break;
	case CQ_STATUS_LOCAL_ACC_ERR:
		ibv_wcst = IBV_WC_LOC_ACCESS_ERR;
		break;
	case CQ_STATUS_REMOTE_INV_REQ_ERR:
		ibv_wcst = IBV_WC_REM_INV_REQ_ERR;
		break;
	case CQ_STATUS_REMOTE_ACC_ERR:
		ibv_wcst = IBV_WC_REM_ACCESS_ERR;
		break;
	case CQ_STATUS_REMOTE_OPER_ERR:
		ibv_wcst = IBV_WC_REM_OP_ERR;
		break;
	case CQ_STATUS_RNR_RETRY_EXCEEDED:
		ibv_wcst = IBV_WC_RNR_RETRY_EXC_ERR;
		break;
	case CQ_STATUS_RETRY_EXCEEDED:
		ibv_wcst = IBV_WC_RETRY_EXC_ERR;
		break;
	default:
		ibv_wcst = IBV_WC_GENERAL_ERR;
		break;
	}

	return ibv_wcst;
}

static inline bool ionic_op_is_local(uint8_t opcd)
{
	return opcd == OP_TYPE_LOCAL_INV || opcd == OP_TYPE_BIND_MW;
}

#endif /* IONIC_H */
