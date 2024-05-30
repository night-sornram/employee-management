package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockNotificationRepo struct {
	GetAllFunc          func() ([]Notification, error)
	GetByIDFunc         func(id int) (Notification, error)
	CreateFunc          func(Notification Notification) (Notification, error)
	UpdateFunc          func(id int, Notification Notification) (Notification, error)
	DeleteFunc          func(id int) error
	GetByEmployeeIDFunc func(employeeID string) ([]Notification, error)
	ReadFunc            func(id int) error
}

func (m *mockNotificationRepo) GetAll() ([]Notification, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *mockNotificationRepo) GetByID(id int) (Notification, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return Notification{}, errors.New("not implemented")
}

func (m *mockNotificationRepo) Create(leave Notification) (Notification, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(leave)
	}
	return Notification{}, errors.New("not implemented")
}

func (m *mockNotificationRepo) Update(id int, leave Notification) (Notification, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, leave)
	}
	return Notification{}, errors.New("not implemented")
}

func (m *mockNotificationRepo) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockNotificationRepo) Read(id int) error {
	if m.ReadFunc != nil {
		return m.ReadFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockNotificationRepo) GetByEmployeeID(employeeID string) ([]Notification, error) {
	if m.GetByEmployeeIDFunc != nil {
		return m.GetByEmployeeIDFunc(employeeID)
	}
	return []Notification{}, errors.New("not implemented")
}

func TestGetNotifications(t *testing.T) {
	t.Run("Valid-GetNotifications", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			GetAllFunc: func() ([]Notification, error) {
				return []Notification{}, nil
			},
		}

		service := NewNotificationService(mockRepo)
		_, err := service.GetNotifications()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetNotifications", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			GetAllFunc: func() ([]Notification, error) {
				return []Notification{}, errors.New("invalid")
			},
		}

		service := NewNotificationService(mockRepo)
		_, err := service.GetNotifications()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetNotification(t *testing.T) {
	t.Run("Valid-GetNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			GetByIDFunc: func(id int) (Notification, error) {
				return Notification{}, nil
			},
		}

		service := NewNotificationService(mockRepo)
		_, err := service.GetNotification(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			GetByIDFunc: func(id int) (Notification, error) {
				return Notification{}, errors.New("invalid")
			},
		}

		service := NewNotificationService(mockRepo)
		_, err := service.GetNotification(1)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreateNotification(t *testing.T) {
	t.Run("Valid-CreateNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			CreateFunc: func(leave Notification) (Notification, error) {
				return Notification{}, nil
			},
		}
		service := NewNotificationService(mockRepo)
		_, err := service.CreateNotification(Notification{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-CreateNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			CreateFunc: func(leave Notification) (Notification, error) {
				return Notification{}, errors.New("invalid")
			},
		}
		service := NewNotificationService(mockRepo)
		_, err := service.CreateNotification(Notification{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdateNotification(t *testing.T) {
	t.Run("Valid-UpdateNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			UpdateFunc: func(id int, leave Notification) (Notification, error) {
				return Notification{}, nil
			},
		}
		service := NewNotificationService(mockRepo)
		_, err := service.UpdateNotification(1, Notification{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-UpdateNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			UpdateFunc: func(id int, leave Notification) (Notification, error) {
				return Notification{}, errors.New("invalid")
			},
		}
		service := NewNotificationService(mockRepo)
		_, err := service.UpdateNotification(1, Notification{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDeleteNotification(t *testing.T) {
	t.Run("Valid-DeleteNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		}
		service := NewNotificationService(mockRepo)
		err := service.DeleteNotification(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-DeleteNotification", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			DeleteFunc: func(id int) error {
				return errors.New("invalid")
			},
		}
		service := NewNotificationService(mockRepo)
		err := service.DeleteNotification(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetNotificationByEmployeeID(t *testing.T) {
	t.Run("Valid-GetNotificationByEmployeeID", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			GetByEmployeeIDFunc: func(employeeID string) ([]Notification, error) {
				return []Notification{}, nil
			},
		}

		service := NewNotificationService(mockRepo)
		_, err := service.GetNotificationByEmployeeID("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetNotificationByEmployeeID", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			GetByEmployeeIDFunc: func(employeeID string) ([]Notification, error) {
				return []Notification{}, errors.New("invalid")
			},
		}

		service := NewNotificationService(mockRepo)
		_, err := service.GetNotificationByEmployeeID("1")
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestReadNotification(t *testing.T) {
	t.Run("Valid-Read", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			ReadFunc: func(id int) error {
				return nil
			},
		}

		service := NewNotificationService(mockRepo)
		err := service.ReadNotification(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Read", func(t *testing.T) {
		mockRepo := &mockNotificationRepo{
			ReadFunc: func(id int) error {
				return errors.New("invalid")
			},
		}

		service := NewNotificationService(mockRepo)
		err := service.ReadNotification(1)
		assert.Equal(t, "invalid", err.Error())
	})
}
