package usecase

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/query"
)

type SearchPlayerUsecase struct {
	playerQuery query.PlayerSearchQuery
}

func NewSearchPlayerUsecase(playerQuery query.PlayerSearchQuery) *SearchPlayerUsecase {
	return &SearchPlayerUsecase{
		playerQuery: playerQuery,
	}
}

func (u *SearchPlayerUsecase) Run(c context.Context, param query.SearchParam) (*query.PlayerSearchQueryDTO, error) {
	result, err := u.playerQuery.Run(c, param)
	if err != nil {
		return nil, err
	}

	return result, nil
}
