package domain

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	UUID  uuid.UUID
	Email string
	Name  string
}

func NewUser(email, name string) (User, error) {

	if email == "" {
		return User{}, errors.New("email must not be empty")
	}

	if name == "" {
		return User{}, errors.New("name must not be empty")
	}

	user := User{
		Email: email,
		Name:  name,
	}

	return user, nil
}

func NewIdentifiedUser(id uuid.UUID, email, name string) (User, error) {
	// Validate UUID
	if id == uuid.Nil {
		return User{}, errors.New("UUID must not be null or empty")
	}

	// Validate email
	if email == "" {
		return User{}, errors.New("email must not be empty")
	}

	// Validate name
	if name == "" {
		return User{}, errors.New("name must not be empty")
	}

	user := User{
		UUID:  id,
		Email: email,
		Name:  name,
	}

	return user, nil
}
