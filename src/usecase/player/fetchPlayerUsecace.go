package usecase

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/transactionImpl"
	"github.com/tomoya_kamaji/go-pkg/src/usecase/transaction"
)

type FetchPlayerUsecase struct {
	txMgr            transaction.TransactionManager
	playerRepository domain.PlayerRepository
}

func NewFetchPlayerUsecase(d *gorm.DBProvider) *FetchPlayerUsecase {
	return &FetchPlayerUsecase{
		txMgr:            transactionImpl.NewTransactionManagerImpl(d),
		playerRepository: repositoryImpl.NewPlayerRepositoryImpl(d),
	}
}

func (u *FetchPlayerUsecase) Run(c context.Context, ID string) (*domain.Player, error) {
	var player *domain.Player
	player, err := u.playerRepository.GetByID(c, domain.PlayerID(ID))
	if err != nil {
		return nil, err
	}
	return player, nil
}
