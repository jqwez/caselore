package model

import (
	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID `json:"id"`
	OwnerID     uuid.UUID `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
}

func (d *Document) CreateUUID() error {
	return ModelCreateUUID(d)
}

func (d *Document) GetID() uuid.UUID {
	return d.ID
}

func (d *Document) SetID(_id uuid.UUID) error {
	d.ID = _id
	return nil
}

func (d *Document) SetIDString(idString string) error {
	return ModelSetUUIDString(d, idString)
}
