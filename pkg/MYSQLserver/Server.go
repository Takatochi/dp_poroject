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
	Server *server.Server
	// server external config

	Config *server.Config

	Engine *sqle.Engine
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
	//ctx := sql.NewEmptyContext()

	engine := sqle.New(this.Analyzer, this.cfg)

	s, err := server.NewDefaultServer(*this.Config, engine)

	defer s.Close()

	if err != nil {
		return err
	}
	//s.Close()
	return s.Start()
}
func (this *MySqli) VirtualRun() (*MySqli, error) {
	//ctx := sql.NewEmptyContext()

	engine := sqle.New(this.Analyzer, this.cfg)

	s, err := server.NewDefaultServer(*this.Config, engine)

	//defer s.Close()

	if err != nil {
		return nil, err
	}
	//s.Close()

	return &MySqli{
		Server: s,
		Engine: engine,
	}, nil
}
func (this *MySqli) Stop() error {
	return this.Server.Close()
}
