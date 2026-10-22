package game

import (
	"encoding/json"
	"errors"
	"game/internal/domain/domainErrors"
	"game/internal/domain/game"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// POST без доски (в т.ч. первый запрос после редиректа с {"game_mode":"online"}) отдаёт состояние игры, с доской делает ход
func (h *GameHandler) OnlineGameStart(w http.ResponseWriter, r *http.Request) {
	uuidGame, err := uuid.Parse(r.PathValue("uuid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idPerson, err := uuid.Parse(r.Header.Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req JsonOnlineGame
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && !errors.Is(err, io.EOF) {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	var onlineGame game.Game
	if req.Board.Grid == nil {
		onlineGame, err = h.gameService.GetOnlineGame(uuidGame, idPerson)
	} else {
		onlineGame = game.Game{ID: uuidGame, Board: req.Board}
		err = h.gameService.ValidateOnlineMove(&onlineGame, idPerson)
		if err == nil {
			onlineGame, err = h.gameService.SaveOnlineMove(onlineGame, idPerson, h.gameService.GameOver(onlineGame))
		}
	}
	if err != nil {
		writeOnlineErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(GetOnlineJson(onlineGame, h.gameService.GameOver(onlineGame), idPerson))
	if err != nil {
		http.Error(w, "sorry.. ", http.StatusInternalServerError)
	}
}

func writeOnlineErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domainErrors.ErrGameNotFound):
		http.Error(w, err.Error(), http.StatusNotFound) // 404
	case errors.Is(err, domainErrors.ErrNotYourGame):
		http.Error(w, err.Error(), http.StatusForbidden) // 403
	case errors.Is(err, domainErrors.ErrInvalidBoard):
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
	case errors.Is(err, domainErrors.ErrGameWaiting),
		errors.Is(err, domainErrors.ErrGameFinished),
		errors.Is(err, domainErrors.ErrNotYourTurn):
		http.Error(w, err.Error(), http.StatusConflict) // 409
	default:
		http.Error(w, "sorry..", http.StatusInternalServerError)
	}
}
