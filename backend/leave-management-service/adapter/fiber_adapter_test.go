package adapter

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

type MockLeaveService struct {
	mock.Mock
}

func (m *MockLeaveService) GetLeaves() ([]repository.Leave, error) {
	args := m.Called()
	return args.Get(0).([]repository.Leave), args.Error(1)
}

func (m *MockLeaveService) GetLeave(id int) (repository.Leave, error) {
	args := m.Called(id)
	return args.Get(0).(repository.Leave), args.Error(1)
}

func (m *MockLeaveService) CreateLeave(leave repository.Leave) (repository.Leave, error) {
	args := m.Called(leave)
	return args.Get(0).(repository.Leave), args.Error(1)
}

func (m *MockLeaveService) UpdateLeave(id int, leave repository.Leave) (repository.Leave, error) {
	args := m.Called(id, leave)
	return args.Get(0).(repository.Leave), args.Error(1)
}

func (m *MockLeaveService) DeleteLeave(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockLeaveService) UpdateStatus(id int, leave repository.LeaveStatus) (repository.Leave, error) {
	args := m.Called(id, leave)
	return args.Get(0).(repository.Leave), args.Error(1)
}

func (m *MockLeaveService) GetAllMe(eid string) ([]repository.Leave, error) {
	args := m.Called(eid)
	return args.Get(0).([]repository.Leave), args.Error(1)
}

func TestGetLeavesHandler(t *testing.T) {
	t.Run("Valid-GetLeaves", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/leaves", handle.GetLeaves)

		mockService.On("GetLeaves").Return([]repository.Leave{}, nil)

		req := httptest.NewRequest("GET", "/leaves", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetLeaves", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/leaves", handle.GetLeaves)

		//mock error
		mockService.On("GetLeaves").Return([]repository.Leave{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/leaves", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetLeaveHandler(t *testing.T) {
	t.Run("Valid-GetLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/leaves/:id", handle.GetLeave)

		mockService.On("GetLeave", 1).Return(repository.Leave{}, nil)

		req := httptest.NewRequest("GET", "/leaves/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-GetLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/leaves/:id", handle.GetLeave)

		//id "one" is invalid
		req := httptest.NewRequest("GET", "/leaves/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-GetLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/leaves/:id", handle.GetLeave)

		//mock error
		mockService.On("GetLeave", 1).Return(repository.Leave{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/leaves/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateLeaveHandler(t *testing.T) {
	t.Run("Valid-CreateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/leaves", handle.CreateLeave)

		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"category": "sick_leave",
			"manager_opinion": "",
			"status": "pending",
			"manager": "E10000"
		}`

		mockService.On("CreateLeave", mock.AnythingOfType("repository.Leave")).Return(repository.Leave{}, nil)

		req := httptest.NewRequest("POST", "/leaves", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CreateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/leaves", handle.CreateLeave)

		//body is missing
		req := httptest.NewRequest("POST", "/leaves", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-Validator-CreateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/leaves", handle.CreateLeave)

		//field status is missing
		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"category": "sick_leave",
			"manager_opinion": "",
			"manager": "E10000"
		}`

		req := httptest.NewRequest("POST", "/leaves", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CreateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/leaves", handle.CreateLeave)

		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"category": "sick_leave",
			"manager_opinion": "",
			"status": "pending",
			"manager": "E10000"
		}`

		//mock error
		mockService.On("CreateLeave", mock.AnythingOfType("repository.Leave")).Return(repository.Leave{}, errors.New("invalid"))

		req := httptest.NewRequest("POST", "/leaves", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateLeaveHandler(t *testing.T) {

	t.Run("Valid-UpdateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/:id", handle.UpdateLeave)

		mockService.On("UpdateLeave", 1, mock.AnythingOfType("repository.Leave")).Return(repository.Leave{}, nil)

		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"status": "pending"
		}`

		req := httptest.NewRequest("PUT", "/leaves/1", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-UpdateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/:id", handle.UpdateLeave)

		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"status": "pending"
		}`

		//id "one" is invalid
		req := httptest.NewRequest("PUT", "/leaves/one", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-BodyParser-UpdateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/:id", handle.UpdateLeave)

		//body is missing
		req := httptest.NewRequest("PUT", "/leaves/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-Validator-UpdateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/:id", handle.UpdateLeave)

		//field status is missing
		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
		}`

		req := httptest.NewRequest("PUT", "/leaves/one", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-UpdateLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/:id", handle.UpdateLeave)

		bodyLeave := `{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"status": "pending"
		}`

		//mock error
		mockService.On("UpdateLeave", 1, mock.AnythingOfType("repository.Leave")).Return(repository.Leave{}, errors.New("invalid"))

		req := httptest.NewRequest("PUT", "/leaves/1", bytes.NewBufferString(bodyLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteLeaveHandler(t *testing.T) {
	t.Run("Valid-DeleteLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/leaves/:id", handle.DeleteLeave)
		mockService.On("DeleteLeave", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/leaves/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-DeleteLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/leaves/:id", handle.DeleteLeave)

		//id "one" is invalid
		req := httptest.NewRequest("DELETE", "/leaves/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-DeleteLeave", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/leaves/:id", handle.DeleteLeave)

		//mock error
		mockService.On("DeleteLeave", 1).Return(errors.New("invalid"))

		req := httptest.NewRequest("DELETE", "/leaves/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateStatusHandler(t *testing.T) {
	t.Run("Valid-UpdateStatus", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/approval/:id", handle.UpdateStatus)

		mockService.On("UpdateStatus", 1, mock.AnythingOfType("repository.LeaveStatus")).Return(repository.Leave{}, nil)

		bodyStatusLeave := `{
			"employee_id": "E12779",
			"status": "approve",
			"manager_opinion": "OK, approve"
		}`

		req := httptest.NewRequest("PUT", "/leaves/approval/1", bytes.NewBufferString(bodyStatusLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-UpdateStatus", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/approval/:id", handle.UpdateStatus)

		bodyStatusLeave := `{
			"employee_id": "E12779",
			"status": "approve",
			"manager_opinion": "OK, approve"
		}`

		//id "one" is invalid
		req := httptest.NewRequest("PUT", "/leaves/approval/one", bytes.NewBufferString(bodyStatusLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-BodyParser-UpdateStatus", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/approval/:id", handle.UpdateStatus)

		//body is missing
		req := httptest.NewRequest("PUT", "/leaves/approval/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-Validator-UpdateStatus", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/approval/:id", handle.UpdateStatus)

		//field status is missing
		bodyStatusLeave := `{
			"employee_id": "E12779",
			"manager_opinion": "OK, approve"
		}`

		req := httptest.NewRequest("PUT", "/leaves/approval/1", bytes.NewBufferString(bodyStatusLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Valid-UpdateStatus", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/leaves/approval/:id", handle.UpdateStatus)

		mockService.On("UpdateStatus", 1, mock.AnythingOfType("repository.LeaveStatus")).Return(repository.Leave{}, errors.New("invalid"))

		bodyStatusLeave := `{
			"employee_id": "E12779",
			"status": "approve",
			"manager_opinion": "OK, approve"
		}`

		req := httptest.NewRequest("PUT", "/leaves/approval/1", bytes.NewBufferString(bodyStatusLeave))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
