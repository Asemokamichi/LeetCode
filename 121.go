//121. Best Time to Buy and Sell Stock
//04/09/2022 15:22	Accepted	112 ms	8 MB	golang

package main

import "fmt"

func main() {
	arr := []int{2, 4, 1}
	fmt.Println(maxProfit(arr))
}

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				break
			} else if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
