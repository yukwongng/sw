/*****************************************************************************
 *  seq_xts_sgl_pdma_xfer: Parse the destination SGL and DMA the data from
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"


struct s3_tbl1_k k;
struct s3_tbl1_seq_comp_sgl_pdma_xfer_d d;
struct phv_ p;

/*
 * Registers usage
 */
#define r_src_len                   r2  // length of data source
#define r_xfer_len                  r3  // current transfer length
#define r_src_addr                  r4  // source address
#define r_dst_addr                  r5  // destination address
#define r_sgl_len                   r6  // current length of an SGL entry

/*
 * Registers reuse, post source transfer completion
 */
#define r_src_qaddr                 r_sgl_len  // for SEQ_METRICS_TABLE_COMMIT
 
%%
   SEQ_METRICS_PARAMS()

storage_seq_xts_sgl_pdma_xfer:

   // Store the data length into r_src_len accouting for the fact that 0 => 64K
   add          r_src_len, SEQ_KIVEC5XTS_DATA_LEN, r0
   add		r_src_addr, SEQ_KIVEC3XTS_DECR_BUF_ADDR, r0

   // Setup the compression data buffer DMA based on flat source buffer 
   // and destination SGL.The macro processes one SGL entry and branches
   // if xfer is complete at any point.
   // Notes: These GPRs are used for input/output to/from this macro
   //  1. r_src_len stores the running count of data remaining to be xfered
   //  2. r_src_addr stores the offeset of the source data buffer from where
   //     the current xfer is to be done.
   //
   /*
    * VERY IMPORTANT NOTE: mem2mem descriptors work in adjacent pair and the
    * pair must not cross flit boundary. P4+ code sets up PHV space which 
    * guarantees that a valid mem2mem pair always starts with an even numbered ID.
    *
    * For example: dma_m2m_2/dma_m2m_3 would be a valid pair, 
    *              but dma_m2m_7/dma_m2m_8 would not necessarily be adjacent.
    */
 
   // Can process the entire PDMA SGL here which holds 4 addr/len pairs,
   // plus padding
   
   CHAIN_SGL_PDMA(dma_m2m_4, dma_m2m_5,
                  d.addr0, d.len0, pdma_xfer_done, pdma_xfer_error)
   CHAIN_SGL_PDMA(dma_m2m_6, dma_m2m_7,
                  d.addr1, d.len1, pdma_xfer_done, pdma_xfer_error)
   CHAIN_SGL_PDMA(dma_m2m_8, dma_m2m_9,
                  d.addr2, d.len2, pdma_xfer_done, pdma_xfer_error)
   CHAIN_SGL_PDMA(dma_m2m_10, dma_m2m_11,
                  d.addr3, d.len3, pdma_xfer_done, pdma_xfer_error)

   // Catch any driver errors here for debugging, i.e., driver did not 
   // provision the SGL correctly relative to comp output data length.
   bne          r_src_len, r0, pdma_xfer_error
   nop
   
pdma_xfer_done:   
   SEQ_METRICS_SET(sgl_pdma_xfers)
   
exit:
   // Relaunch stats commit for table 1
   SEQ_METRICS0_TABLE1_COMMIT_e(SEQ_KIVEC5XTS_SRC_QADDR)

pdma_xfer_error:
   SEQ_METRICS_SET(sgl_pdma_errs)
   SEQ_COMP_SGL_PDMA_XFER_ERROR_TRAP()
   b            exit
   nop
