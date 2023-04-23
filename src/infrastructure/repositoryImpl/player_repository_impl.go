package repositoryImpl

import (
	"context"

	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/entity"
	"gorm.io/gorm"
)

type PlayerRepositoryImpl struct {
	DB *gorm.DB
}

func NewPlayerRepositoryImpl(db *gorm.DB) domain.PlayerRepository {
	return &PlayerRepositoryImpl{DB: db}
}

func (pr *PlayerRepositoryImpl) GetByID(c context.Context, ID domain.PlayerID) (*domain.Player, error) {
	var playerEntity entity.PlayerEntity
	if err := pr.DB.First(&playerEntity, ID).Error; err != nil {
		return nil, err
	}
	return playerEntity.ConvertToModel(), nil
}
