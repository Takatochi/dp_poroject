package sqlBd

import (
	"database/sql"
	"fmt"
	"project/pkg/store"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	listRepository *ListRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Server interface for add parameters Store.Server()....
func (s *Store) Server() store.ListRepository {
	fmt.Println("save")
	if s.listRepository != nil {
		return s.listRepository
	}
	//s.db.Begin()
	s.listRepository = &ListRepository{
		store: s,
	}

	return s.listRepository
}
