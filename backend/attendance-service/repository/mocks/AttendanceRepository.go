// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	repository "github.com/night-sornram/employee-management/attendance-service/repository"
	mock "github.com/stretchr/testify/mock"
)

// AttendanceRepository is an autogenerated mock type for the AttendanceRepository type
type AttendanceRepository struct {
	mock.Mock
}

// CheckIn provides a mock function with given fields: eid
func (_m *AttendanceRepository) CheckIn(eid string) (repository.Attendance, error) {
	ret := _m.Called(eid)

	if len(ret) == 0 {
		panic("no return value specified for CheckIn")
	}

	var r0 repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (repository.Attendance, error)); ok {
		return rf(eid)
	}
	if rf, ok := ret.Get(0).(func(string) repository.Attendance); ok {
		r0 = rf(eid)
	} else {
		r0 = ret.Get(0).(repository.Attendance)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(eid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckOut provides a mock function with given fields: id
func (_m *AttendanceRepository) CheckOut(id int) (repository.Attendance, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for CheckOut")
	}

	var r0 repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (repository.Attendance, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) repository.Attendance); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(repository.Attendance)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckToday provides a mock function with given fields: eid
func (_m *AttendanceRepository) CheckToday(eid string) (repository.Attendance, error) {
	ret := _m.Called(eid)

	if len(ret) == 0 {
		panic("no return value specified for CheckToday")
	}

	var r0 repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (repository.Attendance, error)); ok {
		return rf(eid)
	}
	if rf, ok := ret.Get(0).(func(string) repository.Attendance); ok {
		r0 = rf(eid)
	} else {
		r0 = ret.Get(0).(repository.Attendance)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(eid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: attendance
func (_m *AttendanceRepository) Create(attendance repository.Attendance) (repository.Attendance, error) {
	ret := _m.Called(attendance)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(repository.Attendance) (repository.Attendance, error)); ok {
		return rf(attendance)
	}
	if rf, ok := ret.Get(0).(func(repository.Attendance) repository.Attendance); ok {
		r0 = rf(attendance)
	} else {
		r0 = ret.Get(0).(repository.Attendance)
	}

	if rf, ok := ret.Get(1).(func(repository.Attendance) error); ok {
		r1 = rf(attendance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *AttendanceRepository) Delete(id int) error {
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
func (_m *AttendanceRepository) GetAll(query repository.Query) (repository.DataJson, error) {
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

// GetAllLate provides a mock function with given fields:
func (_m *AttendanceRepository) GetAllLate() ([]repository.Attendance, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllLate")
	}

	var r0 []repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]repository.Attendance, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []repository.Attendance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Attendance)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllMe provides a mock function with given fields: query, eid
func (_m *AttendanceRepository) GetAllMe(query repository.Query, eid string) (repository.DataJson, error) {
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
func (_m *AttendanceRepository) GetByID(id int) (repository.Attendance, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (repository.Attendance, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) repository.Attendance); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(repository.Attendance)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCSV provides a mock function with given fields: query
func (_m *AttendanceRepository) GetCSV(query string) ([]byte, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for GetCSV")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDayLate provides a mock function with given fields:
func (_m *AttendanceRepository) GetDayLate() ([]repository.Attendance, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDayLate")
	}

	var r0 []repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]repository.Attendance, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []repository.Attendance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Attendance)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMonthLate provides a mock function with given fields: month, year
func (_m *AttendanceRepository) GetMonthLate(month int, year int) ([]repository.Attendance, error) {
	ret := _m.Called(month, year)

	if len(ret) == 0 {
		panic("no return value specified for GetMonthLate")
	}

	var r0 []repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]repository.Attendance, error)); ok {
		return rf(month, year)
	}
	if rf, ok := ret.Get(0).(func(int, int) []repository.Attendance); ok {
		r0 = rf(month, year)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Attendance)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(month, year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetYearLate provides a mock function with given fields: year
func (_m *AttendanceRepository) GetYearLate(year int) ([]repository.Attendance, error) {
	ret := _m.Called(year)

	if len(ret) == 0 {
		panic("no return value specified for GetYearLate")
	}

	var r0 []repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]repository.Attendance, error)); ok {
		return rf(year)
	}
	if rf, ok := ret.Get(0).(func(int) []repository.Attendance); ok {
		r0 = rf(year)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Attendance)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, attendance
func (_m *AttendanceRepository) Update(id int, attendance repository.Attendance) (repository.Attendance, error) {
	ret := _m.Called(id, attendance)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 repository.Attendance
	var r1 error
	if rf, ok := ret.Get(0).(func(int, repository.Attendance) (repository.Attendance, error)); ok {
		return rf(id, attendance)
	}
	if rf, ok := ret.Get(0).(func(int, repository.Attendance) repository.Attendance); ok {
		r0 = rf(id, attendance)
	} else {
		r0 = ret.Get(0).(repository.Attendance)
	}

	if rf, ok := ret.Get(1).(func(int, repository.Attendance) error); ok {
		r1 = rf(id, attendance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAttendanceRepository creates a new instance of AttendanceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAttendanceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AttendanceRepository {
	mock := &AttendanceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
