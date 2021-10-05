package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-web/user-api/router"
)

func Routers() *gin.Engine {
	routes := gin.Default()
	ApiGroup := routes.Group("u/v1")
	router.InitUserRouter(ApiGroup)
	return routes
}
