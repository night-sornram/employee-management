package adapter

import (
	"github.com/night-sornram/employee-management/repository"

	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.AttendanceRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll() ([]repository.Attendance, error) {
	var attendances []repository.Attendance
	err := g.db.Find(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (g *GormAdapter) GetByID(id int) (repository.Attendance, error) {
	var attendance repository.Attendance
	err := g.db.First(&attendance, id).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil
}

func (g *GormAdapter) Create(attendance repository.Attendance) (repository.Attendance, error) {
	err := g.db.Create(&attendance).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil
}

func (g *GormAdapter) Update(id int, attendance repository.Attendance) (repository.Attendance, error) {
	err := g.db.Model(&attendance).Where("id = ?", id).Updates(attendance).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil
}

func (g *GormAdapter) Delete(id int) error {
	err := g.db.Delete(&repository.Attendance{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
