package domain

import "context"

// PlayerRepository is a repository interface for player
type PlayerRepository interface {
	GetByID(c context.Context, ID PlayerID) (*Player, error)
}
