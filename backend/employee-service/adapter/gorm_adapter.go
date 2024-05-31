package adapter

import (
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.EmployeeRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll() ([]repository.Employee, error) {
	var Employees []repository.Employee
	if err := g.db.Find(&Employees).Error; err != nil {
		return nil, err
	}
	return Employees, nil
}

func (g *GormAdapter) GetByID(eid string) (repository.Employee, error) {
	var Employee repository.Employee
	if err := g.db.Model(&Employee).Where("employee_id = ?", eid).First(&Employee).Error; err != nil {
		return Employee, err

	}
	return Employee, nil
}

func (g *GormAdapter) Create(Employee repository.Employee) (repository.Employee, error) {
	if err := g.db.Create(&Employee).Error; err != nil {
		return Employee, err
	}
	return Employee, nil
}

func (g *GormAdapter) Update(id string, Employee repository.Employee) (repository.Employee, error) {
	if err := g.db.Model(&Employee).Where("employee_id = ?", id).Updates(Employee).Error; err != nil {
		return Employee, err
	}
	return Employee, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) Login(id string, password string) (repository.Employee, error) {
	var Employee repository.Employee

	if err := g.db.First(&Employee, "employee_id = ?", id).Error; err != nil {
		return repository.Employee{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(Employee.Password), []byte(password)); err != nil {
		return repository.Employee{}, err
	}
	return Employee, nil
}

func (g *GormAdapter) GetMe(id string) (repository.Employee, error) {
	var Employee repository.Employee
	if err := g.db.Where("employee_id = ?", id).First(&Employee).Error; err != nil {
		return repository.Employee{}, err
	}

	return Employee, nil
}

func (g *GormAdapter) ChangePassword(id string, password string, new_password string) (repository.Employee, error) {

	var Employee repository.Employee
	if err := g.db.Where("employee_id = ?", id).First(&Employee).Error; err != nil {
		return repository.Employee{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(Employee.Password), []byte(password)); err != nil {
		return repository.Employee{}, err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(new_password), 14)
	Employee.Password = string(hashedPassword)

	if err := g.db.Save(&Employee).Error; err != nil {
		return repository.Employee{}, err
	}
	return Employee, nil
}
