package configs

type RoutesConfig struct {
	Microservices []MicroserviceConfig `yaml:"microservices" env-required:"true"`
}

type MicroserviceConfig struct {
	Url    string  `yaml:"url" env-required:"true"`
	Name   string  `yaml:"name" env-required:"true"`
	Routes []Route `yaml:"routes" env-required:"true"`
}

type Route struct {
	Path        string   `yaml:"path" env-required:"true"`
	Methods     []string `yaml:"methods"`
	Middlewares []string `yaml:"middlewares"`
}

type RouteMethods string

const (
	GET    RouteMethods = "GET"
	POST   RouteMethods = "POST"
	PUT    RouteMethods = "PUT"
	PATCH  RouteMethods = "PATCH"
	DELETE RouteMethods = "DELETE"
)

type RouteMiddleware string

const (
	Auth RouteMiddleware = "auth"
)
