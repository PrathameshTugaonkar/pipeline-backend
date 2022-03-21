// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/pipeline.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	model "github.com/instill-ai/pipeline-backend/pkg/model"
)

// MockOperations is a mock of Operations interface.
type MockOperations struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsMockRecorder
}

// MockOperationsMockRecorder is the mock recorder for MockOperations.
type MockOperationsMockRecorder struct {
	mock *MockOperations
}

// NewMockOperations creates a new mock instance.
func NewMockOperations(ctrl *gomock.Controller) *MockOperations {
	mock := &MockOperations{ctrl: ctrl}
	mock.recorder = &MockOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperations) EXPECT() *MockOperationsMockRecorder {
	return m.recorder
}

// CreatePipeline mocks base method.
func (m *MockOperations) CreatePipeline(pipeline model.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", pipeline)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockOperationsMockRecorder) CreatePipeline(pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockOperations)(nil).CreatePipeline), pipeline)
}

// DeletePipeline mocks base method.
func (m *MockOperations) DeletePipeline(namespace, pipelineName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", namespace, pipelineName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockOperationsMockRecorder) DeletePipeline(namespace, pipelineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockOperations)(nil).DeletePipeline), namespace, pipelineName)
}

// GetPipelineByName mocks base method.
func (m *MockOperations) GetPipelineByName(namespace, pipelineName string) (model.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineByName", namespace, pipelineName)
	ret0, _ := ret[0].(model.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineByName indicates an expected call of GetPipelineByName.
func (mr *MockOperationsMockRecorder) GetPipelineByName(namespace, pipelineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineByName", reflect.TypeOf((*MockOperations)(nil).GetPipelineByName), namespace, pipelineName)
}

// ListPipelines mocks base method.
func (m *MockOperations) ListPipelines(query model.ListPipelineQuery) ([]model.Pipeline, uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipelines", query)
	ret0, _ := ret[0].([]model.Pipeline)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListPipelines indicates an expected call of ListPipelines.
func (mr *MockOperationsMockRecorder) ListPipelines(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelines", reflect.TypeOf((*MockOperations)(nil).ListPipelines), query)
}

// UpdatePipeline mocks base method.
func (m *MockOperations) UpdatePipeline(pipeline model.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", pipeline)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipeline indicates an expected call of UpdatePipeline.
func (mr *MockOperationsMockRecorder) UpdatePipeline(pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockOperations)(nil).UpdatePipeline), pipeline)
}
