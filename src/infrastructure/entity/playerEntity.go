package entity

import (
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
)

type PlayerEntity struct {
	BaseModel
	ID            string
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
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
