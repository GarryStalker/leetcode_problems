package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	jumps := nums[0]

	for i := 1; i < len(nums); i++ {
		if jumps == 0 {
			return false
		}
		jumps = max(jumps-1, nums[i])
	}

	return true
}
