package main

import "fmt"

func main() {
	height := []int{0, 1, 0}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	water_area := 0
	for i := 0; i < len(height); {
		square := 0
		for j := i + 1; j < len(height); j++ {
			if height[j] >= height[i] {
				if j-i-1 == 0 {
					height = height[j:]
					i = 0
					j = i
					continue
				}
				square += height[i] * (j - i - 1)
				height = height[j:]
				i = -1
				break
			} else {
				square -= height[j]
			}
		}
		if square > 0 {
			water_area += square
			i++
		} else if i != len(height)-1 {
			height = reverse(height)
			i = 0
		} else {
			i++
		}

	}
	return water_area
}

func reverse(nums []int) []int {
	x := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums = append(nums, nums[i])
	}
	return nums[x:]
}
