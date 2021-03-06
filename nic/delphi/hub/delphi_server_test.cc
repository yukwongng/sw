// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

#include <unistd.h>
#include <ev++.h>
#include "gtest/gtest.h"

#include "nic/delphi/utils/log.hpp"
#include "nic/delphi/hub/delphi_server.hpp"
#include "nic/delphi/sdk/delphi_sdk.hpp"

namespace {
using namespace std;
using namespace delphi;

class DelphiServerTest : public testing::Test {
protected:
    DelphiServerPtr    server;
public:
    virtual void SetUp() {
        // create delphi server without the messenger
        server = make_shared<DelphiServer>("/tmp/delphi_test.dat");
    }
    virtual void TearDown() {
    }
};

TEST_F(DelphiServerTest, MountTest) {
    MountReqMsgPtr mountReq = make_shared<MountReqMsg>();
    MountReqMsgPtr mountReq2 = make_shared<MountReqMsg>();
    MountReqMsgPtr mountReq3 = make_shared<MountReqMsg>();
    MountReqMsgPtr mountReq4 = make_shared<MountReqMsg>();
    MountRespMsgPtr mountResp = make_shared<MountRespMsg>();

    // mount with a service name
    mountReq->set_servicename("TestService");
    MountData *mnt = mountReq->add_mounts();
    mnt->set_kind("TestObject");
    mnt->set_mode(delphi::ReadWriteMode);

    error err = server->HandleMountReq(1, mountReq, mountResp);
    ASSERT_EQ(err, error::OK()) << "mount request failed";
    ASSERT_NE(mountResp->serviceid(), 0) << "mount response had invalid service id";
    uint16_t svcID = mountResp->serviceid();

    // send another mount with same service name again and verify we get same service id back
    err = server->HandleMountReq(1, mountReq, mountResp);
    ASSERT_EQ(err, error::OK()) << "duplicate mount request failed";
    ASSERT_EQ(mountResp->serviceid(), svcID) << "duplicate mount response had incorrect svc id";

    // send a mount request with different sock id and verify it fails
    err = server->HandleMountReq(2, mountReq, mountResp);
    ASSERT_NE(err, error::OK()) << "mount request suceeded with incorrect sock ctx";

    // send mount request with same sock but different service name
    mountReq->set_servicename("TestService2");
    err = server->HandleMountReq(1, mountReq, mountResp);
    ASSERT_NE(err, error::OK()) << "multiple mount request on same sock ctx suceeded";

    // trigger unmount
    err = server->HandleSocketClosed(1);
    ASSERT_EQ(err, error::OK()) << "socket disconnect failed";

    // closing unknown socket should fail
    err = server->HandleSocketClosed(2);
    ASSERT_NE(err, error::OK()) << "unknown socket disconnect suceeded";

    // connecting another service on same socket should suceed
    err = server->HandleMountReq(1, mountReq, mountResp);
    ASSERT_EQ(err, error::OK()) << "remounting failed";

    // verify connecting same service on different socket suceeds
    mountReq2->set_servicename("TestService");
    mnt = mountReq2->add_mounts();
    mnt->set_kind("TestObject2");
    mnt->set_mode(delphi::ReadMode);
    err = server->HandleMountReq(2, mountReq2, mountResp);
    ASSERT_EQ(err, error::OK()) << "remounting failed";

    // get a list of all services
    vector<ServiceInfoPtr> svcs = server->ListService();
    ASSERT_EQ(svcs.size(), 2) << "Incorrect number of services";
    for (vector<ServiceInfoPtr>::iterator iter=svcs.begin(); iter!=svcs.end(); ++iter) {
        ASSERT_EQ((*iter)->Mounts.size(), 1) << "Incorrect number of mounts on service";
    }

    // verify two services can not mount same kind in RW mode
    mountReq->set_servicename("TestService3");
    err = server->HandleMountReq(3, mountReq, mountResp);
    ASSERT_NE(err, error::OK()) << "duplicate RW mount suceeded";

    // verify a service can mount a key in RW mode while parent is mounted in RO
    mountReq3->set_servicename("TestService4");
    mnt = mountReq3->add_mounts();
    mnt->set_kind("TestObject2");
    mnt->set_mode(delphi::ReadWriteMode);
    mnt = mountReq3->add_mounts();
    mnt->set_kind("TestObject2");
    mnt->set_key("TestKey2");
    mnt->set_mode(delphi::ReadWriteMode);
    err = server->HandleMountReq(4, mountReq3, mountResp);
    ASSERT_EQ(err, error::OK()) << "remounting failed";

    // verify a child can not mount RW mode while parent ismounted RW
    mountReq4->set_servicename("TestService5");
    mnt = mountReq4->add_mounts();
    mnt->set_kind("TestObject");
    mnt->set_key("TestKey");
    mnt->set_mode(delphi::ReadWriteMode);
    err = server->HandleMountReq(5, mountReq4, mountResp);
    ASSERT_NE(err, error::OK()) << "duplicate parent/child RW mount suceeded";
}

TEST_F(DelphiServerTest, BasicObjectTest) {
    MountReqMsgPtr mountReq = make_shared<MountReqMsg>();
    MountRespMsgPtr mountResp = make_shared<MountRespMsg>();
    vector<ObjectData> objReq;
    vector<ObjectData> objDelReq;
    vector<ObjectData *> objResp;
    const int num_objects = 5;
    error err;

    // create some objects
    for (int i = 0; i < num_objects; i++) {
        ObjectData obj;
        string key = "TestKey-" + to_string(i);
        ObjectMeta *meta = obj.mutable_meta();
        meta->set_kind("TestObject");
        meta->set_key(key);
        meta->set_path(getPath("TestObject", key));
        meta->set_handle(i+1);
        obj.set_op(delphi::SetOp);
        objReq.push_back(obj);
        err = server->HandleChangeReq(1, objReq, &objResp);
        ASSERT_EQ(err, error::OK()) << "change request failed";
    }

    // verify all the objects are created
    DbSubtreePtr subtree = server->GetSubtree("TestObject");
    ASSERT_EQ(subtree->objects.size(), num_objects) << "Not all objects are in the BD";
    ASSERT_NE(subtree->objects["TestKey-0"]->meta().handle(), 0) << "object has incorrect handle";

    // send a mount request and verify we get all the objects back
    mountReq->set_servicename("TestService");
    MountData *mnt = mountReq->add_mounts();
    mnt->set_kind("TestObject");
    mnt->set_mode(delphi::ReadWriteMode);
    err = server->HandleMountReq(1, mountReq, mountResp);
    ASSERT_EQ(err, error::OK()) << "mount request failed";
    ASSERT_EQ(mountResp->objects().size(), num_objects) << "mount response had invalid number of objects";

    // create duplicate object with different handle and verify it overwrites the old object
    vector<ObjectData> objReq2;
    ObjectData obj2;
    ObjectMeta *meta2 = obj2.mutable_meta();
    meta2->set_kind("TestObject");
    meta2->set_key("TestKey-0");
    meta2->set_path(getPath("TestObject", "TestKey-0"));
    meta2->set_handle(2);
    obj2.set_op(delphi::SetOp);
    objReq2.push_back(obj2);
    err = server->HandleChangeReq(1, objReq2, &objResp);
    ASSERT_EQ(err, error::OK()) << "change request failed";
    subtree = server->GetSubtree("TestObject");
    ASSERT_EQ(subtree->objects.size(), num_objects) << "incorrect number of objects in db";
    ASSERT_EQ(subtree->objects["TestKey-0"]->meta().handle(), 2) << "object has incorrect handle";

    // verify object with empty kind, key, handle is not accepted
    meta2->set_kind("");
    meta2->set_key("TestKey-0");
    meta2->set_handle(2);
    meta2->set_path(getPath("TestObject", "TestKey-0"));
    objReq2.clear();
    objReq2.push_back(obj2);
    err = server->HandleChangeReq(1, objReq2, &objResp);
    ASSERT_NE(err, error::OK()) << "object add with no kind suceeded";
    meta2->set_kind("TestObject");
    meta2->set_key("");
    err = server->HandleChangeReq(1, objReq2, &objResp);
    ASSERT_NE(err, error::OK()) << "object add with no key suceeded";
    meta2->set_key("TestKey-0");
    meta2->set_handle(0);
    objReq2.clear();
    objReq2.push_back(obj2);
    err = server->HandleChangeReq(1, objReq2, &objResp);
    ASSERT_NE(err, error::OK()) << "object add with no handle suceeded";
    meta2->set_handle(2);
    meta2->set_path("");
    objReq2.clear();
    objReq2.push_back(obj2);
    err = server->HandleChangeReq(1, objReq2, &objResp);
    ASSERT_NE(err, error::OK()) << "object add with no path suceeded";

    // test deleting the object
    for (int i = 0; i < num_objects; i++) {
        ObjectData obj;
        string key = "TestKey-" + to_string(i);
        ObjectMeta *meta = obj.mutable_meta();
        meta->set_kind("TestObject");
        meta->set_key(key);
        meta->set_path(getPath("TestObject", key));
        meta->set_handle(i+1);
        obj.set_op(delphi::DeleteOp);
        objDelReq.push_back(obj);
        err = server->HandleChangeReq(1, objDelReq, &objResp);
        ASSERT_EQ(err, error::OK()) << "change request failed";
    }

    // make sure objects are gone
    subtree = server->GetSubtree("TestObject");
    ASSERT_EQ(subtree->objects.size(), 0) << "objects are still in DB";
    mountResp->clear_objects();
    err = server->HandleMountReq(1, mountReq, mountResp);
    ASSERT_EQ(err, error::OK()) << "mount request failed";
    ASSERT_EQ(mountResp->objects().size(), 0) << "mount response had objects after delete";
}    
} // namespace

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
