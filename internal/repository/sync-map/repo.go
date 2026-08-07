package sync_map

import (
	"fmt"
	"game/internal/domain"

	"github.com/google/uuid"
)

type gameRepository struct { // структура, которая реализует интерфейс domain.GameRepository
	storage *GameStorage
}

func NewGameRepo(s *GameStorage) domain.GameRepository {
	return &gameRepository{storage: s}
}

func (r *gameRepository) Save(game domain.Game) error {
	ds := GameIntoDS(game)
	r.storage.data.Store(ds.Key, ds)
	return nil
}

func (r *gameRepository) Get(id uuid.UUID) (domain.Game, error) {
	val, ok := r.storage.data.Load(id)
	if !ok {
		return domain.Game{}, fmt.Errorf("failed game with ID %s not found", id)
	}

	ds, ok := val.(Storage)
	if !ok {
		return domain.Game{}, fmt.Errorf("failed to cast data from store")
	}
	return DSIntoGame(ds), nil
}

func (r *gameRepository) Delete(id uuid.UUID) error {
	_, ok := r.storage.data.Load(id)
	if !ok {
		return fmt.Errorf("failed delete game", id)
	}
	r.storage.data.Delete(id)
	return nil
}
