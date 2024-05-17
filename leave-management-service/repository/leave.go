package repository

import "time"

type Leave struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	DateStart  time.Time `db:"date_start" json:"date_start"`
	DateEnd    time.Time `db:"date_end" json:"date_end"`
	Reason     string    `db:"reason" json:"reason"`
	ManagerOpinion string    `db:"manager_opinion" json:"manager_opinion"`
	Status     string    `db:"status" json:"status"`
}

type LeaveStatus struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	ManagerOpinion string    `db:"manager_opinion" json:"manager_opinion"`
	Status     string    `db:"status" json:"status"`
}