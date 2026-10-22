package game

import (
	"game/internal/domain/game"
)

type JsonReqAiGame struct {
	Board  game.GameField `json:"game,omitempty"`
	Winner int            `json:"winner,omitempty"`
}

type JsonOnlineGame struct {
	Board    game.GameField `json:"game"`
	Winner   int            `json:"winner"`    // как в AI: 1 X, -1 O, -3 ничья, 0 игра идёт
	Waiting  bool           `json:"waiting"`   // ждём второго игрока
	YourTurn bool           `json:"your_turn"` // ход того, кто сделал запрос
}

type JsonReqNewGame struct {
	GameMode string `json:"game_mode"`
}
