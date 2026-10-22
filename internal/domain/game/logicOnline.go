package game

import (
	"errors"
	"game/internal/domain/domainErrors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *gameService) GetOnlineGame(id, user uuid.UUID) (Game, error) { // игра из бд, только для её участников
	game, err := s.repo.Get(id)
	if errors.Is(err, pgx.ErrNoRows) {
		return Game{}, domainErrors.ErrGameNotFound
	} else if err != nil {
		return Game{}, err
	}
	if user != game.XPlayer && user != game.OPlayer {
		return Game{}, domainErrors.ErrNotYourGame
	}
	game.IsOnline = true
	return game, nil
}

func (s *gameService) ValidateOnlineMove(game *Game, user uuid.UUID) error { // проверяет ход и дополняет game данными из бд
	legacyGame, err := s.GetOnlineGame(game.ID, user)
	if err != nil {
		return err
	}
	if legacyGame.Winner != uuid.Nil || legacyGame.Draw {
		return domainErrors.ErrGameFinished
	}
	if legacyGame.Waiting {
		return domainErrors.ErrGameWaiting
	}
	if legacyGame.MovePlayer != user {
		return domainErrors.ErrNotYourTurn
	}
	symbol := X
	if user == legacyGame.OPlayer {
		symbol = O
	}
	if !checkMove(legacyGame.Board.Grid, game.Board.Grid, symbol) {
		return domainErrors.ErrInvalidBoard
	}
	legacyGame.Board = game.Board
	*game = legacyGame
	return nil
}

func (s *gameService) SaveOnlineMove(game Game, user uuid.UUID, result int) (Game, error) { // result из GameOver
	switch result {
	case X:
		game.Winner = game.XPlayer
	case O:
		game.Winner = game.OPlayer
	case Draw:
		game.Draw = true
	}
	if user == game.XPlayer { // передаём ход сопернику
		game.MovePlayer = game.OPlayer
	} else {
		game.MovePlayer = game.XPlayer
	}
	return game, s.repo.SaveOnlineGame(game, user)
}

// ровно одна клетка: была пустой, стала symbol
func checkMove(oldGrid, newGrid [][]int, symbol int) bool {
	if !validGrid(newGrid) {
		return false
	}
	changed := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if newGrid[i][j] == oldGrid[i][j] {
				continue
			}
			if oldGrid[i][j] != Void || newGrid[i][j] != symbol {
				return false
			}
			changed++
		}
	}
	return changed == 1
}

func validGrid(grid [][]int) bool { // 3x3 и только X, O, Void
	if len(grid) != 3 {
		return false
	}
	for _, row := range grid {
		if len(row) != 3 {
			return false
		}
		for _, cell := range row {
			if cell != X && cell != O && cell != Void {
				return false
			}
		}
	}
	return true
}
