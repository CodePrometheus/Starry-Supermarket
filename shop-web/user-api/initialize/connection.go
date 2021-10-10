package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop-web/user-api/global"
	"shop-web/user-api/proto"
)

func InitServiceConn() {
	userConnection, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserServiceInfo.Host,
		global.ServerConfig.UserServiceInfo.Port),
		grpc.WithInsecure())
	if err != nil {
		zap.S().Panicf("连接失败: %s", err.Error())
	}
	// 生成grpc的client并调用接口
	userServiceClient := proto.NewUserClient(userConnection)
	global.UserServiceClient = userServiceClient
}
