package sqlBd_test

import (
	"testing"

	"project/pkg/model"
	"project/pkg/store/sqlBd"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlBd.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlBd.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

// func TestUserRepository_Find(t *testing.T) {
// 	db, teardown := sqlBd.TestDB(t, databaseURL)
// 	defer teardown("users")

// 	s := sqlBd.New(db)
// 	u1 := model.TestUser(t)
// 	s.User().Create(u1)
// 	u2, err := s.User().Find(u1.ID)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, u2)
// }
