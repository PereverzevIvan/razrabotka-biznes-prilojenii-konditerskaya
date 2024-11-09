package configs

import (
	"flag"
	"fmt"
)

const (
	CONFIG_PATH_PARAM_NAME   = "config_path"
	CONFIG_ROUTES_PARAM_NAME = "routes_path"
)

var (
	ConfigPath string
	RoutesPath string
)

type cmdParam[T any] struct {
	Name     string
	ValuePtr *T
}

func LoadCmdParams() error {
	stringParams := []cmdParam[string]{
		{CONFIG_PATH_PARAM_NAME, &ConfigPath},
		{CONFIG_ROUTES_PARAM_NAME, &RoutesPath},
	}

	for _, param := range stringParams {
		flag.StringVar(param.ValuePtr, param.Name, "", "")
	}

	flag.Parse()

	for _, param := range stringParams {
		if *param.ValuePtr == "" {
			return fmt.Errorf("cmd param %s is empty", param.Name)
		}
	}
	return nil
}
