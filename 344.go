package main

import "fmt"

func main() {
	s := []string{"h", "e", "l", "l", "o"}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []string) {
	x := len(s)
	for i := 0; i < (x-1)/2; i++ {
		s[i], s[x-1-i] = s[x-1-i], s[i]
	}
}
