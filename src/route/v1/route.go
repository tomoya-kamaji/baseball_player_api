package v1

import (
	"github.com/gin-gonic/gin"
	internal "github.com/tomoya_kamaji/go-pkg/src/internal/player"
)

// @title Baseball API
// @version バージョン(1.0)
// @description APIの説明(必須)
// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /v1
func Init(api *gin.Engine) {
	r := api.Group("/v1")

	pl := r.Group("/players")
	{
		pl.POST("", internal.CreatePlayer)
	}
}
