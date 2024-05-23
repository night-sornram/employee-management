package repository

type LeaveRepository interface {
	GetAll() ([]Leave, error)
	GetByID(id int) (Leave, error)
	Create(leave Leave) (Leave, error)
	Update(id int, Leave Leave) (Leave, error)
	Delete(id int) error
	UpdateStatus(id int, leave Leave) (Leave, error)
	GetAllMe(eid string) ([]Leave, error)
}
