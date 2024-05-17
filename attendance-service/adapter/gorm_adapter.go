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

func (g *GormAdapter) CheckIn(checkIn repository.CheckIn) (repository.Attendance, error) {

	newAttendance := repository.Attendance{
		EmployeeID: checkIn.EmployeeID,
		CheckIn:    checkIn.CheckIn,
		CheckOut:   time.Time{},
		Date:       checkIn.CheckIn,
		LeaveID:    -1,
	}
	err := g.db.Create(&newAttendance).Error
	if err != nil {
		return newAttendance, err
	}

	return newAttendance, nil
}

func (g *GormAdapter) CheckOut(checkOut repository.CheckOut) (repository.Attendance, error) {

	var latestAttendance repository.Attendance
	newAttendance := repository.Attendance{
		CheckOut: checkOut.CheckOut,
	}
	err := g.db.Model(&latestAttendance).
    Where("employee_id = ?", checkOut.EmployeeID).
    Order("check_in desc").
	Limit(1).
	First(&latestAttendance).Error

	if err != nil {
		return newAttendance, err
	}

	err = g.db.Model(&latestAttendance).Updates(newAttendance).Error

	if err != nil {
		return newAttendance, err
	}

	return newAttendance, nil
}
