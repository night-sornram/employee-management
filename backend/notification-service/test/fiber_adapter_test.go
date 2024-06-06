package test

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/notification-service/adapter"
	"github.com/night-sornram/employee-management/notification-service/repository"
	"github.com/night-sornram/employee-management/notification-service/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestGetNotificationsHandler(t *testing.T) {
	t.Run("Valid-GetNotifications", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications", handle.GetNotifications)

		mockService.On("GetNotifications").Return([]repository.Notification{}, nil)

		req := httptest.NewRequest("GET", "/api/notifications", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetNotifications", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications", handle.GetNotifications)

		//mock error
		mockService.On("GetNotifications").Return([]repository.Notification{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/notifications", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetNotificationHandler(t *testing.T) {
	t.Run("Valid-GetNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications/:id", handle.GetNotification)

		mockService.On("GetNotification", 1).Return(repository.Notification{}, nil)

		req := httptest.NewRequest("GET", "/api/notifications/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications/:id", handle.GetNotification)

		mockService.On("GetNotification", 1).Return(repository.Notification{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/notifications/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateNotificationHandler(t *testing.T) {
	t.Run("Valid-CreateNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		mockService.On("CreateNotification", mock.Anything).Return(repository.Notification{}, nil)

		bodyNotification := `{
			"employee_id": "EMP0001",
			"message": "message",
			"title": "title",
			"read": true
		}`

		req := httptest.NewRequest("POST", "/api/notifications", bytes.NewBufferString(bodyNotification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req, 2000)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CreateNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		//body is missing
		req := httptest.NewRequest("POST", "/api/notifications", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CreateNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		mockService.On("CreateNotification", mock.Anything).Return(repository.Notification{}, errors.New("invalid"))

		bodyNotification := `{
			"employee_id": "EMP0001",
			"message": "message",
			"title": "title",
			"read": true
		}`

		req := httptest.NewRequest("POST", "/api/notifications", bytes.NewBufferString(bodyNotification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req, 2000)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateNotificationHandler(t *testing.T) {
	t.Run("Valid-UpdateNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		mockService.On("UpdateNotification", 1, mock.Anything).Return(repository.Notification{}, nil)

		bodyNotification := `{
			"employee_id": "EMP0001",
			"message": "message",
			"title": "title",
			"read": true
		}`

		req := httptest.NewRequest("PUT", "/api/notifications/1", bytes.NewBufferString(bodyNotification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-UpdateNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		//body is missing
		req := httptest.NewRequest("PUT", "/api/notifications/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-UpdateNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		mockService.On("UpdateNotification", 1, mock.Anything).Return(repository.Notification{}, errors.New("invalid"))

		bodyNotification := `{
			"employee_id": "EMP0001",
			"message": "message",
			"title": "title",
			"read": true
		}`

		req := httptest.NewRequest("PUT", "/api/notifications/1", bytes.NewBufferString(bodyNotification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteNotificationHandler(t *testing.T) {
	t.Run("Valid-DeleteNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/notifications/:id", handle.DeleteNotification)

		mockService.On("DeleteNotification", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/api/notifications/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-DeleteNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/notifications/:id", handle.DeleteNotification)

		//ID "one" is invalid
		req := httptest.NewRequest("DELETE", "/api/notifications/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-DeleteNotification", func(t *testing.T) {
		mockService := new(mocks.NotificationService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/notifications/:id", handle.DeleteNotification)

		mockService.On("DeleteNotification", 1).Return(errors.New("invalid"))

		req := httptest.NewRequest("DELETE", "/api/notifications/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
