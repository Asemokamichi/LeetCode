package main

import "fmt"

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		return atoi(tokens[0])
	}
	for i, w := range tokens {
		if w == "+" || w == "-" || w == "/" || w == "*" || w == "%" {
			a := atoi(tokens[i-2])
			b := atoi(tokens[i-1])
			sum := operator(a, b, w)
			k := tokens[:i-2]
			k = append(k, atoi_str(sum))
			k = append(k, tokens[i+1:]...)
			return evalRPN(k)
		}
	}
	return 0
}

func atoi(s string) int {
	n := 1
	if rune(s[0]) == '-' {
		n = -1
		s = s[1:]
	}
	num := 0
	for _, w := range s {
		num = num*10 + int(w-48)
	}
	return num * n
}

func atoi_str(n int) string {
	if n == 0 {
		return "0"
	}
	k := []rune{}
	if n < 0 {
		k = append(k, '-')
		n = n * -1
	}
	bnm := []rune{}
	for n > 0 {
		bnm = append(bnm, rune(n%10+48))
		n = n / 10
	}
	for i := len(bnm) - 1; i >= 0; i-- {
		k = append(k, bnm[i])
	}
	return string(k)

}

func operator(a, b int, s string) int {
	switch rune(s[0]) {
	case '+':
		return a + b
	case '-':
		return a - b
	case '/':
		return a / b
	case '*':
		return a * b
	case '%':
		return a % b
	}
	return 0
}
