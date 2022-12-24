package Database

import (
	"database/sql"
	"project/app/server"
	"time"
)

type Database interface {
	Open(c *server.Config) (*sql.DB, error)
}
type MySQLDatabase struct{}

func (d *MySQLDatabase) Open(c *server.Config) (*sql.DB, error) {
	db, err := sql.Open(c.DriverName, c.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	return db, nil
}
