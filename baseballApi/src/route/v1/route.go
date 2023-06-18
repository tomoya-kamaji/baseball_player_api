package v1

import (
	"github.com/gin-gonic/gin"
	healthIn "github.com/tomoya_kamaji/go-pkg/src/internal/health"
	playerIn "github.com/tomoya_kamaji/go-pkg/src/internal/player"
)

func Init(api *gin.Engine) {
	r := api.Group("/v1")

	pl := r.Group("/players")
	{
		r.GET("/health", healthIn.HealthCheck)
		pl.GET("/search", playerIn.SearchPlayer)
		pl.GET("/:id", playerIn.FetchPlayerById)
		pl.POST("", playerIn.CreatePlayer)
		pl.POST("/crawl", playerIn.Crawler)
	}
}
