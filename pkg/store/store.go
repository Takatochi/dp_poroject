package store

import (
	"project/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

// TokenRepository ...
type TokenRepository interface {
	Find(int) (*model.Token, error)
}

// Store ...
type Store interface {
	User() UserRepository
	Token() TokenRepository
}
