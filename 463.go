package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sum += 4 - check_horizontally(grid, i, j) - check_vertically(grid, i, j)
			}
		}
	}
	return sum
}

func check(grid [][]int, i, j int) int {
	if grid[i][j] == 1 {
		return 1
	} else {
		return 0
	}
}

func check_horizontally(grid [][]int, i, j int) int {
	if j == 0 && j == len(grid[i])-1 {
		return 0
	} else if j == 0 {
		return check(grid, i, j+1)
	} else if j == len(grid[i])-1 {
		return check(grid, i, j-1)
	} else {
		return check(grid, i, j-1) + check(grid, i, j+1)
	}
}

func check_vertically(grid [][]int, i, j int) int {
	if i == 0 && i == len(grid)-1 {
		return 0
	} else if i == 0 {
		return check(grid, i+1, j)
	} else if i == len(grid)-1 {
		return check(grid, i-1, j)
	} else {
		return check(grid, i-1, j) + check(grid, i+1, j)
	}
}
