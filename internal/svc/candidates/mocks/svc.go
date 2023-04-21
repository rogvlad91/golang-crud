// Code generated by MockGen. DO NOT EDIT.
// Source: ./types.go

// Package mock_candidate_svc is a generated GoMock package.
package mock_candidate_svc

import (
	context "context"
	candidates "homework-7/internal/svc/candidates"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCandidateProcessor is a mock of CandidateProcessor interface.
type MockCandidateProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockCandidateProcessorMockRecorder
}

// MockCandidateProcessorMockRecorder is the mock recorder for MockCandidateProcessor.
type MockCandidateProcessorMockRecorder struct {
	mock *MockCandidateProcessor
}

// NewMockCandidateProcessor creates a new mock instance.
func NewMockCandidateProcessor(ctrl *gomock.Controller) *MockCandidateProcessor {
	mock := &MockCandidateProcessor{ctrl: ctrl}
	mock.recorder = &MockCandidateProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCandidateProcessor) EXPECT() *MockCandidateProcessorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCandidateProcessor) Create(ctx context.Context, createDTO candidates.CreateCandidateDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, createDTO)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCandidateProcessorMockRecorder) Create(ctx, createDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCandidateProcessor)(nil).Create), ctx, createDTO)
}

// Delete mocks base method.
func (m *MockCandidateProcessor) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCandidateProcessorMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCandidateProcessor)(nil).Delete), ctx, id)
}

// GetAll mocks base method.
func (m *MockCandidateProcessor) GetAll(ctx context.Context) ([]*candidates.Candidate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*candidates.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCandidateProcessorMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCandidateProcessor)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockCandidateProcessor) GetById(ctx context.Context, id string) (*candidates.Candidate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*candidates.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCandidateProcessorMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCandidateProcessor)(nil).GetById), ctx, id)
}

// Update mocks base method.
func (m *MockCandidateProcessor) Update(ctx context.Context, id string, dto candidates.UpdateCandidateDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCandidateProcessorMockRecorder) Update(ctx, id, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCandidateProcessor)(nil).Update), ctx, id, dto)
}