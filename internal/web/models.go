package web

import (
	"game/internal/domain"
)

// @von (Petya Mishin)

type JsonRequest struct {
	Board  domain.GameField `json:"game,omitempty"`
	Winner int              `json:"winner,omitempty"`
}
