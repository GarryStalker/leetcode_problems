package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 3, 2, 2, 3}
	val := 3
	fmt.Println(RemoveElement2(nums, val))

	/* a := []string{"A", "B", "C", "D", "E"}
	i := 0

	// Удалить элемент по индексу i из a.

	// 1. Копировать последний элемент в индекс i.
	a[i] = a[len(a)-1]

	// 2. Удалить последний элемент (записать нулевое значение).
	a[len(a)-1] = ""

	// 3. Усечь срез.
	a = a[:len(a)-1]

	fmt.Println(a) // [A B E D] */
}

func RemoveElement(nums []int, val int) (k int) {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			k++
		}
	}
	return k
}

func RemoveElement2(nums []int, val int) int {
	var j int

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	fmt.Println(nums)
	return j
}
