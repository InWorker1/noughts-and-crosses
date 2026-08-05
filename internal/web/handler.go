package web

import (
	"encoding/json"
	"game/internal/domain"
	"net/http"

	"github.com/google/uuid"
)

type GameHandler struct {
	service domain.GameService
}

func NewGameHandler(s domain.GameService) *GameHandler {
	return &GameHandler{service: s}
}

func (h *GameHandler) Move(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("uuid")

	var req JsonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	gameId, err := uuid.Parse(Id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	game := GetDomain(req, gameId)

	if flag, err := h.service.ValidateBoard(&game); err != nil { // ошибка плохо
		http.Error(w, "Invalid game", http.StatusBadRequest)
		return
	} else if !flag { // флаг на то есть вообще эта игра или ее нет, если нет то мы ее создали
		w.Header().Set("Content-Type", "application/json")
		req = GetJson(game, 0)
		w.WriteHeader(http.StatusOK)          // хэд статуса
		err := json.NewEncoder(w).Encode(req) // отправляем данные пользователю(клиенту)
		if err != nil {
			http.Error(w, "sorry.. ", 500)
			return
		}
		return
	}

	game = h.service.GetNextMove(game)
	winner := h.service.GameOver(game)

	req = GetJson(game, winner)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)         // хэд статуса
	err = json.NewEncoder(w).Encode(req) // отправляем данные пользователю(клиенту)
	if err != nil {
		http.Error(w, "sorry.. ", 500)
		return
	}

}
