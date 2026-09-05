package game

import "github.com/google/uuid"

const (
	X       = 1
	O       = -1
	Void    = 0
	Draw    = -3
	BotUUID = "00000000-0000-0000-0000-000000000001"
)

type Game struct {
	ID         uuid.UUID
	Board      GameField
	Waiting    bool
	MovePlayer uuid.UUID
	Winner     uuid.UUID
	Draw       bool
	XPlayer    uuid.UUID
	OPlayer    uuid.UUID
	IsOnline   bool // для определения онл игры или нет
}

type GameField struct {
	Grid [][]int
}
