package outbox

import (
	"context"
	"time"

	"github.com/Ali127Dev/xoutbox"
)

type Worker[T comparable] struct {
	store     xoutbox.Store[T]
	publisher xoutbox.Publisher[T]

	limit    int
	interval time.Duration
}

func NewWorker[T comparable](
	store xoutbox.Store[T],
	publisher xoutbox.Publisher[T],
	limit int,
	interval time.Duration,
) *Worker[T] {
	return &Worker[T]{
		store:     store,
		publisher: publisher,
		limit:     limit,
		interval:  interval,
	}
}

func (w *Worker[T]) Run(ctx context.Context) error {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-ticker.C:
			w.process(ctx)
		}
	}
}

func (w *Worker[T]) process(ctx context.Context) {
	events, err := w.store.FetchPending(ctx, w.limit)
	if err != nil {
		return
	}

	for _, evt := range events {
		w.handle(ctx, evt)
	}
}

func (w *Worker[T]) handle(ctx context.Context, evt xoutbox.Event[T]) {
	err := w.publisher.Publish(ctx, evt)
	if err != nil {
		_ = w.store.MarkFailed(ctx, evt.ID, evt.RetryCount+1)
		return
	}

	_ = w.store.MarkPublished(ctx, evt.ID)
}
