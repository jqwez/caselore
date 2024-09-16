package model

import (
	"github.com/google/uuid"
)

type IModel interface {
	GetID() uuid.UUID
	SetID(uuid.UUID) error
	SetIDString(string) error
}

func ModelCreateUUID(m IModel) error {
	currentID, err := uuid.Parse(m.GetID().String())
	if err == nil && currentID != uuid.Nil {
		return nil
	}
	newID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return m.SetID(newID)
}

func ModelSetUUIDString(m IModel, idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}
	return m.SetID(id)
}
