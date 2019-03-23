package board

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Board [][]rune

func New(size int) Board {
	b := make(Board, size)
	for i := 0; i < len(b); i++ {
		b[i] = make([]rune, size)
	}
	return b
}

func (b *Board) Set(row int, col int, char rune) (bool, error) {
	size := len(*b)
	if row < 0 || col < 0 || row > size || col > size || (*b)[row][col] != 0 {
		return false, errors.New("Invalid position")
	}
	(*b)[row][col] = char
	diaRes := 0
	reverseDiaRes := 0
	for r := 0; r < size; r++ {
		colRes := 0
		rowRes := 0
		if (*b)[r][r] == char {
			diaRes++
		}
		if (*b)[2-r][r] == char {
			reverseDiaRes++
		}
		for c := 0; c < size; c++ {
			if (*b)[r][c] == char {
				colRes++
			}
			if (*b)[c][r] == char {
				rowRes++
			}
		}
		if colRes == size || rowRes == size || diaRes == size || reverseDiaRes == size {
			return true, nil
		}
	}
	return false, nil
}

func clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	_ = c.Run()
}

func (b *Board) PrintBoard() {
	clear()
	for r := 0; r < len(*b); r++ {
		for c := 0; c < len((*b)[r]); c++ {
			fmt.Printf("[%c]", (*b)[r][c])
		}
		println()
	}
}
