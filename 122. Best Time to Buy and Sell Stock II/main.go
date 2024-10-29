package main

import "fmt"

func main() {
	prices := []int{1, 4, 2}
	fmt.Println(maxProfitII(prices))
}

func maxProfitII(prices []int) int {
	stock, profit := 0, 0
	haveStock := false

	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[stock] && haveStock == true {
			profit += prices[i] - prices[stock]
			stock = 0
			haveStock = false
			fmt.Printf("sell on day: %d\nprofit: %d\n\n", i, profit)
		}
		if i != len(prices)-1 && prices[i] < prices[i+1] && haveStock == false {
			stock = i
			haveStock = true
			fmt.Printf("buy on day: %d\nprofit: %d\n\n", i, profit)
		}
	}
	return profit
}

func maxProfitII_2(prices []int) int {
	profit := 0

	for i := 0; i < len(prices)-1; i++ {
		if prof := prices[i+1] - prices[i]; prof > 0 {
			profit += prof
		}
	}
	return profit
}
