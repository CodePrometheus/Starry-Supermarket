package main

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	port       int         `mapstructure:"port"`
	MysqlInfo  MysqlConfig `mapstructure:"mysql"`
}

func TestConfig(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
}
