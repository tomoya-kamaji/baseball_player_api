package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	"github.com/tomoya_kamaji/go-pkg/src/pkg/http"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
)

func CreatePlayer(ctx *gin.Context) {
	var req createPlayerRequest
	param := usecase.CreatePlayerUsecaseParam{
		UniformNumber: req.UniformNumber,
		Name:          req.Name,
		AtBats:        req.AtBats,
		Hits:          req.Hits,
		Walks:         req.Walks,
		HomeRuns:      req.HomeRuns,
		RunsBattedIn:  req.RunsBattedIn,
	}

	db := gorm.NewMainDB()
	dto, err := usecase.NewCreatePlayerUsecase(
		repositoryImpl.NewTransactionManagerImpl(db),
		repositoryImpl.NewPlayerRepositoryImpl(db),
	).Run(ctx, param)
	if err != nil {
		http.Return500(ctx, err)
		return
	}

	res := convertcreatePlayerResponse(dto)
	http.Return201(ctx, res)
}
