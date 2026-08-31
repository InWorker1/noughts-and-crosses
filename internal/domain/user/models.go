package user

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Login string    `json:"login" db:"login"`
	Pass  string    `json:"pass" db:"pass"`
}
