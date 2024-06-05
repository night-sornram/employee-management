package test

import (
	"errors"
	"github.com/night-sornram/employee-management/employee-service/repository"
	"github.com/night-sornram/employee-management/employee-service/repository/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllEmployee(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("GetAll").Return([]repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.GetEmployees()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("GetAll").Return([]repository.Employee{}, errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.GetEmployees()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetByIDEmployee(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("GetByID", "EMP0001").Return(repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.GetEmployee("EMP0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("GetByID", "1").Return(repository.Employee{}, errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.GetEmployee("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreateEmployee(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.CreateEmployee(repository.Employee{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Create", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Employee{}, errors.New("invalid"))
		service := repository.NewEmployeeService(mockRepo)
		_, err := service.CreateEmployee(repository.Employee{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdateEmployee(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Update", "EMP0001", mock.Anything).Return(repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.UpdateEmployee("EMP0001", repository.Employee{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Update", "EMP0001", mock.Anything).Return(repository.Employee{}, errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.UpdateEmployee("EMP0001", repository.Employee{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDeleteEmployee(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Delete", 1).Return(nil)

		service := repository.NewEmployeeService(mockRepo)
		err := service.DeleteEmployee(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Delete", 1).Return(errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		err := service.DeleteEmployee(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestLoginEmployee(t *testing.T) {
	t.Run("Valid-Login", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Login", "1", "password").Return(repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.Login("1", "password")
		assert.NoError(t, err)
	})
	t.Run("Invalid-Login", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("Login", "1", "password").Return(repository.Employee{}, errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.Login("1", "password")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestLogoutEmployee(t *testing.T) {
	t.Run("Valid-Logout", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		service := repository.NewEmployeeService(mockRepo)
		err := service.Logout()
		assert.NoError(t, err)
	})
}

func TestGetMeEmployee(t *testing.T) {
	t.Run("Valid-GetMe", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("GetMe", "1").Return(repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.GetMe("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMe", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("GetMe", "1").Return(repository.Employee{}, errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.GetMe("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestChangePasswordEmployee(t *testing.T) {
	t.Run("Valid-ChangePassword", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("ChangePassword", "1", "password", "newPassword").Return(repository.Employee{}, nil)

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.ChangePassword("1", "password", "newPassword")
		assert.NoError(t, err)
	})
	t.Run("Invalid-ChangePassword", func(t *testing.T) {
		mockRepo := new(mocks.EmployeeRepository)
		mockRepo.On("ChangePassword", "1", "password", "newPassword").Return(repository.Employee{}, errors.New("invalid"))

		service := repository.NewEmployeeService(mockRepo)
		_, err := service.ChangePassword("1", "password", "newPassword")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
