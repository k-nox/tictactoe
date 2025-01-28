package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board [][]string

func new() board {
	return board{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
}

func (b board) move(cell string, mark string) {
	rowLabel := strings.ToLower(string(cell[0]))
	var row int8
	switch rowLabel {
	case "a":
		row = 0
	case "b":
		row = 1
	case "c":
		row = 2
	}
	col, err := strconv.Atoi(string(cell[1]))
	if err != nil {
		panic(err)
	}
	col -= 1

	b[row][col] = mark
}

func (b board) print() {
	fmt.Println("    1   2   3")
	fmt.Println("  ┌───┬───┬───┐")
	fmt.Printf("A │ %s │ %s │ %s │\n", b[0][0], b[0][1], b[0][2])
	fmt.Println("  ├───┼───┼───┤")
	fmt.Printf("B │ %s │ %s │ %s │\n", b[1][0], b[1][1], b[1][2])
	fmt.Println("  ├───┼───┼───┤")
	fmt.Printf("C │ %s │ %s │ %s │\n", b[2][0], b[2][1], b[2][2])
	fmt.Println("  └───┴───┴───┘")
}

func main() {
	fmt.Print("Player 1 name: ")
	reader := bufio.NewReader(os.Stdin)
	p1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	p1 = strings.TrimSpace(p1)

	fmt.Print("Player 2 name: ")
	p2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	p2 = strings.TrimSpace(p2)

	game := new()
	game.print()

	fmt.Printf("%s's turn, pick a spot: ", p1)
	cell, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	game.move(cell, "X")
	game.print()
}
