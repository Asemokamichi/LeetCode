//136. Single Number
package main

import "fmt"

func main() {
	arr := []int{1}
	fmt.Println(singleNumber(arr))
}
func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		flags := true
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				flags = false
				break
			}
		}
		if flags {
			return nums[i]
		}
	}
	return 0
}
