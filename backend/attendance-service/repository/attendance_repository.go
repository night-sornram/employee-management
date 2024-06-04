package repository

type AttendanceRepository interface {
	GetAll(query Query) (DataJson, error)
	GetByID(id int) (Attendance, error)
	Create(attendance Attendance) (Attendance, error)
	Update(id int, attendance Attendance) (Attendance, error)
	Delete(id int) error
	CheckIn(eid string) (Attendance, error)
	CheckOut(id int) (Attendance, error)
	GetAllMe(query Query, eid string) (DataJson, error)
	CheckToday(eid string) (Attendance, error)
}
