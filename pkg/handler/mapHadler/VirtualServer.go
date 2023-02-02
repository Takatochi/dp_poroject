package mapHadler

import (
	"fmt"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql/analyzer"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/emirpasic/gods/trees/redblacktree"
	"project/pkg/Database/VirtualSql"
	"project/pkg/MYSQLserver"
	"project/pkg/logger"
	"project/pkg/stringFMT"
	"time"
)

const (
	Version = "Virtual server"
)

type ListServerSql struct {
	Port   int32
	Name   string
	Server *MYSQLserver.MySqli
	Config *VirtualSql.ConfigVirtual
}

//var listServerSql []ListServerSql

var serverTree = redblacktree.NewWithIntComparator()

func NewServerSql(address, dbname string, port int32) (*redblacktree.Tree, error) {

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
	nameDB := stringFMT.StringTitleJoin(dbname)
	db := memory.NewDatabase(nameDB)
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

	//listServerSql = append(listServerSql, ListServerSql{Port: port, Server: srv})
	serverTree.Put(int(port), ListServerSql{Port: port, Name: dbname, Server: srv, Config: &VirtualSql.ConfigVirtual{
		DatabaseURL: fmt.Sprintf("root:root@tcp(127.0.0.1:%d)/%s", port, nameDB),
		DriverName:  "mysql",
	}})

	select {
	case err, open := <-errs:
		if open {
			return nil, err
		}
	case <-time.After(time.Millisecond):

	}

	return serverTree, nil
}

func startVirtualSqlserver(srv *MYSQLserver.MySqli) error {

	return srv.Server.Start()
}
