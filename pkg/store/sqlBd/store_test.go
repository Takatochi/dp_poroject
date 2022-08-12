package sqlBd_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "root:@tcp(127.0.0.1:3306)/golang"
	}

	os.Exit(m.Run())
}
