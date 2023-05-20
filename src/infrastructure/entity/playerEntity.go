package entity

type PlayerEntity struct {
	BaseModel
	ID            string `gorm:"primary_key"`
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}

func (e *PlayerEntity) TableName() string {
	return "players"
}
