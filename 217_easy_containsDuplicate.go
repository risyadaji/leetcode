package main

// ref: https://leetcode.com/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}

		m[nums[i]] = struct{}{}
	}

	return false
}
