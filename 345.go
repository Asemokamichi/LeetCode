package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	arr := []int{}
	for i, w := range s {
		for _, q := range vowels {
			if w == q {
				arr = append(arr, i)
				break
			}
		}
	}
	k := []rune(s)
	x := len(arr)
	for i := 0; i < (x-1)/2; i++ {
		k[arr[i]], k[arr[x-1-i]] = k[arr[x-1-i]], k[arr[i]]
	}
	return string(k)
}
