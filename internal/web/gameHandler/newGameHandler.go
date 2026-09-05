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
	}
	request.GameMode = strings.TrimSpace(request.GameMode)
	request.GameMode = strings.ToLower(request.GameMode)
	switch request.GameMode {
	//case "online":
	//	id := uuid.New()
	case "offline":
		id := uuid.New()
		http.Redirect(w, r, "/game/ai/"+id.String(), http.StatusTemporaryRedirect) // 307
	default:
		http.Error(w, "invalid game mode", http.StatusBadRequest) // 400
	}
}
