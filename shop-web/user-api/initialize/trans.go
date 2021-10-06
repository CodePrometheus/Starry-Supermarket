package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTran "github.com/go-playground/validator/v10/translations/en"
	zhTran "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"shop-web/user-api/global"
	"strings"
)

func InitTrans(tran string) (err error) {
	// 修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(tran)
		if !ok {
			zap.S().Errorf("uni.GetTranslator(%s)", tran)
		}
		switch tran {
		case "en":
			if err := enTran.RegisterDefaultTranslations(v, global.Trans); err != nil {
				zap.S().Panic("国际化en失败: ", err.Error())
			}
		case "zh":
			if err := zhTran.RegisterDefaultTranslations(v, global.Trans); err != nil {
				zap.S().Panic("国际化zh失败: ", err.Error())
			}
		default:
			if err := enTran.RegisterDefaultTranslations(v, global.Trans); err != nil {
				zap.S().Panic("国际化en失败: ", err.Error())
			}
		}
		return
	}
	return
}
