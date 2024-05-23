package adapter

import (
	"employee/repository"

	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.NotificationRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll() ([]repository.Notification, error) {
	var notifications []repository.Notification
	if err := g.db.Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (g *GormAdapter) GetByID(id int) (repository.Notification, error) {
	var notification repository.Notification
	if err := g.db.First(&notification, id).Error; err != nil {
		return notification, err
	}
	return notification, nil
}

func (g *GormAdapter) Create(notification repository.Notification) (repository.Notification, error) {
	if err := g.db.Create(&notification).Error; err != nil {
		return notification, err
	}
	return notification, nil
}

func (g *GormAdapter) Update(id int, notification repository.Notification) (repository.Notification, error) {
	if err := g.db.Model(&notification).Where("id = ?", id).Updates(notification).Error; err != nil {
		return notification, err
	}
	return notification, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Notification{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) GetByEmployeeID(employeeID string) ([]repository.Notification, error) {
	var notifications []repository.Notification
	if err := g.db.Where("employee_id = ? AND read = false", employeeID).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (g *GormAdapter) Read(id int) error {
	if err := g.db.Model(&repository.Notification{}).Where("id = ?", id).Update("read", true).Error; err != nil {
		return err
	}
	return nil
}
