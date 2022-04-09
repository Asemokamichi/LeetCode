package main

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	var bnm []rune
	for n > 0 {
		bnm = append(bnm, rune(n%10+48))
		n = n / 10
	}
	var arr []rune
	for i := len(bnm) - 1; i >= 0; i-- {
		arr = append(arr, bnm[i])
	}

	if string(bnm) == string(arr) {
		return true
	}
	return false
}
