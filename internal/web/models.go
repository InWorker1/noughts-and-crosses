package web

import (
	"game/internal/domain"

	"github.com/google/uuid"
)

type JsonRequest struct {
	Id     uuid.UUID        `json:"id"`
	Board  domain.GameField `json:"game"`
	Winner int              `json:"winner,omitempty"`
}
