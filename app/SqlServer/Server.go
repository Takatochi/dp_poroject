package SqlServer

import (
	"fmt"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/go-kit/kit/metrics/prometheus"
	promopts "github.com/prometheus/client_golang/prometheus"
)

// Example of how to implement a MySQL server based on a Engine:
//
// ```
// > mysql --host=127.0.0.1 --port=3309 -u root mydb -e "SELECT * FROM mytable" mysql --host=127.0.0.1 --port=3309 -u root mydb <J:\go-mysql-server-master\dumps\wds_1212yy12-20221024T001234.sql
// +----------+-------------------+-------------------------------+---------------------+
// | name     | email             | phone_numbers                 | created_at          |
// +----------+-------------------+-------------------------------+---------------------+
// | John Doe | john@doe.com      | ["555-555-555"]               | 2018-04-18 09:41:13 |
// | John Doe | johnalt@doe.com   | []                            | 2018-04-18 09:41:13 |
// | Jane Doe | jane@doe.com      | []                            | 2018-04-18 09:41:13 |
// | Evil Bob | evilbob@gmail.com | ["555-666-555","666-666-666"] | 2018-04-18 09:41:13 |
// +----------+-------------------+-------------------------------+---------------------+
// ```
const (
	dbName    = "ww3y-34"
	tableName = "mytable"
)

func Start() error {
	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			createTestDatabase(),
			information_schema.NewInformationSchemaDatabase(),
		))
	engine.Analyzer.Catalog.MySQLDb.AddSuperUser("root", "root")
	server.QueryCounter = prometheus.NewCounterFrom(promopts.CounterOpts{
		Namespace: "go_mysql_server",
		Subsystem: "engine",
		Name:      "query_counter",
	}, []string{
		"query",
	})
	fmt.Println(server.QueryCounter)

	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3309",
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	return s.Start()

}

func createTestDatabase() *memory.Database {

	db := memory.NewDatabase(dbName)
	//ctx := sql.NewEmptyContext()
	//s := sql.Schema{{AutoIncrement: true, Name: "id", Type:sql.int, Nullable: false, Source: tableName}}
	//
	//db.CreateTable(ctx, tableName, sql.NewPrimaryKeySchema(s))

	//table := memory.NewTable(tableName, sql.NewPrimaryKeySchema(s), nil)
	//table.CreateIndex(ctx, "id")
	//table := memory.NewTable(tableName, sql.NewPrimaryKeySchema(sql.Schema{
	//	{Name: "name", Type: sql.Text, Nullable: false, Source: tableName},
	//	{Name: "email", Type: sql.Text, Nullable: false, Source: tableName},
	//	{Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
	//	{Name: "created_at", Type: sql.Datetime, Nullable: false, Source: tableName},
	//}), nil)

	//db.AddTable(tableName, table)
	//
	//_ = table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", sql.MustJSON(`["555-555-555"]`), time.Now()))
	//_ = table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", sql.MustJSON(`[]`), time.Now()))
	//_ = table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", sql.MustJSON(`[]`), time.Now()))
	//_ = table.Insert(ctx, sql.NewRow("Jane Deo", "janedeo@gmail.com", sql.MustJSON(`["556-565-566", "777-777-777"]`), time.Now()))

	return db
}
