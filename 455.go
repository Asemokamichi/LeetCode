package main

import "fmt"

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort(g)
	sort(s)
	i, j, count := 0, 0, 0
	fmt.Println(g, s)
	for i < len(g) && j < len(s) {
		fmt.Println(g[i], s[j])
		if g[i] <= s[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}

func sort(arr []int) {
	for {
		flags := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flags = false
			}
		}
		if flags {
			return
		}
	}
}
