package main

func lengthOfLastWord(s string) int {
	var arr []string
	var str string
	for i, w := range s {
		if w != ' ' {
			str = str + string(w)
		}
		if w == ' ' || i == len(s)-1 {
			arr = append(arr, str)
			str = ""
		}
	}
	var bnm []string
	for _, w := range arr {
		if w != "" {
			bnm = append(bnm, w)
		}
	}
	return len(bnm[len(bnm)-1])
}
