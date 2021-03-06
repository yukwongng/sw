package services

import (
	"testing"

	"github.com/pensando/sw/venice/cmd/services/mock"
	"github.com/pensando/sw/venice/cmd/types"
)

func setupNode(testName string) (types.SystemdService, types.NodeService) {
	s := NewSystemdService(WithSysIfSystemdSvcOption(&mock.SystemdIf{}))
	n := NewNodeService(testName, WithSystemdSvcNodeOption(s), WithConfigsNodeOption(&mock.Configs{}))
	return s, n
}

func checkAllNodeServiceStarted(t *testing.T, s types.SystemdService) {
	checkAllServiceStarted(t, s, nodeServices)
}

func TestNodeServiceStopServices(t *testing.T) {
	t.Parallel()
	s, n := setupNode("TestNodeServiceStopServices")

	s.Start()
	n.Start()

	checkAllNodeServiceStarted(t, s)
	n.Stop()
	checkAllServiceStopped(t, s)
}

func TestNodeServiceStartBeforeSystemdService(t *testing.T) {
	t.Parallel()
	s, n := setupNode("TestNodeServiceStartBeforeSystemdService")

	n.Start() // node service in turn starts systemd service. so, the nodes services should come up fine
	s.Start()

	checkAllNodeServiceStarted(t, s)

	// should not crash
	n.FileBeatConfig([]string{"dummy"})
	n.ElasticDiscoveryConfig([]string{"dummy"})
	n.ElasticMgmtConfig()
	n.KubeletConfig("dummy")
}
