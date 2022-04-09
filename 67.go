package main

func addBinary(a string, b string) string {
	max_len := 0
	if len(a) > len(b) {
		max_len = len(a)
	} else {
		max_len = len(b)
	}
	var arr []int
	x := atoibase(a, max_len)
	y := atoibase(b, max_len)
	num := 0
	for i := 0; i < max_len; i++ {
		if x[i]+y[i]+num == 2 {
			arr = append(arr, 0)
			num = 1
		} else if x[i]+y[i]+num == 3 {
			arr = append(arr, 1)
			num = 1
		} else {
			arr = append(arr, x[i]+y[i]+num)
			num = 0
		}
		if i == max_len-1 && num > 0 {
			arr = append(arr, 1)
		}
	}
	str := ""
	for i := len(arr) - 1; i >= 0; i-- {
		str = str + string(rune(arr[i]+48))
	}
	return str

}

func atoibase(s string, n int) []int {
	num := make([]int, n)
	for i, j := len(s)-1, 0; i >= 0 && j < len(s); i-- {
		num[j] = int(rune(s[i]) - 48)
		j++
	}
	return num
}
