package v1

import (
	"github.com/gin-gonic/gin"
	healthIn "github.com/tomoya_kamaji/go-pkg/src/internal/health"
)

func Init(api *gin.Engine) {
	r := api.Group("/v1")
	r.GET("/health", healthIn.HealthCheck)

	// pl := r.Group("/players")
	{
		// pl.GET("/search", playerIn.SearchPlayer)
		// pl.GET("/:id", playerIn.FetchPlayerById)
		// pl.POST("", playerIn.CreatePlayer)
		// pl.POST("/crawl", playerIn.Crawler)
	}
}
