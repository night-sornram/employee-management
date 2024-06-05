// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	repository "github.com/night-sornram/employee-management/leave-management-service/repository"
	mock "github.com/stretchr/testify/mock"
)

// LeaveService is an autogenerated mock type for the LeaveService type
type LeaveService struct {
	mock.Mock
}

// CreateLeave provides a mock function with given fields: leave
func (_m *LeaveService) CreateLeave(leave repository.Leave) (repository.Leave, error) {
	ret := _m.Called(leave)

	if len(ret) == 0 {
		panic("no return value specified for CreateLeave")
	}

	var r0 repository.Leave
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Leave) (repository.Leave, error)); ok {
		return rf(leave)
	}
	if rf, ok := ret.Get(0).(func(repository.Leave) repository.Leave); ok {
		r0 = rf(leave)
	} else {
		r0 = ret.Get(0).(repository.Leave)
	}

	if rf, ok := ret.Get(1).(func(repository.Leave) error); ok {
		r1 = rf(leave)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLeave provides a mock function with given fields: id
func (_m *LeaveService) DeleteLeave(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLeave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllMe provides a mock function with given fields: query, eid
func (_m *LeaveService) GetAllMe(query repository.Query, eid string) (repository.DataJson, error) {
	ret := _m.Called(query, eid)

	if len(ret) == 0 {
		panic("no return value specified for GetAllMe")
	}

	var r0 repository.DataJson
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Query, string) (repository.DataJson, error)); ok {
		return rf(query, eid)
	}
	if rf, ok := ret.Get(0).(func(repository.Query, string) repository.DataJson); ok {
		r0 = rf(query, eid)
	} else {
		r0 = ret.Get(0).(repository.DataJson)
	}

	if rf, ok := ret.Get(1).(func(repository.Query, string) error); ok {
		r1 = rf(query, eid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeave provides a mock function with given fields: id
func (_m *LeaveService) GetLeave(id int) (repository.Leave, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetLeave")
	}

	var r0 repository.Leave
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (repository.Leave, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) repository.Leave); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(repository.Leave)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeaves provides a mock function with given fields: query
func (_m *LeaveService) GetLeaves(query repository.Query) (repository.DataJson, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for GetLeaves")
	}

	var r0 repository.DataJson
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Query) (repository.DataJson, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(repository.Query) repository.DataJson); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(repository.DataJson)
	}

	if rf, ok := ret.Get(1).(func(repository.Query) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLeave provides a mock function with given fields: id, leave
func (_m *LeaveService) UpdateLeave(id int, leave repository.Leave) (repository.Leave, error) {
	ret := _m.Called(id, leave)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLeave")
	}

	var r0 repository.Leave
	var r1 error
	if rf, ok := ret.Get(0).(func(int, repository.Leave) (repository.Leave, error)); ok {
		return rf(id, leave)
	}
	if rf, ok := ret.Get(0).(func(int, repository.Leave) repository.Leave); ok {
		r0 = rf(id, leave)
	} else {
		r0 = ret.Get(0).(repository.Leave)
	}

	if rf, ok := ret.Get(1).(func(int, repository.Leave) error); ok {
		r1 = rf(id, leave)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: id, leave
func (_m *LeaveService) UpdateStatus(id int, leave repository.LeaveStatus) (repository.Leave, error) {
	ret := _m.Called(id, leave)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
	}

	var r0 repository.Leave
	var r1 error
	if rf, ok := ret.Get(0).(func(int, repository.LeaveStatus) (repository.Leave, error)); ok {
		return rf(id, leave)
	}
	if rf, ok := ret.Get(0).(func(int, repository.LeaveStatus) repository.Leave); ok {
		r0 = rf(id, leave)
	} else {
		r0 = ret.Get(0).(repository.Leave)
	}

	if rf, ok := ret.Get(1).(func(int, repository.LeaveStatus) error); ok {
		r1 = rf(id, leave)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLeaveService creates a new instance of LeaveService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLeaveService(t interface {
	mock.TestingT
	Cleanup(func())
}) *LeaveService {
	mock := &LeaveService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
