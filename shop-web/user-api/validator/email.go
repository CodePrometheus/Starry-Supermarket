package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if _, err := regexp.MatchString(`/^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/`, email); err != nil {
		zap.S().Errorw("邮箱不合法 ", err.Error())
		return false
	}
	return true
}
