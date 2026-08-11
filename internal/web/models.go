package web

import (
	"game/internal/domain/game"
)

type JsonRequest struct {
	Board  game.GameField `json:"game,omitempty"`
	Winner int            `json:"winner,omitempty"`
}
