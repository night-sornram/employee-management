package adapter

import (
	"bytes"
	"employee/repository"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

type MockLeaveService struct {
	mock.Mock
}

func (m *MockLeaveService) GetNotifications() ([]repository.Notification, error) {
	args := m.Called()
	return args.Get(0).([]repository.Notification), args.Error(1)
}

func (m *MockLeaveService) GetNotification(id int) (repository.Notification, error) {
	args := m.Called(id)
	return args.Get(0).(repository.Notification), args.Error(1)
}

func (m *MockLeaveService) CreateNotification(notification repository.Notification) (repository.Notification, error) {
	args := m.Called(notification)
	return args.Get(0).(repository.Notification), args.Error(1)
}

func (m *MockLeaveService) UpdateNotification(id int, notification repository.Notification) (repository.Notification, error) {
	args := m.Called(id, notification)
	return args.Get(0).(repository.Notification), args.Error(1)
}

func (m *MockLeaveService) DeleteNotification(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockLeaveService) GetNotificationByEmployeeID(employeeID string) ([]repository.Notification, error) {
	args := m.Called(employeeID)
	return args.Get(0).([]repository.Notification), args.Error(1)
}

func (m *MockLeaveService) ReadNotification(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetNotificationsHandler(t *testing.T) {
	t.Run("Valid-GetNotifications", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
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
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications", handle.GetNotifications)

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
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
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
	t.Run("Invalid-ID-GetNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications/:id", handle.GetNotification)

		//id "one" is invalid
		req := httptest.NewRequest("GET", "/api/notifications/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-GetNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications/:id", handle.GetNotification)

		//mock error
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
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		mockService.On("CreateNotification", mock.AnythingOfType("repository.Notification")).Return(repository.Notification{}, nil)

		//IDK why I can't set "read": false
		notification := `{
			"employee_id": "E12779",
			"message": "Test",
			"title": "Test",
			"read": true
		}`

		req := httptest.NewRequest("POST", "/api/notifications", bytes.NewBufferString(notification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CreateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		//body is missing
		req := httptest.NewRequest("POST", "/api/notifications", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-Validator-CreateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		//employee_id is missing
		notification := `{
			"message": "Test",
			"title": "Test",
			"read": true
		}`

		req := httptest.NewRequest("POST", "/api/notifications", bytes.NewBufferString(notification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CreateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/notifications", handle.CreateNotification)

		//mock error
		mockService.On("CreateNotification", mock.AnythingOfType("repository.Notification")).Return(repository.Notification{}, errors.New("invalid"))

		notification := `{
			"employee_id": "E12779",
			"message": "Test",
			"title": "Test",
			"read": true
		}`

		req := httptest.NewRequest("POST", "/api/notifications", bytes.NewBufferString(notification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateNotificationHandler(t *testing.T) {
	t.Run("Valid-UpdateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		mockService.On("UpdateNotification", 1, mock.AnythingOfType("repository.Notification")).Return(repository.Notification{}, nil)

		notification := `{
			"employee_id": "E12779",
			"message": "Test",
			"title": "Test",
			"read": true
		}`

		req := httptest.NewRequest("PUT", "/api/notifications/1", bytes.NewBufferString(notification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-UpdateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		attendance := `{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"notification_id": 0
		}`

		//ID "one" is invalid
		req := httptest.NewRequest("PUT", "/api/notifications/one", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-BodyParser-UpdateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		//body is missing
		req := httptest.NewRequest("PUT", "/api/notifications/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-Validator-UpdateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		//field employee_id is missing
		attendance := `{
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"notification_id": 0
		}`
		req := httptest.NewRequest("PUT", "/api/notifications/1", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-UpdateNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/:id", handle.UpdateNotification)

		mockService.On("UpdateNotification", 1, mock.AnythingOfType("repository.Notification")).Return(repository.Notification{}, errors.New("invalid"))

		notification := `{
			"employee_id": "E12779",
			"message": "Test",
			"title": "Test",
			"read": true
		}`

		req := httptest.NewRequest("PUT", "/api/notifications/1", bytes.NewBufferString(notification))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteNotificationHandler(t *testing.T) {
	t.Run("Valid-DeleteNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
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
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/notifications/:id", handle.DeleteNotification)

		req := httptest.NewRequest("DELETE", "/api/notifications/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-DeleteNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
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

func TestGetNotificationByEmployeeIDHandler(t *testing.T) {
	t.Run("Valid-GetNotificationByEmployeeID", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications/employee/:employeeID", handle.GetNotificationByEmployeeID)

		mockService.On("GetNotificationByEmployeeID", "E12777").Return([]repository.Notification{}, nil)

		req := httptest.NewRequest("GET", "/api/notifications/employee/E12777", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetNotificationByEmployeeID", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/notifications/employee/:employeeID", handle.GetNotificationByEmployeeID)

		//mock error
		mockService.On("GetNotificationByEmployeeID", "E12777").Return([]repository.Notification{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/notifications/employee/E12777", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestReadNotificationHandler(t *testing.T) {
	t.Run("Valid-ReadNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/read/:id", handle.ReadNotification)

		mockService.On("ReadNotification", 1).Return(nil)

		req := httptest.NewRequest("PUT", "/api/notifications/read/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-ReadNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/read/:id", handle.ReadNotification)

		//ID "one" is invalid
		req := httptest.NewRequest("PUT", "/api/notifications/read/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-ReadNotification", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewhandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/notifications/read/:id", handle.ReadNotification)

		mockService.On("ReadNotification", 1).Return(errors.New("invalid"))

		req := httptest.NewRequest("PUT", "/api/notifications/read/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
