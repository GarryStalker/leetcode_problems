package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3999))
}

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string

	for num > 0 {
		for i := range arabic {
			if num >= arabic[i] {
				result += roman[i]
				num -= arabic[i]
				break
			}
		}
	}

	return result
}
