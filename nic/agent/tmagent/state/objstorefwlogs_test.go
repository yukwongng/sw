// {C} Copyright 2020 Pensando Systems Inc. All rights reserved.

package state

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/csv"
	"fmt"
	"net"
	"net/http"
	_ "os"
	_ "os/exec"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/pensando/sw/api"
	_ "github.com/pensando/sw/nic/agent/dscagent"
	"github.com/pensando/sw/nic/agent/dscagent/types"
	halproto "github.com/pensando/sw/nic/agent/dscagent/types/irisproto"
	_ "github.com/pensando/sw/nic/agent/protos/netproto"
	_ "github.com/pensando/sw/nic/agent/protos/tpmprotos"
	servicetypes "github.com/pensando/sw/venice/cmd/types/protos"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/netutils"
	objstore "github.com/pensando/sw/venice/utils/objstore/client"
	"github.com/pensando/sw/venice/utils/resolver/mock"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func listBuckets(w http.ResponseWriter, r *http.Request) {
	resp := `{<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/"><Owner><ID>02d6176db174dc93cb1b899f7c6078f08654445fe8cf1b6ce98d8855f66bdbf4</ID><DisplayName></DisplayName></Owner><Buckets><Bucket><Name>tenant1-pktcap</Name><CreationDate>2018-06-05T00:44:14.880Z</CreationDate></Bucket><Bucket><Name>tenant1.pktcap</Name><CreationDate>2018-06-05T07:05:22.601Z</CreationDate></Bucket></Buckets></ListAllMyBucketsResult>} `
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
	w.Header().Set("Content-Type", "application/xml")
}

func minioServer(l net.Listener) {
	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("/"), listBuckets).Methods("GET")
	go http.Serve(l, router)
}

// TestProcessFWEventForObjStore tests the fwlog to objstore pipeline
func TestProcessFWEventForObjStore(t *testing.T) {
	// Start local minio server
	l, err := net.Listen("tcp", "127.0.0.1:")
	AssertOk(t, err, "failed listen")
	minioServer(l)
	defer l.Close()
	url := l.Addr().(*net.TCPAddr).String()
	// url := "127.0.0.1:19001"
	// url := "127.0.0.1:9000"

	c := gomock.NewController(t)
	defer c.Finish()

	r := mock.New()
	retryOpt := objstore.WithConnectRetries(1)
	_, err = objstore.NewClient("ten1", "svc1", r, retryOpt)
	Assert(t, err != nil, "failed test client error ")

	err = r.AddServiceInstance(&servicetypes.ServiceInstance{
		TypeMeta: api.TypeMeta{
			Kind: "ServiceInstance",
		},
		ObjectMeta: api.ObjectMeta{
			Name: globals.VosMinio,
		},
		Service: globals.VosMinio,
		URL:     url,
	})
	AssertOk(t, err, "failed to add objstore sercvice")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ps, err := NewTpAgent(ctx, strings.Split(types.DefaultAgentRestURL, ":")[1])
	AssertOk(t, err, "failed to create tp agent")
	Assert(t, ps != nil, "invalid policy state received")
	defer ps.Close()

	err = ps.FwlogInit(FwlogIpcShm)
	AssertOk(t, err, "failed to init FwLog")

	testChannel := make(chan TestObject, 10000)
	err = ps.ObjStoreInit("1", r, time.Duration(1)*time.Second, testChannel)
	AssertOk(t, err, "failed to init objectstore")

	srcIPStr := "192.168.10.1"
	srcIP, err := netutils.IPv4ToUint32(srcIPStr)
	AssertOk(t, err, "failed to convert ip address")

	destIPStr := "192.168.20.1"
	destIP, err := netutils.IPv4ToUint32(destIPStr)
	AssertOk(t, err, "failed to convert ip address")

	// send logs
	events := []struct {
		msgBsd     string
		msgRfc5424 string
		fwEvent    *halproto.FWEvent
	}{
		{
			fwEvent: &halproto.FWEvent{
				SourceVrf:       65,
				DestVrf:         65,
				Sipv4:           srcIP,
				Dipv4:           destIP,
				Dport:           10000,
				IpProt:          1,
				Fwaction:        halproto.SecurityAction_SECURITY_RULE_ACTION_DENY,
				Direction:       0,
				Alg:             halproto.ALGName_APP_SVC_TFTP,
				SessionId:       uint64(1000),
				ParentSessionId: uint64(1001),
				RuleId:          uint64(1000),
				Icmptype:        0,
				Icmpcode:        0,
				Icmpid:          0,
				AppId:           32,
			},
		},
	}

	for _, e := range events {
		ps.ProcessFWEvent(e.fwEvent, time.Now())
	}

	testObject := <-testChannel
	Assert(t, testObject.ObjectName != "", "object name is empty")
	Assert(t, testObject.BucketName == "default.fwlogs", "bucket name is not correct")
	Assert(t, testObject.IndexBucketName == "meta-default.fwlogs", "index bucket name is not correct")
	verifyData(t, testObject.Data)
	Assert(t, testObject.Index != "", "index is empty")
	Assert(t, len(testObject.Meta) == 7, "object meta is empty %s", testObject.Meta)
	Assert(t, testObject.Meta["startts"] != "", "object meta's startts is empty")
	Assert(t, testObject.Meta["endts"] != "", "object meta's endts is empty")
	Assert(t, testObject.Meta["logcount"] == "1",
		"object meta's log count is not correct, expectecd 1, received %d", testObject.Meta["logcount"])
	Assert(t, testObject.Meta["nodeid"] == "1",
		"object meta's nodeid is not correct, expectecd 1, received %d", testObject.Meta["nodeid"])
	Assert(t, testObject.Meta["csvversion"] == "v1",
		"object meta's data csv version is not correct, expectecd v1, received %d", testObject.Meta["csvversion"])
	Assert(t, testObject.Meta["metaversion"] == "v1",
		"object meta's meta csv version is not correct, expectecd v1, received %d", testObject.Meta["metaversion"])

	// Check metrics
	Assert(t, metric.fwlogDrops.Value() == int64(0), "fwlog object got dropped")
	Assert(t, metric.fwlogSuccess.Value() == int64(1), "fwlog object success metric is 0, expected 1")

	// TestObjStoreErrors tests the error scenarios in the object store pipeline
	t.Run("TestObjStoreErrors", func(t *testing.T) {
		// Pass nil resolver
		err = ps.ObjStoreInit("1", nil, time.Duration(1)*time.Second, nil)
		Assert(t, err != nil, "failed to init objectstore")
	})

	// TestInternalUntriggeredFunctions tests the functions that are not triggered by the testcases written above
	t.Run("TestInternalUntriggeredFunctions", func(t *testing.T) {
		metric.addDrop()
		Assert(t, metric.fwlogDrops.Value() == 1, "fwlogDrops metric did not update")
		metric.addSuccess()
		Assert(t, metric.fwlogSuccess.Value() == 2, "fwlogSuccess metric did not update")
		metric.addRetries(5)
		Assert(t, metric.fwlogRetries.Get("5").String() == "1", "fwlogRetries metric did not update")
	})

	// done := make(chan bool)
	// <-done
}

func verifyData(t *testing.T, data bytes.Buffer) {
	reader := bytes.NewReader(data.Bytes())
	Assert(t, data.Len() != 0, "data is empty")
	zipReader, err := gzip.NewReader(reader)
	AssertOk(t, err, "error in unzipping data")

	rd := csv.NewReader(zipReader)
	AssertOk(t, err, "error in reading csv file")
	lines, err := rd.ReadAll()
	Assert(t, len(lines) == 2, "incorrect number of logs in the file")

	Assert(t, lines[1][0] == "65", "incorrect src vrf")
	Assert(t, lines[1][1] == "65", "incorrect dest vrf")
	Assert(t, lines[1][2] == "192.168.10.1", "incorrect source ip")
	Assert(t, lines[1][3] == "192.168.20.1", "incorrect dest ip")
	Assert(t, lines[1][16] == "32", "incorrect app id")
	Assert(t, lines[1][17] == "TFTP", "incorrect alg, %s", lines[1][17])
	Assert(t, lines[1][18] == "1", "incorrect count")
}
