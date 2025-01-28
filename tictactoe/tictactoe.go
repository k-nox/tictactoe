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
	name  string
	mark  string
	color func(s string) string
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
			color: magenta,
		},
		player2: player{
			color: cyan,
		},
	}

	game.player1.name = game.playerInput(game.player1.color("Player 1, please enter your name: "), func(s string) bool {
		if len(s) < 1 {
			fmt.Print(red("Please enter a name at least 1 character long: "))
			return false
		}
		return true
	})

	game.player1.mark = game.getMarkForPlayer("X", game.player1)

	game.player2.name = game.playerInput(game.player2.color("Player 2, please enter your name: "), func(s string) bool {
		if len(s) < 1 {
			fmt.Print(red("Please enter a name at least 1 character long: "))
			return false
		}
		if s == game.player1.name {
			fmt.Print(red("Please choose a different name than player 1: "))
			return false
		}

		return true
	})

	game.player2.mark = game.getMarkForPlayer("O", game.player2)

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
	g.instructions()
	maxTurns := len(g.board) * len(g.board[0])
	isPlayer1 := true
	for i := 0; i < maxTurns; i++ {
		activePlayer := g.player1
		if !isPlayer1 {
			activePlayer = g.player2
		}

		move := activePlayer.takeTurn(g)
		g.board[move.row][move.col] = move.mark
		fmt.Println(g)
		if g.isWon(move) {
			fmt.Printf(green("Congrats to the winner %s!\n"), activePlayer.name)
			return
		}

		isPlayer1 = !isPlayer1
	}
	fmt.Println(blue("It's a draw!"))
}

func (p player) takeTurn(g Game) move {
	var (
		row int
		col int
	)
	g.playerInput(p.color(fmt.Sprintf("%s's turn, pick a cell: ", p.name)), func(s string) bool {
		if len(s) != 2 {
			fmt.Print(red("Please enter two characters: "))
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
			fmt.Print(red("Please indicate row with a, b, or c: "))
			return false
		}

		if rawCol != "1" && rawCol != "2" && rawCol != "3" {
			fmt.Print(red("Please indicate column with 1, 2, or 3: "))
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
			fmt.Print(red("That cell has already been played, please pick another one: "))
			return false
		}
		return true
	})

	return move{
		row:  row,
		col:  col - 1,
		mark: p.mark,
	}
}

func (g Game) instructions() {
	fmt.Println(yellow("Welcome to TicTacToe ") + g.player1.color(g.player1.name) + yellow(" and ") + g.player2.color(g.player2.name) + yellow("!"))
	fmt.Println(yellow("This is what the board looks like:"))
	fmt.Println(g)
	fmt.Println(yellow("To take your turn, you'll input the cell you want to play."))
	fmt.Println(yellow("For example, to play the top left cell, you can input 'a1', 'A1', '1a', or '1A'."))
	fmt.Println(yellow("Have fun!"))
}

func (g Game) isWon(lastMove move) bool {
	boardSize := len(g.board)
	// check horizontal
	for col := 0; col < boardSize; col++ {
		if g.board[lastMove.row][col] != lastMove.mark {
			break
		}
		if col == boardSize-1 {
			return true
		}
	}

	// check vertical
	for row := 0; row < boardSize; row++ {
		if g.board[row][lastMove.col] != lastMove.mark {
			break
		}
		if row == boardSize-1 {
			return true
		}
	}

	// check diagnoal
	if lastMove.row == lastMove.col {
		for i := 0; i < boardSize; i++ {
			if g.board[i][i] != lastMove.mark {
				break
			}
			if i == boardSize-1 {
				return true
			}
		}
	}

	// check anti-diagonal
	if lastMove.col+lastMove.row == boardSize-1 {
		for i := 0; i < boardSize; i++ {
			if g.board[i][boardSize-1-i] != lastMove.mark {
				break
			}
			if i == boardSize-1 {
				return true
			}
		}
	}
	return false
}

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

func (g Game) getMarkForPlayer(defaultMark string, player player) string {
	mark := g.playerInput(player.color(fmt.Sprintf("Hit enter to use %s or type a custom mark now: ", defaultMark)), func(s string) bool {
		if len(s) > 1 {
			fmt.Print(red("Please input only one character: "))
			return false
		}
		return true
	})

	if mark == "" {
		mark = defaultMark
	}

	return player.color(mark)
}

const reset = "\033[0m"

func red(s string) string {
	return "\033[31m" + s + reset
}

func green(s string) string {
	return "\033[32m" + s + reset
}

func yellow(s string) string {
	return "\033[33m" + s + reset
}

func blue(s string) string {
	return "\033[34m" + s + reset
}

func magenta(s string) string {
	return "\033[35m" + s + reset
}

func cyan(s string) string {
	return "\033[36m" + s + reset
}
