// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "iHR/db/model"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: username, password
func (_m *AccountRepository) Authenticate(username string, password string) (*model.Account, error) {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for Authenticate")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.Account, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.Account); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccount provides a mock function with given fields: _a0
func (_m *AccountRepository) CreateAccount(_a0 *model.Account) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
