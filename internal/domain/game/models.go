package game

import "github.com/google/uuid"

const (
	X    = 1
	O    = -1
	Void = 0
	Draw = -3
)

type Game struct {
	ID         uuid.UUID
	Board      GameField
	Waiting    bool
	MovePlayer uuid.UUID
	Winner     uuid.UUID
	Draw       bool
	IsOnline   bool // для определения онл игры или нет
}

type GameField struct {
	Grid [][]int
}
