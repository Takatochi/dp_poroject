package sqlBd

import (
	"database/sql"
	"project/app/model"
	"project/pkg/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

func (r UserRepository) Create(u *model.User) error {

	return r.store.db.QueryRow(
		"INSERT INTO `user` (Email, Password) VALUES (?, ?)",
		u.Email, u.Password).Err()
}
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT Id, email, encrypted_password FROM user WHERE id = ?",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
