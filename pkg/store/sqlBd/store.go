package sqlBd

import (
	"database/sql"
	"project/pkg/store"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db              *sql.DB
	userRepository  *UserRepository
	tokenRepository *TokenRepository
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

// Token ...
func (s *Store) Token() store.TokenRepository {
	if s.userRepository != nil {
		return s.tokenRepository
	}

	s.tokenRepository = &TokenRepository{
		store: s,
	}

	return s.tokenRepository
}
