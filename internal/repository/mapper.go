package repository

import (
	"game/internal/domain"

	"github.com/google/uuid"
)

type Storage struct {
	Value domain.GameField
	Key   uuid.UUID
}

func GameIntoDS(game domain.Game) Storage {
	return Storage{Key: game.ID, Value: game.Board}
}

func DSIntoGame(store Storage) domain.Game {
	return domain.Game{ID: store.Key, Board: store.Value}
}
