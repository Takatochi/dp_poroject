package VirtualSql

import (
	"database/sql"
	"project/pkg/Database/ConnBD"
)

type Database interface {
	Open(c *ConfigVirtual) (*sql.DB, error)
}

type VirtualMySQLDatabase struct{}

func (d *VirtualMySQLDatabase) Open(c *ConfigVirtual) (*sql.DB, error) {

	db, err := ConnBD.Conn(c.DriverName, c.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
