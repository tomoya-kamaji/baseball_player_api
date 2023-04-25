package usecase

import (
	"context"

	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/usecase/transaction"
)

type CreatePlayerUsecase struct {
	txMgr            transaction.TransactionManager
	playerRepository domain.PlayerRepository
}

func NewCreatePlayerUsecase(txMgr transaction.TransactionManager, playerRepository domain.PlayerRepository) *CreatePlayerUsecase {
	return &CreatePlayerUsecase{
		txMgr:            txMgr,
		playerRepository: playerRepository,
	}
}

type CreatePlayerUsecaseParam struct {
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
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
