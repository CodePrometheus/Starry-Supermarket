package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop-web/user-api/global"
)

func InitConfig() {
	configFileName := fmt.Sprintf("user-api/config-dev.yaml")
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic("读取配置信息失败: ", err.Error())
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panic("解析配置信息失败: ", err.Error())
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)
}
