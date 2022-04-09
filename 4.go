package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := []int{}
	arr = append(nums1, nums2...)
	sort(arr)
	if len(arr)%2 != 0 {
		return float64(arr[len(arr)/2])
	} else {
		sum := (float64(arr[len(arr)/2-1]) + float64(arr[len(arr)/2])) / 2
		return sum
	}
}

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
