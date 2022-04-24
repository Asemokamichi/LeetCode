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
	fmt.Println(isValidSudoku(arr))

}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if !check(board, i, j) {
				return false
			}
		}
	}
	return true
}

func check(board [][]byte, row, col int) bool {
	for i := 0; i <= 8; i++ {
		if board[row][col] == board[row][i] && col != i || board[row][col] == board[i][col] && row != i {
			return false
		}
	}

	firstCol := col - col%3
	firstRow := row - row%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+firstRow][j+firstCol] == board[row][col] && row != i+firstRow && col != j+firstCol {
				return false
			}
		}
	}
	return true
}
