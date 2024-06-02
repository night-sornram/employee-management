package adapter

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

type MockLeaveService struct {
	mock.Mock
}

func (m *MockLeaveService) GetAttendances() ([]repository.Attendance, error) {
	args := m.Called()
	return args.Get(0).([]repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) GetAttendance(id int) (repository.Attendance, error) {
	args := m.Called(id)
	return args.Get(0).(repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) CreateAttendance(leave repository.Attendance) (repository.Attendance, error) {
	args := m.Called(leave)
	return args.Get(0).(repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) UpdateAttendance(id int, leave repository.Attendance) (repository.Attendance, error) {
	args := m.Called(id, leave)
	return args.Get(0).(repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) DeleteAttendance(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockLeaveService) CheckIn(eid string) (repository.Attendance, error) {
	args := m.Called(eid)
	return args.Get(0).(repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) CheckOut(id int) (repository.Attendance, error) {
	args := m.Called(id)
	return args.Get(0).(repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) CheckToday(eid string) (repository.Attendance, error) {
	args := m.Called(eid)
	return args.Get(0).(repository.Attendance), args.Error(1)
}

func (m *MockLeaveService) GetMyAttendances(eid string) ([]repository.Attendance, error) {
	args := m.Called(eid)
	return args.Get(0).([]repository.Attendance), args.Error(1)
}

func TestGetAttendancesHandler(t *testing.T) {
	t.Run("Valid-GetAttendances", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances", handle.GetAttendances)

		mockService.On("GetAttendances").Return([]repository.Attendance{}, nil)

		req := httptest.NewRequest("GET", "/api/attendances", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetAttendances", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances", handle.GetAttendances)

		mockService.On("GetAttendances").Return([]repository.Attendance{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/attendances", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAttendanceHandler(t *testing.T) {
	t.Run("Valid-GetAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/:id", handle.GetAttendance)

		mockService.On("GetAttendance", 1).Return(repository.Attendance{}, nil)

		req := httptest.NewRequest("GET", "/api/attendances/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-GetAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/:id", handle.GetAttendance)

		//id "one" is invalid
		req := httptest.NewRequest("GET", "/api/attendances/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-GetAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/:id", handle.GetAttendance)

		//mock error
		mockService.On("GetAttendance", 1).Return(repository.Attendance{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/attendances/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateAttendanceHandler(t *testing.T) {
	t.Run("Valid-CreateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances", handle.CreateAttendance)

		mockService.On("CreateAttendance", mock.AnythingOfType("repository.Attendance")).Return(repository.Attendance{}, nil)

		attendance := `{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`

		req := httptest.NewRequest("POST", "/api/attendances", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CreateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances", handle.CreateAttendance)

		//body is missing
		req := httptest.NewRequest("POST", "/api/attendances", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-Validator-CreateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances", handle.CreateAttendance)

		//employee_id is missing
		attendance := `{
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`

		req := httptest.NewRequest("POST", "/api/attendances", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CreateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances", handle.CreateAttendance)

		//mock error
		mockService.On("CreateAttendance", mock.AnythingOfType("repository.Attendance")).Return(repository.Attendance{}, errors.New("invalid"))

		attendance := `{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`

		req := httptest.NewRequest("POST", "/api/attendances", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateAttendanceHandler(t *testing.T) {
	t.Run("Valid-UpdateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/:id", handle.UpdateAttendance)

		mockService.On("UpdateAttendance", 1, mock.AnythingOfType("repository.Attendance")).Return(repository.Attendance{}, nil)

		attendance := `{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`
		req := httptest.NewRequest("PUT", "/api/attendances/1", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-UpdateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/:id", handle.UpdateAttendance)

		attendance := `{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`

		//ID "one" is invalid
		req := httptest.NewRequest("PUT", "/api/attendances/one", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-BodyParser-UpdateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/:id", handle.UpdateAttendance)

		//body is missing
		req := httptest.NewRequest("PUT", "/api/attendances/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-Validator-UpdateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/:id", handle.UpdateAttendance)

		//field employee_id is missing
		attendance := `{
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`
		req := httptest.NewRequest("PUT", "/api/attendances/1", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-UpdateAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/:id", handle.UpdateAttendance)

		mockService.On("UpdateAttendance", 1, mock.AnythingOfType("repository.Attendance")).Return(repository.Attendance{}, errors.New("invalid"))

		attendance := `{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`
		req := httptest.NewRequest("PUT", "/api/attendances/1", bytes.NewBufferString(attendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAttendanceHandler(t *testing.T) {
	t.Run("Valid-DeleteAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/attendances/:id", handle.DeleteAttendance)

		mockService.On("DeleteAttendance", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/api/attendances/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-DeleteAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/attendances/:id", handle.DeleteAttendance)

		req := httptest.NewRequest("DELETE", "/api/attendances/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-DeleteAttendance", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Delete("/api/attendances/:id", handle.DeleteAttendance)

		mockService.On("DeleteAttendance", 1).Return(errors.New("invalid"))

		req := httptest.NewRequest("DELETE", "/api/attendances/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCheckInHandler(t *testing.T) {
	t.Run("Valid-CheckIn", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances/check-in", handle.CheckIn)

		bodyEid := `{"eid": "E12777"}`

		mockService.On("CheckIn", mock.Anything).Return(repository.Attendance{}, nil)

		req := httptest.NewRequest("POST", "/api/attendances/check-in", bytes.NewBufferString(bodyEid))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CheckIn", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances/check-in", handle.CheckIn)

		//body is missing
		req := httptest.NewRequest("POST", "/api/attendances/check-in", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CheckIn", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Post("/api/attendances/check-in", handle.CheckIn)

		bodyEid := `{"eid": "E12777"}`

		//mock error
		mockService.On("CheckIn", mock.Anything).Return(repository.Attendance{}, errors.New("invalid"))

		req := httptest.NewRequest("POST", "/api/attendances/check-in", bytes.NewBufferString(bodyEid))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

}

func TestCheckOutHandler(t *testing.T) {
	t.Run("Valid-CheckOut", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/check-out", handle.CheckOut)

		mockService.On("CheckOut", mock.Anything).Return(repository.Attendance{}, nil)

		bodyID := `{
			"id": 1
		}`
		req := httptest.NewRequest("PUT", "/api/attendances/check-out", bytes.NewBufferString(bodyID))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CheckOut", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/check-out", handle.CheckOut)

		//body is missing
		req := httptest.NewRequest("PUT", "/api/attendances/check-out", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CheckOut", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Put("/api/attendances/check-out", handle.CheckOut)

		mockService.On("CheckOut", mock.Anything).Return(repository.Attendance{}, errors.New("invalid"))

		bodyID := `{
			"id": 1
		}`
		req := httptest.NewRequest("PUT", "/api/attendances/check-out", bytes.NewBufferString(bodyID))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetMyAttendancesHandler(t *testing.T) {
	t.Run("Valid-GetMyAttendances", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/me/:eid", handle.GetMyAttendances)

		mockService.On("GetMyAttendances", "E12777").Return([]repository.Attendance{}, nil)

		req := httptest.NewRequest("GET", "/api/attendances/me/E12777", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-GetMyAttendances", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/me/", handle.GetMyAttendances)

		//eid is missing
		req := httptest.NewRequest("GET", "/api/attendances/me", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-GetMyAttendances", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/me/:eid", handle.GetMyAttendances)

		//mock error
		mockService.On("GetMyAttendances", "E12777").Return([]repository.Attendance{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/attendances/me/E12777", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCheckTodayHandler(t *testing.T) {
	t.Run("Valid-CheckToday", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/check-today/:eid", handle.CheckToday)

		mockService.On("CheckToday", "E12777").Return(repository.Attendance{
			ID: 1,
		}, nil)

		req := httptest.NewRequest("GET", "/api/attendances/check-today/E12777", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Valid-ID-0-CheckToday", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/check-today/:eid", handle.CheckToday)

		mockService.On("CheckToday", "E12777").Return(repository.Attendance{
			ID: 0,
		}, nil)

		req := httptest.NewRequest("GET", "/api/attendances/check-today/E12777", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-CheckToday", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/check-today/", handle.CheckToday)

		//id is missing
		req := httptest.NewRequest("GET", "/api/attendances/check-today/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CheckToday", func(t *testing.T) {
		mockService := new(MockLeaveService)
		handle := NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/attendances/check-today/:eid", handle.CheckToday)

		//mock error
		mockService.On("CheckToday", "E1277").Return(repository.Attendance{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/attendances/check-today/E1277", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
