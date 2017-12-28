#include <asm/byteorder.h>
#include <stdio.h>
#include <strings.h>
#include <string>
#include <execinfo.h>
#include <signal.h>
#include <stdlib.h>
#include <unistd.h>
#include <vector>
#include <iostream>

#include "gflags/gflags.h"

#include "dol/test/storage/tests.hpp"
#include "dol/test/storage/rdma.hpp"
#include "dol/test/storage/compression_test.hpp"

namespace queues {
void queues_shutdown();
}

DEFINE_uint64(hal_port, 50054, "TCP port of the HAL's gRPC server");
DEFINE_string(test_group, "", "Test group to run");
DEFINE_uint64(poll_interval, 60, "Polling interval in seconds");

bool run_unit_tests = false;
bool run_nvme_tests = false;
bool run_nvme_be_tests = false;
bool run_local_e2e_tests = false;
bool run_comp_tests = false;
bool run_xts_tests = false;
bool run_rdma_tests = false;

std::vector<tests::TestEntry> test_suite;

std::vector<tests::TestEntry> unit_tests = {
  {&tests::test_run_nvme_pvm_admin_cmd, "NVME->PVM Admin Cmd", false},
};

std::vector<tests::TestEntry> nvme_tests = {
  {&tests::test_run_nvme_pvm_read_cmd, "NVME->PVM Read Cmd", false},
  {&tests::test_run_nvme_pvm_write_cmd, "NVME->PVM Write Cmd", false},
  {&tests::test_run_nvme_pvm_hashing1, "NVME->PVM Hashing 1", false},
  {&tests::test_run_nvme_pvm_hashing2, "NVME->PVM Hashing 2", false},
  {&tests::test_run_pvm_nvme_admin_status, "PVM->NVME Admin Status", false},
  {&tests::test_run_pvm_nvme_read_status, "PVM->NVME Read Status", false},
  {&tests::test_run_pvm_nvme_write_status, "PVM->NVME Write Status", false},
};

std::vector<tests::TestEntry> nvme_be_tests = {
  {&tests::test_run_r2n_read_cmd, "R2N -> SSD Read Cmd", false},
  {&tests::test_run_r2n_write_cmd, "R2N -> SSD Write Cmd", false},
  {&tests::test_run_r2n_ssd_pri1, "R2N -> SSD Pri Cmd 1", false},
  {&tests::test_run_r2n_ssd_pri2, "R2N -> SSD Pri Cmd 2", false},
  {&tests::test_run_r2n_ssd_pri3, "R2N -> SSD Pri Cmd 3", false},
  {&tests::test_run_r2n_ssd_pri4, "R2N -> SSD Pri Cmd 4", false},
  {&tests::test_run_nvme_be_wrr1, "NVME Backend WRR 1", false},
  {&tests::test_run_nvme_be_wrr2, "NVME Backend WRR 2", false},
  {&tests::test_run_nvme_be_wrr3, "NVME Backend WRR 3", false},
  {&tests::test_run_nvme_be_wrr4, "NVME Backend WRR 4", false},
  {&tests::test_run_nvme_be_wrr5, "NVME Backend WRR 5", false},
  {&tests::test_run_nvme_be_wrr6, "NVME Backend WRR 6", false},
  {&tests::test_run_nvme_read_comp1, "PVM Local Read Comp 1", false},
  {&tests::test_run_nvme_write_comp1, "PVM Local Write Comp 1", false},
  {&tests::test_run_nvme_read_comp2, "PVM Local Read Comp 2", false},
  {&tests::test_run_nvme_write_comp2, "PVM Local Write Comp 2", false},
  {&tests::test_run_nvme_read_comp3, "PVM Local Read Comp 3", false},
  {&tests::test_run_nvme_write_comp3, "PVM Local Write Comp 3", false},
};

std::vector<tests::TestEntry> local_e2e_tests = {
  {&tests::test_run_nvme_local_e2e1, "NVME Local Tgt E2E 1", false},
  {&tests::test_run_nvme_local_e2e2, "NVME Local Tgt E2E 2", false},
  {&tests::test_run_nvme_local_e2e3, "NVME Local Tgt E2E 3", false},
  {&tests::test_run_seq_write1, "Seq Local Tgt Write 1", false},
  {&tests::test_run_seq_write2, "Seq Local Tgt Write 2", false},
  {&tests::test_run_seq_write3, "Seq Local Tgt Write 3", false},
  {&tests::test_run_seq_write4, "Seq Local Tgt Write 4", false},
  {&tests::test_run_seq_read1, "Seq Local Tgt Read 1", false},
  {&tests::test_run_seq_read2, "Seq Local Tgt Read 2", false},
  {&tests::test_run_seq_read3, "Seq Local Tgt Read 3", false},
  {&tests::test_run_seq_read4, "Seq Local Tgt Read 4", false},
  {&tests::test_run_seq_e2e1, "Seq Local Tgt E2E 1", false},
  {&tests::test_run_seq_e2e2, "Seq Local Tgt E2E 2", false},
  {&tests::test_run_seq_e2e3, "Seq Local Tgt E2E 3", false},
  {&tests::test_run_seq_e2e4, "Seq Local Tgt E2E 4", false},
  {&tests::test_seq_e2e_xts_r2n1, "PDMA->XTS->R2N", false},
};

std::vector<tests::TestEntry> comp_tests = {
  {&tests::compress_host_flat, "Compress Host->Host flat buf", false},
  {&tests::compress_hbm_flat, "Compress HBM->HBM flat buf", false},
  {&tests::compress_host_to_hbm_flat, "Compress Host->HBM flat buf", false},
  {&tests::compress_hbm_to_host_flat, "Compress HBM->Host flat buf", false},
  {&tests::compress_host_sgl, "Compress Host->Host sgl buf", false},
  {&tests::compress_hbm_sgl, "Compress HBM->HBM sgl buf", false},
  {&tests::compress_host_nested_sgl, "Compress Host->Host nested sgl buf", false},
  {&tests::compress_hbm_nested_sgl, "Compress HBM->HBM nested sgl buf", false},
  {&tests::compress_nested_sgl_in_hbm, "Compress Nested sgl buf in HBM", false},
  {&tests::compress_return_through_hbm, "Compress Status/dest-buf in HBM", false},
  {&tests::compress_adler_sha256, "Compress with Adler32 and SHA256", false},
  {&tests::compress_crc_sha512, "Compress with CRC32 and SHA512", false},
  {&tests::compress_only_sha512, "Compress with only SHA512", false},
  {&tests::compress_doorbell_odata, "Compress with DMA end writes", false},
  {&tests::compress_max_features, "Compress with multiple features", false},
  {&tests::compress_output_through_sequencer, "Compress with out through seq", false},
  {&tests::decompress_host_flat, "Decompress Host->Host flat buf", false},
  {&tests::decompress_hbm_flat, "Decompress HBM->HBM flat buf", false},
  {&tests::decompress_host_to_hbm_flat, "Decompress Host->HBM flat buf", false},
  {&tests::decompress_hbm_to_host_flat, "Decompress HBM->Host flat buf", false},
  {&tests::decompress_host_sgl, "Decompress Host->Host sgl buf", false},
  {&tests::decompress_hbm_sgl, "Decompress HBM->HBM sgl buf", false},
  {&tests::decompress_host_nested_sgl, "Decompress Host->Host nested sgl buf", false},
  {&tests::decompress_hbm_nested_sgl, "Decompress HBM->HBM nested sgl buf", false},
  {&tests::decompress_nested_sgl_in_hbm, "Decompress Nested sgl buf in HBM", false},
  {&tests::decompress_return_through_hbm, "Decompress Status/dest-buf in HBM", false},
  {&tests::decompress_adler, "Decompress with Adler32", false},
  {&tests::decompress_crc, "Decompress with CRC32", false},
  {&tests::decompress_doorbell_odata, "Decompress with DMA end writes", false},
};

std::vector<tests::TestEntry> rdma_tests = {
  {&tests::test_run_rdma_e2e_write, "E2E write over RDMA", false},
  {&tests::test_run_rdma_e2e_read, "E2E read over RDMA", false},
  {&tests::test_run_rdma_lif_override, "E2E read LIF override", false},
};

void sig_handler(int sig) {
  void *array[16];
  size_t size;

  // get void*'s for all entries on the stack
  size = backtrace(array, 16);

  // print out all the frames to stderr
  fprintf(stderr, "Error: signal %d:\n", sig);
  backtrace_symbols_fd(array, size, STDERR_FILENO);
  exit(1);
}

int main(int argc, char**argv) {
  gflags::ParseCommandLineFlags(&argc, &argv, true);
  signal(SIGSEGV, sig_handler);

  std::cout << "Input - hal_port: "   << FLAGS_hal_port 
            << ", test group: "       << FLAGS_test_group
            << ", polling interval: " << FLAGS_poll_interval 
            << std::endl;

  // Set the test group based on flags. Default is to allow all.
  if (FLAGS_test_group == "") {
      run_unit_tests = true;
      run_nvme_tests = true;
      run_nvme_be_tests = true;
      run_local_e2e_tests = true;
      run_comp_tests = true;
      run_xts_tests = true;
      run_rdma_tests = true;
  } else if (FLAGS_test_group == "unit") {
      run_unit_tests = true;
  } else if (FLAGS_test_group == "nvme") {
      run_nvme_tests = true;
  } else if (FLAGS_test_group == "nvme_be") {
      run_nvme_be_tests = true;
  } else if (FLAGS_test_group == "local_e2e") {
      run_local_e2e_tests = true;
  } else if (FLAGS_test_group == "comp") {
      run_comp_tests = true;
  } else if (FLAGS_test_group == "xts") {
      run_xts_tests = true;
  } else if (FLAGS_test_group == "rdma") {
      run_rdma_tests = true;
  } else {
    printf("Usage: ./storage_test [--hal_port <xxx>] [--test_group unit|nvme|nvme_be|local_e2e|comp|xts|rdma] "
           " [--poll_interval <yyy> \n");
    return -1;
  }

  printf("Starting configuration \n");
  if (tests::test_setup() < 0) {
    printf("Setup failed\n");
    return 1;
  }
  printf("Base configuration completed \n");

  printf("Going to init compression\n");
  tests::compression_init();
  printf("Compression configuration completed \n");

  printf("Going to init decompression\n");
  tests::decompression_init();
  printf("Decompression configuration completed \n");

  if (rdma_init() < 0) {
    printf("RDMA Setup failed\n");
    return 1;
  }
  printf("RDMA configuration completed \n");

  printf("Forming test suite based on configuration\n");
  // Add unit tests
  if (run_unit_tests || run_nvme_tests) {
    for (size_t i = 0; i < unit_tests.size(); i++) {
      test_suite.push_back(unit_tests[i]);
    }
    printf("Added unit tests \n");
  }

  // Add NVME tests
  if (run_nvme_tests) {
    for (size_t i = 0; i < nvme_tests.size(); i++) {
      test_suite.push_back(nvme_tests[i]);
    }
    printf("Added nvme tests \n");
  }

  // Add nvme_be tests
  if (run_nvme_be_tests) {
    for (size_t i = 0; i < nvme_be_tests.size(); i++) {
      test_suite.push_back(nvme_be_tests[i]);
    }
    printf("Added nvme_be tests \n");
  }

  // Add local_e2e tests
  if (run_local_e2e_tests) {
    for (size_t i = 0; i < local_e2e_tests.size(); i++) {
      test_suite.push_back(local_e2e_tests[i]);
    }
    printf("Added local_e2e tests \n");
  }

  // Add comp tests
  if (run_comp_tests) {
    for (size_t i = 0; i < comp_tests.size(); i++) {
      test_suite.push_back(comp_tests[i]);
    }
    printf("Added comp tests \n");
  }

  // Add xts tests
  if (run_xts_tests) {
    tests::add_xts_tests(test_suite);
    printf("Added XTS tests \n");
  }

  // Add rdma tests
  if (run_rdma_tests) {
    for (size_t i = 0; i < rdma_tests.size(); i++) {
      test_suite.push_back(rdma_tests[i]);
    }
    printf("Added RDMA tests \n");
  }
  printf("Formed test suite with %d cases \n", (int) test_suite.size());

  printf("Running test cases \n");
  for (size_t i = 0; i < test_suite.size(); i++) {
    printf(" Starting test #: %d name: %s \n", (int) i, test_suite[i].test_name.c_str());
    if (test_suite[i].test_fn() < 0)
      test_suite[i].test_succeded = false;
    else
      test_suite[i].test_succeded = true;
  }
  printf("Test case run complete, shutting down queues \n");
  queues::queues_shutdown();

  printf("\nConsolidated Test Report \n");
  printf("--------------------------------------------------------------\n");
  printf("Number\t\tName\t\t\tResult\n");
  printf("--------------------------------------------------------------\n");
  
  int rc = 0;
  for (size_t i = 0; i < test_suite.size(); i++) {
    printf("%lu\t", i+1);
    printf("%s\t\t\t\t", test_suite[i].test_name.c_str());
    printf("%s\n", test_suite[i].test_succeded ? "Success" : "Failure");
    if (!test_suite[i].test_succeded) rc = 1;
  }
  if (rc != 0) { 
    printf("\nOverall Report: FAILURE \n");
  } else {
    printf("\nOverall Report: SUCCESS \n");
  }
  fflush(stdout);
  if (rc != 0) return rc;
  exit(0);
}
