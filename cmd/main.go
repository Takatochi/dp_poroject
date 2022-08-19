package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"project/app"
	srv "project/app/server"
	"project/pkg/logger"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {

	flag.Parse()
	//init config ...
	config := srv.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		logger.Warnf("Warn to stop config: %v", err)
		return
	}
	// Run App ...
	app.Run(config)

}
