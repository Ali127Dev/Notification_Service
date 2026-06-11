package port

import (
	"context"
)

type Outbox interface {
	WithTx(tx Transaction) Outbox
	Insert(ctx context.Context, event any) error
}
