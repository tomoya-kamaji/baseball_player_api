package usecase

import (
	"context"

	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
)

type SearchPlayerUsecase struct {
	playerRepository domain.PlayerRepository
}

func NewSearchPlayerUsecase(playerRepository domain.PlayerRepository) *SearchPlayerUsecase {
	return &SearchPlayerUsecase{
		playerRepository: playerRepository,
	}
}

type playerDTO struct {
	ID            string
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}

type SearchPlayerUsecaseDTO struct {
	Players []playerDTO
}

func (u *SearchPlayerUsecase) Run(c context.Context, ID string) (*SearchPlayerUsecaseDTO, error) {
	// var player *domain.Player
	// player, err := u.playerRepository.GetByID(c, domain.PlayerID(ID))
	// if err != nil {
	// 	return nil, err
	// }

	var players []playerDTO
	return &SearchPlayerUsecaseDTO{
		Players: players,
	}, nil

}
