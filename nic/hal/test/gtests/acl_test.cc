#include "nic/include/base.h"
#include "nic/hal/src/acl.hpp"
#include "nic/proto/types.pb.h"
#include "nic/proto/hal/acl.pb.h"
#include "nic/hal/hal.hpp"
#include <gtest/gtest.h>
#include <stdio.h>
#include <stdlib.h>
#include "nic/p4/nw/include/defines.h"

using acl::AclSpec;
using acl::AclResponse;
using acl::AclKeyHandle;
using acl::AclSelector;
using acl::AclActionInfo;
using acl::AclDeleteRequest;
using acl::AclDeleteResponse;

void
hal_initialize()
{
    char 			cfg_file[] = "hal.json";
	char 			*cfg_path;
    std::string     full_path;
    hal::hal_cfg_t  hal_cfg = { 0 };

    cfg_path = std::getenv("HAL_CONFIG_PATH");
    if (cfg_path) {
        full_path =  std::string(cfg_path) + "/" + std::string(cfg_file);
        std::cerr << "full path " << full_path << std::endl;
    } else {
        full_path = std::string(cfg_file);
    }

    // make sure cfg file exists
    if (access(full_path.c_str(), R_OK) < 0) {
        fprintf(stderr, "config file %s has no read permissions\n",
                full_path.c_str());
        exit(1);
    }

    printf("Json file: %s\n", full_path.c_str());

    if (hal::hal_parse_cfg(full_path.c_str(), &hal_cfg) != HAL_RET_OK) {
        fprintf(stderr, "HAL config file parsing failed, quitting ...\n");
        ASSERT_TRUE(0);
    }
    printf("Parsed cfg json file \n");

    // initialize HAL
    if (hal::hal_init(&hal_cfg) != HAL_RET_OK) {
        fprintf(stderr, "HAL initialization failed, quitting ...\n");
        exit(1);
    }
    printf("HAL Initialized \n");
}

class acl_test : public ::testing::Test {
protected:
  acl_test() {
  }

  virtual ~acl_test() {
  }

  // will be called immediately after the constructor before each test
  virtual void SetUp() {
  }

  // will be called immediately after each test before the destructor
  virtual void TearDown() {
  }

  // Will be called at the beginning of all test cases in this class
  static void SetUpTestCase() {
    hal_initialize();
  }
  // Will be called at the end of all test cases in this class
  static void TearDownTestCase() {
  }
};

// Creating acls
TEST_F(acl_test, test1)
{
    hal_ret_t     ret;
    AclSpec       spec;
    AclResponse   rsp;

    AclSelector   *match;
    AclActionInfo *action;

    match = spec.mutable_match();
    match->mutable_ip_selector()->set_ip_af(types::IPAddressFamily::IP_AF_INET);
    match->mutable_ip_selector()->mutable_src_prefix()->mutable_address()->set_ip_af(types::IPAddressFamily::IP_AF_INET);
    match->mutable_ip_selector()->mutable_src_prefix()->mutable_address()->set_v4_addr(1);
    match->mutable_ip_selector()->mutable_src_prefix()->set_prefix_len(32);

    action = spec.mutable_action();
    action->set_action(acl::AclAction::ACL_ACTION_DENY);

    spec.mutable_key_or_handle()->set_acl_id(1);
    spec.set_priority(100);

    hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
    ret = hal::acl_create(spec, &rsp);
    hal::hal_cfg_db_close();
    ASSERT_TRUE(ret == HAL_RET_OK);
    spec.Clear();
}

TEST_F(acl_test, test2)
{
    hal_ret_t     ret;
    AclSpec       spec;
    AclResponse   rsp;

    AclSelector   *match; 
    AclActionInfo *action;

    for (int i = 10; i < 20; i++) {
        match = spec.mutable_match();
        match->mutable_ip_selector()->set_ip_af(types::IPAddressFamily::IP_AF_INET);
        match->mutable_ip_selector()->mutable_src_prefix()->mutable_address()->set_ip_af(types::IPAddressFamily::IP_AF_INET);
        match->mutable_ip_selector()->mutable_src_prefix()->mutable_address()->set_v4_addr(1);
        match->mutable_ip_selector()->mutable_src_prefix()->set_prefix_len(32);

        action = spec.mutable_action();
        action->set_action(acl::AclAction::ACL_ACTION_DENY);

        spec.mutable_key_or_handle()->set_acl_id(i);
        spec.set_priority(100);

        hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
        ret = hal::acl_create(spec, &rsp);
        hal::hal_cfg_db_close();
        ASSERT_TRUE(ret == HAL_RET_OK);
        spec.Clear();
    }
}

// Create and Delete ACLs
TEST_F(acl_test, test3)
{
    hal_ret_t     ret;
    AclSpec       spec;
    AclResponse   rsp;
    AclDeleteRequest del_req;
    AclDeleteResponse del_rsp;

    AclSelector   *match; 
    AclActionInfo *action;

    std::vector<hal_handle_t> entries;

    for (int i = 100; i < 110; i++) {
        match = spec.mutable_match();
        match->mutable_ip_selector()->set_ip_af(types::IPAddressFamily::IP_AF_INET6);
        match->mutable_ip_selector()->mutable_src_prefix()->mutable_address()->set_ip_af(types::IPAddressFamily::IP_AF_INET6);
        match->mutable_ip_selector()->mutable_src_prefix()->mutable_address()->set_v6_addr("00010001000100010001000100010001");
        match->mutable_ip_selector()->mutable_src_prefix()->set_prefix_len(64);

        action = spec.mutable_action();
        action->set_action(acl::AclAction::ACL_ACTION_DENY);

        spec.mutable_key_or_handle()->set_acl_id(i);
        spec.set_priority(100);

        hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
        ret = hal::acl_create(spec, &rsp);
        hal::hal_cfg_db_close();
        ASSERT_TRUE(ret == HAL_RET_OK);

        entries.push_back(rsp.status().acl_handle().handle());
        spec.Clear();
        rsp.Clear();
    }

    for (auto &entry : entries) {
        del_req.mutable_key_or_handle()->mutable_acl_handle()->set_handle(entry);
        hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
        ret = hal::acl_delete(del_req, &del_rsp);
        hal::hal_cfg_db_close();
        ASSERT_EQ(ret, HAL_RET_OK);
        ASSERT_EQ(del_rsp.api_status(), types::API_STATUS_OK);
        del_req.Clear();
        del_rsp.Clear();
    }

}


int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
