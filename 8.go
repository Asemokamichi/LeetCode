package main

import "fmt"

func main() {
	s := "  -458"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	for i, w := range s {
		if w != ' ' {
			s = s[i:]
			break
		}
	}
	isnegative := 1
	if s[0] == '-' {
		isnegative = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	num := 0
	for _, w := range s {
		if 2147483647 < num*isnegative {
			return 2147483647
		} else if num*isnegative < -2147483648 {
			return -2147483648
		}
		if '0' <= w && w <= '9' {
			num = num*10 + int(w-'0')
		} else {
			return num * isnegative
		}
	}

	return num * isnegative
}
