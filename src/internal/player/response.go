package internal

import usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"

type createPlayerResponse struct {
	Id            string `json:"id"`
	UniformNumber int64  `json:"uniform_number"`
	Name          string `json:"name"`
	AtBats        int64  `json:"at_bats"`
	Hits          int64  `json:"hits"`
	Walks         int64  `json:"walks"`
	HomeRuns      int64  `json:"home_runs"`
	RunsBattedIn  int64  `json:"runs_batted_in"`
}

func convertcreatePlayerResponse(d *usecase.CreatePlayerUsecaseDto) *createPlayerResponse {
	return &createPlayerResponse{
		Id:            d.ID,
		UniformNumber: d.UniformNumber,
		Name:          d.Name,
		AtBats:        d.AtBats,
		Hits:          d.Hits,
		Walks:         d.Walks,
		HomeRuns:      d.HomeRuns,
		RunsBattedIn:  d.RunsBattedIn,
	}
}
