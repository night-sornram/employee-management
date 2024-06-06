package test

import (
	"bytes"
	"errors"
	"github.com/night-sornram/employee-management/employee-service/adapter"
	"github.com/night-sornram/employee-management/employee-service/repository/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/night-sornram/employee-management/employee-service/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEmployeesHandler(t *testing.T) {
	t.Run("Valid-GetEmployees", func(t *testing.T) {
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees/:id", handle.GetEmployee)

		mockService.On("GetEmployee", "EMP0001").Return(repository.Employee{}, nil)

		req := httptest.NewRequest("GET", "/api/employees/EMP0001", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-GetEmployee", func(t *testing.T) {
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/api/employees/:id", handle.GetEmployee)

		mockService.On("GetEmployee", "EMP0001").Return(repository.Employee{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/api/employees/EMP0001", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCreateEmployeeHandler(t *testing.T) {
	t.Run("Valid-CreateEmployee", func(t *testing.T) {
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		mockService.On("GetMe", mock.Anything).Return(repository.Employee{}, nil)

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3NTk5NDc5LCJpc3MiOiJBRE1JTiIsInJvbGUiOiJhZG1pbiJ9.e-7Fg8KuhYxgokKsSjQyjDqh_Lu720yt_YiKER7vMig")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
	t.Run("Invalid-Token-GetMe", func(t *testing.T) {
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		//Bearer token is missing
		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3NTk5NDc5LCJpc3MiOiJBRE1JTiIsInJvbGUiOiJhZG1pbiJ9.e-7Fg8KuhYxgokKsSjQyjDqh_Lu720yt_YiKER7vMig")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("Invalid-JWT-GetMe", func(t *testing.T) {
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
		app := fiber.New()
		app.Get("/me", handle.GetMe)

		mockService.On("GetMe", mock.Anything).Return(repository.Employee{}, errors.New("invalid"))

		req := httptest.NewRequest("GET", "/me", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3NTk5NDc5LCJpc3MiOiJBRE1JTiIsInJvbGUiOiJhZG1pbiJ9.e-7Fg8KuhYxgokKsSjQyjDqh_Lu720yt_YiKER7vMig")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestChangePasswordHandler(t *testing.T) {
	t.Run("Valid-ChangePassword", func(t *testing.T) {
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
		mockService := new(mocks.EmployeeService)
		handle := adapter.NewHandlerFiber(mockService)
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
