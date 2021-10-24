package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop-web/user-api/config"
	"shop-web/user-api/proto"
)

var (
	ServerConfig = &config.ServerConfig{}

	Trans ut.Translator

	UserServiceClient proto.UserClient

	NacosConfig = &config.NacosConfig{}
)
