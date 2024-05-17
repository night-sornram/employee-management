package repository

type LeaveService interface {
	GetLeaves() ([]Leave, error)
	GetLeave(id int) (Leave, error)
	CreateLeave(Leave Leave) (Leave, error)
	UpdateLeave(id int, Leave Leave) (Leave, error)
	DeleteLeave(id int) error
	GetMyLeaves(eid string) ([]Leave, error)
}

type LeaveServiceDB struct {
	repo LeaveRepository
}

func NewLeaveService(repo LeaveRepository) LeaveService {
	return &LeaveServiceDB{
		repo: repo,
	}
}

func (u *LeaveServiceDB) GetLeaves() ([]Leave, error) {
	return u.repo.GetAll()
}

func (u *LeaveServiceDB) GetLeave(id int) (Leave, error) {
	return u.repo.GetByID(id)
}

func (u *LeaveServiceDB) CreateLeave(Leave Leave) (Leave, error) {
	return u.repo.Create(Leave)
}

func (u *LeaveServiceDB) UpdateLeave(id int, Leave Leave) (Leave, error) {
	return u.repo.Update(id, Leave)
}

func (u *LeaveServiceDB) DeleteLeave(id int) error {
	return u.repo.Delete(id)
}

func (u *LeaveServiceDB) GetMyLeaves(eid string) ([]Leave, error) {
	return u.repo.GetAllMe(eid)
}
