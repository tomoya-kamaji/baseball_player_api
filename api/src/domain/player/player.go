package domain

import "github.com/google/uuid"

type PlayerID string

func (id PlayerID) String() string {
	return string(id)
}

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

type CreatePlayerParam struct {
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}

func CreatePlayer(param CreatePlayerParam) *Player {
	return &Player{
		ID:            NewPlayerId(),
		UniformNumber: param.UniformNumber,
		Name:          param.Name,
		AtBats:        param.AtBats,
		Hits:          param.Hits,
		Walks:         param.Walks,
		HomeRuns:      param.HomeRuns,
		RunsBattedIn:  param.RunsBattedIn,
	}
}

type ReConstractPlayerParam struct {
	ID            PlayerID
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}

func ReConstractPlayer(param ReConstractPlayerParam) *Player {
	return &Player{
		ID:            param.ID,
		UniformNumber: param.UniformNumber,
		Name:          param.Name,
		AtBats:        param.AtBats,
		Hits:          param.Hits,
		Walks:         param.Walks,
		HomeRuns:      param.HomeRuns,
		RunsBattedIn:  param.RunsBattedIn,
	}
}
