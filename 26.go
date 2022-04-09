package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 2, 2, 4}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}

func removeDuplicates(arr []int) int {
	for {
		flags := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flags = false
			} else if arr[j-1] == arr[j] {
				arr[len(arr)-1], arr[j-1] = arr[j-1], arr[len(arr)-1]
				arr = arr[:len(arr)-1]
				j--
				flags = false
			}
		}
		if flags {
			break
		}
	}
	return len(arr)
}
