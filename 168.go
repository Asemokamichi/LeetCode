package main

import "fmt"

func main() {
	columnNumber := 52
	fmt.Println(convertToTitle(columnNumber))
}

func convertToTitle(columnNumber int) string {
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if columnNumber == 0 {
		return ""
	} else if columnNumber < 27 {
		return alphabet[columnNumber-1]
	}
	k := ""
	if columnNumber%26 > 0 && columnNumber%26 < 26 {
		k = alphabet[columnNumber%26-1]
	}
	if columnNumber%26 == 0 {
		return convertToTitle(columnNumber/26-1) + "Z"
	}
	return convertToTitle(columnNumber/26) + k
}
