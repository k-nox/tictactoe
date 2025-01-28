package tictactoe

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	board   [][]string
	player1 player
	player2 player
	scanner bufio.Scanner
}

type player struct {
	name string
	mark string
}

type move struct {
	row  int
	col  int
	mark string
}

func New(r io.Reader) Game {
	game := Game{
		scanner: *bufio.NewScanner(r),
		board: [][]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
		player1: player{
			mark: "X",
		},
		player2: player{
			mark: "O",
		},
	}

	game.player1.name = game.playerInput("Player 1, please enter your name: ", func(s string) bool {
		if len(s) < 1 {
			fmt.Print("Please enter a name at least 1 character long: ")
			return false
		}
		return true
	})

	game.player2.name = game.playerInput("Player 2, please enter your name: ", func(s string) bool {
		if len(s) < 1 {
			fmt.Print("Please enter a name at least 1 character long: ")
			return false
		}
		if s == game.player1.name {
			fmt.Print("Please choose a different name than player 1: ")
			return false
		}

		return true
	})

	return game
}

func (g Game) String() string {
	return "    1   2   3\n" +
		"  ┌───┬───┬───┐\n" +
		fmt.Sprintf("A │ %s │ %s │ %s │\n", g.board[0][0], g.board[0][1], g.board[0][2]) +
		"  ├───┼───┼───┤\n" +
		fmt.Sprintf("B │ %s │ %s │ %s │\n", g.board[1][0], g.board[1][1], g.board[1][2]) +
		"  ├───┼───┼───┤\n" +
		fmt.Sprintf("C │ %s │ %s │ %s │\n", g.board[2][0], g.board[2][1], g.board[2][2]) +
		"  └───┴───┴───┘"
}

func (g Game) Play() {
	// game loop: until there is a winner OR all cells have been played, keep taking turns
	fmt.Println(g)
	g.player1.takeTurn(g)
	fmt.Println(g)
	g.player1.takeTurn(g)
	fmt.Println(g)
	g.player1.takeTurn(g)
	fmt.Println(g)
}

func (p player) takeTurn(g Game) {
	fmt.Printf("%s's turn, pick a cell: ", p.name)
	var (
		row int
		col int
	)
	g.playerInput(fmt.Sprintf("%s's turn, pick a cell: ", p.name), func(s string) bool {
		if len(s) != 2 {
			fmt.Printf("Please enter two characters: ")
			return false
		}

		rawMove := strings.ToLower(s)
		var (
			rawRow string
			rawCol string
		)
		if unicode.IsDigit(rune(rawMove[0])) {
			rawCol = string(rawMove[0])
			rawRow = string(rawMove[1])
		} else {
			rawRow = string(rawMove[0])
			rawCol = string(rawMove[1])
		}

		if rawRow != "a" && rawRow != "b" && rawRow != "c" {
			fmt.Print("Please indicate row with a, b, or c: ")
			return false
		}

		if rawCol != "1" && rawCol != "2" && rawCol != "3" {
			fmt.Print("Please indicate column with 1, 2, or 3: ")
			return false
		}

		switch rawRow {
		case "a":
			row = 0
		case "b":
			row = 1
		case "c":
			row = 2
		}

		col, _ = strconv.Atoi(rawCol)
		if g.board[row][col-1] != " " {
			fmt.Print("That cell has already been played, please pick another one: ")
			return false
		}
		return true
	})

	g.board[row][col-1] = p.mark
}

// func (g Game) isWon() bool {
//
// }

func (g Game) playerInput(prompt string, validator func(s string) bool) string {
	fmt.Print(prompt)
	g.scanner.Scan()
	response := g.scanner.Text()
	for !validator(response) {
		g.scanner.Scan()
		response = g.scanner.Text()
	}
	return response
}
