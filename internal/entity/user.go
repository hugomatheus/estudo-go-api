package entity

import (
	"errors"

	"github.com/hugomatheus/go-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserErrIDIsRequired       = errors.New("ID is required")
	UserErrInvalidID          = errors.New("Invalid id")
	UserErrNameIsRequired     = errors.New("Name is required")
	UserErrEmailIsRequired    = errors.New("Email is required")
	UserErrPasswordIsRequired = errors.New("Password is required")
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return user, nil
}

func (u *User) Validate() error {
	if u.ID.String() == "" {
		return UserErrIDIsRequired
	}
	if _, err := entity.ParseID(u.ID.String()); err != nil {
		return UserErrInvalidID
	}
	if u.Name == "" {
		return UserErrNameIsRequired
	}

	if u.Email == "" {
		return UserErrEmailIsRequired
	}

	if u.Password == "" {
		return UserErrPasswordIsRequired
	}
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
