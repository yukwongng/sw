#include <stdio.h>
#include <strings.h>
#include <byteswap.h>

#include "dol/test/storage/hal_if.hpp"
#include "dol/test/storage/qstate_if.hpp"
#include "dol/test/storage/nvme.hpp"
#include "dol/test/storage/queues.hpp"
#include "dol/test/storage/ssd.hpp"
#include "nic/utils/host_mem/c_if.h"
#include "nic/model_sim/include/lib_model_client.h"

#define	NUM_TO_VAL(num)		(1 << (num))

const static uint32_t	kNvmeNumSQs		 = 3;
const static uint32_t	kNvmeNumCQs		 = 3;
const static uint32_t	kPvmNumSQs		 = 7;
const static uint32_t	kPvmNumCQs		 = 5;
const static uint32_t	kPvmNumHQs		 = 0; // 2^0 => 1 queue
const static uint32_t	kPvmNumEQs		 = 1;
const static uint32_t	kPvmNumNvmeSQs		 = 1;
const static uint32_t	kPvmNumR2nSQs		 = 0; // 2^0 => 1 queue
const static uint32_t	kPvmNumNvmeBeSQs	 = 4;
const static uint32_t	kPvmNumSsdSQs	 	 = 4;
const static uint32_t	kPvmNumSeqPdmaSQs	 = 3;
const static uint32_t	kPvmNumSeqR2nSQs	 = 3;
const static uint32_t	kPvmNumNvmeCQs		 = 1;
const static uint32_t	kPvmNumR2nCQs		 = 0; // 2^0 => 1 queue
const static uint32_t	kPvmNumNvmeBeCQs	 = 4;
const static uint32_t	kDefaultEntrySize	 = 6; // Default is 64 bytes
const static uint32_t	kPvmNvmeSQEntrySize	 = 7; // PVM SQ is 128 bytes (NVME command + PVM header)
const static uint32_t	kNvmeCQEntrySize	 = 4; // NVME CQ is 16 bytes
const static uint32_t	kNvmeNumEntries		 = 6;
const static uint32_t	kPvmNumEntries		 = 6;
const static uint32_t	kPvmNumSeqEntries	 = 1;
const static char	*kNvmeSqHandler		 = "storage_tx_nvme_sq_handler.bin";
const static char	*kPvmCqHandler		 = "storage_tx_pvm_cq_handler.bin";
const static char	*kR2nSqHandler		 = "storage_tx_r2n_sq_handler.bin";
const static char	*kNvmeBeSqHandler	 = "storage_tx_nvme_be_sq_handler.bin";
const static char	*kNvmeBeCqHandler	 = "storage_tx_nvme_be_cq_handler.bin";
const static char	*kSeqPdmaSqHandler	 = "storage_tx_seq_pdma_entry_handler.bin";
const static char	*kSeqR2nSqHandler	 = "storage_tx_seq_r2n_entry_handler.bin";
const static uint32_t	kDefaultTotalRings	 = 1;
const static uint32_t	kDefaultHostRings	 = 1;
const static uint32_t	kDefaultNoHostRings	 = 0;
const static uint32_t	kNvmeBeTotalRings	 = 3;


const static uint32_t	kDbAddrHost		 = 0x400000;
const static uint32_t	kDbAddrCapri		 = 0x68800000;
const static uint32_t	kDbAddrUpdate		 = 0xB;
const static uint32_t	kDbQidShift		 = 24;
const static uint32_t	kDbRingShift		 = 16;
const static uint32_t	kDbUpdateShift		 = 17;
const static uint32_t	kDbLifShift		 = 6;
const static uint32_t	kDbTypeShift		 = 3;

// Special SSD marked for running E2e traffic with queues from 
// NvmeSsd class
const static int	kE2eSsdhandle		 = 2;


// Offset of the Cndx in the Qstate
const static uint32_t	kQstateCndxOffset	 = 10;

namespace queues {

// Special SSD marked for running E2e traffic with queues from 
// NvmeSsd class
std::unique_ptr<storage_test::NvmeSsd> nvme_e2e_ssd;

uint64_t ssd_cndx_addr[NUM_TO_VAL(kPvmNumNvmeBeSQs)];

typedef struct queues_ {
  void *vaddr;
  uint64_t paddr;
  uint16_t index; // p_ndx for tx queue, c_ndx for rx queue
  uint16_t entry_size;
  uint16_t num_entries;
} queues_t;

// NVME Submission, Completion queues
queues_t nvme_sqs[NUM_TO_VAL(kNvmeNumSQs)];
queues_t nvme_cqs[NUM_TO_VAL(kNvmeNumCQs)];

// PVM Submission, Completion, Host(for R2N), Error queues
queues_t pvm_sqs[NUM_TO_VAL(kPvmNumSQs)];
queues_t pvm_cqs[NUM_TO_VAL(kPvmNumCQs)];
queues_t pvm_rqs[NUM_TO_VAL(kPvmNumHQs)];
queues_t pvm_eqs[NUM_TO_VAL(kPvmNumEQs)];

uint64_t nvme_lif, pvm_lif;

uint64_t storage_hbm_addr;
uint32_t storage_hbm_size;

void *pndx_data_va;
uint64_t pndx_data_pa;

void lif_params_init(hal_if::lif_params_t *params, uint16_t type,
                     uint16_t num_entries, uint16_t num_queues) {
  params->type[type].valid = true; 
  params->type[type].queue_size = num_entries; // you get 2^num_entries
  params->type[type].num_queues = num_queues;  // you get 2^num_queues 
}

int nvme_e2e_ssd_sq_init(queues_t *queue, uint16_t num_entries, uint16_t entry_size) {
  // NOTE: Assumes that nvme_e2e SSD SQ is initated with 64 entries x 64 bytes size
  // TODO: Fix to make this consistent or assert on the values
  storage_test::SsdWorkingParams params;
  nvme_e2e_ssd->GetWorkingParams(&params);
  queue->vaddr = params.subq_va;
  queue->paddr = params.subq_pa;
  queue->index = 0;
  queue->entry_size = entry_size;
  queue->num_entries = num_entries;
  return 0;
}

int nvme_e2e_ssd_cq_init(queues_t *queue, uint16_t num_entries, uint16_t entry_size) {
  // NOTE: Assumes that nvme_e2e SSD SQ is initated with 64 entries x 16 bytes size
  // TODO: Fix to make this consistent or assert on the values
  storage_test::SsdWorkingParams params;
  nvme_e2e_ssd->GetWorkingParams(&params);
  queue->vaddr = params.compq_va;
  queue->paddr = params.compq_pa;
  queue->index = 0;
  queue->entry_size = entry_size;
  queue->num_entries = num_entries;
  return 0;
}

void nvme_e2e_ssd_db_init(uint64_t db_addr, uint64_t db_data) {
  nvme_e2e_ssd->EnableInterrupt(db_addr, db_data);
}

int queue_init(queues_t *queue, uint16_t num_entries, uint16_t entry_size) {
  queue->vaddr = (uint8_t *)alloc_host_mem(num_entries * entry_size);
  if (queue->vaddr == nullptr) {
    printf("Unable to allocate host memory for queue\n");
    return -1;
  }
  queue->paddr = host_mem_v2p(queue->vaddr);
  queue->index = 0;
  queue->entry_size = entry_size;
  queue->num_entries = num_entries;
  return 0;
}

void *queue_consume_entry(queues_t *queue, uint16_t *index) {
  if (!queue->vaddr || !index) return nullptr;
  uint16_t curr_index = queue->index;
  queue->index = ((queue->index + 1)  % queue->num_entries);
  *index = queue->index;
  return (void *) ((uint8_t *) queue->vaddr + (curr_index * queue->entry_size));
}



int seq_queue_setup(queues_t *q_ptr, uint32_t qid, char *pgm_bin, 
                    uint16_t total_rings, uint16_t host_rings) {

  // Initialize the queue in the DOL enviroment
  if (queue_init(q_ptr, NUM_TO_VAL(kPvmNumSeqEntries),
                 NUM_TO_VAL(kDefaultEntrySize)) < 0) {
    printf("Unable to allocate host memory for PVM Seq SQ %d\n", qid);
    return -1;
  }
  printf("Initialized PVM Seq SQ %d \n", qid);

  // Setup the queue state in Capri:
  if (qstate_if::setup_q_state(pvm_lif, SQ_TYPE, qid, pgm_bin, 
                               total_rings, host_rings, 
                               kPvmNumSeqEntries, q_ptr->paddr,
                               kDefaultEntrySize, false, 0, 0,
                               0, 0, 0, storage_hbm_addr, 0, 0, 0) < 0) {
    printf("Failed to setup PVM Seq SQ %d state \n", qid);
    return -1;
  }
  return 0;
}

int queues_setup() {
  // Allocatge HBM address for storage
  nvme_e2e_ssd.reset(new storage_test::NvmeSsd());

  // Allocatge HBM address for storage
  if (hal_if::alloc_hbm_address(&storage_hbm_addr, &storage_hbm_size) < 0) {
    printf("can't allocate HBM address for storage \n");
    return -1;
  }
  printf("Storage HBM address %lx size %d KB\n", storage_hbm_addr, storage_hbm_size);

  // Create the just one global 32 bit pndx counter for all SSDs, NVME driver, 
  // R2N/PVM SQ etc.  // This is needed as a dummy for the P4+ program to write 
  // the pndx to a valid host address. The SSD emulation layer implements this 
  // for the E2E cases.
  pndx_data_va = alloc_host_mem(sizeof(uint32_t));
  if (pndx_data_va == nullptr) {
    printf("Unable to allocate host memory for queue\n");
    return -1;
  }
  pndx_data_pa = host_mem_v2p(pndx_data_va);

  // Create NVME and PVM LIFs
  hal_if::lif_params_t nvme_lif_params, pvm_lif_params;

  bzero(&nvme_lif_params, sizeof(nvme_lif_params));
  lif_params_init(&nvme_lif_params, SQ_TYPE, kNvmeNumEntries, kNvmeNumSQs);
  lif_params_init(&nvme_lif_params, CQ_TYPE, kNvmeNumEntries, kNvmeNumCQs);

  if (hal_if::create_lif(&nvme_lif_params, &nvme_lif) < 0) {
    printf("can't create nvme lif \n");
    return -1;
  }
  printf("NVME LIF created\n");

  bzero(&pvm_lif_params, sizeof(pvm_lif_params));
  lif_params_init(&pvm_lif_params, SQ_TYPE, kPvmNumEntries, kPvmNumSQs);
  lif_params_init(&pvm_lif_params, CQ_TYPE, kPvmNumEntries, kPvmNumCQs);
  lif_params_init(&pvm_lif_params, HQ_TYPE, kPvmNumEntries, kPvmNumHQs);
  lif_params_init(&pvm_lif_params, EQ_TYPE, kPvmNumEntries, kPvmNumEQs);

  if (hal_if::create_lif(&pvm_lif_params, &pvm_lif) < 0) {
    printf("can't create pvm lif \n");
    return -1;
  }
  printf("PVM LIF created\n");

  int i, j;

  // Initialize NVME SQs
  for (i = 0; i < (int) NUM_TO_VAL(kNvmeNumSQs); i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&nvme_sqs[i], NUM_TO_VAL(kNvmeNumEntries),
                   NUM_TO_VAL(kDefaultEntrySize)) < 0) {
      printf("Unable to allocate host memory for NVME SQ %d\n", i);
      return -1;
    }
    printf("Initialized NVME SQ %d \n", i);

    // Setup the queue state in Capri
    if (qstate_if::setup_q_state(nvme_lif, SQ_TYPE, i, (char *) kNvmeSqHandler, 
                                 kDefaultTotalRings, kDefaultHostRings, 
                                 kNvmeNumEntries, nvme_sqs[i].paddr,
                                 kDefaultEntrySize, true, pvm_lif, SQ_TYPE, 
                                 (i%2), 0, 0, storage_hbm_addr, 0, 0, 0) < 0) {
      printf("Failed to setup NVME SQ %d state \n", i);
      return -1;
    }
  }


  // Initialize NVME CQs
  for (i = 0; i < (int) NUM_TO_VAL(kNvmeNumCQs); i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&nvme_cqs[i], NUM_TO_VAL(kNvmeNumEntries),
                   NUM_TO_VAL(kNvmeCQEntrySize)) < 0) {
      printf("Unable to allocate host memory for NVME CQ %d\n", i);
      return -1;
    }
    printf("Initialized NVME CQ %d \n", i);

    // Setup the queue state in Capri:
    // 1. no dst queues for these as these are in host
    // 2. no program address for these as these are in host
    // 3. The push address is that of the SSD's SQ's PI 
    if (qstate_if::setup_pci_q_state(nvme_lif, CQ_TYPE, i,
                                     kDefaultTotalRings, kDefaultHostRings, 
                                     kNvmeNumEntries, nvme_cqs[i].paddr,
                                     kNvmeCQEntrySize, pndx_data_pa, 
                                     0, 0, 0) < 0) {
      printf("Failed to setup NVME CQ %d state \n", i);
      return -1;
    }
  }

  // Initialize PVM SQs for processing commands from NVME VF only
  // Note: i is overall index across PVM SQs, j iterates the loop
  for (j = 0, i = 0; j < (int) NUM_TO_VAL(kPvmNumNvmeSQs); j++, i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&pvm_sqs[i], NUM_TO_VAL(kPvmNumEntries),
                   NUM_TO_VAL(kPvmNvmeSQEntrySize)) < 0) {
      printf("Unable to allocate host memory for PVM NVME SQ %d\n", i);
      return -1;
    }
    printf("Initialized PVM NVME SQ %d \n", i);

    // Setup the queue state in Capri:
    // 1. no dst queues for these as these are in host
    // 2. no program address for these as these are in host
    // 3. The push address is that of the SSD's SQ's PI 
    if (qstate_if::setup_pci_q_state(pvm_lif, SQ_TYPE, i,
                                     kDefaultTotalRings, kDefaultHostRings, 
                                     kPvmNumEntries, pvm_sqs[i].paddr,
                                     kPvmNvmeSQEntrySize, pndx_data_pa, 
                                     0, 0, 0) < 0) {
      printf("Failed to setup PVM NVME SQ %d state \n", i);
      return -1;
    }
  }

  // Initialize PVM SQs for processing R2N commands 
  // Note: Not incrementing nvme_be_q in for loop as the SSD handle is added to this
  //       by R2N module in datapath.
  uint32_t nvme_be_q = i + NUM_TO_VAL(kPvmNumR2nSQs);
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumR2nSQs); j++, i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&pvm_sqs[i], NUM_TO_VAL(kPvmNumEntries),
                   NUM_TO_VAL(kDefaultEntrySize)) < 0) {
      printf("Unable to allocate host memory for PVM R2N SQ %d\n", i);
      return -1;
    }
    printf("Initialized PVM R2N SQ %d \n", i);

    // Setup the queue state in Capri:
    if (qstate_if::setup_q_state(pvm_lif, SQ_TYPE, i, (char *) kR2nSqHandler, 
                                 kDefaultTotalRings, kDefaultHostRings, 
                                 kPvmNumEntries, pvm_sqs[i].paddr,
                                 kDefaultEntrySize, true, pvm_lif, SQ_TYPE,
                                 nvme_be_q, 0, 0, storage_hbm_addr, 
                                 kPvmNumNvmeBeSQs, kDefaultEntrySize, 0) < 0) {
      printf("Failed to setup PVM R2N SQ %d state \n", i);
      return -1;
    }
  }

  // Initialize PVM SQs for processing NVME backend commands 
  // Note: Incrementing ssd_q in the for loop as the NVME backend corresponds 1:1 
  //       with the SSD
  uint32_t ssd_q = i + NUM_TO_VAL(kPvmNumNvmeBeSQs);
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumNvmeBeSQs); j++, i++, ssd_q++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&pvm_sqs[i], NUM_TO_VAL(kPvmNumEntries),
                   NUM_TO_VAL(kDefaultEntrySize)) < 0) {
      printf("Unable to allocate host memory for PVM NVME backend SQ %d\n", i);
      return -1;
    }
    printf("Initialized PVM NVME backend SQ %d \n", i);

    // Setup the queue state in Capri: 
    if (qstate_if::setup_pri_q_state(pvm_lif, SQ_TYPE, i, (char *) kNvmeBeSqHandler, 
                                     kNvmeBeTotalRings, kDefaultNoHostRings, 
                                     kPvmNumEntries, pvm_sqs[i].paddr,
                                     kDefaultEntrySize, true, pvm_lif, SQ_TYPE,
                                     ssd_q, 0, 0, storage_hbm_addr) < 0) {
      printf("Failed to setup PVM NVME backend SQ %d state \n", i);
      return -1;
    }
  }

  // Initialize PVM SQs for processing SSD commands 
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumSsdSQs); j++, i++) {
    // Physical address to update pindex
    uint64_t pi_pa;

    // For the special E2E SSD, use NvmeSsd class to initialize the queue
    if  (j == kE2eSsdhandle) {
      // Initialize the queue in the DOL enviroment
      if (nvme_e2e_ssd_sq_init(&pvm_sqs[i], NUM_TO_VAL(kPvmNumEntries),
                               NUM_TO_VAL(kDefaultEntrySize)) < 0) {
        printf("Unable to allocate host memory for PVM E2E SSD SQ %d\n", i);
        return -1;
      }
      printf("Initialized PVM SSD E2E SQ %d \n", i);

      // Get the PI's physical address from the SSD emulation layer
      storage_test::SsdWorkingParams params;
      nvme_e2e_ssd->GetWorkingParams(&params);
      pi_pa = params.subq_pi_pa;

      // Initialize the doorbell of the CQ
      uint64_t db_addr;
      uint64_t db_data;
      uint32_t qid = (int) NUM_TO_VAL(kPvmNumNvmeCQs) + (int) NUM_TO_VAL(kPvmNumR2nCQs) + kE2eSsdhandle;
      queues::get_doorbell(pvm_lif, CQ_TYPE, qid, 0, 0, &db_addr, &db_data);
      nvme_e2e_ssd_db_init(db_addr, db_data);
      printf("Initialized backend doorbell for SSD %d \n", i);

    } else {
      // Initialize the queue in the DOL enviroment
      if (queue_init(&pvm_sqs[i], NUM_TO_VAL(kPvmNumEntries),
                     NUM_TO_VAL(kDefaultEntrySize)) < 0) {
        printf("Unable to allocate host memory for PVM SSD SQ %d\n", i);
        return -1;
      }
      printf("Initialized PVM SSD SQ %d \n", i);

      // Store the default value of the PI's physical address
      pi_pa = pndx_data_pa;
    }

    // Setup the queue state in Capri:
    // 1. no dst queues for these as these are in host
    // 2. no program address for these as these are in host
    // 3. The push address is that of the SSD's SQ's PI 
    if (qstate_if::setup_pci_q_state(pvm_lif, SQ_TYPE, i,
                                     kDefaultTotalRings, kDefaultNoHostRings, 
                                     kPvmNumEntries, pvm_sqs[i].paddr,
                                     kDefaultEntrySize, pi_pa, 
                                     0, 0, 0) < 0) {
      printf("Failed to setup PVM SSD SQ %d state \n", i);
      return -1;
    }

    // Store the address of the CI for this SSD
    uint64_t qaddr;
    if (qstate_if::get_qstate_addr(pvm_lif, SQ_TYPE, i, &qaddr) < 0) {
        printf("Failed to get PVM SSD SQ %d address \n", i);
        return -1;
    }
    ssd_cndx_addr[j] = qaddr + kQstateCndxOffset;
    printf("SSD %d qaddr %lx cndx_addr %lx \n", j, qaddr, ssd_cndx_addr[j]);
  }

  // Initialize PVM SQs for processing Sequencer commands for PDMA
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumSeqPdmaSQs); j++, i++) {
    if (seq_queue_setup(&pvm_sqs[i], i, (char *) kSeqPdmaSqHandler,
                        kDefaultTotalRings, kDefaultHostRings) < 0) {
      printf("Failed to setup PVM Seq PDMA queue %d \n", i);
      return -1;
    }
    printf("Setup PVM Seq PDMA queue %d \n", i);
  }

  // Initialize PVM SQs for processing Sequencer commands for R2N
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumSeqR2nSQs); j++, i++) {
    if (seq_queue_setup(&pvm_sqs[i], i, (char *) kSeqR2nSqHandler,
                        kDefaultTotalRings, kDefaultNoHostRings) < 0) {
      printf("Failed to setup PVM Seq R2n queue %d \n", i);
      return -1;
    }
    printf("Setup PVM Seq R2n queue %d \n", i);
  }

  // Initialize PVM SQs for processing R2N commands from the Sequencer.
  // This is strictly to avoid queue sharing between PVM and P4+ code.
  // Note: This is different from the Sequencer R2N entry handler queue 
  // created above.
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumR2nSQs); j++, i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&pvm_sqs[i], NUM_TO_VAL(kPvmNumEntries),
                   NUM_TO_VAL(kDefaultEntrySize)) < 0) {
      printf("Unable to allocate host memory for PVM R2N SQ for Seq %d\n", i);
      return -1;
    }
    printf("Initialized PVM R2N SQ %d for Seq \n", i);

    // Setup the queue state in Capri:
    if (qstate_if::setup_q_state(pvm_lif, SQ_TYPE, i, (char *) kR2nSqHandler, 
                                 kDefaultTotalRings, kDefaultHostRings, 
                                 kPvmNumEntries, pvm_sqs[i].paddr,
                                 kDefaultEntrySize, true, pvm_lif, SQ_TYPE,
                                 nvme_be_q, 0, 0, storage_hbm_addr, 
                                 kPvmNumNvmeBeSQs, kDefaultEntrySize, 0) < 0) {
      printf("Failed to setup PVM R2N SQ %d state for Seq \n", i);
      return -1;
    }
  }

  // Initialize PVM CQs for processing commands from NVME VF only
  // Note: i is overall index across PVM CQs, j iterates the loop
  for (j = 0, i = 0; j < (int) NUM_TO_VAL(kPvmNumNvmeCQs); j++, i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&pvm_cqs[i], NUM_TO_VAL(kPvmNumEntries),
                   NUM_TO_VAL(kDefaultEntrySize)) < 0) {
      printf("Unable to allocate host memory for PVM NVME CQ %d\n", i);
      return -1;
    }
    printf("Initialized PVM NVME CQ %d \n", i);

    // Setup the queue state in Capri:
    // 1. no dst queues for these as these are in host
    // 2. no program address for these as these are in host
    if (qstate_if::setup_q_state(pvm_lif, CQ_TYPE, i, (char *) kPvmCqHandler,
                                 kDefaultTotalRings, kDefaultHostRings, 
                                 kPvmNumEntries, pvm_cqs[i].paddr,
                                 kDefaultEntrySize, false, 0, 0,
                                 0, 0, 0, storage_hbm_addr, 0, 0, 0) < 0) {
      printf("Failed to setup PVM NVME CQ %d state \n", i);
      return -1;
    }
  }

  // Initialize PVM CQs for processing R2N commands 
  uint32_t pvm_r2n_cq = i;
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumR2nCQs); j++, i++) {
    // Initialize the queue in the DOL enviroment
    if (queue_init(&pvm_cqs[i], NUM_TO_VAL(kPvmNumEntries),
                   NUM_TO_VAL(kDefaultEntrySize)) < 0) {
      printf("Unable to allocate host memory for PVM R2N CQ %d\n", i);
      return -1;
    }
    printf("Initialized PVM R2N CQ %d \n", i);

    // Setup the queue state in Capri:
    // 1. no dst queues for these as these are in host
    // 2. no program address for these as these are in host
    // 3. The push address is that of the SSD's SQ's PI 
    if (qstate_if::setup_pci_q_state(pvm_lif, CQ_TYPE, i,
                                     kDefaultTotalRings, kDefaultNoHostRings, 
                                     kPvmNumEntries, pvm_cqs[i].paddr,
                                     kDefaultEntrySize, pndx_data_pa, 
                                     0, 0, 0) < 0) {
      printf("Failed to setup PVM NVME SQ %d state \n", i);
      return -1;
    }
  }

  // Initialize PVM CQs for processing NVME backend commands 
  // Note: Incrementing ssd_q in the for loop as the NVME backend corresponds 1:1 
  //       with the SSD
  for (j = 0; j < (int) NUM_TO_VAL(kPvmNumNvmeBeCQs); j++, i++) {
    // For the special E2E SSD, use NvmeSsd class to initialize the queue
    if  (j == kE2eSsdhandle) {
      // Initialize the queue in the DOL enviroment
      if (nvme_e2e_ssd_cq_init(&pvm_cqs[i], NUM_TO_VAL(kPvmNumEntries),
                               NUM_TO_VAL(kNvmeCQEntrySize)) < 0) {
        printf("Unable to allocate host memory for PVM E2E SSD CQ %d\n", i);
        return -1;
      }
      printf("Initialized PVM NVME backend E2E CQ %d \n", i);
    } else {
      // Initialize the queue in the DOL enviroment
      if (queue_init(&pvm_cqs[i], NUM_TO_VAL(kPvmNumEntries),
                     NUM_TO_VAL(kNvmeCQEntrySize)) < 0) {
        printf("Unable to allocate host memory for PVM NVME backend CQ %d\n", i);
        return -1;
      }
      printf("Initialized PVM NVME backend CQ %d \n", i);
    }

    // Setup the queue state in Capri: 
    if (qstate_if::setup_q_state(pvm_lif, CQ_TYPE, i, (char *) kNvmeBeCqHandler, 
                                 kDefaultTotalRings, kDefaultHostRings, 
                                 kPvmNumEntries, pvm_cqs[i].paddr,
                                 kNvmeCQEntrySize, true, pvm_lif, CQ_TYPE,
                                 pvm_r2n_cq, 0, 0, storage_hbm_addr, 0, 0,
                                 ssd_cndx_addr[j]) < 0) {
      printf("Failed to setup PVM NVME backend CQ %d state \n", i);
      return -1;
    }
  }

  return 0;
}

void *nvme_sq_consume_entry(uint16_t qid, uint16_t *index) {
  if (qid >= NUM_TO_VAL(kNvmeNumSQs)) return nullptr;
  return queue_consume_entry(&nvme_sqs[qid], index);
}

void *pvm_sq_consume_entry(uint16_t qid, uint16_t *index) {
  if (qid >= NUM_TO_VAL(kPvmNumSQs)) return nullptr;
  return queue_consume_entry(&pvm_sqs[qid], index);
}

void *nvme_cq_consume_entry(uint16_t qid, uint16_t *index) {
  if (qid >= NUM_TO_VAL(kNvmeNumCQs)) return nullptr;
  return queue_consume_entry(&nvme_cqs[qid], index);
}

void *pvm_cq_consume_entry(uint16_t qid, uint16_t *index) {
  if (qid >= NUM_TO_VAL(kPvmNumCQs)) return nullptr;
  return queue_consume_entry(&pvm_cqs[qid], index);
}

uint16_t get_nvme_lif() {
  return (uint16_t) nvme_lif;
}

uint16_t get_pvm_lif() {
  return (uint16_t) pvm_lif;
}

void ring_nvme_e2e_ssd() {
  printf("*********** RINGING SSD DOORBELL for testing *********** \n");
  storage_test::SsdWorkingParams params;
  nvme_e2e_ssd->GetWorkingParams(&params);
  *((uint32_t *) params.subq_pi_va) =   (*((uint32_t *) params.subq_pi_va)) + 1;
}

void get_doorbell(uint16_t lif, uint8_t qtype, uint32_t qid, 
                  uint8_t ring, uint16_t index, 
                  uint64_t *db_addr, uint64_t *db_data) {

  if (!db_addr || !db_data) return;

  *db_data = (qid << kDbQidShift) | (ring << kDbRingShift) | bswap_16(index);
  *db_addr = kDbAddrHost |  (kDbAddrUpdate << kDbUpdateShift) | 
             (lif << kDbLifShift) | (qtype << kDbTypeShift);
}

void get_capri_doorbell(uint16_t lif, uint8_t qtype, uint32_t qid, 
                        uint8_t ring, uint16_t index, 
                        uint64_t *db_addr, uint64_t *db_data) {

  if (!db_addr || !db_data) return;

  *db_data = (qid << kDbQidShift) | (ring << kDbRingShift) | bswap_16(index);
  *db_addr = kDbAddrCapri |  (kDbAddrUpdate << kDbUpdateShift) | 
             (lif << kDbLifShift) | (qtype << kDbTypeShift);
}

uint64_t get_storage_hbm_addr() {
  return storage_hbm_addr;
}

void queues_shutdown() {
  nvme_e2e_ssd.reset();
}

}  // namespace queues
