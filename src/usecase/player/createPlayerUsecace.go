package usecase

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/transactionImpl"
	"github.com/tomoya_kamaji/go-pkg/src/usecase/transaction"
)

type CreatePlayerUsecase struct {
	txMgr            transaction.TransactionManager
	playerRepository domain.PlayerRepository
}

func NewCreatePlayerUsecase(d *gorm.DBProvider) *CreatePlayerUsecase {
	return &CreatePlayerUsecase{
		txMgr:            transactionImpl.NewTransactionManagerImpl(d),
		playerRepository: repositoryImpl.NewPlayerRepositoryImpl(d),
	}
}

type CreatePlayerUsecaseParam struct {
	UniformNumber int
	Name          string
	AtBats        int
	Hits          int
	Walks         int
	HomeRuns      int
	RunsBattedIn  int
}

func (u *CreatePlayerUsecase) Run(c context.Context, param CreatePlayerUsecaseParam) (*domain.Player, error) {
	player := domain.CreatePlayer(
		param.UniformNumber,
		param.Name,
		param.AtBats,
		param.Hits,
		param.Walks,
		param.HomeRuns,
		param.RunsBattedIn,
	)
	err := u.txMgr.RunInTransaction(c, func(ctx context.Context) error {
		_, err := u.playerRepository.Create(ctx, player)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return player, nil
}
