package model

import "github.com/google/uuid"

type DocumentPermission struct {
	ID             uuid.UUID      `json:"id"`
	DocumentID     uuid.UUID      `json:"document_id"`
	UserPermitted  uuid.UUID      `json:"user_id"`
	TypePermission TypePermission `json:"type_permission"`
}

type TypePermission int

const (
	NoPermission TypePermission = iota
	ViewOnly
	ViewShare
	EditOnly
	EditShare
	PublicView
)

func NewDocumentPermission(doc *Document, user *User, typePermission TypePermission) *DocumentPermission {
	dp := &DocumentPermission{}
	dp.DocumentID = doc.GetID()
	dp.UserPermitted = user.GetID()
	dp.SetPermission(typePermission)
	return dp
}

func (dp *DocumentPermission) GetID() uuid.UUID {
	return dp.ID
}

func (dp *DocumentPermission) SetID(_id uuid.UUID) error {
	dp.ID = _id
	return nil
}

func (dp *DocumentPermission) SetIDString(idString string) error {
	return ModelSetUUIDString(dp, idString)
}

func (dp *DocumentPermission) SetPermission(tp TypePermission) {
	if tp < NoPermission || tp > PublicView {
		dp.TypePermission = NoPermission
	} else {
		dp.TypePermission = tp
	}
}
