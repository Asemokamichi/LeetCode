package main

import "fmt"

func main() {
	arr := []int{6, 5, 5}
	fmt.Println(majorityElement(arr))
}

func majorityElement(nums []int) int {
	rev := comb(nums)
	for _, w := range rev {
		count := 0
		for j := 0; j < len(nums); j++ {
			if w == nums[j] {
				count++
			}
		}
		if count >= len(nums)-len(nums)/2 {
			return w
		}
	}
	return 0
}

func comb(nums []int) []int {
	arr := []int{nums[0]}
	for _, w := range nums {
		flags := true
		for _, q := range arr {
			if w == q {
				flags = false
			}
		}
		if flags {
			arr = append(arr, w)
		}
	}
	fmt.Println(arr)
	return arr
}
