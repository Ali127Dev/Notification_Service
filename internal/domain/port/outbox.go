package port

import "context"

type OutboxPublisher interface {
	Publish(ctx context.Context, event any) error
}
