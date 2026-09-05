package di

import (
	"context"
	"fmt"
	"game/internal/domain/auth"
	"game/internal/domain/game"
	"game/internal/domain/user"
	"game/internal/repository/postgres"
	"game/internal/web/authHandler"
	game2 "game/internal/web/gameHandler"
	"game/internal/web/middleware"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	return http.NewServeMux()
}

func NewDSN() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	user, pass, host, port, dbname := os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DB_NAME")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			postgres.NewGameDataBase,
			postgres.NewGameRepository,
			postgres.NewUserRepo,
			game.NewGameService,
			game2.NewGameHandler,
			user.NewUserService,
			auth.NewAuthService,
			authHandler.NewAuthHandler,
			middleware.NewUserAuthenticator,
			NewMux,
			NewDSN,
		),
		fx.Invoke(
			RegisterRout,
			ListenServer,
		),
	)
}

func RegisterRout(mux *http.ServeMux, gameHand *game2.GameHandler, authHand *authHandler.AuthHandler, middleHand *middleware.UserAuthenticator) {
	//mux.HandleFunc("/user/info", )
	mux.HandleFunc("POST /game/newgame", middleHand.Authorization(gameHand.CreateGame))
	//mux.HandleFunc("POST /game/online/{uuid}", middleHand.Authorization())
	//mux.HandleFunc("/user/info", )
	//mux.HandleFunc("")
	mux.HandleFunc("POST /game/ai/{uuid}", middleHand.Authorization(gameHand.Move))
	mux.HandleFunc("POST /auth/reg", authHand.Register)
	mux.HandleFunc("GET /auth/log", authHand.Login)
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
