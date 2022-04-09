package main

import "fmt"

func main() {
	arr := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(arr))
}

func longestCommonPrefix(strs []string) string {
	min_obj, index := min_arr(strs)
	k := comb_str(min_obj)
	for _, w := range k {
		flags := true
		for j, q := range strs {
			if j != index && !check_toFind(q, w) {
				flags = false
				break
			}
		}
		if flags {
			return w
		}
	}
	return ""
}

func comb_str(s string) []string {
	arr := []string{}
	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			arr = append(arr, s[j:j+i])
		}
	}
	return arr
}

func min_arr(arr []string) (string, int) {
	min := arr[0]
	index_min := 0
	for i, w := range arr {
		if len(w) < len(min) {
			min = w
			index_min = i
		}
	}
	return min, index_min
}

func check_toFind(str, toFind string) bool {
	if str[0:len(toFind)] == toFind {
		return true
	}
	return false
}
