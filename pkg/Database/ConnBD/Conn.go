package ConnBD

import (
	"database/sql"
	"time"
)

func Conn(DriverName, DatabaseURL string) (*sql.DB, error) {
	db, err := sql.Open(DriverName, DatabaseURL)
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
