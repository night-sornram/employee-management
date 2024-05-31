package repository

type NotificationRepository interface {
	GetAll() ([]Notification, error)
	GetByID(id int) (Notification, error)
	GetByEmployeeID(employeeID string) ([]Notification, error)
	Create(Notification Notification) (Notification, error)
	Update(id int, Notification Notification) (Notification, error)
	Delete(id int) error
	Read(id int) error
}
