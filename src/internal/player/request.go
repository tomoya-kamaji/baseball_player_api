package internal

type createPlayerRequest struct {
	UniformNumber int64  `json:"uniform_number" binding:"required" format:"int64" example:"1"`
	Name          string `json:"name" binding:"required" format:"string" example:"田中太郎"`
	AtBats        int64  `json:"at_bats" binding:"required" format:"int64" example:"100"`
	Hits          int64  `json:"hits" binding:"required" format:"int64" example:"30"`
	Walks         int64  `json:"walks" binding:"required" format:"int64" example:"10"`
	HomeRuns      int64  `json:"home_runs" binding:"required" format:"int64" example:"5"`
	RunsBattedIn  int64  `json:"runs_batted_in" binding:"required" format:"int64" example:"20"`
}
