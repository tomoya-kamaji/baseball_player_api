package entity

import "gorm.io/gorm"

type PlayerEntity struct {
	gorm.Model
	ID            uint
	UniformNumber int
	Name          string
	AtBats        int
	Hits          int
	Walks         int
	HomeRuns      int
	RunsBattedIn  int
	// TeamID        int
	// Positions     []*Position `gorm:"many2many:player_positions"`
}

func (h *PlayerEntity) ConvertToModel()  {
	// TODO: ConvertToModel
}

func (e *PlayerEntity) TableName() string {
	return "players"
}
