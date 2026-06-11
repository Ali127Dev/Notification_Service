package adapter

import (
	"context"

	"github.com/Ali127Dev/xerr"
	"github.com/Ali127Dev/xoutbox"
)

type XOutboxAdapter[T comparable] struct {
	store xoutbox.Store[T]
}

func (o *XOutboxAdapter[T]) InsertEvent(ctx context.Context, event any) error {
	e, ok := event.(xoutbox.Event[T])
	if !ok {
		return xerr.New(xerr.CodeInternalError,
			xerr.WithMessage("invalid event type"))
	}

	return o.store.InsertEvent(ctx, e)
}
