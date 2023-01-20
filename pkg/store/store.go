package store

import (
	"fmt"
	"project/app/model"
)

type ListenStore interface {
	StoreBD() Store
}

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
}

// ListRepository ...
type ListRepository interface {
	Find() (*[]model.Server, error)
	AddServer(u *model.Server) error
	DeleteServerFromDB(id int) error
}

// Store ...
type Store interface {
	User() UserRepository
	Server() ListRepository
}
type Listen struct {
	Store
}

func (l *Listen) StoreBD() Store {
	fmt.Println("load")
	return l
}
