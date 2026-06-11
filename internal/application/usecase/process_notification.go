package usecase

import (
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

func (uc *ProcessNotification) Execute(req ProcessNotificationRequest) error {
	notification, err := uc.repo.FindByID(req.NotificationID)
	if err != nil {
		return err
	}

	if !notification.IsPending() {
		return nil
	}

	notification.IncrementAttempts()
	err = uc.consumerRetry.Do(func() error {
		return uc.sender.Send(notification)
	})
	if err != nil {
		notification.MarkAsFailed()

		if saveErr := uc.repo.Update(notification); saveErr != nil {
			return saveErr
		}

		return err
	}

	if err := notification.MarkAsSent(); err != nil {
		return err
	}

	return uc.repo.Update(notification)
}
