package entity

import (
	"errors"
	"time"
)

var (
	ErrNotificationIDRequired   = errors.New("notification id is required")
	ErrNotificationToRequired   = errors.New("notification to is required")
	ErrNotificationBodyRequired = errors.New("notification body is required")

	ErrNotificationKindInvalid = errors.New("notification kind is invalid")

	ErrNotificationFromRequiredForEmail    = errors.New("notification from is required for email")
	ErrNotificationSubjectRequiredForEmail = errors.New("notification subject is required for email")

	ErrNotificationSubjectRequiredForPush = errors.New("notification subject is required for push notification")

	ErrNotificationAlreadySent = errors.New("notification already sent")
)

type NotificationKind uint
type NotificationStatus uint

const (
	Email NotificationKind = iota
	SMS
	Push
)

const (
	Pending NotificationStatus = iota
	Sent
	Failed
)

var notificationKinds = [...]string{"Email", "SMS", "Push"}
var notificationStatuses = [...]string{"Pending", "Sent", "Failed"}

func (k NotificationKind) String() string {
	if k >= NotificationKind(len(notificationKinds)) {
		return "Unknown"
	}
	return notificationKinds[k]
}
func (k NotificationKind) Validate() error {
	if k >= NotificationKind(len(notificationKinds)) {
		return ErrNotificationKindInvalid
	}
	return nil
}

func (k NotificationStatus) String() string {
	if k >= NotificationStatus(len(notificationStatuses)) {
		return "Unknown"
	}
	return notificationStatuses[k]
}

type Notification struct {
	id        string
	kind      NotificationKind
	status    NotificationStatus
	attempts  uint
	body      []byte
	to        string
	from      string
	subject   string
	createdAt time.Time
}

func NewNotification(
	id, to, from, subject string,
	body []byte, kind NotificationKind,
) (*Notification, error) {
	if id == "" {
		return nil, ErrNotificationIDRequired
	}
	if to == "" {
		return nil, ErrNotificationToRequired
	}
	if len(body) <= 0 {
		return nil, ErrNotificationBodyRequired
	}
	if err := kind.Validate(); err != nil {
		return nil, err
	}
	switch kind {
	case Email:
		if from == "" {
			return nil, ErrNotificationFromRequiredForEmail
		}

		if subject == "" {
			return nil, ErrNotificationSubjectRequiredForEmail
		}

	case Push:
		if subject == "" {
			return nil, ErrNotificationSubjectRequiredForPush
		}
	}

	return &Notification{
		id: id, kind: kind, status: Pending, to: to,
		body: body, createdAt: time.Now().UTC(),
		from: from, subject: subject, attempts: 0,
	}, nil
}

func (n *Notification) ID() string             { return n.id }
func (n *Notification) Kind() NotificationKind { return n.kind }
func (n *Notification) Body() []byte {
	cp := make([]byte, len(n.body))
	copy(cp, n.body)
	return cp
}
func (n *Notification) To() string           { return n.to }
func (n *Notification) From() string         { return n.from }
func (n *Notification) Subject() string      { return n.subject }
func (n *Notification) Attempts() uint       { return n.attempts }
func (n *Notification) CreatedAt() time.Time { return n.createdAt }
func (n *Notification) IsSent() bool         { return n.status == Sent }
func (n *Notification) IsFailed() bool       { return n.status == Failed }
func (n *Notification) IsPending() bool      { return n.status == Pending }

func (n *Notification) MarkAsSent() error {
	if n.status == Sent {
		return ErrNotificationAlreadySent
	}
	n.status = Sent
	return nil
}
func (n *Notification) MarkAsFailed() error {
	if n.status == Sent {
		return ErrNotificationAlreadySent
	}
	n.status = Failed
	return nil
}
func (n *Notification) IncrementAttempts() {
	n.attempts++
}
