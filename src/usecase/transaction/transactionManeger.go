package transaction

import (
	"context"
)

type TransactionManager interface {
	RunInTransaction(ctx context.Context, funcs ...func(ctx context.Context) error) error
}
