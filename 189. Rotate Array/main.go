package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateArray(nums, 3)
	fmt.Println(nums)
}

func rotateArray(nums []int, k int) {
	k = -k % len(nums)
	if k < 0 {
		k += len(nums)
	}
	rotated := make([]int, len(nums))
	copy(rotated, nums[k:])
	copy(rotated[len(nums)-k:], nums[:k])
	copy(nums, rotated)
}

func rotateArray2(nums []int, k int) {
	size := len(nums)
	k %= size

	var count, start = 0, 0

	for count < size {
		nextIdx := start
		temp := nums[start]
		for {
			nextIdx = (nextIdx + k) % size
			temp, nums[nextIdx] = nums[nextIdx], temp
			count++
			if start == nextIdx {
				break
			}
		}
		start++
	}

}

//0  1  2  3  4  5  6
//1, 2, 3, 4, 5, 6, 7
//5, 6, 7, 1, 2, 3, 4
