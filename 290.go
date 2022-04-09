package main

import "fmt"

func main() {
	pattern := "abba"
	s := "dog cat cat fish"
	fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
	a := unique(pattern)
	b := unique_str(s)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func unique(s string) []string {
	k := []rune(s)
	arr := []rune{k[0]}
	for _, w := range s {
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

	index := []string{}
	for _, w := range arr {
		rev := []rune{}
		for j, q := range k {
			if w == q {
				rev = append(rev, rune(j+48))
			}
		}
		index = append(index, string(rev))
	}
	return index
}

func unique_str(s string) []string {
	k := split(s)
	arr := []string{k[0]}
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

	index := []string{}
	for _, w := range arr {
		rev := []rune{}
		for j, q := range k {
			if w == q {
				rev = append(rev, rune(j+48))
			}
		}
		index = append(index, string(rev))
	}
	return index
}

func split(s string) []string {
	var str string
	var arr []string
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
