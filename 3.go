package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	k := []rune(s)
	for i := len(k); i > 0; i-- {
		for j := 0; j <= len(k)-i; j++ {
			if unicale(k[j : j+i]) {
				return len(k[j : j+i])
			}
		}
	}
	return 0
}

func unicale(k []rune) bool {
	if len(k) > 256 {
		return false
	}
	arr := []rune{k[0]}
	for _, w := range k {
		flags := true
		for _, q := range arr {
			if w == q {
				flags = false
				break
			}
		}
		if flags {
			arr = append(arr, w)
		}
	}
	if len(arr) == len(k) {
		return true
	}
	return false
}
