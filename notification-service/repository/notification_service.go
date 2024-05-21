package repository

type NotificationService interface {
	GetNotifications() ([]Notification, error)
	GetNotification(id int) (Notification, error)
	CreateNotification(Notification Notification) (Notification, error)
	UpdateNotification(id int, Notification Notification) (Notification, error)
	DeleteNotification(id int) error
	GetNotificationByEmployeeID(employeeID string) ([]Notification, error)
	ReadNotification(id int) error
}

type NotificationRepositoryDB struct {
	repo NotificationRepository
}

func NewNotificationService(repo NotificationRepository) NotificationService {
	return &NotificationRepositoryDB{
		repo: repo,
	}
}

func (n *NotificationRepositoryDB) GetNotifications() ([]Notification, error) {
	return n.repo.GetAll()
}

func (n *NotificationRepositoryDB) GetNotification(id int) (Notification, error) {
	return n.repo.GetByID(id)
}

func (n *NotificationRepositoryDB) CreateNotification(Notification Notification) (Notification, error) {
	return n.repo.Create(Notification)
}

func (n *NotificationRepositoryDB) UpdateNotification(id int, Notification Notification) (Notification, error) {
	return n.repo.Update(id, Notification)
}

func (n *NotificationRepositoryDB) DeleteNotification(id int) error {
	return n.repo.Delete(id)
}

func (n *NotificationRepositoryDB) GetNotificationByEmployeeID(employeeID string) ([]Notification, error) {
	return n.repo.GetByEmployeeID(employeeID)
}

func (n *NotificationRepositoryDB) ReadNotification(id int) error {
	return n.repo.Read(id)
}
