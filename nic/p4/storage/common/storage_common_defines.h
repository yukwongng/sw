/*****************************************************************************
 * storage_common_defines.h: Common storage defintions across P4+ and assembly
 *****************************************************************************/
 
#ifndef STORAGE_COMMON_DEFINES_H
#define STORAGE_COMMON_DEFINES_H

// Default table load size (used in cases where table entry size > 64 bytes)
// The sizes are expressed as a power of 2, hence using 6 here
#define STORAGE_DEFAULT_TBL_LOAD_SIZE	6

// Offset in bytes of the start of actual NVME command in the NVME backend command
#define NVME_BE_NVME_CMD_OFFSET		64

// Maximum number of commands outstanding to an SSD
#define SSD_MAX_CMDS			64

// Keep the header and entry size of the SSD's list of saved commands to be 
// aligned at 64 bytes for easy table read access
// Header: 64-bit bitmap 
// Entry:  Full R2N WQE (r2n_wqe_t) is stored here
#define SSD_CMDS_HEADER_SIZE		64
#define SSD_CMDS_ENTRY_SIZE		64

// Offset in memory to store the command handle in the SSD's list of saved commands
#define SSD_CMDS_ENTRY_OFFSET(bm_pos)	(SSD_CMDS_HEADER_SIZE + (bm_pos * SSD_CMDS_ENTRY_SIZE))

// Capri doorbell address/data formation values
#define DOORBELL_UPDATE_NONE		0
#define DOORBELL_UPDATE_C_NDX		1
#define DOORBELL_UPDATE_P_NDX		2
#define DOORBELL_UPDATE_P_NDX_INCR	3
#define DOORBELL_SCHED_WR_NONE		0
#define DOORBELL_SCHED_WR_EVAL		1
#define DOORBELL_SCHED_WR_RESET		2
#define DOORBELL_SCHED_WR_SET		3

#define DOORBELL_ADDR_QTYPE_SHIFT       3
#define DOORBELL_ADDR_LIF_SHIFT         6
#define DOORBELL_ADDR_SCHED_WR_SHIFT    17
#define DOORBELL_ADDR_UPD_SHIFT         19
#define DOORBELL_ADDR_WA_LOCAL_BASE     0x68800000

#define DOORBELL_DATA_RING_SHIFT	16
#define DOORBELL_DATA_QID_SHIFT		24
#define DOORBELL_DATA_PID_SHIFT		48

// Each queue state can be 64 bytes at most (from d-vector in Capri)
// The sizes are expressed as a power of 2, hence using 6 here
#define Q_STATE_SIZE     	6      

// Offsets (in bytes) of fields within a structure
#define Q_STATE_C_NDX_OFFSET	10

// PRP and LB size calculation assists
#define PRP_SIZE_MASK		0xFFF
#define PRP_SIZE_SUB		0x1000
#define LB_SIZE_SHIFT		12

// Max h/w assist size for PRP list download
#define MAX_ASSIST_SIZE		65536

// Max number of additional PRPs in assist
#define NVME_MAX_XTRA_PRPS	16

// NVME write and read command opcodes (from NVME spec)
#define NVME_WRITE_CMD_OPCODE			1
#define NVME_READ_CMD_OPCODE			2

#define NVME_BE_PRIORITY_HI			0
#define NVME_BE_PRIORITY_MED			1
#define NVME_BE_PRIORITY_LO			2
#define NVME_BE_MAX_PRIORITIES			3

#endif     // STORAGE_COMMON_DEFINES_H
