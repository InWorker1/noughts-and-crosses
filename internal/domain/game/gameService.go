package game

import (
	"github.com/google/uuid"
)

type GameService interface {
	GetNextMove(game Game) (Game, error)
	ValidateBoard(game *Game) (bool, error)
	GameOver(game Game) int

	SetRolesWithAi(game Game, user uuid.UUID) Game
	FindOnlineGame(user uuid.UUID) (uuid.UUID, error)
	GetOnlineGame(id, user uuid.UUID) (Game, error)
	ValidateOnlineMove(game *Game, user uuid.UUID) error
	SaveOnlineMove(game Game, user uuid.UUID, result int) (Game, error)
	//SetRolesWithRP(game Game, ...) // тут реализовать рандом с раздачей ролей сразу. придумать как получить uuid сразу двух пользователей
}

type GameRepository interface {
	Create(game Game) error
	SaveAiGame(game Game) error
	Get(id uuid.UUID) (Game, error)
	//SaveOnlineGame(game Game) error

	SetRoles(game Game) error
	CreateOnline(game Game) error
	JoinWaitingGame(user uuid.UUID) (uuid.UUID, error)
	SaveOnlineGame(game Game, mover uuid.UUID) error
	//Delete(id uuid.UUID) error
}
