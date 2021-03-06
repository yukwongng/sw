package metrics

import (
	"fmt"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/netutils"
)

type macMetricsXlate struct {
	portToInterfaceMap map[uint32]string // maintains portID to name
}

func newMacMetricsTranslator() *macMetricsXlate {
	m := &macMetricsXlate{
		portToInterfaceMap: make(map[uint32]string),
	}
	return m
}

// KeyToMeta converts network key to meta
func (n *macMetricsXlate) KeyToMeta(key interface{}) *api.ObjectMeta {
	portID, ok := key.(uint32)
	if ok {
		// Check cache.
		_, found := n.portToInterfaceMap[portID]
		if !found {
			// build cache
			reqURL := fmt.Sprintf("http://127.0.0.1:%s/api/interfaces/", globals.AgentRESTPort)
			var interfaces []netproto.Interface
			if err := netutils.HTTPGet(reqURL, &interfaces); err != nil {
				log.Errorf("Failed to Get Interfaces from Agent. Err: %v", err)
			}

			for _, intf := range interfaces {
				if intf.Spec.Type == "UPLINK_ETH" || intf.Spec.Type == "UPLINK_MGMT" {
					if intf.Status.InterfaceID == uint64(portID) {
						n.portToInterfaceMap[portID] = intf.GetName()

					}
				}
			}
		}

		intfName, ok := n.portToInterfaceMap[portID]
		if !ok {
			log.Errorf("Failed to find the interface with port ID %d in DB: %v", portID, n.portToInterfaceMap)
			return nil
		}

		return &api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      intfName,
		}

	}
	return nil
}

// MetaToKey converts meta to network key
func (n *macMetricsXlate) MetaToKey(meta *api.ObjectMeta) interface{} {
	return meta.Name
}
