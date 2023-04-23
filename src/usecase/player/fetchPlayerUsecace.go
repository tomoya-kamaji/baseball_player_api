package usecase

import (
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
)

type FetchPlayerUsecase struct {
	playerRepository domain.PlayerRepository
}

// NewFetchPlayerUsecase func
func NewFetchPlayerUsecase(d *gorm.DBProvider) *FetchPlayerUsecase {
	return &FetchPlayerUsecase{
		playerRepository: repositoryImpl.NewPlayerRepositoryImpl(d),
	}
}
