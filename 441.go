package main

func arrangeCoins(n int) int {
	for i := 1; n > 0; i++ {
		if n-i == 0 {
			return i
		} else if n-i < 0 {
			return i - 1
		}
		n = n - i
	}
	return 0
}
