package SqlServer

import (
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/sql/analyzer"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	server2 "project/app/server"
	"project/pkg/MYSQLserver"
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

func Start(MLR *server2.MysqliConfig) error {

	db := memory.NewDatabase(MLR.DbName)

	analyzer := analyzer.NewDefault(analyzer.NewDatabaseProvider(db, information_schema.NewInformationSchemaDatabase()))

	MYs := MYSQLserver.NewMySqliDefault(MLR.Config, analyzer, MLR.Cfg)

	return MYs.Run()

}
