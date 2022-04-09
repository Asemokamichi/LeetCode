package main

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	for i := 1; i <= x; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
	return 1
}
