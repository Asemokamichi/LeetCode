package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 2, 2, 4}
	val := 1
	fmt.Println(removeElement(arr, val))
	fmt.Println(arr)
}
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return len(nums)
}
