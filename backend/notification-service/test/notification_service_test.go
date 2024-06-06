package test

import (
	"errors"
	"github.com/night-sornram/employee-management/notification-service/repository"
	"github.com/night-sornram/employee-management/notification-service/repository/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNotificationsService(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("GetAll").Return([]repository.Notification{}, nil)

		service := repository.NewNotificationService(mockRepo)
		_, err := service.GetNotifications()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("GetAll").Return([]repository.Notification{}, errors.New("invalid"))

		service := repository.NewNotificationService(mockRepo)
		_, err := service.GetNotifications()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetNotificationService(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("GetByID", 1).Return(repository.Notification{}, nil)

		service := repository.NewNotificationService(mockRepo)
		_, err := service.GetNotification(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("GetByID", 1).Return(repository.Notification{}, errors.New("invalid"))

		service := repository.NewNotificationService(mockRepo)
		_, err := service.GetNotification(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreateNotificationService(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Notification{}, nil)

		service := repository.NewNotificationService(mockRepo)
		_, err := service.CreateNotification(repository.Notification{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Create", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Notification{}, errors.New("invalid"))
		service := repository.NewNotificationService(mockRepo)
		_, err := service.CreateNotification(repository.Notification{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdateNotificationService(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Update", 1, mock.Anything).Return(repository.Notification{}, nil)

		service := repository.NewNotificationService(mockRepo)
		_, err := service.UpdateNotification(1, repository.Notification{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Update", 1, mock.Anything).Return(repository.Notification{}, errors.New("invalid"))

		service := repository.NewNotificationService(mockRepo)
		_, err := service.UpdateNotification(1, repository.Notification{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDeleteNotificationService(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Delete", 1).Return(nil)

		service := repository.NewNotificationService(mockRepo)
		err := service.DeleteNotification(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Delete", 1).Return(errors.New("invalid"))

		service := repository.NewNotificationService(mockRepo)
		err := service.DeleteNotification(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetNotificationByEmployeeIDService(t *testing.T) {
	t.Run("Valid-GetByEmployeeID", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("GetByEmployeeID", "EMP0001").Return([]repository.Notification{}, nil)

		service := repository.NewNotificationService(mockRepo)
		_, err := service.GetNotificationByEmployeeID("EMP0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByEmployeeID", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("GetByEmployeeID", "EMP0001").Return([]repository.Notification{}, errors.New("invalid"))

		service := repository.NewNotificationService(mockRepo)
		_, err := service.GetNotificationByEmployeeID("EMP0001")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestReadNotificationService(t *testing.T) {
	t.Run("Valid-Read", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Read", 1).Return(nil)

		service := repository.NewNotificationService(mockRepo)
		err := service.ReadNotification(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Read", func(t *testing.T) {
		mockRepo := new(mocks.NotificationRepository)
		mockRepo.On("Read", 1).Return(errors.New("invalid"))

		service := repository.NewNotificationService(mockRepo)
		err := service.ReadNotification(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
