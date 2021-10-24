package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"os"
	"os/signal"
	"shop-web/user-api/global"
	"shop-web/user-api/proto"
	"shop-web/user-api/utils"
	"syscall"
)

func InitConsulClient() {
	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			consulInfo.Host, consulInfo.Port,
			global.ServerConfig.UserServiceInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	if err != nil {
		zap.S().Panicw("[InitConsul] 连接 【用户服务失败】")
	}
	userSvsClient := proto.NewUserClient(conn)
	global.UserServiceClient = userSvsClient
	zap.S().Infow("[InitConsul] 连接 【用户服务成功】")
}

func InitConsul() (utils.RegisterClient, string, string, int) {
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	// Register
	Name := global.ServerConfig.Name
	Host := global.ServerConfig.Host
	Port := global.ServerConfig.Port
	client := utils.NewRegisterClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	err := client.Register(
		Host,
		Port,
		Name,
		global.ServerConfig.Tags,
		serviceId,
	)
	if err != nil {
		zap.S().Panicf("服务: %s 注册失败", Name)
	} else {
		zap.S().Infof("服务: %s 注册成功", Name)
	}
	return client, serviceId, Name, Port
}

func OnQuit(client utils.RegisterClient, serviceId string, Name string) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	if err := client.Deregister(serviceId); err != nil {
		zap.S().Errorf("服务: %s 注销失败", Name)
	} else {
		zap.S().Infof("服务: %s 注销成功", Name)
	}
}
