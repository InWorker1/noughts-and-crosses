package di

import (
	"context"
	"fmt"
	"game/internal/domain"
	"game/internal/repository/sync-map"
	"game/internal/web"
	"net/http"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	return http.NewServeMux()
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			sync_map.NewGameStorage,
			sync_map.NewGameRepo,
			domain.NewGameService,
			web.NewGameHandler,
			NewMux,
		),
		fx.Invoke(
			RegisterRout,
			ListenServer,
		),
	)
}

func RegisterRout(mux *http.ServeMux, gameHand *web.GameHandler) {
	mux.HandleFunc("/game/{uuid}", gameHand.Move)
}

func ListenServer(lc fx.Lifecycle, mux *http.ServeMux) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := server.ListenAndServe()
				if err != nil {
					fmt.Errorf("Error ListenServer: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Server Stop")
			return server.Shutdown(context.Background())
		},
	})
}
