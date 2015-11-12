// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/statemanager (interfaces: StateManager)

package mock_statemanager

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of StateManager interface
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *_MockStateManagerRecorder
}

// Recorder for MockStateManager (not exported)
type _MockStateManagerRecorder struct {
	mock *MockStateManager
}

func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &_MockStateManagerRecorder{mock}
	return mock
}

func (_m *MockStateManager) EXPECT() *_MockStateManagerRecorder {
	return _m.recorder
}

func (_m *MockStateManager) ForceSave() error {
	ret := _m.ctrl.Call(_m, "ForceSave")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStateManagerRecorder) ForceSave() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForceSave")
}

func (_m *MockStateManager) Load() error {
	ret := _m.ctrl.Call(_m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStateManagerRecorder) Load() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Load")
}

func (_m *MockStateManager) Save() error {
	ret := _m.ctrl.Call(_m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStateManagerRecorder) Save() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save")
}