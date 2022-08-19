package sqlBd

import (
	"database/sql"
	"project/app/model"
	"project/pkg/store"
)

// TokenRepository ...
type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Find(id int) (*model.Token, error) {
	u := &model.Token{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM `token` WHERE `Id` = ?",
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Token,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
