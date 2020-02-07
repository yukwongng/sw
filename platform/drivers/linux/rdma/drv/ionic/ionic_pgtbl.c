// SPDX-License-Identifier: GPL-2.0 OR Linux-OpenIB
/*
 * Copyright (c) 2018-2020 Pensando Systems, Inc.  All rights reserved.
 */

#include <linux/mman.h>

#include "ionic_fw.h"
#include "ionic_ibdev.h"

__le64 ionic_pgtbl_dma(struct ionic_tbl_buf *buf, u64 va)
{
	u64 dma;

	if (!buf->tbl_pages)
		return cpu_to_le64(0);

	if (buf->tbl_pages > 1)
		return cpu_to_le64(buf->tbl_dma);

	if (buf->tbl_buf)
		dma = le64_to_cpu(buf->tbl_buf[0]);
	else
		dma = buf->tbl_dma;

	return cpu_to_le64(dma + (va & ~PAGE_MASK));
}

__be64 ionic_pgtbl_off(struct ionic_tbl_buf *buf, u64 va)
{
	if (buf->tbl_pages > 1)
		return cpu_to_be64(va & ~PAGE_MASK);
	else
		return 0;
}

int ionic_pgtbl_page(struct ionic_tbl_buf *buf, u64 dma)
{
	if (unlikely(buf->tbl_pages == buf->tbl_limit))
		return -ENOMEM;

	if (buf->tbl_buf)
		buf->tbl_buf[buf->tbl_pages] = cpu_to_le64(dma);
	else
		buf->tbl_dma = dma;

	++buf->tbl_pages;

	return 0;
}

int ionic_pgtbl_umem(struct ionic_tbl_buf *buf, struct ib_umem *umem)
{
	struct scatterlist *sg;
	u64 page_dma;
	int rc = 0, sg_i, page_i, page_count;

	for_each_sg(umem->sg_head.sgl, sg, umem->nmap, sg_i) {
		page_dma = sg_dma_address(sg);
		page_count = sg_dma_len(sg) >> PAGE_SHIFT;

		for (page_i = 0; page_i < page_count; ++page_i) {
			rc = ionic_pgtbl_page(buf, page_dma);
			if (rc)
				goto out;

			page_dma += PAGE_SIZE;
		}
	}

	buf->page_size_log2 = PAGE_SHIFT;

out:
	return rc;
}

int ionic_pgtbl_init(struct ionic_ibdev *dev, struct ionic_tbl_res *res,
		     struct ionic_tbl_buf *buf, struct ib_umem *umem,
		     dma_addr_t dma, int limit)
{
	int rc;

	if (umem)
		limit = ib_umem_num_pages(umem);

	if (limit < 1)
		return -EINVAL;

	res->tbl_order = ionic_res_order(limit, dev->pte_stride,
					 dev->cl_stride);
	res->tbl_pos = 0;

	buf->tbl_limit = limit;
	buf->tbl_pages = 0;
	buf->tbl_size = 0;
	buf->tbl_buf = NULL;
	buf->page_size_log2 = 0;

	/* skip pgtbl if contiguous / direct translation */
	if (limit > 1) {
		/* A reservation will be made for page table resources.
		 *
		 * The page table reservation must be large enough to account
		 * for alignment of the first page within the page table.  The
		 * requirement is therefore greater-than-page-size alignment.
		 * If not accounted for, the placement of pages in the page
		 * table could extend beyond the bounds of the reservation.
		 *
		 * To account for the alignment of the page in the reservation,
		 * extend the reservation based on the worst case first page
		 * alignment.  The worst case is required for alloc_mr, because
		 * only the number of pages is given, but the alignment is not
		 * known until registration.
		 */
		if (likely(dev->cl_stride > dev->pte_stride)) {
			/* increase the limit to account for worst case */
			buf->tbl_limit =
				1 + ALIGN(limit - 1, BIT(dev->cl_stride -
							 dev->pte_stride));

			/* and recalculate the reservation dimensions */
			res->tbl_order = ionic_res_order(buf->tbl_limit,
							 dev->pte_stride,
							 dev->cl_stride);
		}

		rc = ionic_get_res(dev, res);
		if (rc)
			goto err_res;

		/* limit for pte reservation should not affect anything else */
		buf->tbl_limit = limit;
	} else {
		res->tbl_order = IONIC_RES_INVALID;
	}

	if (limit > 1) {
		buf->tbl_size = buf->tbl_limit * sizeof(*buf->tbl_buf);
		buf->tbl_buf = kmalloc(buf->tbl_size, GFP_KERNEL);
		if (!buf->tbl_buf) {
			rc = -ENOMEM;
			goto err_buf;
		}

		buf->tbl_dma = dma_map_single(dev->hwdev, buf->tbl_buf,
					      buf->tbl_size, DMA_FROM_DEVICE);
		rc = dma_mapping_error(dev->hwdev, buf->tbl_dma);
		if (rc)
			goto err_dma;
	}

	if (umem) {
		rc = ionic_pgtbl_umem(buf, umem);
		if (rc)
			goto err_umem;
	} else {
		rc = ionic_pgtbl_page(buf, dma);
		if (rc)
			goto err_umem;

		buf->page_size_log2 = 0;
	}

	return 0;

err_umem:
	if (buf->tbl_buf) {
		dma_unmap_single(dev->hwdev, buf->tbl_dma, buf->tbl_size,
				 DMA_FROM_DEVICE);
err_dma:
		kfree(buf->tbl_buf);
	}
err_buf:
	ionic_put_res(dev, res);
err_res:
	buf->tbl_buf = NULL;
	buf->tbl_limit = 0;
	buf->tbl_pages = 0;
	return rc;
}

void ionic_pgtbl_unbuf(struct ionic_ibdev *dev, struct ionic_tbl_buf *buf)
{
	if (buf->tbl_buf) {
		dma_unmap_single(dev->hwdev, buf->tbl_dma, buf->tbl_size,
				 DMA_FROM_DEVICE);
		kfree(buf->tbl_buf);
	}

	buf->tbl_buf = NULL;
	buf->tbl_limit = 0;
	buf->tbl_pages = 0;
}
