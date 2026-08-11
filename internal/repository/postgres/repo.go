package postgres

import (
	"context"
	"game/internal/domain"

	"github.com/google/uuid"
)

type gameRepository struct {
	data *gameDataBase
}

func NewGameRepository(store *gameDataBase) domain.GameRepository {
	return &gameRepository{data: store}
}

//Save(game Game) error
//Get(id uuid.UUID) (Game, error)
//Delete(id uuid.UUID) error

func (repo *gameRepository) Save(game domain.Game) error {
	unit := DomainIntoRepo(game)
	_, err := repo.data.db.Exec(context.Background(), "INSERT INTO games(id, grid) VALUES ($1, $2)", unit.id, unit.grid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *gameRepository) Get(id uuid.UUID) (domain.Game, error) {
	row := repo.data.db.QueryRow(context.Background(), "SELECT * FROM games WHERE id = $1", id)
	var unit sqlStore
	if err := row.Scan(&unit.id, &unit.grid); err != nil {
		return domain.Game{}, err
	}
	return RepoIntoDomain(unit), nil
}

func (repo *gameRepository) Delete(id uuid.UUID) error {
	_, err := repo.data.db.Exec(context.Background(), "DELETE FROM games WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
