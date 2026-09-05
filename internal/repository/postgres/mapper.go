package postgres

import (
	"game/internal/domain/game"

	"github.com/google/uuid"
)

type sqlStore struct {
	id         uuid.UUID `db:"id"`
	grid       [][]int   `db:"grid"`
	waiting    bool      `db:"waiting"`
	movePlayer uuid.UUID `db:"move_player"`
	winner     uuid.UUID `db:"winner"`
	draw       bool      `db:"draw"`
}

func DomainIntoRepo(game game.Game) sqlStore {
	return sqlStore{id: game.ID,
		grid:       game.Board.Grid,
		waiting:    game.Waiting,
		movePlayer: game.MovePlayer,
		winner:     game.Winner,
		draw:       game.Draw}
}

func RepoIntoDomain(store sqlStore) game.Game {
	return game.Game{ID: store.id,
		Board:      game.GameField{Grid: store.grid},
		Waiting:    store.waiting,
		MovePlayer: store.movePlayer,
		Winner:     store.winner,
		Draw:       store.draw}
}
