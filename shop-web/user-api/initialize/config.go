package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop-web/user-api/global"
)

func InitConfig() {
	configFileName := fmt.Sprintf("user-api/config.yaml")
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Errorf("读取配置信息失败: %s", err.Error())
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Errorf("解析配置信息失败: %s", err.Error())
	}
}
