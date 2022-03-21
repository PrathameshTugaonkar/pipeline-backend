// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/protogen-go/model (interfaces: ModelClient)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	model "github.com/instill-ai/protogen-go/model/v1alpha"
)

// MockModelClient is a mock of ModelClient interface.
type MockModelClient struct {
	ctrl     *gomock.Controller
	recorder *MockModelClientMockRecorder
}

// MockModelClientMockRecorder is the mock recorder for MockModelClient.
type MockModelClientMockRecorder struct {
	mock *MockModelClient
}

// NewMockModelClient creates a new mock instance.
func NewMockModelClient(ctrl *gomock.Controller) *MockModelClient {
	mock := &MockModelClient{ctrl: ctrl}
	mock.recorder = &MockModelClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelClient) EXPECT() *MockModelClientMockRecorder {
	return m.recorder
}

// Liveness mocks base method.
func (m *MockModelClient) Liveness(arg0 context.Context, arg1 *model.LivenessRequest, arg2 ...grpc.CallOption) (*model.LivenessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Liveness", varargs...)
	ret0, _ := ret[0].(*model.LivenessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Liveness indicates an expected call of Liveness.
func (mr *MockModelClientMockRecorder) Liveness(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Liveness", reflect.TypeOf((*MockModelClient)(nil).Liveness), varargs...)
}

// Readiness mocks base method.
func (m *MockModelClient) Readiness(arg0 context.Context, arg1 *model.ReadinessRequest, arg2 ...grpc.CallOption) (*model.ReadinessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Readiness", varargs...)
	ret0, _ := ret[0].(*model.ReadinessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readiness indicates an expected call of Readiness.
func (mr *MockModelClientMockRecorder) Readiness(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readiness", reflect.TypeOf((*MockModelClient)(nil).Readiness), varargs...)
}

// CreateModel mocks base method.
func (m *MockModelClient) CreateModel(arg0 context.Context, arg1 *model.CreateModelBinaryFileUploadRequest, arg2 ...grpc.CallOption) (*model.CreateModelBinaryFileUploadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateModel", varargs...)
	ret0, _ := ret[0].(*model.CreateModelBinaryFileUploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateModel indicates an expected call of CreateModel.
func (mr *MockModelClientMockRecorder) CreateModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModel", reflect.TypeOf((*MockModelClient)(nil).CreateModel), varargs...)
}

// CreateModelBinaryFileUpload mocks base method.
func (m *MockModelClient) CreateModelBinaryFileUpload(arg0 context.Context, arg1 ...grpc.CallOption) (model.ModelService_CreateModelBinaryFileUploadClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateModelByUpload", varargs...)
	ret0, _ := ret[0].(model.ModelService_CreateModelBinaryFileUploadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateModelBinaryFileUpload indicates an expected call of CreateModelBinaryFileUpload.
func (mr *MockModelClientMockRecorder) CreateModelBinaryFileUpload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModelBinaryFileUpload", reflect.TypeOf((*MockModelClient)(nil).CreateModelBinaryFileUpload), varargs...)
}

// DeleteModel mocks base method.
func (m *MockModelClient) DeleteModel(arg0 context.Context, arg1 *model.DeleteModelRequest, arg2 ...grpc.CallOption) (*model.DeleteModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteModel", varargs...)
	ret0, _ := ret[0].(*model.DeleteModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteModel indicates an expected call of DeleteModel.
func (mr *MockModelClientMockRecorder) DeleteModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteModel", reflect.TypeOf((*MockModelClient)(nil).DeleteModel), varargs...)
}

// DeleteModelVersion mocks base method.
func (m *MockModelClient) DeleteModelVersion(arg0 context.Context, arg1 *model.DeleteModelVersionRequest, arg2 ...grpc.CallOption) (*model.DeleteModelVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteModelVersion", varargs...)
	ret0, _ := ret[0].(*model.DeleteModelVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteModelVersion indicates an expected call of DeleteModelVersion.
func (mr *MockModelClientMockRecorder) DeleteModelVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteModelVersion", reflect.TypeOf((*MockModelClient)(nil).DeleteModelVersion), varargs...)
}

// GetModel mocks base method.
func (m *MockModelClient) GetModel(arg0 context.Context, arg1 *model.GetModelRequest, arg2 ...grpc.CallOption) (*model.GetModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModel", varargs...)
	ret0, _ := ret[0].(*model.GetModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModel indicates an expected call of GetModel.
func (mr *MockModelClientMockRecorder) GetModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModel", reflect.TypeOf((*MockModelClient)(nil).GetModel), varargs...)
}

// ListModel mocks base method.
func (m *MockModelClient) ListModel(arg0 context.Context, arg1 *model.ListModelRequest, arg2 ...grpc.CallOption) (*model.ListModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListModel", varargs...)
	ret0, _ := ret[0].(*model.ListModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModel indicates an expected call of ListModel.
func (mr *MockModelClientMockRecorder) ListModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModel", reflect.TypeOf((*MockModelClient)(nil).ListModel), varargs...)
}

// TriggerModel mocks base method.
func (m *MockModelClient) TriggerModel(arg0 context.Context, arg1 *model.TriggerModelRequest, arg2 ...grpc.CallOption) (*model.TriggerModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerModel", varargs...)
	ret0, _ := ret[0].(*model.TriggerModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerModel indicates an expected call of TriggerModel.
func (mr *MockModelClientMockRecorder) TriggerModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerModel", reflect.TypeOf((*MockModelClient)(nil).TriggerModel), varargs...)
}

// TriggerModelBinaryFileUpload mocks base method.
func (m *MockModelClient) TriggerModelBinaryFileUpload(arg0 context.Context, arg1 ...grpc.CallOption) (model.ModelService_TriggerModelBinaryFileUploadClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PredictModelByUpload", varargs...)
	ret0, _ := ret[0].(model.ModelService_TriggerModelBinaryFileUploadClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerModelBinaryFileUpload indicates an expected call of PredictModelByUpload.
func (mr *MockModelClientMockRecorder) TriggerModelBinaryFileUpload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PredictModelByUpload", reflect.TypeOf((*MockModelClient)(nil).TriggerModelBinaryFileUpload), varargs...)
}

// UpdateModelVersion mocks base method.
func (m *MockModelClient) UpdateModelVersion(arg0 context.Context, arg1 *model.UpdateModelVersionRequest, arg2 ...grpc.CallOption) (*model.UpdateModelVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateModelVersion", varargs...)
	ret0, _ := ret[0].(*model.UpdateModelVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateModelVersion indicates an expected call of UpdateModelVersion.
func (mr *MockModelClientMockRecorder) UpdateModelVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateModelVersion", reflect.TypeOf((*MockModelClient)(nil).UpdateModelVersion), varargs...)
}
