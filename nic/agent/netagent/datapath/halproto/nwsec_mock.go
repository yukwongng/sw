// Code generated by MockGen. DO NOT EDIT.
// Source: nwsec.pb.go

package halproto

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockisSecurityProfileKeyHandle_KeyOrHandle is a mock of isSecurityProfileKeyHandle_KeyOrHandle interface
type MockisSecurityProfileKeyHandle_KeyOrHandle struct {
	ctrl     *gomock.Controller
	recorder *MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder
}

// MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder is the mock recorder for MockisSecurityProfileKeyHandle_KeyOrHandle
type MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder struct {
	mock *MockisSecurityProfileKeyHandle_KeyOrHandle
}

// NewMockisSecurityProfileKeyHandle_KeyOrHandle creates a new mock instance
func NewMockisSecurityProfileKeyHandle_KeyOrHandle(ctrl *gomock.Controller) *MockisSecurityProfileKeyHandle_KeyOrHandle {
	mock := &MockisSecurityProfileKeyHandle_KeyOrHandle{ctrl: ctrl}
	mock.recorder = &MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockisSecurityProfileKeyHandle_KeyOrHandle) EXPECT() *MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder {
	return _m.recorder
}

// isSecurityProfileKeyHandle_KeyOrHandle mocks base method
func (_m *MockisSecurityProfileKeyHandle_KeyOrHandle) isSecurityProfileKeyHandle_KeyOrHandle() {
	_m.ctrl.Call(_m, "isSecurityProfileKeyHandle_KeyOrHandle")
}

// isSecurityProfileKeyHandle_KeyOrHandle indicates an expected call of isSecurityProfileKeyHandle_KeyOrHandle
func (_mr *MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder) isSecurityProfileKeyHandle_KeyOrHandle() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "isSecurityProfileKeyHandle_KeyOrHandle", reflect.TypeOf((*MockisSecurityProfileKeyHandle_KeyOrHandle)(nil).isSecurityProfileKeyHandle_KeyOrHandle))
}

// MarshalTo mocks base method
func (_m *MockisSecurityProfileKeyHandle_KeyOrHandle) MarshalTo(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "MarshalTo", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (_mr *MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MarshalTo", reflect.TypeOf((*MockisSecurityProfileKeyHandle_KeyOrHandle)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (_m *MockisSecurityProfileKeyHandle_KeyOrHandle) Size() int {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (_mr *MockisSecurityProfileKeyHandle_KeyOrHandleMockRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Size", reflect.TypeOf((*MockisSecurityProfileKeyHandle_KeyOrHandle)(nil).Size))
}

// MockisService_L4Info is a mock of isService_L4Info interface
type MockisService_L4Info struct {
	ctrl     *gomock.Controller
	recorder *MockisService_L4InfoMockRecorder
}

// MockisService_L4InfoMockRecorder is the mock recorder for MockisService_L4Info
type MockisService_L4InfoMockRecorder struct {
	mock *MockisService_L4Info
}

// NewMockisService_L4Info creates a new mock instance
func NewMockisService_L4Info(ctrl *gomock.Controller) *MockisService_L4Info {
	mock := &MockisService_L4Info{ctrl: ctrl}
	mock.recorder = &MockisService_L4InfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockisService_L4Info) EXPECT() *MockisService_L4InfoMockRecorder {
	return _m.recorder
}

// isService_L4Info mocks base method
func (_m *MockisService_L4Info) isService_L4Info() {
	_m.ctrl.Call(_m, "isService_L4Info")
}

// isService_L4Info indicates an expected call of isService_L4Info
func (_mr *MockisService_L4InfoMockRecorder) isService_L4Info() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "isService_L4Info", reflect.TypeOf((*MockisService_L4Info)(nil).isService_L4Info))
}

// MarshalTo mocks base method
func (_m *MockisService_L4Info) MarshalTo(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "MarshalTo", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (_mr *MockisService_L4InfoMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MarshalTo", reflect.TypeOf((*MockisService_L4Info)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (_m *MockisService_L4Info) Size() int {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (_mr *MockisService_L4InfoMockRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Size", reflect.TypeOf((*MockisService_L4Info)(nil).Size))
}

// MockisDoSService_L4Info is a mock of isDoSService_L4Info interface
type MockisDoSService_L4Info struct {
	ctrl     *gomock.Controller
	recorder *MockisDoSService_L4InfoMockRecorder
}

// MockisDoSService_L4InfoMockRecorder is the mock recorder for MockisDoSService_L4Info
type MockisDoSService_L4InfoMockRecorder struct {
	mock *MockisDoSService_L4Info
}

// NewMockisDoSService_L4Info creates a new mock instance
func NewMockisDoSService_L4Info(ctrl *gomock.Controller) *MockisDoSService_L4Info {
	mock := &MockisDoSService_L4Info{ctrl: ctrl}
	mock.recorder = &MockisDoSService_L4InfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockisDoSService_L4Info) EXPECT() *MockisDoSService_L4InfoMockRecorder {
	return _m.recorder
}

// isDoSService_L4Info mocks base method
func (_m *MockisDoSService_L4Info) isDoSService_L4Info() {
	_m.ctrl.Call(_m, "isDoSService_L4Info")
}

// isDoSService_L4Info indicates an expected call of isDoSService_L4Info
func (_mr *MockisDoSService_L4InfoMockRecorder) isDoSService_L4Info() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "isDoSService_L4Info", reflect.TypeOf((*MockisDoSService_L4Info)(nil).isDoSService_L4Info))
}

// MarshalTo mocks base method
func (_m *MockisDoSService_L4Info) MarshalTo(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "MarshalTo", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (_mr *MockisDoSService_L4InfoMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MarshalTo", reflect.TypeOf((*MockisDoSService_L4Info)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (_m *MockisDoSService_L4Info) Size() int {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (_mr *MockisDoSService_L4InfoMockRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Size", reflect.TypeOf((*MockisDoSService_L4Info)(nil).Size))
}

// MockisSecurityGroupKeyHandle_KeyOrHandle is a mock of isSecurityGroupKeyHandle_KeyOrHandle interface
type MockisSecurityGroupKeyHandle_KeyOrHandle struct {
	ctrl     *gomock.Controller
	recorder *MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder
}

// MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder is the mock recorder for MockisSecurityGroupKeyHandle_KeyOrHandle
type MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder struct {
	mock *MockisSecurityGroupKeyHandle_KeyOrHandle
}

// NewMockisSecurityGroupKeyHandle_KeyOrHandle creates a new mock instance
func NewMockisSecurityGroupKeyHandle_KeyOrHandle(ctrl *gomock.Controller) *MockisSecurityGroupKeyHandle_KeyOrHandle {
	mock := &MockisSecurityGroupKeyHandle_KeyOrHandle{ctrl: ctrl}
	mock.recorder = &MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockisSecurityGroupKeyHandle_KeyOrHandle) EXPECT() *MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder {
	return _m.recorder
}

// isSecurityGroupKeyHandle_KeyOrHandle mocks base method
func (_m *MockisSecurityGroupKeyHandle_KeyOrHandle) isSecurityGroupKeyHandle_KeyOrHandle() {
	_m.ctrl.Call(_m, "isSecurityGroupKeyHandle_KeyOrHandle")
}

// isSecurityGroupKeyHandle_KeyOrHandle indicates an expected call of isSecurityGroupKeyHandle_KeyOrHandle
func (_mr *MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder) isSecurityGroupKeyHandle_KeyOrHandle() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "isSecurityGroupKeyHandle_KeyOrHandle", reflect.TypeOf((*MockisSecurityGroupKeyHandle_KeyOrHandle)(nil).isSecurityGroupKeyHandle_KeyOrHandle))
}

// MarshalTo mocks base method
func (_m *MockisSecurityGroupKeyHandle_KeyOrHandle) MarshalTo(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "MarshalTo", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (_mr *MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MarshalTo", reflect.TypeOf((*MockisSecurityGroupKeyHandle_KeyOrHandle)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (_m *MockisSecurityGroupKeyHandle_KeyOrHandle) Size() int {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (_mr *MockisSecurityGroupKeyHandle_KeyOrHandleMockRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Size", reflect.TypeOf((*MockisSecurityGroupKeyHandle_KeyOrHandle)(nil).Size))
}

// MockNwSecurityClient is a mock of NwSecurityClient interface
type MockNwSecurityClient struct {
	ctrl     *gomock.Controller
	recorder *MockNwSecurityClientMockRecorder
}

// MockNwSecurityClientMockRecorder is the mock recorder for MockNwSecurityClient
type MockNwSecurityClientMockRecorder struct {
	mock *MockNwSecurityClient
}

// NewMockNwSecurityClient creates a new mock instance
func NewMockNwSecurityClient(ctrl *gomock.Controller) *MockNwSecurityClient {
	mock := &MockNwSecurityClient{ctrl: ctrl}
	mock.recorder = &MockNwSecurityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNwSecurityClient) EXPECT() *MockNwSecurityClientMockRecorder {
	return _m.recorder
}

// SecurityProfileCreate mocks base method
func (_m *MockNwSecurityClient) SecurityProfileCreate(ctx context.Context, in *SecurityProfileRequestMsg, opts ...grpc.CallOption) (*SecurityProfileResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityProfileCreate", _s...)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileCreate indicates an expected call of SecurityProfileCreate
func (_mr *MockNwSecurityClientMockRecorder) SecurityProfileCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileCreate), _s...)
}

// SecurityProfileUpdate mocks base method
func (_m *MockNwSecurityClient) SecurityProfileUpdate(ctx context.Context, in *SecurityProfileRequestMsg, opts ...grpc.CallOption) (*SecurityProfileResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityProfileUpdate", _s...)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileUpdate indicates an expected call of SecurityProfileUpdate
func (_mr *MockNwSecurityClientMockRecorder) SecurityProfileUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileUpdate), _s...)
}

// SecurityProfileDelete mocks base method
func (_m *MockNwSecurityClient) SecurityProfileDelete(ctx context.Context, in *SecurityProfileDeleteRequestMsg, opts ...grpc.CallOption) (*SecurityProfileDeleteResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityProfileDelete", _s...)
	ret0, _ := ret[0].(*SecurityProfileDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileDelete indicates an expected call of SecurityProfileDelete
func (_mr *MockNwSecurityClientMockRecorder) SecurityProfileDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileDelete), _s...)
}

// SecurityProfileGet mocks base method
func (_m *MockNwSecurityClient) SecurityProfileGet(ctx context.Context, in *SecurityProfileGetRequestMsg, opts ...grpc.CallOption) (*SecurityProfileGetResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityProfileGet", _s...)
	ret0, _ := ret[0].(*SecurityProfileGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileGet indicates an expected call of SecurityProfileGet
func (_mr *MockNwSecurityClientMockRecorder) SecurityProfileGet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileGet", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileGet), _s...)
}

// SecurityGroupCreate mocks base method
func (_m *MockNwSecurityClient) SecurityGroupCreate(ctx context.Context, in *SecurityGroupRequestMsg, opts ...grpc.CallOption) (*SecurityGroupResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityGroupCreate", _s...)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupCreate indicates an expected call of SecurityGroupCreate
func (_mr *MockNwSecurityClientMockRecorder) SecurityGroupCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupCreate), _s...)
}

// SecurityGroupUpdate mocks base method
func (_m *MockNwSecurityClient) SecurityGroupUpdate(ctx context.Context, in *SecurityGroupRequestMsg, opts ...grpc.CallOption) (*SecurityGroupResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityGroupUpdate", _s...)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupUpdate indicates an expected call of SecurityGroupUpdate
func (_mr *MockNwSecurityClientMockRecorder) SecurityGroupUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupUpdate), _s...)
}

// SecurityGroupDelete mocks base method
func (_m *MockNwSecurityClient) SecurityGroupDelete(ctx context.Context, in *SecurityGroupDeleteRequestMsg, opts ...grpc.CallOption) (*SecurityGroupDeleteResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityGroupDelete", _s...)
	ret0, _ := ret[0].(*SecurityGroupDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupDelete indicates an expected call of SecurityGroupDelete
func (_mr *MockNwSecurityClientMockRecorder) SecurityGroupDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupDelete), _s...)
}

// SecurityGroupGet mocks base method
func (_m *MockNwSecurityClient) SecurityGroupGet(ctx context.Context, in *SecurityGroupGetRequestMsg, opts ...grpc.CallOption) (*SecurityGroupGetResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SecurityGroupGet", _s...)
	ret0, _ := ret[0].(*SecurityGroupGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupGet indicates an expected call of SecurityGroupGet
func (_mr *MockNwSecurityClientMockRecorder) SecurityGroupGet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupGet", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupGet), _s...)
}

// DoSPolicyCreate mocks base method
func (_m *MockNwSecurityClient) DoSPolicyCreate(ctx context.Context, in *DoSPolicyRequestMsg, opts ...grpc.CallOption) (*DoSPolicyResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DoSPolicyCreate", _s...)
	ret0, _ := ret[0].(*DoSPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyCreate indicates an expected call of DoSPolicyCreate
func (_mr *MockNwSecurityClientMockRecorder) DoSPolicyCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).DoSPolicyCreate), _s...)
}

// DoSPolicyUpdate mocks base method
func (_m *MockNwSecurityClient) DoSPolicyUpdate(ctx context.Context, in *DoSPolicyRequestMsg, opts ...grpc.CallOption) (*DoSPolicyResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DoSPolicyUpdate", _s...)
	ret0, _ := ret[0].(*DoSPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyUpdate indicates an expected call of DoSPolicyUpdate
func (_mr *MockNwSecurityClientMockRecorder) DoSPolicyUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).DoSPolicyUpdate), _s...)
}

// DoSPolicyDelete mocks base method
func (_m *MockNwSecurityClient) DoSPolicyDelete(ctx context.Context, in *DoSPolicyDeleteRequestMsg, opts ...grpc.CallOption) (*DoSPolicyDeleteResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DoSPolicyDelete", _s...)
	ret0, _ := ret[0].(*DoSPolicyDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyDelete indicates an expected call of DoSPolicyDelete
func (_mr *MockNwSecurityClientMockRecorder) DoSPolicyDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).DoSPolicyDelete), _s...)
}

// DoSPolicyGet mocks base method
func (_m *MockNwSecurityClient) DoSPolicyGet(ctx context.Context, in *DoSPolicyGetRequestMsg, opts ...grpc.CallOption) (*DoSPolicyGetResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DoSPolicyGet", _s...)
	ret0, _ := ret[0].(*DoSPolicyGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyGet indicates an expected call of DoSPolicyGet
func (_mr *MockNwSecurityClientMockRecorder) DoSPolicyGet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyGet", reflect.TypeOf((*MockNwSecurityClient)(nil).DoSPolicyGet), _s...)
}

// MockNwSecurityServer is a mock of NwSecurityServer interface
type MockNwSecurityServer struct {
	ctrl     *gomock.Controller
	recorder *MockNwSecurityServerMockRecorder
}

// MockNwSecurityServerMockRecorder is the mock recorder for MockNwSecurityServer
type MockNwSecurityServerMockRecorder struct {
	mock *MockNwSecurityServer
}

// NewMockNwSecurityServer creates a new mock instance
func NewMockNwSecurityServer(ctrl *gomock.Controller) *MockNwSecurityServer {
	mock := &MockNwSecurityServer{ctrl: ctrl}
	mock.recorder = &MockNwSecurityServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockNwSecurityServer) EXPECT() *MockNwSecurityServerMockRecorder {
	return _m.recorder
}

// SecurityProfileCreate mocks base method
func (_m *MockNwSecurityServer) SecurityProfileCreate(_param0 context.Context, _param1 *SecurityProfileRequestMsg) (*SecurityProfileResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityProfileCreate", _param0, _param1)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileCreate indicates an expected call of SecurityProfileCreate
func (_mr *MockNwSecurityServerMockRecorder) SecurityProfileCreate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileCreate), arg0, arg1)
}

// SecurityProfileUpdate mocks base method
func (_m *MockNwSecurityServer) SecurityProfileUpdate(_param0 context.Context, _param1 *SecurityProfileRequestMsg) (*SecurityProfileResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityProfileUpdate", _param0, _param1)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileUpdate indicates an expected call of SecurityProfileUpdate
func (_mr *MockNwSecurityServerMockRecorder) SecurityProfileUpdate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileUpdate), arg0, arg1)
}

// SecurityProfileDelete mocks base method
func (_m *MockNwSecurityServer) SecurityProfileDelete(_param0 context.Context, _param1 *SecurityProfileDeleteRequestMsg) (*SecurityProfileDeleteResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityProfileDelete", _param0, _param1)
	ret0, _ := ret[0].(*SecurityProfileDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileDelete indicates an expected call of SecurityProfileDelete
func (_mr *MockNwSecurityServerMockRecorder) SecurityProfileDelete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileDelete), arg0, arg1)
}

// SecurityProfileGet mocks base method
func (_m *MockNwSecurityServer) SecurityProfileGet(_param0 context.Context, _param1 *SecurityProfileGetRequestMsg) (*SecurityProfileGetResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityProfileGet", _param0, _param1)
	ret0, _ := ret[0].(*SecurityProfileGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileGet indicates an expected call of SecurityProfileGet
func (_mr *MockNwSecurityServerMockRecorder) SecurityProfileGet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityProfileGet", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileGet), arg0, arg1)
}

// SecurityGroupCreate mocks base method
func (_m *MockNwSecurityServer) SecurityGroupCreate(_param0 context.Context, _param1 *SecurityGroupRequestMsg) (*SecurityGroupResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityGroupCreate", _param0, _param1)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupCreate indicates an expected call of SecurityGroupCreate
func (_mr *MockNwSecurityServerMockRecorder) SecurityGroupCreate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupCreate), arg0, arg1)
}

// SecurityGroupUpdate mocks base method
func (_m *MockNwSecurityServer) SecurityGroupUpdate(_param0 context.Context, _param1 *SecurityGroupRequestMsg) (*SecurityGroupResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityGroupUpdate", _param0, _param1)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupUpdate indicates an expected call of SecurityGroupUpdate
func (_mr *MockNwSecurityServerMockRecorder) SecurityGroupUpdate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupUpdate), arg0, arg1)
}

// SecurityGroupDelete mocks base method
func (_m *MockNwSecurityServer) SecurityGroupDelete(_param0 context.Context, _param1 *SecurityGroupDeleteRequestMsg) (*SecurityGroupDeleteResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityGroupDelete", _param0, _param1)
	ret0, _ := ret[0].(*SecurityGroupDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupDelete indicates an expected call of SecurityGroupDelete
func (_mr *MockNwSecurityServerMockRecorder) SecurityGroupDelete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupDelete), arg0, arg1)
}

// SecurityGroupGet mocks base method
func (_m *MockNwSecurityServer) SecurityGroupGet(_param0 context.Context, _param1 *SecurityGroupGetRequestMsg) (*SecurityGroupGetResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "SecurityGroupGet", _param0, _param1)
	ret0, _ := ret[0].(*SecurityGroupGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupGet indicates an expected call of SecurityGroupGet
func (_mr *MockNwSecurityServerMockRecorder) SecurityGroupGet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SecurityGroupGet", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupGet), arg0, arg1)
}

// DoSPolicyCreate mocks base method
func (_m *MockNwSecurityServer) DoSPolicyCreate(_param0 context.Context, _param1 *DoSPolicyRequestMsg) (*DoSPolicyResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "DoSPolicyCreate", _param0, _param1)
	ret0, _ := ret[0].(*DoSPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyCreate indicates an expected call of DoSPolicyCreate
func (_mr *MockNwSecurityServerMockRecorder) DoSPolicyCreate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).DoSPolicyCreate), arg0, arg1)
}

// DoSPolicyUpdate mocks base method
func (_m *MockNwSecurityServer) DoSPolicyUpdate(_param0 context.Context, _param1 *DoSPolicyRequestMsg) (*DoSPolicyResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "DoSPolicyUpdate", _param0, _param1)
	ret0, _ := ret[0].(*DoSPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyUpdate indicates an expected call of DoSPolicyUpdate
func (_mr *MockNwSecurityServerMockRecorder) DoSPolicyUpdate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).DoSPolicyUpdate), arg0, arg1)
}

// DoSPolicyDelete mocks base method
func (_m *MockNwSecurityServer) DoSPolicyDelete(_param0 context.Context, _param1 *DoSPolicyDeleteRequestMsg) (*DoSPolicyDeleteResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "DoSPolicyDelete", _param0, _param1)
	ret0, _ := ret[0].(*DoSPolicyDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyDelete indicates an expected call of DoSPolicyDelete
func (_mr *MockNwSecurityServerMockRecorder) DoSPolicyDelete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).DoSPolicyDelete), arg0, arg1)
}

// DoSPolicyGet mocks base method
func (_m *MockNwSecurityServer) DoSPolicyGet(_param0 context.Context, _param1 *DoSPolicyGetRequestMsg) (*DoSPolicyGetResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "DoSPolicyGet", _param0, _param1)
	ret0, _ := ret[0].(*DoSPolicyGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSPolicyGet indicates an expected call of DoSPolicyGet
func (_mr *MockNwSecurityServerMockRecorder) DoSPolicyGet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSPolicyGet", reflect.TypeOf((*MockNwSecurityServer)(nil).DoSPolicyGet), arg0, arg1)
}
