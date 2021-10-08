package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop-web/user-api/global"
	"shop-web/user-api/initialize"
	"shop-web/user-api/utils"
)

func main() {
	// 初始化logger
	initialize.Logger()
	// 初始化config
	initialize.InitConfig()
	// 初始化routers
	Router := initialize.Routers()
	// 初始化国际化
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorw("国际化失败: ", err.Error())
	}
	// 注册验证器
	initialize.BindingValidate()
	// 初始化连接
	initialize.InitServiceConn()
	// 初始化Redis
	utils.InitRedis()

	port := global.ServerConfig.Port
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Errorw("启动失败: ", err.Error())
	}

}
