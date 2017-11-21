// Code generated by MockGen. DO NOT EDIT.
// Source: l2segment.pb.go

package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockL2SegmentClient is a mock of L2SegmentClient interface
type MockL2SegmentClient struct {
	ctrl     *gomock.Controller
	recorder *MockL2SegmentClientMockRecorder
}

// MockL2SegmentClientMockRecorder is the mock recorder for MockL2SegmentClient
type MockL2SegmentClientMockRecorder struct {
	mock *MockL2SegmentClient
}

// NewMockL2SegmentClient creates a new mock instance
func NewMockL2SegmentClient(ctrl *gomock.Controller) *MockL2SegmentClient {
	mock := &MockL2SegmentClient{ctrl: ctrl}
	mock.recorder = &MockL2SegmentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockL2SegmentClient) EXPECT() *MockL2SegmentClientMockRecorder {
	return _m.recorder
}

// L2SegmentCreate mocks base method
func (_m *MockL2SegmentClient) L2SegmentCreate(ctx context.Context, in *L2SegmentRequestMsg, opts ...grpc.CallOption) (*L2SegmentResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "L2SegmentCreate", _s...)
	ret0, _ := ret[0].(*L2SegmentResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentCreate indicates an expected call of L2SegmentCreate
func (_mr *MockL2SegmentClientMockRecorder) L2SegmentCreate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentCreate", reflect.TypeOf((*MockL2SegmentClient)(nil).L2SegmentCreate), _s...)
}

// L2SegmentUpdate mocks base method
func (_m *MockL2SegmentClient) L2SegmentUpdate(ctx context.Context, in *L2SegmentRequestMsg, opts ...grpc.CallOption) (*L2SegmentResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "L2SegmentUpdate", _s...)
	ret0, _ := ret[0].(*L2SegmentResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentUpdate indicates an expected call of L2SegmentUpdate
func (_mr *MockL2SegmentClientMockRecorder) L2SegmentUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentUpdate", reflect.TypeOf((*MockL2SegmentClient)(nil).L2SegmentUpdate), _s...)
}

// L2SegmentDelete mocks base method
func (_m *MockL2SegmentClient) L2SegmentDelete(ctx context.Context, in *L2SegmentDeleteRequestMsg, opts ...grpc.CallOption) (*L2SegmentDeleteResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "L2SegmentDelete", _s...)
	ret0, _ := ret[0].(*L2SegmentDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentDelete indicates an expected call of L2SegmentDelete
func (_mr *MockL2SegmentClientMockRecorder) L2SegmentDelete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentDelete", reflect.TypeOf((*MockL2SegmentClient)(nil).L2SegmentDelete), _s...)
}

// L2SegmentGet mocks base method
func (_m *MockL2SegmentClient) L2SegmentGet(ctx context.Context, in *L2SegmentGetRequestMsg, opts ...grpc.CallOption) (*L2SegmentGetResponseMsg, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "L2SegmentGet", _s...)
	ret0, _ := ret[0].(*L2SegmentGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentGet indicates an expected call of L2SegmentGet
func (_mr *MockL2SegmentClientMockRecorder) L2SegmentGet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentGet", reflect.TypeOf((*MockL2SegmentClient)(nil).L2SegmentGet), _s...)
}

// MockL2SegmentServer is a mock of L2SegmentServer interface
type MockL2SegmentServer struct {
	ctrl     *gomock.Controller
	recorder *MockL2SegmentServerMockRecorder
}

// MockL2SegmentServerMockRecorder is the mock recorder for MockL2SegmentServer
type MockL2SegmentServerMockRecorder struct {
	mock *MockL2SegmentServer
}

// NewMockL2SegmentServer creates a new mock instance
func NewMockL2SegmentServer(ctrl *gomock.Controller) *MockL2SegmentServer {
	mock := &MockL2SegmentServer{ctrl: ctrl}
	mock.recorder = &MockL2SegmentServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockL2SegmentServer) EXPECT() *MockL2SegmentServerMockRecorder {
	return _m.recorder
}

// L2SegmentCreate mocks base method
func (_m *MockL2SegmentServer) L2SegmentCreate(_param0 context.Context, _param1 *L2SegmentRequestMsg) (*L2SegmentResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "L2SegmentCreate", _param0, _param1)
	ret0, _ := ret[0].(*L2SegmentResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentCreate indicates an expected call of L2SegmentCreate
func (_mr *MockL2SegmentServerMockRecorder) L2SegmentCreate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentCreate", reflect.TypeOf((*MockL2SegmentServer)(nil).L2SegmentCreate), arg0, arg1)
}

// L2SegmentUpdate mocks base method
func (_m *MockL2SegmentServer) L2SegmentUpdate(_param0 context.Context, _param1 *L2SegmentRequestMsg) (*L2SegmentResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "L2SegmentUpdate", _param0, _param1)
	ret0, _ := ret[0].(*L2SegmentResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentUpdate indicates an expected call of L2SegmentUpdate
func (_mr *MockL2SegmentServerMockRecorder) L2SegmentUpdate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentUpdate", reflect.TypeOf((*MockL2SegmentServer)(nil).L2SegmentUpdate), arg0, arg1)
}

// L2SegmentDelete mocks base method
func (_m *MockL2SegmentServer) L2SegmentDelete(_param0 context.Context, _param1 *L2SegmentDeleteRequestMsg) (*L2SegmentDeleteResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "L2SegmentDelete", _param0, _param1)
	ret0, _ := ret[0].(*L2SegmentDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentDelete indicates an expected call of L2SegmentDelete
func (_mr *MockL2SegmentServerMockRecorder) L2SegmentDelete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentDelete", reflect.TypeOf((*MockL2SegmentServer)(nil).L2SegmentDelete), arg0, arg1)
}

// L2SegmentGet mocks base method
func (_m *MockL2SegmentServer) L2SegmentGet(_param0 context.Context, _param1 *L2SegmentGetRequestMsg) (*L2SegmentGetResponseMsg, error) {
	ret := _m.ctrl.Call(_m, "L2SegmentGet", _param0, _param1)
	ret0, _ := ret[0].(*L2SegmentGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// L2SegmentGet indicates an expected call of L2SegmentGet
func (_mr *MockL2SegmentServerMockRecorder) L2SegmentGet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "L2SegmentGet", reflect.TypeOf((*MockL2SegmentServer)(nil).L2SegmentGet), arg0, arg1)
}
