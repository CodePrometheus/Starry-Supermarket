package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-web/user-api/middlewares"
	"shop-web/user-api/router"
)

func Routers() *gin.Engine {
	routes := gin.Default()
	// 配置跨域
	routes.Use(middlewares.Cors())
	ApiGroup := routes.Group("u/v1")
	router.InitUserRouter(ApiGroup)
	return routes
}
