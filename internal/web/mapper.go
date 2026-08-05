package web

import (
	"game/internal/domain"

	"github.com/google/uuid"
)

func GetJson(game domain.Game, winner int) JsonRequest {
	return JsonRequest{Board: game.Board, Winner: winner}
}

func GetDomain(req JsonRequest, id uuid.UUID) domain.Game {
	return domain.Game{ID: id, Board: req.Board}
}
