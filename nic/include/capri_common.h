#ifndef __CAPRI_COMMON_H
#define __CAPRI_COMMON_H

/*
 * #defines shared by ASM and C code
 */

#define CAPRI_HBM_OFFSET(x)     (0x80000000 + (x))

// Memory regions
#define CAPRI_MEM_SECURE_RAM_START          0
#define CAPRI_MEM_FLASH_START               0x20000
#define CAPRI_MEM_INTR_START                0x40000
#define CAPRI_MEM_SEMA_START                0x40000000

// Semaphores
#define CAPRI_SEM_RAW_OFFSET                0x1000
#define CAPRI_SEM_INC_OFFSET                0x2000
#define CAPRI_SEM_INC_NOT_FULL_OFFSET       0x4000
#define CAPRI_SEM_ATOMIC_ADD_BYTE_OFFSET    0x8000000
#define CAPRI_SEM_INF_OFFSET                CAPRI_SEM_INC_NOT_FULL_OFFSET
#define CAPRI_SEM_INC_NOT_FULL_SIZE         8
#define CAPRI_SEM_INC_NOT_FULL_PI_OFFSET    0
#define CAPRI_SEM_INC_NOT_FULL_CI_OFFSET    4

#define CAPRI_MEM_SEM_RAW_START \
                    (CAPRI_MEM_SEMA_START + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_MEM_SEM_INC_START \
                    (CAPRI_MEM_SEMA_START + CAPRI_SEM_INC_OFFSET)
#define CAPRI_MEM_SEM_INF_START \
                    (CAPRI_MEM_SEMA_START + CAPRI_SEM_INC_NOT_FULL_OFFSET)
#define CAPRI_MEM_SEM_ATOMIC_ADD_START \
                    (CAPRI_MEM_SEMA_START + CAPRI_SEM_ATOMIC_ADD_BYTE_OFFSET)

// 1K 32 bit semaphores or 512 PI/CI (inc not full) pairs
#define CAPRI_RNMDR_RING_SIZE               16384
#define CAPRI_RNMDR_RING_SHIFT              14
#define CAPRI_SEM_RNMDR_ALLOC_ADDR          (CAPRI_MEM_SEMA_START + 8 * 0)
#define CAPRI_SEM_RNMDR_ALLOC_RAW_ADDR \
                            (CAPRI_SEM_RNMDR_ALLOC_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_RNMDR_ALLOC_INF_ADDR \
                            (CAPRI_SEM_RNMDR_ALLOC_ADDR + CAPRI_SEM_INF_OFFSET)

#define CAPRI_SEM_RNMDR_FREE_ADDR          (CAPRI_MEM_SEMA_START + 8 * 1)
#define CAPRI_SEM_RNMDR_FREE_RAW_ADDR \
                            (CAPRI_SEM_RNMDR_FREE_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_RNMDR_FREE_INF_ADDR \
                            (CAPRI_SEM_RNMDR_FREE_ADDR + CAPRI_SEM_INF_OFFSET)


#define CAPRI_RNMPR_RING_SIZE               16384
#define CAPRI_RNMPR_RING_SHIFT              14
#define CAPRI_SEM_RNMPR_ALLOC_ADDR          (CAPRI_MEM_SEMA_START + 8 * 2)
#define CAPRI_SEM_RNMPR_ALLOC_RAW_ADDR \
                            (CAPRI_SEM_RNMPR_ALLOC_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_RNMPR_ALLOC_INF_ADDR \
                            (CAPRI_SEM_RNMPR_ALLOC_ADDR + CAPRI_SEM_INF_OFFSET)

#define CAPRI_SEM_RNMPR_FREE_ADDR          (CAPRI_MEM_SEMA_START + 8 * 3)
#define CAPRI_SEM_RNMPR_FREE_RAW_ADDR \
                            (CAPRI_SEM_RNMPR_FREE_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_RNMPR_FREE_INF_ADDR \
                            (CAPRI_SEM_RNMPR_FREE_ADDR + CAPRI_SEM_INF_OFFSET)


#define CAPRI_TNMDR_RING_SIZE               16384
#define CAPRI_TNMDR_RING_SHIFT              14
#define CAPRI_SEM_TNMDR_ALLOC_ADDR          (CAPRI_MEM_SEMA_START + 8 * 4)
#define CAPRI_SEM_TNMDR_ALLOC_RAW_ADDR \
                            (CAPRI_SEM_TNMDR_ALLOC_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_TNMDR_ALLOC_INF_ADDR \
                            (CAPRI_SEM_TNMDR_ALLOC_ADDR + CAPRI_SEM_INF_OFFSET)

#define CAPRI_SEM_TNMDR_FREE_ADDR          (CAPRI_MEM_SEMA_START + 8 * 5)
#define CAPRI_SEM_TNMDR_FREE_RAW_ADDR \
                            (CAPRI_SEM_TNMDR_FREE_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_TNMDR_FREE_INF_ADDR \
                            (CAPRI_SEM_TNMDR_FREE_ADDR + CAPRI_SEM_INF_OFFSET)


#define CAPRI_TNMPR_RING_SIZE               16384
#define CAPRI_TNMPR_RING_SHIFT              14
#define CAPRI_SEM_TNMPR_ALLOC_ADDR          (CAPRI_MEM_SEMA_START + 8 * 6)
#define CAPRI_SEM_TNMPR_ALLOC_RAW_ADDR \
                            (CAPRI_SEM_TNMPR_ALLOC_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_TNMPR_ALLOC_INF_ADDR \
                            (CAPRI_SEM_TNMPR_ALLOC_ADDR + CAPRI_SEM_INF_OFFSET)
#define CAPRI_SEM_TNMPR_FREE_ADDR           (CAPRI_MEM_SEMA_START + 8 * 7)
#define CAPRI_SEM_TNMPR_FREE_RAW_ADDR \
                            (CAPRI_SEM_TNMPR_FREE_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_TNMPR_FREE_INF_ADDR \
                            (CAPRI_SEM_TNMPR_FREE_ADDR + CAPRI_SEM_INF_OFFSET)



#define CAPRI_RNMPR_SMALL_RING_SIZE         16384
#define CAPRI_SEM_RNMPR_SMALL_ALLOC_ADDR    (CAPRI_MEM_SEMA_START + 8 * 8)
#define CAPRI_SEM_RNMPR_SMALL_ALLOC_RAW_ADDR \
                            (CAPRI_SEM_RNMPR_SMALL_ALLOC_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_RNMPR_SMALL_ALLOC_INF_ADDR \
                            (CAPRI_SEM_RNMPR_SMALL_ALLOC_ADDR + CAPRI_SEM_INF_OFFSET)

#define CAPRI_SEM_RNMPR_SMALL_FREE_ADDR    (CAPRI_MEM_SEMA_START + 8 * 9)
#define CAPRI_SEM_RNMPR_SMALL_FREE_RAW_ADDR \
                            (CAPRI_SEM_RNMPR_SMALL_FREE_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_RNMPR_SMALL_FREE_INF_ADDR \
                            (CAPRI_SEM_RNMPR_SMALL_FREE_ADDR + CAPRI_SEM_INF_OFFSET)


#define CAPRI_TNMPR_SMALL_RING_SIZE         16384

#define CAPRI_SEM_TNMPR_SMALL_ALLOC_ADDR    (CAPRI_MEM_SEMA_START + 8 * 10)
#define CAPRI_SEM_TNMPR_SMALL_ALLOC_RAW_ADDR \
                            (CAPRI_SEM_TNMPR_SMALL_ALLOC_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_TNMPR_SMALL_ALLOC_INF_ADDR \
                            (CAPRI_SEM_TNMPR_SMALL_ALLOC_ADDR + CAPRI_SEM_INF_OFFSET)
#define CAPRI_SEM_TNMPR_SMALL_FREE_ADDR    (CAPRI_MEM_SEMA_START + 8 * 11)
#define CAPRI_SEM_TNMPR_SMALL_FREE_RAW_ADDR \
                            (CAPRI_SEM_TNMPR_SMALL_FREE_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_TNMPR_SMALL_FREE_INF_ADDR \
                            (CAPRI_SEM_TNMPR_SMALL_FREE_ADDR + CAPRI_SEM_INF_OFFSET)

#define CAPRI_SEM_IPSECCB_ADDR                (CAPRI_MEM_SEMA_START + 8 * 6)
#define CAPRI_IPSECCB_RING_SIZE               16384
#define CAPRI_SEM_IPSECCB_RAW_ADDR \
                            (CAPRI_SEM_IPSECCB_ADDR + CAPRI_SEM_RAW_OFFSET)
#define CAPRI_SEM_IPSECCB_INF_ADDR \
                            (CAPRI_SEM_IPSECCB_ADDR + CAPRI_SEM_INF_OFFSET)
#endif
