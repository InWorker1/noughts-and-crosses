package game

import (
	"game/internal/domain/game"

	"github.com/google/uuid"
)

func GetJson(game game.Game, winner int) JsonReqAiGame {
	return JsonReqAiGame{Board: game.Board, Winner: winner}
}

func GetDomain(req JsonReqAiGame, id uuid.UUID) game.Game {
	return game.Game{ID: id, Board: req.Board}
}
