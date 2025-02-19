// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "iHR/repositories/model"

	mock "github.com/stretchr/testify/mock"
)

// ResetPasswordRepository is an autogenerated mock type for the ResetPasswordRepository type
type ResetPasswordRepository struct {
	mock.Mock
}

// CreatePasswordReset provides a mock function with given fields: reset
func (_m *ResetPasswordRepository) CreatePasswordReset(reset *model.PasswordReset) error {
	ret := _m.Called(reset)

	if len(ret) == 0 {
		panic("no return value specified for CreatePasswordReset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.PasswordReset) error); ok {
		r0 = rf(reset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByEmail provides a mock function with given fields: email
func (_m *ResetPasswordRepository) FindByEmail(email string) (*model.Account, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Account, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Account); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPasswordResetByToken provides a mock function with given fields: token
func (_m *ResetPasswordRepository) FindPasswordResetByToken(token string) (*model.PasswordReset, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for FindPasswordResetByToken")
	}

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.PasswordReset, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *model.PasswordReset); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePasswordReset provides a mock function with given fields: reset
func (_m *ResetPasswordRepository) UpdatePasswordReset(reset *model.PasswordReset) error {
	ret := _m.Called(reset)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePasswordReset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.PasswordReset) error); ok {
		r0 = rf(reset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewResetPasswordRepository creates a new instance of ResetPasswordRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResetPasswordRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResetPasswordRepository {
	mock := &ResetPasswordRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
