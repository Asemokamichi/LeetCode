package main

import "fmt"

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))

}

func isIsomorphic(s string, t string) bool {
	a := unique(s)
	b := unique(t)
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
