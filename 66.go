package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
			return digits
		} else if i != 0 {
			digits[i] = 0
		} else {
			digits[i] = 0
			digits = append([]int{1}, digits...)
			return digits
		}
	}
	return digits
}
