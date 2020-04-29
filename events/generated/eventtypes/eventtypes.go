// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package eventtypes is a auto generated package.
Input file: github.com/pensando/sw/events/protos/eventtypes.proto
*/

package eventtypes

import "sort"

// EventTypeAttributes attributes belong to each event type
type EventTypeAttributes struct {
	EType      string
	Category   string
	Severity   string
	Desc       string
	SuppressMM bool
}

// EventDetails contains all the attributes of an events
type EventDetails struct {
	Name     string `json:name`
	Severity string `json:severity`
	Desc     string `json:desc`
}

// map containing the list of all event types and it's associated attributes
var eventTypes map[EventType]*EventTypeAttributes

// map containing the list of all events grouped by category
var eventTypesByCategory map[string][]*EventDetails

func init() {
	eventTypes = map[EventType]*EventTypeAttributes{}

	eventTypes[ELECTION_STARTED] = &EventTypeAttributes{
		EType:      ELECTION_STARTED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Leader election started in the cluster",
		SuppressMM: true}

	eventTypes[ELECTION_CANCELLED] = &EventTypeAttributes{
		EType:      ELECTION_CANCELLED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Leader election cancelled",
		SuppressMM: true}

	eventTypes[ELECTION_NOTIFICATION_FAILED] = &EventTypeAttributes{
		EType:      ELECTION_NOTIFICATION_FAILED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Failed to send leader election notification",
		SuppressMM: true}

	eventTypes[ELECTION_STOPPED] = &EventTypeAttributes{
		EType:      ELECTION_STOPPED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Leader election stopped",
		SuppressMM: true}

	eventTypes[LEADER_ELECTED] = &EventTypeAttributes{
		EType:      LEADER_ELECTED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Leader elected for the cluster",
		SuppressMM: true}

	eventTypes[LEADER_LOST] = &EventTypeAttributes{
		EType:      LEADER_LOST.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Node lost leadership during the election",
		SuppressMM: true}

	eventTypes[LEADER_CHANGED] = &EventTypeAttributes{
		EType:      LEADER_CHANGED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Leader changed in the election",
		SuppressMM: true}

	eventTypes[NODE_JOINED] = &EventTypeAttributes{
		EType:      NODE_JOINED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Node joined the cluster",
		SuppressMM: true}

	eventTypes[NODE_DISJOINED] = &EventTypeAttributes{
		EType:      NODE_DISJOINED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Node disjoined from the cluster",
		SuppressMM: false}

	eventTypes[NODE_HEALTHY] = &EventTypeAttributes{
		EType:      NODE_HEALTHY.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Node is healthy",
		SuppressMM: true}

	eventTypes[NODE_UNREACHABLE] = &EventTypeAttributes{
		EType:      NODE_UNREACHABLE.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "Node is unreachable",
		SuppressMM: true}

	eventTypes[CLOCK_SYNC_FAILED] = &EventTypeAttributes{
		EType:      CLOCK_SYNC_FAILED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Node failed to synchronize clock",
		SuppressMM: true}

	eventTypes[QUORUM_MEMBER_ADD] = &EventTypeAttributes{
		EType:      QUORUM_MEMBER_ADD.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Added member to quorum",
		SuppressMM: true}

	eventTypes[QUORUM_MEMBER_REMOVE] = &EventTypeAttributes{
		EType:      QUORUM_MEMBER_REMOVE.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Removed member from quorum",
		SuppressMM: true}

	eventTypes[QUORUM_MEMBER_HEALTHY] = &EventTypeAttributes{
		EType:      QUORUM_MEMBER_HEALTHY.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Quorum member is now healthy",
		SuppressMM: true}

	eventTypes[QUORUM_MEMBER_UNHEALTHY] = &EventTypeAttributes{
		EType:      QUORUM_MEMBER_UNHEALTHY.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "Quorum member is now unhealthy",
		SuppressMM: true}

	eventTypes[UNSUPPORTED_QUORUM_SIZE] = &EventTypeAttributes{
		EType:      UNSUPPORTED_QUORUM_SIZE.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "Quorum size is below supported minimum",
		SuppressMM: true}

	eventTypes[QUORUM_UNHEALTHY] = &EventTypeAttributes{
		EType:      QUORUM_UNHEALTHY.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "Quorum does not have enough healthy members",
		SuppressMM: true}

	eventTypes[MODULE_CREATION_FAILED] = &EventTypeAttributes{
		EType:      MODULE_CREATION_FAILED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Module creation for diagnostics failed",
		SuppressMM: true}

	eventTypes[CONFIG_RESTORED] = &EventTypeAttributes{
		EType:      CONFIG_RESTORED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "Configuration was restored",
		SuppressMM: true}

	eventTypes[CONFIG_RESTORE_ABORTED] = &EventTypeAttributes{
		EType:      CONFIG_RESTORE_ABORTED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Configuration restore operation was aborted",
		SuppressMM: true}

	eventTypes[CONFIG_RESTORE_FAILED] = &EventTypeAttributes{
		EType:      CONFIG_RESTORE_FAILED.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "Configuration restore operation failed",
		SuppressMM: true}

	eventTypes[DSC_ADMITTED] = &EventTypeAttributes{
		EType:      DSC_ADMITTED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "DSC admitted to the cluster",
		SuppressMM: true}

	eventTypes[DSC_REJECTED] = &EventTypeAttributes{
		EType:      DSC_REJECTED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "DSC admission request was rejected",
		SuppressMM: false}

	eventTypes[DSC_UNREACHABLE] = &EventTypeAttributes{
		EType:      DSC_UNREACHABLE.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "DSC is unreachable",
		SuppressMM: true}

	eventTypes[DSC_HEALTHY] = &EventTypeAttributes{
		EType:      DSC_HEALTHY.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "DSC is healthy",
		SuppressMM: true}

	eventTypes[DSC_UNHEALTHY] = &EventTypeAttributes{
		EType:      DSC_UNHEALTHY.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "DSC is unhealthy",
		SuppressMM: true}

	eventTypes[HOST_DSC_SPEC_CONFLICT] = &EventTypeAttributes{
		EType:      HOST_DSC_SPEC_CONFLICT.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "The system has detected a conflict between the DSC specifications of different Host objects",
		SuppressMM: false}

	eventTypes[DSC_DEADMITTED] = &EventTypeAttributes{
		EType:      DSC_DEADMITTED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "DSC de-admitted from the cluster",
		SuppressMM: false}

	eventTypes[DSC_DECOMMISSIONED] = &EventTypeAttributes{
		EType:      DSC_DECOMMISSIONED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "DSC decommissioned from the cluster",
		SuppressMM: false}

	eventTypes[AUTO_GENERATED_TLS_CERT] = &EventTypeAttributes{
		EType:      AUTO_GENERATED_TLS_CERT.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Auto generated certificate is being used for API Gateway TLS",
		SuppressMM: true}

	eventTypes[LOGIN_FAILED] = &EventTypeAttributes{
		EType:      LOGIN_FAILED.String(),
		Severity:   "info",
		Category:   "cluster",
		Desc:       "User login failed",
		SuppressMM: false}

	eventTypes[AUDITING_FAILED] = &EventTypeAttributes{
		EType:      AUDITING_FAILED.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "Writing of AuditEvent failed",
		SuppressMM: true}

	eventTypes[PASSWORD_CHANGED] = &EventTypeAttributes{
		EType:      PASSWORD_CHANGED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Password changed",
		SuppressMM: false}

	eventTypes[PASSWORD_RESET] = &EventTypeAttributes{
		EType:      PASSWORD_RESET.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Password reset",
		SuppressMM: false}

	eventTypes[LINK_UP] = &EventTypeAttributes{
		EType:      LINK_UP.String(),
		Severity:   "info",
		Category:   "network",
		Desc:       "Port is linked up",
		SuppressMM: true}

	eventTypes[LINK_DOWN] = &EventTypeAttributes{
		EType:      LINK_DOWN.String(),
		Severity:   "warn",
		Category:   "network",
		Desc:       "Port link status is down",
		SuppressMM: true}

	eventTypes[SERVICE_STARTED] = &EventTypeAttributes{
		EType:      SERVICE_STARTED.String(),
		Severity:   "debug",
		Category:   "system",
		Desc:       "Service started",
		SuppressMM: true}

	eventTypes[SERVICE_STOPPED] = &EventTypeAttributes{
		EType:      SERVICE_STOPPED.String(),
		Severity:   "warn",
		Category:   "system",
		Desc:       "Service stopped",
		SuppressMM: true}

	eventTypes[NAPLES_SERVICE_STOPPED] = &EventTypeAttributes{
		EType:      NAPLES_SERVICE_STOPPED.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "Naples service stopped",
		SuppressMM: true}

	eventTypes[SERVICE_PENDING] = &EventTypeAttributes{
		EType:      SERVICE_PENDING.String(),
		Severity:   "debug",
		Category:   "system",
		Desc:       "Service pending",
		SuppressMM: false}

	eventTypes[SERVICE_RUNNING] = &EventTypeAttributes{
		EType:      SERVICE_RUNNING.String(),
		Severity:   "debug",
		Category:   "system",
		Desc:       "Service running",
		SuppressMM: true}

	eventTypes[SERVICE_UNRESPONSIVE] = &EventTypeAttributes{
		EType:      SERVICE_UNRESPONSIVE.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "Service unresponsive due to lack of system resources",
		SuppressMM: false}

	eventTypes[SYSTEM_COLDBOOT] = &EventTypeAttributes{
		EType:      SYSTEM_COLDBOOT.String(),
		Severity:   "warn",
		Category:   "system",
		Desc:       "System cold booted",
		SuppressMM: true}

	eventTypes[SYSTEM_RESOURCE_USAGE] = &EventTypeAttributes{
		EType:      SYSTEM_RESOURCE_USAGE.String(),
		Severity:   "warn",
		Category:   "system",
		Desc:       "System resource usage is high",
		SuppressMM: false}

	eventTypes[NAPLES_FATAL_INTERRUPT] = &EventTypeAttributes{
		EType:      NAPLES_FATAL_INTERRUPT.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "Naples has a fatal interrupt",
		SuppressMM: false}

	eventTypes[NAPLES_CATTRIP_INTERRUPT] = &EventTypeAttributes{
		EType:      NAPLES_CATTRIP_INTERRUPT.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "System encountered cattrip resetting system",
		SuppressMM: false}

	eventTypes[NAPLES_OVER_TEMP] = &EventTypeAttributes{
		EType:      NAPLES_OVER_TEMP.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "System temperature is above threshold.",
		SuppressMM: false}

	eventTypes[NAPLES_OVER_TEMP_EXIT] = &EventTypeAttributes{
		EType:      NAPLES_OVER_TEMP_EXIT.String(),
		Severity:   "info",
		Category:   "system",
		Desc:       "System temperature is below threshold.",
		SuppressMM: false}

	eventTypes[NAPLES_PANIC_EVENT] = &EventTypeAttributes{
		EType:      NAPLES_PANIC_EVENT.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "System panic on the previous boot",
		SuppressMM: false}

	eventTypes[NAPLES_POST_DIAG_FAILURE_EVENT] = &EventTypeAttributes{
		EType:      NAPLES_POST_DIAG_FAILURE_EVENT.String(),
		Severity:   "warn",
		Category:   "system",
		Desc:       "System post diag test failed",
		SuppressMM: false}

	eventTypes[NAPLES_INFO_PCIEHEALTH_EVENT] = &EventTypeAttributes{
		EType:      NAPLES_INFO_PCIEHEALTH_EVENT.String(),
		Severity:   "info",
		Category:   "system",
		Desc:       "System has detected a pcie link health event",
		SuppressMM: false}

	eventTypes[NAPLES_WARN_PCIEHEALTH_EVENT] = &EventTypeAttributes{
		EType:      NAPLES_WARN_PCIEHEALTH_EVENT.String(),
		Severity:   "warn",
		Category:   "system",
		Desc:       "System has detected a pcie link health warning",
		SuppressMM: false}

	eventTypes[NAPLES_ERR_PCIEHEALTH_EVENT] = &EventTypeAttributes{
		EType:      NAPLES_ERR_PCIEHEALTH_EVENT.String(),
		Severity:   "critical",
		Category:   "system",
		Desc:       "System has detected a pcie link health error",
		SuppressMM: false}

	eventTypes[DISK_THRESHOLD_EXCEEDED] = &EventTypeAttributes{
		EType:      DISK_THRESHOLD_EXCEEDED.String(),
		Severity:   "critical",
		Category:   "",
		Desc:       "Disk threshold exceeded",
		SuppressMM: false}

	eventTypes[ROLLOUT_STARTED] = &EventTypeAttributes{
		EType:      ROLLOUT_STARTED.String(),
		Severity:   "info",
		Category:   "rollout",
		Desc:       "Rollout initiated",
		SuppressMM: false}

	eventTypes[ROLLOUT_SUCCESS] = &EventTypeAttributes{
		EType:      ROLLOUT_SUCCESS.String(),
		Severity:   "info",
		Category:   "rollout",
		Desc:       "Rollout successfully completed",
		SuppressMM: false}

	eventTypes[ROLLOUT_FAILED] = &EventTypeAttributes{
		EType:      ROLLOUT_FAILED.String(),
		Severity:   "critical",
		Category:   "rollout",
		Desc:       "Rollout failed",
		SuppressMM: false}

	eventTypes[ROLLOUT_SUSPENDED] = &EventTypeAttributes{
		EType:      ROLLOUT_SUSPENDED.String(),
		Severity:   "info",
		Category:   "rollout",
		Desc:       "Rollout suspended",
		SuppressMM: false}

	eventTypes[CONFIG_FAIL] = &EventTypeAttributes{
		EType:      CONFIG_FAIL.String(),
		Severity:   "warn",
		Category:   "config",
		Desc:       "Configuration failed",
		SuppressMM: false}

	eventTypes[TCP_HALF_OPEN_SESSION_LIMIT_APPROACH] = &EventTypeAttributes{
		EType:      TCP_HALF_OPEN_SESSION_LIMIT_APPROACH.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "TCP Session Limit approaching",
		SuppressMM: false}

	eventTypes[TCP_HALF_OPEN_SESSION_LIMIT_REACHED] = &EventTypeAttributes{
		EType:      TCP_HALF_OPEN_SESSION_LIMIT_REACHED.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "TCP Session Limit reached",
		SuppressMM: false}

	eventTypes[UDP_ACTIVE_SESSION_LIMIT_APPROACH] = &EventTypeAttributes{
		EType:      UDP_ACTIVE_SESSION_LIMIT_APPROACH.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "UDP Session Limit approaching",
		SuppressMM: false}

	eventTypes[UDP_ACTIVE_SESSION_LIMIT_REACHED] = &EventTypeAttributes{
		EType:      UDP_ACTIVE_SESSION_LIMIT_REACHED.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "UDP Session Limit reached",
		SuppressMM: false}

	eventTypes[ICMP_ACTIVE_SESSION_LIMIT_APPROACH] = &EventTypeAttributes{
		EType:      ICMP_ACTIVE_SESSION_LIMIT_APPROACH.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "ICMP Session Limit approaching",
		SuppressMM: false}

	eventTypes[ICMP_ACTIVE_SESSION_LIMIT_REACHED] = &EventTypeAttributes{
		EType:      ICMP_ACTIVE_SESSION_LIMIT_REACHED.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "ICMP Session Limit reached",
		SuppressMM: false}

	eventTypes[OTHER_ACTIVE_SESSION_LIMIT_APPROACH] = &EventTypeAttributes{
		EType:      OTHER_ACTIVE_SESSION_LIMIT_APPROACH.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "OTHER Session Limit approaching",
		SuppressMM: false}

	eventTypes[OTHER_ACTIVE_SESSION_LIMIT_REACHED] = &EventTypeAttributes{
		EType:      OTHER_ACTIVE_SESSION_LIMIT_REACHED.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "OTHER Session Limit reached",
		SuppressMM: false}

	eventTypes[ORCH_CONNECTION_ERROR] = &EventTypeAttributes{
		EType:      ORCH_CONNECTION_ERROR.String(),
		Severity:   "critical",
		Category:   "",
		Desc:       "Failed to connect to orchestrator",
		SuppressMM: false}

	eventTypes[ORCH_LOGIN_FAILURE] = &EventTypeAttributes{
		EType:      ORCH_LOGIN_FAILURE.String(),
		Severity:   "critical",
		Category:   "",
		Desc:       "Login credentials were invalid for orchestrator.",
		SuppressMM: false}

	eventTypes[ORCH_CONFIG_PUSH_FAILURE] = &EventTypeAttributes{
		EType:      ORCH_CONFIG_PUSH_FAILURE.String(),
		Severity:   "warn",
		Category:   "",
		Desc:       "Failed to push some configurations to orchestrator",
		SuppressMM: false}

	eventTypes[ORCH_INVALID_ACTION] = &EventTypeAttributes{
		EType:      ORCH_INVALID_ACTION.String(),
		Severity:   "warn",
		Category:   "",
		Desc:       "User performed an action that is not supported",
		SuppressMM: false}

	eventTypes[MIGRATION_FAILED] = &EventTypeAttributes{
		EType:      MIGRATION_FAILED.String(),
		Severity:   "warn",
		Category:   "",
		Desc:       "Migration Failed",
		SuppressMM: false}

	eventTypes[MIGRATION_TIMED_OUT] = &EventTypeAttributes{
		EType:      MIGRATION_TIMED_OUT.String(),
		Severity:   "warn",
		Category:   "",
		Desc:       "Migration Timed out",
		SuppressMM: false}

	eventTypes[ORCH_ALREADY_MANAGED] = &EventTypeAttributes{
		EType:      ORCH_ALREADY_MANAGED.String(),
		Severity:   "warn",
		Category:   "",
		Desc:       "Another PSM may be managing the same namespace",
		SuppressMM: false}

	eventTypes[ORCH_UNSUPPORTED_VERSION] = &EventTypeAttributes{
		EType:      ORCH_UNSUPPORTED_VERSION.String(),
		Severity:   "critical",
		Category:   "",
		Desc:       "Unsupported orchestrator version",
		SuppressMM: false}

	eventTypes[ORCH_DSC_NOT_ADMITTED] = &EventTypeAttributes{
		EType:      ORCH_DSC_NOT_ADMITTED.String(),
		Severity:   "warn",
		Category:   "cluster",
		Desc:       "Workload provisioned on a host with no admitted DSC",
		SuppressMM: false}

	eventTypes[ORCH_DSC_MODE_INCOMPATIBLE] = &EventTypeAttributes{
		EType:      ORCH_DSC_MODE_INCOMPATIBLE.String(),
		Severity:   "critical",
		Category:   "cluster",
		Desc:       "DSC mode is incompatible.",
		SuppressMM: false}

	eventTypes[COLLECTOR_REACHABLE] = &EventTypeAttributes{
		EType:      COLLECTOR_REACHABLE.String(),
		Severity:   "info",
		Category:   "network",
		Desc:       "Collector reachable from DSC",
		SuppressMM: true}

	eventTypes[COLLECTOR_UNREACHABLE] = &EventTypeAttributes{
		EType:      COLLECTOR_UNREACHABLE.String(),
		Severity:   "warn",
		Category:   "network",
		Desc:       "Collector not reachable from DSC",
		SuppressMM: false}

	eventTypes[BGP_SESSION_ESTABLISHED] = &EventTypeAttributes{
		EType:      BGP_SESSION_ESTABLISHED.String(),
		Severity:   "info",
		Category:   "network",
		Desc:       "BGP session is established",
		SuppressMM: true}

	eventTypes[BGP_SESSION_DOWN] = &EventTypeAttributes{
		EType:      BGP_SESSION_DOWN.String(),
		Severity:   "warn",
		Category:   "network",
		Desc:       "BGP session is down",
		SuppressMM: true}

}

func GetEventTypeAttrs(eType EventType) *EventTypeAttributes {
	if attrs, ok := eventTypes[eType]; ok {
		return attrs
	}

	return nil
}

func SuppressDuringMaintenance(eTypeStr string) bool {
	if eType, found := EventType_value[eTypeStr]; found {
		if attrs, ok := eventTypes[EventType(eType)]; ok {
			return attrs.SuppressMM
		}
	}
	return false
}

func GetEventsByCategory() map[string][]*EventDetails {
	var eTypes []string
	for _, eType := range eventTypes {
		eTypes = append(eTypes, eType.EType)
	}
	sort.Strings(eTypes)

	eventTypesByCategory = map[string][]*EventDetails{}
	for _, eTypeStr := range eTypes {
		eType, ok := EventType_value[eTypeStr]
		if ok {
			eAttrs := eventTypes[EventType(eType)]
			eventTypesByCategory[eAttrs.Category] = append(eventTypesByCategory[eAttrs.Category],
				&EventDetails{Name: eAttrs.EType, Desc: eAttrs.Desc, Severity: eAttrs.Severity})
		}
	}

	return eventTypesByCategory
}
