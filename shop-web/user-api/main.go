package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop-web/user-api/global"
	"shop-web/user-api/initialize"
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
		zap.S().Panic("国际化失败: ", err.Error())
	}
	// 注册验证器
	initialize.BindingValidate()
	// 初始化连接
	initialize.InitServiceConn()

	port := global.ServerConfig.Port
	zap.S().Infof("启动服务器, 端口: %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}

}
