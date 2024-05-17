package repository

import (
	"time"
)

type Attendance struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckIn    time.Time `db:"check_in" json:"check_in" validate:"required"`
	CheckOut   time.Time `db:"check_out" json:"check_out" validate:"required"`
	Date       time.Time `db:"date" json:"date" validate:"required"`
	LeaveID    int       `db:"leave_id" json:"leave_id" validate:"required"`
}

type CheckIn struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckIn    time.Time `db:"check_in" json:"check_in" validate:"required"`
}

type CheckOut struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckOut   time.Time `db:"check_out" json:"check_out" validate:"required"`
}
