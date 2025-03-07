package app

import (
	"fmt"
	"net"
)

func (app *App) initNetListener() error {
	var err error
	app.lis, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", app.config.ServerConfig.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return nil
}
