package database

import (
	"database/sql"
	"os"

	"github.com/jqwez/caselore/database/repositories"
	"github.com/jqwez/caselore/database/sqlite"
)

type Repository struct {
	Database       *sql.DB
	UserRepository repositories.UserRepository
}

type DatabaseType interface {
	GetUserRepository() repositories.UserRepository
	GetDatabase() *sql.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) MustSetFromEnv() *Repository {
	dbType := MustGetRepositoryFromEnv()
	return r.SetFromDBType(dbType)
}

func (r *Repository) SetFromDBType(dbType DatabaseType) *Repository {
	r.Database = dbType.GetDatabase()
	r.UserRepository = dbType.GetUserRepository()
	return r
}

func MustGetRepositoryFromEnv() DatabaseType {
	databaseEnv := os.Getenv("DATABASE_IS")
	if databaseEnv == "" || databaseEnv == "sqlite" {
		dbName := os.Getenv("DATABASE_NAME")
		if dbName == "" {
			dbName = "app.dat"
		}
		return sqlite.NewSQLiteRepository()
	}
	return sqlite.NewSQLiteRepository()
}
