package test

import (
	"errors"
	"github.com/night-sornram/employee-management/attendance-service/repository"
	"github.com/night-sornram/employee-management/attendance-service/repository/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllService(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetAll", mock.Anything).Return(repository.DataJson{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendances(repository.Query{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetAll", mock.Anything).Return(repository.DataJson{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendances(repository.Query{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetByIDService(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetByID", 1).Return(repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetByID", 1).Return(repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreateService(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(repository.Attendance{})
		assert.NoError(t, err)
	})
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(repository.Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdateService(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("Update", 1, mock.Anything).Return(repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, repository.Attendance{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("Update", 1, mock.Anything).Return(repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, repository.Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDeleteService(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("Delete", 1).Return(nil)

		service := repository.NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("Delete", 1).Return(errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckInService(t *testing.T) {
	t.Run("Valid-CheckIn", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("CheckIn", "EMP0001").Return(repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckIn("EMP0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckIn", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("CheckIn", "EMP0001").Return(repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckIn("EMP0001")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckOutService(t *testing.T) {
	t.Run("Valid-CheckOut", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("CheckOut", 1).Return(repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckOut(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckOut", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("CheckOut", 1).Return(repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckOut(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetMyAttendancesService(t *testing.T) {
	t.Run("Valid-GetMyAttendances", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetAllMe", mock.Anything, "EMP0001").Return(repository.DataJson{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetMyAttendances(repository.Query{}, "EMP0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMyAttendances", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetAllMe", mock.Anything, "EMP0001").Return(repository.DataJson{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetMyAttendances(repository.Query{}, "EMP0001")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckTodayService(t *testing.T) {
	t.Run("Valid-CheckToday", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("CheckToday", "EMP0001").Return(repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckToday("EMP0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckToday", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("CheckToday", "EMP0001").Return(repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckToday("EMP0001")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetDayLateService(t *testing.T) {
	t.Run("Valid-GetDayLate", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetDayLate").Return([]repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetDayLate()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetDayLate", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetDayLate").Return([]repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetDayLate()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetMonthLateService(t *testing.T) {
	t.Run("Valid-GetMonthLate", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetMonthLate", mock.Anything, mock.Anything).Return([]repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetMonthLate(repository.GetMonth{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMonthLate", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetMonthLate", mock.Anything, mock.Anything).Return([]repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetMonthLate(repository.GetMonth{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetYearLateService(t *testing.T) {
	t.Run("Valid-GetYearLate", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetYearLate", 2024).Return([]repository.Attendance{}, nil)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetYearLate(2024)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetYearLate", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)
		mockRepo.On("GetYearLate", 2024).Return([]repository.Attendance{}, errors.New("invalid"))

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetYearLate(2024)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
