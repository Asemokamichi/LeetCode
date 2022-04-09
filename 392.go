package main

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}
	for _, w := range s {
		flags := false
		for j, q := range t {
			if w == q {
				t = t[j+1:]
				flags = true
				break
			}
		}
		if !flags {
			return false
		}
	}
	return true
}
