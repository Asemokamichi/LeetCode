package main

import "fmt"

func main() {
	s := "Hello how are you Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	arr := split(s)[:k]
	s = ""
	for i, w := range arr {
		s += w
		if i != len(arr)-1 {
			s += " "
		}
	}
	return s
}

func split(s string) []string {
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
	return arr
}
