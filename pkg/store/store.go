package store

import (
	"database/sql"
	"project/pkg/server"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db     *sql.DB
	config *server.Config
}

func New(config *server.Config) *Store {
	return &Store{
		config: config,
	}
}

func (this *Store) Open() error {
	db, err := sql.Open("mysql", this.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err

	}
	return nil
}
func (this *Store) Close() {
	this.db.Close()
}
