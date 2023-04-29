package repositoryImpl

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/usecase/transaction"
)

type TransactionManagerImpl struct {
	db *gorm.DBProvider
}

func NewTransactionManagerImpl(db *gorm.DBProvider) transaction.TransactionManager {
	return &TransactionManagerImpl{db: db}
}

func (t *TransactionManagerImpl) RunInTransaction(ctx context.Context, funcs ...func(ctx context.Context) error) error {
	tx := t.db.Begin()
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
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}
