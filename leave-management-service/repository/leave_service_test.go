package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockLeaveRepo struct {
	GetAllFunc       func() ([]Leave, error)
	GetByIDFunc      func(id int) (Leave, error)
	CreateFunc       func(leave Leave) (Leave, error)
	UpdateFunc       func(id int, leave Leave) (Leave, error)
	DeleteFunc       func(id int) error
	UpdateStatusFunc func(id int, leave Leave) (Leave, error)
	GetAllMeFunc  func(eid string) ([]Leave, error)
}

func (m *mockLeaveRepo) GetAll() ([]Leave, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *mockLeaveRepo) GetByID(id int) (Leave, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) Create(leave Leave) (Leave, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(leave)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) Update(id int, leave Leave) (Leave, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, leave)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockLeaveRepo) UpdateStatus(id int, leave Leave) (Leave, error) {
	if m.UpdateStatusFunc != nil {
		return m.UpdateStatusFunc(id, leave)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) GetAllMe(eid string) ([]Leave, error) {
	if m.GetAllMeFunc != nil {
		return m.GetAllMeFunc(eid)
	}
	return []Leave{}, errors.New("not implemented")
}

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetAllFunc: func() ([]Leave, error) {
				return []Leave{}, nil
			},
		}

		service := NewLeaveService(mockRepo)
		_, err := service.GetLeaves()
		assert.NoError(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, nil
			},
		}

		service := NewLeaveService(mockRepo)
		_, err := service.GetLeave(1)
		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			CreateFunc: func(leave Leave) (Leave, error) {
				return Leave{}, nil
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.CreateLeave(Leave{})
		assert.NoError(t, err)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.UpdateLeave(1, Leave{})
		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		}
		service := NewLeaveService(mockRepo)
		err := service.DeleteLeave(1)
		assert.NoError(t, err)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Valid-UpdateStatus", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, nil
			},
		}

		mockUpdateStatus := LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.NoError(t, err)
	})

	t.Run("Invalid-ID UpdateStatus", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, errors.New("ID not found")
			},
		}

		mockUpdateStatus := LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.Error(t, err)
	})
}
