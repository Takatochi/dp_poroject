package mysqldump

import (
	"database/sql"
	"io/ioutil"
	"log"
	"project/pkg/logger"
	"strings"
)

func Load(db *sql.DB, path string) error {

	sqlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Split the SQL file into individual statements
	statements := parseSQL(sqlFile)

	// Execute each statement individually
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			return err
		}
	}
	logger.Info("Database load and go for job")
	return nil
}
func parseSQL(file []byte) []string {
	return strings.Split(string(file), ";")
}
