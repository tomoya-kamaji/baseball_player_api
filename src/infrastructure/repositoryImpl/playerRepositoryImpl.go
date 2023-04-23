package repositoryImpl

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/entity"
)

type playerRepositoryImpl struct {
	db *gorm.DBProvider
}

func NewPlayerRepositoryImpl(db *gorm.DBProvider) domain.PlayerRepository {
	return &playerRepositoryImpl{db: db}
}

func (repo *playerRepositoryImpl) GetByID(c context.Context, ID domain.PlayerID) (*domain.Player, error) {
	var playerEntity entity.PlayerEntity

	return playerEntity.ConvertToModel(), nil
}
