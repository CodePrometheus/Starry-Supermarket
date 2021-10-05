package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop-web/user-api/initialize"
)

func main() {
	// 初始化logger
	initialize.Logger()
	// 初始化routers
	Router := initialize.Routers()
	port := 9000
	zap.S().Infof("启动服务器, 端口: %d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}

}
