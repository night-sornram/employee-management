package repository

type EmployeeService interface {
	GetEmployees() ([]Employee, error)
	GetEmployee(id int) (Employee, error)
	CreateEmployee(Employee Employee) (Employee, error)
	UpdateEmployee(id int, Employee Employee) (Employee, error)
	DeleteEmployee(id int) error
	Login(email string, password string) (Employee, error)
	Logout() error
	GetMe(id string) (Employee, error)
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

func (u *EmployeeServiceDB) Login(email string, password string) (Employee, error) {
	return u.repo.Login(email, password)
}

func (u *EmployeeServiceDB) Logout() error {
	return nil
}

func (u *EmployeeServiceDB) GetMe(id string) (Employee, error) {
	return u.repo.GetMe(id)
}
