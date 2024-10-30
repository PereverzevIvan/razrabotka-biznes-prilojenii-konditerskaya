package configs

type Config struct {
	ConfigServer   ServerConfig    `yaml:"server" env-required:"true"`
	AllowedOrigins []AllowedOrigin `yaml:"allowed_origins" env-required:"true"`
}

type ServerConfig struct {
	Port int `yaml:"port" env-required:"true"`
}

type AllowedOrigin struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port"`
}
