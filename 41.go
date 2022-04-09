package main

import "fmt"

func main() {
	nums := []int{-5}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	sort(nums)
	k := 0
	num := 1
	for _, w := range nums {
		if w > 0 && w != k && w == num {
			k = w
			num++
		} else if w > 0 && w != k && w != num {
			return num
		}
	}
	if nums[len(nums)-1] >= 0 {
		return nums[len(nums)-1] + 1
	}
	return num
}

func sort(arr []int) {
	for {
		flags := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flags = false
			}
		}
		if flags {
			return
		}
	}
}
