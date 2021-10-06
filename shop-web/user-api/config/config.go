package config

type UserServiceConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name            string            `mapstructure:"name" json:"name"`
	Port            int               `mapstructure:"port" json:"port"`
	UserServiceInfo UserServiceConfig `mapstructure:"user-service" json:"user_service"`
}
