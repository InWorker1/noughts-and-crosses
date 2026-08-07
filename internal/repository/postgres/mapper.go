package postgres

import (
	"game/internal/domain"

	"github.com/google/uuid"
)

type sqlStore struct {
	id   uuid.UUID `db:"id"`
	grid [][]int   `db:"grid"`
}

func DomainIntoRepo(game domain.Game) sqlStore {
	return sqlStore{id: game.ID, grid: game.Board.Grid}
}

func RepoIntoDomain(store sqlStore) domain.Game {
	return domain.Game{ID: store.id, Board: domain.GameField{Grid: store.grid}}
}
