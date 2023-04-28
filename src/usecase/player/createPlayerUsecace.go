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
type CreatePlayerUsecaseDto struct {
	ID            string
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}

func (u *CreatePlayerUsecase) Run(ctx context.Context, param CreatePlayerUsecaseParam) (*CreatePlayerUsecaseDto, error) {
	player := domain.CreatePlayer(
		param.UniformNumber,
		param.Name,
		param.AtBats,
		param.Hits,
		param.Walks,
		param.HomeRuns,
		param.RunsBattedIn,
	)
	err := u.txMgr.RunInTransaction(ctx, func(ctx context.Context) error {
		if err := u.playerRepository.Create(ctx, player); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	createPlayer, err := u.playerRepository.GetByID(ctx, player.ID)
	if err != nil {
		return nil, err
	}

	return &CreatePlayerUsecaseDto{
		ID:            createPlayer.ID.String(),
		UniformNumber: createPlayer.UniformNumber,
		Name:          createPlayer.Name,
		AtBats:        createPlayer.AtBats,
		Hits:          createPlayer.Hits,
		Walks:         createPlayer.Walks,
		HomeRuns:      createPlayer.HomeRuns,
		RunsBattedIn:  createPlayer.RunsBattedIn,
	}, nil
}
