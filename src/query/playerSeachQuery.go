package query

import "context"

// SearchParam is a parameter for search
type SearchParam struct {
	MinHits         int64  // ヒット数の最小値
	MaxHits         int64  // ヒット数の最大値
	MinHomeRuns     int64  // 本塁打数の最小値
	MaxHomeRuns     int64  // 本塁打数の最大値
	MinRunsBattedIn int64  // 打点数の最小値
	MaxRunsBattedIn int64  // 打点数の最大値
	SortField       string // ソートフィールド
	SortOrder       string // ソート順 (asc: 昇順, desc: 降順)
	Limit           int    // 取得件数の制限
}

type PlayerSearchQuery interface {
	Run(c context.Context, param SearchParam) (*PlayerSearchQueryDTO, error)
}

type PlayerSearchQueryDTO struct {
	Players []*Player
}

type Player struct {
	ID            string
	UniformNumber int64
	Name          string
	AtBats        int64
	Hits          int64
	Walks         int64
	HomeRuns      int64
	RunsBattedIn  int64
}
