package adapter

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/employee-service/repository"
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

func (m *MockEmployeeService) GetEmployee(eid string) (repository.Employee, error) {
	args := m.Called(eid)
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
	t.Run("Valid-GetEmployees", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees", handle.GetEmployees)

		mockService.On("GetEmployees").Return([]repository.Employee{}, nil)

		req := httptest.NewRequest("GET", "/api/employees", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetEmployees", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees", handle.GetEmployees)

		//mock error
		mockService.On("GetEmployees").Return([]repository.Employee{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/employees", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetEmployeeHandler(t *testing.T) {
	t.Run("Valid-GetEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees/:id", handle.GetEmployee)

		mockService.On("GetEmployee", 1).Return(repository.Employee{}, nil)

		req := httptest.NewRequest("GET", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-GetEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees/:id", handle.GetEmployee)

		//ID "one" is invalid
		req := httptest.NewRequest("GET", "/api/employees/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-GetEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees/:id", handle.GetEmployee)

		mockService.On("GetEmployee", 1).Return(repository.Employee{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateEmployeeHandler(t *testing.T) {
	t.Run("Valid-CreateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/api/employees", handle.CreateEmployee)

		mockService.On("CreateEmployee", mock.AnythingOfType("repository.Employee")).Return(repository.Employee{}, nil)

		bodyEmployee := `{
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
		}`

		req := httptest.NewRequest("POST", "/api/employees", bytes.NewBufferString(bodyEmployee))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req, 2000)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-CreateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/api/employees", handle.CreateEmployee)

		//body is missing
		req := httptest.NewRequest("POST", "/api/employees", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-CreateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/api/employees", handle.CreateEmployee)

		mockService.On("CreateEmployee", mock.AnythingOfType("repository.Employee")).Return(repository.Employee{}, errors.New("invalid"))

		bodyEmployee := `{
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
		}`

		req := httptest.NewRequest("POST", "/api/employees", bytes.NewBufferString(bodyEmployee))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req, 2000)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateEmployeeHandler(t *testing.T) {
	t.Run("Valid-UpdateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Put("/api/employees/:id", handle.UpdateEmployee)

		mockService.On("UpdateEmployee", "1", mock.AnythingOfType("repository.Employee")).Return(repository.Employee{}, nil)

		bodyEmployee := `{
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
		}`

		req := httptest.NewRequest("PUT", "/api/employees/1", bytes.NewBufferString(bodyEmployee))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-UpdateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Put("/api/employees/:id", handle.UpdateEmployee)

		//body is missing
		req := httptest.NewRequest("PUT", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-UpdateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Put("/api/employees/:id", handle.UpdateEmployee)

		mockService.On("UpdateEmployee", "1", mock.AnythingOfType("repository.Employee")).Return(repository.Employee{}, errors.New("invalid"))

		bodyEmployee := `{
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
		}`

		req := httptest.NewRequest("PUT", "/api/employees/1", bytes.NewBufferString(bodyEmployee))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteEmployeeHandler(t *testing.T) {
	t.Run("Valid-DeleteEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Delete("/api/employees/:id", handle.DeleteEmployee)

		mockService.On("DeleteEmployee", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-ID-DeleteEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Delete("/api/employees/:id", handle.DeleteEmployee)

		//ID "one" is invalid
		req := httptest.NewRequest("DELETE", "/api/employees/one", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-DeleteEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Delete("/api/employees/:id", handle.DeleteEmployee)

		mockService.On("DeleteEmployee", 1).Return(errors.New("invalid"))

		req := httptest.NewRequest("DELETE", "/api/employees/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestLoginHandler(t *testing.T) {
	t.Run("Valid-LoginFiber", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/login", handle.Login)

		//? ID ???
		mockService.On("Login", "", "123456").Return(repository.Employee{}, nil)

		bodyAttendance := `{
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
		}`

		req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(bodyAttendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-LoginFiber", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/login", handle.Login)

		//body is missing
		req := httptest.NewRequest("POST", "/login", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-LoginFiber", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/login", handle.Login)

		//mock error
		mockService.On("Login", "", "123456").Return(repository.Employee{}, errors.New("invalid"))

		bodyAttendance := `{
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
		}`

		req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(bodyAttendance))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

}

func TestLogoutHandler(t *testing.T) {
	t.Run("Valid-Logout", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/logout", handle.Logout)

		req := httptest.NewRequest("POST", "/logout", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestGetMeHandler(t *testing.T) {
	t.Run("Valid-GetMe", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		mockService.On("GetMe", mock.Anything).Return(repository.Employee{}, nil)

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3MTQwNjg2LCJpc3MiOiJBRE1JTiIsInJvbGUiOiJhZG1pbiJ9.ViTG1GkFIQYknj2msORXASJsStL93oD-JT2y3oxS_Jw")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-Token-GetMe", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		//Header Authorization is missing
		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-Bearer-GetMe", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		//Bearer token is missing
		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3MTQwNjg2LCJpc3MiOiJBRE1JTiIsInJvbGUiOiJhZG1pbiJ9.ViTG1GkFIQYknj2msORXASJsStL93oD-JT2y3oxS_Jw")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("Invalid-JWT-GetMe", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		//Token is not JWT
		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer token")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("Invalid-GetMe", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		mockService.On("GetMe", mock.Anything).Return(repository.Employee{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3MTQwNjg2LCJpc3MiOiJBRE1JTiIsInJvbGUiOiJhZG1pbiJ9.ViTG1GkFIQYknj2msORXASJsStL93oD-JT2y3oxS_Jw")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestChangePasswordHandler(t *testing.T) {
	t.Run("Valid-ChangePassword", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/api/changePassword", handle.ChangePassword)

		mockService.On("ChangePassword", "1", "123456", "654321").Return(repository.Employee{}, nil)

		reqBody := `{
				"id": "1",
				"password": "123456",
				"new_password": "654321"
			}`

		req := httptest.NewRequest("POST", "/api/changePassword", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-BodyParser-ChangePassword", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/api/changePassword", handle.ChangePassword)

		//Body is missing
		req := httptest.NewRequest("POST", "/api/changePassword", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid-ChangePassword", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		handle := NewHandleFiber(mockService)
		app := fiber.New()
		app.Post("/api/changePassword", handle.ChangePassword)

		//mock error
		mockService.On("ChangePassword", "1", "123456", "654321").Return(repository.Employee{}, errors.New("invalid"))

		reqBody := `{
				"id": "1",
				"password": "123456",
				"new_password": "654321"
			}`

		req := httptest.NewRequest("POST", "/api/changePassword", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
