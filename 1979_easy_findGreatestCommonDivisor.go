package main

// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
func findGCD(nums []int) int {
	// constraint max: 1000, min: 1
	min, max := 1001, 0

	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	for i := min; i >= 1; i-- {
		if min%i == 0 && max%i == 0 {
			return i
		}
	}
	return 1
}
