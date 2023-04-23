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

func (repo *playerRepositoryImpl) Create(c context.Context, player *domain.Player) (*domain.Player, error) {
	playerEntity := entity.PlayerEntity{
		UniformNumber: player.UniformNumber,
		Name:          player.Name,
		AtBats:        player.AtBats,
		Hits:          player.Hits,
		Walks:         player.Walks,
		HomeRuns:      player.HomeRuns,
		RunsBattedIn:  player.RunsBattedIn,
	}

	repo.db.Provide(c).Create(&playerEntity)
	return playerEntity.ConvertToModel(), nil
}

func (repo *playerRepositoryImpl) GetByID(c context.Context, ID domain.PlayerID) (*domain.Player, error) {
	var playerEntity entity.PlayerEntity
	repo.db.Provide(c).First(&playerEntity, ID)
	return playerEntity.ConvertToModel(), nil
}
