package configs

import "flag"

const (
	CONFIG_PATH_PARAM_NAME   = "config_path"
	CONFIG_ROUTES_PARAM_NAME = "routes_path"
)

var (
	ConfigPath string
	RoutesPath string
)

func MustLoadCmdParams() {
	flag.StringVar(&ConfigPath, CONFIG_PATH_PARAM_NAME, "", "")
	flag.StringVar(&RoutesPath, CONFIG_ROUTES_PARAM_NAME, "", "")
	flag.Parse()
}
