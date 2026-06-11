package port

import (
	"context"

	"github.com/Ali127Dev/Notification_Service/internal/domain/entity"
)

type NotificationSender interface {
	Send(context.Context, *entity.Notification) error
}

type NotificationRepository interface {
	WithTx(tx Transaction) NotificationRepository

	Save(context.Context, *entity.Notification) error
	Update(context.Context, *entity.Notification) error

	FindByID(context.Context, string) (*entity.Notification, error)
}

type EventHandler func(ctx context.Context, event any) error

type NotificationConsumer interface {
	Consume(context.Context, EventHandler) error
}
