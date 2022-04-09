//125. Valid Palindrome
//03/30/2022 23:59	Accepted	4 ms	6.7 MB	golang

package main

import "fmt"

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
}
func isPalindrome(s string) bool {
	var k []rune
	for _, w := range s {
		if w >= 'a' && w <= 'z' {
			k = append(k, w)
		} else if w >= 'A' && w <= 'Z' {
			k = append(k, w+32)
		} else if w >= '0' && w <= '9' {
			k = append(k, w)
		}
	}
	x := len(k)
	for i := len(k) - 1; i >= 0; i-- {
		k = append(k, k[i])
	}
	if string(k[:x]) == string(k[x:]) {
		return true
	}
	return false
}
