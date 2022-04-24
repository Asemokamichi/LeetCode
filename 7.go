package main

import "fmt"

func main() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	num := 0
	k := 1
	if x < 0 {
		k = -1
		x = -x
	}
	for x > 0 {
		num = num*10 + x%10
		x /= 10
	}

	if -2147483648 > num || num > 2147483647 {
		return 0
	}

	return num * k
}
