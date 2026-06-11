package port

import "context"

type OutboxPublisher[T comparable] interface {
	Publish(ctx context.Context, event any) error
}
