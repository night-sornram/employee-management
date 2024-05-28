package repository

import "time"

type Leave struct {
	ID               int       `db:"id" json:"id" validate:"gte=0"`
	EmployeeID       string    `db:"employee_id" json:"employee_id" validate:"required"`
	DateStart        time.Time `db:"date_start" json:"date_start" validate:"required"`
	DateEnd          time.Time `db:"date_end" json:"date_end" validate:"required"`
	Reason           string    `db:"reason" json:"reason"`
	Category         string    `db:"category" json:"category" validate:"required"`
	ManagerOpinion   string    `db:"manager_opinion" json:"manager_opinion"`
	Status           string    `db:"status" json:"status" validate:"required"`
	Manager          string    `db:"manager" json:"manager"`
	EmployeeName     string    `db:"employee_name" json:"employee_name"`
	EmployeeLastname string    `db:"employee_lastname" json:"employee_lastname"`
}

type LeaveStatus struct {
	ID             int    `db:"id" json:"id"`
	EmployeeID     string `db:"employee_id" json:"employee_id" validate:"required"`
	ManagerOpinion string `db:"manager_opinion" json:"manager_opinion"`
	Status         string `db:"status" json:"status" validate:"required"`
}
