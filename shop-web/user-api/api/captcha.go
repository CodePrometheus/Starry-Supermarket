package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"image/color"
	"net/http"
	"shop-web/user-api/utils"
)

var store = utils.RedisStore{}

func GetCaptcha(ctx *gin.Context) {
	rgba := color.RGBA{
		R: 3,
		G: 102,
		B: 214,
		A: 254,
	}
	fonts := []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverMath(50, 140, 0, 0, &rgba, nil, fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	fmt.Println("res: " + b64s)
	if err != nil {
		zap.S().Errorw("生成验证码错误: ", err.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"pic":       b64s,
	})
}
