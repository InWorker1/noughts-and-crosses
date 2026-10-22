package game

import "testing"

func TestCheckMove(t *testing.T) {
	old := [][]int{{X, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	cases := []struct {
		name   string
		grid   [][]int
		symbol int
		want   bool
	}{
		{"ok", [][]int{{X, O, 0}, {0, 0, 0}, {0, 0, 0}}, O, true},
		{"чужой символ", [][]int{{X, X, 0}, {0, 0, 0}, {0, 0, 0}}, O, false},
		{"два хода", [][]int{{X, O, O}, {0, 0, 0}, {0, 0, 0}}, O, false},
		{"без хода", [][]int{{X, 0, 0}, {0, 0, 0}, {0, 0, 0}}, O, false},
		{"перезапись", [][]int{{O, 0, 0}, {0, 0, 0}, {0, 0, 0}}, O, false},
		{"не 3x3", [][]int{{X, O}}, O, false},
		{"мусор в клетке", [][]int{{X, 5, 0}, {0, 0, 0}, {0, 0, 0}}, O, false},
	}
	for _, c := range cases {
		if got := checkMove(old, c.grid, c.symbol); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}
