package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type RegisterInfo struct {
	Host string
	Port int
}

type RegisterClient interface {
	Register(addr string, port int, name string,
		tags []string, id string) error
	Deregister(serviceId string) error
}

func NewRegisterClient(host string, port int) RegisterClient {
	return &RegisterInfo{
		Host: host,
		Port: port,
	}
}

func (r *RegisterInfo) Register(addr string, port int, name string,
	tags []string, id string) error {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Panicf("Consul[NewClient]错误: %s", err.Error())
	}

	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", addr, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	regis := new(api.AgentServiceRegistration)
	regis.Address = addr
	regis.Port = port
	regis.Name = name
	regis.Tags = tags
	regis.ID = id
	regis.Check = check

	if err := client.Agent().ServiceRegister(regis); err != nil {
		zap.S().Panicf("Consul[ServiceRegister] 注册失败: %s", err.Error())
	}
	return nil
}

func (r *RegisterInfo) Deregister(serviceId string) error {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Errorf("Consul[NewClient]错误: %s", err.Error())
		return err
	}
	if err := client.Agent().ServiceDeregister(serviceId); err != nil {
		zap.S().Errorf("Consul[ServiceDeregister]注销失败: %s", err.Error())
	}
	return err
}
