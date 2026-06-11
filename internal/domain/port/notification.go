package port

import (
	"context"

	"github.com/Ali127Dev/Notification_Service/internal/domain/entity"
	"github.com/Ali127Dev/Notification_Service/internal/domain/event"
)

type NotificationSender interface {
	Send(context.Context, *entity.Notification) error
}

type NotificationRepository interface {
	Save(context.Context, *entity.Notification) error
	Update(context.Context, *entity.Notification) error

	FindByID(context.Context, string) (*entity.Notification, error)
}

type NotificationProducer interface {
	Publish(context.Context, event.NotificationCreated) error
}

type NotificationConsumer interface {
	Consume(context.Context, func(event.NotificationCreated) error) error
}
