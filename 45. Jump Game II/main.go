package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	var j, maxDistance, currPos int

	for i := 0; i < len(nums)-1; i++ {
		maxDistance = max(maxDistance, nums[i]+i)

		if currPos == i {
			j++
			currPos = maxDistance
		}
	}

	return j
}
