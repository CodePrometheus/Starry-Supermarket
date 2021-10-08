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
		_ = v.RegisterValidation("email", customValidator.ValidateEmail)
		_ = v.RegisterTranslation("email", global.Trans, func(ut ut.Translator) error {
			return ut.Add("email", "{0} 非法的邮箱", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			tag, _ := ut.T("email", fe.Field())
			return tag
		})
	}
}
