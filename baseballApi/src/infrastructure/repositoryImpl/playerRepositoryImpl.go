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

func (repo *playerRepositoryImpl) Create(c context.Context, player *domain.Player) error {
	playerEntity := entity.PlayerEntity{
		ID:            player.ID.String(),
		UniformNumber: player.UniformNumber,
		Name:          player.Name,
		AtBats:        player.AtBats,
		Hits:          player.Hits,
		Walks:         player.Walks,
		HomeRuns:      player.HomeRuns,
		RunsBattedIn:  player.RunsBattedIn,
	}

	return repo.db.Provide(c).Save(&playerEntity).Error
}

func (repo *playerRepositoryImpl) BulkUpsert(c context.Context, players []*domain.Player) error {
	var playerEntities []entity.PlayerEntity
	for _, player := range players {
		playerEntities = append(playerEntities, entity.PlayerEntity{
			ID:            player.ID.String(),
			UniformNumber: player.UniformNumber,
			Name:          player.Name,
			AtBats:        player.AtBats,
			Hits:          player.Hits,
			Walks:         player.Walks,
			HomeRuns:      player.HomeRuns,
			RunsBattedIn:  player.RunsBattedIn,
		})
	}
	return repo.db.Provide(c).Save(&playerEntities).Error
}

func (repo *playerRepositoryImpl) GetByID(c context.Context, ID domain.PlayerID) (*domain.Player, error) {
	var playerEntity entity.PlayerEntity
	if err := repo.db.Provide(c).First(&playerEntity, ID).Error; err != nil {
		return nil, err
	}
	return repo.ConvertToModel(c, playerEntity), nil
}

func (repo *playerRepositoryImpl) ConvertToModel(c context.Context, playerEntity entity.PlayerEntity) *domain.Player {
	return domain.ReConstractPlayer(
		domain.ReConstractPlayerParam{
			ID:            domain.PlayerID(playerEntity.ID),
			UniformNumber: playerEntity.UniformNumber,
			Name:          playerEntity.Name,
			AtBats:        playerEntity.AtBats,
			Hits:          playerEntity.Hits,
			Walks:         playerEntity.Walks,
			HomeRuns:      playerEntity.HomeRuns,
			RunsBattedIn:  playerEntity.RunsBattedIn,
		},
	)
}
