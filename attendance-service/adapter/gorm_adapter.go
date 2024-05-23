package adapter

import (
	"time"

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

func (g *GormAdapter) CheckIn(eid string) (repository.Attendance, error) {

	newAttendance := repository.Attendance{
		EmployeeID: eid,
		CheckIn:    time.Now(),
		CheckOut:   time.Time{},
		Date:       time.Now().Format("2006-01-02"),
		LeaveID:    -1,
	}
	err := g.db.Create(&newAttendance).Error
	if err != nil {
		return newAttendance, err
	}

	return newAttendance, nil
}

func (g *GormAdapter) CheckOut(id int) (repository.Attendance, error) {

	var attendance repository.Attendance
	err := g.db.First(&attendance, id).Error
	if err != nil {
		return attendance, err
	}
	attendance.CheckOut = time.Now()
	err = g.db.Save(&attendance).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil

}

func (g *GormAdapter) GetAllMe(eid string) ([]repository.Attendance, error) {
	var attendances []repository.Attendance
	err := g.db.Where("employee_id = ?", eid).Find(&attendances).Order("date DESC").Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (g *GormAdapter) CheckToday(eid string) (repository.Attendance, error) {
	var attendance repository.Attendance
	err := g.db.Where("employee_id = ? AND date = ?", eid, time.Now().Format(time.DateOnly)).First(&attendance).Error
	if err != nil {
		return attendance, nil
	}
	return attendance, nil
}
