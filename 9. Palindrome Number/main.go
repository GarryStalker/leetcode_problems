package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

//Given an integer x, return true if x is a palindrome , and false otherwise.
//convert number int to []int and read from start and end and compare values
//if number is negative, return false
//if number%10 == 0, return true

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return true
	}
	numberSlice := intToIntSlice(x)
	fmt.Println(len(numberSlice))
	for i := range numberSlice {
		if numberSlice[i] == numberSlice[len(numberSlice)-1-i] {
			continue
		}
		return false
	}
	return true
}

func intToIntSlice(x int) []int {
	var slice []int

	for x > 0 {
		slice = append(slice, x%10)
		x /= 10
	}

	return slice
}
