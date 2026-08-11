package sync_map

import (
	"game/internal/domain/game"

	"github.com/google/uuid"
)

type Storage struct {
	Value game.GameField
	Key   uuid.UUID
}

func GameIntoDS(game game.Game) Storage {
	return Storage{Key: game.ID, Value: game.Board}
}

func DSIntoGame(store Storage) game.Game {
	return game.Game{ID: store.Key, Board: store.Value}
}
