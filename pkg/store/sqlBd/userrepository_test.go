package sqlBd_test

import (
	"project/pkg/model"
	"project/pkg/store/sqlBd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlBd.TestDB(t, databaseURL)
	defer teardown("user")

	s := sqlBd.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}
