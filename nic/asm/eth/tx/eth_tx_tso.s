
#include "INGRESS_p.h"
#include "ingress.h"
#include "INGRESS_tx_table_s3_t2_k.h"

#include "../../asm/eth/tx/defines.h"
#include "nic/p4/common/defines.h"

struct phv_ p;
struct tx_table_s3_t2_k_ k;
struct tx_table_s3_t2_eth_tx_tso_d d;

#define   _r_addr       r1        // Current buffer/descriptor address
#define   _r_num_sg     r2        // Remaining number of SG elements
#define   _r_ptr        r5        // Current DMA byte offset in PHV
#define   _r_index      r6        // Current DMA command index in PHV

%%

.param eth_tx_completion

.align
eth_tx_tso_start:

  bcf             [c2 | c3 | c7], eth_tx_tso_error
  nop

  // Load DMA command pointer
  add             _r_index, r0, k.eth_tx_global_dma_cur_index
  add             _r_num_sg, r0, k.eth_tx_global_num_sg_elems

  // Are we in the middle of SG?
  seq             c1, k.eth_tx_global_sg_in_progress, 1
  bcf             [c1], eth_tx_tso_cont
  nop

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_INTRINSIC(0, _r_ptr)
  DMA_CMD_NEXT(_r_index)

eth_tx_tso:
  bbeq            k.eth_tx_global_tso_sot, 1, eth_tx_tso_sot
  nop

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_TSO_HDR(0, _r_addr, _r_ptr, to_s3)
  DMA_CMD_NEXT(_r_index)

eth_tx_tso_sot:

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_HDR(0, _r_addr, _r_ptr, to_s3)
  DMA_CMD_NEXT(_r_index)

eth_tx_tso_cont:

  beqi            _r_num_sg, 0, eth_tx_tso_done

  subi            _r_num_sg, _r_num_sg, 1
  seq             c1, _r_num_sg, 0

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_FRAG(0, c1, _r_addr, _r_ptr)
  DMA_CMD_NEXT(_r_index)

  bcf             [c1], eth_tx_tso_done

  subi            _r_num_sg, _r_num_sg, 1
  seq             c1, _r_num_sg, 0

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_FRAG(1, c1, _r_addr, _r_ptr)
  DMA_CMD_NEXT(_r_index)

  bcf             [c1], eth_tx_tso_done

  subi            _r_num_sg, _r_num_sg, 1
  seq             c1, _r_num_sg, 0

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_FRAG(2, c1, _r_addr, _r_ptr)
  DMA_CMD_NEXT(_r_index)

  bcf             [c1], eth_tx_tso_done

  subi            _r_num_sg, _r_num_sg, 1
  seq             c1, _r_num_sg, 0

  DMA_CMD_PTR(_r_ptr, _r_index, r7)
  DMA_FRAG(3, c1, _r_addr, _r_ptr)
  DMA_CMD_NEXT(_r_index)

  bcf             [c1], eth_tx_tso_done
  nop

eth_tx_tso_next:   // Continue SG in next stage
  phvwri          p.eth_tx_global_sg_in_progress, 1

  // Calculate the next SG descriptor address
  add             _r_addr, k.eth_tx_global_sg_desc_addr, 1, LG2_TX_SG_MAX_READ_SIZE

  // Update the remaining number of SG elements & SG descriptor address
  phvwrpair       p.eth_tx_global_sg_desc_addr, _r_addr, p.eth_tx_global_num_sg_elems, _r_num_sg

  // Save DMA command pointer
  phvwr.e         p.eth_tx_global_dma_cur_index, _r_index

  // Launch eth_tx_tso stage
  phvwrpair       p.common_te2_phv_table_raw_table_size, LG2_TX_SG_MAX_READ_SIZE, p.common_te2_phv_table_addr, _r_addr

eth_tx_tso_done:   // We are done with SG
  phvwri          p.eth_tx_global_sg_in_progress, 0

  // Save DMA command pointer
  phvwr           p.eth_tx_global_dma_cur_index, _r_index

  // Launch eth_completion stage
  phvwri          p.common_te0_phv_table_pc, eth_tx_completion[38:6]
  phvwri.e        p.common_te0_phv_table_raw_table_size, CAPRI_RAW_TABLE_SIZE_MPU_ONLY
  phvwri.f        p.{app_header_table0_valid...app_header_table3_valid}, (1 << 3)

eth_tx_tso_error:
  phvwri.e        p.{app_header_table0_valid...app_header_table3_valid}, 0
  phvwri.f        p.p4_intr_global_drop, 1
