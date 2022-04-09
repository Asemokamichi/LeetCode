package main

import "fmt"

func main() {
	num1 := "1"
	num2 := "9"
	fmt.Println(num1, num2)
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	var len_arr int = len(num1)
	if len_arr < len(num2) {
		len_arr = len(num2)
	}
	digits := convert(num1, len_arr)
	digits2 := convert(num2, len_arr)
	n := 0
	for i := 0; i < len(digits); i++ {
		if digits[i]+digits2[i]+n < 10 {
			digits[i] = digits[i] + digits2[i] + n
			n = 0
		} else if i != len_arr-1 {
			digits[i] = digits2[i] + digits[i] + n - 10
			n = 1
		} else {
			digits[i] = digits2[i] + digits[i] + n - 10
			digits = append(digits, 1)
			break
		}
	}
	var s string
	for i := len(digits) - 1; i >= 0; i-- {
		s += string(rune(digits[i] + 48))
	}
	return s
}

func convert(num string, len_arr int) []int {
	k := []rune(num)
	nums := make([]int, len_arr)
	for i := len(k) - 1; i >= 0; i-- {
		nums[len(k)-1-i] = int(k[i] - 48)
	}
	return nums
}
