package domain

import "github.com/tomoya_kamaji/go-pkg/src/infrastructure/entity"

// PlayerRepository is a repository interface for player
type PlayerRepository interface {
	GetByID(id string) (*entity.PlayerEntity, error)
}
