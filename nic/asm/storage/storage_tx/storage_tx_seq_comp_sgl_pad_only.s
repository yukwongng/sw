/*****************************************************************************
 *  seq_comp_sgl_pad_only: on entry, d points to the last Barco SGL descriptor
 *                         where the comp engine has written the last chunk of
 *                         output compressed data (to d.addr0). Here, we'll 
 *                         calculate how much it has written and pad the
 *                         remaining buffer with pad data.
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"


struct s3_tbl2_k k;
struct s3_tbl2_seq_comp_sgl_pad_only_d d;
struct phv_ p;

/*
 * Registers usage
 * Note: avoid r1 as it is implicitly used by DMA_MEM2MEM_SETUP() 
 */
#define r_pad_len                   r3  // padding length
#define r_pad_boundary              r4  // user specified padding boundary
#define r_pad_dest                  r5  // pointer to destination to apply padding

%%

storage_tx_seq_comp_sgl_pad_only:

    // An assumption we make here is, the last SGL describes a buffer of a fixed
    // block size which is 1 << STORAGE_KIVEC5_PAD_LEN_SHIFT. So, given the 
    // actual padding length in STORAGE_KIVEC3ACC_PAD_LEN, we can calculate
    // how much HW has already written to d.addr0, i.e.
    //     written_len = block_size - pad_len
    //     pad_dest = d.addr0 + written_len
    
    add         r_pad_len, STORAGE_KIVEC3ACC_PAD_LEN, r0
    beq         r_pad_len, r0, exit
    sll         r_pad_boundary, 1, STORAGE_KIVEC5_PAD_LEN_SHIFT // delay slot
    sub         r_pad_dest, r_pad_boundary, r_pad_len
    add         r_pad_dest, r_pad_dest, d.addr0
    
    DMA_MEM2MEM_SETUP(CAPRI_DMA_M2M_TYPE_SRC, STORAGE_KIVEC3ACC_PAD_BUF_ADDR,
                      r_pad_len, 0, 0, dma_m2m_2)
    DMA_MEM2MEM_SETUP_REG_ADDR(CAPRI_DMA_M2M_TYPE_DST, r_pad_dest,
                               r_pad_len, 0, 0, dma_m2m_3)

exit:
                      
   // Setup the start and end DMA pointers
   DMA_PTR_SETUP(dma_p2m_0_dma_cmd_pad, dma_p2m_11_dma_cmd_eop,
                 p4_txdma_intr_dma_cmd_ptr)

   // Exit the pipeline here
   LOAD_NO_TABLES

