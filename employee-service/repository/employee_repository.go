package repository

type EmployeeRepository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(Employee Employee) (Employee, error)
	Update(id int, Employee Employee) (Employee, error)
	Delete(id int) error
}
