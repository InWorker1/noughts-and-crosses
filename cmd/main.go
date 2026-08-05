package main

import (
	"fmt"
	"game/internal/domain"
	"game/internal/repository"
	"game/internal/web"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	gameStorage := repository.NewGameStorage()
	gameRepo := repository.NewGameRepo(gameStorage)
	gameService := domain.NewGameService(gameRepo)
	gameHandler := web.NewGameHandler(gameService)

	mux.HandleFunc("/game/{uuid}", gameHandler.Move)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Errorf("Error LAS: %v", err)
	}
}
