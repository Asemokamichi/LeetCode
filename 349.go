package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
func intersection(nums1 []int, nums2 []int) []int {
	a := unicale(nums1)
	b := unicale(nums2)
	arr := []int{}
	for _, w := range a {
		for _, q := range b {
			if w == q {
				arr = append(arr, w)
			}
		}
	}
	return arr
}

func unicale(nums []int) []int {
	arr := []int{}
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
	return arr
}
