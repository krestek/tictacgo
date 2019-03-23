package main

import (
	"fmt"

	"./board"
)

func main() {
	b := board.New(3)
	var player rune = 'X'
	b.PrintBoard()
	for {
		fmt.Printf("Player %c\n", player)
		print("Row: ")
		var x, y int
		_, _ = fmt.Scan(&x)
		print("Column: ")
		_, _ = fmt.Scan(&y)
		isWon, err := b.Set(x-1, y-1, player)
		if err != nil {
			b.PrintBoard()
			println("Please enter a valid position.")
			continue
		}
		if isWon {
			break
		}
		if player == 'X' {
			player = 'O'
		} else {
			player = 'X'
		}
		b.PrintBoard()
	}

	b.PrintBoard()
	fmt.Printf("We have a winner! Player %c!", player)
}
