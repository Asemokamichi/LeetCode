package main

import "fmt"

func main() {
	board := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	arr := [][]byte{}
	for i := 0; i < 9; i++ {
		rev := []byte{}
		for j := 0; j < 9; j++ {
			rev = append(rev, board[i][j][0])
		}
		arr = append(arr, rev)
	}
	solveSudoku(arr)
	fmt.Println(arr)

}

func solveSudoku(board [][]byte) {
	sudoku(&board, 0, 0)
}

func sudoku(board *[][]byte, row, col int) bool {
	if row == 8 && col == 9 {
		return true
	} else if col == 9 {
		row++
		col = 0
	}

	if (*board)[row][col] != '.' {
		return sudoku(board, row, col+1)
	}
	for i := '1'; i <= '9'; i++ {
		if check(board, row, col, byte(i)) {
			(*board)[row][col] = byte(i)
			if sudoku(board, row, col+1) {
				return true
			}
		}
		(*board)[row][col] = '.'
	}
	return false
}

func check(board *[][]byte, row, col int, num byte) bool {
	for i := 0; i <= 8; i++ {
		if num == (*board)[row][i] || num == (*board)[i][col] {
			return false
		}
	}

	firstCol := col - col%3
	firstRow := row - row%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[i+firstRow][j+firstCol] == num {
				return false
			}
		}
	}
	return true
}
