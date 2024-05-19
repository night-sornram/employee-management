package repository

type AttendanceRepository interface {
	GetAll() ([]Attendance, error)
	GetByID(id int) (Attendance, error)
	Create(attendance Attendance) (Attendance, error)
	Update(id int, attendance Attendance) (Attendance, error)
	Delete(id int) error
	CheckIn(checkIn CheckIn) (Attendance, error)
	CheckOut(id int) (Attendance, error)
	GetAllMe(eid string) ([]Attendance, error)
	CheckToday(eid string) (Attendance, error)
}
