package repository

type EmployeeService interface {
	GetEmployees() ([]Employee, error)
	GetEmployee(id int) (Employee, error)
	CreateEmployee(Employee Employee) (Employee, error)
	UpdateEmployee(id int, Employee Employee) (Employee, error)
	DeleteEmployee(id int) error
}

type EmployeeServiceDB struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeService {
	return &EmployeeServiceDB{
		repo: repo,
	}
}

func (u *EmployeeServiceDB) GetEmployees() ([]Employee, error) {
	return u.repo.GetAll()
}

func (u *EmployeeServiceDB) GetEmployee(id int) (Employee, error) {
	return u.repo.GetByID(id)
}

func (u *EmployeeServiceDB) CreateEmployee(Employee Employee) (Employee, error) {
	return u.repo.Create(Employee)
}

func (u *EmployeeServiceDB) UpdateEmployee(id int, Employee Employee) (Employee, error) {
	return u.repo.Update(id, Employee)
}

func (u *EmployeeServiceDB) DeleteEmployee(id int) error {
	return u.repo.Delete(id)
}
