package router

import (
	"github.com/gin-gonic/gin"
	"shop-web/user-api/api"
	"shop-web/user-api/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JwtAuth(),
			middlewares.IsAdmin(), api.GetUserList)
		UserRouter.POST("login", api.Login)
		UserRouter.POST("register", api.Register)
	}
}
