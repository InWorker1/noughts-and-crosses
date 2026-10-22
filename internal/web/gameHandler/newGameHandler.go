package game

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var request JsonReqNewGame
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	request.GameMode = strings.TrimSpace(request.GameMode)
	request.GameMode = strings.ToLower(request.GameMode)
	switch request.GameMode {
	case "offline":
		id := uuid.New()
		http.Redirect(w, r, "/game/ai/"+id.String(), http.StatusTemporaryRedirect) // 307
	case "online":
		userId, err := uuid.Parse(r.Header.Get("id"))
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		id, err := h.gameService.FindOnlineGame(userId) // зашли в ожидающую игру или создали новую с waiting
		if err != nil {
			http.Error(w, "sorry..", http.StatusInternalServerError)
			return
		}
		// TODO: Нужно сделать рандом роли (кто чем ходит), сейчас ожидающий = X, вошедший = O
		http.Redirect(w, r, "/game/online/"+id.String(), http.StatusTemporaryRedirect)
	default:
		http.Error(w, "invalid game mode", http.StatusBadRequest) // 400
	}
}
