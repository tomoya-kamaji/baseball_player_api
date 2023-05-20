package route

import (
	"github.com/gin-gonic/gin"
)

// NewEngine is init gin engine
func NewEngine() *gin.Engine {
	api := gin.New()
	api.MaxMultipartMemory = 2 << 20 // 2 MB
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	return api
}
