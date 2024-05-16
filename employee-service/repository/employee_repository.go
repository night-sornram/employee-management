package repository

type EmployeeRepository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(Employee Employee) (Employee, error)
	Update(id string, Employee Employee) (Employee, error)
	Delete(id int) error
	Login(email string, password string) (Employee, error)
	GetMe(id string) (Employee, error)
	ChangePassword(id string, password string, new_password string) (Employee, error)
}
