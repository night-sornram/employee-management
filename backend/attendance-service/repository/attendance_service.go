package repository

type AttendanceService interface {
	GetAttendances() ([]Attendance, error)
	GetAttendance(id int) (Attendance, error)
	CreateAttendance(attendance Attendance) (Attendance, error)
	UpdateAttendance(id int, attendance Attendance) (Attendance, error)
	DeleteAttendance(id int) error
	CheckIn(eid string) (Attendance, error)
	CheckOut(id int) (Attendance, error)
	GetMyAttendances(eid string) ([]Attendance, error)
	CheckToday(eid string) (Attendance, error)
	GetDayLate() ([]Attendance, error)
	GetMonthLate(date GetMonth) ([]Attendance, error)
	GetYearLate(year int) ([]Attendance, error)
	GetAllLate() ([]Attendance, error)
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

func (a *AttendanceServiceDB) CheckIn(eid string) (Attendance, error) {
	return a.repo.CheckIn(eid)
}

func (a *AttendanceServiceDB) CheckOut(id int) (Attendance, error) {
	return a.repo.CheckOut(id)
}

func (a *AttendanceServiceDB) GetMyAttendances(eid string) ([]Attendance, error) {
	return a.repo.GetAllMe(eid)
}

func (a *AttendanceServiceDB) CheckToday(eid string) (Attendance, error) {
	return a.repo.CheckToday(eid)
}

func (a *AttendanceServiceDB) GetDayLate() ([]Attendance, error) {
	return a.repo.GetDayLate()
}

func (a *AttendanceServiceDB) GetMonthLate(date GetMonth) ([]Attendance, error) {
	return a.repo.GetMonthLate(date.Month, date.Year)
}

func (a *AttendanceServiceDB) GetYearLate(year int) ([]Attendance, error) {
	return a.repo.GetYearLate(year)
}

func (a *AttendanceServiceDB) GetAllLate() ([]Attendance, error) {
	return a.repo.GetAllLate()
}
