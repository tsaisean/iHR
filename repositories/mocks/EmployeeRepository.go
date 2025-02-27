// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	model "iHR/repositories/model"

	mock "github.com/stretchr/testify/mock"

	repositories "iHR/repositories"
)

// EmployeeRepository is an autogenerated mock type for the EmployeeRepository type
type EmployeeRepository struct {
	mock.Mock
}

// Autocomplete provides a mock function with given fields: ctx, query
func (_m *EmployeeRepository) Autocomplete(ctx context.Context, query string) ([]repositories.Suggestion, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Autocomplete")
	}

	var r0 []repositories.Suggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]repositories.Suggestion, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []repositories.Suggestion); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repositories.Suggestion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEmployee provides a mock function with given fields: ctx, employee
func (_m *EmployeeRepository) CreateEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error) {
	ret := _m.Called(ctx, employee)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmployee")
	}

	var r0 *model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Employee) (*model.Employee, error)); ok {
		return rf(ctx, employee)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Employee) *model.Employee); ok {
		r0 = rf(ctx, employee)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Employee) error); ok {
		r1 = rf(ctx, employee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEmployee provides a mock function with given fields: ctx, id
func (_m *EmployeeRepository) DeleteEmployee(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllEmployeesAfter provides a mock function with given fields: ctx, id, pageSize
func (_m *EmployeeRepository) GetAllEmployeesAfter(ctx context.Context, id int, pageSize int) ([]model.Employee, error) {
	ret := _m.Called(ctx, id, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for GetAllEmployeesAfter")
	}

	var r0 []model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]model.Employee, error)); ok {
		return rf(ctx, id, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []model.Employee); ok {
		r0 = rf(ctx, id, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, id, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllEmployeesFrom provides a mock function with given fields: ctx, offset, pageSize
func (_m *EmployeeRepository) GetAllEmployeesFrom(ctx context.Context, offset int, pageSize int) ([]model.Employee, error) {
	ret := _m.Called(ctx, offset, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for GetAllEmployeesFrom")
	}

	var r0 []model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]model.Employee, error)); ok {
		return rf(ctx, offset, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []model.Employee); ok {
		r0 = rf(ctx, offset, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, offset, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployeeByAccID provides a mock function with given fields: id
func (_m *EmployeeRepository) GetEmployeeByAccID(id uint) (*model.Employee, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetEmployeeByAccID")
	}

	var r0 *model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*model.Employee, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *model.Employee); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployeeByID provides a mock function with given fields: ctx, id
func (_m *EmployeeRepository) GetEmployeeByID(ctx context.Context, id uint) (*model.Employee, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetEmployeeByID")
	}

	var r0 *model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*model.Employee, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *model.Employee); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotal provides a mock function with given fields:
func (_m *EmployeeRepository) GetTotal() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotal")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEmployeeByID provides a mock function with given fields: ctx, id, updated
func (_m *EmployeeRepository) UpdateEmployeeByID(ctx context.Context, id uint, updated *model.Employee) (*model.Employee, error) {
	ret := _m.Called(ctx, id, updated)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmployeeByID")
	}

	var r0 *model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, *model.Employee) (*model.Employee, error)); ok {
		return rf(ctx, id, updated)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, *model.Employee) *model.Employee); ok {
		r0 = rf(ctx, id, updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, *model.Employee) error); ok {
		r1 = rf(ctx, id, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEmployeeRepository creates a new instance of EmployeeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmployeeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmployeeRepository {
	mock := &EmployeeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
