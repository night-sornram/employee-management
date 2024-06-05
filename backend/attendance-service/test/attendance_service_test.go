package test

import (
	"errors"
	"github.com/night-sornram/employee-management/attendance-service/repository"
	"github.com/night-sornram/employee-management/attendance-service/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.AttendanceRepository)

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendances()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllFunc: func() ([]repository.Attendance, error) {
				return []repository.Attendance{}, errors.New("invalid")
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendances()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetByIDFunc: func(id int) (repository.Attendance, error) {
				return repository.Attendance{}, nil
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetByIDFunc: func(id int) (repository.Attendance, error) {
				return repository.Attendance{}, errors.New("invalid")
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreate(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CreateFunc: func(attendance repository.Attendance) (repository.Attendance, error) {
				return repository.Attendance{}, nil
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(repository.Attendance{})
		assert.NoError(t, err)
	})
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CreateFunc: func(attendance repository.Attendance) (repository.Attendance, error) {
				return repository.Attendance{}, errors.New("invalid")
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(repository.Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			UpdateFunc: func(id int, attendance repository.Attendance) (repository.Attendance, error) {
				return repository.Attendance{}, nil
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, repository.Attendance{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			UpdateFunc: func(id int, attendance repository.Attendance) (repository.Attendance, error) {
				return repository.Attendance{}, errors.New("invalid")
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, repository.Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			DeleteFunc: func(id int) error {
				return errors.New("invalid")
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckIn(t *testing.T) {
	t.Run("Valid-CheckIn", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckInFunc: func(eid string) (repository.Attendance, error) {
				return repository.Attendance{}, nil
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckIn("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckIn", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckInFunc: func(eid string) (repository.Attendance, error) {
				return repository.Attendance{}, errors.New("invalid")
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckIn("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckOut(t *testing.T) {
	t.Run("Valid-CheckOut", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckOutFunc: func(id int) (repository.Attendance, error) {
				return repository.Attendance{}, nil
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckOut(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckOut", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckOutFunc: func(id int) (repository.Attendance, error) {
				return repository.Attendance{}, errors.New("invalid")
			},
		}
		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckOut(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetMyAttendances(t *testing.T) {
	t.Run("Valid-GetMyAttendances", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllMeFunc: func(eid string) ([]repository.Attendance, error) {
				return []repository.Attendance{}, nil
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetMyAttendances("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMyAttendances", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAllMeFunc: func(eid string) ([]repository.Attendance, error) {
				return []repository.Attendance{}, errors.New("invalid")
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.GetMyAttendances("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckToday(t *testing.T) {
	t.Run("Valid-CheckToday", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckTodayFunc: func(eid string) (repository.Attendance, error) {
				return repository.Attendance{}, nil
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckToday("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-CheckToday", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckTodayFunc: func(eid string) (repository.Attendance, error) {
				return repository.Attendance{}, errors.New("invalid")
			},
		}

		service := repository.NewAttendanceService(mockRepo)
		_, err := service.CheckToday("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
