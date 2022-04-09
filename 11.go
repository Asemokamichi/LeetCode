package main

import "fmt"

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	max := 0

	for i := 0; i < len(height); i++ {
		square := 0
		for j := len(height) - 1; j >= 0; j-- {
			if height[j] >= height[i] && j > i {
				square = height[i] * (j - i)
				break
			}

		}
		if max < square {
			max = square
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		square := 0
		for j := 0; j < len(height); j++ {
			if height[j] > height[i] && i > j {
				square = height[i] * (i - j)
				break
			}
		}
		if max < square {
			max = square
		}
	}
	return max
}
