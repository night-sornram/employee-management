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

type Query struct {
	Date    string `json:"date"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Option  string `json:"option"`
}

type DataJson struct {
	Data     []Leave `json:"data"`
	Total    int     `json:"total"`
	Page     int     `json:"page"`
	LastPage int     `json:"last_page"`
}
