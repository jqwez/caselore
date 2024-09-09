package repo

import (
	"database/sql"
	"os"

	"github.com/jqwez/caselore/repo/sqlite"
)

func GetRepo() *sql.DB {
	return GetRepositoryFromEnv()
}

func GetRepositoryFromEnv() *sql.DB {
	databaseEnv := os.Getenv("DATABASE_SYSTEM")
	if databaseEnv == "" || databaseEnv == "sqlite" {
		dbName := os.Getenv("DATABASE_NAME")
		if dbName == "" {
			dbName = "app.dat"
		}
		return sqlite.NewSQLiteConnection(dbName)
	}
	return sqlite.NewSQLiteConnection("app.dat")
}
