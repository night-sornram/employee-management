package repository

type LeaveRepository interface {
	GetAll() ([]Leave, error)
	GetByID(id int) (Leave, error)
	Create(Leave Leave) (Leave, error)
	Update(id int, Leave Leave) (Leave, error)
	Delete(id int) error
	GetAllMe(eid string) ([]Leave, error)
}
