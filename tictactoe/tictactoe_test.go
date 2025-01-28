package tictactoe

import "testing"

func TestGameIsWon(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]string
		lastMove move
		expect   bool
	}{
		{
			name: "should detect horizontal win",
			board: [][]string{
				{" ", "O", " "},
				{"X", "X", "X"},
				{" ", " ", "O"},
			},
			lastMove: move{
				row:  1,
				col:  2,
				mark: "X",
			},
			expect: true,
		},
		{
			name: "should detect vertical win",
			board: [][]string{
				{"X", " ", " "},
				{"X", "O", " "},
				{"X", " ", "O"},
			},
			lastMove: move{
				row:  1,
				col:  0,
				mark: "X",
			},
			expect: true,
		},
		{
			name: "should detect diagnoal win",
			board: [][]string{
				{"X", " ", " "},
				{"O", "X", " "},
				{"O", " ", "X"},
			},
			lastMove: move{
				row:  2,
				col:  2,
				mark: "X",
			},
			expect: true,
		},
		{
			name: "should detect anti-diagonal win",
			board: [][]string{
				{" ", " ", "X"},
				{" ", "X", "O"},
				{"X", "O", " "},
			},
			lastMove: move{
				row:  2,
				col:  0,
				mark: "X",
			},
			expect: true,
		},
		{
			name: "should return false if there is no winner",
			board: [][]string{
				{" ", " ", "O"},
				{" ", "X", "O"},
				{"X", "O", " "},
			},
			lastMove: move{
				row:  2,
				col:  0,
				mark: "X",
			},
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := Game{
				board: test.board,
			}

			isWon := game.isWon(test.lastMove)
			if isWon != test.expect {
				t.Errorf("expected isWon to be %t but got %t", test.expect, isWon)
			}
		})
	}
}
