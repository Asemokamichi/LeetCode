//88. Merge Sorted Array
//04/07/2022 14:35	Accepted	4 ms	2.5 MB	golang
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	var arr []int
	arr = append(nums1, nums2...)
	arr = sort(arr)
	fmt.Println(arr)
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
