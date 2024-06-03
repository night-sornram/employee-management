package repository

type Employee struct {
	ID          int    `db:"id" json:"id" gorm:"primaryKey"`
	EmployeeID  string `db:"employee_id" json:"employee_id" gorm:"unique"`
	TitleTH     string `db:"title_th" json:"title_th"`
	FirstNameTH string `db:"first_name_th" json:"first_name_th"`
	LastNameTH  string `db:"last_name_th" json:"last_name_th"`
	TitleEN     string `db:"title_en" json:"title_en"`
	FirstNameEN string `db:"first_name_en" json:"first_name_en"`
	LastNameEN  string `db:"last_name_en" json:"last_name_en"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Gender      string `db:"gender" json:"gender"`
	Department  string `db:"department" json:"department"`
	Role        string `db:"role" json:"role"`
	Phone       string `db:"phone" json:"phone"`
	Email       string `db:"email" json:"email" gorm:"unique"`
	Password    string `db:"password" json:"password"`
}
