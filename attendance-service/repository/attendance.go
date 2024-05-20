package repository

import (
	"time"
)

type Attendance struct {
	ID         int       `db:"id" json:"id" gorm:"primaryKey"`
	EmployeeID string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckIn    time.Time `db:"check_in" json:"check_in" `
	CheckOut   time.Time `db:"check_out" json:"check_out" `
	Date       string    `db:"date" json:"date" validate:"required"`
	LeaveID    int       `db:"leave_id" json:"leave_id" `
}

type CheckIn struct {
	ID         int       `db:"id" json:"id" gor:"primaryKey"`
	EmployeeID string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckIn    time.Time `db:"check_in" json:"check_in"`
}
