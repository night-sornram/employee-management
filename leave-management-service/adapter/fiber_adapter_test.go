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
	return repository.Leave{}, nil
}

func (m *MockLeaveService) CreateLeave(leave repository.Leave) (repository.Leave, error) {
	args := m.Called(leave)
	return leave, args.Error(0)
}

func (m *MockLeaveService) UpdateLeave(id int, leave repository.Leave) (repository.Leave, error) {
	args := m.Called(leave)
	return leave, args.Error(0)
}

func (m *MockLeaveService) DeleteLeave(id int) error {
	return nil
}

func (m *MockLeaveService) UpdateStatus(id int, leave repository.Leave) (repository.Leave, error) {
	args := m.Called(leave)
	return leave, args.Error(0)
}

func TestGetLeavesHandler(t *testing.T){
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
  	app.Get("/leaves", handle.GetLeaves)

	  t.Run("Valid GetLeaves", func(t *testing.T) {
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

func TestCreateLeaveHandler(t *testing.T) {
	mockService := new(MockLeaveService)
	handle := NewhandlerFiber(mockService)
	app := fiber.New()
  	app.Post("/leaves", handle.CreateLeave)

	  t.Run("Valid leave", func(t *testing.T) {
		mockService.On("CreateLeave", mock.AnythingOfType("repository.Leave")).Return(nil)
	
		req := httptest.NewRequest("POST", "/leaves", bytes.NewBufferString(`{
			"employee_id": "E12779",
			"date_start": "2024-05-14T08:00:00Z",
			"date_end": "2024-05-16T08:00:00Z",
			"reason": "I am suffering from a severe flu with high fever, body aches, chills, and persistent coughing, and need to take 3-5 days of sick leave to recover fully and avoid spreading the illness.",
			"status": "pending"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
	
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	  })
}
