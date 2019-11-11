// // Copyright 2019 Antrea Authors
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //      http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/agent/openflow (interfaces: FlowOperations)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	openflow "github.com/vmware-tanzu/antrea/pkg/ovs/openflow"
	reflect "reflect"
)

// MockFlowOperations is a mock of FlowOperations interface
type MockFlowOperations struct {
	ctrl     *gomock.Controller
	recorder *MockFlowOperationsMockRecorder
}

// MockFlowOperationsMockRecorder is the mock recorder for MockFlowOperations
type MockFlowOperationsMockRecorder struct {
	mock *MockFlowOperations
}

// NewMockFlowOperations creates a new mock instance
func NewMockFlowOperations(ctrl *gomock.Controller) *MockFlowOperations {
	mock := &MockFlowOperations{ctrl: ctrl}
	mock.recorder = &MockFlowOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlowOperations) EXPECT() *MockFlowOperationsMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockFlowOperations) Add(arg0 openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockFlowOperationsMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFlowOperations)(nil).Add), arg0)
}

// Delete mocks base method
func (m *MockFlowOperations) Delete(arg0 openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFlowOperationsMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFlowOperations)(nil).Delete), arg0)
}

// Modify mocks base method
func (m *MockFlowOperations) Modify(arg0 openflow.Flow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Modify indicates an expected call of Modify
func (mr *MockFlowOperationsMockRecorder) Modify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockFlowOperations)(nil).Modify), arg0)
}