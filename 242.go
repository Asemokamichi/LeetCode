package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
func isAnagram(s string, t string) bool {
	a, x := sort(s)
	b, y := sort(t)
	if a != b {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func sort(s string) (string, []int) {
	k := []rune{rune(s[0])}
	for _, w := range s {
		flags := true
		for _, q := range k {
			if w == q {
				flags = false
			}
		}
		if flags {
			k = append(k, w)
		}
	}

	for i := 1; i < len(k); i++ {
		for j := 1; j < len(k); j++ {
			if k[j-1] > k[j] {
				k[j-1], k[j] = k[j], k[j-1]
			}
		}
	}

	bnm := []int{}
	for _, q := range k {
		count := 0
		for _, w := range s {
			if q == w {
				count++
			}
		}
		bnm = append(bnm, count)
	}
	return string(k), bnm
}
