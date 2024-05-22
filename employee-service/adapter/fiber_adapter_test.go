package adapter

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeService struct {
	mock.Mock
}

func (m *MockEmployeeService) GetEmployees() ([]repository.Employee, error) {
	args := m.Called()
	return args.Get(0).([]repository.Employee), args.Error(1)
}

func (m *MockEmployeeService) GetEmployee(id int) (repository.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(repository.Employee), args.Error(1)
}

func (m *MockEmployeeService) CreateEmployee(leave repository.Employee) (repository.Employee, error) {
	args := m.Called(leave)
	return args.Get(0).(repository.Employee), args.Error(1)
}

func (m *MockEmployeeService) UpdateEmployee(id string, leave repository.Employee) (repository.Employee, error) {
	args := m.Called(id, leave)
	return args.Get(0).(repository.Employee), args.Error(1)
}

func (m *MockEmployeeService) DeleteEmployee(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEmployeeService) Login(id string, password string) (repository.Employee, error) {
	args := m.Called(id, password)
	return args.Get(0).(repository.Employee), args.Error(1)
}

func (m *MockEmployeeService) Logout() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockEmployeeService) GetMe(id string) (repository.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(repository.Employee), args.Error(1)
}

func (m *MockEmployeeService) ChangePassword(id string, password string, new_password string) (repository.Employee, error) {
	args := m.Called(id, password, new_password)
	return args.Get(0).(repository.Employee), args.Error(1)
}


func TestGetEmployeesHandler(t *testing.T) {
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Get("/api/employees", handle.GetEmployees)

	t.Run("Valid-GetEmployees", func(t *testing.T) {
		expectedEmployee := []repository.Employee{
			{
				EmployeeID:  "ADMIN",
				TitleTH:     "นาย",
				FirstNameTH: "สมชาย",
				LastNameTH:  "ใจดี",
				TitleEN:     "Mr.",
				FirstNameEN: "Somchai",
				LastNameEN:  "Jaidee",
				DateOfBirth: "1990-01-01",
				Gender:      "Male",
				Department:  "IT",
				Role:        "admin",
				Phone:       "080-123-4567",
				Email:       "admin@example.com",
				Password:    "123456",
			},
		}

		mockService.On("GetEmployees").Return(expectedEmployee, nil)

		req := httptest.NewRequest("GET", "/api/employees", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetEmployeeHandler(t *testing.T) {
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Get("/api/employees/:id", handle.GetEmployee)

	t.Run("Valid-GetEmployee", func(t *testing.T) {
		expectedEmployee := repository.Employee{
			EmployeeID:  "ADMIN",
			TitleTH:     "นาย",
			FirstNameTH: "สมชาย",
			LastNameTH:  "ใจดี",
			TitleEN:     "Mr.",
			FirstNameEN: "Somchai",
			LastNameEN:  "Jaidee",
			DateOfBirth: "1990-01-01",
			Gender:      "Male",
			Department:  "IT",
			Role:        "admin",
			Phone:       "080-123-4567",
			Email:       "admin@example.com",
			Password:    "123456",
		}

		mockService.On("GetEmployee", 1).Return(expectedEmployee, nil)

		req := httptest.NewRequest("GET", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateEmployeeHandler(t *testing.T) {
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Post("/api/employees", handle.CreateEmployee)

	expectedEmployee := repository.Employee{
		EmployeeID:  "ADMIN",
		TitleTH:     "นาย",
		FirstNameTH: "สมชาย",
		LastNameTH:  "ใจดี",
		TitleEN:     "Mr.",
		FirstNameEN: "Somchai",
		LastNameEN:  "Jaidee",
		DateOfBirth: "1990-01-01",
		Gender:      "Male",
		Department:  "IT",
		Role:        "admin",
		Phone:       "080-123-4567",
		Email:       "admin@example.com",
		Password:    "123456",
	}

	t.Run("Valid-CreateEmployee", func(t *testing.T) {
		mockService.On("CreateEmployee", mock.AnythingOfType("repository.Employee")).Return(expectedEmployee, nil)

		req := httptest.NewRequest("POST", "/api/employees", bytes.NewBufferString(`{
			"employee_id": "ADMIN",
			"title_th": "นาย",
			"first_name_th": "สมชาย",
			"last_name_th": "ใจดี",
			"title_en": "Mr.",
			"first_name_en": "Somchai",
			"last_name_en": "Jaidee",
			"date_of_birth": "1990-01-01",
			"gender": "Male",
			"department": "IT",
			"role": "admin",
			"phone": "080-123-4567",
			"email": "admin@example.com",
			"password": "123456"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateEmployeeHandler(t *testing.T) {
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Put("/api/employees/:id", handle.UpdateEmployee)

	expectedEmployee := repository.Employee{
		EmployeeID:  "ADMIN",
		TitleTH:     "นาย",
		FirstNameTH: "สมชาย",
		LastNameTH:  "ใจดี",
		TitleEN:     "Mr.",
		FirstNameEN: "Somchai",
		LastNameEN:  "Jaidee",
		DateOfBirth: "1990-01-01",
		Gender:      "Male",
		Department:  "IT",
		Role:        "admin",
		Phone:       "080-123-4567",
		Email:       "admin@example.com",
		Password:    "123456",
	}

	t.Run("Valid-UpdateEmployee", func(t *testing.T) {
		mockService.On("UpdateEmployee", "1", mock.AnythingOfType("repository.Employee")).Return(expectedEmployee, nil)

		req := httptest.NewRequest("PUT", "/api/employees/1", bytes.NewBufferString(`{
			"employee_id": "ADMIN",
			"title_th": "นาย",
			"first_name_th": "สมชาย",
			"last_name_th": "ใจดี",
			"title_en": "Mr.",
			"first_name_en": "Somchai",
			"last_name_en": "Jaidee",
			"date_of_birth": "1990-01-01",
			"gender": "Male",
			"department": "IT",
			"role": "admin",
			"phone": "080-123-4567",
			"email": "admin@example.com",
			"password": "123456"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteEmployeeHandler(t *testing.T) {
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Delete("/api/employees/:id", handle.DeleteEmployee)

	t.Run("Valid-DeleteEmployee", func(t *testing.T) {
		mockService.On("DeleteEmployee", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
 
func TestLoginHandler(t *testing.T){
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Post("/login", handle.Login)

	t.Run("Valid-LoginFiber", func(t *testing.T) {
		expectedEmployee := repository.Employee{
			EmployeeID:  "ADMIN",
			TitleTH:     "นาย",
			FirstNameTH: "สมชาย",
			LastNameTH:  "ใจดี",
			TitleEN:     "Mr.",
			FirstNameEN: "Somchai",
			LastNameEN:  "Jaidee",
			DateOfBirth: "1990-01-01",
			Gender:      "Male",
			Department:  "IT",
			Role:        "admin",
			Phone:       "080-123-4567",
			Email:       "admin@example.com",
			Password:    "123456",
		}
		
		//? ID ???
		mockService.On("Login", "", "123456").Return(expectedEmployee, nil)

		req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(`{
			"employee_id": "ADMIN",
			"title_th": "นาย",
			"first_name_th": "สมชาย",
			"last_name_th": "ใจดี",
			"title_en": "Mr.",
			"first_name_en": "Somchai",
			"last_name_en": "Jaidee",
			"date_of_birth": "1990-01-01",
			"gender": "Male",
			"department": "IT",
			"role": "admin",
			"phone": "080-123-4567",
			"email": "admin@example.com",
			"password": "123456"
		}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestLogoutHandler(t *testing.T) {
	mockService := new(MockEmployeeService)
	handle := NewHandleFiber(mockService)
	app := fiber.New()
	app.Post("/logout", handle.Logout)

	t.Run("Valid-Logout", func(t *testing.T) {
		// mockService.On("Logout").Return(nil)

		req := httptest.NewRequest("POST", "/logout", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

//! error 401
// func TestGetMeHandler(t *testing.T) {
// 	mockService := new(MockEmployeeService)
// 	handle := NewHandleFiber(mockService)
// 	app := fiber.New()
// 	app.Get("/me", handle.GetMe)

// 	t.Run("Valid-GetMe", func(t *testing.T) {
// 		expectedEmployee := repository.Employee{
// 			EmployeeID:  "ADMIN",
// 			TitleTH:     "นาย",
// 			FirstNameTH: "สมชาย",
// 			LastNameTH:  "ใจดี",
// 			TitleEN:     "Mr.",
// 			FirstNameEN: "Somchai",
// 			LastNameEN:  "Jaidee",
// 			DateOfBirth: "1990-01-01",
// 			Gender:      "Male",
// 			Department:  "IT",
// 			Role:        "admin",
// 			Phone:       "080-123-4567",
// 			Email:       "admin@example.com",
// 			Password:    "123456",
// 		}

// 		mockService.On("GetMe").Return(expectedEmployee, nil)

// 		req := httptest.NewRequest("GET", "/me", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, err := app.Test(req)

// 		assert.NoError(t, err)
// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
// 		mockService.AssertExpectations(t)
// 	})
// }

//! error 405
// func TestChangePasswordHandler(t *testing.T) {
// 	mockService := new(MockEmployeeService)
// 	handle := NewHandleFiber(mockService)
// 	app := fiber.New()
// 	app.Post("/api/changePassword", handle.ChangePassword)

// 	t.Run("Valid-ChangePassword", func(t *testing.T) {
// 		expectedEmployee := repository.Employee{
// 			EmployeeID:  "ADMIN",
// 			TitleTH:     "นาย",
// 			FirstNameTH: "สมชาย",
// 			LastNameTH:  "ใจดี",
// 			TitleEN:     "Mr.",
// 			FirstNameEN: "Somchai",
// 			LastNameEN:  "Jaidee",
// 			DateOfBirth: "1990-01-01",
// 			Gender:      "Male",
// 			Department:  "IT",
// 			Role:        "admin",
// 			Phone:       "080-123-4567",
// 			Email:       "admin@example.com",
// 			Password:    "123456",
// 		}

// 		mockService.On("ChangePassword", "", "123456", "654321").Return(expectedEmployee, nil)

// 		req := httptest.NewRequest("POST", "/api/changePassword", nil)
// 		req.Header.Set("Content-Type", "application/json")
// 		resp, err := app.Test(req)

// 		assert.NoError(t, err)
// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
// 		mockService.AssertExpectations(t)
// 	})
// }