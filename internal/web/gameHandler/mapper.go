package game

import (
	"game/internal/domain/game"

	"github.com/google/uuid"
)

func GetJson(game game.Game, winner int) JsonReqAiGame {
	return JsonReqAiGame{Board: game.Board, Winner: winner}
}

func GetOnlineJson(game game.Game, winner int, user uuid.UUID) JsonOnlineGame {
	return JsonOnlineGame{Board: game.Board, Winner: winner, Waiting: game.Waiting, YourTurn: game.MovePlayer == user}
}

func GetDomain(req JsonReqAiGame, id uuid.UUID) game.Game {
	return game.Game{ID: id, Board: req.Board}
}
