package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAttendanceRepo struct {
	GetAttendancesFunc   func() ([]Attendance, error)
	GetAttendanceFunc    func(id int) (Attendance, error)
	CreateAttendanceFunc func(attendance Attendance) (Attendance, error)
	UpdateAttendanceFunc func(id int, attendance Attendance) (Attendance, error)
	DeleteAttendanceFunc func(id int) error
	CheckInFunc          func(checkIn CheckIn) (Attendance, error)
	CheckOutFunc         func(checkOut CheckOut) (Attendance, error)
}

func (m *mockAttendanceRepo) GetAll()([]Attendance, error) {
	if m.GetAttendancesFunc != nil {
		return m.GetAttendancesFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *mockAttendanceRepo) GetByID(id int) (Attendance, error) {
	if m.GetAttendanceFunc != nil {
		return m.GetAttendanceFunc(id)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) Create(attendance Attendance) (Attendance, error) {
	if m.CreateAttendanceFunc != nil {
		return m.CreateAttendanceFunc(attendance)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) Update(id int, attendance Attendance) (Attendance, error) {
	if m.UpdateAttendanceFunc != nil {
		return m.UpdateAttendanceFunc(id, attendance)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) Delete(id int) error {
	if m.DeleteAttendanceFunc != nil {
		return m.DeleteAttendanceFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockAttendanceRepo) CheckIn(checkIn CheckIn) (Attendance, error) {
	if m.CheckInFunc != nil {
		return m.CheckInFunc(checkIn)
	}
	return Attendance{}, errors.New("not implemented")
}

func (m *mockAttendanceRepo) CheckOut(checkOut CheckOut) (Attendance, error) {
	if m.CheckOutFunc != nil {
		return m.CheckOutFunc(checkOut)
	}
	return Attendance{}, errors.New("not implemented")
}

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAttendancesFunc: func() ([]Attendance, error) {
				return []Attendance{}, nil
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetAttendances()
		assert.NoError(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			GetAttendanceFunc: func(id int) (Attendance, error) {
				return Attendance{}, nil
			},
		}

		service := NewAttendanceService(mockRepo)
		_, err := service.GetAttendance(1)
		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CreateAttendanceFunc: func(attendance Attendance) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CreateAttendance(Attendance{})
		assert.NoError(t, err)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			UpdateAttendanceFunc: func(id int, attendance Attendance) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.UpdateAttendance(1, Attendance{})
		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			DeleteAttendanceFunc: func(id int) error {
				return nil
			},
		}
		service := NewAttendanceService(mockRepo)
		err := service.DeleteAttendance(1)
		assert.NoError(t, err)
	})
}

func TestCheckIn(t *testing.T) {
	t.Run("Valid-CheckIn", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckInFunc: func(checkIn CheckIn) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CheckIn(CheckIn{})
		assert.NoError(t, err)
	})
}

func TestCheckOut(t *testing.T) {
	t.Run("Valid-CheckOut", func(t *testing.T) {
		mockRepo := &mockAttendanceRepo{
			CheckOutFunc: func(checkOut CheckOut) (Attendance, error) {
				return Attendance{}, nil
			},
		}
		service := NewAttendanceService(mockRepo)
		_, err := service.CheckOut(CheckOut{})
		assert.NoError(t, err)
	})
}



