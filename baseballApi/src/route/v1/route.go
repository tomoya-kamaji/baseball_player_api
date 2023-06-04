package v1

import (
	"github.com/gin-gonic/gin"
	internal "github.com/tomoya_kamaji/go-pkg/src/internal/player"
)

func Init(api *gin.Engine) {
	r := api.Group("/v1")

	pl := r.Group("/players")
	{
		pl.GET("/search", internal.SearchPlayer)
		pl.GET("/:id", internal.FetchPlayerById)
		pl.POST("", internal.CreatePlayer)
		pl.POST("/crawl", internal.Crawler)
	}
}
