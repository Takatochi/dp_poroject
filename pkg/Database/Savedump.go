package Database

import (
	"database/sql"
	"fmt"
	"os"
	mysqldump "project/pkg/Database/dumpmysql"
	"project/pkg/logger"
)

func SaveFile(dumpDir, dbname string, db *sql.DB) error {
	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", dbname)
	if err := os.MkdirAll(dumpDir, 0755); err != nil {
		logger.Error("Error mkdir:", err)
		return err
	}
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		logger.Error("Error registering database:", err)
		return err
	}
	// Dump database to file
	if err := dumper.Dump(); err != nil {
		logger.Error("Error dumping:", err)
		return err
	}
	if file, ok := dumper.Out.(*os.File); ok {
		logger.Info("File is saved to", file.Name())
	} else {
		logger.Info("It's not part of *os.File, but dump is done")
	}
	return nil
}
