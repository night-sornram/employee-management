package adapter

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	mockService := new(MockLeaveService)
	handle := NewHandlerFiber(mockService)
	app := fiber.New()
	app.Get("/leaves", handle.GetLeaves)

	t.Run("Valid-GetLeaves", func(t *testing.T) {
		expectedLeaves := []repository.Leave{
			{
				ID:             1,
				EmployeeID:     "E12779",
				DateStart:      time.Date(2024, time.May, 14, 8, 0, 0, 0, time.UTC),
				DateEnd:        time.Date(2024, time.May, 16, 8, 0, 0, 0, time.UTC),
				Reason:         "Sick leave",
				Status:         "approved",
				ManagerOpinion: "",
			},
		}

		mockService.On("GetLeaves").Return(expectedLeaves, nil)

		req := httptest.NewRequest("GET", "/leaves", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetLeaveHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewHandlerFiber(mockService)
	app := fiber.New()
	app.Get("/leaves/:id", handle.GetLeave)

	t.Run("Valid-GetLeave", func(t *testing.T) {
		expectedLeave := repository.Leave{
			ID:             1,
			EmployeeID:     "E12779",
			DateStart:      time.Date(2024, time.May, 14, 8, 0, 0, 0, time.UTC),
			DateEnd:        time.Date(2024, time.May, 16, 8, 0, 0, 0, time.UTC),
			Reason:         "Sick leave",
			Status:         "approved",
			ManagerOpinion: "",
		}

		mockService.On("GetLeave", 1).Return(expectedLeave, nil)

		req := httptest.NewRequest("GET", "/leaves/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateLeaveHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewHandlerFiber(mockService)
	app := fiber.New()
	app.Post("/leaves", handle.CreateLeave)

	expectedLeave := repository.Leave{
		ID:             1,
		EmployeeID:     "E12779",
		DateStart:      time.Date(2024, time.May, 14, 8, 0, 0, 0, time.UTC),
		DateEnd:        time.Date(2024, time.May, 16, 8, 0, 0, 0, time.UTC),
		Reason:         "employee_reason",
		Status:         "pending",
		ManagerOpinion: "",
	}

	t.Run("Valid-CreateLeave", func(t *testing.T) {
		mockService.On("CreateLeave", mock.AnythingOfType("repository.Leave")).Return(expectedLeave, nil)

		req := httptest.NewRequest("POST", "/leaves", bytes.NewBufferString(`{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"status": "pending"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateLeaveHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewHandlerFiber(mockService)
	app := fiber.New()
	app.Put("/leaves/:id", handle.UpdateLeave)

	expectedLeave := repository.Leave{
		ID:             1,
		EmployeeID:     "E12779",
		DateStart:      time.Date(2024, time.May, 14, 8, 0, 0, 0, time.UTC),
		DateEnd:        time.Date(2024, time.May, 16, 8, 0, 0, 0, time.UTC),
		Reason:         "employee_reason",
		Status:         "pending",
		ManagerOpinion: "",
	}

	t.Run("Valid-UpdateLeave", func(t *testing.T) {
		mockService.On("UpdateLeave", 1, mock.AnythingOfType("repository.Leave")).Return(expectedLeave, nil)

		req := httptest.NewRequest("PUT", "/leaves/1", bytes.NewBufferString(`{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "employee_reason",
			"status": "pending"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteLeaveHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewHandlerFiber(mockService)
	app := fiber.New()
	app.Delete("/leaves/:id", handle.DeleteLeave)

	t.Run("Valid-DeleteLeave", func(t *testing.T) {
		mockService.On("DeleteLeave", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/leaves/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateStatusHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewHandlerFiber(mockService)
	app := fiber.New()
	app.Put("/leaves/approval/:id", handle.UpdateStatus)

	expectedLeave := repository.Leave{
		ID:             1,
		EmployeeID:     "E12779",
		DateStart:      time.Date(2024, time.May, 14, 8, 0, 0, 0, time.UTC),
		DateEnd:        time.Date(2024, time.May, 16, 8, 0, 0, 0, time.UTC),
		Reason:         "employee_reason",
		Status:         "approve",
		ManagerOpinion: "OK, approve",
	}

	t.Run("Valid-UpdateStatus", func(t *testing.T) {
		mockService.On("UpdateStatus", 1, mock.AnythingOfType("repository.LeaveStatus")).Return(expectedLeave, nil)

		req := httptest.NewRequest("PUT", "/leaves/approval/1", bytes.NewBufferString(`{
			"employee_id": "E12779",
			"status": "approve",
			"manager_opinion": "OK, approve"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
