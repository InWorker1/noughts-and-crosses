package domain

import (
	"testing"

	"github.com/google/uuid"
)

// mockRepository - моковая реализация GameRepository для тестов
type mockRepository struct {
	savedGames map[uuid.UUID]*Game
}

func (m *mockRepository) Save(game Game) error {
	uuid, _ := uuid.NewUUID()
	m.savedGames[uuid] = &game
	return nil
}

func (m *mockRepository) Get(id uuid.UUID) (Game, error) { return *m.savedGames[id], nil }

func (m *mockRepository) Delete(id uuid.UUID) error { return nil }

// createEmptyGame создает пустую игру с пустым полем
func createEmptyGame() Game {
	grid := make([][]int, 3)
	for i := 0; i < 3; i++ {
		grid[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			grid[i][j] = Void
		}
	}
	return Game{
		ID: uuid.New(),
		Board: GameField{
			Grid: grid,
		},
	}
}

// TestMinimax_PCWinningMove тестирует ситуацию, когда PC может выиграть следующим ходом
func TestMinimax_PCWinningMove(t *testing.T) {
	repo := &mockRepository{}
	service := NewGameService(repo)
	game := createEmptyGame()

	// Расставляем так, что у PC есть выигрышный ход
	// X . X
	// . O .
	// . . .
	game.Board.Grid[0][0] = X
	game.Board.Grid[0][2] = X
	game.Board.Grid[1][1] = O

	// PC (X) должен найти выигрышный ход в позиции [0][1]
	result := service.(*gameService).minimax(game, X, 0, X)

	// Выигрышный ход должен дать скор 10 - dep (в данном случае dep=0, но будет больше из-за глубины)
	if result <= 0 {
		t.Errorf("Ожидали положительный скор для выигрышной позиции PC, получили %d", result)
	}
}

// TestMinimax_HumanWinningMove тестирует ситуацию, когда человек может выиграть следующим ходом
func TestMinimax_HumanWinningMove(t *testing.T) {
	repo := &mockRepository{}
	service := NewGameService(repo)
	game := createEmptyGame()

	// Расставляем так, что у человека есть выигрышный ход
	// O . O
	// . X .
	// . . .
	game.Board.Grid[0][0] = O
	game.Board.Grid[0][2] = O
	game.Board.Grid[1][1] = X

	// Человек (O) должен найти выигрышный ход в позиции [0][1]
	result := service.(*gameService).minimax(game, O, 0, X)

	// Выигрышный ход человека должен дать отрицательный скор
	if result >= 0 {
		t.Errorf("Ожидали отрицательный скор для выигрышной позиции человека, получили %d", result)
	}
}

// TestMinimax_DrawPosition тестирует позицию, которая приведет к ничьей
func TestMinimax_DrawPosition(t *testing.T) {
	repo := &mockRepository{}
	service := NewGameService(repo)
	game := createEmptyGame()

	// Заполняем доску так, чтобы получилась ничья
	// X O X
	// X O O
	// O X X
	game.Board.Grid[0][0] = X
	game.Board.Grid[0][1] = O
	game.Board.Grid[0][2] = X
	game.Board.Grid[1][0] = X
	game.Board.Grid[1][1] = O
	game.Board.Grid[1][2] = O
	game.Board.Grid[2][0] = O
	game.Board.Grid[2][1] = X
	game.Board.Grid[2][2] = X

	result := service.(*gameService).minimax(game, X, 0, X)

	// Ничья должна вернуть 0
	if result != 0 {
		t.Errorf("Ожидали скор 0 для ничейной позиции, получили %d", result)
	}
}

// TestMinimax_DepthPenalty тестирует штраф за глубину
func TestMinimax_DepthPenalty(t *testing.T) {
	repo := &mockRepository{}
	service := NewGameService(repo)
	game := createEmptyGame()

	// Простая позиция где PC может выиграть, но чем дольше тем хуже
	game.Board.Grid[0][0] = X
	game.Board.Grid[0][1] = X
	// Пустая позиция [0][2] даст быстрый выигрыш

	fastWin := service.(*gameService).minimax(game, X, 0, X)

	// Тот же выигрыш но на большей глубине должен быть хуже
	slowWin := service.(*gameService).minimax(game, X, 5, X)

	if fastWin <= slowWin {
		t.Errorf("Быстрый выигрыш (%d) должен быть лучше медленного (%d)", fastWin, slowWin)
	}
}

// TestGetNextMove_BasicMove тестирует базовый ход GetNextMove
func TestGetNextMove_BasicMove(t *testing.T) {
	repo := &mockRepository{}
	service := NewGameService(repo)
	game := createEmptyGame()

	nextMove := service.GetNextMove(game)

	// Проверяем что ход был сделан
	hasMove := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if nextMove.Board.Grid[i][j] != Void {
				hasMove = true
				break
			}
		}
		if hasMove {
			break
		}
	}

	if !hasMove {
		t.Error("GetNextMove должен вернуть игру с сделанным ходом")
	}

	// Проверяем что игра была сохранена в репозиторий
	if len(repo.savedGames) == 0 {
		t.Error("GetNextMove должен сохранить игру в репозиторий")
	}
}

// TestMinimax_OptimalPlay тестирует оптимальную игру
func TestMinimax_OptimalPlay(t *testing.T) {
	repo := &mockRepository{}
	service := NewGameService(repo)

	// Тестируем что PC выбирает выигрышный ход когда он доступен
	game := createEmptyGame()
	game.Board.Grid[0][0] = X // PC уже сделал ход

	// Если человек не блокирует, PC должен выиграть
	game.Board.Grid[1][0] = O // человек ходит
	game.Board.Grid[2][0] = X // PC продолжает

	// Теперь если PC ходит в [0][0], он должен увидеть потенциальный выигрыш
	result := service.(*gameService).minimax(game, X, 0, X)

	// Должны получить положительный скор так как у PC преимущество
	if result <= 0 {
		t.Errorf("Ожидали положительный скор для преимущественной позиции PC, получили %d", result)
	}
}
