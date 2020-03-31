package enterprise

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pensando/sw/api"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/objects"
)

// GetFwLogObjectCount gets the object count for firewall logs under the bucket with the given name
func (sm *SysModel) GetFwLogObjectCount(
	tenantName string, bucketName string, objectKeyPrefix string) (int, error) {
	opts := api.ListWatchOptions{
		ObjectMeta: api.ObjectMeta{
			Tenant:    tenantName,
			Namespace: bucketName,
		},
	}

	ctx, err := sm.VeniceLoggedInCtx(context.Background())
	if err != nil {
		return 0, err
	}

	restClients, err := sm.VeniceRestClient()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, restClient := range restClients {
		list, err := restClient.ObjstoreV1().Object().List(ctx, &opts)
		if err != nil {
			return 0, err
		}

		if len(list) != 0 {
			if objectKeyPrefix == "" {
				return len(list), nil
			}

			for _, o := range list {
				if strings.Contains(o.Name, objectKeyPrefix) {
					count++
				}
			}

			return count, nil
		}
	}

	return 0, nil
}

// getLatestObjectName gets the last entry from the list
// Minio lists objects in chronological order, hence its safe to return the last
// entry from the list.
func (sm *SysModel) getLatestObjectName(tenantName, bucketName, objectKeyPrefix string) (string, error) {
	temp := []string{}

	opts := api.ListWatchOptions{
		ObjectMeta: api.ObjectMeta{
			Tenant:    tenantName,
			Namespace: bucketName,
		},
	}

	ctx, err := sm.VeniceLoggedInCtx(context.Background())
	if err != nil {
		return "", err
	}

	restClients, err := sm.VeniceRestClient()
	if err != nil {
		return "", err
	}

	for _, restClient := range restClients {
		list, err := restClient.ObjstoreV1().Object().List(ctx, &opts)
		if err != nil {
			return "", err
		}

		if len(list) != 0 {
			for _, o := range list {
				if objectKeyPrefix == "" || strings.Contains(o.Name, objectKeyPrefix) {
					temp = append(temp, o.Name)
				}
			}
		}
	}

	if len(temp) == 0 {
		return "", fmt.Errorf("no objects found for prefix %s", objectKeyPrefix)
	}

	return temp[len(temp)-1], nil
}

// FindFwlogForWorkloadPairsFromObjStore finds workload ip addresses in firewall log
func (sm *SysModel) FindFwlogForWorkloadPairsFromObjStore(
	tenantName, protocol string, port uint32, fwaction string, wpc *objects.WorkloadPairCollection) error {
	for _, wp := range wpc.Pairs {
		ipA := wp.First.GetIP()
		ipB := wp.Second.GetIP()
		aMac := wp.First.NaplesMAC()
		bMac := wp.Second.NaplesMAC()
		return sm.findFwlogForWorkloadPairsFromObjStore(tenantName,
			ipA, ipB, protocol, port, fwaction, aMac, bMac)
	}
	return nil
}

func (sm *SysModel) findFwlogForWorkloadPairsFromObjStore(
	tenantName, srcIP, destIP, protocol string, port uint32, fwaction, naplesA, naplesB string) error {
	latestObjectNameA, err := sm.getLatestObjectName(tenantName, "fwlogs", naplesA)
	if err != nil {
		return fmt.Errorf("could not find latest object for naples %s", naplesA)
	}

	latestObjectNameB, err := sm.getLatestObjectName(tenantName, "fwlogs", naplesB)
	if err != nil {
		return fmt.Errorf("could not find latest object for naples %s", naplesB)
	}

	url := sm.GetVeniceURL()[0]

	dataNaplesA, err := sm.downloadCsvFileViaPSMRESTAPI("fwlogs", latestObjectNameA, url)
	if err != nil {
		return err
	}

	fmt.Println("Data naplesA", dataNaplesA)

	dataNaplesB, err := sm.downloadCsvFileViaPSMRESTAPI("fwlogs", latestObjectNameB, url)
	if err != nil {
		return err
	}

	fmt.Println("Data naplesB", dataNaplesB)

	// reject or deny logs will appear only on one DSC.
	shouldVerifyBothNaples := true
	if fwaction == "reject" || fwaction == "deny" {
		shouldVerifyBothNaples = false
	}

	// enhance it to match the given log info
	errA := sm.isLogPresent(dataNaplesA, srcIP, destIP, protocol, port, fwaction)
	errB := sm.isLogPresent(dataNaplesB, srcIP, destIP, protocol, port, fwaction)

	if shouldVerifyBothNaples {
		if errA != nil {
			return fmt.Errorf("error for naplesA %s, err %s", naplesA, errA.Error())
		}
		if errB != nil {
			return fmt.Errorf("error for naplesB %s, err %s", naplesB, errB.Error())
		}
	} else if errA != nil && errB != nil {
		return fmt.Errorf("error on one of the naples, err %s", errA.Error())
	}

	return nil
}

func (sm *SysModel) isLogPresent(data [][]string,
	srcIP, destIP, protocol string, port uint32, fwaction string) error {
	for _, line := range data {
		if (line[2] == srcIP || line[2] == destIP) &&
			(line[3] == destIP || line[3] == srcIP) &&
			line[6] == strconv.Itoa(int(port)) &&
			line[7] == protocol &&
			line[8] == fwaction {
			return nil
		}
	}
	return fmt.Errorf("log not found for srcIP %s, destIP %s, protocol %s, port %d, fwAction %s",
		srcIP, destIP, protocol, port, fwaction)
}

func (sm *SysModel) downloadCsvFileViaPSMRESTAPI(bucketName, objectName string, url string) ([][]string, error) {
	ctx, err := sm.VeniceLoggedInCtx(context.Background())
	if err != nil {
		return nil, err
	}

	// replace first 5 "/" with "_"
	name := strings.Replace(objectName, "/", "_", 5)
	uri := fmt.Sprintf("https://%s/objstore/v1/downloads/%s/%s", url, bucketName, name)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	authzHeader, ok := loginctx.AuthzHeaderFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("error in getting authorization header")
	}

	req.Header.Set("Authorization", authzHeader)
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("returned error code from PSM %d", resp.StatusCode)
	}

	// body is a zipped file
	reader := bytes.NewReader(body)
	zipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	rd := csv.NewReader(zipReader)
	lines, err := rd.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
