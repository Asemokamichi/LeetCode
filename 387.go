package main

func firstUniqChar(s string) int {
	arr := unicale(s)
	for _, w := range arr {
		count, i := 0, 0
		for j, q := range s {
			if w == q {
				count++
				i = j
			}
		}
		if count == 1 {
			return i
		}
	}
	return -1
}

func unicale(s string) []rune {
	arr := []rune{}
	for _, w := range s {
		flags := true
		for _, q := range arr {
			if w == q {
				flags = false
			}
		}
		if flags {
			arr = append(arr, w)
		}
	}
	return arr
}
