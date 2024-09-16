package model

import (
	"github.com/google/uuid"
)

type Case struct {
	ID            uuid.UUID
	Idea          string
	Title         string
	Description   string
	ImageLocation string
}

func (c *Case) GetID() uuid.UUID {
	return c.ID
}

func (c *Case) SetID(_id uuid.UUID) error {
	c.ID = _id
	return nil
}

func (c *Case) SetIDString(idString string) error {
	return ModelSetUUIDString(c, idString)
}
