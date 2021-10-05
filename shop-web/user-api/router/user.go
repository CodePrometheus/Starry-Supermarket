package router

import (
	"github.com/gin-gonic/gin"
	"shop-web/user-api/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GetUserList)
	}
}
