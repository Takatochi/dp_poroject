package store

import (
	"project/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

// ListRepository ...
type ListRepository interface {
	Find() (*[]model.Server, error)
	AddServer(u *model.Server) error
}

// Store ...
type Store interface {
	User() UserRepository
	Server() ListRepository
}
