package main

func countSegments(s string) int {
	str := ""
	arr := []string{}
	for i, w := range s {
		if w != ' ' {
			str += string(w)
		}
		if w == ' ' || i == len(s)-1 {
			if str != "" {
				arr = append(arr, str)
				str = ""
			}
		}
	}
	return len(arr)

}
