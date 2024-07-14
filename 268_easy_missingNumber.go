package main

// ref: https://leetcode.com/problems/missing-number/?envType=daily-question&envId=2024-07-13

// first solution
func missingNumber(nums []int) int {
	m := make(map[int]struct{})

	for i := range nums {
		m[nums[i]] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 1
}
