package port

import "github.com/Ali127Dev/Notification_Service/internal/domain/entity"

type NotificationSender interface {
	Send(*entity.Notification) error
}

type NotificationRepository interface {
	Save(*entity.Notification) error
	Update(*entity.Notification) error

	FindByID(id string) (*entity.Notification, error)
}

type NotificationProducer interface {
	Publish(*entity.Notification) error
}
