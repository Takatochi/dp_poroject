package store

import "project/pkg/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

// Store ...
type Store interface {
	User() UserRepository
}
