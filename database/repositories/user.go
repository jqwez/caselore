package repositories

import (
	"github.com/google/uuid"
	"github.com/jqwez/caselore/model"
)

type UserRepository interface {
	Create(user *model.User) error
	UpdateUsername(user *model.User, username string) error
	GetByID(id uuid.UUID) (*model.User, error)
	Delete(id uuid.UUID) error
	CreateTable() error
}
