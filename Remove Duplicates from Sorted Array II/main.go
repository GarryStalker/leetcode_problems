package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(removeDuplicatesII(nums))
}

func removeDuplicatesII(nums []int) int {
	left, counter := 1, 1
	for right := 1; right < len(nums); right++ {
		if nums[right] == nums[right-1] {
			counter++
		} else {
			counter = 1
		}

		if counter <= 2 {
			nums[left] = nums[right]
			left++
		}
	}
	fmt.Println(nums)

	return left
}

/* var left int

for right := 0; right < len(nums); right++ {
	if nums[left] != nums[right] {
		nums[left+1] = nums[right]
		left++
	}
}
fmt.Println(nums)
return len(nums[:left+1]) */
