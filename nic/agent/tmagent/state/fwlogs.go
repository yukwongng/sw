package state

import (
	"fmt"
	"strings"
	"time"

	"github.com/pensando/sw/venice/utils/netutils"

	halproto "github.com/pensando/sw/nic/agent/dscagent/types/irisproto"
	"github.com/pensando/sw/nic/agent/ipc"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"
	"github.com/pensando/sw/venice/utils/tsdb"
)

// FwlogIpcShm is the path to fw logs in shared memory
const FwlogIpcShm = "/fwlog_ipc_shm"

var flowDirectionName map[uint32]string

type fwevent struct {
	*halproto.FWEvent
}

// String show json string
func (ev *fwevent) String() string {
	return fmt.Sprintf("[{"+
		"\"time\":\"%v\",\"destaddr\":\"%v\",\"destport\":%v,\"srcaddr\":\"%v\",\"srcport\":%v,\"protocol\":\"%v\",\"action\":\"%v\","+
		"\"direction\":\"%v\",\"rule-id\":%v,\"session-id\":%v,\"session-state\":\"%v\",\"app-id\":\"%v\"}]",
		time.Unix(0, ev.GetTimestamp()).Format(time.RFC3339),
		netutils.IPv4Uint32ToString(ev.GetDipv4()),
		ev.GetDport(),
		netutils.IPv4Uint32ToString(ev.GetSipv4()),
		ev.GetSport(),
		strings.TrimPrefix(ev.GetIpProt().String(), "IPPROTO_"),
		strings.ToLower(strings.TrimPrefix(ev.GetFwaction().String(), "SECURITY_RULE_ACTION_")),
		flowDirectionName[ev.GetDirection()],
		ev.GetRuleId(),
		ev.GetSessionId(),
		strings.ToLower(strings.Replace(halproto.FlowLogEventType_name[int32(ev.GetFlowaction())], "LOG_EVENT_TYPE_", "", 1)),
		ev.GetAppId())
}

// TsdbInit initilaizes tsdb and fwlog object
func (s *PolicyState) TsdbInit(nodeUUID string, rc resolver.Interface) error {
	opts := &tsdb.Opts{
		ClientName:              nodeUUID,
		Collector:               globals.Collector,
		ResolverClient:          rc,
		DBName:                  "default",
		SendInterval:            time.Duration(30) * time.Second,
		ConnectionRetryInterval: 100 * time.Millisecond,
	}

	// Init the TSDB
	tsdb.Init(s.ctx, opts)

	fwTable, err := tsdb.NewObj("Fwlogs", map[string]string{}, nil, &tsdb.ObjOpts{})
	if err != nil {
		return err
	}
	s.fwTable = fwTable

	return nil
}

// TsdbCleanup free up tsdb resources
func (s *PolicyState) TsdbCleanup() {
	tsdb.Cleanup()
}

// FwlogInit starts processing fwlogs
func (s *PolicyState) FwlogInit(path string) error {
	mSize := int(ipc.GetSharedConstant("IPC_MEM_SIZE"))
	instCount := int(ipc.GetSharedConstant("IPC_INSTANCES"))
	log.Infof("init fwlog shared memory, total memory: %d instances: %d", mSize, instCount)
	shm, err := ipc.NewSharedMem(mSize, instCount, path)
	if err != nil {
		return err
	}
	s.shm = shm
	log.Infof("allocated shared memory: %v", shm)

	flowDirectionName = make(map[uint32]string, len(halproto.FlowDirection_name))
	// create mapping
	//    FLOW_DIRECTION_FROM_HOST -> from-host
	//    FLOW_DIRECTION_TO_HOST -> to-host
	for k, d := range halproto.FlowDirection_name {
		flowDirectionName[uint32(k)] = strings.ToLower(strings.Replace(strings.TrimPrefix(d, "FLOW_DIRECTION_"), "_", "-", 1))
	}

	for ix := 0; ix < instCount; ix++ {
		ipc := shm.IPCInstance()
		log.Infof("IPC[%d] %s", ix, ipc)
		s.ipc = append(s.ipc, ipc)
		s.wg.Add(1)
		go func(ix int) {
			defer s.wg.Done()
			log.Infof("start ProcessFWEvent_%d()", ix)
			ipc.Receive(s.ctx, s.ProcessFWEvent)
			log.Infof("exit ProcessFWEvent_%d()", ix)
		}(ix)
	}

	return nil
}

func (s *PolicyState) isAtleastOneFwLogCollectorPresent() bool {
	// Doing it this way becuase sync.Map does not provide Len method and
	// the regular len method does not support sync.Map type.
	present := false
	s.fwLogCollectors.Range(func(k interface{}, v interface{}) bool {
		present = true
		return true
	})
	return present
}

// ProcessFWEvent process fwlog event received from ipc
func (s *PolicyState) ProcessFWEvent(ev *halproto.FWEvent, ts time.Time) {
	// return if no collectors are present
	if !s.isAtleastOneFwLogCollectorPresent() {
		return
	}

	fwev := &fwevent{FWEvent: ev}

	vrfSrc := fmt.Sprintf("%v", ev.GetSourceVrf())
	vrfDest := fmt.Sprintf("%v", ev.GetDestVrf())
	ipSrc := netutils.IPv4Uint32ToString(ev.GetSipv4())
	ipDest := netutils.IPv4Uint32ToString(ev.GetDipv4())
	dPort := fmt.Sprintf("%v", ev.GetDport())
	sPort := fmt.Sprintf("%v", ev.GetSport())
	ipProt := fmt.Sprintf("%v", strings.TrimPrefix(ev.GetIpProt().String(), "IPPROTO_"))
	action := fmt.Sprintf("%v", strings.ToLower(strings.TrimPrefix(ev.GetFwaction().String(), "SECURITY_RULE_ACTION_")))
	dir := flowDirectionName[ev.GetDirection()]
	ruleID := fmt.Sprintf("%v", ev.GetRuleId())
	sessionID := fmt.Sprintf("%v", ev.GetSessionId())
	state := strings.ToLower(strings.Replace(halproto.FlowLogEventType_name[int32(ev.GetFlowaction())], "LOG_EVENT_TYPE_", "", 1))
	alg := fmt.Sprintf("%v", strings.TrimPrefix(ev.GetAlg().String(), "APP_SVC_"))
	icmpType := fmt.Sprintf("%v", int64(ev.GetIcmptype()))
	icmpID := fmt.Sprintf("%v", int64(ev.GetIcmpid()))
	icmpCode := fmt.Sprintf("%v", int64(ev.GetIcmpcode()))
	appID := fmt.Sprintf("%v", ev.GetAppId()) // TODO: praveen convert to enum
	iflowBytes := fmt.Sprintf("%v", ev.GetIflowBytes())
	rflowBytes := fmt.Sprintf("%v", ev.GetRflowBytes())
	unixnano := ev.GetTimestamp()
	if unixnano != 0 {
		// if a timestamp was specified in the msg, use it
		ts = time.Unix(0, unixnano)
	}

	// CSV file format
	// Since no aggregation is done as fo now, just report count=1 for every log.
	count := "1"

	// CSV file format
	fwLog := []string{
		vrfSrc,
		vrfDest,
		ipSrc,
		ipDest,
		ts.Format(time.RFC3339),
		sPort,
		dPort,
		ipProt,
		action,
		dir,
		ruleID,
		sessionID,
		state,
		icmpType,
		icmpID,
		icmpCode,
		appID,
		alg,
		iflowBytes,
		rflowBytes,
		count,
	}

	// set src/dest vrf
	vrfList := map[uint64]bool{
		ev.SourceVrf: true,
		ev.DestVrf:   true,
	}

	// used for sending the log only once to PSM
	logSentToPSM := false

	// check dest/src vrf
	s.fwLogCollectors.Range(func(k interface{}, v interface{}) bool {
		if col, ok := v.(*syslogFwlogCollector); ok {
			col.Lock()
			if _, ok := vrfList[col.vrf]; ok {
				if col.syslogFd != nil && col.filter&(1<<uint32(ev.Fwaction)) != 0 {
					s.sendFwLog(col, fwev)
				}
			} else {
				log.Errorf("invalid collector")
			}
			col.Unlock()
		} else if col, ok := v.(*psmFwlogCollector); ok && !logSentToPSM {
			if _, ok := vrfList[col.vrf]; ok && col.filter&(1<<uint32(ev.Fwaction)) != 0 {
				// TODO: use sync pool
				s.logsChannel <- singleLog{ts, ev.GetSourceVrf(), fwLog}
				logSentToPSM = true
			}
		}

		return true
	})
}
