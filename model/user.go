package model

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"regexp"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string
	Role     string `json:"role"`
}

func (u *User) CreateUUID() error {
	_, err := uuid.Parse(u.ID.String())
	if err == nil && u.ID != uuid.Nil {
		return err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (u *User) SetIDString(idString string) error {
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (u *User) Validate() (*User, error) {
	if err := u.CreateUUID(); err != nil {
		return &User{}, err
	}
	if err := u.isValidUsername(); err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) isValidUsername() error {
	return validation.Validate(u.Username,
		validation.Required,
		validation.Length(3, 20),
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9_-]+$")).Error("invalid characters"),
		validation.By(checkNoSpaces),
	)
}

func checkNoSpaces(value interface{}) error {
	s, _ := value.(string)
	if len(s) != len(s) {
		return fmt.Errorf("username should not contain spaces")
	}
	return nil
}

func (u *User) isValidEmail() error {
	return validation.Validate(u.Email, is.Email)
}
