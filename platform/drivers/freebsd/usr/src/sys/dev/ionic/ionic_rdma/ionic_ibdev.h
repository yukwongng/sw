#ifndef IONIC_IBDEV_H
#define IONIC_IBDEV_H

#include <linux/device.h>
#include <linux/netdevice.h>
#include <rdma/ib_umem.h>
#include <rdma/ib_verbs.h>
#include <rdma/ib_pack.h>

//#include <rdma/ionic-abi.h>
#include <ionic_api.h>

#include "ionic_fw.h"
#include "ionic_queue.h"
#include "table.h"

#define IONIC_MIN_RDMA_VERSION 0
#define IONIC_MAX_RDMA_VERSION 1

#define IONIC_DBPAGE_SIZE	0x1000
#define IONIC_MAX_MRID		0xffffff
#define IONIC_MAX_QPID		0xffffff
#define IONIC_MAX_HBM_ORDER	15

struct ionic_aq;
struct ionic_cq;
struct ionic_eq;

enum ionic_admin_state {
	IONIC_ADMIN_ACTIVE, /* subbmitting admin commands to queue */
	IONIC_ADMIN_PAUSED, /* not submitting, but may complete normally */
	IONIC_ADMIN_KILLED, /* not submitting, fake normal completion */
	IONIC_ADMIN_FAILED, /* not submitting, failed completion */
};

struct ionic_mmap_info {
	struct list_head ctx_ent;
	unsigned long offset;
	unsigned long size;
	unsigned long pfn;
};

struct ionic_ibdev {
	struct ib_device	ibdev;

	struct list_head	driver_ent;

	struct ib_device_attr	dev_attr;
	struct ib_port_attr	port_attr;

	struct device		*hwdev;
	struct net_device	*ndev;
	struct lif		*lif;
	int			lif_id;

	u32		__iomem *intr_ctrl;

	phys_addr_t		phys_dbpage_base;
	u64		__iomem *dbpage;
	u32			dbid;

	u16			rdma_version;
	u8			qp_opcodes;
	u8			admin_opcodes;

	u8			admin_qtype;
	u8			sq_qtype;
	u8			rq_qtype;
	u8			cq_qtype;
	u8			eq_qtype;

	struct mutex		tbl_lock; /* for modify cq_tbl, qp_tbl */
	rwlock_t		tbl_rcu; /* instead of synchronize_rcu() */

	struct tbl_root		qp_tbl;
	struct tbl_root		cq_tbl;

	struct mutex		inuse_lock; /* for id reservation */

	unsigned long		*inuse_pdid;
	u32			size_pdid;
	u32			next_pdid;

	unsigned long		*inuse_ahid;
	u32			size_ahid;
	u32			next_ahid;

	unsigned long		*inuse_mrid;
	u32			size_mrid;
	u32			next_mrid;
	u8			next_rkey_key;

	unsigned long		*inuse_eqid;
	u32			size_eqid;
	u32			next_eqid;

	unsigned long		*inuse_cqid;
	u32			size_cqid;
	u32			next_cqid;

	unsigned long		*inuse_qpid;
	u32			size_qpid;
	u32			next_qpid;
	u32			size_srqid;
	u32			next_srqid;

	unsigned long		*inuse_pgtbl;
	u32			size_pgtbl;

	struct work_struct	admin_work;
	spinlock_t		admin_lock;
	struct ionic_aq		*adminq;
	struct ionic_cq		*admincq;
	bool			admin_armed;
	enum ionic_admin_state	admin_state;

	struct ionic_eq		**eq_vec;
	int			eq_count;
};

struct ionic_eq {
	struct ionic_ibdev	*dev;

	u32			eqid;
	u32			intr;

	struct ionic_queue	q;

	bool			enable;
	bool			armed;

	struct work_struct	work;

	int			vec;
	int			cpu;
	int			irq;
	char			name[32];
};

struct ionic_admin_wr {
	struct completion	work;
	struct list_head	aq_ent;
	struct ionic_v1_admin_wqe wqe;
	struct ionic_v1_cqe	cqe;
	int			status;
};

struct ionic_aq {
	struct ionic_ibdev	*dev;

	u32			aqid;
	u32			cqid;

	spinlock_t		lock; /* for posting */
	struct ionic_queue	q;
	struct ionic_admin_wr	**q_wr;
	struct list_head	wr_prod;
	struct list_head	wr_post;
};

struct ionic_ctx {
	struct ib_ucontext	ibctx;

	/* fallback to kernel-impl verbs for this user context */
	bool			fallback;

	u32			dbid;

	struct mutex		mmap_mut; /* for mmap_list */
	unsigned long long	mmap_off;
	struct list_head	mmap_list;
	struct ionic_mmap_info	mmap_dbell;
};

struct ionic_pd {
	struct ib_pd		ibpd;

	u32			pdid;
};

struct ionic_cq {
	struct ib_cq		ibcq;

	u32			cqid;
	u32			eqid;

	spinlock_t		lock; /* for polling */
	struct list_head	qp_poll;
	struct ionic_queue	q;

	/* infrequently accessed, keep at end */
	struct ib_umem		*umem;
	u32			tbl_pos;
	int			tbl_order;

	/* XXX cleanup */
	u32			lkey;
	struct delayed_work	notify_work;
};

struct ionic_sq_meta {
	u64			wrid;
	u32			len;
	u16			seq;
	u8			op;
	u8			status;
	bool			remote;
	bool			signal;
};

struct ionic_rq_meta {
	u64			wrid;
	u32			len; /* XXX byte_len must come from cqe */
};

struct ionic_qp {
	union {
		struct ib_qp	ibqp;
		struct ib_srq	ibsrq;
	};

	u32			qpid;
	u32			ahid;

	bool			has_ah;
	bool			has_sq;
	bool			has_rq;
	bool			is_srq;

	bool			sig_all;

	struct list_head	cq_poll_ent;

	spinlock_t		sq_lock; /* for posting and polling */
	struct ionic_queue	sq;
	struct ionic_sq_meta	*sq_meta;
	u16			*sq_msn_idx;
	u16			sq_msn_prod;
	u16			sq_msn_cons;
	u16			sq_npg_cons;

	void			__iomem *sq_hbm_ptr;
	u16			sq_hbm_prod;

	spinlock_t		rq_lock; /* for posting and polling */
	struct ionic_queue	rq;
	struct ionic_rq_meta	*rq_meta; /* XXX this rq_meta will go away */

	/* infrequently accessed, keep at end */
	bool			sq_is_hbm;
	int			sq_hbm_order;
	u32			sq_hbm_pgid;
	phys_addr_t		sq_hbm_addr;
	struct ionic_mmap_info	sq_hbm_mmap;

	struct ib_umem		*sq_umem;
	int			sq_tbl_order;
	u32			sq_tbl_pos;

	struct ib_umem		*rq_umem;
	int			rq_tbl_order;
	u32			rq_tbl_pos;

	/* XXX cleanup */
	u32			sq_lkey;
	u32			rq_lkey;
};

struct ionic_ah {
	struct ib_ah		ibah;
	u32			ahid;
};

struct ionic_mr {
	struct ib_mr		ibmr;

	struct ib_umem		*umem;

	u32			tbl_pos;
	int			tbl_order;
};

static inline struct ionic_ibdev *to_ionic_ibdev(struct ib_device *ibdev)
{
	return container_of(ibdev, struct ionic_ibdev, ibdev);
}

static inline struct ionic_ctx *to_ionic_ctx(struct ib_ucontext *ibctx)
{
	return container_of(ibctx, struct ionic_ctx, ibctx);
}

static inline struct ionic_ctx *to_ionic_ctx_fb(struct ib_ucontext *ibctx)
{
	struct ionic_ctx *ctx;

	if (!ibctx)
		return NULL;

	ctx = to_ionic_ctx(ibctx);

	if (ctx->fallback)
		return NULL;

	return ctx;
}

static inline struct ionic_ctx *to_ionic_ctx_uobj(struct ib_uobject *uobj)
{
	struct ionic_ctx *ctx;

	if (!uobj)
		return NULL;

	ctx = to_ionic_ctx(uobj->context);

	if (ctx->fallback)
		return NULL;

	return ctx;
}

static inline struct ionic_pd *to_ionic_pd(struct ib_pd *ibpd)
{
	return container_of(ibpd, struct ionic_pd, ibpd);
}

static inline struct ionic_mr *to_ionic_mr(struct ib_mr *ibmr)
{
	return container_of(ibmr, struct ionic_mr, ibmr);
}

static inline struct ionic_cq *to_ionic_cq(struct ib_cq *ibcq)
{
	return container_of(ibcq, struct ionic_cq, ibcq);
}

static inline struct ionic_qp *to_ionic_qp(struct ib_qp *ibqp)
{
	return container_of(ibqp, struct ionic_qp, ibqp);
}

static inline struct ionic_qp *to_ionic_srq(struct ib_srq *ibsrq)
{
	return container_of(ibsrq, struct ionic_qp, ibsrq);
}

static inline struct ionic_ah *to_ionic_ah(struct ib_ah *ibah)
{
        return container_of(ibah, struct ionic_ah, ibah);
}

static inline u32 ionic_dbid(struct ionic_ibdev *dev,
					   struct ib_uobject *uobj)
{
	struct ionic_ctx *ctx = to_ionic_ctx_uobj(uobj);

	if (!ctx)
		return dev->dbid;

	return ctx->dbid;
}

static inline u32 ionic_idver(struct ionic_ibdev *dev, u32 id)
{
	return id | (dev->rdma_version << 24);
}

enum ionic_intr_bits {
	/* register set */
	IONIC_INTR_REGS_PER		= 8, /* eight registers per intr */
	IONIC_INTR_REG_COALESCE_INIT	= 0, /* coalesce timer reset value */
	IONIC_INTR_REG_MASK		= 1, /* mask interrupts */
	IONIC_INTR_REG_CREDITS		= 2, /* update credits */
	IONIC_INTR_REG_MASK_ASSERT	= 3, /* mask interrupts on assert */
	IONIC_INTR_REG_COALESCE		= 4, /* coalesce timer */

	/* mask (and mask-on-assert) values */
	IONIC_INTR_MASK_CLEAR		= 0u, /* unmask interrupts */
	IONIC_INTR_MASK_SET		= 1u, /* mask interrupts */

	/* credits values */
	IONIC_INTR_CRED_UNMASK		= 0x10000u, /* unmask interrupts */
	IONIC_INTR_CRED_RESET_COALESCE	= 0x20000u, /* reset coalesce timer */
};

static inline void ionic_intr_coalesce_init(struct ionic_ibdev *dev,
					    int intr, u32 usec)
{
	intr *= IONIC_INTR_REGS_PER;
	intr += IONIC_INTR_REG_COALESCE_INIT;

	iowrite32(usec, &dev->intr_ctrl[intr]);
}

static inline void ionic_intr_mask(struct ionic_ibdev *dev,
				   int intr, u32 mask)
{
	intr *= IONIC_INTR_REGS_PER;
	intr += IONIC_INTR_REG_MASK;

	iowrite32(mask, &dev->intr_ctrl[intr]);
}

static inline void ionic_intr_credits(struct ionic_ibdev *dev,
				      int intr, u32 cred)
{
	intr *= IONIC_INTR_REGS_PER;
	intr += IONIC_INTR_REG_CREDITS;

	iowrite32(cred, &dev->intr_ctrl[intr]);
}

static inline void ionic_intr_mask_assert(struct ionic_ibdev *dev,
					  int intr, u32 mask)
{
	intr *= IONIC_INTR_REGS_PER;
	intr += IONIC_INTR_REG_MASK_ASSERT;

	iowrite32(mask, &dev->intr_ctrl[intr]);
}

#endif /* IONIC_IBDEV_H */
