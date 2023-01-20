package mapHadler

import (
	"fmt"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql/analyzer"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"project/pkg/MYSQLserver"
	"project/pkg/logger"
	"time"
)

const (
	Version = "Virtual server"
)

type ListServerSql struct {
	Port   int32
	Server *MYSQLserver.MySqli
}

var listServerSql []ListServerSql

func NewServerSql(address, dbname string, port int32) []ListServerSql {
	config := &sqle.Config{
		VersionPostfix:     Version,
		IsReadOnly:         false,
		IsServerLocked:     false,
		IncludeRootAccount: false,
	}

	cfg := &server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
		Version:  Version,
	}
	db := memory.NewDatabase(dbname)
	analyzer := analyzer.NewDefault(analyzer.NewDatabaseProvider(db, information_schema.NewInformationSchemaDatabase()))

	srv := MYSQLserver.NewMySqliDefault(cfg, analyzer, config)
	srv, err := srv.VirtualRun()
	if err != nil {
		logger.Error(err)
	}
	errs := make(chan error, 1)
	go func() {
		err := startVirtualSqlserver(srv)
		if err != nil {
			errs <- err
		}
		close(errs)
	}()

	listServerSql = append(listServerSql, ListServerSql{Port: port, Server: srv})

	select {
	case err, open := <-errs:
		if open {
			logger.Error(err)
		}
	case <-time.After(time.Second):
		// handle timeout
	}
	return listServerSql
}

func startVirtualSqlserver(srv *MYSQLserver.MySqli) error {

	return srv.Server.Start()
}
