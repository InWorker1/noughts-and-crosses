package game

import (
	"game/internal/domain/game"

	"github.com/google/uuid"
)

func GetJson(game game.Game, winner int) JsonRequest {
	return JsonRequest{Board: game.Board, Winner: winner}
}

func GetDomain(req JsonRequest, id uuid.UUID) game.Game {
	return game.Game{ID: id, Board: req.Board}
}
