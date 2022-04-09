package main

import "fmt"

func main() {
	nums1 := 2
	fmt.Println(isPowerOfTwo(nums1))
}

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if n%2 != 0 || n < 1 {
		return false
	}
	return isPowerOfTwo(n / 2)
}
