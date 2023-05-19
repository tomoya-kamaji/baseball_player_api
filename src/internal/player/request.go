package internal

type createPlayerRequest struct {
	UniformNumber int64  `json:"uniform_number" binding:"required"`
	Name          string `json:"name" binding:"required"`
	AtBats        int64  `json:"at_bats" binding:"required"`
	Hits          int64  `json:"hits" binding:"required"`
	Walks         int64  `json:"walks" binding:"required"`
	HomeRuns      int64  `json:"home_runs" binding:"required"`
	RunsBattedIn  int64  `json:"runs_batted_in" binding:"required"`
}
