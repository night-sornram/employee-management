package repository

type AttendanceRepository interface {
	GetAll() ([]Attendance, error)
	GetByID(id int) (Attendance, error)
	Create(attendance Attendance) (Attendance, error)
	Update(id int, attendance Attendance) (Attendance, error)
	Delete(id int) error
	CheckIn(checkIn CheckIn) (Attendance, error)
	CheckOut(checkOut CheckOut) (Attendance, error)
}
