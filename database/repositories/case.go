package repositories

import (
	"github.com/google/uuid"
	"github.com/jqwez/caselore/model"
)

type CaseRepository interface {
	Create(*model.Case) error
	UpdateUsername(*model.Case, string) error
	GetByID(id uuid.UUID) (*model.Case, error)
	Delete(id uuid.UUID) error
	CreateTable() error
}
