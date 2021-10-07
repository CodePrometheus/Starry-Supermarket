package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-web/user-api/global"
	"shop-web/user-api/models"
	"time"
)

type JWT struct {
	Key []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("it's not even a token")
	TokenInvalid     = errors.New("could not handle the token")
)

// JwtAuth Jwt认证
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "您还未登录",
			})
			c.Abort()
			return
		}
		j := NewJwt()
		claim, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, map[string]string{
						"msg": "授权已过期",
					})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, "未登录")
			c.Abort()
			return
		}
		// 存入上下文
		c.Set("claims", claim)
		c.Set("userId", claim.Id)
		c.Next()
	}
}

// NewJwt 实例Token
func NewJwt() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JwtInfo.Key),
	}
}

// CreateToken 创建Token
func (j *JWT) CreateToken(claims models.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Key)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*models.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.Key, nil
		})
	if err != nil {
		if value, ok := err.(*jwt.ValidationError); ok {
			if value.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if value.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else if value.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 刷新Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.Key, nil
		})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
