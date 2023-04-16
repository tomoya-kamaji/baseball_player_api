package entity

import (
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"gorm.io/gorm"
)

type PlayerEntity struct {
	gorm.Model
	ID            string
	UniformNumber int
	Name          string
	AtBats        int
	Hits          int
	Walks         int
	HomeRuns      int
	RunsBattedIn  int
	// TODO: Add more fields
	// TeamID        int
	// Positions     []*Position `gorm:"many2many:player_positions"`
}

func (h *PlayerEntity) ConvertToModel() *domain.Player {
	return &domain.Player{
		ID:            domain.PlayerID(h.ID),
		UniformNumber: h.UniformNumber,
		Name:          h.Name,
		AtBats:        h.AtBats,
		Hits:          h.Hits,
		Walks:         h.Walks,
		HomeRuns:      h.HomeRuns,
		RunsBattedIn:  h.RunsBattedIn,
	}
}

func (e *PlayerEntity) TableName() string {
	return "players"
}
