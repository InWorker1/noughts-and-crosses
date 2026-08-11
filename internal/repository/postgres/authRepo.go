package postgres

import (
	"context"
	"database/sql"
	"errors"
	"game/internal/domain/auth"
	"game/internal/domain/domainErrors"

	"github.com/google/uuid"
)

type authRepo struct {
	data *dataBase
}

func NewAuthRepo(data *dataBase) auth.AuthRepository {
	return &authRepo{data: data}
}

func (a *authRepo) Create(request auth.SignUpRequest) error {
	id := uuid.New()
	query := "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)"
	_, err := a.data.db.Exec(context.Background(), query, id, request.Login, request.Pass)
	if err != nil {
		return err
	}
	return nil
}

func (a *authRepo) GetByUsername(username string) (string, error) {
	var hash string
	query := "SELECT password FROM users WHERE username = $1"
	row := a.data.db.QueryRow(context.Background(), query, username)
	if err := row.Scan(&hash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domainErrors.ErrPersonNotFound
		}
		return "", err
	}
	return hash, nil
}
