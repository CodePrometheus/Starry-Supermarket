package config

type ServerConfig struct {
	Name            string            `mapstructure:"name" json:"name"`
	Host            string            `mapstructure:"host" json:"host"`
	Port            int               `mapstructure:"port" json:"port"`
	Tags            []string          `mapstructure:"tags" json:"tags"`
	UserServiceInfo UserServiceConfig `mapstructure:"user-service" json:"user_service"`
	JwtInfo         JwtConfig         `mapstructure:"jwt" json:"jwt"`
	RedisInfo       RedisConfig       `mapstructure:"redis" json:"redis"`
	EmailInfo       EmailConfig       `mapstructure:"email" json:"email"`
	ConsulInfo      ConsulConfig      `mapstructure:"consul" json:"consul"`
}

type UserServiceConfig struct {
	Name string `mapstructure:"name" json:"name"`
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
	Expires  int64  `mapstructure:"expires" json:"expires"`
}

type EmailConfig struct {
	From     string `mapstructure:"from" json:"from"`
	Subject  string `mapstructure:"subject" json:"subject"`
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port""`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
