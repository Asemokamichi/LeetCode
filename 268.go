//268. Missing Number

package main

import "fmt"

func main() {
	arr := []int{3, 0, 1}
	fmt.Println(missingNumber(arr))
}
func missingNumber(nums []int) int {
	nums = sort(nums)
	for i := 0; i <= len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return 0
}

func sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
