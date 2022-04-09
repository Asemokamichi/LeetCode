package main

func romanToInt(s string) int {
	var num int = 0
	k := []rune(s)
	for i := 0; i < len(k); i++ {
		if i != len(k)-1 && roman_numerals(k[i]) < roman_numerals(k[i+1]) {
			num = num + roman_numerals(k[i+1]) - roman_numerals(k[i])
			i = i + 1
		} else {
			num = num + roman_numerals(k[i])
		}
	}
	return num
}

func roman_numerals(roman_num rune) int {
	switch roman_num {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
