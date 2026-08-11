package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type gameDataBase struct {
	db *pgxpool.Pool
}

func NewGameDataBase(dsn string) *gameDataBase {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	return &gameDataBase{db: pool}
}

//postgres://user:password@localhost:5432/dbname?sslmode=off
