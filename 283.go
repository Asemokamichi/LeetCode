package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	fmt.Println(nums1)
}

func moveZeroes(nums []int) {
	index := 0
	for i, w := range nums {
		if w != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}

}
