package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
)

func NewSQLiteConnection(data string) *sql.DB {
	db, err := sql.Open("sqlite3", withExecPath(data))
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func withExecPath(data string) string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal("Error getting executable path:", err)
	}
	execDir := filepath.Dir(execPath)
	fullPath := filepath.Join(execDir, data)
	return fullPath
}
