package sync_map

import "sync"

type GameStorage struct {
	data sync.Map
}

func NewGameStorage() *GameStorage {
	return &GameStorage{}
}
