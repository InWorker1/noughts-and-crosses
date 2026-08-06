package main

import (
	"game/internal/di"

	"go.uber.org/fx"
)

// @von (Petya Mishin)

func main() {
	fx.New(di.CreateApp()).Run()

	//mux := http.NewServeMux() LEGACY
	//
	//gameStorage := repository.NewGameStorage()
	//gameRepo := repository.NewGameRepo(gameStorage)
	//gameService := domain.NewGameService(gameRepo)
	//gameHandler := web.NewGameHandler(gameService)
	//
	//mux.HandleFunc("/game/{uuid}", gameHandler.Move)
	//
	//err := http.ListenAndServe(":8080", mux)
	//if err != nil {
	//	fmt.Errorf("Error LAS: %v", err)
	//}
}
