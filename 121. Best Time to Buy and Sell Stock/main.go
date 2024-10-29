package main

func main() {

}

func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for _, price := range prices {
		if min > price {
			min = price
			continue
		}
		if price-min > profit {
			profit = price - min
		}
	}
	return profit
}
