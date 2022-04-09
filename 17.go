package main

import "fmt"

func main() {
	digits := "2"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		k := []string{}
		return k
	}
	digit := atoi(digits)
	s := Phone_Number(digit[0])
	for i := 1; i < len(digit); i++ {
		s = comb(s, Phone_Number(digit[i]))
	}
	return s

}

func atoi(str string) []int {
	bnm := []int{}
	for _, w := range str {
		bnm = append(bnm, int(w-48))
	}
	//x := len(bnm)
	//for i := x - 1; i >= 0; i-- {
	//	bnm = append(bnm, bnm[i])
	//}
	return bnm
}

func comb(s []string, bnm []string) []string {
	k := []string{}
	for _, w := range s {
		for _, q := range bnm {
			k = append(k, w+q)
		}
	}
	return k
}

func Phone_Number(n int) []string {
	k := []string{}
	switch n {
	case 2:
		k = []string{"a", "b", "c"}
		return k
	case 3:
		k = []string{"d", "e", "f"}
		return k
	case 4:
		k = []string{"g", "h", "i"}
		return k
	case 5:
		k = []string{"j", "k", "l"}
		return k
	case 6:
		k = []string{"m", "n", "o"}
		return k
	case 7:
		k = []string{"p", "q", "r", "s"}
		return k
	case 8:
		k = []string{"t", "u", "v"}
		return k
	case 9:
		k = []string{"w", "x", "y", "z"}
		return k
	}
	return k
}
