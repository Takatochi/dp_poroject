package sqlBd

import (
	"database/sql"
	"project/app/model"
	"project/pkg/store"
)

// ListRepository ...
type ListRepository struct {
	store *Store
}

func (r *ListRepository) Find() (*[]model.Server, error) {
	var Uarr []model.Server

	u := &model.Server{}

	res, err := r.store.db.Query(
		"SELECT * FROM `Server`",
	)
	for res.Next() {

		err := res.Scan(
			&u.Id,
			&u.Name,
			&u.Port,
		)
		if err != nil {
			return nil, err

		}

		Uarr = append(Uarr, *u)
	}

	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	return &Uarr, nil
}
func (r *ListRepository) AddServer(u *model.Server) error {

	res, err := r.store.db.Exec(
		"INSERT INTO Server (name, port) VALUES (?, ?)",
		u.Name, u.Port)
	if err != nil {
		return err
	}
	u.Id, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
