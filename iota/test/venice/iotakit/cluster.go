// {C} Copyright 2019 Pensando Systems Inc. All rights reserved.

package iotakit

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/cluster"
	iota "github.com/pensando/sw/iota/protos/gogen"
	"github.com/pensando/sw/venice/utils/log"
)

// GetVeniceURL returns venice URL for the testbed
func (tb *TestBed) GetVeniceURL() []string {
	var veniceURL []string

	if tb.mockMode {
		return []string{mockVeniceURL}
	}

	// walk all venice nodes
	for _, node := range tb.Nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			veniceURL = append(veniceURL, fmt.Sprintf("%s:9000", node.NodeMgmtIP))
		}
	}

	return veniceURL
}

// getVeniceHostNames returns a list of venice host names
func (tb *TestBed) getVeniceHostNames() []string {
	var hostNames []string

	// walk all venice nodes
	for _, node := range tb.Nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			hostNames = append(hostNames, node.NodeName)
		}
	}

	return hostNames
}

// CheckIotaClusterHealth checks iota cluster health
func (tb *TestBed) CheckIotaClusterHealth() error {
	// check cluster health
	topoClient := iota.NewTopologyApiClient(tb.iotaClient.Client)
	healthResp, err := topoClient.CheckClusterHealth(context.Background(), tb.addNodeResp)
	if err != nil {
		log.Errorf("Failed to check health of the cluster. Err: %v", err)
		return fmt.Errorf("Cluster health check failed. %v", err)
	} else if healthResp.ApiResponse.ApiStatus != iota.APIResponseType_API_STATUS_OK {
		log.Errorf("Failed to check health of the cluster {%+v}. Err: %v", healthResp.ApiResponse, err)
		return fmt.Errorf("Cluster health check failed %v", healthResp.ApiResponse.ApiStatus)
	}

	for _, h := range healthResp.Health {
		if h.HealthCode != iota.NodeHealth_HEALTH_OK {
			log.Errorf("Testbed unhealthy. HealthCode: %v | Node: %v", h.HealthCode, h.NodeName)
			return fmt.Errorf("Cluster health check failed")
		}
	}

	log.Debugf("Got cluster health resp: %+v", healthResp)

	return nil
}

// MakeVeniceCluster inits the venice cluster
func (tb *TestBed) MakeVeniceCluster() error {
	// get CMD URL URL
	var veniceCmdURL []string
	for _, node := range tb.Nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			veniceCmdURL = append(veniceCmdURL, fmt.Sprintf("%s:9001", node.NodeMgmtIP))
		}
	}

	// cluster message to init cluster
	clusterCfg := cluster.Cluster{
		TypeMeta: api.TypeMeta{Kind: "Cluster"},
		ObjectMeta: api.ObjectMeta{
			Name: "iota-cluster",
		},
		Spec: cluster.ClusterSpec{
			AutoAdmitNICs: true,
			QuorumNodes:   tb.getVeniceHostNames(),
		},
	}

	// make cluster message to be sent to API server
	clusterStr, _ := json.Marshal(clusterCfg)
	makeCluster := iota.MakeClusterMsg{
		ApiResponse: &iota.IotaAPIResponse{},
		Endpoint:    veniceCmdURL[0] + "/api/v1/cluster",
		Config:      string(clusterStr),
	}

	log.Infof("Making Venice cluster..")

	// ask iota server to make cluster
	log.Debugf("Making cluster with params: %+v", makeCluster)
	cfgClient := iota.NewConfigMgmtApiClient(tb.iotaClient.Client)
	resp, err := cfgClient.MakeCluster(context.Background(), &makeCluster)
	if err != nil {
		log.Errorf("Error initing venice cluster. Err: %v", err)
		return err
	}
	if resp.ApiResponse.ApiStatus != iota.APIResponseType_API_STATUS_OK {
		log.Errorf("Error making venice cluster: ApiResp: %+v. Err %v", resp.ApiResponse, err)
		return fmt.Errorf("Error making venice cluster")
	}
	tb.makeClustrResp = resp

	// done
	return nil
}

// SetupVeniceNodes sets up some test tools on venice nodes
func (tb *TestBed) SetupVeniceNodes() error {

	// walk all venice nodes
	trig := tb.NewTrigger()
	for _, node := range tb.Nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			entity := node.NodeName + "_venice"
			trig.AddCommand(fmt.Sprintf("mkdir -p /pensando/iota/k8s/"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("sudo cp -r /var/lib/pensando/pki/kubernetes/apiserver-client /pensando/iota/k8s/"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("sudo chmod -R 777 /pensando/iota/k8s/"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("if ! [ -f /tmp/kubernetes-server-linux-amd64.tar.gz ]; then curl -fL --retry 3 --keepalive-time 2  -o /tmp/kubernetes-server-linux-amd64.tar.gz  https://storage.googleapis.com/kubernetes-release/release/v1.7.14/kubernetes-server-linux-amd64.tar.gz; fi"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("tar xvf /tmp/kubernetes-server-linux-amd64.tar.gz  kubernetes/server/bin/kubectl"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("chmod 755 kubernetes/server/bin/kubectl"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("sudo cp kubernetes/server/bin/kubectl /usr/local/bin/kubectl"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf(`echo '/pensando/iota/kubernetes/server/bin/kubectl config set-cluster e2e --server=https://%s:6443 --certificate-authority=/pensando/iota/k8s/apiserver-client/ca-bundle.pem; 
				/pensando/iota/kubernetes/server/bin/kubectl config set-context e2e --cluster=e2e --user=admin;
				/pensando/iota/kubernetes/server/bin/kubectl config use-context e2e;
				/pensando/iota/kubernetes/server/bin/kubectl config set-credentials admin --client-certificate=/pensando/iota/k8s/apiserver-client/cert.pem --client-key=/pensando/iota/k8s/apiserver-client/key.pem;
				' > /pensando/iota/setup_kubectl.sh
				`, node.NodeName), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("chmod +x /pensando/iota/setup_kubectl.sh"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("/pensando/iota/setup_kubectl.sh"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf(`echo 'docker run --rm --name kibana --network host \
				-v /var/lib/pensando/pki/shared/elastic-client-auth:/usr/share/kibana/config/auth \
				-e ELASTICSEARCH_URL=https://%s:9200 \
				-e ELASTICSEARCH_SSL_CERTIFICATEAUTHORITIES="config/auth/ca-bundle.pem" \
				-e ELASTICSEARCH_SSL_CERTIFICATE="config/auth/cert.pem" \
				-e ELASTICSEARCH_SSL_KEY="config/auth/key.pem" \
				-e xpack.security.enabled=false \
				-e xpack.logstash.enabled=false \
				-e xpack.graph.enable=false \
				-e xpack.watcher.enabled=false \
				-e xpack.ml.enabled=false \
				-e xpack.monitoring.enabled=false \
				-d docker.elastic.co/kibana/kibana:6.3.0
				' > /pensando/iota/start_kibana.sh
				`, node.NodeName), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("chmod +x /pensando/iota/start_kibana.sh"), entity, node.NodeName)
		}
	}

	// trigger commands
	triggerResp, err := trig.RunSerial()
	if err != nil {
		log.Errorf("Failed to setup venice node. Err: %v", err)
		return fmt.Errorf("Error triggering commands on venice nodes: %v", err)
	}

	for _, cmdResp := range triggerResp {
		if cmdResp.ExitCode != 0 {
			return fmt.Errorf("Venice trigger %v failed. code %v, Out: %v, StdErr: %v", cmdResp.Command, cmdResp.ExitCode, cmdResp.Stdout, cmdResp.Stderr)
		}
	}

	return nil
}

// CheckVeniceServiceStatus checks if all services are running on venice nodes
func (tb *TestBed) CheckVeniceServiceStatus(leaderNode string) error {

	trig := tb.NewTrigger()
	for _, node := range tb.Nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			entity := node.NodeName + "_venice"
			trig.AddCommand(fmt.Sprintf("docker ps -q -f Name=pen-cmd"), entity, node.NodeName)
			trig.AddCommand(fmt.Sprintf("docker ps -q -f Name=pen-etcd"), entity, node.NodeName)
		}
	}

	// trigger commands
	triggerResp, err := trig.Run()
	if err != nil {
		log.Errorf("Failed to check cmd/etcd service status. Err: %v", err)
		return err
	}

	for _, cmdResp := range triggerResp {
		if cmdResp.ExitCode != 0 {
			return fmt.Errorf("Venice trigger %v failed. code %v, Out: %v, StdErr: %v", cmdResp.Command, cmdResp.ExitCode, cmdResp.Stdout, cmdResp.Stderr)
		}
		if cmdResp.Stdout == "" {
			return fmt.Errorf("Venice required service not running: %v", cmdResp.Command)
		}
	}

	// check all pods on leader node

	for _, node := range tb.Nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_VENICE && node.NodeName == leaderNode {
			trig = tb.NewTrigger()
			entity := node.NodeName + "_venice"
			trig.AddCommand(fmt.Sprintf("/pensando/iota/kubernetes/server/bin/kubectl get pods -owide --no-headers"), entity, node.NodeName)

			// trigger commands
			triggerResp, err = trig.Run()
			if err != nil {
				log.Errorf("Failed to get k8s service status Err: %v", err)
				return err
			}

			for _, cmdResp := range triggerResp {
				if cmdResp.ExitCode != 0 {
					return fmt.Errorf("Venice trigger %v failed. code %v, Out: %v, StdErr: %v", cmdResp.Command, cmdResp.ExitCode, cmdResp.Stdout, cmdResp.Stderr)
				}
				if cmdResp.Stdout == "" {
					return fmt.Errorf("Could not get any information from k8s: %v", cmdResp.Command)
				}
				log.Debugf("Got kubectl resp\n%v", cmdResp.Stdout)
				out := strings.Split(cmdResp.Stdout, "\n")
				for _, line := range out {
					if line != "" && !strings.Contains(line, "Running") {
						fmt.Printf("Some kuberneted services were not running: %v", cmdResp.Stdout)
						return fmt.Errorf("Some pods not running: %v", line)
					}
				}
			}
		}
	}

	return nil
}
