package config

type ServerConfig struct {
	Name            string            `mapstructure:"name" json:"name"`
	Port            int               `mapstructure:"port" json:"port"`
	UserServiceInfo UserServiceConfig `mapstructure:"user-service" json:"user_service"`
	JwtInfo         JwtConfig         `mapstructure:"jwt" json:"jwt"`
	RedisInfo       RedisConfig       `mapstructure:"redis" json:"redis"`
}

type UserServiceConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type JwtConfig struct {
	Key     string `mapstructure:"key" json:"key"`
	Expires int64  `mapstructure:"expires" json:"expires"`
	Issuer  string `mapstructure:"issuer" json:"issuer"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}
