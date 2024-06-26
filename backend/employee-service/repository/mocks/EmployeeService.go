// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	repository "github.com/night-sornram/employee-management/leave-management-service/repository"
	mock "github.com/stretchr/testify/mock"
)

// EmployeeService is an autogenerated mock type for the EmployeeService type
type EmployeeService struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: id, password, new_password
func (_m *EmployeeService) ChangePassword(id string, password string, new_password string) (repository.Employee, error) {
	ret := _m.Called(id, password, new_password)

	if len(ret) == 0 {
		panic("no return value specified for ChangePassword")
	}

	var r0 repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (repository.Employee, error)); ok {
		return rf(id, password, new_password)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) repository.Employee); ok {
		r0 = rf(id, password, new_password)
	} else {
		r0 = ret.Get(0).(repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(id, password, new_password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEmployee provides a mock function with given fields: Employee
func (_m *EmployeeService) CreateEmployee(Employee repository.Employee) (repository.Employee, error) {
	ret := _m.Called(Employee)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmployee")
	}

	var r0 repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Employee) (repository.Employee, error)); ok {
		return rf(Employee)
	}
	if rf, ok := ret.Get(0).(func(repository.Employee) repository.Employee); ok {
		r0 = rf(Employee)
	} else {
		r0 = ret.Get(0).(repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(repository.Employee) error); ok {
		r1 = rf(Employee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEmployee provides a mock function with given fields: id
func (_m *EmployeeService) DeleteEmployee(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEmployee provides a mock function with given fields: eid
func (_m *EmployeeService) GetEmployee(eid string) (repository.Employee, error) {
	ret := _m.Called(eid)

	if len(ret) == 0 {
		panic("no return value specified for GetEmployee")
	}

	var r0 repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (repository.Employee, error)); ok {
		return rf(eid)
	}
	if rf, ok := ret.Get(0).(func(string) repository.Employee); ok {
		r0 = rf(eid)
	} else {
		r0 = ret.Get(0).(repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(eid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployees provides a mock function with given fields:
func (_m *EmployeeService) GetEmployees() ([]repository.Employee, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEmployees")
	}

	var r0 []repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]repository.Employee, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []repository.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMe provides a mock function with given fields: id
func (_m *EmployeeService) GetMe(id string) (repository.Employee, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetMe")
	}

	var r0 repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (repository.Employee, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) repository.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: id, password
func (_m *EmployeeService) Login(id string, password string) (repository.Employee, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (repository.Employee, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) repository.Employee); ok {
		r0 = rf(id, password)
	} else {
		r0 = ret.Get(0).(repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields:
func (_m *EmployeeService) Logout() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEmployee provides a mock function with given fields: id, Employee
func (_m *EmployeeService) UpdateEmployee(id string, Employee repository.Employee) (repository.Employee, error) {
	ret := _m.Called(id, Employee)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmployee")
	}

	var r0 repository.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(string, repository.Employee) (repository.Employee, error)); ok {
		return rf(id, Employee)
	}
	if rf, ok := ret.Get(0).(func(string, repository.Employee) repository.Employee); ok {
		r0 = rf(id, Employee)
	} else {
		r0 = ret.Get(0).(repository.Employee)
	}

	if rf, ok := ret.Get(1).(func(string, repository.Employee) error); ok {
		r1 = rf(id, Employee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEmployeeService creates a new instance of EmployeeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmployeeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmployeeService {
	mock := &EmployeeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
