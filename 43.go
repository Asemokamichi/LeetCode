package main

import "fmt"

func main() {
	num1 := "25"
	num2 := "36"
	fmt.Println(multiply(num1, num2))
}

func multiply(num1 string, num2 string) string {
	if len(num2) == 0 {
		return "0"
	}
	a := atoi(num1)
	b := runInt(rune(num2[len(num2)-1]))
	fmt.Println(a)
	fmt.Println(b)
	n := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i]*b+n < 10 {
			a[i] = a[i]*b + n
			n = 0
		} else if i != 0 {
			k := a[i]*b + n
			n = k / 10
			a[i] = k % 10
		} else {
			k := a[i]*b + n
			a[i] = k % 10
			a = append([]int{k / 10}, a...)
			break
		}
	}
	return addStrings(atoi_str(a), multiply(num1, num2[:len(num2)-1])+"0")
}

func atoi_str(nums []int) string {
	k := []rune{}
	for _, w := range nums {
		k = append(k, rune(w+48))
	}
	return string(k)
}

func atoi(s string) []int {
	nums := []int{}
	for _, w := range s {
		nums = append(nums, int(w-48))
	}
	return nums
}

func runInt(num rune) int {
	return int(num - 48)
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
