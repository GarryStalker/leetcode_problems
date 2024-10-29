package main

func main() {

}

func majorityElement(nums []int) int {
	majorityMap := make(map[int]int)
	max := 0
	res := 0
	for _, elem := range nums {
		_, ok := majorityMap[elem]
		if !ok {
			majorityMap[elem] = 1
		} else {
			majorityMap[elem]++
		}
		if majorityMap[elem] > max {
			max = majorityMap[elem]
			res = elem
		}
	}
	return res
}

func majorityElement2(nums []int) int {
	majority := nums[0]
	count := 1

	for _, elem := range nums {
		if elem == majority {
			count++
		} else {
			count--
		}
		if count == 0 {
			majority = elem
			count = 1
		}
	}

	return majority
}
