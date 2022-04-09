package main

func convertToBase7(num int) string {
	k := []rune{}
	s := ""
	if num < 0 {
		s += "-"
		num *= -1
	} else if num == 0 {
		return "0"
	}
	for num > 0 {
		k = append(k, rune(num%7+48))
		num = num / 7
	}
	for i := len(k) - 1; i >= 0; i-- {
		s += string(k[i])
	}
	return s
}
