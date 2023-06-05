package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/config"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/queryImpl"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	"github.com/tomoya_kamaji/go-pkg/src/pkg/http"
	"github.com/tomoya_kamaji/go-pkg/src/query"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
)

// SearchPlayer godoc
// @Summary 選手を検索する
// @Description　選手を作成する
// @Tags players
// @Accept json
// @Produce json
// @Param query query searchPlayerRequerst false "検索クエリ"
// @Success 200 {object} searchPlayerResponse
// @Router /v1/players/search [get]
func SearchPlayer(ctx *gin.Context) {
	var req searchPlayerRequerst
	if err := http.ValidateBindJSON(ctx, &req); err != nil {
		http.Return400(ctx, err)
		return
	}
	db := gorm.NewMainDB()

	sortField := query.SortField(req.SortField)
	sortOrder := query.SortOrder(req.SortOrder)

	dto, err := usecase.NewSearchPlayerUsecase(
		queryImpl.NewPlayerSearchQueryImpl(db),
	).Run(ctx, query.SearchParam{
		MinHits:         req.MinHits,
		MaxHits:         req.MaxHits,
		MinHomeRuns:     req.MinHomeRuns,
		MaxHomeRuns:     req.MaxHomeRuns,
		MinRunsBattedIn: req.MinRunsBattedIn,
		MaxRunsBattedIn: req.MaxRunsBattedIn,
		SortField:       &sortField,
		SortOrder:       &sortOrder,
	})

	if err != nil {
		config.GetLogger().Error(ctx, err.Error())
		http.Return500(ctx, err)
		return
	}

	res := searchPlayerResponse{}
	for _, p := range dto.Players {
		res.Players = append(res.Players, convertPlayerResponseModel(
			p.ID,
			p.UniformNumber,
			p.Name,
			p.AtBats,
			p.Hits,
			p.Walks,
			p.HomeRuns,
			p.RunsBattedIn,
		))
	}
	http.Return200(ctx, res)
}

// FetchPlayerById godoc
// @Summary 選手を取得する
// @Description　選手を作成する
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} fetchPlayerResponse
// @Router /v1/players/{id} [get]
func FetchPlayerById(ctx *gin.Context) {
	id := ctx.Param("id")
	db := gorm.NewMainDB()
	dto, err := usecase.NewFetchPlayerUsecase(
		repositoryImpl.NewPlayerRepositoryImpl(db),
	).Run(ctx, id)
	if err != nil {
		config.GetLogger().Error(ctx, err.Error())
		http.Return404(ctx, err)
		return
	}
	res := fetchPlayerResponse{
		Player: convertPlayerResponseModel(
			dto.ID,
			dto.UniformNumber,
			dto.Name,
			dto.AtBats,
			dto.Hits,
			dto.Walks,
			dto.HomeRuns,
			dto.RunsBattedIn,
		),
	}
	http.Return200(ctx, res)
}

// CreateUser godoc
// @Summary 選手を作成する
// @Description　選手を作成する
// @Tags players
// @Accept json
// @Produce json
// @Param player body createPlayerRequest true "User information"
// @Success 201 {object} createPlayerResponse
// @Router /v1/players [post]
func CreatePlayer(ctx *gin.Context) {
	var req createPlayerRequest
	if err := http.ValidateBindJSON(ctx, &req); err != nil {
		http.Return400(ctx, err)
		return
	}

	db := gorm.NewMainDB()
	param := usecase.CreatePlayerUsecaseParam{
		UniformNumber: req.UniformNumber,
		Name:          req.Name,
		AtBats:        req.AtBats,
		Hits:          req.Hits,
		Walks:         req.Walks,
		HomeRuns:      req.HomeRuns,
		RunsBattedIn:  req.RunsBattedIn,
	}
	dto, err := usecase.NewCreatePlayerUsecase(
		repositoryImpl.NewTransactionManagerImpl(db),
		repositoryImpl.NewPlayerRepositoryImpl(db),
	).Run(ctx, param)
	if err != nil {
		http.Return500(ctx, err)
		return
	}
	res := createPlayerResponse{
		Player: convertPlayerResponseModel(
			dto.ID,
			dto.UniformNumber,
			dto.Name,
			dto.AtBats,
			dto.Hits,
			dto.Walks,
			dto.HomeRuns,
			dto.RunsBattedIn,
		),
	}
	http.Return201(ctx, res)
}

// CreateUser godoc
// @Summary 選手情報のクロール
// @Description　"https://baseball-data.com/stats/hitter-%v/tpa-1.html" から選手情報をクロールする
// @Tags players
// @Accept json
// @Produce json
// @Success 204
// @Router /v1/players/crawl [post]
func Crawler(ctx *gin.Context) {
	db := gorm.NewMainDB()
	f := func() {
		err := usecase.NewCrawlPlayerUsecase(
			repositoryImpl.NewTransactionManagerImpl(db),
			repositoryImpl.NewPlayerRepositoryImpl(db),
		).Run(ctx)
		if err != nil {
			config.GetLogger().Error(ctx, err.Error())
			return
		}
	}
	go f()
	http.Return204(ctx)
}
