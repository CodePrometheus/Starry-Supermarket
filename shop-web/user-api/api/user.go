package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"shop-web/user-api/forms"
	"shop-web/user-api/global"
	"shop-web/user-api/global/reponse"
	"shop-web/user-api/middlewares"
	"shop-web/user-api/models"
	"shop-web/user-api/proto"
	"strconv"
	"strings"
	"time"
)

func GetUserList(ctx *gin.Context) {
	// login
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.JwtClaims)
	zap.S().Infof("访问用户: %d", currentUser.Id)

	// 转换
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("pSize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.UserServiceClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})

	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表] 失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	// 返回数据
	res := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := reponse.UserResponse{
			Id:       value.Id,
			Nickname: value.Nickname,
			Birthday: reponse.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender:   value.Gender,
			Email:    value.Email,
		}
		res = append(res, user)
	}
	ctx.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	loginForm := forms.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		HandleFormError(c, err)
		return
	}

	// captcha
	if !store.Verify(loginForm.CaptchaId, loginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
	}

	// login
	if rsp, err := global.UserServiceClient.GetUserByEmail(context.Background(), &proto.EmailRequest{
		Email: loginForm.Email,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "该用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"msg": "登录失败",
				})
			}
			return
		}
	} else {
		// 查询到对应的用户，校验密码
		if passRsp, passErr := global.UserServiceClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          loginForm.Password,
			EncryptedPassword: rsp.Password,
		}); passErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"msg": "登录失败",
			})
		} else {
			if passRsp.Success {
				// token
				j := middlewares.NewJwt()
				now := time.Now().Unix()
				jwtInfo := global.ServerConfig.JwtInfo
				claims := models.JwtClaims{
					Id:          uint(rsp.Id),
					Nickname:    rsp.Nickname,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: now + jwtInfo.Expires, // 过期时间
						NotBefore: now,                   // 生效时间
						Issuer:    jwtInfo.Issuer,        // 签发者
					},
				}

				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "Token颁发失败",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"id":       rsp.Id,
					"nickname": rsp.Nickname,
					"token":    token,
					"expires":  (time.Now().Unix() + jwtInfo.Expires) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}
}

func Register(c *gin.Context) {

}

func HandleFormError(c *gin.Context, error error) {
	// 拦截表单错误
	err, ok := error.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": error.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(err.Translate(global.Trans)),
	})
	return
}

func removeTopStruct(fileds map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fileds {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换成http的状态码
	if err != nil {
		if se, ok := status.FromError(err); ok {
			switch se.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": se.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": se.Code(),
				})
			}
			return
		}
	}
}
