package domain

import "context"

// PlayerRepository is a repository interface for player
type PlayerRepository interface {
	Create(c context.Context, player *Player) (*Player, error)
	GetByID(c context.Context, ID PlayerID) (*Player, error)
}
