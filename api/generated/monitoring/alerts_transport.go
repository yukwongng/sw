// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoring is a auto generated package.
Input file: alerts.proto
*/
package monitoring

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pensando/sw/api"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

func encodeHTTPAlert(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlert(_ context.Context, r *http.Request) (interface{}, error) {
	var req Alert
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlert encodes GRPC request
func EncodeGrpcReqAlert(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*Alert)
	return req, nil
}

// DecodeGrpcReqAlert decodes GRPC request
func DecodeGrpcReqAlert(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*Alert)
	return req, nil
}

// EncodeGrpcRespAlert encodes GRC response
func EncodeGrpcRespAlert(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlert decodes GRPC response
func DecodeGrpcRespAlert(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertDestination(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertDestination(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertDestination
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertDestination encodes GRPC request
func EncodeGrpcReqAlertDestination(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertDestination)
	return req, nil
}

// DecodeGrpcReqAlertDestination decodes GRPC request
func DecodeGrpcReqAlertDestination(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertDestination)
	return req, nil
}

// EncodeGrpcRespAlertDestination encodes GRC response
func EncodeGrpcRespAlertDestination(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertDestination decodes GRPC response
func DecodeGrpcRespAlertDestination(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertDestinationSpec(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertDestinationSpec(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertDestinationSpec
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertDestinationSpec encodes GRPC request
func EncodeGrpcReqAlertDestinationSpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertDestinationSpec)
	return req, nil
}

// DecodeGrpcReqAlertDestinationSpec decodes GRPC request
func DecodeGrpcReqAlertDestinationSpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertDestinationSpec)
	return req, nil
}

// EncodeGrpcRespAlertDestinationSpec encodes GRC response
func EncodeGrpcRespAlertDestinationSpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertDestinationSpec decodes GRPC response
func DecodeGrpcRespAlertDestinationSpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertDestinationStatus(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertDestinationStatus(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertDestinationStatus
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertDestinationStatus encodes GRPC request
func EncodeGrpcReqAlertDestinationStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertDestinationStatus)
	return req, nil
}

// DecodeGrpcReqAlertDestinationStatus decodes GRPC request
func DecodeGrpcReqAlertDestinationStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertDestinationStatus)
	return req, nil
}

// EncodeGrpcRespAlertDestinationStatus encodes GRC response
func EncodeGrpcRespAlertDestinationStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertDestinationStatus decodes GRPC response
func DecodeGrpcRespAlertDestinationStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertPolicy(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertPolicy(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertPolicy
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertPolicy encodes GRPC request
func EncodeGrpcReqAlertPolicy(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertPolicy)
	return req, nil
}

// DecodeGrpcReqAlertPolicy decodes GRPC request
func DecodeGrpcReqAlertPolicy(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertPolicy)
	return req, nil
}

// EncodeGrpcRespAlertPolicy encodes GRC response
func EncodeGrpcRespAlertPolicy(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertPolicy decodes GRPC response
func DecodeGrpcRespAlertPolicy(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertPolicySpec(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertPolicySpec(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertPolicySpec
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertPolicySpec encodes GRPC request
func EncodeGrpcReqAlertPolicySpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertPolicySpec)
	return req, nil
}

// DecodeGrpcReqAlertPolicySpec decodes GRPC request
func DecodeGrpcReqAlertPolicySpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertPolicySpec)
	return req, nil
}

// EncodeGrpcRespAlertPolicySpec encodes GRC response
func EncodeGrpcRespAlertPolicySpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertPolicySpec decodes GRPC response
func DecodeGrpcRespAlertPolicySpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertPolicyStatus(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertPolicyStatus(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertPolicyStatus
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertPolicyStatus encodes GRPC request
func EncodeGrpcReqAlertPolicyStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertPolicyStatus)
	return req, nil
}

// DecodeGrpcReqAlertPolicyStatus decodes GRPC request
func DecodeGrpcReqAlertPolicyStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertPolicyStatus)
	return req, nil
}

// EncodeGrpcRespAlertPolicyStatus encodes GRC response
func EncodeGrpcRespAlertPolicyStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertPolicyStatus decodes GRPC response
func DecodeGrpcRespAlertPolicyStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertReason(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertReason(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertReason
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertReason encodes GRPC request
func EncodeGrpcReqAlertReason(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertReason)
	return req, nil
}

// DecodeGrpcReqAlertReason decodes GRPC request
func DecodeGrpcReqAlertReason(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertReason)
	return req, nil
}

// EncodeGrpcRespAlertReason encodes GRC response
func EncodeGrpcRespAlertReason(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertReason decodes GRPC response
func DecodeGrpcRespAlertReason(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertSource(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertSource(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertSource
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertSource encodes GRPC request
func EncodeGrpcReqAlertSource(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertSource)
	return req, nil
}

// DecodeGrpcReqAlertSource decodes GRPC request
func DecodeGrpcReqAlertSource(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertSource)
	return req, nil
}

// EncodeGrpcRespAlertSource encodes GRC response
func EncodeGrpcRespAlertSource(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertSource decodes GRPC response
func DecodeGrpcRespAlertSource(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertSpec(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertSpec(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertSpec
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertSpec encodes GRPC request
func EncodeGrpcReqAlertSpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertSpec)
	return req, nil
}

// DecodeGrpcReqAlertSpec decodes GRPC request
func DecodeGrpcReqAlertSpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertSpec)
	return req, nil
}

// EncodeGrpcRespAlertSpec encodes GRC response
func EncodeGrpcRespAlertSpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertSpec decodes GRPC response
func DecodeGrpcRespAlertSpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAlertStatus(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAlertStatus(_ context.Context, r *http.Request) (interface{}, error) {
	var req AlertStatus
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAlertStatus encodes GRPC request
func EncodeGrpcReqAlertStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertStatus)
	return req, nil
}

// DecodeGrpcReqAlertStatus decodes GRPC request
func DecodeGrpcReqAlertStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AlertStatus)
	return req, nil
}

// EncodeGrpcRespAlertStatus encodes GRC response
func EncodeGrpcRespAlertStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAlertStatus decodes GRPC response
func DecodeGrpcRespAlertStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPAuditInfo(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPAuditInfo(_ context.Context, r *http.Request) (interface{}, error) {
	var req AuditInfo
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqAuditInfo encodes GRPC request
func EncodeGrpcReqAuditInfo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AuditInfo)
	return req, nil
}

// DecodeGrpcReqAuditInfo decodes GRPC request
func DecodeGrpcReqAuditInfo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*AuditInfo)
	return req, nil
}

// EncodeGrpcRespAuditInfo encodes GRC response
func EncodeGrpcRespAuditInfo(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespAuditInfo decodes GRPC response
func DecodeGrpcRespAuditInfo(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPEmailExport(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPEmailExport(_ context.Context, r *http.Request) (interface{}, error) {
	var req EmailExport
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqEmailExport encodes GRPC request
func EncodeGrpcReqEmailExport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*EmailExport)
	return req, nil
}

// DecodeGrpcReqEmailExport decodes GRPC request
func DecodeGrpcReqEmailExport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*EmailExport)
	return req, nil
}

// EncodeGrpcRespEmailExport encodes GRC response
func EncodeGrpcRespEmailExport(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespEmailExport decodes GRPC response
func DecodeGrpcRespEmailExport(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPMatchedRequirement(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPMatchedRequirement(_ context.Context, r *http.Request) (interface{}, error) {
	var req MatchedRequirement
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqMatchedRequirement encodes GRPC request
func EncodeGrpcReqMatchedRequirement(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*MatchedRequirement)
	return req, nil
}

// DecodeGrpcReqMatchedRequirement decodes GRPC request
func DecodeGrpcReqMatchedRequirement(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*MatchedRequirement)
	return req, nil
}

// EncodeGrpcRespMatchedRequirement encodes GRC response
func EncodeGrpcRespMatchedRequirement(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespMatchedRequirement decodes GRPC response
func DecodeGrpcRespMatchedRequirement(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPMeasurementCriteria(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPMeasurementCriteria(_ context.Context, r *http.Request) (interface{}, error) {
	var req MeasurementCriteria
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqMeasurementCriteria encodes GRPC request
func EncodeGrpcReqMeasurementCriteria(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*MeasurementCriteria)
	return req, nil
}

// DecodeGrpcReqMeasurementCriteria decodes GRPC request
func DecodeGrpcReqMeasurementCriteria(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*MeasurementCriteria)
	return req, nil
}

// EncodeGrpcRespMeasurementCriteria encodes GRC response
func EncodeGrpcRespMeasurementCriteria(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespMeasurementCriteria decodes GRPC response
func DecodeGrpcRespMeasurementCriteria(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPMetricIdentifier(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPMetricIdentifier(_ context.Context, r *http.Request) (interface{}, error) {
	var req MetricIdentifier
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqMetricIdentifier encodes GRPC request
func EncodeGrpcReqMetricIdentifier(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*MetricIdentifier)
	return req, nil
}

// DecodeGrpcReqMetricIdentifier decodes GRPC request
func DecodeGrpcReqMetricIdentifier(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*MetricIdentifier)
	return req, nil
}

// EncodeGrpcRespMetricIdentifier encodes GRC response
func EncodeGrpcRespMetricIdentifier(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespMetricIdentifier decodes GRPC response
func DecodeGrpcRespMetricIdentifier(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPSNMPExport(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPSNMPExport(_ context.Context, r *http.Request) (interface{}, error) {
	var req SNMPExport
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqSNMPExport encodes GRPC request
func EncodeGrpcReqSNMPExport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*SNMPExport)
	return req, nil
}

// DecodeGrpcReqSNMPExport decodes GRPC request
func DecodeGrpcReqSNMPExport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*SNMPExport)
	return req, nil
}

// EncodeGrpcRespSNMPExport encodes GRC response
func EncodeGrpcRespSNMPExport(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespSNMPExport decodes GRPC response
func DecodeGrpcRespSNMPExport(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPStatsAlertPolicy(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPStatsAlertPolicy(_ context.Context, r *http.Request) (interface{}, error) {
	var req StatsAlertPolicy
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqStatsAlertPolicy encodes GRPC request
func EncodeGrpcReqStatsAlertPolicy(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*StatsAlertPolicy)
	return req, nil
}

// DecodeGrpcReqStatsAlertPolicy decodes GRPC request
func DecodeGrpcReqStatsAlertPolicy(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*StatsAlertPolicy)
	return req, nil
}

// EncodeGrpcRespStatsAlertPolicy encodes GRC response
func EncodeGrpcRespStatsAlertPolicy(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespStatsAlertPolicy decodes GRPC response
func DecodeGrpcRespStatsAlertPolicy(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPStatsAlertPolicySpec(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPStatsAlertPolicySpec(_ context.Context, r *http.Request) (interface{}, error) {
	var req StatsAlertPolicySpec
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqStatsAlertPolicySpec encodes GRPC request
func EncodeGrpcReqStatsAlertPolicySpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*StatsAlertPolicySpec)
	return req, nil
}

// DecodeGrpcReqStatsAlertPolicySpec decodes GRPC request
func DecodeGrpcReqStatsAlertPolicySpec(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*StatsAlertPolicySpec)
	return req, nil
}

// EncodeGrpcRespStatsAlertPolicySpec encodes GRC response
func EncodeGrpcRespStatsAlertPolicySpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespStatsAlertPolicySpec decodes GRPC response
func DecodeGrpcRespStatsAlertPolicySpec(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPStatsAlertPolicyStatus(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPStatsAlertPolicyStatus(_ context.Context, r *http.Request) (interface{}, error) {
	var req StatsAlertPolicyStatus
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqStatsAlertPolicyStatus encodes GRPC request
func EncodeGrpcReqStatsAlertPolicyStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*StatsAlertPolicyStatus)
	return req, nil
}

// DecodeGrpcReqStatsAlertPolicyStatus decodes GRPC request
func DecodeGrpcReqStatsAlertPolicyStatus(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*StatsAlertPolicyStatus)
	return req, nil
}

// EncodeGrpcRespStatsAlertPolicyStatus encodes GRC response
func EncodeGrpcRespStatsAlertPolicyStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespStatsAlertPolicyStatus decodes GRPC response
func DecodeGrpcRespStatsAlertPolicyStatus(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPSyslogExport(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPSyslogExport(_ context.Context, r *http.Request) (interface{}, error) {
	var req SyslogExport
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqSyslogExport encodes GRPC request
func EncodeGrpcReqSyslogExport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*SyslogExport)
	return req, nil
}

// DecodeGrpcReqSyslogExport decodes GRPC request
func DecodeGrpcReqSyslogExport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*SyslogExport)
	return req, nil
}

// EncodeGrpcRespSyslogExport encodes GRC response
func EncodeGrpcRespSyslogExport(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespSyslogExport decodes GRPC response
func DecodeGrpcRespSyslogExport(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPThreshold(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPThreshold(_ context.Context, r *http.Request) (interface{}, error) {
	var req Threshold
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqThreshold encodes GRPC request
func EncodeGrpcReqThreshold(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*Threshold)
	return req, nil
}

// DecodeGrpcReqThreshold decodes GRPC request
func DecodeGrpcReqThreshold(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*Threshold)
	return req, nil
}

// EncodeGrpcRespThreshold encodes GRC response
func EncodeGrpcRespThreshold(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespThreshold decodes GRPC response
func DecodeGrpcRespThreshold(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPThresholds(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPThresholds(_ context.Context, r *http.Request) (interface{}, error) {
	var req Thresholds
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqThresholds encodes GRPC request
func EncodeGrpcReqThresholds(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*Thresholds)
	return req, nil
}

// DecodeGrpcReqThresholds decodes GRPC request
func DecodeGrpcReqThresholds(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*Thresholds)
	return req, nil
}

// EncodeGrpcRespThresholds encodes GRC response
func EncodeGrpcRespThresholds(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespThresholds decodes GRPC response
func DecodeGrpcRespThresholds(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
