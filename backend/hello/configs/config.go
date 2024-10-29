package configs

type Config struct {
	ConfigServer ServerConfig `yaml:"server" env-required:"true"`
}

type ServerConfig struct {
	Port int `yaml:"port" env-required:"true"`
}
