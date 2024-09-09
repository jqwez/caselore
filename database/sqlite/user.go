package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jqwez/caselore/model"
)

type SQLiteUserRepository struct {
	conn *sql.DB
}

func NewUserRepository(c *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{
		conn: c,
	}
}

func (ur *SQLiteUserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users(id, username, password, role)
		VALUES (?, ?, ?, ?)
		`
	user, err := user.Validate()
	user.HashPassword()
	if err != nil {
		return err
	}
	_, err = ur.conn.Exec(query, user.ID.String(), user.Username, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (ur *SQLiteUserRepository) UpdateUsername(u *model.User, username string) error {
	query := "UPDATE users set username=? WHERE id=?"
	u.Username = username
	_, err := u.Validate()
	if err != nil {
		return err
	}
	_, err = ur.conn.Exec(query, username, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *SQLiteUserRepository) GetByID(id uuid.UUID) (*model.User, error) {
	query := "SELECT * FROM users WHERE id=?"
	row := ur.conn.QueryRow(query, id.String())
	if err := row.Err(); err != nil {
		return &model.User{}, err
	}
	u := &model.User{}
	var idString string
	err := row.Scan(&idString, &u.Username, &u.Password, &u.Role)
	if err != nil {
		return &model.User{}, err
	}
	err = u.SetIDString(idString)
	if err != nil {
		return &model.User{}, err
	}
	return u, nil
}

func (ur *SQLiteUserRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := ur.conn.Exec(query, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (ur *SQLiteUserRepository) CreateTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username TEXT UNIQUE,
			password TEXT,
			role	TEXT
	);
	`
	_, err := ur.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
