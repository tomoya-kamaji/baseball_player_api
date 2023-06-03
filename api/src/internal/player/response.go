package internal

type createPlayerResponse struct {
	Player playerResponseModel `json:"player"`
}

type fetchPlayerResponse struct {
	Player playerResponseModel `json:"player"`
}

type searchPlayerResponse struct {
	Players []playerResponseModel `json:"players"`
}

type playerResponseModel struct {
	Id            string `json:"id"`
	UniformNumber int64  `json:"uniform_number"`
	Name          string `json:"name"`
	AtBats        int64  `json:"at_bats"`
	Hits          int64  `json:"hits"`
	Walks         int64  `json:"walks"`
	HomeRuns      int64  `json:"home_runs"`
	RunsBattedIn  int64  `json:"runs_batted_in"`
}

func convertPlayerResponseModel(
	id string,
	uniformNumber int64,
	name string,
	atBats int64,
	hits int64,
	walks int64,
	homeRuns int64,
	runsBattedIn int64,
) playerResponseModel {
	return playerResponseModel{
		Id:            id,
		UniformNumber: uniformNumber,
		Name:          name,
		AtBats:        atBats,
		Hits:          hits,
		Walks:         walks,
		HomeRuns:      homeRuns,
		RunsBattedIn:  runsBattedIn,
	}
}
