package user

import "github.com/google/uuid"

type UserService interface {
	SaveNewPerson(login, password string) uuid.UUID
	GetPerson(login, password string) (User, error)
}

type UserRepository interface {
	Create(user User) error
	GetByUsername(username string) (User, error)
}
