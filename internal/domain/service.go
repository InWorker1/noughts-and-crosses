package domain

import "github.com/google/uuid"

type GameService interface {
	GetNextMove(game Game) Game
	ValidateBoard(game *Game) (bool, error)
	GameOver(game Game) int
}

type GameRepository interface {
	Save(game Game) error
	Get(id uuid.UUID) (Game, error)
	Delete(id uuid.UUID) error
}
