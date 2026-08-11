package postgres

import (
	"game/internal/domain/game"

	"github.com/google/uuid"
)

type sqlStore struct {
	id   uuid.UUID `db:"id"`
	grid [][]int   `db:"grid"`
}

func DomainIntoRepo(game game.Game) sqlStore {
	return sqlStore{id: game.ID, grid: game.Board.Grid}
}

func RepoIntoDomain(store sqlStore) game.Game {
	return game.Game{ID: store.id, Board: game.GameField{Grid: store.grid}}
}
