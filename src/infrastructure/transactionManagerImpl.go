package repositoryImpl

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/usecase"
)

type TransactionManagerImpl struct{}

func NewTransactionManagerImpl() usecase.TransactionManager {
	return &TransactionManagerImpl{}
}

func (t *TransactionManagerImpl) RunInTransaction(ctx context.Context, db *gorm.DBProvider, funcs ...func(ctx context.Context, tx *gorm.DBProvider) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else if ctx.Err() != nil {
			tx.Rollback()
		} else if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	for _, f := range funcs {
		if err := f(ctx, db); err != nil {
			return err
		}
	}

	return nil
}
