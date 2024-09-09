package sqlite

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/jqwez/caselore/database/repositories"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	Database       *sql.DB
	UserRepository *SQLiteUserRepository
}

func NewSQLiteRepository() *SQLiteRepository {
	r := &SQLiteRepository{}
	r.Database = NewSQLiteConnection("data.dat")
	r.UserRepository = NewUserRepository(r.Database)
	return r
}

func (sr *SQLiteRepository) GetUserRepository() repositories.UserRepository {
	return sr.UserRepository
}

func (sr *SQLiteRepository) GetDatabase() *sql.DB {
	return sr.Database
}

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
