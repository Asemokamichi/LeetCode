package main

func fizzBuzz(n int) []string {
	arr := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, atoi_str(i))
		}
	}
	return arr
}

func atoi_str(num int) string {
	k := []rune{}
	for num > 0 {
		k = append(k, rune(num%10+48))
		num = num / 10
	}
	x := len(k)
	for i := len(k) - 1; i >= 0; i-- {
		k = append(k, k[i])
	}
	return string(k[x:])
}
