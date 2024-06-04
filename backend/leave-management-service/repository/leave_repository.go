package repository

type LeaveRepository interface {
	GetAll(query Query) (DataJson, error)
	GetByID(id int) (Leave, error)
	Create(leave Leave) (Leave, error)
	Update(id int, Leave Leave) (Leave, error)
	Delete(id int) error
	UpdateStatus(id int, leave Leave) (Leave, error)
	GetAllMe(query Query, eid string) (DataJson, error)
}
