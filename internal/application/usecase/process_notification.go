package usecase

import (
	"context"

	"github.com/Ali127Dev/Notification_Service/internal/application/retry"
	"github.com/Ali127Dev/Notification_Service/internal/domain/port"
)

type ProcessNotification struct {
	repo          port.NotificationRepository
	sender        port.NotificationSender
	consumerRetry *retry.Runner
}

func NewProcessNotification(
	repo port.NotificationRepository,
	sender port.NotificationSender,
	consumerRetry *retry.Runner,
) *ProcessNotification {
	return &ProcessNotification{repo: repo, sender: sender, consumerRetry: consumerRetry}
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
	err = uc.consumerRetry.Do(func() error {
		return uc.sender.Send(ctx, notification)
	})
	if err != nil {
		notification.MarkAsFailed()

		if saveErr := uc.repo.Update(ctx, notification); saveErr != nil {
			return saveErr
		}

		return err
	}

	if err := notification.MarkAsSent(); err != nil {
		return err
	}

	return uc.repo.Update(ctx, notification)
}
