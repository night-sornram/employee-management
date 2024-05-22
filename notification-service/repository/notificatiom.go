package repository

type Notification struct {
	ID         int    `db:"id" json:"id" gorm:"primaryKey"`
	EmployeeID string `db:"employee_id" json:"employee_id" validate:"required"`
	Message    string `db:"message" json:"message" validate:"required"`
	Title      string `db:"title" json:"title" validate:"required"`
	Read       bool   `db:"read" json:"read" validate:"required"`
}
