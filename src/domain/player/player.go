package domain

import "github.com/google/uuid"

type PlayerID string

func NewPlayerId() PlayerID {
	playerId := PlayerID(uuid.New().String())
	return playerId
}

type Player struct {
	ID            PlayerID
	UniformNumber int    // 背番号
	Name          string // 名前
	AtBats        int    // 打席数
	Hits          int    // 安打数
	Walks         int    // 四球数
	HomeRuns      int    // ホームラン数
	RunsBattedIn  int    // 打点数
}

func CreatePlayer(
	UniformNumber int,
	Name string,
	AtBats int,
	Hits int,
	Walks int,
	HomeRuns int,
	RunsBattedIn int,
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
