package repository

import "time"

type Attendance struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	CheckIn    time.Time `db:"check_in" json:"check_in"`
	CheckOut   time.Time `db:"check_out" json:"check_out"`
	Date       time.Time `db:"date" json:"date"`
	LeaveID    int       `db:"leave_id" json:"leave_id"`
}
