package domain

import "github.com/google/uuid"

type PlayerID string

func NewPlayerId() PlayerID {
	playerId := PlayerID(uuid.New().String())
	return playerId
}

type Player struct {
	ID            PlayerID
	UniformNumber int64  // 背番号
	Name          string // 名前
	AtBats        int64  // 打席数
	Hits          int64  // 安打数
	Walks         int64  // 四球数
	HomeRuns      int64  // ホームラン数
	RunsBattedIn  int64  // 打点数
}

func CreatePlayer(
	UniformNumber int64,
	Name string,
	AtBats int64,
	Hits int64,
	Walks int64,
	HomeRuns int64,
	RunsBattedIn int64,
) *Player {
	player := Player{
		ID:            NewPlayerId(),
		UniformNumber: UniformNumber,
		Name:          Name,
		AtBats:        AtBats,
		Hits:          Hits,
		Walks:         Walks,
		HomeRuns:      HomeRuns,
		RunsBattedIn:  RunsBattedIn,
	}
	return &player
}
