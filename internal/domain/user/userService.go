package user

import "github.com/google/uuid"

type UserService interface {
	Register(login, password string) uuid.UUID
}

type UserRepository interface {
	Create(user User) error
	GetByUsername(username string) (User, error)
}
