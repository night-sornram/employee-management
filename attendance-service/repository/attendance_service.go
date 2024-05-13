package repository

type AttendanceService interface {
	GetAttendances() ([]Attendance, error)
	GetAttendance(id int) (Attendance, error)
	CreateAttendance(attendance Attendance) (Attendance, error)
	UpdateAttendance(id int, attendance Attendance) (Attendance, error)
	DeleteAttendance(id int) error
}

type AttendanceServiceDB struct {
	repo AttendanceRepository
}

func NewAttendanceService(repo AttendanceRepository) AttendanceService {
	return &AttendanceServiceDB{
		repo: repo,
	}
}

func (a *AttendanceServiceDB) GetAttendances() ([]Attendance, error) {
	return a.repo.GetAll()
}

func (a *AttendanceServiceDB) GetAttendance(id int) (Attendance, error) {
	return a.repo.GetByID(id)
}

func (a *AttendanceServiceDB) CreateAttendance(attendance Attendance) (Attendance, error) {
	return a.repo.Create(attendance)
}

func (a *AttendanceServiceDB) UpdateAttendance(id int, attendance Attendance) (Attendance, error) {
	return a.repo.Update(id, attendance)
}

func (a *AttendanceServiceDB) DeleteAttendance(id int) error {
	return a.repo.Delete(id)
}
