package game

import (
	"encoding/json"
	"game/internal/domain/game"
	"game/internal/domain/user"
	"net/http"

	"github.com/google/uuid"
)

type GameHandler struct {
	gameService game.GameService
	userService user.UserService
}

func NewGameHandler(gs game.GameService, us user.UserService) *GameHandler {
	return &GameHandler{gameService: gs, userService: us}
}

func (h *GameHandler) Move(w http.ResponseWriter, r *http.Request) {
	Id := r.PathValue("uuid")

	gameId, err := uuid.Parse(Id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req JsonReqAiGame
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	game := GetDomain(req, gameId)
	game.IsOnline = false

	if flag, err := h.gameService.ValidateBoard(&game); err != nil { // ошибка плохо
		http.Error(w, "Invalid game", http.StatusBadRequest)
		return
	} else if !flag { // флаг на то есть вообще эта игра или ее нет, если нет то мы ее создали
		w.Header().Set("Content-Type", "application/json")
		req = GetJson(game, 0)
		w.WriteHeader(http.StatusCreated)     // хэд статуса
		err := json.NewEncoder(w).Encode(req) // отправляем данные пользователю(клиенту)
		if err != nil {
			http.Error(w, "sorry.. ", 500)
			return
		}
		return
	}

	userId, err := uuid.Parse(r.Header.Get("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	game = h.gameService.SetRolesWithAi(game, userId) // будет только в случае первого хода
	game, err = h.gameService.GetNextMove(game)
	if err != nil {
		http.Error(w, "sorry..", 500)
		return
	}
	req = GetJson(game, h.gameService.GameOver(game))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)         // хэд статуса
	err = json.NewEncoder(w).Encode(req) // отправляем данные пользователю(клиенту)
	if err != nil {
		http.Error(w, "sorry.. ", 500)
		return
	}

}
