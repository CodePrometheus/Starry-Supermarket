package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"net/smtp"
	"shop-web/user-api/forms"
	"shop-web/user-api/global"
	"shop-web/user-api/utils"
	"strings"
	"time"
)

func GetEmail(ctx *gin.Context) {
	emailForm := forms.EmailForm{}
	if err := ctx.ShouldBind(&emailForm); err != nil {
		utils.HandleFormError(ctx, err)
		return
	}
	code := GenerateCode(6)
	e := &email.Email{
		To:      []string{emailForm.Email},
		From:    global.ServerConfig.EmailInfo.From,
		Text:    []byte("code -> " + code),
		Subject: global.ServerConfig.EmailInfo.Subject,
	}
	host := global.ServerConfig.EmailInfo.Host
	port := global.ServerConfig.EmailInfo.Port
	if err := e.Send(host+":"+port, smtp.PlainAuth("",
		global.ServerConfig.EmailInfo.Username,
		global.ServerConfig.EmailInfo.Password, host)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "发送邮件失败",
		})
		zap.S().Errorf("发送邮件失败: %s", err.Error())
		return
	}
	utils.SetKey(emailForm.Email, code)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送邮件成功",
	})
}

func GenerateCode(width int) string {
	// 生成width长度的验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			zap.S().Errorf("生成验证码错误: %s", err.Error())
		}
	}
	return sb.String()
}
