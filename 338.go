package main

func countBits(n int) []int {
	bnm := []int{}
	for i := 0; i <= n; i++ {
		bnm = append(bnm, sum(i))
	}
	return bnm
}

func sum(n int) int {
	sum := 0
	for n > 0 {
		if n%2 == 1 {
			sum++
		}
		n = n / 2
	}
	return sum
}
