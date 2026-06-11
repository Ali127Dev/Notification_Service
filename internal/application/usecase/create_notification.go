package usecase

import (
	"github.com/Ali127Dev/Notification_Service/internal/domain/entity"
	"github.com/Ali127Dev/Notification_Service/internal/domain/event"
	"github.com/Ali127Dev/Notification_Service/internal/domain/port"
)

type CreateNotification struct {
	repo     port.NotificationRepository
	producer port.NotificationProducer
	idGen    port.IDGenerator
}

func NewCreateNotification(
	repo port.NotificationRepository,
	producer port.NotificationProducer,
	idGen port.IDGenerator,
) *CreateNotification {
	return &CreateNotification{repo: repo, producer: producer, idGen: idGen}
}

type CreateNotificationRequest struct {
	Kind    entity.NotificationKind
	To      string
	From    string
	Subject string
	Body    []byte
}

func (uc *CreateNotification) Execute(req CreateNotificationRequest) (*entity.Notification, error) {
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

	if err := uc.repo.Save(notification); err != nil {
		return nil, err
	}

	event := event.NotificationCreated{
		ID: notification.ID(),
	}

	if err := uc.producer.Publish(event); err != nil {
		return nil, err
	}

	return notification, nil
}
