package usecase

import (
	"context"

	"github.com/Ali127Dev/Notification_Service/internal/domain/entity"
	"github.com/Ali127Dev/Notification_Service/internal/domain/event"
	"github.com/Ali127Dev/Notification_Service/internal/domain/port"
)

type CreateNotification struct {
	repo   port.NotificationRepository
	idGen  port.IDGenerator
	outbox port.Outbox
	txMgr  port.TransactionManager
}

func NewCreateNotification(
	repo port.NotificationRepository,
	idGen port.IDGenerator,
	outbox port.Outbox,
	txMgr port.TransactionManager,
) *CreateNotification {
	return &CreateNotification{
		repo: repo, idGen: idGen,
		outbox: outbox, txMgr: txMgr,
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
	tx, err := uc.txMgr.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	repo := uc.repo.WithTx(tx)
	outbox := uc.outbox.WithTx(tx)

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

	if err := repo.Save(ctx, notification); err != nil {
		return nil, err
	}

	evt := event.NotificationCreated{
		ID: notification.ID(),
	}
	if err := outbox.Insert(ctx, evt); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return notification, nil
}
