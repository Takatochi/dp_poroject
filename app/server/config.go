package server

import (
	"fmt"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/server"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	DriverName  string `toml:"driverName"`
}

type MysqliConfig struct {
	DbName    string
	TableName string
	USER      string
	PASSWORD  string
	Address   string
	PORT      int64
	Version   string
	Cfg       *sqle.Config
	Config    *server.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8000",
	}
}

func NewMysqliConfig(DbName string, address string, PORT int64, Version string, cfg *sqle.Config, config *server.Config) *MysqliConfig {
	if cfg == nil {
		cfg = &sqle.Config{
			VersionPostfix:     Version,
			IsReadOnly:         false,
			IsServerLocked:     false,
			IncludeRootAccount: false,
		}
	}
	if config == nil {
		config = &server.Config{
			Protocol: "tcp",
			Address:  fmt.Sprintf("%s:%d", address, PORT),
			Version:  Version,
		}
	}
	return &MysqliConfig{
		DbName:  DbName,
		Address: address,
		PORT:    PORT,
		Cfg:     cfg,
		Config:  config,
	}
}
