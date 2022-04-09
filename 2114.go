package main

import "fmt"

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}

func mostWordsFound(sentences []string) int {
	max := 0
	for _, w := range sentences {
		if split(w) > max {
			max = split(w)
		}
	}
	return max
}

func split(s string) int {
	str := ""
	arr := []string{}
	for i, w := range s {
		if w != ' ' {
			str += string(w)
		}
		if w == ' ' || i == len(s)-1 {
			arr = append(arr, str)
			str = ""
		}
	}
	return len(arr)
}
