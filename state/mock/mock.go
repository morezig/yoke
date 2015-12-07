// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/nanopack/yoke/state (interfaces: State,Store)

package mock_state

import (
	gomock "github.com/golang/mock/gomock"
	state "github.com/nanopack/yoke/state"
)

// Mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *_MockStateRecorder
}

// Recorder for MockState (not exported)
type _MockStateRecorder struct {
	mock *MockState
}

func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &_MockStateRecorder{mock}
	return mock
}

func (_m *MockState) EXPECT() *_MockStateRecorder {
	return _m.recorder
}

func (_m *MockState) Bounce(_param0 string) state.State {
	ret := _m.ctrl.Call(_m, "Bounce", _param0)
	ret0, _ := ret[0].(state.State)
	return ret0
}

func (_mr *_MockStateRecorder) Bounce(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bounce", arg0)
}

func (_m *MockState) GetDBRole() (string, error) {
	ret := _m.ctrl.Call(_m, "GetDBRole")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStateRecorder) GetDBRole() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDBRole")
}

func (_m *MockState) GetDataDir() (string, error) {
	ret := _m.ctrl.Call(_m, "GetDataDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStateRecorder) GetDataDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDataDir")
}

func (_m *MockState) GetRole() (string, error) {
	ret := _m.ctrl.Call(_m, "GetRole")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStateRecorder) GetRole() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRole")
}

func (_m *MockState) HasSynced() (bool, error) {
	ret := _m.ctrl.Call(_m, "HasSynced")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStateRecorder) HasSynced() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasSynced")
}

func (_m *MockState) Location() string {
	ret := _m.ctrl.Call(_m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockStateRecorder) Location() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Location")
}

func (_m *MockState) Ready() {
	_m.ctrl.Call(_m, "Ready")
}

func (_mr *_MockStateRecorder) Ready() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ready")
}

func (_m *MockState) SetDBRole(_param0 string) error {
	ret := _m.ctrl.Call(_m, "SetDBRole", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStateRecorder) SetDBRole(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDBRole", arg0)
}

func (_m *MockState) SetSynced(_param0 bool) error {
	ret := _m.ctrl.Call(_m, "SetSynced", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStateRecorder) SetSynced(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSynced", arg0)
}

// Mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *_MockStoreRecorder
}

// Recorder for MockStore (not exported)
type _MockStoreRecorder struct {
	mock *MockStore
}

func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &_MockStoreRecorder{mock}
	return mock
}

func (_m *MockStore) EXPECT() *_MockStoreRecorder {
	return _m.recorder
}

func (_m *MockStore) Read(_param0 string, _param1 string, _param2 interface{}) error {
	ret := _m.ctrl.Call(_m, "Read", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoreRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1, arg2)
}

func (_m *MockStore) Write(_param0 string, _param1 string, _param2 interface{}) error {
	ret := _m.ctrl.Call(_m, "Write", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoreRecorder) Write(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2)
}