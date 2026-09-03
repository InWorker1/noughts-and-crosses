package user

import "github.com/google/uuid"

type User struct {
	Id      uuid.UUID `json:"id" db:"id"`
	Login   string    `json:"login" db:"login"`
	Pass    string    `json:"pass" db:"password"`
	GameNow uuid.UUID `json:"game_now" db:"game_now"`
	Role    string    `json:"role" db:"role"` // p - person; a - adm; (будущее)
	//Games   []uuid.UUID `json:"games" db:"games"`
}
