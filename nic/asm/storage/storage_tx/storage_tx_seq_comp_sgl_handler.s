/*****************************************************************************
 *  seq_comp_sgl_handler: Parse the destination SGL and DMA the status,
 *                        data (if status was success) and set the interrupt.
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"


struct s3_tbl_k k;
struct s3_tbl_seq_comp_sgl_handler_d d;
struct phv_ p;

%%

storage_tx_seq_comp_sgl_handler_start:

   // Store the data length into r6 accouting for the fact that 0 => 64K
   add          r6, STORAGE_KIVEC5_DATA_LEN, r0

   // Store the data address into r7
   // TODO: handle the case of data_addr being an SGL address!
   
   add		r7, STORAGE_KIVEC3_DATA_ADDR, r0

   // Setup the compression data buffer DMA based on flat source buffer 
   // and destination SGL.The macro processes one SGL entry and branches
   // if xfer is complete at any point.
   // Notes: These GPRs are used for input/output to/from this macro
   //  1. r6 stores the running count of data remaining to be xfered
   //  2. r7 stores the offeset of the source data buffer from where
   //     the current xfer is to be done.
   COMP_SGL_DMA(dma_m2m_3, dma_m2m_4, d.addr0, d.len0, complete_dma)
   COMP_SGL_DMA(dma_m2m_5, dma_m2m_6, d.addr1, d.len1, complete_dma)
   COMP_SGL_DMA(dma_m2m_7, dma_m2m_8, d.addr2, d.len2, complete_dma)
   COMP_SGL_DMA(dma_m2m_9, dma_m2m_10, d.addr3, d.len3, complete_dma)

complete_dma:
   // Setup the start and end DMA pointers
   DMA_PTR_SETUP(dma_p2m_0_dma_cmd_pad, dma_p2m_11_dma_cmd_eop,
                 p4_txdma_intr_dma_cmd_ptr)

   // Exit the pipeline here
   LOAD_NO_TABLES

