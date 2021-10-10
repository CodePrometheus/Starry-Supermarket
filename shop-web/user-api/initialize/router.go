package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shop-web/user-api/middlewares"
	"shop-web/user-api/router"
)

func Routers() *gin.Engine {
	routes := gin.Default()

	routes.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域`
	routes.Use(middlewares.Cors())
	ApiGroup := routes.Group("u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return routes
}

func GoRouters(Router *gin.Engine, Port int) {
	if err := Router.Run(fmt.Sprintf(":%d", Port)); err != nil {
		zap.S().Panicf("启动失败: %s", err.Error())
	}
}
