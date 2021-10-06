package initialize

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"shop-web/user-api/global"
	customValidator "shop-web/user-api/validator"
)

func BindingValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", customValidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			tag, _ := ut.T("mobile", fe.Field())
			return tag
		})
	}
}
