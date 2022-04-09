package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {
	arr := unicale(nums)
	sort(arr)
	if len(arr)-1 >= 2 {
		return arr[2]
	}
	return arr[0]
}

func sort(arr []int) {
	for {
		flags := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] < arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flags = false
			}
		}
		if flags {
			return
		}
	}
}

func unicale(arr []int) []int {
	k := []int{}
	for _, w := range arr {
		flags := true
		for _, q := range k {
			if w == q {
				flags = false
			}
		}
		if flags {
			k = append(k, w)
		}
	}
	return k
}
