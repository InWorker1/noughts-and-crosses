package web

import "game/internal/domain"

func GetJson(game domain.Game, winner int) JsonRequest {
	return JsonRequest{Id: game.ID, Board: game.Board, Winner: winner}
}

func GetDomain(req JsonRequest) domain.Game {
	return domain.Game{ID: req.Id, Board: req.Board}
}
