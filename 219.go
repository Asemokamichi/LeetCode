package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums1, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i-j <= k {
				return true
			}
		}
	}
	return false
}
