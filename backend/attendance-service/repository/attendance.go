package repository

import (
	"time"
)

type Attendance struct {
	ID               int       `db:"id" json:"id" gorm:"primaryKey"`
	EmployeeID       string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckIn          time.Time `db:"check_in" json:"check_in" `
	CheckOut         time.Time `db:"check_out" json:"check_out" `
	Date             string    `db:"date" json:"date" validate:"required"`
	LeaveID          int       `db:"leave_id" json:"leave_id" `
	EmployeeName     string    `db:"employee_name" json:"employee_name"`
	EmployeeLastname string    `db:"employee_lastname" json:"employee_lastname"`
}

type CheckIn struct {
	EmployeeID string    `db:"employee_id" json:"employee_id" validate:"required"`
	CheckIn    time.Time `db:"check_in" json:"check_in"`
}

type Query struct {
	Date    string `json:"date"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Name    string `json:"name"`
	Option  string `json:"option"`
	LeaveID int    `json:"leave_id"`
}

type DataJson struct {
	Data     []Attendance `json:"data"`
	Total    int          `json:"total"`
	Page     int          `json:"page"`
	LastPage int          `json:"last_page"`
}
type GetMonth struct {
	Month int `db:"month" json:"month" validate:"required"`
	Year  int `db:"year" json:"year" validate:"required"`
}
