package main

import "fmt"

func main() {
	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(arr))
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxSubArray(nums []int) int {
	var s int = 0
	var best int = -100000
	for i := 0; i < len(nums); i++ {
		if s+nums[i] < nums[i] {
			s = 0
		}
		s = s + nums[i]
		best = max(s, best)

	}
	return best
}
