package adapter

import (
	"github.com/night-sornram/employee-management/repository"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.LeaveRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll() ([]repository.Leave, error) {
	var leaves []repository.Leave
	query := `select * from leaves l join dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
	as employees(employee_id text, employee_name text, employee_lastname text) on l.employee_id = employees.employee_id;`
	if err := g.db.Raw(query).Scan(&leaves).Error; err != nil {
		return nil, err
	}
	return leaves, nil
}

func (g *GormAdapter) GetByID(id int) (repository.Leave, error) {
	var leave repository.Leave
	if err := g.db.First(&leave, id).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Create(leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Create(&leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Update(id int, leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Model(&leave).Where("id = ?", id).Updates(leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Leave{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) UpdateStatus(id int, leave repository.Leave) (repository.Leave, error) {
	var existingLeave repository.Leave
	if err := g.db.Where("id = ?", id).First(&existingLeave).Error; err != nil {
		return leave, err
	}

	existingLeave.Status = leave.Status
	existingLeave.ManagerOpinion = leave.ManagerOpinion

	if err := g.db.Save(&existingLeave).Error; err != nil {
		return existingLeave, err
	}
	return existingLeave, nil
}

func (g *GormAdapter) GetAllMe(eid string) ([]repository.Leave, error) {
	var Leaves []repository.Leave
	err := g.db.Where("employee_id = ?", eid).Find(&Leaves).Order("id DESC").Error
	if err != nil {
		return nil, err
	}
	return Leaves, nil
}
