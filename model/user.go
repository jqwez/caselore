package model

import (
	"fmt"
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string
	Role     string `json:"role"`
}

func (u *User) CreateUUID() error {
	return ModelCreateUUID(u)
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) SetID(_id uuid.UUID) error {
	u.ID = _id
	return nil
}

func (u *User) SetIDString(idString string) error {
	return ModelSetUUIDString(u, idString)
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

func (u *User) HashPassword() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hashedPassword)

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
