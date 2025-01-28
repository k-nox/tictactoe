package main

import (
	"os"

	"github.com/k-nox/tictactoe/tictactoe"
)

func main() {
	game := tictactoe.New(os.Stdin)
	game.Play()
}
