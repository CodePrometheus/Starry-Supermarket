package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop-web/user-api/global"
)

func InitConfig() {
	configFileName := fmt.Sprintf("user-api/config.yaml")
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("读取配置信息失败: %s", err.Error())
	}
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		zap.S().Panicf("解析Nacos信息失败: %s", err.Error())
	}

	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	NacosClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		zap.S().Panicf("Nacos客户端连接失败: %s", err.Error())
	}

	data, error := NacosClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})
	if error != nil {
		zap.S().Fatalf("获取Nacos配置信息失败: %s", err.Error())
	}

	// 解析所有配置信息
	if err := json.Unmarshal([]byte(data), &global.ServerConfig); err != nil {
		zap.S().Fatalf("解析Nacos配置信息失败: %s", err.Error())
	}
}
