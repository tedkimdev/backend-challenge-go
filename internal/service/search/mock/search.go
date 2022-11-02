// Code generated by MockGen. DO NOT EDIT.
// Source: search.go

// Package mock_search is a generated GoMock package.
package mock_search

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	eth "github.com/rarecircles/backend-challenge-go/internal/pkg/eth"
)

// MockSearchService is a mock of SearchService interface.
type MockSearchService struct {
	ctrl     *gomock.Controller
	recorder *MockSearchServiceMockRecorder
}

// MockSearchServiceMockRecorder is the mock recorder for MockSearchService.
type MockSearchServiceMockRecorder struct {
	mock *MockSearchService
}

// NewMockSearchService creates a new mock instance.
func NewMockSearchService(ctrl *gomock.Controller) *MockSearchService {
	mock := &MockSearchService{ctrl: ctrl}
	mock.recorder = &MockSearchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchService) EXPECT() *MockSearchServiceMockRecorder {
	return m.recorder
}

// SearchToken mocks base method.
func (m *MockSearchService) SearchToken(ctx context.Context, name string) ([]eth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchToken", ctx, name)
	ret0, _ := ret[0].([]eth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchToken indicates an expected call of SearchToken.
func (mr *MockSearchServiceMockRecorder) SearchToken(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchToken", reflect.TypeOf((*MockSearchService)(nil).SearchToken), ctx, name)
}
