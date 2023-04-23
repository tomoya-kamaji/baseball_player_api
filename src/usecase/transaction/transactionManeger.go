package transaction

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
)

type TransactionManager interface {
	RunInTransaction(ctx context.Context, funcs ...func(ctx context.Context, tx *gorm.DBProvider) error) error
}
