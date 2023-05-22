package internal

type searchPlayerRequerst struct {
	MinHits         int64  `json:"min_hits" binding:"omitempty" format:"int64" example:"10"`
	MaxHits         int64  `json:"max_hits" binding:"omitempty" format:"int64" example:"20"`
	MinHomeRuns     int64  `json:"min_home_runs" binding:"omitempty" format:"int64" example:"5"`
	MaxHomeRuns     int64  `json:"max_home_runs" binding:"omitempty" format:"int64" example:"10"`
	MinRunsBattedIn int64  `json:"min_runs_batted_in" binding:"omitempty" format:"int64" example:"10"`
	MaxRunsBattedIn int64  `json:"max_runs_batted_in" binding:"omitempty" format:"int64" example:"20"`
	SortField       string `json:"sort_field" binding:"omitempty" format:"string" example:"hits"`
	SortOrder       string `json:"sort_order" binding:"omitempty" format:"string" example:"asc"`
	Limit           int    `json:"limit" binding:"omitempty" format:"int" example:"10"`
}
type createPlayerRequest struct {
	UniformNumber int64  `json:"uniform_number" binding:"required" format:"int64" example:"1"`
	Name          string `json:"name" binding:"required" format:"string" example:"田中太郎"`
	AtBats        int64  `json:"at_bats" binding:"required" format:"int64" example:"100"`
	Hits          int64  `json:"hits" binding:"required" format:"int64" example:"30"`
	Walks         int64  `json:"walks" binding:"required" format:"int64" example:"10"`
	HomeRuns      int64  `json:"home_runs" binding:"required" format:"int64" example:"5"`
	RunsBattedIn  int64  `json:"runs_batted_in" binding:"required" format:"int64" example:"20"`
}
