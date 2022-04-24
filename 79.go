package main

import "fmt"

//79. Word Search

func main() {
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	board := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	//board := [][]byte{{'a', 'a'}}
	word := "AAB"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			k := board
			if k[i][j] == word[0] {
				if checkSearch(k, word[1:], i, j) {
					return true
				}
			}
		}
	}
	return false
}

func checkSearch(board [][]byte, word string, row, col int) bool {
	if len(word) == 0 {
		return true
	}
	symbol := board[row][col]
	board[row][col] = ' '

	if col != 0 && board[row][col-1] == word[0] {
		if checkSearch(board, word[1:], row, col-1) {
			return true
		}
	}

	if col != len(board[row])-1 && board[row][col+1] == word[0] {
		if checkSearch(board, word[1:], row, col+1) {
			return true
		}
	}

	if row != 0 && board[row-1][col] == word[0] {
		if checkSearch(board, word[1:], row-1, col) {
			return true
		}
	}

	if row != len(board)-1 && board[row+1][col] == word[0] {
		if checkSearch(board, word[1:], row+1, col) {
			return true
		}
	}
	board[row][col] = symbol
	return false
}
