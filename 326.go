package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(25))
}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	} else if n%3 != 0 || n < 1 {
		return false
	}
	return isPowerOfThree(n / 3)
}
