package usecase

import (
	"context"

	"github.com/Ali127Dev/Notification_Service/internal/domain/entity"
	"github.com/Ali127Dev/Notification_Service/internal/domain/event"
	"github.com/Ali127Dev/Notification_Service/internal/domain/port"
)

type CreateNotification struct {
	repo            port.NotificationRepository
	idGen           port.IDGenerator
	outboxPublisher port.OutboxPublisher
}

func NewCreateNotification(
	repo port.NotificationRepository,
	idGen port.IDGenerator,
	outboxPublisher port.OutboxPublisher,
) *CreateNotification {
	return &CreateNotification{
		repo: repo, idGen: idGen,
		outboxPublisher: outboxPublisher,
	}
}

type CreateNotificationRequest struct {
	Kind    entity.NotificationKind
	To      string
	From    string
	Subject string
	Body    []byte
}

func (uc *CreateNotification) Execute(ctx context.Context, req CreateNotificationRequest) (*entity.Notification, error) {
	notification, err := entity.NewNotification(
		uc.idGen.NewID(),
		req.To,
		req.From,
		req.Subject,
		req.Body,
		req.Kind,
	)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Save(ctx, notification); err != nil {
		return nil, err
	}

	evt := event.NotificationCreated{
		ID: notification.ID(),
	}
	if err := uc.outboxPublisher.Publish(ctx, evt); err != nil {
		return nil, err
	}

	return notification, nil
}
