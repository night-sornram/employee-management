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
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
	app.Get("/attendance", handle.GetAttendances)

	t.Run("Valid-GetAttendances", func(t *testing.T) {
		checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
		expectedAttendance := []repository.Attendance{
			{
				ID:         1,
				EmployeeID: "E12779",
				CheckIn:    checkIn,
				CheckOut:   checkOut,
				Date:       "today",
				LeaveID:    1,
			},
		}

		mockService.On("GetAttendances").Return(expectedAttendance, nil)

		req := httptest.NewRequest("GET", "/attendance", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetAttendanceHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
	app.Get("/attendance/:id", handle.GetAttendance)

	t.Run("Valid-GetAttendance", func(t *testing.T) {
		checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
		expectedAttendance := repository.Attendance{
			ID:         1,
			EmployeeID: "E12779",
			CheckIn:    checkIn,
			CheckOut:   checkOut,
			Date:       "today",
			LeaveID:    1,
		}

		mockService.On("GetAttendance", 1).Return(expectedAttendance, nil)

		req := httptest.NewRequest("GET", "/attendance/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateAttendanceHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
	app.Post("/attendance", handle.CreateAttendance)

	checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
	checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
	expectedAttendance := repository.Attendance{
		ID:         1,
		EmployeeID: "E12779",
		CheckIn:    checkIn,
		CheckOut:   checkOut,
		Date:       "today",
		LeaveID:    1,
	}

	t.Run("Valid-CreateAttendance", func(t *testing.T) {
		mockService.On("CreateAttendance", mock.AnythingOfType("repository.Attendance")).Return(expectedAttendance, nil)

		req := httptest.NewRequest("POST", "/attendance", bytes.NewBufferString(`{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateAttendanceHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
	app.Put("/attendance/:id", handle.UpdateAttendance)

	checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
	checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
	expectedAttendance := repository.Attendance{
		ID:         1,
		EmployeeID: "E12779",
		CheckIn:    checkIn,
		CheckOut:   checkOut,
		Date:       "today",
		LeaveID:    1,
	}

	t.Run("Valid-UpdateAttendance", func(t *testing.T) {
		mockService.On("UpdateAttendance", 1, mock.AnythingOfType("repository.Attendance")).Return(expectedAttendance, nil)

		req := httptest.NewRequest("PUT", "/attendance/1", bytes.NewBufferString(`{
			"employee_id": "E12779",
			"check_in": "2024-05-14T08:00:00Z",
			"check_out": "2024-05-14T17:00:00Z",
			"date": "today",
			"leave_id": 0
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteAttendanceHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
	app.Delete("/attendance/:id", handle.DeleteAttendance)

	t.Run("Valid-DeleteAttendance", func(t *testing.T) {
		mockService.On("DeleteAttendance", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/attendance/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

// func TestUpdateStatusHandler(t *testing.T) {
// 	mockService := new(MockLeaveService)
// 	handle := NewhandlerFiber(mockService)
// 	app := fiber.New()
// 	app.Put("/attendance/approval/:id", handle.UpdateStatus)

// 	expectedAttendance := repository.Attendance{
// 		ID:             1,
// 		EmployeeID:     "E12779",
// 		DateStart:      time.Date(2024, time.May, 14, 8, 0, 0, 0, time.UTC),
// 		DateEnd:        time.Date(2024, time.May, 16, 8, 0, 0, 0, time.UTC),
// 		Reason:         "employee_reason",
// 		Status:         "approve",
// 		ManagerOpinion: "OK, approve",
// 	}

// 	t.Run("Valid-UpdateStatus", func(t *testing.T) {
// 		mockService.On("UpdateStatus", 1, mock.AnythingOfType("repository.AttendanceStatus")).Return(expectedAttendance, nil)

// 		req := httptest.NewRequest("PUT", "/attendance/approval/1", bytes.NewBufferString(`{
// 			"employee_id": "E12779",
// 			"status": "approve",
// 			"manager_opinion": "OK, approve"
// 		}`))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, err := app.Test(req)

// 		assert.NoError(t, err)
// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
// 		mockService.AssertExpectations(t)
// 	})
// }
