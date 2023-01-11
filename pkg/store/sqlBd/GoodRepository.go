package sqlBd

//
//import (
//	"database/sql"
//	"project/app/model"
//	"project/pkg/store"
//)
//
//// GoodsRepository ...
//type GoodsRepository struct {
//	store *Store
//}
//
//func (r *GoodsRepository) Find(id int) (*model.Goods, error) {
//	u := &model.Goods{}
//	if err := r.store.db.QueryRow(
//		"SELECT * FROM `goods` WHERE `Id` = ?",
//		id,
//	).Scan(
//		&u.Id,
//		&u.Title,
//		&u.Description,
//	); err != nil {
//		if err == sql.ErrNoRows {
//			return nil, store.ErrRecordNotFound
//		}
//
//		return nil, err
//	}
//
//	return u, nil
//}
//func (r *GoodsRepository) GoodsTable() (*[]model.Goods, error) {
//	arru := []model.Goods{}
//	u := &model.Goods{}
//	res, err := r.store.db.Query(
//		"SELECT * FROM `goods` ",
//	)
//	for res.Next() {
//
//		err := res.Scan(
//			&u.Id,
//			&u.Title,
//			&u.Description,
//			&u.Price,
//		)
//		if err != nil {
//			return nil, err
//		}
//		arru = append(arru, *u)
//	}
//	if err != nil && err != sql.ErrNoRows {
//		return nil, err
//	}
//
//	return &arru, nil
//}
