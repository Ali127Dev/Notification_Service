package usecase

import (
	"context"

	"github.com/Ali127Dev/Notification_Service/internal/domain/port"
)

type ProcessNotification struct {
	repo   port.NotificationRepository
	sender port.NotificationSender
}

func NewProcessNotification(
	repo port.NotificationRepository,
	sender port.NotificationSender,
) *ProcessNotification {
	return &ProcessNotification{repo: repo, sender: sender}
}

type ProcessNotificationRequest struct {
	NotificationID string
}

func (uc *ProcessNotification) Execute(ctx context.Context, req ProcessNotificationRequest) error {
	notification, err := uc.repo.FindByID(ctx, req.NotificationID)
	if err != nil {
		return err
	}

	if !notification.IsPending() {
		return nil
	}

	notification.IncrementAttempts()
	if err := uc.repo.Update(ctx, notification); err != nil {
		return err
	}

	err = uc.sender.Send(ctx, notification)
	if err != nil {
		notification.MarkAsFailed()

		if updateErr := uc.repo.Update(ctx, notification); updateErr != nil {
			return updateErr
		}

		return err
	}

	if err := notification.MarkAsSent(); err != nil {
		return err
	}

	return uc.repo.Update(ctx, notification)
}
