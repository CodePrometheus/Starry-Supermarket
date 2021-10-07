package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-web/user-api/models"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.JwtClaims)
		// 2为管理员
		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
