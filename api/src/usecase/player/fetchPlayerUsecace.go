package usecase

import (
	"context"

	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
)

type FetchPlayerUsecase struct {
	playerRepository domain.PlayerRepository
}

func NewFetchPlayerUsecase(playerRepository domain.PlayerRepository) *FetchPlayerUsecase {
	return &FetchPlayerUsecase{
		playerRepository: playerRepository,
	}
}

type FetchPlayerUsecaseDTO struct {
	ID            string
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}

func (u *FetchPlayerUsecase) Run(c context.Context, ID string) (*FetchPlayerUsecaseDTO, error) {
	var player *domain.Player
	player, err := u.playerRepository.GetByID(c, domain.PlayerID(ID))
	if err != nil {
		return nil, err
	}
	return &FetchPlayerUsecaseDTO{
		ID:            player.ID.String(),
		UniformNumber: player.UniformNumber,
		Name:          player.Name,
		AtBats:        player.AtBats,
		Hits:          player.Hits,
		Walks:         player.Walks,
		HomeRuns:      player.HomeRuns,
		RunsBattedIn:  player.RunsBattedIn,
	}, nil
}
