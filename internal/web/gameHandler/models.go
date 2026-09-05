package game

import (
	"game/internal/domain/game"
)

type JsonReqAiGame struct {
	Board  game.GameField `json:"game,omitempty"`
	Winner int            `json:"winner,omitempty"`
}

type JsonReqNewGame struct {
	GameMode string `json:"game_mode"`
}
