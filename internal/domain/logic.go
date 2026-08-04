package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type gameService struct {
	repo GameRepository
}

func NewGameService(r GameRepository) GameService {
	return &gameService{
		repo: r,
	}
}

func (s *gameService) GetNextMove(game Game) Game {
	var bestVar struct {
		g     Game
		score int
	}
	bestVar.score = -1000

	PC := whoPC(game)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board.Grid[i][j] == Void {
				gameClone := game
				gameClone.Board.Grid[i][j] = PC
				result := s.minimax(gameClone, -PC, 1, PC)
				if result > bestVar.score {
					bestVar.g = gameClone
					bestVar.score = result
				}
			}
		}
	}
	s.repo.Save(bestVar.g)
	return bestVar.g
}

func (s *gameService) minimax(game Game, player, dep, symPC int) int {
	winner := s.GameOver(game)
	if winner != 0 {
		switch winner {
		case symPC:
			return 10 - dep
		default:
			return -10 + dep
		}
	} else if winner == 0 && gridIsFull(game) {
		return 0
	}
	var bestScore int
	if player == symPC {
		bestScore = -1000 // Для ПК ищем максимум, стартуем со дна
	} else {
		bestScore = 1000 // Для человека ищем минимум, стартуем с потолка
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board.Grid[i][j] == Void {
				gameClone := game
				gameClone.Board.Grid[i][j] = player
				result := s.minimax(gameClone, -player, dep+1, symPC)
				if result > bestScore && player == symPC {
					bestScore = result
				} else if result < bestScore && player != symPC {
					bestScore = result
				}
			}
		}
	}
	return bestScore
}

func (s *gameService) ValidateBoard(game Game) (bool, error) { // false когда новая игра или ошибка
	legacyGame, err := s.repo.Get(game.ID)
	if err != nil {
		id, _ := uuid.NewUUID()
		s.repo.Save(Game{ID: id,
			Board: GameField{
				Grid: [][]int{
					{Void, Void, Void},
					{Void, Void, Void},
					{Void, Void, Void},
				},
			},
		})
		return false, nil
	}
	counter := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			symbolNew := game.Board.Grid[i][j]
			symbolLeg := legacyGame.Board.Grid[i][j]

			if symbolNew != symbolLeg && symbolLeg != Void {
				return false, fmt.Errorf("Invalid NewBoard")
			} else if symbolNew != symbolLeg && symbolLeg == Void {
				counter++
			} else if counter > 1 {
				return false, fmt.Errorf("Invalid NewBoard")
			}
		}
	}
	return true, nil
}

func (s *gameService) GameOver(game Game) int { // выдает победителя или 0 (null тип)
	hor := 0  // горизонтально
	ver := 0  // вертикально
	diaL := 0 // диагональ слева
	diaR := 0 // диагональ справа
	for i := 0; i < 3; i++ {
		diaL += game.Board.Grid[i][i]
		diaR += game.Board.Grid[i][2-i]
	}
	if diaL == 3 || diaR == 3 {
		return X
	} else if diaL == -3 || diaR == -3 {
		return O
	}

	for i := 0; i < 3; i++ {
		hor = 0
		ver = 0
		for j := 0; j < 3; j++ {
			hor += game.Board.Grid[i][j]
			ver += game.Board.Grid[j][i]
		}
		if hor == 3 || ver == 3 {
			return X
		} else if hor == -3 || ver == -3 {
			return O
		}
	}
	return 0
}

func gridIsFull(game Game) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board.Grid[i][j] == Void {
				return false
			}
		}
	}
	return true
}

func whoPC(game Game) int {
	counterX, counterO := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board.Grid[i][j] == X {
				counterX++
			} else if game.Board.Grid[i][j] == O {
				counterO++
			}
		}
	}
	if counterX == counterO {
		return X
	}
	return O
}
