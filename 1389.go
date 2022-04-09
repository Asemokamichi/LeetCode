package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	fmt.Println(createTargetArray(nums, index))
}

func createTargetArray(nums []int, index []int) []int {
	arr := []int{}
	for i, w := range nums {
		arr = add(arr, w, index[i])
	}
	return arr
}

func add(nums []int, num, index int) []int {
	arr := make([]int, len(nums)+1)
	k := 0
	for i := 0; i <= len(nums); i++ {
		if i == index {
			arr[i] = num
			k = 1
		} else {
			arr[i] = nums[i-k]
		}

	}
	return arr
}
