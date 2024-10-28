package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	nums2 := []int{3}
	m := 5
	n := 1
	//MergeSortedArray(nums1, m, nums2, n)
	//MergeSortedArray2(nums1, m, nums2, n)
	MergeSortedArray3(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < m+n; i++ {
		if m == 0 {
			nums1[i] = nums2[i]
			continue
		}
		if i > m-1 {
			nums1[i] = nums2[i-m]
		}

	}

	sort.Ints(nums1)
}

func MergeSortedArray2(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func MergeSortedArray3(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}
