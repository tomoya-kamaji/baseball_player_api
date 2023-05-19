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
	// api.Use(apperror.Handle400())
	// api.Use(apperror.Handle500())
	// api.NoRoute(apperror.NotFound)
	// api.NoMethod(apperror.NoMethod)

	return api
}
