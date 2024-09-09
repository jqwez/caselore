package repo

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jqwez/caselore/model"
	"github.com/jqwez/caselore/repo/sqlite"
)

type UserRepository interface {
	Create(user *model.User) error
	UpdateUsername(user *model.User, username string) error
	GetByID(id uuid.UUID) (*model.User, error)
	Delete(id uuid.UUID) error
	CreateTable() error
}

func NewUserRepository(database string) (UserRepository, *sql.DB) {
	if database == "sqlite" {
		db := sqlite.NewSQLiteConnection("app.dat")
		return sqlite.NewUserRepository(db), db
	}
	if database == "mem_sqlite" {
		db := sqlite.NewSQLiteConnection(":memory:")
		return sqlite.NewUserRepository(db), db
	}
	db := sqlite.NewSQLiteConnection("app.dat")
	return sqlite.NewUserRepository(db), db
}
