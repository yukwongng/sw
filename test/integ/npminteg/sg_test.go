// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package npminteg

import (
	"fmt"

	. "gopkg.in/check.v1"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/security"
	"github.com/pensando/sw/api/labels"
	. "github.com/pensando/sw/venice/utils/testutils"
)

// TestNpmSgCreateDelete
func (it *integTestSuite) TestNpmSgCreateDelete(c *C) {
	// create sg in watcher
	err := it.ctrler.Watchr.CreateSecurityGroup("default", "default", "testsg", labels.SelectorFromSet(labels.Set{"env": "production", "app": "procurement"}))
	c.Assert(err, IsNil)

	// verify all agents have the security group
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			return len(ag.datapath.DB.SgDB) == 1, nil
		}, "Sg not found on agent", "10ms", it.pollTimeout())
		c.Assert(len(ag.datapath.DB.SgDB[fmt.Sprintf("%s|%s", "default", "testsg")].Request), Equals, 1)
	}

	// incoming rule
	rules := []*security.SGRule{
		{
			Apps:   []string{"tcp/80"},
			Action: "PERMIT",
		},
	}

	// create sg policy
	err = it.ctrler.Watchr.CreateSgpolicy("default", "default", "testpolicy", true, []string{}, rules)
	c.Assert(err, IsNil)

	// construct object meta
	policyMeta := api.ObjectMeta{
		Tenant:    "default",
		Namespace: "default",
		Name:      "testpolicy",
	}

	// verify agent state has the policy has the rules
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			_, err := ag.nagent.NetworkAgent.FindSGPolicy(policyMeta)
			if err != nil {
				return false, nil
			}
			return true, nil
		}, fmt.Sprintf("Sg rules not found on agent. DB: %v", ag.datapath.DB.SgPolicyDB), "10ms", it.pollTimeout())
	}

	// delete the sg policy
	err = it.ctrler.Watchr.DeleteSgpolicy("default", "default", "testpolicy")
	c.Assert(err, IsNil)

	// verify rules are gone from agent
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			_, err := ag.nagent.NetworkAgent.FindSGPolicy(policyMeta)
			if err == nil {
				return false, nil
			}
			return true, nil
		}, "Sg rules still found on agent", "10ms", it.pollTimeout())
	}
	// delete the security group
	err = it.ctrler.Watchr.DeleteSecurityGroup("default", "testsg")
	c.Assert(err, IsNil)

	// verify sg is removed from datapath
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			return (len(ag.datapath.DB.SgDB) == 0), nil
		}, "Sg still found on agent", "10ms", it.pollTimeout())
	}
}

func (it *integTestSuite) TestNpmSgEndpointAttach(c *C) {
	// create a network in controller
	err := it.ctrler.Watchr.CreateNetwork("default", "default", "testNetwork", "10.1.0.0/16", "10.1.1.254")
	c.Assert(err, IsNil)
	AssertEventually(c, func() (bool, interface{}) {
		_, nerr := it.ctrler.StateMgr.FindNetwork("default", "testNetwork")
		return (nerr == nil), nil
	}, "Network not found in statemgr")

	// create sg in watcher
	err = it.ctrler.Watchr.CreateSecurityGroup("default", "default", "testsg", labels.SelectorFromSet(labels.Set{"env": "production", "app": "procurement"}))
	c.Assert(err, IsNil)
	AssertEventually(c, func() (bool, interface{}) {
		_, serr := it.ctrler.StateMgr.FindSecurityGroup("default", "testsg")
		return (serr == nil), nil
	}, "Sg not found in statemgr")

	// create endpoint
	err = it.ctrler.Watchr.CreateEndpoint("default", "default", "testNetwork", "testEndpoint1", "testVm1", "01:01:01:01:01:01", "host1", "20.1.1.1", map[string]string{"env": "production", "app": "procurement"}, 2)
	c.Assert(err, IsNil)

	// verify endpoint is present in all agents
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			return (len(ag.datapath.DB.EndpointDB) == 1), nil
		}, "endpoint not found on agent", "10ms", it.pollTimeout())
		ep, ok := ag.datapath.DB.EndpointDB[fmt.Sprintf("%s|%s", "default", "testEndpoint1")]
		c.Assert(ok, Equals, true)
		c.Assert(len(ep.Request), Equals, 1)
		c.Assert(len(ep.Request[0].EndpointAttrs.SgKeyHandle), Equals, 1)
		c.Assert(ep.Request[0].EndpointAttrs.UsegVlan, Equals, uint32(2))
	}

	// create second endpoint
	err = it.ctrler.Watchr.CreateEndpoint("default", "default", "testNetwork", "testEndpoint2", "testVm2", "02:02:02:02:02:02", "host2", "20.2.2.2", map[string]string{"env": "production", "app": "procurement"}, 3)
	c.Assert(err, IsNil)

	// verify new endpoint is present in all agents
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			return (len(ag.datapath.DB.EndpointDB) == 2), nil
		}, "endpoint not found on agent", "10ms", it.pollTimeout())
		ep, ok := ag.datapath.DB.EndpointDB[fmt.Sprintf("%s|%s", "default", "testEndpoint2")]
		c.Assert(ok, Equals, true)
		c.Assert(len(ep.Request), Equals, 1)
		c.Assert(len(ep.Request[0].EndpointAttrs.SgKeyHandle), Equals, 1)
		c.Assert(ep.Request[0].EndpointAttrs.UsegVlan, Equals, uint32(3))
	}

	// delete the second endpoint
	err = it.ctrler.Watchr.DeleteEndpoint("default", "testNetwork", "testEndpoint2", "testVm2", "02:02:02:02:02:02", "host2", "20.2.2.2")
	c.Assert(err, IsNil)
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			return (len(ag.datapath.DB.EndpointDB) == 1), nil
		}, "endpoint still found on agent", "10ms", it.pollTimeout())
	}

	// delete the security group
	err = it.ctrler.Watchr.DeleteSecurityGroup("default", "testsg")
	c.Assert(err, IsNil)

	// verify sg is removed from the endpoint
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			ep, ok := ag.datapath.DB.EndpointUpdateDB[fmt.Sprintf("%s|%s", "default", "testEndpoint1")]
			if ok && len(ep.Request) == 1 && len(ep.Request[0].EndpointAttrs.SgKeyHandle) == 0 {
				return true, nil
			}
			return false, nil
		}, "endpoint still found on agent", "10ms", it.pollTimeout())
	}

	// delete endpoint
	err = it.ctrler.Watchr.DeleteEndpoint("default", "testNetwork", "testEndpoint1", "testVm1", "01:01:01:01:01:01", "host1", "20.1.1.1")
	c.Assert(err, IsNil)
	for _, ag := range it.agents {
		AssertEventually(c, func() (bool, interface{}) {
			return (len(ag.datapath.DB.EndpointDB) == 0), nil
		}, "endpoint still found on agent", "10ms", it.pollTimeout())
	}

	// delete the network
	err = it.ctrler.Watchr.DeleteNetwork("default", "testNetwork")
	c.Assert(err, IsNil)
	AssertEventually(c, func() (bool, interface{}) {
		_, nerr := it.ctrler.StateMgr.FindNetwork("default", "testNetwork")
		return (nerr != nil), nil
	}, "endpoint still found on agent", "10ms", it.pollTimeout())
}
