package port

import "github.com/Ali127Dev/Notification_Service/internal/domain/entity"

type NotificationSender interface {
	Send(*entity.Notification) error
}
