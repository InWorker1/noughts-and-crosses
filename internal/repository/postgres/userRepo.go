package postgres

import (
	"context"
	"database/sql"
	"errors"
	"game/internal/domain/domainErrors"
	"game/internal/domain/user"

	"github.com/jackc/pgx/v5"
)

type userRepo struct {
	data *dataBase
}

func NewAuthRepo(data *dataBase) user.UserRepository {
	return &userRepo{data: data}
}

func (a *userRepo) Create(user user.User) error {
	query := "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)"
	_, err := a.data.db.Exec(context.Background(), query, user.Id, user.Login, user.Pass)
	if err != nil {
		return err
	}
	return nil
}

func (a *userRepo) GetByUsername(username string) (user.User, error) {
	query := "SELECT password FROM users WHERE username = $1"

	rows, err := a.data.db.Query(context.Background(), query, username)
	if err != nil {
		panic(err)
	}

	userLoc, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[user.User])

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, domainErrors.ErrPersonNotFound
		}
		return user.User{}, err
	}
	return userLoc, nil
}
