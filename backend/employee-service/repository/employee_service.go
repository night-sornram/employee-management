package repository

type EmployeeService interface {
	GetEmployees() ([]Employee, error)
	GetEmployee(eid string) (Employee, error)
	CreateEmployee(Employee Employee) (Employee, error)
	UpdateEmployee(id string, Employee Employee) (Employee, error)
	DeleteEmployee(id int) error
	Login(id string, password string) (Employee, error)
	Logout() error
	GetMe(id string) (Employee, error)
	ChangePassword(id string, password string, new_password string) (Employee, error)
	DownloadCSV(query string) ([]byte, error)
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

func (u *EmployeeServiceDB) GetEmployee(eid string) (Employee, error) {
	return u.repo.GetByID(eid)
}

func (u *EmployeeServiceDB) CreateEmployee(Employee Employee) (Employee, error) {
	return u.repo.Create(Employee)
}

func (u *EmployeeServiceDB) UpdateEmployee(id string, Employee Employee) (Employee, error) {
	return u.repo.Update(id, Employee)
}

func (u *EmployeeServiceDB) DeleteEmployee(id int) error {
	return u.repo.Delete(id)
}

func (u *EmployeeServiceDB) Login(id string, password string) (Employee, error) {
	return u.repo.Login(id, password)
}

func (u *EmployeeServiceDB) Logout() error {
	return nil
}

func (u *EmployeeServiceDB) GetMe(id string) (Employee, error) {
	return u.repo.GetMe(id)
}

func (u *EmployeeServiceDB) ChangePassword(id string, password string, new_password string) (Employee, error) {
	return u.repo.ChangePassword(id, password, new_password)
}

func (u *EmployeeServiceDB) DownloadCSV(query string) ([]byte, error) {
	return u.repo.GetCSV(query)
}
