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

//Save(game Game) error
//Get(id uuid.UUID) (Game, error)
//Delete(id uuid.UUID) error

func (repo *gameRepository) Save(game game.Game) error {
	unit := DomainIntoRepo(game)
	query := "INSERT INTO games(id, grid) VALUES ($1, $2)"
	_, err := repo.data.db.Exec(context.Background(), query, unit.id, unit.grid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *gameRepository) Get(id uuid.UUID) (game.Game, error) {
	query := "SELECT * FROM games WHERE id = $1"
	row := repo.data.db.QueryRow(context.Background(), query, id)
	var unit sqlStore
	if err := row.Scan(&unit.id, &unit.grid); err != nil {
		return game.Game{}, err
	}
	return RepoIntoDomain(unit), nil
}

func (repo *gameRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM games WHERE id = $1"
	_, err := repo.data.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
