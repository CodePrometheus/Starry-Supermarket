package main

import (
	"shop-web/user-api/initialize"
	"shop-web/user-api/utils"
)

var (
	tran = "zh"
)

func main() {
	// 初始化logger
	initialize.Logger()
	// 初始化config
	initialize.InitConfig()
	// 初始化routers
	Router := initialize.Routers()
	// 初始化国际化
	initialize.InitTrans(tran)
	// 注册验证器
	initialize.BindingValidate()
	// 初始化连接
	initialize.InitServiceConn()
	// 初始化Redis
	utils.InitRedis()
	// 初始化Consul客户端
	initialize.InitConsulClient()
	// 初始化Consul服务
	client, serviceId, Name, Port := initialize.InitConsul()
	// 启动路由
	go func() {
		initialize.GoRouters(Router, Port)
	}()
	initialize.OnQuit(client, serviceId, Name)
}
