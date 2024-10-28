package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	var left int

	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	fmt.Println(nums)
	return len(nums[:left+1])
}
