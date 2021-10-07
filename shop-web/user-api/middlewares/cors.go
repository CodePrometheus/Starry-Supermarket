package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 在非简单请求且跨域的情况下，浏览器会发起options预检请求，为了判断实际发送的请求是否安全
// 复杂请求包括 PUT和DELETE，application/json，额外的header
// 解决跨域的是先发一次options请求，获取Access-Control-Allow-Headers，允许跨域之后才会再请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
