package game

import (
	"github.com/google/uuid"
)

type GameService interface {
	GetNextMove(game Game) (Game, error)
	ValidateBoard(game *Game) (bool, error)
	GameOver(game Game) int

	SetRolesWithAi(game Game, user uuid.UUID) Game
}

type GameRepository interface {
	Create(game Game) error
	SaveAiGame(game Game) error
	Get(id uuid.UUID) (Game, error)
	//SaveOnlineGame(game Game) error

	SetRoles(game Game) error
	//Delete(id uuid.UUID) error
}
