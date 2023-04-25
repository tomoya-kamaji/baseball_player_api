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
func (u *FetchPlayerUsecase) Run(c context.Context, ID string) (*domain.Player, error) {
	var player *domain.Player
	player, err := u.playerRepository.GetByID(c, domain.PlayerID(ID))
	if err != nil {
		return nil, err
	}
	return player, nil
}
