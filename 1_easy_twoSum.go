package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		if idx, ok := hash[num]; ok {
			return []int{idx, i}
		}
		hash[target-num] = i
	}
	return []int{}
}
