package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewEngine is init gin engine
func NewEngine() *gin.Engine {
	api := gin.New()
	api.MaxMultipartMemory = 2 << 20 // 2 MB
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	api.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // TODO：本番環境では許可するオリジンを指定する
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With"},
	}))
	return api
}
