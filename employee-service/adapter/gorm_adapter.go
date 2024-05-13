package adapter

import (
	"github.com/night-sornram/employee-management/repository"
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

func (g *GormAdapter) GetByID(id int) (repository.Employee, error) {
	var Employee repository.Employee
	if err := g.db.First(&Employee, id).Error; err != nil {
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

func (g *GormAdapter) Update(id int, Employee repository.Employee) (repository.Employee, error) {
	if err := g.db.Model(&Employee).Where("id = ?", id).Updates(Employee).Error; err != nil {
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
