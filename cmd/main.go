package main

import (
	"flag"
	"log"
	"project/pkg/server"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	//server ...
	srv := new(server.Server)

	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(config); err != nil {
		log.Fatal(err)
	}

}
