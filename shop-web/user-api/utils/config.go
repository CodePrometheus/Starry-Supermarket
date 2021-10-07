package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetConfigKey() *viper.Viper {
	configFileName := fmt.Sprintf("user-api/config-dev.yaml")
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Error(err)
	}
	return v
}
