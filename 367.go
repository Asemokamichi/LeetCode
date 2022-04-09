package main

import "fmt"

func main() {
	num := 15
	fmt.Println(isPerfectSquare(num))
}

func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	} else if num < 3 {
		return true
	}
	for i := 2; i <= num/2; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
