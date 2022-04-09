package main

import "fmt"

func main() {
	columnTitle := "AZ"
	fmt.Println(titleToNumber(columnTitle))
}

func titleToNumber(columnTitle string) int {
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	num := 0
	for _, w := range columnTitle {
		num = num*26 + index(alphabet, w) + 1
	}
	return num
}

func index(s []rune, toFind rune) int {
	for i := 0; i < len(s); i++ {
		if s[i] == toFind {
			return i
		}
	}
	return -1
}
