// Code generated by MockGen. DO NOT EDIT.
// Source: tcp_proxy.pb.go

// Package halproto is a generated GoMock package.
package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockisTlsProxyAction_Keys is a mock of isTlsProxyAction_Keys interface
type MockisTlsProxyAction_Keys struct {
	ctrl     *gomock.Controller
	recorder *MockisTlsProxyAction_KeysMockRecorder
}

// MockisTlsProxyAction_KeysMockRecorder is the mock recorder for MockisTlsProxyAction_Keys
type MockisTlsProxyAction_KeysMockRecorder struct {
	mock *MockisTlsProxyAction_Keys
}

// NewMockisTlsProxyAction_Keys creates a new mock instance
func NewMockisTlsProxyAction_Keys(ctrl *gomock.Controller) *MockisTlsProxyAction_Keys {
	mock := &MockisTlsProxyAction_Keys{ctrl: ctrl}
	mock.recorder = &MockisTlsProxyAction_KeysMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisTlsProxyAction_Keys) EXPECT() *MockisTlsProxyAction_KeysMockRecorder {
	return m.recorder
}

// isTlsProxyAction_Keys mocks base method
func (m *MockisTlsProxyAction_Keys) isTlsProxyAction_Keys() {
	m.ctrl.Call(m, "isTlsProxyAction_Keys")
}

// isTlsProxyAction_Keys indicates an expected call of isTlsProxyAction_Keys
func (mr *MockisTlsProxyAction_KeysMockRecorder) isTlsProxyAction_Keys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTlsProxyAction_Keys", reflect.TypeOf((*MockisTlsProxyAction_Keys)(nil).isTlsProxyAction_Keys))
}

// MarshalTo mocks base method
func (m *MockisTlsProxyAction_Keys) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisTlsProxyAction_KeysMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTlsProxyAction_Keys)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisTlsProxyAction_Keys) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisTlsProxyAction_KeysMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTlsProxyAction_Keys)(nil).Size))
}

// MockisTcpProxyAction_ProxyConfig is a mock of isTcpProxyAction_ProxyConfig interface
type MockisTcpProxyAction_ProxyConfig struct {
	ctrl     *gomock.Controller
	recorder *MockisTcpProxyAction_ProxyConfigMockRecorder
}

// MockisTcpProxyAction_ProxyConfigMockRecorder is the mock recorder for MockisTcpProxyAction_ProxyConfig
type MockisTcpProxyAction_ProxyConfigMockRecorder struct {
	mock *MockisTcpProxyAction_ProxyConfig
}

// NewMockisTcpProxyAction_ProxyConfig creates a new mock instance
func NewMockisTcpProxyAction_ProxyConfig(ctrl *gomock.Controller) *MockisTcpProxyAction_ProxyConfig {
	mock := &MockisTcpProxyAction_ProxyConfig{ctrl: ctrl}
	mock.recorder = &MockisTcpProxyAction_ProxyConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisTcpProxyAction_ProxyConfig) EXPECT() *MockisTcpProxyAction_ProxyConfigMockRecorder {
	return m.recorder
}

// isTcpProxyAction_ProxyConfig mocks base method
func (m *MockisTcpProxyAction_ProxyConfig) isTcpProxyAction_ProxyConfig() {
	m.ctrl.Call(m, "isTcpProxyAction_ProxyConfig")
}

// isTcpProxyAction_ProxyConfig indicates an expected call of isTcpProxyAction_ProxyConfig
func (mr *MockisTcpProxyAction_ProxyConfigMockRecorder) isTcpProxyAction_ProxyConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTcpProxyAction_ProxyConfig", reflect.TypeOf((*MockisTcpProxyAction_ProxyConfig)(nil).isTcpProxyAction_ProxyConfig))
}

// MarshalTo mocks base method
func (m *MockisTcpProxyAction_ProxyConfig) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisTcpProxyAction_ProxyConfigMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTcpProxyAction_ProxyConfig)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisTcpProxyAction_ProxyConfig) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisTcpProxyAction_ProxyConfigMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTcpProxyAction_ProxyConfig)(nil).Size))
}

// MockisTcpProxyCbKeyHandle_KeyOrHandle is a mock of isTcpProxyCbKeyHandle_KeyOrHandle interface
type MockisTcpProxyCbKeyHandle_KeyOrHandle struct {
	ctrl     *gomock.Controller
	recorder *MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder
}

// MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder is the mock recorder for MockisTcpProxyCbKeyHandle_KeyOrHandle
type MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder struct {
	mock *MockisTcpProxyCbKeyHandle_KeyOrHandle
}

// NewMockisTcpProxyCbKeyHandle_KeyOrHandle creates a new mock instance
func NewMockisTcpProxyCbKeyHandle_KeyOrHandle(ctrl *gomock.Controller) *MockisTcpProxyCbKeyHandle_KeyOrHandle {
	mock := &MockisTcpProxyCbKeyHandle_KeyOrHandle{ctrl: ctrl}
	mock.recorder = &MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisTcpProxyCbKeyHandle_KeyOrHandle) EXPECT() *MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder {
	return m.recorder
}

// isTcpProxyCbKeyHandle_KeyOrHandle mocks base method
func (m *MockisTcpProxyCbKeyHandle_KeyOrHandle) isTcpProxyCbKeyHandle_KeyOrHandle() {
	m.ctrl.Call(m, "isTcpProxyCbKeyHandle_KeyOrHandle")
}

// isTcpProxyCbKeyHandle_KeyOrHandle indicates an expected call of isTcpProxyCbKeyHandle_KeyOrHandle
func (mr *MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder) isTcpProxyCbKeyHandle_KeyOrHandle() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTcpProxyCbKeyHandle_KeyOrHandle", reflect.TypeOf((*MockisTcpProxyCbKeyHandle_KeyOrHandle)(nil).isTcpProxyCbKeyHandle_KeyOrHandle))
}

// MarshalTo mocks base method
func (m *MockisTcpProxyCbKeyHandle_KeyOrHandle) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTcpProxyCbKeyHandle_KeyOrHandle)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisTcpProxyCbKeyHandle_KeyOrHandle) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisTcpProxyCbKeyHandle_KeyOrHandleMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTcpProxyCbKeyHandle_KeyOrHandle)(nil).Size))
}

// MockTcpProxyClient is a mock of TcpProxyClient interface
type MockTcpProxyClient struct {
	ctrl     *gomock.Controller
	recorder *MockTcpProxyClientMockRecorder
}

// MockTcpProxyClientMockRecorder is the mock recorder for MockTcpProxyClient
type MockTcpProxyClientMockRecorder struct {
	mock *MockTcpProxyClient
}

// NewMockTcpProxyClient creates a new mock instance
func NewMockTcpProxyClient(ctrl *gomock.Controller) *MockTcpProxyClient {
	mock := &MockTcpProxyClient{ctrl: ctrl}
	mock.recorder = &MockTcpProxyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTcpProxyClient) EXPECT() *MockTcpProxyClientMockRecorder {
	return m.recorder
}

// TcpProxyRuleCreate mocks base method
func (m *MockTcpProxyClient) TcpProxyRuleCreate(ctx context.Context, in *TcpProxyRuleRequestMsg, opts ...grpc.CallOption) (*TcpProxyRuleResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyRuleCreate", varargs...)
	ret0, _ := ret[0].(*TcpProxyRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleCreate indicates an expected call of TcpProxyRuleCreate
func (mr *MockTcpProxyClientMockRecorder) TcpProxyRuleCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleCreate", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyRuleCreate), varargs...)
}

// TcpProxyRuleUpdate mocks base method
func (m *MockTcpProxyClient) TcpProxyRuleUpdate(ctx context.Context, in *TcpProxyRuleRequestMsg, opts ...grpc.CallOption) (*TcpProxyRuleResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyRuleUpdate", varargs...)
	ret0, _ := ret[0].(*TcpProxyRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleUpdate indicates an expected call of TcpProxyRuleUpdate
func (mr *MockTcpProxyClientMockRecorder) TcpProxyRuleUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleUpdate", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyRuleUpdate), varargs...)
}

// TcpProxyRuleDelete mocks base method
func (m *MockTcpProxyClient) TcpProxyRuleDelete(ctx context.Context, in *TcpProxyRuleDeleteRequestMsg, opts ...grpc.CallOption) (*TcpProxyRuleDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyRuleDelete", varargs...)
	ret0, _ := ret[0].(*TcpProxyRuleDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleDelete indicates an expected call of TcpProxyRuleDelete
func (mr *MockTcpProxyClientMockRecorder) TcpProxyRuleDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleDelete", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyRuleDelete), varargs...)
}

// TcpProxyRuleGet mocks base method
func (m *MockTcpProxyClient) TcpProxyRuleGet(ctx context.Context, in *TcpProxyRuleGetRequestMsg, opts ...grpc.CallOption) (*TcpProxyRuleGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyRuleGet", varargs...)
	ret0, _ := ret[0].(*TcpProxyRuleGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleGet indicates an expected call of TcpProxyRuleGet
func (mr *MockTcpProxyClientMockRecorder) TcpProxyRuleGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleGet", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyRuleGet), varargs...)
}

// TcpProxyCbCreate mocks base method
func (m *MockTcpProxyClient) TcpProxyCbCreate(ctx context.Context, in *TcpProxyCbRequestMsg, opts ...grpc.CallOption) (*TcpProxyCbResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyCbCreate", varargs...)
	ret0, _ := ret[0].(*TcpProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbCreate indicates an expected call of TcpProxyCbCreate
func (mr *MockTcpProxyClientMockRecorder) TcpProxyCbCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbCreate", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyCbCreate), varargs...)
}

// TcpProxyCbUpdate mocks base method
func (m *MockTcpProxyClient) TcpProxyCbUpdate(ctx context.Context, in *TcpProxyCbRequestMsg, opts ...grpc.CallOption) (*TcpProxyCbResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyCbUpdate", varargs...)
	ret0, _ := ret[0].(*TcpProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbUpdate indicates an expected call of TcpProxyCbUpdate
func (mr *MockTcpProxyClientMockRecorder) TcpProxyCbUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbUpdate", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyCbUpdate), varargs...)
}

// TcpProxyCbDelete mocks base method
func (m *MockTcpProxyClient) TcpProxyCbDelete(ctx context.Context, in *TcpProxyCbDeleteRequestMsg, opts ...grpc.CallOption) (*TcpProxyCbDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyCbDelete", varargs...)
	ret0, _ := ret[0].(*TcpProxyCbDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbDelete indicates an expected call of TcpProxyCbDelete
func (mr *MockTcpProxyClientMockRecorder) TcpProxyCbDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbDelete", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyCbDelete), varargs...)
}

// TcpProxyCbGet mocks base method
func (m *MockTcpProxyClient) TcpProxyCbGet(ctx context.Context, in *TcpProxyCbGetRequestMsg, opts ...grpc.CallOption) (*TcpProxyCbGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyCbGet", varargs...)
	ret0, _ := ret[0].(*TcpProxyCbGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbGet indicates an expected call of TcpProxyCbGet
func (mr *MockTcpProxyClientMockRecorder) TcpProxyCbGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbGet", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyCbGet), varargs...)
}

// TcpProxySessionGet mocks base method
func (m *MockTcpProxyClient) TcpProxySessionGet(ctx context.Context, in *TcpProxySessionGetRequestMsg, opts ...grpc.CallOption) (*TcpProxySessionGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxySessionGet", varargs...)
	ret0, _ := ret[0].(*TcpProxySessionGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxySessionGet indicates an expected call of TcpProxySessionGet
func (mr *MockTcpProxyClientMockRecorder) TcpProxySessionGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxySessionGet", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxySessionGet), varargs...)
}

// TcpProxyGlobalStatsGet mocks base method
func (m *MockTcpProxyClient) TcpProxyGlobalStatsGet(ctx context.Context, in *TcpProxyGlobalStatsGetRequestMsg, opts ...grpc.CallOption) (*TcpProxyGlobalStatsGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TcpProxyGlobalStatsGet", varargs...)
	ret0, _ := ret[0].(*TcpProxyGlobalStatsGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyGlobalStatsGet indicates an expected call of TcpProxyGlobalStatsGet
func (mr *MockTcpProxyClientMockRecorder) TcpProxyGlobalStatsGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyGlobalStatsGet", reflect.TypeOf((*MockTcpProxyClient)(nil).TcpProxyGlobalStatsGet), varargs...)
}

// MockTcpProxyServer is a mock of TcpProxyServer interface
type MockTcpProxyServer struct {
	ctrl     *gomock.Controller
	recorder *MockTcpProxyServerMockRecorder
}

// MockTcpProxyServerMockRecorder is the mock recorder for MockTcpProxyServer
type MockTcpProxyServerMockRecorder struct {
	mock *MockTcpProxyServer
}

// NewMockTcpProxyServer creates a new mock instance
func NewMockTcpProxyServer(ctrl *gomock.Controller) *MockTcpProxyServer {
	mock := &MockTcpProxyServer{ctrl: ctrl}
	mock.recorder = &MockTcpProxyServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTcpProxyServer) EXPECT() *MockTcpProxyServerMockRecorder {
	return m.recorder
}

// TcpProxyRuleCreate mocks base method
func (m *MockTcpProxyServer) TcpProxyRuleCreate(arg0 context.Context, arg1 *TcpProxyRuleRequestMsg) (*TcpProxyRuleResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyRuleCreate", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleCreate indicates an expected call of TcpProxyRuleCreate
func (mr *MockTcpProxyServerMockRecorder) TcpProxyRuleCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleCreate", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyRuleCreate), arg0, arg1)
}

// TcpProxyRuleUpdate mocks base method
func (m *MockTcpProxyServer) TcpProxyRuleUpdate(arg0 context.Context, arg1 *TcpProxyRuleRequestMsg) (*TcpProxyRuleResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyRuleUpdate", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleUpdate indicates an expected call of TcpProxyRuleUpdate
func (mr *MockTcpProxyServerMockRecorder) TcpProxyRuleUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleUpdate", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyRuleUpdate), arg0, arg1)
}

// TcpProxyRuleDelete mocks base method
func (m *MockTcpProxyServer) TcpProxyRuleDelete(arg0 context.Context, arg1 *TcpProxyRuleDeleteRequestMsg) (*TcpProxyRuleDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyRuleDelete", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyRuleDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleDelete indicates an expected call of TcpProxyRuleDelete
func (mr *MockTcpProxyServerMockRecorder) TcpProxyRuleDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleDelete", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyRuleDelete), arg0, arg1)
}

// TcpProxyRuleGet mocks base method
func (m *MockTcpProxyServer) TcpProxyRuleGet(arg0 context.Context, arg1 *TcpProxyRuleGetRequestMsg) (*TcpProxyRuleGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyRuleGet", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyRuleGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyRuleGet indicates an expected call of TcpProxyRuleGet
func (mr *MockTcpProxyServerMockRecorder) TcpProxyRuleGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyRuleGet", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyRuleGet), arg0, arg1)
}

// TcpProxyCbCreate mocks base method
func (m *MockTcpProxyServer) TcpProxyCbCreate(arg0 context.Context, arg1 *TcpProxyCbRequestMsg) (*TcpProxyCbResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyCbCreate", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbCreate indicates an expected call of TcpProxyCbCreate
func (mr *MockTcpProxyServerMockRecorder) TcpProxyCbCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbCreate", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyCbCreate), arg0, arg1)
}

// TcpProxyCbUpdate mocks base method
func (m *MockTcpProxyServer) TcpProxyCbUpdate(arg0 context.Context, arg1 *TcpProxyCbRequestMsg) (*TcpProxyCbResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyCbUpdate", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbUpdate indicates an expected call of TcpProxyCbUpdate
func (mr *MockTcpProxyServerMockRecorder) TcpProxyCbUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbUpdate", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyCbUpdate), arg0, arg1)
}

// TcpProxyCbDelete mocks base method
func (m *MockTcpProxyServer) TcpProxyCbDelete(arg0 context.Context, arg1 *TcpProxyCbDeleteRequestMsg) (*TcpProxyCbDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyCbDelete", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyCbDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbDelete indicates an expected call of TcpProxyCbDelete
func (mr *MockTcpProxyServerMockRecorder) TcpProxyCbDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbDelete", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyCbDelete), arg0, arg1)
}

// TcpProxyCbGet mocks base method
func (m *MockTcpProxyServer) TcpProxyCbGet(arg0 context.Context, arg1 *TcpProxyCbGetRequestMsg) (*TcpProxyCbGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyCbGet", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyCbGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyCbGet indicates an expected call of TcpProxyCbGet
func (mr *MockTcpProxyServerMockRecorder) TcpProxyCbGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyCbGet", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyCbGet), arg0, arg1)
}

// TcpProxySessionGet mocks base method
func (m *MockTcpProxyServer) TcpProxySessionGet(arg0 context.Context, arg1 *TcpProxySessionGetRequestMsg) (*TcpProxySessionGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxySessionGet", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxySessionGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxySessionGet indicates an expected call of TcpProxySessionGet
func (mr *MockTcpProxyServerMockRecorder) TcpProxySessionGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxySessionGet", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxySessionGet), arg0, arg1)
}

// TcpProxyGlobalStatsGet mocks base method
func (m *MockTcpProxyServer) TcpProxyGlobalStatsGet(arg0 context.Context, arg1 *TcpProxyGlobalStatsGetRequestMsg) (*TcpProxyGlobalStatsGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "TcpProxyGlobalStatsGet", arg0, arg1)
	ret0, _ := ret[0].(*TcpProxyGlobalStatsGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TcpProxyGlobalStatsGet indicates an expected call of TcpProxyGlobalStatsGet
func (mr *MockTcpProxyServerMockRecorder) TcpProxyGlobalStatsGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TcpProxyGlobalStatsGet", reflect.TypeOf((*MockTcpProxyServer)(nil).TcpProxyGlobalStatsGet), arg0, arg1)
}
