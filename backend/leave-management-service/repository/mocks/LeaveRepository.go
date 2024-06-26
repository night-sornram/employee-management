// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	repository "github.com/night-sornram/employee-management/leave-management-service/repository"
	mock "github.com/stretchr/testify/mock"
)

// LeaveRepository is an autogenerated mock type for the LeaveRepository type
type LeaveRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: leave
func (_m *LeaveRepository) Create(leave repository.Leave) (repository.Leave, error) {
	ret := _m.Called(leave)

	if len(ret) == 0 {
		panic("no return value specified for Create")
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

// Delete provides a mock function with given fields: id
func (_m *LeaveRepository) Delete(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: query
func (_m *LeaveRepository) GetAll(query repository.Query) (repository.DataJson, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
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

// GetAllMe provides a mock function with given fields: query, eid
func (_m *LeaveRepository) GetAllMe(query repository.Query, eid string) (repository.DataJson, error) {
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

// GetByID provides a mock function with given fields: id
func (_m *LeaveRepository) GetByID(id int) (repository.Leave, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
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

// Update provides a mock function with given fields: id, Leave
func (_m *LeaveRepository) Update(id int, Leave repository.Leave) (repository.Leave, error) {
	ret := _m.Called(id, Leave)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 repository.Leave
	var r1 error
	if rf, ok := ret.Get(0).(func(int, repository.Leave) (repository.Leave, error)); ok {
		return rf(id, Leave)
	}
	if rf, ok := ret.Get(0).(func(int, repository.Leave) repository.Leave); ok {
		r0 = rf(id, Leave)
	} else {
		r0 = ret.Get(0).(repository.Leave)
	}

	if rf, ok := ret.Get(1).(func(int, repository.Leave) error); ok {
		r1 = rf(id, Leave)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: id, leave
func (_m *LeaveRepository) UpdateStatus(id int, leave repository.Leave) (repository.Leave, error) {
	ret := _m.Called(id, leave)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatus")
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

// NewLeaveRepository creates a new instance of LeaveRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLeaveRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LeaveRepository {
	mock := &LeaveRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
