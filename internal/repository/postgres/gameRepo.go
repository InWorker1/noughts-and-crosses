package postgres

import (
	"context"
	"game/internal/domain/game"

	"github.com/google/uuid"
)

type gameRepository struct {
	data *dataBase
}

func NewGameRepository(store *dataBase) game.GameRepository {
	return &gameRepository{data: store}
}

func (repo *gameRepository) Create(game game.Game) error {
	query := `INSERT INTO games (id, grid) VALUES ($1, $2)`
	_, err := repo.data.db.Exec(context.Background(), query, game.ID, game.Board.Grid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *gameRepository) Save(game game.Game) error {
	unit := DomainIntoRepo(game)
	query := `UPDATE games SET grid = $1 WHERE id = $2`
	_, err := repo.data.db.Exec(context.Background(), query, unit.grid, unit.id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *gameRepository) Get(id uuid.UUID) (game.Game, error) {
	query := `SELECT * FROM games WHERE id = $1`
	row := repo.data.db.QueryRow(context.Background(), query, id)
	var unit sqlStore
	if err := row.Scan(&unit.id, &unit.grid); err != nil {
		return game.Game{}, err
	}
	return RepoIntoDomain(unit), nil
}

//func (repo *gameRepository) Delete(id uuid.UUID) error {
//	query := "DELETE FROM games WHERE id = $1"
//	_, err := repo.data.db.Exec(context.Background(), query, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
