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
	var Leaves []repository.Leave
	if err := g.db.Find(&Leaves).Error; err != nil {
		return nil, err
	}
	return Leaves, nil
}

func (g *GormAdapter) GetByID(id int) (repository.Leave, error) {
	var Leave repository.Leave
	if err := g.db.First(&Leave, id).Error; err != nil {
		return Leave, err
	}
	return Leave, nil
}

func (g *GormAdapter) Create(Leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Create(&Leave).Error; err != nil {
		return Leave, err
	}
	return Leave, nil
}

func (g *GormAdapter) Update(id int, Leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Model(&Leave).Where("id = ?", id).Updates(Leave).Error; err != nil {
		return Leave, err
	}
	return Leave, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Leave{}, id).Error; err != nil {
		return err
	}
	return nil
}
