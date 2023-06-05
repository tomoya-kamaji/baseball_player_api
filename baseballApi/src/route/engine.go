package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewEngine is init gin engine
func NewEngine() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8002"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type"}

	router.Use(cors.New(config))
	return router
}
