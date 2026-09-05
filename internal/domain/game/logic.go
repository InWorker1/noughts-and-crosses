package game

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type gameService struct {
	repo GameRepository
}

func NewGameService(r GameRepository) GameService {
	return &gameService{
		repo: r,
	}
}

func (s *gameService) GetNextMove(game Game) (Game, error) {
	var filledGrid bool

	var bestVar struct {
		g     Game
		score int
	}
	bestVar.score = -1000

	PC := whoPC(game)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board.Grid[i][j] == Void {
				filledGrid = true
				clone := gameClone(game)
				clone.Board.Grid[i][j] = PC
				result := s.minimax(clone, -PC, 1, PC)
				if result > bestVar.score {
					bestVar.g = clone
					bestVar.score = result
				}
			}
		}
	}
	if !filledGrid {
		return game, nil
	}
	bestVar.g.ID = game.ID // присваиваем ему ID оригинала
	err := s.repo.Save(bestVar.g)
	if err != nil {
		return Game{}, err
	}
	return bestVar.g, nil
}

func (s *gameService) ValidateBoard(game *Game) (bool, error) { // false когда новая игра или ошибка
	legacyGame, err := s.repo.Get(game.ID)
	if errors.Is(err, pgx.ErrNoRows) {
		grid := make([][]int, 3, 3)
		for i := 0; i < 3; i++ {
			grid[i] = make([]int, 3)
			for j := 0; j < 3; j++ {
				grid[i][j] = Void
			}
		}
		game.Board.Grid = grid
		err := s.repo.Create(Game{
			ID: game.ID,
			Board: GameField{
				Grid: grid,
			},
		})
		if err != nil {
			return false, err
		}
		return false, nil
	}
	//if whoPC(*game) == -1 {
	//	return false, fmt.Errorf("Invalid game")
	//} 									я в душе не знаю откуда это появилось и зачем. не разрешает пользователю быть крестиком
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
		//_ = s.repo.Delete(game.ID) // удаление игры из памяти
		return X
	} else if diaL == -3 || diaR == -3 {
		//_ = s.repo.Delete(game.ID) // удаление игры из памяти
		return O
	}
	var counterVoid int
	for i := 0; i < 3; i++ {
		hor = 0
		ver = 0
		for j := 0; j < 3; j++ {
			if game.Board.Grid[i][j] == Void {
				counterVoid++
			}
			hor += game.Board.Grid[i][j]
			ver += game.Board.Grid[j][i]
		}
		if hor == 3 || ver == 3 {
			//_ = s.repo.Delete(game.ID) // удаление игры из памяти
			return X
		} else if hor == -3 || ver == -3 {
			//_ = s.repo.Delete(game.ID) // удаление игры из памяти
			return O
		}
	}
	if counterVoid == 0 {
		return Draw
	}
	return 0
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
				clone := gameClone(game)
				clone.Board.Grid[i][j] = player
				result := s.minimax(clone, -player, dep+1, symPC)
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
	} else if counterO > counterX || counterX-counterO > 1 {
		return -1 // обработка случая, что кто то сходил 2 раза подряд
	}
	return O
}

func gameClone(game Game) Game {
	var clone Game
	clone.ID = uuid.Nil
	clone.Board.Grid = make([][]int, 3)
	for i := 0; i < 3; i++ {
		clone.Board.Grid[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			clone.Board.Grid[i][j] = game.Board.Grid[i][j]
		}
	}
	return clone
}
