package web

import (
	"game/internal/domain"
)

type JsonRequest struct {
	Board  domain.GameField `json:"game,omitempty"`
	Winner int              `json:"winner,omitempty"`
}
