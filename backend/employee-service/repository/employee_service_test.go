package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEmployeeRepo struct {
	GetAllFunc         func() ([]Employee, error)
	GetByIDFunc        func(eid string) (Employee, error)
	CreateFunc         func(employee Employee) (Employee, error)
	UpdateFunc         func(id string, Employee Employee) (Employee, error)
	DeleteFunc         func(id int) error
	LoginFunc          func(id string, password string) (Employee, error)
	GetMeFunc          func(id string) (Employee, error)
	ChangePasswordFunc func(id string, password string, new_password string) (Employee, error)
}

func (m *mockEmployeeRepo) GetAll() ([]Employee, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *mockEmployeeRepo) GetByID(eid string) (Employee, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(eid)
	}
	return Employee{}, errors.New("not implemented")
}

func (m *mockEmployeeRepo) Create(employee Employee) (Employee, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(employee)
	}
	return Employee{}, errors.New("not implemented")
}

func (m *mockEmployeeRepo) Update(id string, employee Employee) (Employee, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, employee)
	}
	return Employee{}, errors.New("not implemented")
}

func (m *mockEmployeeRepo) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockEmployeeRepo) Login(id string, password string) (Employee, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(id, password)
	}
	return Employee{}, errors.New("not implemented")
}

func (m *mockEmployeeRepo) GetMe(id string) (Employee, error) {
	if m.GetMeFunc != nil {
		return m.GetMeFunc(id)
	}
	return Employee{}, errors.New("not implemented")
}

func (m *mockEmployeeRepo) ChangePassword(id string, password string, new_password string) (Employee, error) {
	if m.ChangePasswordFunc != nil {
		return m.ChangePasswordFunc(id, password, new_password)
	}
	return Employee{}, errors.New("not implemented")
}

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			GetAllFunc: func() ([]Employee, error) {
				return []Employee{}, nil
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.GetEmployees()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			GetAllFunc: func() ([]Employee, error) {
				return []Employee{}, errors.New("invalid")
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.GetEmployees()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			GetByIDFunc: func(eid string) (Employee, error) {
				return Employee{}, nil
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.GetEmployee("E0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			GetByIDFunc: func(eid string) (Employee, error) {
				return Employee{}, errors.New("invalid")
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.GetEmployee("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreate(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			CreateFunc: func(employee Employee) (Employee, error) {
				return Employee{}, nil
			},
		}
		service := NewEmployeeService(mockRepo)
		_, err := service.CreateEmployee(Employee{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Create", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			CreateFunc: func(employee Employee) (Employee, error) {
				return Employee{}, errors.New("invalid")
			},
		}
		service := NewEmployeeService(mockRepo)
		_, err := service.CreateEmployee(Employee{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			UpdateFunc: func(id string, employee Employee) (Employee, error) {
				return Employee{}, nil
			},
		}
		service := NewEmployeeService(mockRepo)
		_, err := service.UpdateEmployee("E0001", Employee{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			UpdateFunc: func(id string, employee Employee) (Employee, error) {
				return Employee{}, errors.New("invalid")
			},
		}
		service := NewEmployeeService(mockRepo)
		_, err := service.UpdateEmployee("E0001", Employee{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		}
		service := NewEmployeeService(mockRepo)
		err := service.DeleteEmployee(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			DeleteFunc: func(id int) error {
				return errors.New("invalid")
			},
		}
		service := NewEmployeeService(mockRepo)
		err := service.DeleteEmployee(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestLogin(t *testing.T) {
	t.Run("Valid-Login", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			LoginFunc: func(id string, password string) (Employee, error) {
				return Employee{}, nil
			},
		}
		service := NewEmployeeService(mockRepo)
		_, err := service.Login("1", "password")
		assert.NoError(t, err)
	})
	t.Run("Invalid-Login", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			LoginFunc: func(id string, password string) (Employee, error) {
				return Employee{}, errors.New("invalid")
			},
		}
		service := NewEmployeeService(mockRepo)
		_, err := service.Login("1", "password")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestLogout(t *testing.T) {
	t.Run("Valid-Logout", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{}
		service := NewEmployeeService(mockRepo)
		err := service.Logout()
		assert.NoError(t, err)
	})
}

func TestGetMe(t *testing.T) {
	t.Run("Valid-GetMe", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			GetMeFunc: func(id string) (Employee, error) {
				return Employee{}, nil
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.GetMe("1")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMe", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			GetMeFunc: func(id string) (Employee, error) {
				return Employee{}, errors.New("invalid")
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.GetMe("1")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("Valid-ChangePassword", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			ChangePasswordFunc: func(id string, password string, new_password string) (Employee, error) {
				return Employee{}, nil
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.ChangePassword("1", "password", "newPassword")
		assert.NoError(t, err)
	})
	t.Run("Invalid-ChangePassword", func(t *testing.T) {
		mockRepo := &mockEmployeeRepo{
			ChangePasswordFunc: func(id string, password string, new_password string) (Employee, error) {
				return Employee{}, errors.New("invalid")
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.ChangePassword("1", "password", "newPassword")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
