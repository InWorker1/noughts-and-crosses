package postgres

import (
	"context"
	"game/internal/domain/domainErrors"
	"game/internal/domain/game"

	"github.com/google/uuid"
)

type gameRepository struct {
	data *dataBase
}

func NewGameRepository(store *dataBase) game.GameRepository { // тут возвращался interface, я поменял на структуру
	return &gameRepository{data: store}
}

func (repo *gameRepository) Create(game game.Game) error {
	query := `INSERT INTO games (id, grid, waiting) VALUES ($1, $2, $3)`
	_, err := repo.data.db.Exec(context.Background(), query, game.ID, game.Board.Grid, game.Waiting)
	if err != nil {
		return err
	}
	return nil
}

func (repo *gameRepository) SaveAiGame(game game.Game) error {
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
	if err := row.Scan(&unit.id, &unit.grid, &unit.waiting, &unit.movePlayer, &unit.winner, &unit.draw, &unit.xPlayer, &unit.oPlayer); err != nil {
		return game.Game{}, err
	}
	return RepoIntoDomain(unit), nil
}

func (repo *gameRepository) SetRoles(game game.Game) error {
	query := `UPDATE games SET x_player = $1, o_player = $2 WHERE id = $3`
	_, err := repo.data.db.Exec(context.Background(), query, game.XPlayer, game.OPlayer, game.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *gameRepository) CreateOnline(game game.Game) error {
	query := `INSERT INTO games (id, grid, waiting, x_player) VALUES ($1, $2, $3, $4)`
	_, err := repo.data.db.Exec(context.Background(), query, game.ID, game.Board.Grid, game.Waiting, game.XPlayer)
	return err
}

// одним запросом находим ожидающую игру и занимаем её, SKIP LOCKED не даст двум игрокам зайти в одну игру
func (repo *gameRepository) JoinWaitingGame(user uuid.UUID) (uuid.UUID, error) {
	query := `UPDATE games SET o_player = $1, waiting = false, move_player = x_player
		WHERE id = (SELECT id FROM games WHERE waiting AND x_player <> $1 LIMIT 1 FOR UPDATE SKIP LOCKED)
		RETURNING id`
	var id uuid.UUID
	err := repo.data.db.QueryRow(context.Background(), query, user).Scan(&id)
	return id, err
}

// сохраняем ход только если ходил тот, чья была очередь, иначе два параллельных запроса одного игрока дадут двойной ход
func (repo *gameRepository) SaveOnlineGame(game game.Game, mover uuid.UUID) error {
	query := `UPDATE games SET grid = $1, move_player = $2, winner = NULLIF($3, '00000000-0000-0000-0000-000000000000'::uuid), draw = $4
		WHERE id = $5 AND move_player = $6`
	tag, err := repo.data.db.Exec(context.Background(), query, game.Board.Grid, game.MovePlayer, game.Winner, game.Draw, game.ID, mover)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domainErrors.ErrNotYourTurn
	}
	return nil
}

//func (repo *gameRepository) Save bool {}

//func (repo *gameRepository) Delete(id uuid.UUID) error {
//	query := "DELETE FROM games WHERE id = $1"
//	_, err := repo.data.db.Exec(context.Background(), query, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
