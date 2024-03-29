package Database

import (
	"database/sql"
	"project/app/server"
	"project/pkg/Database/ConnBD"
)

type Database interface {
	Open(c *server.Config) (*sql.DB, error)
}
type MySQLDatabase struct{}

func (d *MySQLDatabase) Open(c *server.Config) (*sql.DB, error) {

	db, err := ConnBD.Conn(c.DriverName, c.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
