// Code generated by MockGen. DO NOT EDIT.
// Source: ../vcprobe.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	object "github.com/vmware/govmomi/object"
	mo "github.com/vmware/govmomi/vim25/mo"
	types "github.com/vmware/govmomi/vim25/types"

	defs "github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/defs"
	vcprobe "github.com/pensando/sw/venice/ctrler/orchhub/orchestrators/vchub/vcprobe"
)

// MockTagsProbeInf is a mock of TagsProbeInf interface
type MockTagsProbeInf struct {
	ctrl     *gomock.Controller
	recorder *MockTagsProbeInfMockRecorder
}

// MockTagsProbeInfMockRecorder is the mock recorder for MockTagsProbeInf
type MockTagsProbeInfMockRecorder struct {
	mock *MockTagsProbeInf
}

// NewMockTagsProbeInf creates a new mock instance
func NewMockTagsProbeInf(ctrl *gomock.Controller) *MockTagsProbeInf {
	mock := &MockTagsProbeInf{ctrl: ctrl}
	mock.recorder = &MockTagsProbeInfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagsProbeInf) EXPECT() *MockTagsProbeInfMockRecorder {
	return m.recorder
}

// StartWatch mocks base method
func (m *MockTagsProbeInf) StartWatch() {
	m.ctrl.Call(m, "StartWatch")
}

// StartWatch indicates an expected call of StartWatch
func (mr *MockTagsProbeInfMockRecorder) StartWatch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWatch", reflect.TypeOf((*MockTagsProbeInf)(nil).StartWatch))
}

// SetupBaseTags mocks base method
func (m *MockTagsProbeInf) SetupBaseTags() bool {
	ret := m.ctrl.Call(m, "SetupBaseTags")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetupBaseTags indicates an expected call of SetupBaseTags
func (mr *MockTagsProbeInfMockRecorder) SetupBaseTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupBaseTags", reflect.TypeOf((*MockTagsProbeInf)(nil).SetupBaseTags))
}

// GetPensandoTagsOnObjects mocks base method
func (m *MockTagsProbeInf) GetPensandoTagsOnObjects(refs []types.ManagedObjectReference) (map[string]vcprobe.KindTagMapEntry, error) {
	ret := m.ctrl.Call(m, "GetPensandoTagsOnObjects", refs)
	ret0, _ := ret[0].(map[string]vcprobe.KindTagMapEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPensandoTagsOnObjects indicates an expected call of GetPensandoTagsOnObjects
func (mr *MockTagsProbeInfMockRecorder) GetPensandoTagsOnObjects(refs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPensandoTagsOnObjects", reflect.TypeOf((*MockTagsProbeInf)(nil).GetPensandoTagsOnObjects), refs)
}

// ResyncVMTags mocks base method
func (m *MockTagsProbeInf) ResyncVMTags(arg0 string) (defs.Probe2StoreMsg, error) {
	ret := m.ctrl.Call(m, "ResyncVMTags", arg0)
	ret0, _ := ret[0].(defs.Probe2StoreMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResyncVMTags indicates an expected call of ResyncVMTags
func (mr *MockTagsProbeInfMockRecorder) ResyncVMTags(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResyncVMTags", reflect.TypeOf((*MockTagsProbeInf)(nil).ResyncVMTags), arg0)
}

// TagObjsAsManaged mocks base method
func (m *MockTagsProbeInf) TagObjsAsManaged(refs []types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "TagObjsAsManaged", refs)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagObjsAsManaged indicates an expected call of TagObjsAsManaged
func (mr *MockTagsProbeInfMockRecorder) TagObjsAsManaged(refs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagObjsAsManaged", reflect.TypeOf((*MockTagsProbeInf)(nil).TagObjsAsManaged), refs)
}

// TagObjAsManaged mocks base method
func (m *MockTagsProbeInf) TagObjAsManaged(ref types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "TagObjAsManaged", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagObjAsManaged indicates an expected call of TagObjAsManaged
func (mr *MockTagsProbeInfMockRecorder) TagObjAsManaged(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagObjAsManaged", reflect.TypeOf((*MockTagsProbeInf)(nil).TagObjAsManaged), ref)
}

// RemoveTagObjManaged mocks base method
func (m *MockTagsProbeInf) RemoveTagObjManaged(ref types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "RemoveTagObjManaged", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTagObjManaged indicates an expected call of RemoveTagObjManaged
func (mr *MockTagsProbeInfMockRecorder) RemoveTagObjManaged(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagObjManaged", reflect.TypeOf((*MockTagsProbeInf)(nil).RemoveTagObjManaged), ref)
}

// TagObjWithVlan mocks base method
func (m *MockTagsProbeInf) TagObjWithVlan(ref types.ManagedObjectReference, vlan int) error {
	ret := m.ctrl.Call(m, "TagObjWithVlan", ref, vlan)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagObjWithVlan indicates an expected call of TagObjWithVlan
func (mr *MockTagsProbeInfMockRecorder) TagObjWithVlan(ref, vlan interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagObjWithVlan", reflect.TypeOf((*MockTagsProbeInf)(nil).TagObjWithVlan), ref, vlan)
}

// RemoveTagObjVlan mocks base method
func (m *MockTagsProbeInf) RemoveTagObjVlan(ref types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "RemoveTagObjVlan", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTagObjVlan indicates an expected call of RemoveTagObjVlan
func (mr *MockTagsProbeInfMockRecorder) RemoveTagObjVlan(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagObjVlan", reflect.TypeOf((*MockTagsProbeInf)(nil).RemoveTagObjVlan), ref)
}

// RemoveTag mocks base method
func (m *MockTagsProbeInf) RemoveTag(ref types.ManagedObjectReference, tagName string) error {
	ret := m.ctrl.Call(m, "RemoveTag", ref, tagName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTag indicates an expected call of RemoveTag
func (mr *MockTagsProbeInfMockRecorder) RemoveTag(ref, tagName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTag", reflect.TypeOf((*MockTagsProbeInf)(nil).RemoveTag), ref, tagName)
}

// RemovePensandoTags mocks base method
func (m *MockTagsProbeInf) RemovePensandoTags(ref types.ManagedObjectReference) []error {
	ret := m.ctrl.Call(m, "RemovePensandoTags", ref)
	ret0, _ := ret[0].([]error)
	return ret0
}

// RemovePensandoTags indicates an expected call of RemovePensandoTags
func (mr *MockTagsProbeInfMockRecorder) RemovePensandoTags(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePensandoTags", reflect.TypeOf((*MockTagsProbeInf)(nil).RemovePensandoTags), ref)
}

// DestroyTags mocks base method
func (m *MockTagsProbeInf) DestroyTags() error {
	ret := m.ctrl.Call(m, "DestroyTags")
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyTags indicates an expected call of DestroyTags
func (mr *MockTagsProbeInfMockRecorder) DestroyTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyTags", reflect.TypeOf((*MockTagsProbeInf)(nil).DestroyTags))
}

// IsManagedTag mocks base method
func (m *MockTagsProbeInf) IsManagedTag(tag string) bool {
	ret := m.ctrl.Call(m, "IsManagedTag", tag)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsManagedTag indicates an expected call of IsManagedTag
func (mr *MockTagsProbeInfMockRecorder) IsManagedTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManagedTag", reflect.TypeOf((*MockTagsProbeInf)(nil).IsManagedTag), tag)
}

// IsVlanTag mocks base method
func (m *MockTagsProbeInf) IsVlanTag(tag string) (int, bool) {
	ret := m.ctrl.Call(m, "IsVlanTag", tag)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsVlanTag indicates an expected call of IsVlanTag
func (mr *MockTagsProbeInfMockRecorder) IsVlanTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVlanTag", reflect.TypeOf((*MockTagsProbeInf)(nil).IsVlanTag), tag)
}

// GetWriteTags mocks base method
func (m *MockTagsProbeInf) GetWriteTags() map[string]string {
	ret := m.ctrl.Call(m, "GetWriteTags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetWriteTags indicates an expected call of GetWriteTags
func (mr *MockTagsProbeInfMockRecorder) GetWriteTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriteTags", reflect.TypeOf((*MockTagsProbeInf)(nil).GetWriteTags))
}

// SetWriteTags mocks base method
func (m *MockTagsProbeInf) SetWriteTags(arg0 map[string]string) {
	m.ctrl.Call(m, "SetWriteTags", arg0)
}

// SetWriteTags indicates an expected call of SetWriteTags
func (mr *MockTagsProbeInfMockRecorder) SetWriteTags(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteTags", reflect.TypeOf((*MockTagsProbeInf)(nil).SetWriteTags), arg0)
}

// MockProbeInf is a mock of ProbeInf interface
type MockProbeInf struct {
	ctrl     *gomock.Controller
	recorder *MockProbeInfMockRecorder
}

// MockProbeInfMockRecorder is the mock recorder for MockProbeInf
type MockProbeInfMockRecorder struct {
	mock *MockProbeInf
}

// NewMockProbeInf creates a new mock instance
func NewMockProbeInf(ctrl *gomock.Controller) *MockProbeInf {
	mock := &MockProbeInf{ctrl: ctrl}
	mock.recorder = &MockProbeInfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProbeInf) EXPECT() *MockProbeInfMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockProbeInf) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockProbeInfMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProbeInf)(nil).Start))
}

// Stop mocks base method
func (m *MockProbeInf) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockProbeInfMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProbeInf)(nil).Stop))
}

// IsSessionReady mocks base method
func (m *MockProbeInf) IsSessionReady() bool {
	ret := m.ctrl.Call(m, "IsSessionReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSessionReady indicates an expected call of IsSessionReady
func (mr *MockProbeInfMockRecorder) IsSessionReady() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSessionReady", reflect.TypeOf((*MockProbeInf)(nil).IsSessionReady))
}

// StartWatchers mocks base method
func (m *MockProbeInf) StartWatchers() {
	m.ctrl.Call(m, "StartWatchers")
}

// StartWatchers indicates an expected call of StartWatchers
func (mr *MockProbeInfMockRecorder) StartWatchers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWatchers", reflect.TypeOf((*MockProbeInf)(nil).StartWatchers))
}

// ListVM mocks base method
func (m *MockProbeInf) ListVM(dcRef *types.ManagedObjectReference) []mo.VirtualMachine {
	ret := m.ctrl.Call(m, "ListVM", dcRef)
	ret0, _ := ret[0].([]mo.VirtualMachine)
	return ret0
}

// ListVM indicates an expected call of ListVM
func (mr *MockProbeInfMockRecorder) ListVM(dcRef interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVM", reflect.TypeOf((*MockProbeInf)(nil).ListVM), dcRef)
}

// ListDC mocks base method
func (m *MockProbeInf) ListDC() []mo.Datacenter {
	ret := m.ctrl.Call(m, "ListDC")
	ret0, _ := ret[0].([]mo.Datacenter)
	return ret0
}

// ListDC indicates an expected call of ListDC
func (mr *MockProbeInfMockRecorder) ListDC() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDC", reflect.TypeOf((*MockProbeInf)(nil).ListDC))
}

// GetDCMap mocks base method
func (m *MockProbeInf) GetDCMap() map[string]mo.Datacenter {
	ret := m.ctrl.Call(m, "GetDCMap")
	ret0, _ := ret[0].(map[string]mo.Datacenter)
	return ret0
}

// GetDCMap indicates an expected call of GetDCMap
func (mr *MockProbeInfMockRecorder) GetDCMap() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDCMap", reflect.TypeOf((*MockProbeInf)(nil).GetDCMap))
}

// ListDVS mocks base method
func (m *MockProbeInf) ListDVS(dcRef *types.ManagedObjectReference) []mo.VmwareDistributedVirtualSwitch {
	ret := m.ctrl.Call(m, "ListDVS", dcRef)
	ret0, _ := ret[0].([]mo.VmwareDistributedVirtualSwitch)
	return ret0
}

// ListDVS indicates an expected call of ListDVS
func (mr *MockProbeInfMockRecorder) ListDVS(dcRef interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDVS", reflect.TypeOf((*MockProbeInf)(nil).ListDVS), dcRef)
}

// ListPG mocks base method
func (m *MockProbeInf) ListPG(dcRef *types.ManagedObjectReference) []mo.DistributedVirtualPortgroup {
	ret := m.ctrl.Call(m, "ListPG", dcRef)
	ret0, _ := ret[0].([]mo.DistributedVirtualPortgroup)
	return ret0
}

// ListPG indicates an expected call of ListPG
func (mr *MockProbeInfMockRecorder) ListPG(dcRef interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPG", reflect.TypeOf((*MockProbeInf)(nil).ListPG), dcRef)
}

// ListHosts mocks base method
func (m *MockProbeInf) ListHosts(dcRef *types.ManagedObjectReference) []mo.HostSystem {
	ret := m.ctrl.Call(m, "ListHosts", dcRef)
	ret0, _ := ret[0].([]mo.HostSystem)
	return ret0
}

// ListHosts indicates an expected call of ListHosts
func (mr *MockProbeInfMockRecorder) ListHosts(dcRef interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHosts", reflect.TypeOf((*MockProbeInf)(nil).ListHosts), dcRef)
}

// StopWatchForDC mocks base method
func (m *MockProbeInf) StopWatchForDC(dcName, dcID string) {
	m.ctrl.Call(m, "StopWatchForDC", dcName, dcID)
}

// StopWatchForDC indicates an expected call of StopWatchForDC
func (mr *MockProbeInfMockRecorder) StopWatchForDC(dcName, dcID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopWatchForDC", reflect.TypeOf((*MockProbeInf)(nil).StopWatchForDC), dcName, dcID)
}

// StartWatchForDC mocks base method
func (m *MockProbeInf) StartWatchForDC(dcName, dcID string) {
	m.ctrl.Call(m, "StartWatchForDC", dcName, dcID)
}

// StartWatchForDC indicates an expected call of StartWatchForDC
func (mr *MockProbeInfMockRecorder) StartWatchForDC(dcName, dcID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWatchForDC", reflect.TypeOf((*MockProbeInf)(nil).StartWatchForDC), dcName, dcID)
}

// GetVM mocks base method
func (m *MockProbeInf) GetVM(vmID string, retry int) (mo.VirtualMachine, error) {
	ret := m.ctrl.Call(m, "GetVM", vmID, retry)
	ret0, _ := ret[0].(mo.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVM indicates an expected call of GetVM
func (mr *MockProbeInfMockRecorder) GetVM(vmID, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVM", reflect.TypeOf((*MockProbeInf)(nil).GetVM), vmID, retry)
}

// AddPenDC mocks base method
func (m *MockProbeInf) AddPenDC(dcName string, retry int) error {
	ret := m.ctrl.Call(m, "AddPenDC", dcName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPenDC indicates an expected call of AddPenDC
func (mr *MockProbeInfMockRecorder) AddPenDC(dcName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPenDC", reflect.TypeOf((*MockProbeInf)(nil).AddPenDC), dcName, retry)
}

// RenameDC mocks base method
func (m *MockProbeInf) RenameDC(oldName, newName string, retry int) error {
	ret := m.ctrl.Call(m, "RenameDC", oldName, newName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameDC indicates an expected call of RenameDC
func (mr *MockProbeInfMockRecorder) RenameDC(oldName, newName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameDC", reflect.TypeOf((*MockProbeInf)(nil).RenameDC), oldName, newName, retry)
}

// RemovePenDC mocks base method
func (m *MockProbeInf) RemovePenDC(dcName string, retry int) error {
	ret := m.ctrl.Call(m, "RemovePenDC", dcName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePenDC indicates an expected call of RemovePenDC
func (mr *MockProbeInfMockRecorder) RemovePenDC(dcName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePenDC", reflect.TypeOf((*MockProbeInf)(nil).RemovePenDC), dcName, retry)
}

// AddPenPG mocks base method
func (m *MockProbeInf) AddPenPG(dcName, dvsName string, pgConfigSpec *types.DVPortgroupConfigSpec, equalFn vcprobe.IsPGConfigEqual, retry int) error {
	ret := m.ctrl.Call(m, "AddPenPG", dcName, dvsName, pgConfigSpec, equalFn, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPenPG indicates an expected call of AddPenPG
func (mr *MockProbeInfMockRecorder) AddPenPG(dcName, dvsName, pgConfigSpec, equalFn, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPenPG", reflect.TypeOf((*MockProbeInf)(nil).AddPenPG), dcName, dvsName, pgConfigSpec, equalFn, retry)
}

// GetPenPG mocks base method
func (m *MockProbeInf) GetPenPG(dcName, pgName string, retry int) (*object.DistributedVirtualPortgroup, error) {
	ret := m.ctrl.Call(m, "GetPenPG", dcName, pgName, retry)
	ret0, _ := ret[0].(*object.DistributedVirtualPortgroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPenPG indicates an expected call of GetPenPG
func (mr *MockProbeInfMockRecorder) GetPenPG(dcName, pgName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPenPG", reflect.TypeOf((*MockProbeInf)(nil).GetPenPG), dcName, pgName, retry)
}

// GetPGConfig mocks base method
func (m *MockProbeInf) GetPGConfig(dcName, pgName string, ps []string, retry int) (*mo.DistributedVirtualPortgroup, error) {
	ret := m.ctrl.Call(m, "GetPGConfig", dcName, pgName, ps, retry)
	ret0, _ := ret[0].(*mo.DistributedVirtualPortgroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPGConfig indicates an expected call of GetPGConfig
func (mr *MockProbeInfMockRecorder) GetPGConfig(dcName, pgName, ps, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPGConfig", reflect.TypeOf((*MockProbeInf)(nil).GetPGConfig), dcName, pgName, ps, retry)
}

// RenamePG mocks base method
func (m *MockProbeInf) RenamePG(dcName, oldName, newName string, retry int) error {
	ret := m.ctrl.Call(m, "RenamePG", dcName, oldName, newName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenamePG indicates an expected call of RenamePG
func (mr *MockProbeInfMockRecorder) RenamePG(dcName, oldName, newName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenamePG", reflect.TypeOf((*MockProbeInf)(nil).RenamePG), dcName, oldName, newName, retry)
}

// RemovePenPG mocks base method
func (m *MockProbeInf) RemovePenPG(dcName, pgName string, retry int) error {
	ret := m.ctrl.Call(m, "RemovePenPG", dcName, pgName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePenPG indicates an expected call of RemovePenPG
func (mr *MockProbeInfMockRecorder) RemovePenPG(dcName, pgName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePenPG", reflect.TypeOf((*MockProbeInf)(nil).RemovePenPG), dcName, pgName, retry)
}

// AddPenDVS mocks base method
func (m *MockProbeInf) AddPenDVS(dcName string, dvsCreateSpec *types.DVSCreateSpec, equalFn vcprobe.DVSConfigDiff, retry int) error {
	ret := m.ctrl.Call(m, "AddPenDVS", dcName, dvsCreateSpec, equalFn, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPenDVS indicates an expected call of AddPenDVS
func (mr *MockProbeInfMockRecorder) AddPenDVS(dcName, dvsCreateSpec, equalFn, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPenDVS", reflect.TypeOf((*MockProbeInf)(nil).AddPenDVS), dcName, dvsCreateSpec, equalFn, retry)
}

// RenameDVS mocks base method
func (m *MockProbeInf) RenameDVS(dcName, oldName, newName string, retry int) error {
	ret := m.ctrl.Call(m, "RenameDVS", dcName, oldName, newName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameDVS indicates an expected call of RenameDVS
func (mr *MockProbeInfMockRecorder) RenameDVS(dcName, oldName, newName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameDVS", reflect.TypeOf((*MockProbeInf)(nil).RenameDVS), dcName, oldName, newName, retry)
}

// GetPenDVS mocks base method
func (m *MockProbeInf) GetPenDVS(dcName, dvsName string, retry int) (*object.DistributedVirtualSwitch, error) {
	ret := m.ctrl.Call(m, "GetPenDVS", dcName, dvsName, retry)
	ret0, _ := ret[0].(*object.DistributedVirtualSwitch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPenDVS indicates an expected call of GetPenDVS
func (mr *MockProbeInfMockRecorder) GetPenDVS(dcName, dvsName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPenDVS", reflect.TypeOf((*MockProbeInf)(nil).GetPenDVS), dcName, dvsName, retry)
}

// UpdateDVSPortsVlan mocks base method
func (m *MockProbeInf) UpdateDVSPortsVlan(dcName, dvsName string, portsSetting vcprobe.PenDVSPortSettings, forceWrite bool, retry int) error {
	ret := m.ctrl.Call(m, "UpdateDVSPortsVlan", dcName, dvsName, portsSetting, forceWrite, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDVSPortsVlan indicates an expected call of UpdateDVSPortsVlan
func (mr *MockProbeInfMockRecorder) UpdateDVSPortsVlan(dcName, dvsName, portsSetting, forceWrite, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDVSPortsVlan", reflect.TypeOf((*MockProbeInf)(nil).UpdateDVSPortsVlan), dcName, dvsName, portsSetting, forceWrite, retry)
}

// GetPenDVSPorts mocks base method
func (m *MockProbeInf) GetPenDVSPorts(dcName, dvsName string, criteria *types.DistributedVirtualSwitchPortCriteria, retry int) ([]types.DistributedVirtualPort, error) {
	ret := m.ctrl.Call(m, "GetPenDVSPorts", dcName, dvsName, criteria, retry)
	ret0, _ := ret[0].([]types.DistributedVirtualPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPenDVSPorts indicates an expected call of GetPenDVSPorts
func (mr *MockProbeInfMockRecorder) GetPenDVSPorts(dcName, dvsName, criteria, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPenDVSPorts", reflect.TypeOf((*MockProbeInf)(nil).GetPenDVSPorts), dcName, dvsName, criteria, retry)
}

// RemovePenDVS mocks base method
func (m *MockProbeInf) RemovePenDVS(dcName, dvsName string, retry int) error {
	ret := m.ctrl.Call(m, "RemovePenDVS", dcName, dvsName, retry)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePenDVS indicates an expected call of RemovePenDVS
func (mr *MockProbeInfMockRecorder) RemovePenDVS(dcName, dvsName, retry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePenDVS", reflect.TypeOf((*MockProbeInf)(nil).RemovePenDVS), dcName, dvsName, retry)
}

// IsREST401 mocks base method
func (m *MockProbeInf) IsREST401(arg0 error) bool {
	ret := m.ctrl.Call(m, "IsREST401", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsREST401 indicates an expected call of IsREST401
func (mr *MockProbeInfMockRecorder) IsREST401(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsREST401", reflect.TypeOf((*MockProbeInf)(nil).IsREST401), arg0)
}

// StartWatch mocks base method
func (m *MockProbeInf) StartWatch() {
	m.ctrl.Call(m, "StartWatch")
}

// StartWatch indicates an expected call of StartWatch
func (mr *MockProbeInfMockRecorder) StartWatch() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWatch", reflect.TypeOf((*MockProbeInf)(nil).StartWatch))
}

// SetupBaseTags mocks base method
func (m *MockProbeInf) SetupBaseTags() bool {
	ret := m.ctrl.Call(m, "SetupBaseTags")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetupBaseTags indicates an expected call of SetupBaseTags
func (mr *MockProbeInfMockRecorder) SetupBaseTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupBaseTags", reflect.TypeOf((*MockProbeInf)(nil).SetupBaseTags))
}

// GetPensandoTagsOnObjects mocks base method
func (m *MockProbeInf) GetPensandoTagsOnObjects(refs []types.ManagedObjectReference) (map[string]vcprobe.KindTagMapEntry, error) {
	ret := m.ctrl.Call(m, "GetPensandoTagsOnObjects", refs)
	ret0, _ := ret[0].(map[string]vcprobe.KindTagMapEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPensandoTagsOnObjects indicates an expected call of GetPensandoTagsOnObjects
func (mr *MockProbeInfMockRecorder) GetPensandoTagsOnObjects(refs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPensandoTagsOnObjects", reflect.TypeOf((*MockProbeInf)(nil).GetPensandoTagsOnObjects), refs)
}

// ResyncVMTags mocks base method
func (m *MockProbeInf) ResyncVMTags(arg0 string) (defs.Probe2StoreMsg, error) {
	ret := m.ctrl.Call(m, "ResyncVMTags", arg0)
	ret0, _ := ret[0].(defs.Probe2StoreMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResyncVMTags indicates an expected call of ResyncVMTags
func (mr *MockProbeInfMockRecorder) ResyncVMTags(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResyncVMTags", reflect.TypeOf((*MockProbeInf)(nil).ResyncVMTags), arg0)
}

// TagObjsAsManaged mocks base method
func (m *MockProbeInf) TagObjsAsManaged(refs []types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "TagObjsAsManaged", refs)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagObjsAsManaged indicates an expected call of TagObjsAsManaged
func (mr *MockProbeInfMockRecorder) TagObjsAsManaged(refs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagObjsAsManaged", reflect.TypeOf((*MockProbeInf)(nil).TagObjsAsManaged), refs)
}

// TagObjAsManaged mocks base method
func (m *MockProbeInf) TagObjAsManaged(ref types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "TagObjAsManaged", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagObjAsManaged indicates an expected call of TagObjAsManaged
func (mr *MockProbeInfMockRecorder) TagObjAsManaged(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagObjAsManaged", reflect.TypeOf((*MockProbeInf)(nil).TagObjAsManaged), ref)
}

// RemoveTagObjManaged mocks base method
func (m *MockProbeInf) RemoveTagObjManaged(ref types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "RemoveTagObjManaged", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTagObjManaged indicates an expected call of RemoveTagObjManaged
func (mr *MockProbeInfMockRecorder) RemoveTagObjManaged(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagObjManaged", reflect.TypeOf((*MockProbeInf)(nil).RemoveTagObjManaged), ref)
}

// TagObjWithVlan mocks base method
func (m *MockProbeInf) TagObjWithVlan(ref types.ManagedObjectReference, vlan int) error {
	ret := m.ctrl.Call(m, "TagObjWithVlan", ref, vlan)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagObjWithVlan indicates an expected call of TagObjWithVlan
func (mr *MockProbeInfMockRecorder) TagObjWithVlan(ref, vlan interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagObjWithVlan", reflect.TypeOf((*MockProbeInf)(nil).TagObjWithVlan), ref, vlan)
}

// RemoveTagObjVlan mocks base method
func (m *MockProbeInf) RemoveTagObjVlan(ref types.ManagedObjectReference) error {
	ret := m.ctrl.Call(m, "RemoveTagObjVlan", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTagObjVlan indicates an expected call of RemoveTagObjVlan
func (mr *MockProbeInfMockRecorder) RemoveTagObjVlan(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagObjVlan", reflect.TypeOf((*MockProbeInf)(nil).RemoveTagObjVlan), ref)
}

// RemoveTag mocks base method
func (m *MockProbeInf) RemoveTag(ref types.ManagedObjectReference, tagName string) error {
	ret := m.ctrl.Call(m, "RemoveTag", ref, tagName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTag indicates an expected call of RemoveTag
func (mr *MockProbeInfMockRecorder) RemoveTag(ref, tagName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTag", reflect.TypeOf((*MockProbeInf)(nil).RemoveTag), ref, tagName)
}

// RemovePensandoTags mocks base method
func (m *MockProbeInf) RemovePensandoTags(ref types.ManagedObjectReference) []error {
	ret := m.ctrl.Call(m, "RemovePensandoTags", ref)
	ret0, _ := ret[0].([]error)
	return ret0
}

// RemovePensandoTags indicates an expected call of RemovePensandoTags
func (mr *MockProbeInfMockRecorder) RemovePensandoTags(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePensandoTags", reflect.TypeOf((*MockProbeInf)(nil).RemovePensandoTags), ref)
}

// DestroyTags mocks base method
func (m *MockProbeInf) DestroyTags() error {
	ret := m.ctrl.Call(m, "DestroyTags")
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyTags indicates an expected call of DestroyTags
func (mr *MockProbeInfMockRecorder) DestroyTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyTags", reflect.TypeOf((*MockProbeInf)(nil).DestroyTags))
}

// IsManagedTag mocks base method
func (m *MockProbeInf) IsManagedTag(tag string) bool {
	ret := m.ctrl.Call(m, "IsManagedTag", tag)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsManagedTag indicates an expected call of IsManagedTag
func (mr *MockProbeInfMockRecorder) IsManagedTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManagedTag", reflect.TypeOf((*MockProbeInf)(nil).IsManagedTag), tag)
}

// IsVlanTag mocks base method
func (m *MockProbeInf) IsVlanTag(tag string) (int, bool) {
	ret := m.ctrl.Call(m, "IsVlanTag", tag)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsVlanTag indicates an expected call of IsVlanTag
func (mr *MockProbeInfMockRecorder) IsVlanTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVlanTag", reflect.TypeOf((*MockProbeInf)(nil).IsVlanTag), tag)
}

// GetWriteTags mocks base method
func (m *MockProbeInf) GetWriteTags() map[string]string {
	ret := m.ctrl.Call(m, "GetWriteTags")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetWriteTags indicates an expected call of GetWriteTags
func (mr *MockProbeInfMockRecorder) GetWriteTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriteTags", reflect.TypeOf((*MockProbeInf)(nil).GetWriteTags))
}

// SetWriteTags mocks base method
func (m *MockProbeInf) SetWriteTags(arg0 map[string]string) {
	m.ctrl.Call(m, "SetWriteTags", arg0)
}

// SetWriteTags indicates an expected call of SetWriteTags
func (mr *MockProbeInfMockRecorder) SetWriteTags(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteTags", reflect.TypeOf((*MockProbeInf)(nil).SetWriteTags), arg0)
}
