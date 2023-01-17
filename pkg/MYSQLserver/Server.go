package MYSQLserver

import (
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql/analyzer"
)

const (
	Version = "8.14 gorgon medusa"
)

type MySqli struct {
	Server server.Server
	// server external config

	Config *server.Config

	// memory database
	Database *memory.Database

	// server internal init data in system
	Analyzer *analyzer.Analyzer

	// server internal cfg
	cfg *sqle.Config
}

// NewMySqli creates a new MySqli struct
func NewMySqliDefault(config *server.Config, analyzer *analyzer.Analyzer, cfg *sqle.Config) *MySqli {
	if cfg == nil {
		cfg = &sqle.Config{
			VersionPostfix:     Version,
			IsReadOnly:         false,
			IsServerLocked:     false,
			IncludeRootAccount: false,
		}
	}
	return &MySqli{
		Config:   config,
		Analyzer: analyzer,
		cfg:      cfg,
	}

}

func (this *MySqli) Run() error {

	engine := sqle.New(this.Analyzer, this.cfg)

	s, err := server.NewDefaultServer(*this.Config, engine)

	defer s.Close()

	if err != nil {
		return err
	}

	return s.Start()
}
