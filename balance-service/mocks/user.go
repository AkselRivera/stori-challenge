// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ports/user.go
//
// Generated by this command:
//
//	mockgen -source pkg/ports/user.go -destination mocks/user.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	domain "github.com/AkselRivera/stori-challenge/balance-service/pkg/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockUserService) GetBalance(userID int, startDate, endDate time.Time) (domain.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", userID, startDate, endDate)
	ret0, _ := ret[0].(domain.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockUserServiceMockRecorder) GetBalance(userID, startDate, endDate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockUserService)(nil).GetBalance), userID, startDate, endDate)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserTransactions mocks base method.
func (m *MockUserRepository) GetUserTransactions(userID int, startDate, endDate time.Time) ([]domain.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTransactions", userID, startDate, endDate)
	ret0, _ := ret[0].([]domain.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTransactions indicates an expected call of GetUserTransactions.
func (mr *MockUserRepositoryMockRecorder) GetUserTransactions(userID, startDate, endDate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTransactions", reflect.TypeOf((*MockUserRepository)(nil).GetUserTransactions), userID, startDate, endDate)
}