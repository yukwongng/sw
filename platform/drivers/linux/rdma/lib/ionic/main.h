#ifndef __MAIN_H__
#define __MAIN_H__

#include <inttypes.h>
#include <stdbool.h>
#include <stddef.h>
#include <endian.h>
#include <pthread.h>

#include <infiniband/driver.h>
#include <util/udma_barrier.h>

#include "ionic_dbg.h"
#include "ionic-abi.h"
#include "ionic_queue.h"
#include "memory.h"
#include "table.h"

#define DEV	"ionic : "

#define IONIC_UD_QP_HW_STALL 0x400000

struct ionic_dpi {
	struct ionic_doorbell    *dbpage;
	pthread_spinlock_t        db_lock;
};

struct ionic_pd {
	struct ibv_pd ibvpd;
	uint32_t pdid;
};

struct ionic_context {
	struct verbs_context vctx;
	uint32_t max_qp;
	struct ionic_dpi udpi;

	pthread_mutex_t		mut;
	struct tbl_root		qp_tbl;
	struct list_head	cq_list;
};

struct ionic_cq {
	struct ibv_cq		ibvcq;
	struct list_node	ctx_ent;

	uint32_t		cqid;

	pthread_spinlock_t	lock;
	struct ionic_queue	q;

	/* XXX cleanup */
	uint8_t			qtype;
	struct ionic_dpi	*udpi;
	struct ionic_context	*cntxt;
};

struct ionic_srq {
	struct ibv_srq ibvsrq;
};

struct ionic_wrid {
	struct ionic_psns *psns;
	uint64_t wrid;
	uint32_t bytes;
	uint8_t sig;
};

struct ionic_qpcap {
	uint32_t max_swr;
	uint32_t max_rwr;
	uint32_t max_ssge;
	uint32_t max_rsge;
	uint32_t max_inline;
	uint8_t	sqsig;
};

struct ionic_sq_meta {
	uint64_t		wrid;
	uint32_t		len;
	uint8_t			op;
};

struct ionic_rq_meta {
	uint64_t		wrid;
	uint32_t		len; /* XXX byte_len must come from cqe */
};

struct ionic_qp {
	struct ibv_qp		ibvqp;

	uint32_t qpid;

	pthread_spinlock_t	sq_lock;
	struct ionic_queue	sq;
	struct ionic_sq_meta	*sq_meta;

	uint32_t		sq_local;
	uint32_t		sq_msn;

	pthread_spinlock_t	rq_lock;
	struct ionic_queue	rq;
	struct ionic_rq_meta	*rq_meta;

	/* XXX cleanup */
	struct ionic_srq *srq;
	struct ionic_cq *scq;
	struct ionic_cq *rcq;
	struct ionic_dpi *udpi;
	struct ionic_qpcap cap;
    struct ionic_context *cntxt;
    
	uint32_t tbl_indx;
	uint32_t sq_psn;
	uint32_t ionicding_db;
	uint64_t wqe_cnt;
	uint16_t mtu;
	uint16_t qpst;
	uint8_t qptype;
    uint8_t sq_qtype;
    uint8_t rq_qtype;
	/* irdord? */
};

struct ionic_mr {
	struct ibv_mr ibvmr;
};

struct ionic_ah {
	struct ibv_ah ibvah;
	uint32_t avid;
};

struct ionic_dev {
	struct verbs_device vdev;
	uint8_t abi_version;
	size_t pg_size;

	uint32_t cqe_size;
	uint32_t max_cq_depth;
};

/* pointer conversion functions*/
static inline struct ionic_dev *to_ionic_dev(struct ibv_device *ibvdev)
{
	return container_of(ibvdev, struct ionic_dev, vdev.device);
}

static inline struct ionic_context *to_ionic_context(
		struct ibv_context *ibvctx)
{
	return container_of(ibvctx, struct ionic_context, vctx.context);
}

static inline struct ionic_pd *to_ionic_pd(struct ibv_pd *ibvpd)
{
	return container_of(ibvpd, struct ionic_pd, ibvpd);
}

static inline struct ionic_cq *to_ionic_cq(struct ibv_cq *ibvcq)
{
	return container_of(ibvcq, struct ionic_cq, ibvcq);
}

static inline struct ionic_qp *to_ionic_qp(struct ibv_qp *ibvqp)
{
	return container_of(ibvqp, struct ionic_qp, ibvqp);
}

static inline struct ionic_ah *to_ionic_ah(struct ibv_ah *ibvah)
{
        return container_of(ibvah, struct ionic_ah, ibvah);
}

static inline uint32_t ionic_get_sqe_size(void)
{
	return sizeof(struct sqwqe_t);
}

static inline uint32_t ionic_get_rqe_size(void)
{
	return sizeof(struct rqwqe_t);
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

#define     CQ_STATUS_SUCCESS               0
#define     CQ_STATUS_LOCAL_LEN_ERR         1
#define     CQ_STATUS_LOCAL_QP_OPER_ERR     2
#define     CQ_STATUS_LOCAL_PROT_ERR        3
#define     CQ_STATUS_WQE_FLUSHED_ERR       4
#define     CQ_STATUS_MEM_MGMT_OPER_ERR     5
#define     CQ_STATUS_BAD_RESP_ERR          6
#define     CQ_STATUS_LOCAL_ACC_ERR         7
#define     CQ_STATUS_REMOTE_INV_REQ_ERR    8
#define     CQ_STATUS_REMOTE_ACC_ERR        9
#define     CQ_STATUS_REMOTE_OPER_ERR       10
#define     CQ_STATUS_RETRY_EXCEEDED        11
#define     CQ_STATUS_RNR_RETRY_EXCEEDED    12
#define     CQ_STATUS_XRC_VIO_ERR           13

#define OP_TYPE_SEND                0
#define OP_TYPE_SEND_INV            1
#define OP_TYPE_SEND_IMM            2
#define OP_TYPE_READ                3
#define OP_TYPE_WRITE               4
#define OP_TYPE_WRITE_IMM           5
#define OP_TYPE_CMP_N_SWAP          6
#define OP_TYPE_FETCH_N_ADD         7
#define OP_TYPE_FRPNR               8
#define OP_TYPE_LOCAL_INV           9
#define OP_TYPE_BIND_MW             10
#define OP_TYPE_SEND_INV_IMM        11 // vendor specific

#define OP_TYPE_RDMA_OPER_WITH_IMM 16
#define OP_TYPE_SEND_RCVD          17
#define OP_TYPE_INVALID            18

static inline enum ibv_wc_opcode ionic_to_ibv_wc_opcd(uint8_t ionic_opcd)
{
	enum ibv_wc_opcode ibv_opcd;

	/* XXX should this use ionic_wc_type instead? */
	switch (ionic_opcd) {
	case OP_TYPE_SEND:
	case OP_TYPE_SEND_INV:
	case OP_TYPE_SEND_IMM:
		ionic_opcd = IBV_WC_SEND;
		break;
	case OP_TYPE_READ:
		ionic_opcd = IBV_WC_RDMA_READ;
		break;
	case OP_TYPE_WRITE:
	case OP_TYPE_WRITE_IMM:
		ionic_opcd = IBV_WC_RDMA_WRITE;
		break;
	case OP_TYPE_CMP_N_SWAP:
		ionic_opcd = IBV_WC_COMP_SWAP;
		break;
	case OP_TYPE_FETCH_N_ADD:
		ionic_opcd = IBV_WC_FETCH_ADD;
		break;
	case OP_TYPE_LOCAL_INV:
		ionic_opcd = IBV_WC_LOCAL_INV;
		break;
	case OP_TYPE_BIND_MW:
		ionic_opcd = IBV_WC_BIND_MW;
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
	return opcd == OP_TYPE_LOCAL_INV ||
		opcd == OP_TYPE_BIND_MW;
}

#if 0
static inline uint8_t ionic_is_cqe_valid(struct ionic_cq *cq,
					   struct ionic_bcqe *hdr)
{
	udma_from_device_barrier();
	return ((le32toh(hdr->flg_st_typ_ph) &
		 IONIC_BCQE_PH_MASK) == cq->phase);
}

static inline void ionic_change_cq_phase(struct ionic_cq *cq)
{
	if (!cq->cqq.head)
		cq->phase = (~cq->phase & IONIC_BCQE_PH_MASK);
}
#endif

static inline void ionic_lock_all_cqs(struct ionic_context *ctx)
{
	struct ionic_cq *cq;

	/* To avoid deadlock, this is the only place where a thread is allowed
	 * to acquire the lock for more than one cq.  The order in the list is
	 * the locking order.  All other cq lock holders may acquire at most
	 * one cq lock at a time.
	 *
	 * All cq locks must be held when modifying the qp_tbl.  Any one cq
	 * lock can be held to exclude modifying the table.  A cq polling
	 * thread should only acquire the lock for that cq.  Other cq may be
	 * polled by other threads in parallel.  Unlike a reader-writer lock,
	 * any one cq may be polled by at most one thread at a time.
	 *
	 * In addition to cq locks, the context mutex must also be held when
	 * modifying the qp table or accessing the cq list.  The mutex must
	 * have already been acquired by the caller of this function.
	 */
	list_for_each(&ctx->cq_list, cq, ctx_ent)
		pthread_spin_lock(&cq->lock);
}

static inline void ionic_unlock_all_cqs(struct ionic_context *ctx)
{
	struct ionic_cq *cq;

	/* unlock in reverse order of acquire */
	list_for_each_rev(&ctx->cq_list, cq, ctx_ent)
		pthread_spin_unlock(&cq->lock);
}

#endif
