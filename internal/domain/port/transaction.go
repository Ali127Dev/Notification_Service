package port

import "context"

type Transaction interface {
	Commit() error
	Rollback() error
}

type TransactionManager interface {
	Begin(ctx context.Context) (Transaction, error)
}
